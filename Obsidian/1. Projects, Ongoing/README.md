## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 데이터가 제때 수집되고 있는지 눈팅 정도만 하면 충분할 것 같다.
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음?

## 작업 예정

### 진행 중
- 남은 작업들
	- 스테이지 1-0 ~ 1-3 밸런싱
	- 튜토리얼 구현 (진행 중)
	- 1-3에 보스 추가
	- 보상 설정

#### 튜토리얼 구현 계획
- 스테이지에서
	- 코스트 소개
	- 배치 방법 소개
		- 하단 클릭 및 드래그, 초록색 칸에 배치 가능
		- 커서를 놓으면 다시 클릭해서 방향을 설정하게 유도
		- 1명은 강제로 배치하게 함 : 뱅가드
		- 뱅가드를 클릭해서 몇 명을 저지할 수 있는지를 알려줌
	- 나머지 2명은 자율로 둠. 스테이지 클리어까지 튜토리얼은 별도의 동작을 하지 않음. 
		- 자율로 둔다면 스테이지 실패도 염두에 둬야 할 듯?
- 스테이지 클리어 후 아이템 지급, 아이템 사용법(레벨업, 정예화) 안내. 튜토리얼에선 레벨업까지만 안내하고 종료



### 부차적인 이슈
-  `FloatingText`가 나타났으면서 오퍼레이터 클릭 등으로 카메라 각도가 돌아간 상태일 때, `FloatingText`의 각도가 틀어지는 현상이 있음

### 구현 예정
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다 (= 오퍼레이터를 배치하지 않고 스테이지 클리어하기)
### 발생 중인 이슈
- (241216) 레벨업 후에 `OperatorLevelUpPanel`의 스냅핑 동작이 정상적으로 동작하지 않을 때가 있다. 
- 간헐적인 문제
	-  `Enemy` 기준, 바리케이드 파괴 로직이 동작하지 않는 현상이 있음

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 


## 3월

## 250317 - 짭명방

### 튜토리얼 만들기
- 오늘은 본격적으로 스테이지 씬에서의 튜토리얼을 만들 예정

- 튜토리얼 시작(DialogueBox) 시 시간이 멈춰야 함  -> 아래의 타임 매니저 구현으로 이동

- 구현 완료. 근데 TimeScale을 0으로 하면 텍스트가 출력되지 않는 듯. 
	- 기존의 `Ienumerator` 들을 전부 `WaitForSecondsRealTime`으로 바꾸되, 시간을 멈추는 건 스테이지 씬에서만 동작시킨다.

- 어떻게 `DeployableBox`의 동작을 튜토리얼과 연결시킬 것인가가 고민됨. 
	- 일단 `Step` 단위로 `TutorialManager`에서 체크를 하고 있기는 하다. (저장은 `TutorialData` 파일 단위긴 한데)

- 그냥 **단순하게 안내만 하고, 스테이지를 클리어하면 다음 튜토리얼로 넘어가는 방식으로 진행하겠음**
	- 다른 요소들을 어떻게 처리해야 할지 감이 안 잡혀서, 텍스트로 안내하는 방식으로 진행함.

#### TutorialPanel.TypeText :  리치 텍스트 처리
- `<color=#ff7485>라이프 포인트</color>`
	- 지금의 구현에서 위 글자 처리는 일일이 다 적음 -> 색이 적용됨 같은 방식으로 구현된다. 이걸 색이 적용된 상태로 라 이 프 포 인 트 같이 나타나게 하고 싶음.

- 아래처럼 구현. 단순히 **리치 텍스트의 문법이 나타나면 텍스트가 한꺼번에 들어가는 방식**이다. 꺾쇠는 쓸 일이 없을 것 같아서 이렇게 구현함.
```cs
// TypeText의 일부로, 아래처럼 수정했다.
// 1. TypeText를 인덱스 기반으로 동작하게 바꿈
// 2. 리치 텍스트 태그에 쓰이는 <을 만나면, >까지는 한꺼번에 입력되는 방식
int index = 0;
while (index < fullText.Length)
{
	// 만약 현재 문자가 '<'이면 태그 시작 -> '>'까지 전체를 한 번에 추가
	if (fullText[index] == '<')
	{
		int tagEndIndex = fullText.IndexOf('>', index);
		if (tagEndIndex != -1)
		{
			string tag = fullText.Substring(index, tagEndIndex - index + 1);
			currentText += tag;
			textComponent.text = currentText;
			index = tagEndIndex + 1;

			yield return new WaitForSecondsRealtime(typingSpeed);
			continue;
		}
	}
	
	// 그냥 글자 출력
	currentText += fullText[index];
	index++;
	textComponent.text = currentText;
	yield return new WaitForSecondsRealtime(typingSpeed);
	
}
```

### TimeManager 구현
- 메인메뉴에서는 시간에 관한 매니저가 없었고(필요해 보이지도 않고), 스테이지에서는 StageManager에서 관리했음.
- 그렇더라도 공통적으로 `TutorialManager`에서 시간을 관리하는 메서드를 호출할 필요는 있어서, `GameManagement`의 하위 오브젝트로 `TimeManager`를 새로 추가하고 시간 관리 로직을 옮김
- `isSpeedUp`의 상태 관리 관련, 일시정지 / 재생 관련 메서드 수정 완료







## 250316 - 짭명방

### 튜토리얼 만들기 - 스테이지

- 스테이지 진입
	- 코스트 상승 안내
	- 뱅가드 배치 코스트에 도달 시 배치 조작 안내
		- 뱅가드 먼저 배치하는 이유) 코스트 수급
		- 하단 패널 클릭
		- 왼쪽 안내 패널에 대한 안내
		- 드래그해서 특정 타일에 놓게 하기
		- 배치된 오퍼레이터를 클릭한 뒤 퇴각 / 스킬 버튼 안내
			- 스킬 SP 버튼도 안내

- 씬 전환에 따른 튜토리얼 시작 타이밍 설정
	- 이거 교통 정리만 해도 복잡하다.
	- 기존 구현은 곧바로 다음 튜토리얼을 실행시키는 방식이었는데, 대신 **각 `TutorialData`가 종료될 때 `PlayerDataManager`에 튜토리얼 진행 상태를 저장**하고, **씬 전환 & 튜토리얼이 시작되기를 원하는 타이밍에 특정 튜토리얼을 실행시키는 방식으로 변경**함.

- 튜토리얼 중 재실행 등의 상황이 있을 때를 대비, `TutorialProgress`를 `PlayerData`에 저장하고, 씬이 로드될 때마다 점검한다.

