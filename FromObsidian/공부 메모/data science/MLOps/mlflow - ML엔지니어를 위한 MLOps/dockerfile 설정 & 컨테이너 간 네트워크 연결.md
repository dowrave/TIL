- 이전 : [[테이블 설정 & 데이터 삽입(~반복)]]

- `data_generator.py` (추가된 부분만 주목)
```python
import time  
from argparse import ArgumentParser  # 추가된 부분
  
import pandas as pd  
import psycopg2  
from sklearn.datasets import load_iris  
  
  
def get_data():  
X, y = load_iris(return_X_y=True, as_frame=True)  
df = pd.concat([X, y], axis="columns")  
rename_rule = {  
"sepal length (cm)": "sepal_length",  
"sepal width (cm)": "sepal_width",  
"petal length (cm)": "petal_length",  
"petal width (cm)": "petal_width",  
}  
df = df.rename(columns=rename_rule)  
return df  
  
  
def create_table(db_connect):  
create_table_query = """  
CREATE TABLE IF NOT EXISTS iris_data (  
id SERIAL PRIMARY KEY,  
timestamp timestamp,  
sepal_length float8,  
sepal_width float8,  
petal_length float8,  
petal_width float8,  
target int  
);"""  
print(create_table_query)  
with db_connect.cursor() as cur:  
cur.execute(create_table_query)  
db_connect.commit()  
  
  
def insert_data(db_connect, data):  
insert_row_query = f"""  
INSERT INTO iris_data  
(timestamp, sepal_length, sepal_width, petal_length, petal_width, target)  
VALUES (  
NOW(),  
{data.sepal_length},  
{data.sepal_width},  
{data.petal_length},  
{data.petal_width},  
{data.target}  
);  
"""  
print(insert_row_query)  
with db_connect.cursor() as cur:  
cur.execute(insert_row_query)  
db_connect.commit()  
  
  
def generate_data(db_connect, df):  
while True:  
insert_data(db_connect, df.sample(1).squeeze())  
time.sleep(1)  
  
  
if __name__ == "__main__":  
# 추가된 부분
parser = ArgumentParser()  
parser.add_argument("--db-host", dest="db_host", type=str, default="localhost")  
args = parser.parse_args()  
  
db_connect = psycopg2.connect(  
user="myuser",  
password="mypassword",  
host=args.db_host,  
port=5432,  
database="mydatabase",  
)  
create_table(db_connect)  
df = get_data()  
generate_data(db_connect, df)
```

> - 위 스크립트(`data.generator.py`)를 실행할 도커파일을 만든다.
		- Base image : `amd64/python:3.9-slim`
		- `postgresql-client` 설치
		- `scikit-learn`, `pandas`, `psycopg2-binary` 설치
		- `ENTRYPOINT`, `CMD`로 스크립트를 실행함

- 파일명 `dockerfile`
```dockerfile
FROM amd64/python:3.9-slim

RUN apt-get update && apt-get install -y \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/app

RUN pip install -U pip &&\
    pip install scikit-learn pandas psycopg2-binary

COPY data_generator.py data_generator.py

ENTRYPOINT ["python", "data_generator.py", "--db-host"]

# Change CMD to solve host finding error
CMD ["localhost"]
```
> - `BASE` : 베이스 이미지
	- `RUN` : 이미지를 만들 때 실행할 코드 지정
		- 1번째 : `postgresql-client` 설치
		- 2번째 : 컨테이너에서 파이썬 스크립트 실행 시 필요한 것들 설치
	- `WORKDIR` : 작업 디렉토리 지정 : 이후 커맨드들은 해당 디렉토리에서 진행됨
	- `COPY` : `COPY (로컬 파일/폴더) (이미지에 저장할 이름)` : 파일이나 폴더를 이미지에 복사함
	- `ENTRYPOINT` : 컨테이너 실행 시 시작할 프로세스
	- `CMD` : 컨테이너 실행 시 `ENTRYPOINT`에 전달할 `ARGUMENT`지정

- **도커 파일 빌드**
```sh
# 규칙
docker build [OPTIONS] PATH | URL | -
```
```sh
docker build -t data-generator .

# 확인 : (data-generator) 보이면 됨
docker images

# 실행
docker run data-generator
```

- 오류
```
psycopg2.OperationalError: connection to server at "localhost" (127.0.0.1), port 5432 failed: Connection refused  
Is the server running on that host and accepting TCP/IP connections?  
connection to server at "localhost" (::1), port 5432 failed: Cannot assign requested address  
Is the server running on that host and accepting TCP/IP connections?
```
> 확인할 사항
> 1. DB 컨테이너가 제대로 실행되고 있는가?
>> `docker ps`로 확인 -> `postgres:14.0`만 제대로 떠 있음

-  왜 로컬은 되는데 `data-generator` 실행 컨테이너에서는 접근이 되지 않을까?
	- 로컬에서 띄운 명령어는 `docker run ... -p 5432:5432 postgres:14.0`이 있었다
		- `-p`는 `DB 컨테이너`의 5432번 포트를 `localhost`의 5432번 포트에 연결하는 것이다.
		- 즉, **로컬호스트의 5432번 포트에 접근하는 것 = 컨테이너의 5432번 포트에 접근하는 것**이다(여기서 그렇게 설정을 했으니까)

![[Pasted image 20221229165101.png]]
- 지금 띄운 컨테이너는 `Data Generator`이다.
- 이 컨테이너에서 `localhost:5432`는 비어있는 포트로 보이기 때문에, DB를 찾지 못한다는 에러와 함께 종료된다.
- 따라서, DB 컨테이너와 `Data Generator` 간의 연결 설정을 해야 한다(`localhost:5432`)

- 네트워크 연결하기 : `docker network`
1) 네트워크 생성
```sh
docker network create my-network
```

2) 실행 중인 DB 컨테이너를 생성된 네트워크와 연결함
```sh
docker network connect my-network postgres-server
```

3) `data-generator`이미지로 `data-generator` 이름의 컨테이너를 생성하는데, 조건을 추가로 준다
```sh
docker run -d --name data-generator --network 'my-network' data-generator 'postgres-server'
```
- `-d` : 백그라운드 모드
- `--network` : 위에서 생성한 네트워크 이름 입력
	- 뒤에는 연결할 두 네트워크의 이름을 설정하는 듯 하며 `--network (네트워크이름) (지금컨테이너) (상대컨테이너)` 인 것 같은데 정확하지 않음
	- 또한 뒤의 `'postgres-server'` 값은 `dockerfile`의 CMD 값 `['localhost']` 에 들어가는 값임(이미 네트워크에는 `postgres-server` 값이 연결되어 있음)

- **`docker network`의 한계**
	- 컨테이너의 이름을 알아야 함
	- 연결이 필요한 컨테이너에 다른 컨테이너의 이름을 전달해야 함
	- 이름을 지정하는 방식이 있으나, 컨테이너가 예상치 못한 이유로 종료된 경우에도 이름은 점유되고 있기 때문에 다시 실행하기 위해서 종료된 컨테이너를 삭제해야 함
	- 컨테이너 실행 순서를 보장할 수 없음 : DB 서버 실행 전에 Data Generator 서버에서 데이터 삽입을 시도하면 해당 컨테이너는 오류와 함께 실행이 종료됨

- 컨테이너 간의 작업을 조율하는 것을 `Container Orchestration`이라고 하며, 이걸 하는게 `k8s`임
- 도커에서는 `Docker Compose`를 제공함

- 다음 : [[docker-compose 파일을 이용한 네트워크 연결]]