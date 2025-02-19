
- 정의 : 클라우드 컴퓨팅에서 사용되는 **가상 서버, 가상 머신**
- 각 인스턴스는 운영 체제를 포함, 독립적으로 운영되며 원하는 SW와 앱을 설치하고 구성할 수 있다.
- 인스턴스 유형에 따라 CPU, 메모리, 스토리지 같은 컴퓨팅 자원의 양이 결정된다.
- 실제 **물리적 서버의 자원(CPU, 메모리, 스토리지)을 분할하여 여러 가상 머신(인스턴스)이 동시에 동작**하도록 한다. 각 인스턴스는 별도의 가상 하드웨어로 구성되며, 독립된 컴퓨터처럼 작동한다.
	- 각 **인스턴스 별로 특화된 기능**이 있더라도, **CPU, 메모리, 스토리지는 모두 갖고 있다.**

### 인스턴스와 컨테이너
- 컨테이너 : **앱과 그 종속성을 패키징**하여, 소프트웨어가 일관된 환경에서 실행될 수 있도록 하는 경량화된 가상화 기술.
	- **Docker는 OS 수준의 가상화**를 제공하며, 1개의 서버에서 여러 컨테이너를 실행할 수 있게 한다.

### AWS 인스턴스의 종류

#### Amazon EC2(Elastic Compute Cloud)
- 가상 서버를 제공한다.
- 필요한 CPU, 메모리, 스토리지를 선택, 원하는 운영 체제와 앱을 설치하고 구동할 수 있다.
- 웹 서버, DB 서버, 배치 처리, 대규모 연산 작업 등에 활용된다.

#### Amazon S3(Simple Storage Service)
- 인터넷으로 어디서나 데이터를 저장하고 검색할 수 있는 오브젝트 스토리지 서버.
- 정적 웹 호스팅, 백업 저장소, 빅데이터 분석을 위한 데이터 저장소 등
- 사용한 만큼 비용을 지불하는 구조

#### Amazon Lambda
- 서버를 직접 관리하지 않고 코드를 실행할 수 있는 서버리스 컴퓨팅 서비스
- 사용자는 코드를 작성하고 실행 조건만 설정하면, AWS가 자동으로 필요한 컴퓨팅 리소스를 할당하고 관리한다.
- 데이터 처리, 실시간 파일 처리, 웹 요청 처리 등에 사용된다.
- 코드 실행 시간에만 비용이 청구된다.


### 현황에 맞는 인스턴스

#### 프론트엔드(React)
- **Amazon S3**
	- 정적 웹 호스팅에 적합하다. React 빌드 + S3 버킷에 업로드해서 웹 사이트로 제공할 수 있다.

#### 백엔드 (Django)
- **Amazon EC2**
	- `t2.micro` 인스턴스는 AWS 프리 티어에서 제공, 소규모나 초기 단계의 웹앱에 적합하다.
	- `t2.small, t3.small` 등은 트래픽이 증가하거나 더 많은 리소스가 필요할 때 이용할 수 있다.

#### 데이터베이스 (MySQL)
- **Amazon RDS**
	- 프리 티어에서는 `db.t2.micro` 인스턴스를 사용할 수 있다.
	- 자동 백업, 패치 관리, 스케일링 등 다양한 관리 기능을 제공한다.

#### 데이터 수집
- **AWS Lambda**
	- 주기적인 데이터 수집 작업에 적합하다.
	- 스크립트를 Lambda 함수로 작성하면 지정된 시간 간격으로 자동으로 실행된다.