- 2번째 `PlayerDataManager` 실행할 때 스택 오버플로우가 나고 있다.
	- ?? 상호 참조도 아닌데 왜 나는 건지 모르겠음.
	- 아 ㅋㅋㅋㅋㅋㅋ `GameManagement`에 `TutorialManager` 프로퍼티가 `TutorialManager => TutorialManager!`로 되어 있었다. `tutorialManager!`로 수정.

- 일단 `StageScene`에 진입한 다음 2번째 `TutorialData`가 실행되는 것까지 완료.

#### 분기 생각
>1. `Vanguard` 배치만 안내하고 이후는 대기?
>2. 편성된 오퍼레이터를 모두 배치하게 한다?

- 이런 것들보다 더 중요한 게 **튜토리얼 스테이지 패배 시의 분기 처리**라고 보는데, 단순하게 접근하면
	- 마지막 로비로 돌아왔을 때의 `TutorialData`는 `Stage1-0`이 클리어됐는지를 점검(3성이 아니어도 무관), 클리어됐을 때에만 이후 튜토리얼이 진행되게 한다

### 이상한 버그
 - [x] 여러 번 클릭했을 때 글자가 깨지면서 다음 스텝도 진행되는 듯 현상이 있음. 그래서 글자가 막 깨지는 것처럼 보인다. 
	 - 예상 후보) `Dialogues`의 수가 2개 이상인 경우 / `Require User Action` 체크 여부 / `Expected Button Name` 여부
	 - **`Expected Button Name`이 `string.Empty`일 때에만 광클했을 때 여러 대화 코루틴이 동시에 실행되는 현상이 있었음.**
```cs
    // 클릭 시 동작
    public void OnClick()
    {
        // 글자가 나타나는 중 : 현재 페이지의 글자를 모두 보여준다
        if (isTyping)
        {
            if (typingCoroutine != null)
            {
                StopCoroutine(typingCoroutine);
                typingCoroutine = null;
            }

            textComponent.text = fullText;
            isTyping = false;
            OnCurrentPageFinish();
            return; // 타이핑 중일 때 클릭해도 다음 대화로 넘어가지 않게끔 함
        }

        // 다음 페이지로 넘어갈 수 있음 : 다음 페이지로 넘어간다
        if (CanMoveToNextPage)
        {
            currentPageIndex++;
            typingCoroutine = StartCoroutine(TypeText());
            return;
        }
    }
```
> 1. 동작 중인 `typingCoroutine`을 종료하는 걸 우선하고 나머지 조작을 수행할 것
> 2. 다음 페이지로 넘어갈 때 `typingCoroutine`을 설정하지 않았던 이슈도 있었다.

## 250314

### 짭명방
#### 튜토리얼 만들기 계속)
- 기존 UI 위에 투명한 버튼을 만드는 `CreateOverlay` : 기존 UI가 `GridLayout` 등 고정된 위치에서 동작하는 경우가 아닐 때 값이 0으로 설정되는 이슈가 있음
	- 일단 `CreateCurrentOverlayAfterDelay`를 만들어서 유니티 내에서 위치 연산이 수행된 후에 해당 위치를 얻도록 수정

 - [x] `Instance(Clone)` 찾게 하기 
	 - `OperatorInventoryPanel`의 경우, 현재 가지고 있는 오퍼레이터들이 나타나도록 구성됨. 이는 프리팹의 인스턴스화가 되는데, `OperatorSlot`이라는 공통된 프리팹을 사용하므로 이름들이 동일한데, 인스턴스를 하는 패널에서 이름 설정을 바꿔주면 됨.

- [x] 경우에 따라 상호작용 버튼을 필요로 하지 않을 수도 있음
	- `RequireUserAction` 체크 여부에 따라 갈림 
	- 이거 구현하는 과정에서 항상 더블클릭을 해야 넘어가는 느낌으로 되어버렸음. 수정 ㄱ
	- `TutorialManager.WaitForUserAction`에서 분기를 나누는데, 기본적인 동작은 `TutorialPanel에서 대화가 모두 나오는 이벤트를 받음 -> TutorialManager에서 다음 스텝으로 넘어가기 위해 클릭되어야 하는 요소를 설정하고 클릭 대기 -> 다음 스텝으로 넘어감`이다. **분기는 단순히 어떤 버튼이 클릭되어야 하는가** 정도임

- 일단 **1번째 Step : 스테이지 진입 전까지 3명의 오퍼레이터를 편성하기** 까지는 완료.

- [x] `DialogueBox`의 텍스트 출력이 종료되고 나서, 클릭 가능함을 나타내기 위한 인디케이터를 추가했다. 
![[Pasted image 20250314212242.png]]
> - 우측 하단에 나타나는, 아래로 향하는 삼각형 이미지.
> - 깜빡거리는 효과를 `BlinkingShader`라는 이름으로 셰이더 그래프를 만든 뒤 머티리얼로 만들어서 넣었다. 시간을 2로 나눈 모듈로 연산이며, 0~1초는 알파값을 1로, 1~2초는 알파값을 0으로 하는 방식.
> - 원래 그랬는지 모르겠는데 **셰이더 그래프의 `AlphaClip`을 켜야 `png`파일의 투명한 배경이 적용된다.** 꺼져 있으면 검은 배경으로 나타남.
> - 원래는 현재 스텝의 마지막 대화일 때는 셰이더와 인디케이터 이미지를 다르게 적용하려고 했는데, **사용자 입장에서는** 개발자 입장에서 봤을 때 **스텝을 구현한 게 그렇게 드러나지 않기 때문에 그냥 저것만 사용하는 방식으로 통일했다.** 이쪽이 훨씬 간편하기도 하다.

## 250313

### 짭명방

#### Step 동작 문제
- `DialogueBox`가 텍스트의 마지막 페이지에 진입하면, 목표로 하는 동작을 바로 가능하게끔 구성하려고 함. 

- 어제 발생했던 문제 정리
	- [x] `DialogueBox`가 나타남 -> 버튼을 클릭해도 `DialogueBox`가 없어지지 않는 문제
		-  `maxPageIndex` 값을 초기화할 때, `step.dialogues.Count`으로 초기화하고 있었다. **마지막 인덱스는 `Count - 1`이어야 함.** 허무하다.

