
[ML 엔지니어를 위한 MLOps](https://mlops-for-mle.github.io/tutorial/)
- `좀 더 효율적인 학습을 위해서는 해설을 먼저 보기 보다는 스펙 명세서를 직접 구현한 뒤 설명하는 내용들을 읽는 것을 권장합니다.`

#### 1. Database
![[Pasted image 20221229180948.png]]
1. [[환경 준비]]
2. [[테이블 설정 & 데이터 삽입(~반복)]]
3. [[dockerfile 설정 & 컨테이너 간 네트워크 연결]]
4. [[docker-compose 파일을 이용한 네트워크 연결]]

- 여기부터는 위에서 작성한 `docker-compose.yaml` 파일로 컨테이너(`postgresql, data-generator` 를 띄운 상태에서 진행함)
```sh
# 경로 /home/dowrave 에서 실행
docker compose up -d
```
#### 2. Model Development
![[Pasted image 20221229181007.png]]


------------
### 재시작 시 발생하는 문제
1. 재시작 시 `Cannot connect to the Docker daemon at unix:///var/run/docker.sock.`가 뜬다
>`service docker start` 
>`service docker status` 로 도커 실행 & 상태 확인
>> 그래도 안되면 `sudo update-alternatives --config iptables`에서 설정을 바꾸자
>>> 완전 처음이 0, 설정 변경 후 1
>>> 재시작 시 1로 유지된 상태에서 도커만 꺼진 상태였음
