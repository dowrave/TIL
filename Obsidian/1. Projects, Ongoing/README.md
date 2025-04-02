## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 데이터가 제때 수집되고 있는지 눈팅 정도만 하면 충분할 것 같다.
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음?

## 작업 예정

### 남은 작업 내용 
- 남은 작업들
	- 스테이지 1-0 ~ 1-3 밸런싱
	- 1-3에 보스 추가
	- 보상 설정

- 테스트 및 수정

### 작업 중 
1. 스테이지 `1-1` 밸런싱 (이후 `1-2`, `1-3`)
	- **스테이지 완성, 보상 설정, 보스 추가** 등으로 게임을 완성하는 게 젤 중요함!! 다른 건 다 부차적인 요소! --> 잘 실천이 안되고 있음
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


## 4월

## 250402 - 짭명방

### 뭔가 허전한 요소들 채우는 중 
- `StageResultPanel` - `StatsContainer`에 제목 텍스트 추가
- `StageSelectPanel` - 스테이지 클릭 시 나타나는 우측 패널에 `StageName`도 추가
- `SkillData`로 생성한 스킬들의 `Description` 정리 
- `OperatorDetailPanel`
	- 스킬 이름 표시 : 스킬 아이콘 아래에
	- 스킬 아이콘 클릭 동작 변경
		- 현재 동작) 스킬 아이콘 클릭 시 - 즉시 해당 스킬을 디폴트 스킬로 지정
		- 하지만 스킬 아이콘을 클릭하는 이유 중에는 단순히 그 스킬에 대한 설명을 보고 싶기 때문일 수도 있다
		- 따라서 원본 게임처럼 **스킬의 설명 아래에 `디폴트로 지정`하는 별도의 버튼을 하나 추가**함
	- 원본 반영하기 : 스킬 아이콘 우측 하단에 지속시간 항목 추가
		- 소모 SP는 표시하지 않음
		- 이거를 구현할거면 아예 `SkillIcon`이라는 항목을 따로 만드는 것도 방법이겠는데?
	- `SkillIconBox.cs`을 따로 만들었음. 
		- `SkillSelectionIndicator`는 아예 패널의 요소로 빼고, 어떤 스킬을 선택하면 해당 `SkillIconBox`의 1번째 자식으로 배치시킴(지속 시간 박스와 겹치면 보이지 않게 렌더링)
		- 이를 사용하는 코드랑 레이아웃도 일부 수정. 
		- 원본 게임은 지속 시간 아이콘 옆에 숫자만 넣었는데, 나는 `s`까지 넣어서 뭘 나타내는지 조금 더 뚜렷하게 해보겠음
		- 스킬 아이콘이 들어가는 요소들을 다 이걸로 대체해야 하나? `OperatorInventoryPanel`까지는 스킬 지속시간이 보이는 게 좋아보이니 수정하고 마무리. 
		- `OperatorSlot`이나 인게임 스킬 아이콘에서도 넣어야 할까? 굳이 그럴 필요는 없을 듯?

- `OperatorDetailPanel`에 넣음
![[Pasted image 20250402220917.png]]

- `OperatorInventoryPanel`에 반영
![[Pasted image 20250403002824.png]]

### 지식이 늘었다
- `Content Size Fitter` 등에 의한 너비/높이 변화 즉시 반영하기
```cs
// ContentSizeFitter에 너비 변화 반영
LayoutRebuilder.ForceRebuildLayoutImmediate(durationBox.rectTransform);
```
> 관련된 항목을 설정 `ex) 텍스트 할당` 하고, 레이아웃을 즉시 업데이트를 한 번 해준다. 

- **`OperatorInventoryPanel`을 활성화한 채로 테스트에 들어갈 때, `UIHelper`보다 `Awake` 메서드가 먼저 동작**하는 이슈
	- 일단 `UIHelper`를 이용하는 메서드와 `UIHelper`의 싱글턴 인스턴스를 설정하는 메서드 모두 `Awake`에 있음
	- **2개의 활성화된 오브젝트의 `Awake`의 순서는 명확히 보장되지 않음**
		1. `Project Settings - Script Execution Order`에서 설정한 순서
		2. 하이어라키에서의 위치 : 상단의 오브젝트가 먼저 실행되는 "경향"이 있음
		3. 오브젝트 생성 시점 : 먼저 생성된 오브젝트가 먼저 실행되는 "경향"이 있음
	- 비활성화된 경우는 호출되지 않다가, 활성화될 때 1회만 호출됨. 재활성화 시에는 호출되지 않음.
	- 일단 이 문제는 해당 메서드를 `OnEnable`으로 빼는 것으로 바꿨음. 패널이 켜질 때마다 실행되긴 해서 좋은 건 아니겠지만..