- 오늘 발견한 문제
	- [x] 마지막 `Dialogue`가 아닐 때에는 클릭했을 때의 동작이 아예 진행되지 않고 있음.
		- 일단 구조 변경 : **`TutorialPanel`을 별도로 구현하고, 그 안에 `TransparentPanel`과 기존에 구현한 `DialogueBox`를 집어넣음**
		- `DialogueBox`를 `TutorialPanel`로 이름 변경. 실질적인 동작은 `DialogueBox`가 대부분. 하지만 `TransparentPanel`이 화면 전체를 차지해야 하기 때문에 지금처럼 변경

	- [x] `TutorialPanel`에서 `Dialogue`가 모두 나온 다음 버튼 클릭을 기다리는 과정에서의 UI 시스템 설계
		- 뒤에 투명한 패널을 구현하고, 거기에 `OnClick` 동작까지는 구현했다. 그런데 아래에서 막힌 상황이다.
			1. 목표 버튼을 클릭했을 때 다음으로 넘어가기
			2. 목표 버튼 외의 다른 것을 클릭했을 때, 아무것도 동작하지 않게 하기
			- 2번은 `TransparentPanel`과 `OnClick`의 동작만으로 충분한데, 1번이 `TransparentPanel`이 있는데도 동작시킬 수 있는가가 중요하겠음.
			- `ClickDetectionSystem`에서의 동작은 주석으로 달아놨듯이 `GraphicRaycaster`에 해당하는 유니티 자체 동작이 먼저라서 상관 없음.
		- AI와의 대화를 꽤 오래 나눠본 결과, **타겟 버튼과 동일한 위치, 동일한 크기, 동일한 기능을 하는 투명한 버튼을 `TutorialPanel` 위에 하나 더 만드는 방식으로 정리**됨

> 지식이 늘었다 : **`Button` 컴포넌트가 클릭을 감지하려면 `Image` 컴포넌트가 함께 있어야 한다.** **`Image.raycastTarget`가 클릭을 감지하기 때문**임.

- 일단 `SquadEditPanel`까지 들어가는 과정은 튜토리얼로 구현됨. 
- 이제 3명의 오퍼레이터를 편성하고, 스테이지에 들어가고, 배치 방법을 알려주고, 게임을 클리어하면 보상이 지급되고, 그 보상을 레벨업과 정예화에 사용하는 방법까지 알려주는 일련의 과정을 구성하면 됨
- `overlay`의 경우 `Button` 컴포넌트가 없는 경우가 있을 거임(스테이지 씬에서 특히) -> 여기서의 별도 구현도 나중에 해야 함!


## 250312

### 짭명방

#### 튜토리얼 만들기..의 일부인 DialogueBox를 겨우 만듦

- 구성요소 만들기
	- 텍스트 박스
```cs
1. 박스에 한글자씩 자동으로 나타남.
2. 유저가 화면을 클릭 시 현재 텍스트 박스에 나타나야 할 모든 글자가 즉시 나타남 
3. 1번의 경우도 아무 동작을 하지 않았을 때 최종적으로는 2번의 결과와 동일함 
4. 모든 글자가 나타난 상태에서 클릭 시, 다음 텍스트 박스로 넘어가거나 다음 동작으로 넘어감. 1번의 상태에서 클릭했을 때 다음 텍스트 박스로 넘어가는 상황은 지양함.
```
> 4번의 경우는 명일방주에서 튜토리얼을 진행할 때 무심코 다음 대화로 넘어가는 현상이 발생한 기억이 있어서 특히 신경쓰는 부분이 될 듯

- 복잡하다. 만들면서 생각해보는 것들.
1. 텍스트를 넣을 때 `DialogueData`에 구분점을 스크립트로 작성해서 넣을지, 아니면 일일이 다음 문장으로 넘어가기 전에 구분하는 포인트를 만들어줄지에 대해
	- ~~스토리 같은 걸 만드는 경우라면 개발자가 일일이 지정하는 게 맞아보임. 보통 스탠딩 스프라이트도 함께 나타날 걸 생각하면, 어떤 텍스트가 나타나면서 어떤 스프라이트를 띄울지를 일일이 결정하는 과정을 거쳐야 하기 때문임~~
	- ~~그런데 **지금처럼 튜토리얼을 작성하는 경우, 현재 수행하고 있는 동작만 명확하다면 한 번에 보여줄 텍스트의 개수는 스크립트로 작성해서 구분해도 될 것 같음.** `텍스트 보여주기 -> 사용자에게 동작 수행을 유도하기`의 반복이기 때문. ~~
	- 만약 문제가 생기면 그때 생각해보기로 하자.
	- 같은 패널에서도 `DialogueBox`의 위치가 바뀐다든가 하는 경우가 생길 수 있겠다. 
		- 예를 들어서 시작 직후에, 
		- `1-0 Button`을 눌러라 -> `부대 편성하기`를 눌러라 로 이어질 때, 1-0은 화면 윗쪽에, 부대 편성하기는 우측 하단에 위치하는 게 좋아보임.
	- 그래서 **스크립트 자체는 구분점 등을 일일이 지정**하는 것으로 수정.

- `TutorialData`는 여러 개의 `TutorialStep`들로 구성되어 있다.
- 여러 개의 `TutorialStep`은 
	- `stepName`
	- `dialogues` : 안내 문장들
	- `dialogueBoxPosition` : `DialogueBox`가 나타날 위치
	- `actionAfterDialogue` : 이후 수행할 액션 유형. 사용할지 안할지 모르겠음.

- `TutorialPanel`이라는 걸 따로 구현을 해둘까 싶다. 거기에 아래의 구성요소들을 추가하면 되지 않을까? 
	- 텍스트 박스
	- 뒷배경 가리기
	- 혹은 포인터 이미지를 하나 넣기

- 개요를 짜고, AI가 던져준 코드를 보면서 너무 많은 요소를 생각하고 구현하려니까 하나도 진행을 못하는 상황이 몇 시간 정도 있었다. 그래서 가급적이면 **최소한으로 필요한 기능들만을 구현해나가는 방식으로 진행**한다. 
	- 이거는 AI 쓸 때 특히 중요한 거 같음. 특히 끌로드가 필요한 기능 이상으로 코드를 주는 경우가 많음...

##### ~~정리~~
- ~~`DialogueBox`는 `Step` 단위로 진행됨. `Step` 내에는 `dialogues` 들이 있고, 이들이 모두 진행되면 현재 스텝이 종료되는 방식.~~
	- ~~이거는 나중에 `TutorialPanel` 등에 포함될 수도 있음. 다른 상호작용 가능한 요소의 클릭을 방지하는 배경을 하나 추가할지도 모르겠다는 생각이 들고 있기는 하다.~~
	- ~~고칠 필요가 있다면 그 때 고치면 된다. 일단은 **기능 구현에 집중.**~~
- ~~이 `Step`들을 전체적으로 관리해야 하는 게 `TutorialManager`의 역할임~~

##### 오늘 내용 정리(진)
- `DialogueBox`는 `Dialogue`들만 관리함. `Step`은 이것보다 살짝 큰 개념으로, **`Dialogue`에서의 내용들이 모두 나타남 -> 사용자의 동작을 기다림 -> 사용자의 동작 완료 시 다음 스텝으로 넘어감** 처럼 구현됨
- 보통 사용자의 동작은 `Button`으로 감지되므로 `DialogueBox`가 만들어질 때 `TutorialManager`에서는 해당 이벤트에 대한 `WaitForUserAction`을 동작시킨다. 
- 머리가 아파서 오늘은 일단 스탑.

