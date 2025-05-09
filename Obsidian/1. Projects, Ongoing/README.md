## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음? `Quill-Markdown`하고 충돌하고 있는 느낌은 있다.

## 작업 예정 - 짭명방

### 남은 작업 내용 
- 남은 작업들
	- 스테이지 1-0 ~ 1-3 밸런싱
	- 1-3에 보스 추가
	- 보상 설정

- 보고 참조할 내용
	- [[오퍼레이터들 스탯 정리]]
	- [[적 스탯 정리]]

- 테스트 및 수정

### 작업 중 
1. 스테이지 `1-1` 밸런싱 (이후 `1-2`, `1-3`)
	- **스테이지 완성, 보상 설정, 보스 추가** 등으로 게임을 완성하는 게 젤 중요함!! 다른 건 다 부차적인 요소! 
	- ...인데 잘 실천이 안 되고 있다. 계속 고칠 점이 보임...

### 구현 예정

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다 (= 오퍼레이터를 배치하지 않고 스테이지 클리어하기)
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 드는 품을 생각하면 미뤄도 될 것 같음. 

## 5월

### 못 끝낸 일

1. `SquadEditPanel`에서의 스쿼드 초기화 버튼 기능 추가

## 250509

### 짭명방 : 스쿼드에 오퍼레이터 한꺼번에 배치

####  1.  `Bulk`에서 `SetEmpty` 클릭 시 `tempSquad, clickedSlot` 비우고 UI도 초기화
> 일단 `tempSquad`와 `clickedSlot`의 `HandleSlotClicked` 내부에서 이뤄지므로, 굳이 해당 부분을 호출해서 적용할 필요는 없다. 단순히 `SetSelected(bool)`만으로 기능시킬 수 있음

- `SetEmpty` 구현 완료. 반복문으로 슬롯 돌면서 `SetSelected(false)`만 구현했음.
- `nowEditingIndex` 관리
	- **클릭된 시점에 `nowEditingIndex`가 설정되어야 함** 
		- 어제는 슬롯을 클릭하면 다음 인덱스를 미리 갖게 구현했었음. 이러면 스킬 설정할 때 `null`인 공간에 사용할 스킬 인덱스만 들어가게 된다.

#### 2. `Bulk`에서 클릭된 `OperatorSlot`의 우측 상단에 번호 표시(현재 인덱스 + 1)
- 구현 완료
![[Pasted image 20250509172331.png]]
- 이거 하려고 며칠을 쓴 거임 아 ㅋㅋㅋ
- 아예 **스쿼드 구조를 고쳤고, 관련된 메서드들을 고쳤고, `OwnedOperator`와 `Squad` 간의 조율 때문에 시간을 오래 잡아먹었다. 근데 바꾸는 것 자체는 잘 한 것 같음.**

### Bulk 인벤토리 접근 버튼 구현
> 원본 명방을 보면..
> 1. 기본적으로 아이콘이 있는 버튼으로 나타남
> 2. 이 아이콘이 있는 버튼을 클릭하면, 버튼의 오른쪽 끝 부분에서 시작해서 우측으로 텍스트 박스가 늘어남. 결과적으로 기존 버튼보다 넓은 너비를 갖게 됨.
> 3. 2번의 상태에서 클릭하면 해당 패널로 이동함

