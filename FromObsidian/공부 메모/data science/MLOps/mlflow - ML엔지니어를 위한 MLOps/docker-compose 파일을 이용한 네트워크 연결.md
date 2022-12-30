- 이전 : [[dockerfile 설정 & 컨테이너 간 네트워크 연결]]

## Docker-compose 파일 작성
- DB 컨테이너와 Data Generator 컨테이너를 함께 띄우기 위한 Docker Compose 파일 작성

- `docker compose`
>-   Postgres Server
    -   Service name : `postgres-server`
    -   Image : `postgres:14.0`
    -   Container name : `postgres-server`
    -   Environment
        -   POSTGRES_USER : `myuser`
        -   POSTGRES_PASSWORD : `mypassword`
        -   POSTGRES_DB : `mydatabase`
    -   Port forwarding : `5432:5432`
>-   Data Generator
    -   Service name : `data-generator`
    -   Image : `Dockerfile`
    -   Container name : `data-generator`
    -   Command : `["postgres-server"]`
>> 아마 postgres server 서비스가 먼저 생겨야 할 거임(데이터를 저장해야 하니까) -> 먼저 생겨야 할 서비스를 Compose 파일에 지정

- `dockerfile`은 1-5에서 사용한 것과 동일함
- `compose` 파일 작성 : `docker-compose`
```yaml
version: '3'
services:
  postgres-server:
    ...
  data-generator:
    ...
```
> version
>> `Compose` 파일의 버전
>>> 작성일 기준 최신 버전이 `3`이래요
>services
>> Compose에 묶일 서비스를 의미함. 1개의 서비스는 1개의 컨테이너
>> postgres-server, data-generator는 서비스 이름


- `postgres-server` 작성
```yaml
services:
  postgres-server:
    image: postgres:14.0
    container_name: postgres-server
    ports:
      - 5432:5432
    environment:
      POSTGRES_USER: myuser
      POSTGRES_PASSWORD: mypassword
      POSTGRES_DB: mydatabase
```
> `postgres-server`
>> 서비스 이름이자 컨테이너의 호스트 이름
> `image`
>> 사용할 컨테이너 이미지
>`ports`
>> 컨테이너에서 외부로 노출할 포트 포워딩
>> `host:container`로 사용되며 여러 개 지정 가능
>`environment`
>> 컨테이너를 실행할 때의 `-e`와 같은 역할

- `data-generator` 작성
```yaml
services:
  data-generator:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: data-generator
    depends_on:
      - postgres-server
    command: ["postgres-server"]
```
> `build`
>> `context` : 이미지 빌드를 위한 `Dockerfile`의 절대 / 상대 경로
>> `dockerfile` : 도커파일의 파일 이름
>`depends_on`
>> `compose`로 띄워지는 서비스 간의 **종속성 순서대로 서비스를 시작**할 때 사용함
>> `postgre`가 먼저, `data generator`가 나중이어야 하므로 **나중에 실행되는 컨테이너에 먼저 실행되어야 하는 컨테이너를 명시**한다.
>`command`
>> `Dockerfile`의 `cmd`를 덮어쓰는 옵션
>> 앞에서 했듯, `postgres server`의 호스트를 변경해야 함(`localhost` -> `postgres-server`)
>> 호스트 이름은 컨테이너의 이름으로 줘야 함

- 실행
```sh
docker compose up -d
```
> 실행 결과 `postgre-server`만 띄워짐(하드웨어 사양에 따라 결과는 다를 수도 있음)
>> `postgre-server`가 **준비되지 않은** 상태에서 `data generator`가 띄워져서 연결을 시도하기 때문에 `data generator`가 exit되는 문제가 발생함

- 이런 문제 해결을 위해 `docker compose healthcheck`가 있는데, `dokcer-compose.yaml` 파일에 `healthcheck`과 `condition`을 추가하면 됨
```yaml
version:

services:
  postgres-server:
    ...
    healthcheck:
      test: ['CMD', 'pg_isready', '-q', '-U', 'myuser', '-d', 'mydatabase']
      interval: 10s
      timeout: 5s
      retries: 5
  data-generator:
    ...
    depends_on:
      postgres-server:
        condition: service_healthy
    command: ['postgres-server']
```
> `test`
>> 테스트할 명령어 설정
>> `pg_isready`를 이용하여 DB가 준비상태인지 테스트하기 위한 커맨드가 본문 커맨드
>`interval`
>> `healthcheck` 간격을 설정함
>`timeout`
>> `healthcheck`의 `timeout` 설정
>`retries`
>> `Timeout`의 횟수를 설정함
>`condition`
>> 1. `healthcheck`를 사용하기 위해 `depends_on`의 파라미터로 `condition:service_healthy`를 넣어줌 
>> 2. `postgres server`의 `healthcheck`를 적용시키기 위해 `postgres-server`에 `condition`을 추가함

- 실행 확인
```sh
docker compose up -d
docker ps
```

#### Docker Compose Network
- 생성된 네트워크 확인
```sh
docker network ls
```
> `part1_default` 생성 확인 (나는 `dowrave_default`로 뜬 듯?)

- 해당 네트워크 확인
```sh
docker network inspect dowrave_default
```
	- 네트워크에 추가된 컨테이너 `postgres-server`, `data-generator` 확인
	- 이거의 의의는 내가 네트워크를 설정하지 않고 컨테이너를 편리하게 연결할 수 있다는 데 있음

- 네트워크의 이름은 특별히 지정하지 않으면 `{디렉토리명}_default`로 자동생성된다. 따라서 서로 같은 디렉토리를 사용하지 않는다면 네트워크 이름이 변할 수 있다.

- 실행 서비스 종료
```sh
docker compose down -v
```
> `-v` : 생성된 볼륨까지 삭제

- 네트워크 생성 : `services`와 같은 레벨에 `networks`를 입력
- `docker-compose.yaml`
```yaml
networks:
  default:
    name: mlops-network
```
> `default` : 서비스 전체의 기본 네트워크를 수정할 수 있음
>> `name` : 네트워크 이름 작성


- 다시 `docker compose up -d`로 `mlops-network`를 확인하자

- 데이터 확인 
	- (로컬) :  `psql`로 DB에 접속 & 데이터 삽입 확인
	- (컨테이너) 
		- `docker exec -it data-generator /bin/bash`
		- `PGPASSWORD=mypassword psql -h postgres-server -p 5432 -U myuser -d mydatabase`
		- `SELECT * FROM iris_data;`

- 다음 : [[기본 모델 생성]]