## 250311

### 짭명방

#### 튜토리얼 만들기
- 튜토리얼의 시작부터 구현
	- `PlayerDataManager`에서 `TutorialData`라는 부분을 만들고, `PlayerData`에 통합
	- 이를 감지해서 최초 실행 시에 튜토리얼 여부를 물어봄

- 물어보는 패널인 `ConfirmationPanel`이라는 요소를 공용으로 만든다.
	- 기존에 구현된 게 `ConfirmationReturnToLobbyPanel`인데, 여기서 공통으로 사용할 요소를 빼는 방식으로 진행.

- `ConfirmationPanel`은
	1. 뒷배경`BlurArea`가 나타날지 여부와
	2. 취소 버튼`CancelButton`이 나타날지 여부
	- 를 초기화 로직에 포함시켰음.
	- 그 외에도 이것저것 있는데 문제가 발생하는 부분은 다 수정했다.

- 스테이지 버튼 안내
- 편성 안내
- 인게임 배치 안내
- 보상과 성장 안내
...등을 구현하면 되겠다. 우선적으로는 이를 위해 `Canvas`에 나타나는 대화상자 패널이나 하이라이트 패널도 구현해야겠음.

## 250310
- 남은 과정
	- 튜토리얼 추가하기
	- 스테이지 다듬기(적의 등장, 맵 설정 등등)
	- 보스 추가하기
	- 보상 설정하기

이 정도일 것 같음.

### 짭명방

#### 프로젝트 실행하면서 nullable 이슈 다듬기
- **유니티의 `null` 체크는 비싸다**
	- [[Unity - Null 체크에 관해]]
	- 그러나 일단은 **`Update` 등 자주 호출되어 사용되는 메서드 정도만 아니면 사용해도 크게 이슈가 발생한 상황은 없는 것 같음**(최소한 지금까지는 그랬다. 경고문이어서 굳이 처리하지 않아도 됐지만 어쨌든 경고문이라서 다듬는 편이 코드상 더 좋을 것 같아서.)

#### 기타 이슈 수정
- `Map`의 초기화 관련
	- 맵 에디터에서의 `Initialize`와, 실제 스테이지에서의 `Initialize`가 혼용되고 있다. 이 스크립트들을 정리해야 함.
	- 정리 완료. 그냥 `SerializedData` 부분만 정리했다. 맵 데이터는 프리팹 내부에 저장이 되니까 굳이 이걸로 따로 빼서 정리할 필요는 없고, `Initialize`를 할 때 타일들만 한번 긁어서 맵 데이터를 초기화하면 됨.

## 250308-250309 새벽
> 코파일럿 프로 사용 시작. 강력하다.
### 짭명방

#### 약 7~800개 nullable 경고문 해결하기
- 유니티의 버전을 업그레이드하면서, 원래는 나타나지 않았던 경고문인 `nullable`과 관련된 경고문들이 700~800개 정도 나타나기 시작했다.
- 어제부터 계속 헤매고 있는데, 일단은 이런 스크립트를 하나 만들어둠.
```cs

public static class InstanceValidator
{
    // [return: NotNull]
    public static T ValidateInstance<T>(T? instance)
    {
        if (instance == null)
        {
            throw new InvalidOperationException($"{nameof(instance)}가 null임");
        }

	// return instance;
    }
}
```
> 최종
> 1. 어떤 변수가 `null`인지를 체크하는 메서드가 여러 스크립트에서 반복적으로 쓰이므로 그러한 부분들을 이 메서드로 처리
> 2. 이 메서드로 처리된 변수는 이후의 스크립트에서 `!`을 추가, `null` 경고를 무시
> 3. 싱글턴 오브젝트의 `Instance`는 이걸로 처리하지 않고, `Awake`에서 전역적인 유일성이 보존되므로 단순히 `Instance!`으로 사용
> 4. `null`이어도 되는 경우들이 있음 : 그런 경우에는 이 메서드를 사용하지 않고 `null` 체크만으로 충분
##### 수정 내역
- `ValidateInstance`는 `nullable`인 인풋을 받아서, 아웃풋으로 `null`이 아님을 이 변수를 사용하는 외부 메서드에 알리는 역할을 한다
-  비슷한 처리가 많아서 아예 `utils/InstanceValidator`로 빼뒀다.
- `return: NotNull` : 단, 어떤 메서드에서 `null` 체크를 직접 할 때는 `CS` 경고문이 없어지지만, 이렇게 구성하는 경우 `nullable` 경고가 다시 뜬다. 그래서 이 함수에서 반환시키도록 하고 `null`이 아님을 외부에 알리도록 함.
- 싱글턴을 사용하는 경우가 애매했는데, `Awake`에서 초기화하는 것을 보장할 수만 있다면(일반적으로는 그렇다고 봐야 하므로) **`Instance`를 사용하는 부분에 `!`을 붙여서 `nullable` 경고문을 무시시킨다. 그 경우는 이 메서드를 사용하지 않음.**
- 주석 추가) `return`을 추가할 경우 이 메서드를 어떻게 활용할지가 애매해지는 부분들이 생긴다. 그냥 `null` 체크를 여기서 하고, `null` 체크가 된 부분은 `!`을 붙이는 식으로 넘어가겠음.

> 어제 하나하나 수정할 때는 '언제 다 수정하냐'였다. `nullable`로 지정하니까 수십 개가 더 생기는 방식. **하지만 대부분의 경고문은 `ScriptableObject`에 있는 필드들에 `nullable`을 달아주니까 사라졌다.**
> 사실 `nullable` 자체는 사용해야 하는 게 맞기는 하다. 지금까지는 모호하게 코드를 작성해 온 면이 있는 것 같음.
> 는 아래 오류를 처리하니까 다시 750개. ㅋㅋㅋㅋㅋㅋ

#### 계속 진행 - nullable이면 안되는 필드 관리
- **`nullable`이면 안되는 필드에 대해**
	- 예를 들면 어떤 엔티티의 이름이라든가? 추후에 인스펙터에서 할당할 요소지만 당장 넣지 못할 요소 같은 게 생길 수 있음
	- **처음에는 '`nullable`로 선언하라는 걸 고려해봐라'라는 경고문의 메시지를 보고 그렇게 수정했는데, 생각해보면 해당 요소들이 `null`인 걸 가정하는 것 자체가 이상함.** `null`인 순간이 있기야 하겠지만, 추후에 할당을 할 예정이기 때문임
