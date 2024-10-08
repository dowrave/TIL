## 1. 추출 필터
#### 데이터 연결 과정
- `라이브`
- `추출`
	- 데이터 원본 소스에서 일부만 쓰고 싶을 때 사용함
	- 만약 클라우드에 있는 데이터를 추출했을 경우, 해당 로컬 데이터를 따로 파일로 저장하는 작업이 추가로 이뤄진다. 
- 맨 처음에 데이터를 불러올 때 적용되며, 어떤 소스에서 가져오느냐에 따라 `라이브`는 뜨지 않을 수도 있음

## 2. 데이터 원본 필터
- 추출한 데이터에서 다시 필터링을 적용할 수 있음(데이터 원본 탭 우측 상단)

- 데이터 용량이 클수록 **추출 필터 + 데이터 원본 필터**에서 필터링을 잘 해야 성능을 개선할 수 있음

## 3. 컨텍스트 필터

## 4. 차원 필터
- **시트에서 `필터` 카드에 파란색 특성(차원)을 올리는 것**을 의미함
- `필터 편집`을 통해 어떻게 적용할 지 결정할 수 있음
- `필터 표시`를 통해 그래프 오른쪽에 필터를 띄워 더 쉽게 살펴볼 수 있음
	- **디폴트로 다중 선택**이 가능하게 되어 있는데, 오른쪽에 떠있는 필터 창은 단일값만을 띄울 수도 있는 등, 옵션을 줄 수 있다.

## 5. 측정값 필터
- **시트에서 `필터` 카드에 초록색 특성(측정값)을 올리는 것**을 의미
- 해당 특성의 집계값을 이용해서 화면을 필터링할 수 있음(집계값 이상, 이하 ...)
- 집계 값 이외에도 **최상단에 `모든 값`이 있는데, 이는 `Row - Level`에서 접근함을 의미함.

## 6. 테이블 필터


## 7. 숨기기
> 일단 예제를 이해해보자
> 1. 각 지역과 서브 카테고리 4개에 대한 합계(Sales)를 이용할 거임
> 2. 비율 계산은 합계(Sales)의 퀵 테이블 계산 - 구성 비율
> 3. 근데 여기서 구한 구성 비율은 전체에 대한 구성 비율이므로, 지역으로 구분하기 위해  `합계(Sales) - 테이블 계산 편집 - 특정 차원 - Region 체크 해제`
> 4. 이 상태에서 서브 카테고리를 필터에 올린 다음, 한 항목만 체크 해제함

- 위의 예제대로 따라가면 서브 카테고리 3개가 남게 되는데, 이 3개의 비율 합이 100%가 되도록 그래프가 구성됨
- 이는 **`차원 필터`가 먼저 적용된 다음에 `측정값 필터`가 적용되기 때문**임. 즉 3개의 서브 카테고리가 남은 상태에서 구성 비율을 따지기 때문에, 3개의 합이 100%이고 각각의 비율이 나타나게 되는것

- 그렇다면 3개의 서브 카테고리의 합이 100%가 되지 않게 하면서 체크 해제된 항목을 보이지 않게 하는 방법은 없을까? 
- 여기서 이용되는 게 **숨기기** 기능이다.
	1. 서브 카테고리 4개 전부를 띄우고, 필터에 서브 카테고리를 넣음
	2. 서브 카테고리 범례에서 숨기고 싶은 항목을 우클릭하면, 가장 아래에 `숨기기`가 있음
	- 이 결과 3개의 항목을 띄우면서 전체 합이 100%가 되지 않게 할 수 있다.

- **숨기기는 테이블 계산 이후에 진행되므로, 계산 결과가 뭉개지지 않는다.**

- 태블로는 원하는 화면을 만들기 위해 필터의 적용 순서를 생각해야 함