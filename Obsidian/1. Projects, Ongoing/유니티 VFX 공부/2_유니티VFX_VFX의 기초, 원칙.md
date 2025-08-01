#VFX #유니티VFX
1. [[#1. 게임 플레이|1. 게임 플레이]]
2. [[#2. 타이밍|2. 타이밍]]
3. [[#3. 형태(Shape)|3. 형태(Shape)]]
4. [[#4. 대조(Contrast)|4. 대조(Contrast)]]
5. [[#5. 색상(Color)|5. 색상(Color)]]


## 1. 게임 플레이
- **무엇을 위한 이펙트를 만들 것인가?**
	- 투사체
	- 범위 이펙트
		- 범위 내의 적을 공격
		- 범위 내의 아군을 힐
	- 환경 이펙트
	- 캐릭터
- 그래서 게임이 어떻게 동작하고, 매커니즘이 어떻고, 캐릭터에 관한 것인지 여부를 파악하는 게 중요하다. 클라이언트도 자신이 원하는 이펙트가 어떤 것인지 모두 모를 수도 있다.
- **게임의 모든 이펙트를 만들어야 한다면, 게임에 관한 모든 정보를 가능한 많이 뽑아내는 게 중요하다.**  

## 2. 타이밍
>[!info]
>- 타이밍의 종류
>1. Anticipation - 시작 
>2. Climax - 중간
>3. Dissipation - 끝
>- 쉽게 말하면 폭발하는 부분을 기점으로 전조와 후처리로 나뉜다 정도?

- **각 파트를 얼마나 길게 가져가느냐 여부도 이펙트의 인상에 영향**을 줌
- 이 부분은 [강의](https://www.udemy.com/course/vfx-for-games-in-unity-beginner-to-intermediate/learn/lecture/15540264#overview)를 보면서 어떤 인상인지 보는 게 더 나을 수도 있겠다.

---
- **예시 1**
- 강의 내용에서 보여주는 폭발 이펙트를 보면
1. `Anticipation`이 없을 때
	-  단순한 피격 이펙트로 생각해도 되는 듯
2. `Anticipation` : 0.5초
	- `Anticipation`을 보여줌으로써 위험이 임박했고 피할 시간이 얼마 없다는 인상
3. `Anticipation` : 1초
	- `Anticipation`이 보여지지만 피할 시간은 있어 보임
	- 모이는 시간이 더 길기 때문에 더 큰 대미지를 입을 것으로 보임

즉, **`Anticipation`은 기대, 예상을 만든다.** 

4. 3과 동일한 조건이되 `Anticipation`이 달라짐(기존엔 커지다가 다시 모이고 폭발, 여기선 큰 이펙트가 작아진 다음에 폭발)
	- 똑같이 1초를 줬지만 이펙트가 달라지면서 달라짐
	- (강의에서도 구체적으로 어떻게 달라지는지는 설명하지 않음. 실제로도 3처럼 구현하지 4처럼 구현하진 않을 것 같다는 느낌?)

5. `Dissipation`을 길게
	- 그렇게 위험해보이지 않음. 
---
- **예시 2**

1. `Anticipation`
	- 바닥 쪽에는 화염 이펙트들이 가운데로 모이는 모양
	- 가운데에도 원기둥이 하나 있는데, 반경이 좁아지면서 가운데로 모인다.
	- 이펙트가 모이면서 에너지가 모이는 듯한 인상을 준다. 겹쳐지면서 더 밝아보이는 효과도 있음
	- 최종적으로 가운데로 모인 이펙트들이 하나 둘 사라짐

2. `Climax`
	- 폭발 효과. **매우 빠르게 일어나며, 이 때의 밝기가 가장 높다.**
	- 이 부분은 매우 짧으며, 곧바로 `Dissipation`으로 넘어감.

3. `Dissipation`
	- `Climax`에서 빠르게 퍼지는 효과들이 사라짐
	- 한편 남아있다가 서서히 꺼지는 효과도 보임 
		- 폭발 자리의 바닥 균열 같은 부분

- `Anticipation`과 잘 연계시킬 수 있는 요소도 있다 : 아래의 `Shape`.

---
## 3. 형태(Shape)
- `Anticipation` 과정에서, 나선형으로 흔적을 남기면서 바깥에서부터 가운데로 들어가는 이펙트가 있다고 생각해보자. 이는 그 위치에서 대미지가 발생할 것이라는 걸 알려주는 전조가 된다. 

- 투사체에 관해서도
	- 끝이 뾰족하다면 상대에게 대미지를 줄 걸로 예상할 수 있고
	- 끝이 둥글다면 기절이나 빙결 등을 가할 걸로 예상할 수 있다. 
	- 노랑-초록의 가만히 떠 있는 오브라면 해로운 효과는 아닐 거라고 예상할 수 있다.
		- 이 노랑-초록색은 `friendly`하다는 시각적 효과를 줌

- **형태는 비쥬얼 이펙트의 의도를 알려준다.** 

## 4. 대조(Contrast)
- **사람의 눈은 가장 밝은 곳을 향한다.** 
- 배경이 어두운 상태라면
	- 밝은 점이 나타났다 사라지는 걸 반복한다면 사람의 눈은 그 점들 하나하나에 눈이 가고
	- 궤적의 경우도 마찬가지. 

- 이거 그 `Colors & Lights`에서 배웠던 명도의 중요성을 얘기하는 것 같음
	- **이미지가 갖고 있는 정보는 명도 > 채도 > 색상 순으로 많다**는 그 내용.

- 어떤 게임에서는 **캐릭터 능력의 중요도를 강조하기 위해 대조를 사용**하기도 한다. 
- 강의에서는 롤로 예를 들었는데, 챔피언은 럭스로 보인다.
	- 이미지를 흑백으로 만들었을 때, **가장 대조가 큰 부분은 궁극기**다. 
	- 단순히 밝은 부분의 면적이 크기 때문은 아닌 듯. 평타나 Q의 명도가 궁극기의 명도보다 확실학게 낮다.

- 따라서, **더 강조하고 싶은 부분에는 더 큰 대조를 쓰라**는 게 주 내용.
- 반대로 환경처럼 크게 강조하고 싶지 않은 부분에는 낮은 대조를 쓰면 된다.

- 대조 개념은 명도 뿐만 아니라 채도에서도 사용할 수 있다. 

## 5. 색상(Color)
- 색은 생각보다 단순하진 않다. 
	- 빨강은 피이고 파랑은 물일까? 그게 전부는 아니다.

- 똑같이 물이 튀는 듯한 2개의 이미지가 있다. 
	- 둘 모두 파랑 계열의 컬러 팔레트를 쓰지만, 한 쪽의 채도가 조금 더 낮다고 하자.
	- 이런 상황에서 한 쪽을 얼음, 다른 한쪽을 물이라고 한다면 어떻게 이들을 구분할 수 있을까?
	- 일반적으로 채도가 낮은 쪽을 얼음이라고 하겠지만, 다소 애매해 보인다. 요령으로, 눈송이 같은 이미지를 추가하는 방식으로 얼음이라는 걸 보여줄 수 있음.

- 초록 계열의 색도 마찬가지다. 
	- 힐
	- 자연
	- 방사능
	- 독
다 가능함. 이것도 위에서 요령을 부린 것과 마찬가지로, 한 쪽 이미지에는 가우시안 블러 처리로 밝아 보이는 효과를 줘서 방사성을 어필할 수 있고, 다른 쪽에는 녹색 십자가를 추가해서 힐임을 강조할 수 있다. 잎을 추가한다면 자연일 수도 있는 것이고.

이렇듯 **컬러 팔레트만 보고 그게 어떤 특성을 나타내는지는 다소 불분명할 수 있다.** 강의에서는 원하는 속성의 다른 이미지를 추가해서 묘사를 보강하는 방식으로 돌파했다.
![[Pasted image 20250724155415.png]]
- 그래서 어떤 인상을 만들려면 **색깔보다는 요소와 형태를 사용**하는 게 훨씬 확실하다. 색깔 자체로도 의미를 줄수 있지만, 요소와 형태를 곁들인다면 테마에 대한 인식을 더 강조할 수 있다.


- 컬러 팔레트의 색깔 숫자를 늘인 예시. 나중에 참고해서 쓰기 좋을 듯.
![[Pasted image 20250724160335.png]]

컬러 팔레트는 조합이 무궁무진하니까 그 외에도 많겠지만, 이런 컬러 팔레트들의 공통점이 있다 : **모든 컬러 팔레트에서는 대조를 사용하고 있다**는 것. 특히 밝음 <-> 어두움에 주목해서 보자.