```cs
    public GameObject deployablePrefab;
    public GameObject deployablePrefab = default!;
```
- 그럴 때 사용할 수 있는게 `default`이다. **추후에 인스펙터 등에서 정상적으로 할당될 것임을 컴파일러에게 알리는 역할을 한다.**  여기에 `!`을 붙여서, `default`로 할당하더라도 똑같이 `null`이 들어가지만, 경고문을 무시하게 만듦.
- 값의 경우에는 `0, string.Empty, false` 등등을 기본값으로 사용함

> - 이게 지금 정리를 해둬도 적용되지 않는 게 더 좋을 상황도 종종 있어서 함부로 정리하기가 그렇다.
> - 날 바뀜. 오전 12시 49분. 대부분이 패널에서 `default!`로 할당하거나, `Instance!`로 넘길 수 있거나 이슈이긴 했다. 그래도 아직 100개 넘게 남아 있음.
> - 새벽 2시 36분. 다 없앴다. 물론 이게 잘 동작하느냐는 또 별개이기는 하지만.. 그래도 다 없앴다. 그런데 뭐 때문인지는 몰라도 **프로젝트 자체가 전반적으로 무거워진 느낌**임. 컴파일이 오래 걸린다는 느낌? null 체크 때문일까? 어쨌든 일단 덮어둔다.



### 업데이트 전부터 떴으나 무시해온 오류 처리
```cs
The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(Operator) <BaseData>k__BackingField

The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(Operator) currentStats

The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(MedicOperator) <BaseData>k__BackingField

The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(MedicOperator) currentStats

The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(Barricade) <CanDeployGround>k__BackingField

The same field name is serialized multiple times in the class or its parent class. This is not supported: Base(Barricade) <CanDeployHill>k__BackingField
```

- 유닛들의 상속 구조를 만들면서, 근간이 되는 `ScriptableObject`를 모두 **`BaseData`라는 프로퍼티 이름으로 사용**했기 때문인 것 같다.
- `new`를 붙이더라도 부모의 `BaseData`와 자식의 `BaseData`가 같은 이름의 백업 필드를 중복해서 직렬화하려고 하기 때문에 발생하는 현상임. 다른 것들도 마찬가지일 듯.
- **근데 이거 수정하니까 다시 `nullable` 경고문이 700개 추가됨. ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ** 다시 돌아간다.


## 250307

### 짭명방

- `DeployableUnitInfo`랑 `DeployableInitState`가 헷갈려서 정리해둠
	- 전자는 단순히 해당 배치 요소에 대한 값이나 상태를 긁어올 수 있게 하기 위한 클래스
		- 쉽게 말하면 `Initialize`의 타이밍 문제 등으로 인해 어떤 정보를 가져오지 못하는 상황이 있을 수 있기 때문에, 이를 사전에 초기화해둔 클래스라고 보면 된다. 이걸로 `BaseData`들에 접근할 수 있음.
	- 후자는 스테이지가 진행됨에 따른 그 배치요소의 상태를 관리하는 클래스임
		- 배치된 횟수, 그에 따른 증가하는 배치 코스트 등등..

#### 콘솔에 뜨는 LogWarning들 수정
- 일단, **IDE에서 뜨는 초록색 경고문(`nullable` 을 사용하는 건 `#nullable` 컨텍스트 내부에서 써라..)은 무시해도 된다.** 프로젝트 전체에서 이미 `nullable`을 사용할 수 있게 설정되어 있음(유니티의 프로젝트 세팅에서 `player - additional complier` 어쩌구를 찾아보면 됨)

- `nullable` 필드들과 관련, 코드 패턴을 정리해두는 게 좋겠다.
1. `SerializeField` 자체는 `Nullable` 필드가 맞다고 본다. : 인스펙터에서 까먹고 할당하지 않을 수 있으니까.
2. `SerializeField`를 사용하는 상황은 보통 `private`인데, 이게 `null`이 아님을 보장할 수 있는 상황이라면 이를 이용하는 프로퍼티는 `Nullable` 필드로 선언하지 않아도 될 것 같다.

#### 시스템 추가 구현
- 오퍼레이터가 퇴각한 적이 있다면, 배치코스트는 최초 배치코스트의 50%씩 올라가는 시스템 구현하기
	- `DeployableUnitState`에 구현이 되어 있기는 한데, 실제로 반영되어 있지는 않았다. `DeployableBox`에서 해당 값을 참조하게 하고, 오퍼레이터가 퇴각할 때마다 값을 갱신하는 식으로 수정함. 


#### 수정 사항
- [x] 뱅가드 스킬 - 코스트 회복 파티클 : 우측 하단에 너무 가까울 경우, 생존 시간이 남아있기 때문인지 파괴됨에도 다시 생성되는 문제
```cs
	// 목표 지점 근처 도달 시 파티클 제거
	if (distanceToTarget < arrivalThreshold)
	{
		//particle.remainingLifetime = 0f; // 이걸로 쓰면 파티클이 다시 생김
		Destroy(particleSystem.gameObject, 0.1f);
		return;
	}
```

- [x] 배치 후 퇴각 : 배치 가능 수가 2씩 올라가는 현상
	- 로깅을 해보니까 `DeployableManager.Instance.OnDeployableRemoved(this)` 가 `Operator.Die`에도 있고 `DeployableUnitEntity`에도 있다. `Operator`에 있는 걸 제거함.


## 250306
#### 수정 사항
- [x] 스테이지 클리어 실패했는데도 `StageResultPanel`에 보상이 나타나는 문제
	- 클리어했을 때 잘 나타남
	- 클리어하고 다시 시도해서 실패했을 때 안 나타남
	- 그냥 실패했을 때 안 나타남
- [x] ItemUIElement 관련 수정 1 - 아이템의 우측에 표시되는 `ItemDetailPanel`의 오른쪽 끝이 화면에서 벗어날 경우, 화면의 왼쪽에 나타나게 하기
	- 이런 메서드를 자세하게 몰라서 AI한테 물어볼 때는, 메서드가 어떻게 동작하는지에 대해서만 알고 직접 만지는 게 훨씬 좋아보인다.
		- 예를 들면 (현재는 Gemini 2.0 Pro를 사용 중) 이거 수정할 때 캔버스의 월드 좌표, 스크린 좌표 등등의 개념을 말하면서 코드를 던져주고 이를 기반으로 수정했지만, 실제로 적용해보면 `ItemUIElement`을 기준으로 좌-우 전환을 하는 게 아니라 패널이 어디로 튀어나가 있음
		- 근데 단순히 `GetWorldCorners` 메서드를 캔버스와 패널에 적용하고, 오른쪽에 있는 `2`번이나 `3`번의 `x`좌표를 얻어서 비교하는 것만으로도 충분했음
