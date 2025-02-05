#tableau_basic 
### 이중 축의 3가지 용도
1. 기본 사용법
2. 강조 포인트
3. 디자인 효과 또는 컨텍스트 부여

## 1. 이중 축의 기본 사용법
- 만드는 방법
	- 방법 1. 2번째 측정 값을 화면 오른쪽으로 드래그해서 점선이 보이면 놓기
	- 방법 2. 2번째 측정 값 우클릭 - > 아래에서 3번째에 `이중 축` 클릭

- 이중 축의 특징 : 마크 카드를 보면 **각 축에 대한 마크 카드가 생긴 것을 알 수 있음**
	- 이게 Combined Axis Chart와 구분되는 중요한 점임

- 그래프가 임의로 바뀌는데, 막대 그래프로 표시하려면 `마크 - 전체 - (원 -> 막대로 변경)`
- 이대로면 두 막대 그래프가 완전히 겹쳐서 보이게 되므로, **두 막대 그래프를 모두 보이게 하기 위해 해당 측정값의 `마크 카드 - 크기`를 줄여준다.**
- 강렬해보이는 색은 불투명도를 줄이는 방식으로 해결할 수 있다.

#### 두 축의 스케일 일치시키기
- `한쪽 축 우클릭 - 축 동기화`

> 항상 축 동기화를 쓰는 게 좋은 것은 아니다
- 두 측정값의 분포 범위가 완전히 다른 경우(ex : 판매액 vs 할인율)에는 동기화를 조심해야 함

#### Combined Axis Chart
- 2개의 측정값을 화면에 표시하는 또다른 방법
- 2번째 측정값을 **1번째 측정값의 축 위로 올려서 만듦**
- 각 피쳐에 대해 2개의 측정값을 별도로 표시함

- 측정값 별로 색 별도로 부여 : `필터 카드의 측정값 이름 ctrl + 드래그 -> 마크 카드의 색상에 드랍`

### 이중 축과 Combined Axis Chart의 장단점

#### 1. 이중 축의 장점
- 마크 카드를 별도로 확보할 수 있다 : 커스터마이징하기 좋음
	- ex) 각 마크 카드의 값을 표시할 때, `글꼴`을 누르면 `마크 색상과 일치`가 있음. 각 측정값 별로 뜨는 레이블의 색을 다르게 부여할 수 있음
	- `Combined Axis Chart`에선 불가능함 : 마크 카드가 없기 때문임
		- 되는데? `필터에서 마크카드로 복붙 - 레이블 - 폰트 - 마크 색상과 일치`

#### 2. Combined Axis Chart의 장점
- 이중 축을 쓸 수 있는 기회를 남겨두었음
	- 이중 축은 일종의 쪼커임

## 2. 강조 포인트 제공
- 예를 들면 2017년 11월의 값을 강조하고 싶다고 하자
- 전체적인 그래프를 나타내는 2개의 특성(날짜, Sales) 이외에도 별도의 특성(2017년 11월에 해당)이 필요함

- 왼쪽 검색 박스 오른쪽 `계산된 필드 만들기`
```tableau
IIF([Order Date] >= DATE('2017-11-01') AND [Order Date] <= DATE('2017-11-30'), [Sales], NULL)
```
- 2017년 11월 1개월 간의 Sales를 얻어라, 나머지 기간은 null이다

- 이후 해당 특성을 이중 축으로 추가하고 `축 동기화`
	- 이후 `마크 카드`를 통해 수정하면 특정한 곳만 강조할 수 있다

## 3. 디자인 효과, 컨텍스트 부여
- ex) 강사님 예제 : 미국 전역과 각 지역의 아보카도 가격 변화
	- 전역을 흐릿하게 표현
	- 선택된 특정 지역만 뚜렷하게 표현

- 예제 :  기간과 Sales(합계)에 따른 라인 차트를 주고, Sub-Category를 색으로 부여
1. Sub-Category 앞의 아이콘 `색상`을 `세부 정보`로 변경 -> 모든 서브 카테고리의 색이 동일
2. 특정 서브 카테고리만 강조하고 싶다 : 새로운 특성을 만들어야 함
```tableau
IIF(
	[Sub-category] = 'Bookcases',
	[Sales],
	NULL
)
```
3. 이렇게 만들어진 특성을 이중 축 & 축 동기화하면 해당 특성만 강조됨
	- 특히 **다른 특성을 회색**으로 만들고, **강조하고 싶은 특성을 빨간색**으로 하는 방법이 괜찮아보임
	- 추가로, 회색으로 만든 **특성 마크카드에 들어가서 불투명도를 낮추면 더 빨간색이 강조**됨