- 지금은 패널 이동 기능만 구현되어 있는데, 이걸 바꿔보자.
	- `ExpandableButton`이라는 스크립트를 하나 만듦
	- 1번째 클릭은 확장, 2번째 클릭은 버튼 동작으로 이어짐
		- 여기서 버튼 동작은 (처음으로) `UnityEvent`를 써봤다. `public` 메서드가 있는 오브젝트를 연결해서 인스펙터에서 직접 붙이는 방식임.
		- 기존에는 `Event<Action>`을 이용했음 (C#)

- 일단 2개의 `ExpandableButton`을 담는 컨테이너를 하나 만들었음. 이벤트도 1번째 클릭과 2번째 클릭 때 다른 이벤트가 발생하도록 했다.
- `MainMenuManager`에서의 1번째 클릭 시 발생하는 이벤트는 반대쪽 버튼을 축소하는 메서드를 구독시켰고, 2번째 클릭은 해당 버튼의 본격적인 기능이 됨

- 지금은 각 버튼이 늘어나거나 줄어들 때 너비값이 정상적으로 반영되지 않는 문제가 있음. 
```
ExpandableButton
- BasicIcon
- ExpandableBox
```
> `Horizontal Layout Group`을 `Content Size Fitter`와 함께 써야 자식 오브젝트들의 너비 요소들이 부모 오브젝트에 제대로 반영됨

- 버튼이 늘어나거나 줄어들 때 다른 버튼의 위치가 꼼지락대는 이슈도 있었다. 부모 오브젝트의 `Horizontal Layout Group`을 제거했음.

- 버튼이 초기에 등장했을 때에는 애니메이션 동작을 막음
	- 버튼에 `isInitializing`이라는 플래그를 넣고, `SetIsExpanded`로 동작시킴
```cs
    public void SetIsExpanded(bool state, bool isInitializing = false)
    {
        IsExpanded = state;
        this.isInitializing = isInitializing;
    }

// 아래처럼 사용
squadBulkEditButton.SetIsExpanded(true, isInitializing: false);
```

- 초기화시킨 다음 클릭은 멀쩡하게 애니메이션으로 동작해야 하는데 그렇지 못한 현상이 있다. 
	- 디버깅해보면 정확히 같은 루트로 호출이 2번 되는 현상이 있음. 2번째 호출의 경우, 버튼의 동작이 정상적으로 진행되지 못하고 멈춘다.
	- 완전히 같은 경로로 2번 호출되는 게 힌트였음. `StageSelectPanel`에 `OnConfirmButtonClicked` 이벤트가 2번 등록되어서 그랬다. 설마 이벤트 등록 2번 했나? 했는데 진짜였다 ㅋㅋㅋㅋㅋ
	- 그 다음에도 바로 동작 안하는 문제가 있어서 수정 완료.


## 250508
### 어제 못한 것들에 대한 계획
##### 스킬 관련 설정
- 지정된 여러 오퍼레이터 외에도 스킬 설정도 함께 들어가 있어야 함.
	- 기존에는 `SelectedSkill` 1개만 관리되고 있었는데, 이걸 `Operator + Skill`의 조합으로 넣어야 하나? 아니면 사이드 패널에 활성화된 상태에서 스킬 클릭했을 때 해당 OwnedOperator의 `StageSelectedSkill`만 바로 바뀌면 되나?
	- 일단 후자의 방법을 시도해봄 

##### 수정 필요
- 일괄 편집에 들어갔을 때, 슬롯이 선택된 걸 취소하고 다시 한 번 클릭했을 때 왜인지 모를 이유로 클릭이 한 번 씹히는 현상이 있음. 그 다음부턴 괜찮지만 거슬린다. 상태 관련 이슈 같음.

##### OnConfirmButtonClicked - 벌크 구현
- 기존 스쿼드가 저장되는 방식을 보면, 오퍼레이터 하나하나에 대해서
	1. 스킬 버튼이 클릭됐을 때 해당 오퍼레이터의 `StageSelectedSkill`이 변경됨
	2. `UserSquadManager`에서는 해당 오퍼레이터 정보만 들어감. 스쿼드에서 사용하는 스킬은 `StageSelectedSkill`을 따라간다. 이는 `DefaultSelectedSkill`과 구분됨.

> - 실제 명방에서는 여러 개의 스쿼드를 쓸 수 있기 때문에, 스쿼드 자체에서 오퍼레이터들과 스킬을 함께 관리하는 방식으로 구현이 되어 있을 거임. 
>   - 그렇게 추측이 된다고 했을 때, 내 구현을 바꿀까 말까에 대한 생각은 든다. 지금처럼 구현해도 큰 상관은 없겠지만.

- 벌크에서도 다르게 구현할 이유가 있을까? 위와 마찬가지로 스쿼드에 넣으면 될 것 같음
	- 대신 `UserSquadManager`에서 `List<OwnedOperator>`를 통째로 받아서 저장하는 로직 하나는 추가해야 한다. 기존엔 저렇게 넣는 식으로 구현한 게 없었..나? 저장 로직은 `PlayerDataManager`에 있기 때문에, 이것까지도 함께 봐야 할 것 같다. 

### 계획 실행

#### 스쿼드 구조 변경
- 기존 스쿼드는 `operatorName`만을 관리하고, 스킬은 각 `operatorName`을 이용해 `OwnedOperator`에서 `squadSelectedSkill`을 가져오는 방식이었음
- (물론 내 프로젝트에서는 1개의 스쿼드만 두겠지만) 이렇게 구현하는 경우 다른 스쿼드에 두더라도 모든 스킬이 통일되는 문제가 발생함
	- 직관성이 떨어진다고 생각해서 시작함. "스쿼드에서 사용하는 오퍼레이터의 스킬"이라는 개념을 어디에 둘 지에 대한 것이다. 
- 그래서 **스쿼드 관리 자체를 `operatorName, skillIndex`의 형태로 저장**하는 방식으로 수정함.

```cs
    // 플레이어가 소유한 데이터 정보
    [Serializable] 
    private class PlayerData
    {
        //public List<string> currentSquadOperatorNames = new List<string>(); // 스쿼드, 직렬화를 위해 이름만 저장
        public List<SquadOperatorInfo> currentSquad = new List<SquadOperatorInfo>();
    }

    [Serializable]
    public class SquadOperatorInfo
    {
        public string operatorName;
        public int skillIndex;
        
        // 생성자
        public SquadOperatorInfo()
        {
            this.operatorName = string.Empty;
            this.skillIndex = 0;
        }

        public SquadOperatorInfo(string name, int index)
        {
            operatorName = name;
            skillIndex = index; 
        }
    }
```

으로 시작, 관련 메서드들을 전부 수정하는 작업을 진행

##### OwnedOperator 부분 수정
`OwnedOperator`에서 현재 선택된 스킬을 어떤 식으로 관리할지가 고민이다.  다시 써봄.

1. 스쿼드에서 사용 중인 스킬 인덱스를 갖게 하겠다면, `OwnedOperator` 자체에서는 `UnlockedSkills`가 있으니까 아무 정보도 갖지 않아도 된다?
2. `defaultSkillIndex`는 `OwnedOperator`에서 갖는 게 맞음. 
3. 기존에 스킬 정보는 `DeployableInfo.OwnedOperator.SelectedSkill` 같은 식으로 전달이 됐을 거임. 그 값이 이젠 `Squad`로 나갔기 때문에 `Squad`에서 선택 중인 스킬 인덱스 값을 가져와야 함.
4. 기존 스쿼드 정보는 `UserSquadManager.GetCurrentSquad()`로 가져왔었다. 
5. `GetCurrentSquad()`는 `List<string: operatorNames>`을 `List<OwnedOperator>`로 바꾸는 로직이었음. 즉, 이번에는 `OwnedOperator, int`를 함께 갖도록 반환시켜야 한다.
	- `GetCurrentSquad, GetCurrentSquadWithNull` 수정 : [[스쿼드 정보 얻어오기 로직 변경]]

- 얼추 정리되고 있다. 스쿼드 여러 개 구현할 것도 아닌데 괜히 했다는 생각이 들지만, 시작했으니까 어쩔 수 없다. (사실 스쿼드가 여러 개여도 OwnedOperator에서 여러 스쿼드에 대한 필드를 따로 구성하면 훨씬 간단하게 구현할 수 있기도 했고.) 기존 방식이 더 나았던 것 같기도..?
- 근데 이 방식이 맞다고 생각했으니까 가야지 뭐..

- 계속 수정 중
	- `TryUpdateSquad` : 중복된 이름의 오퍼레이터를 스쿼드에 넣는 경우는 동작하지 않도록 했으나, 같은 오퍼레이터를 스킬만 바꾸고 싶은 경우에는 동작해야 함
```cs
// 마지막 조건만 새로 추가
if (squadForSave.Any(opInfo => opInfo != null && opInfo.operatorName == operatorName && opInfo.skillIndex == skillIndex)) return false;
```

- 일단 여기까지 구현하면 `OperatorInventoryPanel`에서 오퍼레이터를 "하나씩" 수정하는 케이스는 잘 동작함
	- 위에서 기존 방식이 더 나았던 것 같다고 했는데, 밑작업이 복잡해지면 위에서 작업할 때는 상대적으로 수월해지는 면이 있는 것 같음.

---
- 이제 최근에 구현했던 **한꺼번에 스쿼드에 넣기**를 해봄
	- `PlayerDataManager`에 `UpdateFullSquad`을 구현
	- `OperatorInventoryPanel`에서 단일 슬롯 / 벌크에 따른 인덱스 관리 분리
	- 인덱스 정리가 난리가 났다.
		- `slot.SetSelected()` 부분이랑 이걸로 인해 `HandleSlotClickedForBulk` 이벤트가 발생하는 이슈가 있음. 이전에도 재귀처리 때문에 다뤘던 부분인데, 여기서도 문제가 된다.
		- 요점은 초기화 상황에서는 `HandleSlotClicked`의 동작을 방지하면 되지 않을까?
		- 그러면 `OperatorInventoryPanel`에 `isInitializing` 전역 필드를 하나 추가함
			- 굳. `OnEnable`의 맨 처음과 끝에 필드 두고 `HandleSlotClickedForBulk`에서만 플래그로 사용함.

- 일단 거의 다 오기는 했음
	- 나머지 할 일
		1.  `Bulk`에서 `SetEmpty` 클릭 시 `tempSquad, clickedSlot` 비우고 UI도 초기화
		2. `Bulk`에서 클릭된 `OperatorSlot`의 우측 상단에 번호 표시(현재 인덱스 + 1)


### 지식이 늘었다 

#### LinQ - Select, Where, Any, Contains

1. `Select`
	- 컬렉션의 각 요소를 새로운 형태로 변환한다.
```cs
List<int> numbers = new List<int> {1, 2, 3, 4, 5};
IEnumerable<int> squares = numbers.Select(n => n*n) // {1, 4, 9, 16, 25}
```

2. `Where`
	- 컬렉션에서 특정 조건을 만족하는 요소들만 필터링한다.
```cs
List<int> numbers = new List<int> {1, 2, 3, 4, 5};
IEnumerable<int> evenNumbers = numbers.Where(n => n % 2 == 0); //{2, 4}
```

3. `Any`
	- 컬렉션에서 특정 조건을 만족하는 요소가 하나라도 있는지 확인
	- 매개변수가 없을 때에는 빈 컬렉션이 아니라면 `true`
	- 조건자`Predicate`와 함께 사용할 때는 조건자가 `true`인 게 하나라도 있으면 `true`
```cs
List<int> numbers = new List<int> {1, 2, 3, 4, 5};
bool hasEvenNumber = numbers.Any(n => n %2 == 0); // 조건자가 있는 경우 - true
bool isEmpty = new List<int>.Any(); // 리스트가 비어 있는지 확인(.Any()는 요소가 있으면 true)
```

> 내 경우 `List<string>`을 `List<squadOperatorInfo>`로 바꾸면서 `Contains`을 `Any`로 바꿔야 하는 상황이 발생했는데, 각 요소를 돌면서 특정 값을 비교했던 기존 상황과 달리 특정 필드의 값을 비교하는 방식으로 변경되었기 때문이다.
```cs
// 기존 방식. List<string>이라서 Contains을 사용
safePlayerData.currentSquadOperatorNames.Contains(operatorName)

// 수정된 방식. List<squadOperatorInfo>라서 Any를 사용. squadOperatorInfo 내에 opereatorName이라는 필드가 있다.
safePlayerData.currentSquad.Any(opInfo => opInfo != null && opInfo.operatorName == operatorName)
```
## 250507
저런! 어제 축구를 하다 다쳐서 당분간 축구를 못할 것 같다!
### 스쿼드 일괄 편집 기능 추가(계속)

#### 대충 개요 정리
- 일단 원하는 기능부터 정리해보자. 그냥 진행하려니 복잡해서 계속 정리가 안된다.
```
- SquadEditPanel에서 별도의 일괄 편집 버튼으로 진입

- OperatorInventoryPanel에 진입하면
	1. 기존 스쿼드에 있는 오퍼레이터들이 선택된 상태
		- 스쿼드는 중간에 null을 포함할 수 있는 리스트임
			- 차이) 원작 게임은 어떤 요소가 빠지면 나머지 요소가 앞으로 하나씩 밀림
		- 선택되었을 때, "선택된 상태" 및 "해당 오퍼레이터가 스쿼드에서 차지하고 있는 인덱스 + 1"이 함께 나타나야 함
	2. 오퍼레이터 슬롯을 선택했을 때
		- 이미 선택된 오퍼레이터 슬롯, 즉 스쿼드 내에 있는 오퍼레이터 슬롯이라면 제거
			- 해당 인덱스만 비우고, 현재 편집 중인 인덱스를 비워진 인덱스로 설정
		- 스쿼드 내에 없다면 새로 추가
			- 다음 인덱스를 검사해서 이미 채워진 슬롯이라면 인덱스 +1을 반복.
		- SideView는 현재 클릭된 오퍼레이터의 정보가 나타나야 함. 이건 Slot만 전달하면 됨.
		- 
```

#### 스킬 관련 설정
- 지정된 여러 오퍼레이터 외에도 스킬 설정도 함께 들어가 있어야 함.
	- 기존에는 `SelectedSkill` 1개만 관리되고 있었는데, 이걸 `Operator + Skill`의 조합으로 넣어야 하나? 아니면 사이드 패널에 활성화된 상태에서 스킬 클릭했을 때 해당 OwnedOperator의 `StageSelectedSkill`만 바로 바뀌면 되나?
	- 일단 후자의 방법을 시도해봄 

#### 슬롯 관련 프로퍼티로 변경
- `SelectedSlot`이라는 상태의 업데이트 로직과 사이드 뷰의 업데이트 로직이 분리되어 있어서 프로퍼티로 구현해서 하나로 합침
```cs
    public OperatorSlot? SelectedSlot
    {
        get { return selectedSlot; }
        set
        {
            if (selectedSlot != value)
            {
                selectedSlot = value;
                UpdateSideView();
            }
        }
    }
```

#### 기존 문제 해결
- 기존의 클릭된 슬롯 재클릭시 스택 오버플로우 이슈가 있었는데, 아래 로직이었기 때문에 그랬다. 예전엔 미봉책으로만 뒀는데, 이번에도 비슷한 문제가 발생해서 아예 뜯어봄.
```
1. 패널의 HandleSlotClicked는 슬롯의 OnSlotClicked 이벤트를 구독
2. 슬롯의 SetSelected은 슬롯이 선택된 UI 설정 및 OnSlotClicked 이벤트를 발생
3. 패널의 HandleSlotClicked 내에 슬롯.SetSelected()을 넣으면 함수 무한 호출 & 스택 오버플로우 발생

-- 내 경우는 스택 오버플로우가 한 번 뜨고 나면 유니티 엔진 자체가 이상해지는 이슈도 생겨서 프로젝트를 아예 껐다가 켜야 했음. 인스트럭션이 화면에 막 나타남;
```
> 따라서 `HandleSlotClicked` 내부에 슬롯 클릭 처리를 진행 중인지 아닌지를 나타내는 `boolean` 플래그 하나를 추가해서 수정함. 재귀 호출이 발생하더라도 `boolean` 플래그가 있기 때문에 중복 실행을 방지함.


## 250506

### 스쿼드 일괄 편집 기능 추가(미완)
- 솔직히 내 프로젝트는 한 번 스쿼드 편집하면 더 이상 편집할 이유가 없어서 굳이 구현할 필요는 없지만, 그래도 이런 걸 해봐야 한다고 생각해서 진행한다.
- `SquadEditPanel`에서만 활성화되는 `TopArea`에 `squadBulkEditButton`을 추가
	- 기존엔 `isEditing`으로만 분기를 정했는데, 여기에 하나 더 추가해야 함
	- "현재 편집 중인가"에 대한 상태는 `UserSquadManager`에서 관리하고 있었기 때문에, 여기에 분기 하나를 더 추가함
	- 기존 `OperatorInventoryPanel`에 벌크 편집 기능을 추가하려니까 조건 수정을 어디서부터 들어가야 할지 모르겠어서 헤매는 중. 범위가 좀 넓다?
	- 기존 함수들 정리하고, 초기화까지는 했는데 아직 버튼 클릭 시의 동작이랑 확인 버튼 동작 등을 구현하지는 못했음. 스쿼드에서 인덱스 관리하는 것도 생각해봐야 할 듯.

## 250502

### 깃허브 자기소개 레포지토리 편집
- 정리를 좀 했음.
### 짭명방

#### 밸런싱 계획


#### 이슈 수정
- 다시 돌아왔다. 못 해결했던 문제부터 다시 작업
- [x]  `SlashSkill`
    - 지금은 대미지 판정이 이펙트 기준인데, 이걸 타일 기준으로 변경해야 할 것 같음. 타일에 적이 있는데도 이펙트의 경로에서 벗어나면 맞지 않는 문제가 있다.
	- 기존의 메인 이펙트 위치에 콜라이더를 생성하는 방식에서, 파티클 시스템에 있는 `Collide`를 이용해서 `OnParticleCollision`으로 판정하는 걸 시도해봤는데 `Enemy`에 있는 `Is Trigger` 때문에 트리거 관련 처리로 바꿔서 진행해봤다. 그걸 `OnParticleTrigger`로 바꿔봤는데 근데 잘 동작하지 않았음.
	- 그래서 **그냥 기존의 콜라이더 생성 로직에서 반경만 넓히는 방식으로 수정**했다.`0.25f -> 0.5f`
		- 일단은 성능 상 큰 이슈는 없으니까 이렇게 진행하겠음.

- [x]  로비에서 `ItemUIElement` 클릭 시 나타나는 패널에서 아이템 갯수 어떻게 표현할 것인가
    - 현재 보유 중인 아이템 숫자 말고 있나?
	- 나타나는 패널에서 아이템 갯수 표시 부분만 `갯수 -> 보유`로 수정
- [x] `OperatorInventoryPanel`
	- 슬롯이 정상적으로 초기화되지 않는 문제 수정
		- 스킬 슬롯 지정하는 부분에서 바로 슬롯으로 들어가는 부분에 대한 분기 처리가 안 되었다.
- [x] 2성 클리어 후 3성 클리어를 했는데 `PromotionItem`이 지급되지 않는 현상
	- 정예화 지급 아이템 조건 수정



#### 엑셀 스탯 계산기 작업
- [스탯 계산기](https://docs.google.com/spreadsheets/d/1ToqWbSyobNPP6mjlOIRBKZMIE65FfH6BYlm0ffL4wCE/edit?usp=sharing)
- 밸런싱을 위해 + 게임을 켜지 않고 직관적으로 보기 위해 구글 스프레드 시트로 오퍼레이터의 스탯 계산기를 하나 만듦



## 250501
### 컴퓨터 공부
- [[16. OS의 내부 동작 원리]]

# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
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
	- [[25년 3월]]
	- [[1. Projects, Ongoing/유니티 - 작은 명방 구현하기/작업 일지/직접 작성/25년 4월|25년 4월]]

## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
	- [[1. Projects, Ongoing/포폴 겸 블로그 만들기/월별 작업 기록/25년 4월|25년 4월]]