```cs
    // 디테일 패널이 화면에서 잘리는 경우 패널을 왼쪽으로 띄우는 메서드
    private void AdjustDetailPanel()
    {
        // 월드 좌표를 쓰는 이유 : 캔버스 내의 절대적인 위치 개념임. 스크린 좌표는 해상도에 따라 값이 달라질 수 있다.

        // detailPanel의 월드 좌표의 우측 가장자리 계산
        Vector3[] detailPanelCorners = new Vector3[4];
        detailPanelRectTransform.GetWorldCorners(detailPanelCorners);
        float detailPanelRightEdgeWorldX = detailPanelCorners[2].x; // 우측 상단 코너의 x좌표

        // 캔버스의 월드 좌표에서의 우측 가장자리 계산
        Vector3[] canvasCorners = new Vector3[4];
        canvasRectTransform.GetWorldCorners(canvasCorners);
        float canvasRightEdgeWorldX = canvasCorners[2].x; // 캔버스 우측 상단의 x좌표

        // 디테일 패널의 오른쪽 변이 화면에서 벗어나는 경우, 왼쪽으로 디테일 패널을 옮김
        if (detailPanelRightEdgeWorldX > canvasRightEdgeWorldX)
        {
            Debug.Log("detailPanelRightEdge가 canvasWidth보다 크다");
            SetDetailPanelLeft();
        }
        else // 좌측으로 넘어가지 않더라도 detailPanel의 위치를 잡아준다
        {
            SetDetailPanelRight();
        }

        // 실시간으로 위치를 변화시킬 게 아니라서 이 정도로만 구현함
    }
```

- [x] ItemUIElement 관련 수정 2 - `BackArea` : 기존에는 `3000, 2000`의 너비, 높이를 갖는 이미지가 패널이 나타날 때 패널 뒤를 가리는 역할을 했는데, 이게 나타나는 위치에 따라 화면을 온전히 가리지 못함
	- UI 상에서 구현하고자 한다면 Canvas의 자식 오브젝트로 관리하는 게 최선이기는 한데, 얘는 렌더링 순서를 고려해야 함
	- 위 문제를 해결하면서 `ItemUIElement`가 이제 `Canvas`를 참조하도록 변경했으니, 여기서도 `BackArea`가 나타날 때 캔버스의 너비와 높이를 모두 활용하는 방식으로 수정함
```cs
        Rect canvasRect = canvasRectTransform.rect;
        RectTransform backAreaRectTransform = detailBackArea.GetComponent<RectTransform>();
		Debug.Log($"{canvasRect.position}");
        backAreaRectTransform.position = -canvasRect.position;
        backAreaRectTransform.sizeDelta = new Vector2(canvasRect.width, canvasRect.height);
```
> 특이한 현상으로, `canvasRect.position`이 `960, 540`이 아니라 `-960, -540`으로 잡히고 있다. 왜 그런지 모르겠음. 일단 임시방편으로 저렇게 달아놓음.

- [x] ItemUIElement 관련 수정 3 - `BackArea`가 나타날 때, `ItemUIElement`의 일부 요소도 패널에 뒤덮이는 현상이 있음
	- `ItemUIElement`에 `Image`가 들어가 있는 경우, 자식 오브젝트인 `BackArea`가 더 위에 렌더링되기 때문으로 보인다. Image 부분을 분리해서 자식 오브젝트로 옮겨놓음.
	- `BackArea`의 경우는 아예 `ItemUIElement`의 바로 자식으로 오게끔 수정한다.
```
ItemUIElement
- BackArea
- ItemDetailPanel
	- 구성 요소..
```
> 기존엔 `BackArea`가 `ItemDetailPanel`의 자식에 있었는데, 이 경우 `BackArea`가 디테일 부분도 어둡게 덮는 현상이 있었음. 

- [x] Enemy 충돌 크기 조정
```
ㅁ ㅁ O
X ㅁ
```
- 지금 O에 있는 오퍼레이터가 공격범위 1칸인데도 X에 있는 적을 공격할 수 있음 : 콜라이더가 너무 크기 때문인 듯?
- 캡슐 모델의 스케일이 `0.25, 0.25, 0.25`임. 이거에 맞추면 되는데, 대신 `y`값의 경우는 타일과의 콜라이더 계산도 들어가고 있기 때문에, 콜라이더의 사이즈를 `0.25, 1, 0.25` 정도로 맞추면 될 것 같다.
- 오퍼레이터의 콜라이더도 생각을 해봐야 하겠는데, `0.8, 1, 0.8`이다. 콜라이더로 저지나 공격 로직을 바꿨던 이유가 **두 타일에 걸친 상태**를 구현하기 위함이었는데, 이렇게 해도 크게 상관은 없지 않을까 싶음.
- 테스트해보니 이 정도면 괜찮은 것 같다. 
## 250305
#### 수정 사항
- [x]  `DeployableActionUI`의 SP 상태에 따른 스킬 버튼, SP 수치 표시 등 컨트롤하기
	- 색상 관리도 `ResourceManager`로 빼뒀다. 
		- 예를 들면 SP 게이지에 사용됐던 스킬 사용 중일 때의 색과, 스킬을 사용하지 않는 상태의 색은 `SkillButton` 위에서도 똑같은 색을 사용해야 함 - 같은 상황에서의 게이지의 색을 표시하는 것이기 때문
##### 겹치는 요소 중복 클릭 문제
-  이전과 비슷한 문제 : `DeployableActionUI`의 버튼들과 `Operator`가 겹쳐졌을 때, 버튼 클릭 -> Operator 클릭 동작이 동시에 이뤄지는 문제
	- 저번에 수정한 내용으로, `DeployableActionUI`의 다이아몬드(클릭해도 동작하지 않음)가 `Operator`의 클릭을 빼앗아가는 현상이 있어서 넣은 코드가 있었는데, 얘가 문제를 일으키고 있는 것으로 보임
	- `Button`의 동작은 `EventSystem, GraphicRaycaster`로 동작되는데, 이게 `ClickDetctionSystem`에서 **레이캐스트를 쏘는 로직과는 별도로 동작함.** 따라서, 1번의 클릭만으로 2개의 병렬적인 레이캐스트가 수행되게 되는 것이다.
	- **이상한 현상이 있다.**
		- 일단 레이캐스트 자체는 `RaycastAll`을 쓰고 있기 때문에, 해당 포인터 위치에 있는 **모든** 요소를 가져오는 방식임
		1. 게임 진행 중에 스킬 버튼의 위치를 클릭하면, 레이캐스트에서 감지 못함
		2. **게임이 끝난 상황**에서 `DeployableActionUI`가 떠 있을 때 스킬 버튼의 위치를 클릭하면 **레이캐스트에서 감지**함(???)

