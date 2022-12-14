#EDA
[EDA의 개념과 데이터분석 잘하는 법](https://jalynne-kim.medium.com/%EB%8D%B0%EC%9D%B4%ED%84%B0%EB%B6%84%EC%84%9D-%EA%B8%B0%EC%B4%88-eda%EC%9D%98-%EA%B0%9C%EB%85%90%EA%B3%BC-%EB%8D%B0%EC%9D%B4%ED%84%B0%EB%B6%84%EC%84%9D-%EC%9E%98-%ED%95%98%EB%8A%94-%EB%B2%95-a3cac2cc5ebc)
`EDA(Exploratory Data Analysis)탐색적 데이터 분석`
	- **데이터를 분석하고 결과를 내는 과정**에서  **지속적으로 해당 데이터에 대한 탐색과 이해가 기본적으로 갖춰져야 한다.**
	- 이게 왜 중요하냐면, 사전에 이해도가 떨어지면 나중에 수행한 코드들이 다 휴지조각이 되기 때문임

##### 어떤 과정을 거치는가?
- 처음 `Raw Data`를 접할 때부터 데이터를 잘 이해하고 파악한 뒤, 어떤 결과를 만들지 **여러 Feature로 필터링**을 해본다
- 데이터 분석을 통한 결과값을 출력하기 전에, 어떤 결과값을 낼지 **가설을 설정**하고 **기본적인 표나 그래프를 그리며 사전 검증**을 하는 과정

#### EDA 어케 잘함?
1. `Raw Data`의 `Description, Dictionary`를 통해 각 Column과 Row의 의미 이해하기
	- 잘 읽으라는 뜻. 캐글이라면 데이터셋에 대한 설명을 다 줄 거임

2. 결측치 처리 및 데이터 필터링 기술
	- NaN값 제거
	- 수치형 데이터와 범주형 데이터 처리

3. 누구나 이해하기 쉬운 시각화
	- 이 그래프를 통해 어떤 의미를 얻어야 하는지가 직관적으로 와닿아야 한다.
	- 화려할 필요가 없음.  중요한 건 **의미 전달**이다.

#### EDA가 쉽지 않은 이유
- **인지편향 : 무의식적으로 자신의 경험으로 먼저 판단을 내려버리는 것**
	- 인간의 본능임 : 모든 정보를 처리할 수 없기 때문에 `어림`해버리는 것.
	- 인지편향을 경계하고 타인의 이야기와 데이터에 대해 그대로 이해할 수 있는 능력을 길러야 한다.