## 250401 - 짭명방

### OperatorPromotionPanel 
- 좌측 하단에 사용 아이템 표시
	- 처음에는 아이템이 들어가는 부분에 대한 이미지를 검정 -> 투명 그라데이션 으로 넣어보려고 했는데, 그냥 아무 것도 없는 곳에 아이템이 나타나게 하는 것도 괜찮은 듯?

- 추가로 `OperatorData`에 `Promotion`에 필요한 아이템을 추가함
	- 실제로 정예화에 `PromotionItem 1개` 외의 아이템을 추가할 생각은 없어서 이렇게 작업할 필요는 없다. `OperatorPromotionPanel`의 좌측 하단을 고정시키면 되니까
	- 그래도 개발을 이것만 하고 말 게 아니니까 실제로 데이터를 가져와서 나타내는 것까지는 작업을 해야 공부가 될 것 같다. 
```cs
// OperatorData
    [Header("Promotion Required Items")]
    public List<PromotionItems> itemDatas = default!; 

    [System.Serializable]
    public class PromotionItems
    {
        public ItemData itemData;
        public int count;
    }
```

- 이렇게 수정한 부분을 바탕으로 아래의 과정들이 진행됨
1. `OperatorPromotionPanel`의 좌측 하단에 `ItemUIElement`을 초기화하는 것
	- 처음에는 인스턴스화 -> 파괴를 반복했는데, 그냥 미리 `ItemUIElement`들을 넣어놓고 필요한 만큼만 활성화를 하는 방식으로 수정.
2. 정예화가 진행되는 로직도 `PromotionItems`에 있는 아이템들을 사용하게 만들기

- 코파일럿) 이 때 아이템을 갖고 있는지를 점검하는 메서드는 `PlayerDataManager`에 구현하는 게 더 맞다고 한다. 정예화 시에 갖고 있는 아이템을 검사하는 거니까 `OperatorGrowthSystem`이 맞을 것 같았는데, 아이템을 정리하는 로직 자체가 `PlayerDataManager`에 있기 때문에 거기서 진행하는 게 맞다고 함. 
	- 클로드한테도 물어보니까 비슷한 답변이 돌아왔음. 재활용한 코드가 갯수는 반영되지 않았었는데, 이것도 고쳐줬다. 무서운 시대.

### 지식이 늘었다
1. `IEnumerable.ToDictionary`라는 메서드가 있다. 컬렉션을 딕셔너리로 바꿀 때 유용할 것 같음.
```cs
var itemsToUse = op.OperatorProgressData.promotionItems.ToDictionary(
	promotionItem => promotionItem.itemData.itemName,
	promotionItem => promotionItem.count);
```
> `promotionItems`가 더 여러 개의 필드를 갖고 있더라도, 그 중에서 2개의 필드만 뽑아서 키, 밸류 쌍으로 묶을 수 있다고 한다. 

2. `ItemUIElement`들을 생성 -> 파괴에서 활성화 -> 비활성화로 바꿨다. `ItemUIElement`들을 리스트로 관리하고 관련 오브젝트들을 인스펙터에서 집어넣는 방식으로 구현하고 있다.
- 이 때, **인스펙터에서 리스트의 길이를 5로 하고 일부 값이 비어있다면 리스트의 길이는 여전히 5이다.** 즉,  내부 원소가 null이더라도 리스트는 null인 원소를 갖고 있는 상태가 된다. 이거는 값이 들어갔다가 나간 상태인 `Missing`도 마찬가지.
> 활성화 - 비활성화로 오브젝트들을 다룰 때, 개발 과정에서 일부 오브젝트를 만들거나 없애는 경우에 주의해야 할 듯. 오늘 같은 경우도 5개를 쓰려다가 3개로 수정했는데 리스트는 (지금 보면 당연히) 길이 5로 유지됨.

# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `Archive` 폴더로 들어갔을 듯.
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

## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
