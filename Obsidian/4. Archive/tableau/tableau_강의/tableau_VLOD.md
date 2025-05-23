#tableau_levelup 

## 1. VLOD란?
- `View Level Of Detail`
- 가장 단순화된 레벨부터, 가장 디테일한 레벨까지
	- **단순화**된 레벨을 **VLOD가 높다**
	- **디테일**한 레벨을 **VLOD가 낮다**라고 표현함

#### VLOD에서 집계된다

> ex) Category, 합계(Sales)를 올린다
- Sales는 `Category 수준`에서 `합계로 집계`되고 있음

> ex) 이 상태에서 Sub-category를 올린다
- 이러면 `Sub-Category` 수준에서 집계가 됨

- VLOD에서 표현된다 = VLOD에서 집계된 결과가 눈에 그대로 보인다
---

> 문제 : VLOD를 결정하는 건 오직 차원일까?
- 위 예제에서 `Category, Sub-category`가 화면에 집계되는 방식을 바꿨으니 맞지 않을까?
- **실제로도 맞다. 오직 차원만이 VLOD를 바꿀 수 있다.**

---
- 화면의 왼쪽 아래에 **`n개 마크` 가 변하면, VLOD가 변했다는 말과 같음**
- **`차원`을 행, 열에 올리는 것 이외에도, `마크 카드`에 올리는 것도 VLOD를 바꾼다.**
	- 단, `도구 설명`은 VLOD를 바꾸지 않는다
- `필터`에 차원을 올리는 것은 VLOD를 바꾸지 않는다.
---
- `페이지`도 VLOD를 바꾸지 않는데, 얘는 조금 애매하다
	- 가이드북은 아닌데 온라인 헬프는 바꾼다고 하고 있음
	- 근데 온라인은 업데이트가 되니까
	- 페이지가 많이 쓰이는 기능은 아니기 때문에 넘어가도 무방
---
### VLOD를 변화시키는 위치
- 행
- 열
- 색상
- 크기
- 레이블
- 세부정보
- (페이지)
- ~~필터~~
- ~~도구설명~~

---
## 2. VLOD에서 모든 측정값이 집계되고 표현되는가
- LOD Expression이 들어오면서, **VLOD에서** 모든 집계된 측정값이 **표현되는 건 맞지만**, **집계가 VLOD에서 일어나지 않을 수도 있게 되었다**

> 예시 질문
> 1. 니네 회사 연봉 평균 대비 얼마임?
> 2. 용인에서 경기도 평균 대비 얼마나 더 팔림?
> 3. 광역 지자체의 GDP에 대한 소속 지자체의 기여도는 얼마임?
- LOD Expression 이전에는 Tableau에서 대답하기 쉽지 않은 질문이었음

## 3. LOD Expression

### 표현식
- 중괄호`{}`를 이용함
```tableau
{ Include [차원1], [차원2] : SUM([측정값])}
	1            2               3
```

#### 1. 종류
- `Include` : VLOD에 포함되지 않은 차원을 포함하고 싶을 떄
- `Exclude` : VLOD에 포함된 차원을 제외하고 싶을 떄
- `Fixed` : VLOD에 상관없이 차원을 고정하고 싶을 때

#### 2. 차원
- 복수개가 들어오면 `,`로 구분함

#### 3. 집계값
- 반드시 집계값이 사용되어야 하며, `ATTR`이나 `테이블 계산식`은 사용할 수 없다.

