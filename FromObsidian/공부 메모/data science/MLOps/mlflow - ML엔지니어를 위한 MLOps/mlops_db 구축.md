## 1) 도커 설치 & DB 서버 생성

##### 우분투 설치
- MS store에서 우분투 설치함 끝

##### 도커 설치(공식 문서)
> docker desktop이랑 관계 없이 돌아감

- 이전 버전 제거
```sh
sudo apt-get remove docker docker-engine docker.io containerd runc
```

- https에서 apt를 실행하기 위한 세팅
```sh
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg lsb-release
```

- GPG 키 추가 
```sh
sudo mkdir -p /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
```

- 레포지토리 세팅 과정
```sh
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

- 도커 엔진 설치 & 실행 확인
```sh

sudo chmod a+r /etc/apt/keyrings/docker.gpg # gpg 에러 발생 시 
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

```sh
sudo docker run hello-world
```

> 에러 : `docker: Cannot connect to the Docker daemon at unix:///var/run/docker.sock.` 
>> 1. `systemctl start docker`로 확인해도 `Host is down`이 뜸
>> 2. `service docker start`도 `service docker status` 입력 시 `docker not running`이 뜬다.
>해결 
>> 1. `sudo update-alternatives --config iptables`
>> 2. `1`을 입력하면 `iptables-legacy`로 변경됨
>> 3. `sudo service docker start` & `sudo service docker status`로 확인


- [관리자 권한 없이 실행시키기](https://netmarble.engineering/docker-on-wsl2-without-docker-desktop/)



### postgresql

#### 설치

- 컨테이너 실행하기
```
-   Image : `postgres:14.0`
-   Container name : `postgres-server`
-   POSTGRES_USER : `myuser`
-   POSTGRES_PASSWORD : `mypassword`
-   POSTGRES_DB : `mydatabase`
-   Port forwarding : `5432:5432`
```

```sh
docker run --rm \ 
-p 5432:5432 \ 
--name postgres-server \ 
-e POSTGRES_USER=myuser \
-e POSTGRES_PASSWORD=mypassword \ 
-e POSTGRES_DB=mydatabase \ 
postgres:14.0
```
- 환경변수에 들어가는 string값들은 `''`을 붙일 필요가 없다.
- 이걸 띄운 터미널은 계속 돌아가고 있기 때문에 **별도의 우분투 탭을 띄워서 접근**해야 한다.


#### DB 서버 확인
- `psql`이라는 클라이언트 툴이 필요하다. 
- 가이드에서는 공식 홈페이지에서 버전을 참고해서 받으라고 한다.
- 무지성으로 `apt-get install psql` 을 치면 `apt-get install postgresql-client-common`을 설치하라고 뜨는데,  정확히는 접근할 `postgresql`의 버전에 맞는 `psql`이 필요하다. 우리는 `14.0` 버전의 이미지를 쓸 거니까..

- `psql` 설치
```sh
sudo apt-get install postgresql-client-14
```

- DB에 접근
```sh
PGPASSWORD=mypassword psql -h localhost -p 5432 -U myuser -d mydatabase

# 결과 (입력 가능한 창이 뜸)
mydatabase=# 
```

- `\du`로 `role name`과 `attributes`을 확인하자.
```sh
mydatabase=# \du
```


## 2) 테이블 생성