- 위 이슈 테스트
1. 일단 스테이지 진행 중일 때의 `DeployableActionUI`의  `Canvas, GraphicRaycaster`을 보면
	- 스테이지 진행 중
		- `Layer : UI_WorldSpace`
		- `Canvas`
			- `World Space`
			- `EventCamera : Main`
			- `Sorting Layer : Default`
			- `Order in Layer : 1`
			- `Additional Shader Channels : TexCoord1, Normal, Tangent`
			- `Vertex Color Always ...` : 체크 해제
		- `Graphic Raycaster`
			- `Ignore Reversed Graphics` : `true`
			- `Blocking Objects` : `None`
			- `Blocking Mask` :  `everything`
	- 스테이지 진행 후 비교
		- 별 차이 없음
2. 버튼 위에 다른 `Screen Overlay` 패널이 씌워진 상태에서 버튼의 위치 클릭
	- 버튼 위에 다른 패널이 없을 때 클릭 : 레이캐스트에 나타나지 않음
	- 버튼 위에 다른 패널이 있을 때 클릭 : 레이캐스트에 나타남
	- 저녁 먹고 왔다. 머리를 식히고 오니 약간 이런 현상이 의심스럽다
		- 패널이 있을 때 클릭한다 = 해당 버튼의 동작이 이뤄지지 않는다
		- 패널이 없는 상태에서 버튼 위치를 클릭했을 때, 버튼이 동작하든 하지 않든 `DeployableActionUI`가 사라진다. 
		- 그러면 **`DeployableActionUI`가 사라지기 전에 버튼의 동작이 실행되고, `DeployableActionUI`가 사라진 다음에 레이캐스트 점검 로직이 실행되는 게 아닐까?**
	- 실제로 `DeployableActionUI`의 버튼 동작과, `ClickDetectionSystem`의 동작을 로깅해보면 아래처럼 나타났다.
```cs
1. Button의 클릭 동작 수행
2. Button이 클릭됐을 때, Hide() 메서드에 의해 DeployableActionUI가 파괴됨
3. ClickDetctionSystem의 동작이 수행됨
4. 결과적으로 Button이 사라진 상태에서 ClickDetectionSystem의 동작이 이뤄지게 되고, Button의 레이캐스트가 이뤄지지 않음
```
> 위의 추측이 맞는 것 같음.
> 그렇다면 버튼에 이뤄지는 레이캐스트가 먼저 이뤄지고, 그 다음에 `ClickDetctionSystem`의 클릭 처리 동작이 작동하는 것으로 보인다. 이걸 어떻게 관리하는지가 중요하겠음.

- 이런 스크립트를 `ClickDetectionSystem`에 만들었는데, 가장 위의 로그도 나타나지 않음. 왜인지는 몰라도 제대로 동작하지 않는 듯.
```cs
    public void OnPointerClick(PointerEventData eventData)
    {
        Debug.Log("OnPointerClick 동작");
        // 클릭된 위치의 가장 위에 있는 UI 요소를 반환함
        GameObject clickedObject = EventSystem.current.currentSelectedGameObject;

        Debug.Log($"{clickedObject.name}");

        if (clickedObject != null)
        {
            Button button = clickedObject.GetComponent<Button>();
            if (button != null)
            {
                DeployableActionUI actionUi = clickedObject.GetComponentInParent<DeployableActionUI>();
                if (actionUi != null)
                {
                    uiClickHandled = true;
                    return;
                }
            }
        }
    }
```

- 그래서 좀 지저분한 해결 방식이기는 한데, `DeployableActionUI`의 버튼 클릭 동작 실행 -> `ClickDetctionSystem`의 상태 변경 -> 이를 감지해서 `HandleClick`이 동작하지 않게 하겠음
```cs
// ClickDetectionSystem.cs
    // 이미 실행된 UI가 있는 경우, 이 스크립트가 동작하지 않아도 되게 함
    public bool buttonClickedThisFrame = false;
    private bool shouldSkipHandleClick = false;

	// ...

    private void Update()
    {
        if (Input.GetMouseButtonDown(0))
        {
            HandleMouseDown();
            shouldSkipHandleClick = false; // 매 프레임 초기화
        }
        if (Input.GetMouseButtonUp(0))
        {
            // UI 클릭이 없었을 때에만 HandleClick 동작
            if (!shouldSkipHandleClick)
            {
                HandleClick();
            }

            // 다음 프레임을 위한 초기화
            buttonClickedThisFrame = false;
            shouldSkipHandleClick = false;
        }
    }

// DeployableActionUI.cs
    private void OnSkillButtonClicked()
    {
        if (deployable is Operator op)
        {
            if (op.CurrentSkill.autoActivate == false)
            {
                op.UseSkill();
                UpdateSkillButton();
            }
            ClickDetectionSystem.Instance.OnButtonClicked();
            Hide();
        }
    }

    private void OnRetreatButtonClicked()
    {
        deployable.Retreat();
        ClickDetectionSystem.Instance.OnButtonClicked();
        Hide();
    }
```
> 스킬 버튼을 클릭했음에도 카메라가 원래 위치로 돌아가지 않음 : 이건 `Hide` 메서드만 수정하면 될 듯

일단 `deployableActionUI`의 어떤 버튼을 클릭했을 때, 그 버튼 밑에 가려진 다른 요소들이 함께 클릭되는 현상은 해결했음. 다른 버튼들에도 구현...하는 게 좋겠지만 크게 필요는 없을 듯. 버튼이 사라지는 경우가 그렇게 많지 않아서..





#### 오늘 해결 안된 문제
- [ ] 스테이지 클리어 실패했는데도 `StageResultPanel`에 보상이 나타나는 문제
- [ ] 뱅가드 스킬 - 코스트 회복 파티클 : 우측 하단에 너무 가까울 경우, 생존 시간이 남아있기 때문인지 파괴됨에도 다시 생성되는 문제
- [ ] ItemUIElement 관련 수정
1. `ItemDetailArea`(설명문 부분) - 너비는 일정하게, 높이는 `Detail` 스크립트에 비례하게 수정하기
	- 아 근데 이거 `ScrollRect`라서 그냥 높이를 수정하는 방법이 훨씬 나을 듯.
2. `ItemDetailPanel`(클릭시 나타나는 아이템의 자세한 내용 패널)
	- 아이템이 나타나는 위치에 따라, 이 디테일 패널이 잘릴 우려가 있음
	- 이 패널의 가장 오른쪽이 스크린의 가장 오른쪽을 초과하면 왼쪽에 나타나도록 구현하기

