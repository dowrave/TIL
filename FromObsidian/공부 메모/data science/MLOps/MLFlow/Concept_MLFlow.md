- MLFlow의 구성요소 4가지
	- Tracking
	- Projects
	- Models
	- Model Registry
- 각각의 요소를 따로 진행할 수도, 함께 진행할 수도 있음
- "기존 라이브러리와 함께 작동하기" + "기존 코드에 대한 대부분의 사항을 관례적으로 결정함" + "기존 코드베이스에 통합하기 위해 최소한의 변경만이 필요하도록 설계됨" + 다른 사람이 재현하고 재사용할 수 있도록 만듦


## WorkFlow
1. 데이터 전처리 
2. 알고리즘으로 실험 
3. 모델을 프로덕션 시스템에 배포.
4. 성능 모니터링 
5. 새로운 데이터 재교육 
6. 대체 모델과 비교

- 위 과정에서 어려운 것들
	- 실험 추적하기 : 특정 결과를 얻기 위해 사용된 데이터, 코드, 매개변수 파악하는 일
	- 코드 복제 : 전체 환경(라이브러리 종속성)을 포착해야 동일한 결과를 얻을 수 있음
	- 모델 패키징 & 배포 : 표준 방법이 없기 때문에 모델 & 생성 코드 & 매개 변수 간의 연결은 종종 없어질 때가 있음
	- 모델 관리 저장소가 없음 : 모델 라이프사이클을 공동으로 관리할 수 있는 중앙집중식 저장소가 없음 -> 모델 관리에 어려움이 있음
- MLFlow는 위 과정들을 해결해줄 수 있다

## Components

### 1. Tracking
- 코드 실행 시 매개변수, 코드 버전, Metric, Artifact를 기록하고 결과를 시각화하기 위한 API이자 UI
- 모든 환경에서 Tracking을 사용해 결과를 로컬 파일이나 서버에 저장한 다음 여러 실행을 비교할 수 있다.

### 2. Projects
- 재사용 가능한 데이터 과학 코드를 패키징하기 위한 표준 형식
- 각 프로젝트는 코드나 Git 저장소가 있는 디렉토리이며, 종속성과 코드 실행 방법을 지정하기 위해 기술자 파일 & 단순한 규약을 사용한다.
	- 예를 들면 프로젝트 내에는 Conda 환경을 지정하기 위한 `conda.yaml` 파일이 포함될 수 있다.
	- 프로젝트에서 MLFlow Tracking API를 쓰면 MLFlow는 프로젝트 버전(Git commit 등)과 매개 변수를 자동으로 기억한다. 

### 3. Models
- ML 모델을 패키징하는 관례 & 배포에 도움이 되는 도구를 제공함
- 각 모델은 임의 파일 + 모델이 사용될 수 있는 `Flavor` 를 나열하는 설명자 파일을 포함하는 디렉토리로 저장됨
	- ex) 텐서플로우 모델은 텐서플로우 DAG로 로드하거나 파이썬 함수로 로드될 수 있음
- MLFlow는 다양한 플랫폼에 많은 일반적인 모델 유형을 배포하기 위한 도구를 제공함
	- ex) 파이썬 함수 Flavor를 지원하는 모든 모은 도커 기반 REST 서버, Azure ML, AWS SageMaker 등의 클라우드 플랫폼이나, Batch & Streaming 추론을 위한 Apache Spark의 사용자 정의 함수로 배포할 수 있다.
	- 추적 API를 사용하여 MLFlow 모델을 출력하는 경우 MLFlow는 어떤 프로젝트와 모델에서 왔는지도 자동으로 기억함

### 4. Registry
- MLflow 모델의 전체 라이프사이클을 공동으로 관리할 수 있는 중앙집중식 모델 저장소, API 집합, UI를 제공함
- 모델 계보(MLFlow 실험 & 실험으로 생성된 모델), 모델 버전, 스테이지 전환, 주석 등을 제공함

## Referencing Artifact
```
-   `/Users/me/path/to/local/model`
    
-   `relative/path/to/local/model`

-   `<scheme>/<scheme-dependent-path>`. For example:
    
    -   `s3://my_bucket/path/to/model`
        
    -   `hdfs://<host>:<port>/<path>`
        
    -   `runs:/<mlflow_run_id>/run-relative/path/to/model`
        
    -   `models:/<model_name>/<model_version>`
        
    -   `models:/<model_name>/<stage>`
        
    -   `mlflow-artifacts:/path/to/model` when running the tracking server in `--serve-artifacts` proxy mode.
```
- 대충 이렇게 있다고 하는데 이건 해보면서 알아야 할 듯
- Tracking API : `mlflow.log_artifacts("<mlflow_run_id>", "/path/to/artifact")`
- Model API : `mlflow.pytorch.log_model("runs:/<mlflow_run_id>/run-relative/path/to/model", registered_model_name="mymodel")`


## Scalability and Big Data
- 스케일링(대규모 데이터, 출력 파일, 실험 등)도 가능함
- 분산 처리(`Apache Spark`)로 개별 클러스터에서 MLFlow 실행 가능
- 하이퍼파라미터 튜닝도 가능
- Model Registry를 통해 대규모 조직에 전체 라이프사이클을 공동으로 관리할 수 있는 중앙 허브를 제공함

