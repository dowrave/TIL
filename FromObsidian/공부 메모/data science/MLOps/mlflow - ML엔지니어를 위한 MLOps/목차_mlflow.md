
[ML 엔지니어를 위한 MLOps](https://mlops-for-mle.github.io/tutorial/) 를 따라간 내용
- `좀 더 효율적인 학습을 위해서는 해설을 먼저 보기 보다는 스펙 명세서를 직접 구현한 뒤 설명하는 내용들을 읽는 것을 권장합니다.`

## 1. Database
- 리눅스, 도커, 데이터 저장할 DB 세팅
- 데이터 DB에 삽입 & `자동 삽입`
- 이미지 생성`Dockerfile`
- 컨테이너 간 네트워크 형성 `Docker Compose`

![[Pasted image 20221229180948.png]]
1. [[환경 준비]]
2. [[테이블 설정 & 데이터 삽입(~반복)]]
3. [[dockerfile 설정 & 컨테이너 간 네트워크 연결]]
4. [[docker-compose 파일을 이용한 네트워크 연결]]

- 여기부터는  `1_make_db_insert_data/docker-compose.yaml` 파일로 컨테이너(`postgresql, data-generator` 를 띄운 상태에서 진행함)
```sh
# 경로 /home/dowrave 에서 실행
docker compose up -d
```

## 2. Model Development
- 데이터 불러오기 & 모델 학습 및 저장 & 저장된 모델 불러와서 재현성 검사
![[Pasted image 20221229181007.png]]
1. [[기본 모델 생성]]
2. [[모델 파이프라인]]
3. [[DB에서 불러온 데이터로 모델 학습]]

## 3. Model Registry
- `2. Model Development` 파트의 모델을 저장 & 관리
- `MLFlow` 서버 구축 & 서버에 모델 저장
![[Pasted image 20221230151634.png]]
1. [[MLflow Backend Store 구축]]
2. [[MLflow 레지스트리에 모델 저장]]
3. [[MLflow 레지스트리에서 모델 불러오기]]

- `docker compose up -d`를 2번 해야 함 : 1번 폴더, 3번 폴더

## 4. Model Deployment
[[모델 배포]]

---

## 5. FastAPI
- API 학습
- `FastAPI` : 파이썬을 이용해 API를 만들 수 있는 웹 프레임워크
![[Pasted image 20221230181547.png]]
- 공부 예정..


## 6. API Serving
- `request driven` 방식을 통해 학습한 모델을 사용함
- `FastAPI`를 이용해 데이터를 입력받고 모델 예측값을 반환하는 REST API를 구현함
![[Pasted image 20221230181645.png]]

## 7. Kafka
- `Streaming Serving`을 구현하기 위한 실시간 데이터 파이프라인을 `Kafka`를 이용하여 구축
- 2가지 가정
	- `Source DB` : 데이터가 계속 쌓이는 외부 DB
	- `Target DB` : 외부 데이터를 처리한 뒤 쌓이는 내부 DB
	- `Source DB`는 1장의 `PostgreSQL DB`를 사용
![[Pasted image 20221230181800.png]]
## 8. Stream
- `Event-Driven` 방식 구현
- 7장의 `Kafka System` + `Target DB`에 더해, `Consumer`를 통해 토픽에서 데이터를 읽어 API 서버의 입력으로 전달, 추론 결과를 Target DB로 전달하는 Data Subscriber를 구현함
- `Grafana`를 이용해 원본 데이터와 예측 결과값을 실시간으로 시각화하는 대시보드를 만들고, Stream Serving이 잘 되고 있는지를 확인함.
![[Pasted image 20221230182009.png]]



---
### 재시작 시 발생하는 문제
1. 재시작 시 `Cannot connect to the Docker daemon at unix:///var/run/docker.sock.`가 뜬다
>`service docker start` 
>`service docker status` 로 도커 실행 & 상태 확인
>> 그래도 안되면 `sudo update-alternatives --config iptables`에서 설정을 바꾸자
>>> 완전 처음이 0, 설정 변경 후 1
>>> 재시작 시 1로 유지된 상태에서 도커만 꺼진 상태였음