> 지식이 늘었다
> - `rectTransform.position`과 `rectTransform.anchoredPosition`
> - 전자는 월드 기준, 후자는 로컬 기준임
> - 보통 로컬로 쓸 상황이 많은 듯?

## 250304

#### 자잘?한 수정 사항
- [x] `DeployableBox`에 쿨타임을 기다리는 원형 게이지 추가
![[Pasted image 20250304162840.png]]

- [ ]  `DeployableActionUI`의 SP 상태에 따른 스킬 버튼, SP 수치 표시 등 컨트롤하기
	- 하다가 마무리 못함 -> 내일 계속

#### 기타 이슈 수정
- [x] 1회 퇴각 후에 오퍼레이터를 다시 배치했을 때, `DeployableBox`가 사라지지 않는 현상
	- 오퍼레이터가 배치 가능한지 여부는 `DeployableUnitState`에서 정보를 갖고 있도록 했다. 이는 `DeployableInfo`를 키로 접근이 가능한데, `IsDeployed`로만 박스의 표시 여부를 결정하게 수정했음
	- `DeployableManager`에서 `DeployableBox`의 표시 여부`gameObject.SetActive()`를 결정시키는 대신, `DeployableBox.UpdateDisplay(DeployableUnitState unitState)`로 관련된 로직을 싹 다 옮겼다. 

## 250303

#### 밸런싱 시작
- `1-0`은 이 정도면 된 듯 함

#### 튜토리얼 구현 계획
- 1-0 진입 ~ 레벨업 및 정예화까지
	- 레벨업, 정예화까지 하면 몰려오는 상황에 대응할 수 있는 오퍼레이터를 강제로 1정예화를 만들어줄 수도 있을 듯
	- 가장 유능한 캐릭이 뱅가드 같아서 뱅가드한테 주면 좋지 않을까 싶다.
- 대충 계획
	- 코스트 회복
	- 오퍼레이터 배치법
	- 스테이지 클리어 후 아이템 사용법
#### 스테이지 클리어 후 경험치 아이템 지급하기
- `StageData`에 `ItemData`들을 넣고, 스테이지 클리어 시 유저에게 아이템이 지급됨
- 관련 로직을 구현하고, UI도 `StageResultPanel`에 구현하기

> 여담) 귀찮아서 AI한테 코드 짜달라고 해야지 -> 이제 `Claude`를 쓰고 있지 않아서 관련 정보들을 다 찾아서 전달해줘야 함 -> 관련 정보를 찾다보니 이건 이렇게 하고 저건 저렇게 하면 되겠네?가 계획됨 -> 질문을 안 하고 직접 코드 작성
> - 초반에 유니티에 대해서 잘 모를 때는 그래도 질문을 하는 편이었는데, 최근엔 그래도 익숙해졌음을 느낀다.  
> - 문제는 UI 구현할 때마다 뇌가 리셋됨 ㅋㅋㅋㅋ 그래도 만들어 놓은 게 있어서 그걸 쓰면 되기는 함.

- 이전에 가져왔던 `UniversalBlurUI`의 경우, 캔버스의 다른 요소들을 블러처리해주지는 않는다.(기억력 이슈) 따라서 뒷배경에 클릭해서 패널이 나타난 상태를 취소시키는 버튼을 추가하고 싶다면 그냥 어둡고 알파값을 1 미만으로 놓는 게 차선책으로 보임

- UI 구현하는 게 생각보다 오래 걸렸다. **`StageResultPanel`의 나가는 버튼 기능을 별도의 내부 버튼을 구현해서 그 부분을 빼놓고, 레이아웃 정리했고, `ItemUIElement`의 로비 팝업이랑 스테이지에서 나오는 패널 구분하고, 스테이지의 패널 구현했음.**

#### 기타 이슈 수정
- [x] `EnemySpawner`에서 `EnemyPathIndicator`와 `Enemy`를 구분해서 생성하지만, 어떤 필드를 사용하느냐에 관계 없음. 즉, `EnemyPathIndicator`를 지정했지만 프리팹에는 `enemy`를 할당하는 것도 가능한 상황
	- 에디터를 고치는 것도 가능하지만, 이게 이슈가 될 수 있는 상황은 개발 상황 뿐이기 때문에, `PathIndicator`를 지정했는데 `Enemy`가 나오거나 혹은 그 역인 상황일 때 오류문구를 출력하는 정도로 충분할 듯
	- `EnemySpawner`에서 로깅 처리만 추가함

- [x] 저지한 적을 제거함 -> `Operator`의 콜라이더 내에 충돌 중인 적이 있음에도 해당 적을 저지하지 않는 문제 발생 중
	- `OnTriggerEnter`, `OnTriggerExit` 외에도, 콜라이더가 충돌 중일 때마다 호출되는 `OnTriggerStay(Collider other)`가 있다. 이를 이용함.
```cs
    private void OnTriggerEnter(Collider other)
    {
        ProcessEnemyCollision(other);
    }

    private void OnTriggerStay(Collider other)
    {
        ProcessEnemyCollision(other);
    }

    private void ProcessEnemyCollision(Collider other)
    {
        if (IsDeployed && blockedEnemies.Count < currentStats.MaxBlockableEnemies)
        { 
            Enemy collidedEnemy = other.GetComponent<Enemy>();

            if (collidedEnemy != null &&
                CanBlockEnemy(collidedEnemy.BlockCount) && // 이 오퍼레이터가 이 적을 저지할 수 있을 때 
                collidedEnemy.BlockingOperator == null) // 해당 적을 저지 중인 아군 오퍼레이터가 없을 때 
            {
                BlockEnemy(collidedEnemy); // 적을 저지
                Debug.Log($"{this.BaseData.entityName}이 {collidedEnemy}을 저지하기 시작함, 현재 저지 수 : {blockedEnemies.Count}");
                collidedEnemy.SetBlockingOperator(this);
            }
        }
    }
```
> 로직 자체는 동일하기 때문에 `ProcessEnemyCollision`이라는 별도의 메서드로 빼두고, `OnTriggerEnter`와 `OnTriggerStay`를 수정한다. 

- [x] `Sniper` 에서 맵 범위를 벗어난 곳을 공격 범위로 지정할 경우, `NullReferenceException`이 발생하며 공격을 시행하지 않는 문제
	- 그냥 `Nullable` 추가하니까 잘 됨. 엌ㅋㅋ
```cs
// ? 추가
	Tile? eachTile = MapManager.Instance.GetTile(eachPos.x, eachPos.y);
```
# 이전 일지

## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- 옵시디언 링크
	- [[24년 7월]]
	- [[24년 8월]]
	- [[24년 9월]]
	- [[24년 10월]]
	- [[24년 11월]]
	- [[24년 12월]]
	- [[25년 1월]]
	- [[25년 2월]]