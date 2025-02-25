# 일지
## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 데이터가 제때 수집되고 있는지 눈팅 정도만 하면 충분할 것 같다.
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음?

## 작업 예정


### 진행 중
- 스테이지 1-1 ~ 1-3 밸런싱
- 튜토리얼?
### 이슈 : 해결 필요



### 구현 예정
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다
	- 혹은 오퍼레이터를 배치하지 않고 3성 클리어하기
	- 이거는 명방 본게임의 4-4에서 쇼만 배치하고 몹들의 동선을 꼬는 방식으로 

### 발생 중인 이슈
- (241216) 레벨업 후에 `OperatorLevelUpPanel`의 스냅핑 동작이 정상적으로 동작하지 않을 때가 있다. 
- 간헐적인 문제
	-  `Enemy` 기준, 바리케이드 파괴 로직이 동작하지 않는 현상이 있음

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다.

# 2월

## 250225
### 짭명방

#### 맵 하나 더 생성
- `1-0`이라는 맵을 하나 만듦. 튜토리얼 느낌의 맵. (근데 튜토리얼을 넣을지 말지는 고민)
- 경로 2개 생성 : 직진으로 달리는 경로 하나, 빙글빙글 돌아서 가는 경로 하나.
	- 경로 테스트도 완료
	- 
#### 구조 수정
- `Map` 스크립트 수정
	- 특정 맵에서만 사용할 수 있는 `MapDeployable` 필드의 위치 수정
		- 기존엔 `Map` 프리팹(스크립트) 내에 저장을 해뒀는데, 이걸 `StageData`로 뺴두는 게 맞지 않나라는 생각이 들어서 그렇게 진행해둠.
	- `tilePrefab`의 위치도 `Map`에서 `TileData`로 옮김

#### 이슈 수정
- [x] `Enemy`가 빈 타일을 지나가고 있는 상황 -> 오퍼레이터 배치 -> 저지가 동작하지 않는 현상
	- 오퍼레이터가 배치된 다음에 적이 들어오면 저지가 잘 동작함
	- 미리보기 중에 **콜라이더가 활성화되는 것을 끄면 해결되지 않을까?** 라는 생각이 들었다. 
		- 디테일) `Operator`에는 `Box Collider` 가 있는데, 이게 미리보기 및 스냅핑 중에 `Enemy`와 충돌했을 때 `OnTriggerEnter`가 동작하지만, 스냅핑된 위치에 그대로 배치가 됐을 때는 이미 두 개의 콜라이더가 충돌된 상태에서 `Operator`의 상태`IsDeployed`만 바뀌기 때문에, 새롭게 콜라이더가 들어온 상태는 감지하지 못하는 게 아닐까?
	- 구조 생각)
		1. `UnitEntity`에서 `BoxCollider`를 관리 : `Enemy`를 포함한 모든 자식 클래스에서 박 콜라이더를 갖고 있다고 봐야 함
		2. `IsDeployed`라는 상태는 `DeployableUnitEntity`부터 가지는 상태이다.
			- `IsDeployed` 프로퍼티의 세터에서 `BoxCollider`의 활성화 여부를 결정하는 메서드를 넣는 것도 고려해봤는데, (실제로 배치 여부에 따라 박스 콜라이더 활성화가 결정이 되기 때문에) 이거는 **기능을 숨기는** 것이기는 해서 메서드의 동작을 볼 때 직관성이 떨어지는 면이 있기는 함
			- 어차피 `IsDeployed`나 `Initialize`에서 박스 콜라이더의 활성화 여부가 결정되기 때문에 프로퍼티에 작성하지 않고 메서드에 작성해도 크게 상관은 없지 않을까? 라는 생각이다.
	- 최종)
		- `UnitEntity`에서 `BoxCollider` 및 `SetColliderState` 메서드 생성
			- `SetColliderState`는 비어 있는 `virtual`로 선언. `abstract`가 아닌 이유는 `deployableUnitEntity`에서 오버라이드할 건데, 저 클래스는 추상 클래스이기 때문이다. `abstract`로 선언하면 구체적으로 구현해야 하는 클래스인 `Operator`, `Barricade` 등에서 별도로 구현해야 함
		- `SetColliderState` 
			- `Enemy`에서는 항상 콜라이더를 켬. 꺼져야 하는 경우가 아직까지는 없는 듯?
			- `deployableUnitEntity`에서는 `IsDeployed`를 따라가도록 메서드를 구현하고, `SetDeployState`에 함께 들어가도록 함

- [x] `Vanguard` - `CostParticleMotion`에서 `private RectTransform deploymentCostIconTransform;`을 찾지 못하는 현상
	- **프리팹에서는 하이어라키에 있는 특정 필드를 할당할 수 없음.** 당연하다. 독립적이니까.
	- 따라서 저 필드를 할당하는 시점 자체는 코스트 파티클이 생성되는 시점이어야만 함. 따라서 `SerializeField`로 설정하는 것도 의미가 없음.
	- 문제 자체는 어제 남은 배치수 패널을 구현하는 과정에서 오브젝트의 이름을 바꿨기 때문에 발생한 것으로 보임.

## 250224

### 짭명방
#### 게임 로직 수정
- [x] 편성된 모든 오퍼레이터를 배치할 수 없게 만들기 
	- 편성은 항상 6명이니까 5명 배치 가능 같은 식으로 변경하겠음
	- 전원을 배치할 수 있는 건 제한적인 상황을 구현하기 어렵기 때문에 이렇게 바꾸는 게 맞는 것 같음. 자원을 한정시켜야 선택하는 재미가 늘기 때문.
	- `StageData`에 해당 정보를 넣고, 관리는 `DeployableManager`로 전달하게 만듦. `Deploy` 및 오퍼레이터가 사망하는 상황에서 `CurrentOperatorDeploymentCount` 값이 관리됨

- [x] 현재 선택된 `DeployableBox`의 경우 살짝 위로 솟아오르는 효과 부여
	- 생각보다 한 방에 잘 되는 느낌이라 오히려 신기하다?
	- 하단 박스를 클릭해서 오퍼레이터가 선택된 상황에서, 배치된 오퍼레이터를 클릭했을 때에도 박스 선택이 취소되도록 수정
	- 가끔 하단 UI가 왼쪽으로 튀어나가는 현상이 있음. 
		- 배치된 오퍼레이터를 클릭한 다음에 빈 공간을 클릭해서 취소할 경우 하단 바가 왼쪽으로 튀어나가는 현상임
		- `DeployableBox`의 `ResetPosition`에서 원래 위치로 돌아가는 로직이, 위치가 변한 적 없는 박스도 이동시키기 때문에 발생했음
		- `originalPosition`이 초기화된 적 있는지를 체크하는 별도의 `boolean` 필드를 하나 두고, 이게 `true`인 것만 원래 위치로 되돌린다. 

- [x] 배치 중인 상태에서 `InstageInfoPanel`이 나타났을 때, 드래그가 아닌 화면의 빈 공간을 클릭했을 때에는 배치 동작을 취소시킴
	- `CancelPanel(Button)`을 `InstageInfoPanel`의 가장 윗쪽에 추가, 클릭 시 `DeployableManager`의 `CancelDeployableSelection`을 동작시킨다.
	- 해당 오퍼레이터가 배치되었는지 여부를 체크해서 `CancelPanel`을 나타나도록 구현했음. 단순히 이렇게만 하면 타일이 정해지고 방향을 설정하는 과정에서 `CancelPanel`이 클릭되기 때문에 배치 동작이 취소되는 현상이 있었다.
	- `DeployableManager`에서는 현재 배치 로직 중 어떤 상태에 있는지를 체크하고 있음
```cs
    public bool IsDeployableSelecting { get; private set; } = false; // 하단 UI에서 오퍼레이터를 클릭한 상태
    public bool IsDraggingDeployable { get; private set; } = false; 
    public bool IsSelectingDirection { get; private set; } = false;
```
> 이들 중 `IsDeployableSelecting` 부분에서만 `CancelPanel`을 클릭했을 때 동작하도록 했음.


#### 이슈 수정하기
- [x] 모든 `Enemy`가 파괴되었거나 도착 지점에 도착했는데도 `Stage`가 끝나지 않는 문제 수정
	- 코루틴이 엉키고 있나? 
	- `OnEnemyDefeated`에는 `StartCoroutine`을 넣었는데, `OnEnemyReachDestination`에는 `StartCoroutine`을 넣지 않아서 발생한 문제.
	- 계속 착각하고 있었네 아..

- [x] 맵 에디터에서 타일 만들 때 콜라이더 크기 조정
	- `Tile`의 스케일 구조 자체를 수정함
		- 기존엔 부모 타일의 스케일이 `1, 1, 1`이고 `Cube`의 스케일이 `0.98, 0.1(0.5), 0.98`였으며 콜라이더는 양쪽 모두에 구현이 되어 있었음
		- `Cube`를 시각화만 담당한다고 설정했기 떄문에, 콜라이더는 부모에만 있으면 됨.
	- 따라서, 부모 타일의 스케일만 `0.98, 0.1(0.5), 0.98`로 수정하면 나머지 콜라이더를 수정한다든가 자식 오브젝트의 스케일을 수정한다든가 하는 과정이 모두 필요 없게 된다. 다 부모 타일을 따라가기 때문임
	- `Cube`에 있는 `BoxCollider`도 제거함
	- 기존에 구현된 맵들에 있는 타일들의 스케일도 수정함

## 250223

### 짭명방

#### 이슈 해결하기
- [ ] `OperatorInfoPanel`이 나타난 상태에서, 다른 오퍼레이터를 클릭했을 때 오퍼레이터 변환이 될 때도 있고 안될 때도 있음
      - 일단 레이캐스트 자체가 타일 a를 클릭해도 타일 b로 인식되는 경우가 있다.
	- 모든 맵의 `Tile`에 있는 콜라이더의 `Size`를 `Cube`의 크기와 동일하게 수정함
	- `ClickDetectionSystem`의 `handleClick`에서 가장 우선적으로 타일 클릭 시 오퍼레이터 / 배치 가능 요소 클릭한 것부터 처리하게끔 수정 (임시로 복붙해놓음)
	- 일단 이 정도로 해두고 무슨 이슈가 발생하면 그때 수정하자
##### MeteorSkill에 의한 킬카운트 중복 현상
- [x] `MeteorSkill`이 있는 상태에서 `Enemy`가 죽을 때 킬 카운트가 여러 개 올라가는 현상
	- **`MeteorController`에서 구독 해제 메서드는 전부 `OnDestroyed`로 빼두고** `HandleCasterDeath(Operator op)`나 `HandleTargetDeath(UnitEntity enemy)` 등은 전부 `DestroySelf()`라는 메서드로 연결함(단순히 `Destroy(gameObject)`의 역할만 함)
	- 위 과정은 진행 상황에서 접근한 방법이고, **이 상황의 문제는 적이 파괴되었거나, 체력이 0인 상황에서 메테오 오브젝트에 의해 `TakeDamage` 메서드가 다시 발생**한 것임
		- 일단 가장 쉬운 해결책은 `TakeDamage` 메서드가 체력이 0일 때에는 대미지 적용 동작이 작동하지 않도록 하는 것임.
		- 근데 "파괴됐는데 대미지를 주는 동작이 발생한다"는 것 자체가 좀 이상하니까, 그 쪽으로 디버깅을 시도해봄.
	- 범위에 1개가 들어올 때랑 2개가 들어올 때랑 생성되는 메테오 수가 달라짐 - 적당 2개씩 생겨야 하는데, 적 2개가 범위에 있다면 메테오가 적당 3~4개씩 생기는 경우가 있음. ????
		- 추정해보면 "적"을 판단하는 기준이 "타일 위에 있는"일 때 이 메테오 스킬을 작성했다. 이 때 적은 모든 타일 중에 1개 위에만 있는 상황이었음. 근데 지금은 콜라이더로 타일 위에 있는지 여부를 판단하고 있기 때문에 2개의 타일, 3개의 타일 등등 위에도 `Enemy`가 존재할 수 있는 상황이 되었음
		- 따라서 단순히 타일 위에 있는 Enemy를 추적하는 식의 구현이라면, **여러 타일에 걸친 Enemy의 경우에는 스킬이 여러 번 적용될 것이다.** <-- 이거 좀 중요해보임. 현재 구현된 범위 관련 스킬에 다 적용할 수 있는 얘기임
			- 그래서 두 개의 타일에 걸친 적은 한번에 2개의 메테오를 완벽히 같은 프레임에 동시에 맞게 되면서 TakeDamage나 Die 메서드도 2번 발동하는 게 아닌가..라고 생각한다. 
	- **타일 기반 로직 자체는 유지해야 하기 때문에, 타일 위의 적을 얻되 추가로 중복은 제거해야 함.**
		- 이를 위해 `MeteorSkill`을 아래처럼 수정했다.
```cs
// 적의 인스턴스 id를 관리하는 해쉬셋
private HashSet<int> enemyIdSet = new HashSet<int>();

protected override GameObject CreateEffectField(Operator op, Vector2Int centerPos)
{
	enemyIdSet.Clear();

	...

	foreach (Enemy enemy in tile.GetEnemiesOnTile())
	{
		int enemyId = enemy.GetInstanceID();
		if (enemyIdSet.Add(enemyId)) // 중복된 enemy가 아니어서 해쉬셋에 추가를 성공하면
		{
			CreateMeteorSequence(op, enemy);
		}
	}
}
```

> 일단 3번 정도 테스트를 더 해봤고, 적을 두 타일에 걸치게 한 상황에서 제거했을 때 카운트가 중복해서 올라가는 현상이 발생하지 않았음. 
> 한편 `enemyIdSet`을 쓰는 부분을 주석처리하고 다시 테스트 해봄 : 걸쳐진 상황에서 MeteorSkill이 발생시키는 대미지로 Enemy가 사망했을 때 2킬이 올라가는 걸 확인했음. 아마 이게 원인이 맞을 듯?

- 타일을 기반으로 하는 범위 스킬 : `ArcaneFieldSkill`이나 `SlashSkill`도 수정이 필요한지 점검함
	- 저 해쉬셋을 부모 클래스인 `AreaEffectSkill`로 뺐음.
	- 단, `ArcaneFieldSkill`의 경우는 유지되는 스킬이니까 `Controller`에서 수정한다. - 이미 `dict`로 중복이 방지되고 있음 : `!affectTargets.ContainsKey(enemy)` : 같은 `key = enemy`가 없을 때에만 적용되는 방식.
	- `SlashSkill`도 이미 대미지를 입은 적을 `damagedEnemies` 해쉬셋으로 관리하고 있어서 괜찮았다. 

## 250222

### 짭명방

#### 이슈 해결하기
- [x] `OperatorInfoPanel` : 배치된 다른 오퍼레이터 클릭 시 패널 업데이트되지 않는 현상
	- 확인해보니 스탯은 업데이트되는데, 기본 정보와 스탯이 업데이트되지 않음
	- `currentOperator`가 갱신될 때, 배치된 오퍼레이터에서는 `currentDeployableInfo`가 갱신되지 않는 차이가 있음
	- 그렇다고 배치된 오퍼레이터인 `Operator`를 그대로 쓰자니 `BaseData`가 초기화되지 않는 시점에서 불러오는 이슈가 있음
	- 조건문으로 나눠서 처리하는 방법도 있고, `Operator`나 `DeployableUnitEntity` 자체에 `DeployableInfo` 정보를 포함시키는 방법도 있는데, 아무래도 후자가 더 깔끔할 것 같음. 생각은 조금 더 해봐야겠지만..
	- 생각보다 쉬울 듯. `Initialize`를 들어갈 때 `DeployableInfo`에 있는 정보로 들어가고 있어서, 전달하는 타입만 변경해주면 된다. 
	- **`Initialize`에 `DeployableInfo` 정보를 넣음. 이 정보는 `DeployableManager`에 의해 객체의 종류마다 1개씩 관리되고 있음** (참조가 전달됨)

- 일단 `OperatorInfoPanel`의 전체적인 구조를 조금 더 보기 편하게 고쳤는데, 새로운 이슈들이 발생함
	- [x] 배치된 오퍼레이터를 클릭했을 때 시점은 전환되나 다이아몬드 패널이 나타나지 않음
	 - [x] 다른 오퍼레이터를 클릭했을 때 전환되지 않음
		 - 위 2개의 문제는 비슷한 이슈로 발생함 : `DeployableInfo`를 이제 `DeployableUnitEntity`에 할당하기로 했는데, `Operator`에서 `DeployableInfo`를 새로 `new`로 선언하면서 벌어진 현상이다.
		 - `DeployableInfo`을 `DeployableUnitEntity`를 상속받기만 하면 다 사용할 수 있는 구조로 만들었기 때문에 굳이 `Operator`에서 `new`로 선언할 필요가 없었음.
	- `RecoverSP`가 동작하지 않는 문제도 있었는데 일시적인 증상이었나보다.




## 250221

### 짭명방

#### 이슈들 
- [ ] `Enemy`가 사라지지 않았는데 승리 판정이 나는 상황 (특정 조건에서 `Enemy` 1개체가 죽었음에도 2킬이 올라가는 현상)

- [ ] `OperatorInfoPanel` : 다른 오퍼레이터 클릭 시 패널 업데이트되지 않는 현상
	- 이거는 정리하는 김에 오퍼레이터 클릭 시 나타나는 사각형 내외부 클릭 로직도 함께 만지면 좋을 것 같다. 오퍼레이터가 클릭된 상태에서 다른 오퍼레이터가 클릭될 때의 전환이라든가, 외부를 클릭하면 다시 원래의 카메라로 돌아온다든가 하는 등등.

- [x] 유니티 버전 업그레이드하면서 아틀라스 텍스쳐를 불러오지 못하는 오류
```
The character with Unicode value \u25B6 was not found in the [SUITE-Regular SDF] font asset or any potential fallbacks. It was replaced by Unicode character \u0020 in text object [CurrentSpeedIcon].
UnityEngine.Debug:LogWarning (object,UnityEngine.Object)
TMPro.TextMeshProUGUI:SetArraySizes (TMPro.TMP_Text/TextProcessingElement[]) (at ./Library/PackageCache/com.unity.ugui@03407c6d8751/Runtime/TMP/TextMeshProUGUI.cs:2005)
TMPro.TMP_Text:ParseInputText () (at ./Library/PackageCache/com.unity.ugui@03407c6d8751/Runtime/TMP/TMP_Text.cs:2017)
TMPro.TextMeshProUGUI:OnPreRenderCanvas () (at ./Library/PackageCache/com.unity.ugui@03407c6d8751/Runtime/TMP/TextMeshProUGUI.cs:2471)
TMPro.TextMeshProUGUI:Rebuild (UnityEngine.UI.CanvasUpdate) (at ./Library/PackageCache/com.unity.ugui@03407c6d8751/Runtime/TMP/TextMeshProUGUI.cs:227)
UnityEngine.Canvas:SendWillRenderCanvases ()
```
- 예전에 적어놓은 거 중에 [[Unity - Font Atlas]] 가 있음. 이거 참고해서 수정해본다
	- 아틀라스의 문제가 아닐 수도 있겠다. 배속에 사용되는 ▶ <-- 이거가 폰트에 없기는 했음
	- 그래서 **2배속, 1배속, 일시 정지 상태 등을 TMP에서 이미지로 변경해서 넣음**
- 추가로 `UIManager`에서 직접 개별 버튼들의 토글 기능 등을 담당했는데 이들을 3개의 버튼이 들어가는 `InGameTopButtonContainer`로 빼둠. `UIManager`에서는 이들을 연결해주는 역할로 수정.
- 스테이지의 `UI` 렌더링 순서도 조절
> 폰트를 이미지로 바꾸니까 경고문 자체는 없어진 관계로 체크

- [x] `Pause` 상태에서 로비로 돌아가기 버튼을 눌렀을 때의 동작
	1. PauseOverlay가 뜬 상태에서 메뉴로 돌아가기 버튼을 눌렀을 때, 오버레이를 없애고 패널이 나타나도록 구현
	2. 메뉴로 돌아가는 패널은 애니메이션이 구현되어 있는데, `DOFade` 등 **애니메이션 메서드의 뒤에 `.SetUpdate(true)`를 추가해서 `Time.timeScale`을 무시**하도록 함 (게임의 timeScale을 멈춘 상태이기 때문에 이와는 별개로 동작해야 함. 예전에 한 적 있는데 까먹었다.)

## 250220
> 버그가 계속 나오네 ㅋㅋㅋㅋㅋㅋㅋ 미치겠다
### 짭명방

#### 기타 이슈 수정

- [x] `MeteorSkill`이 사용되던 중에 게임이 끝나고, 로비로 나갔다가 다시 돌아와서 스킬이 사용된 경우 이전 게임에서 사용된 스킬 범위가 다시 나타나는 현상
	- `StageManager`에서 `OnGameEnded`라는 이벤트를 추가
	- 이거의 경우 위의 이벤트 추가보다는 `AreaEffectSkill`에서 스킬이 끝나고 사라져야 하는 상황에서, 해쉬셋으로 구현된 오브젝트 풀은 제거하는데 정작 해쉬셋을 초기화하지 않아서 발생하는 문제였음
	- `BaseSkill`의 `CleanupSkillObjectPool`을 `CleanupSkill`로 수정, 오브젝트 풀을 수정하는 로직은 `AreaEffectSkill`에 들어가 있다.
	- 좀 특이한 건 이게 **씬이 전환되어도 정보가 유지가 되고 있다**는 거? `스테이지 -> 메인 메뉴 -> 스테이지`로 돌아와도 이전 스테이지 씬의 정보가 유지되는 다소 이상한 상황이 있다. 지금은 초기화 로직을 끼워넣었으니 됐겠지만..


- [ ] `Enemy`가 사라지지 않았는데 승리 판정이 나는 상황도 있음
	- `Enemy`의 체력이 사라진 게 이슈가 아닌 듯??
	- **정확한 발동 조건을 파악하기 어려운데**, 일단 확실해보이는 상황은 Enemy가 Vanguard에 의해 저지된 상황에서 `Meteor` 스킬로 죽었을 때는 확실히 2킬씩 올라가는 듯
![[Pasted image 20250220175243.png]]
> 이런 느낌의 문제가 발생 중이다

- [x] `StatsPanel` : 퇴각 후 재배치했을 때, 같은 오퍼레이터인데도 다른 항목으로 나타나는 현상
	- `Operator`로 관리되던 `OperatorStats`을 `OwnedOperator`로 바꾸면 되지 않을까로 접근 시작
	- 완료 : `Operator` 대신 `OperatorData`를 사용하면 됨 - 한번에 한 개가 배치되고, 사라졌다가 나타났다가 하는 요소도 아니며, `Operator` 내부에 `BaseData`로 사용 중이었기 때문에 수정도 편했음

- [x] 오퍼레이터가 배치 중인 상황에서 타일 위로 스냅핑됐을 때, `Enemy`가 오퍼레이터를 공격하는 현상
	- `250214`(한 달 단위로 일지를 옮기므로 날짜만 써놓음)에 **저지 관련 로직에 콜라이더와 트리거 판정을 넣으면서 발생하기 시작**한 것으로 보인다.
	- 현재 저지 판정은 `Operator`에서 `Enemy` 콜라이더가 들어왔을 때를 기준으로 함
	- **여기에 `IsDeployed`라는, 오퍼레이터가 현재 배치되었을 때 이 `OnTriggerEnter`가 작동해야 하도록 구성해야 하는데, 해당 조건이 없었어서 이를 추가했다.**
	- `Operator`에 있는, 관련하여 중복된 필드와 메서드를 `DeployableUnitEntity`로 빼뒀음. `IsDeployed, IsPreviewMode` 등등.

## 250219

### 짭명방

#### 스테이지 씬 로비로 나가는 패널 추가
- 구조
```
ConfirmationToLobbyPanel(Empty)
- BlurArea
- ContentBox
	- TextArea
	- ButtonArea
		- Button_ReturnToLobby
		- Button_Cancel
```
- 가운데 줄이 뜨고, 뒷배경은 블러 처리가 되어야 함
	- UI 블러의 경우, [이 동영상](https://www.youtube.com/watch?v=CFcGRE1DJRQ)을 참고.
	- [`Unified-Universal-Blur`](https://github.com/lukakldiashvili/Unified-Universal-Blur)으로 접근, 패키지 매니저에서 깃 링크로 설치하는 항목이 있음
	- **근데 `유니티 6`부터 지원한다.** 
		- 설명 내용을 보면 개발 버전의 업그레이드...를 굳이 할 필요는 없어보이기는 한다. 없어도 구현 자체가 불가능한 것도 아닐 것이다. 
		- 그런데 필요해진 김에 백업하고 개발 버전의 업그레이드를 진행해봄.

> 블러 자체는 잘 동작하지만, 같은 캔버스에 있는 다른 패널들이 블러 처리되지는 않는 듯. 블러는 아니지만, 이들은 BlurArea의 알파값을 높이는 방법으로 잘 보이지 않게 했다.

> 유니티 버전 업그레이드 때문인지는 모르겠는데, `DG.Tweening`의 `CanvasGroup.DOFade(endValue, duration)`이 살짝 이상하게 동작한다. `duration`에 따라 `endValue` 값이 변하는 이슈가 있음. 그래서 애니메이션인데 애니메이션이 아닌... 이상한 현상이 있다. 어쨌든 지금 상태가 잘 동작하니까 이대로 둠.



##### Unified-Universal-Blur 설치법
- 유니티 6에서 동작함
- 동영상의 경우 Assets에 URP 폴더가 따로 있는데, 내 경우는 아님(Assets/)에 있기는 하다
- 만약 못 찾겠다면, `Project Settings`에서 `Render Pipeline Asset`의 URP Asset으로 접근, `Renderer List`에 있는 **`Universal Renderer Data`로 접근한 뒤, `Add Renderer Feature - Universal Blur Feature`을 추가한다.**

#### 유니티 6으로 개발 버전 업그레이드
-  위에서 블러 라이브러리를 가져오기 위해, 기존 프로젝트를 백업하고 유니티 6으로 업그레이드를 시도해봄 
- **대부분의 기능이 잘 작동함.** 텍스쳐를 가져오지 못한다는 오류가 발생하는데, 실행에 치명적이지는 않음
```
Graphics.CopyTexture called with null source texture
UnityEngine.GUIUtility:ProcessEvent (int,intptr,bool&)
```
- 유니티 허브 설치 드라이브의 용량이 꽉 차는 이슈가 있어서 설치 드라이브를 바꿈. 지우고 재설치한다는 뜻.

#### 기타 이슈 수정
- [ ] `OperatorInfoPanel` : 다른 오퍼레이터 클릭 시 패널 업데이트되지 않는 현상
	- 배치 전의 요소를 클릭할 때는 업데이트가 잘 되고, 배치된 요소를 클릭했을 때 업데이트가 되지 않는 듯.


## 250218

### 짭명방
#### 인게임 정보 패널 재조정
- 스킬 정보 보는 패널 추가
![[Pasted image 20250218174635.png]]
> 이런 느낌인데, 깔끔하게 떨어지는 것 같아서 이 정도로 하겠음

- 로비로 나가는 버튼까지만 추가하고, 기능의 상세한 구현은 내일 진행


## 250217

### 짭명방

#### 마감 X) 스테이지 씬 - 오퍼레이터 클릭 시 패널 재조정
- 이제 스킬이 추가되었기 때문에 스킬에 대한 설명도 함께 들어감
- 기존 레이아웃도 원작을 참고해서 변경하겠음

~~그 과정에서 기존 `InfoPanel`, 현재 `InStageInfoPanel`의 코드를 수정함 - 배치 여부에 따른 정보를 가져올 때 다른 곳에서 가져오는데, 좀 헷갈리게 짜여져 있어서 이를 고치고 있다.~~ 
> 고치니까 더 이상해진다. 너무 길어지는 부분만 별도의 메서드로 뺐음.


#### 이슈 수정
- [x] `OperatorDetailPanel`에 들어갔다가 다시 `OperatorInventoryPanel`로 나왔을 때 스킬 아이콘이 초기화되지 않는 현상
	- 로직을 보면 인벤토리 패널의 슬롯들 초기화는 `OnEnable`에서 이뤄짐
	- 근데 `SquadEditPanel`에서 들어갈 때는 이제 잘 나타나는데(예전에 수정), `InventoryPanel`에서 나올 때는 오히려 동작하지 않음.
	- **초기화 시점에 이미지가 없는가? : 스프라이트 인식 잘 됨**
	- 패널 전환 과정에서 `CurrentPanel`을 바꾼 다음에 전환 동작이 이뤄지게 변경하니까 스킬 아이콘이 잘 나타나기는 한다. 바뀌는 과정이 살짝 불편한 느낌은 있는데, 그대로 진행하겠음
	- `InitializeSkillIcon`과 `Update` 메서드를 별도로 구현했다. 

- [x] `OperatorInventoryPanel` : 패널 활성화 동안, 변경한 스킬이 있다면 유지시킴
> 추가로, `OperatorInventoryPanel`에서 어떤 슬롯을 클릭하고, 해당 슬롯의 스킬 아이콘을 바꾼 상태에서 다른 오퍼레이터를 클릭하면 바뀐 스킬 아이콘이 유지된다.
> - 이는 해당 오퍼레이터의 `DetailPanel`에 들어갔다가 나와도 유지되며, `Default` 스킬과는 별도로 동작하는 걸 확인했음
> - 그리고 다른 오퍼레이터의 `DetailPanel`에 들어갔다 나와도 기존에 오퍼레이터들을 클릭해가면서 바꿨던 스킬들도 유지됐다. 즉, **기존 패널이 비활성화됐다기 보다는 단순히 투명하게 처리된 것으로 보임** 
> - 즉 지금의 구조랑 원작의 구조가 살짝 안 맞는 느낌이 있다. **전체적으로 바꾸는 것보다는 단순히 `OperatorInventoryPanel`에서 스킬을 바꾸고 다른 오퍼레이터를 클릭하더라도 해당 스킬이 유지되는 정도로만 구현하겠음**
> `OperatorSlot`에서 `SelectedSkill`이라는 필드를 추가, 현재 선택된 스킬을 관리하도록 함


## 250214

### 짭명방

#### 이슈 수정 - 저지 매커니즘 수정하기
- [x] **저지 조건 변경하기 : 같은 타일 위에 있음 -> 두 엔티티의 콜라이더 충돌 감지**
```
ㅁ O
(O는 저지 가능한 오퍼레이터)
```
> 만약 `오퍼레이터 X`의 공격 범위가 ㅁ은 포함하고 O는 포함하지 않을 때, Enemy가 ㅁ에서 O로 진행하는 경우 저지되는 위치는 `O` 타일이 된다. 따라서 `X`는 `O`에서 저지되는 적을 공격하지 못함.
> - 원본 명방의 경우, 저지되는 위치는 `ㅁ`이 되므로 `X`가 공격할 수 있음
> - 또, 원본 명방에는 `O`를 지나고 있는 상황에서 `O`에 오퍼레이터를 배치하면서 `Enemy`가 `ㅁ`과 `O` 2개의 타일에 걸치는 상황도 만들 수도 있음

- 기존의 저지 로직은 특정 타일에 Enemy와 Operator가 모두 있을 때를 기준으로 판단함. 특정 타일에 Enemy가 있는지 없는지 여부는 타일들에서 판단한다.
- 저지 판정의 기준을 오퍼레이터의 콜라이더와 충돌했을 때를 기준으로 함
	- 대신 지름이 1보다는 조금 작아야 할 것 같다. 만에 하나이기는 하지만, 지상 타일과 언덕 타일이 붙어 있는 상황에서 언덕에 있는 오퍼레이터에 의한 저지 상황이 발생할 수도 있을 것 같음. 기우일 수도 있고.
	- 모두 `Box Collider`를 놓고, `0.8, 1, 0.8`로 설정함
- 각 `Operator`나 `Enemy`에 **`rigidBody`도 추가. `Collider` 간의 충돌을 감지하려면 필요함.**
	- `Is Kinematic` 비활성화 -> 물리엔진 무시
	- `Freeze Position` - `Y` 체크 (체크 안하면 떠오르는 듯한 현상이 있음)
- 추가로 공격 가능한 적에 대한 조건을 바꿔줬음
	- `IsCurrentTargetInRange`도 그리드 포지션이 아니라 타일 위에 적이 있는지 여부를 체크
	-  `ValidateCurrentTarget`도 현재 저지 중인 상태라면 공격 범위를 점검하지 않음

- [x] **두 타일에 걸쳐있는 Enemy 구현하기**
- 또 이런 처리에서, **Enemy가 두 타일에 걸쳐있는 상황**도 콜라이더로 만들 수 있을 것 같아서 시도해봄
	- 콜라이더의 y값이 낭낭한 1이므로 타일 위에 있다면 무조건 충돌이 발생하는 상황임
	- 타일에도 `rigidBody`랑 `트리거 콜라이더`를 추가한 다음, 충돌이 발생하는 상황일 때 해당 타일 위에 있다고 하면 될 듯?
	 - 대신 이 경우 Enemy의 `CurrentTile`을 기반으로 한 로직들을 싹 뜯어고쳐야 함
		 - `UnitEntity.CurrentTile`을 `DeployableUnitEntity`로 옮김. 관련 메서드들도 함께 옮김.
> 잠깐 밥 먹고 와서 처리함. 앞으로 할 일 : `Enemy`의 `CurrentTile` 관련 로직들을 고치고, `OnTrigger` 메서드들 추가하기. 타일 위의 적 리스트는 타일에서 관리하면 될 듯.
	- `rigidBody`를 추가한 이후로 이상한 상황들이 많이 나온다. 자기 혼자 돌아간다든가, 어디로 날아간다든가.. 전부 `freeze`를 걸어뒀음.

- 구현 자체는 됐고, 트리거 이벤트의 발생은 `Enemy`에서 컨트롤한다.. `contactedTiles`라는 충돌이 발생 중인 타일들의 리스트를 관리하고, `OnTriggerEnter` 및 `OnTriggerExit` 내에서 해당 타일의 `EnemyEntered` 및 `EnemyExited`을 통해 `Tile`에서 `Enemy` 리스트를 관리하는 메서드를 실행시키는 방식이다.
	- 굳이 이중으로 구현할 필요를 못 느끼겠어서 이렇게 구현했음.
	- 이렇게 구현하면 `Enemy`가 파괴되는 상황에서도 `contactedTiles`에 있는 모든 타일에 대해 `EnemyExited`을 실행시키면 되니까 괜찮은 것 같음.

## 250213
- 치지직을 켜놓고 작업할 때가 많은데, 유니티에서 게임 테스트할 때 렉이 걸린다. 유튜브는 켜놓고 작업해도 게임 플레이 상황에서 버벅임이 발생하지는 않음. 그리드 때문인가? 

### 짭명방
#### 밸런싱 
- 적이 몰려나오는 스테이지 : 이에 대응할 범위 기술이 `Guard`의 1정예화 스킬, `Caster`의 1정예화 스킬 정도밖에 없음 - 따라서 이를 구현하겠다면 가능한 뒷쪽에 모는 게 좋을 듯

#### 이슈 수정

- [x] 아이템 계산 로직 오류
	- 예를 들어서 중급 경험치 1개, 하급 경험치 5개가 있는 상황일 때, 기본적인 계산 결과는 이들을 모두 사용했을 때 최대 0정예화 9레벨에 도달할 수 있다고 나온다. 그런데 4레벨을 찍고, 다시 9레벨을 찍으면 갑자기 10레벨이 활성화됨. 
	- 비교 실험 ON
		1. 0정예화 4레벨 도달 후 9레벨 도달하기 
			1. 하급 2개 사용. 결과 상태 : 4레벨 49/151
			2. 중급 1개, 하급 2개 사용. 결과 상태 : 9레벨, 73/236
			- 하급 1개가 남음. 경험치 200이므로 10레벨 도달 가능
		2. 바로 9레벨 도달하기
			- 중급 1개, 하급 4개 사용. 결과 상태 : 9레벨, 24/236
	- **사용한 아이템이 똑같은데 올라가는 경험치가 다르다.**
		1. 어떤 게 맞는 계산인가?
			 - 0정예화는 경험치 요구량이 1레벨 올라갈 때마다 17씩 증가함 `100, 117, 134, ...`
			 - 등차수열의 합 공식을 까먹었는데, $\frac{n}{2} *[2*a + (n-1)d]$ 이라고 함(a는 초항, d는 공차). 4레벨 도달에 필요한 경험치량은 351, 9레벨은 1276이다.
		 - ~~원인을 알겠다 : 4레벨에서 시작해서 9레벨에 도달하는 경우의 계산은 현재 경험치량이 반영이 안된 것으로 보인다. 두 케이스에서 중급 1개, 하급 4개를 썼을 때를 비교하면 정확히 49의 차이가 나는데, 이는 4레벨에 도달했을 때의 현재 경험치와 동일함~~
			 - 중급 1개, 하급 4개를 쓰면 500 + 800 = 1300이므로, 2번의 케이스가 더 정확한 계산임
			- ~~즉 코드에서 수정할 부분은 **레벨업 시 현재 경험치가 반영이 안되는 부분**임~~
				- 이게 DP가 들어간 재귀식을 뜯는 과정이라 살짝 복잡하다. 이거만 해결해도 오늘은 괜찮겠음. 더 하면 좋고..
				- 일단 AI한테 물어보면 현재 경험치량 정보가 소실된다는 얘기를 하면서 코드 전체를 고치려고 하는데, 내 생각에는 그게 본질은 아닌 것 같다. 왜냐하면 DP와 재귀식을 활용한 코드는 목표 경험치**량**을 달성하기 위한 최적의 아이템 사용 계획을 계산하는 방식이기 때문임 
					- 따라서 시작점이랑 도착점이 정해져 있고, 그 도착점을 넘되 가능한 살짝 넘는, 사용할 아이템 조합의 경우의 수를 찾는 문제라고 봄. 그래서 시작 경험치량이 소실된다라고 지적하는 AI의 말은 정확히 들어맞지는 않는 것 같다. 
		 - **현재 경험치량이 반영이 안된 게 아니라, 2번 반영된 게 문제임!!**
			 - 문제를 잘못 파악했다. 현재 경험치가 있는 상황이라면, 그 현재 경험치가 2번 반영되는 문제 라는 게 더 정확한 설명으로 보인다. 왜 착각했지?
```cs
        // 필요한 총 경험치 계산
        int requiredExp = OperatorGrowthSystem.GetTotalExpRequiredForLevel(phase, currentLevel, targetLevel, currentExp);

        // DP를 위한 캐시 테이블 - key : 경험치 / value : 경험치에 도달하기 위한 아이템 사용 계획 (a 아이템 x개, b 아이템 y개..)

        // 재귀적으로 최적의 조합 찾기
        ExpItemUsagePlan optimalPlan = FindOptimalCombination(
            requiredExp,
            availableItems.OrderBy(x => x.item.expAmount).ToList(),
            new Dictionary<int, ExpItemUsagePlan>()
        );
        
        // 경험치 오버플로우 처리 : 플랜 결과 후 남은 현재 경험치량이 최대 경험치량을 초과할 경우
        //int remainingExp = (currentExp + optimalPlan.totalExp) - requiredExp;
        int remainingExp = optimalPlan.totalExp - requiredExp;
        int finalLevel = targetLevel; 
```
> `ExpCalculationSystem`에 있는 `remainingExp`를 계산하는 부분에 문제가 있었음. `requiredExp`를 계산할 때 `currentExp`가 이미 들어가므로, 여기서 추가로 더할 필요가 없다.
> - 그래서 아래처럼 고치고,  `requiredExp`이라는 변수 이름을 `requiredExpForExactLevel`로 변경, `totalExp` 같은 애매한 변수명도 `totalItemExp`로 수정했음.

## 250212

### 짭명방

#### 밸런싱
- 적의 스탯 조정
	- 사실 적의 스탯이 문제였다기보다는 아래의 **스탯이 정상적으로 적용되지 않는 상황 때문이었던 것으로 보임.** 지금 구현한 `Enemy` 중, `Defender`에게 들어가는 대미지는 `원거리`만 제대로 넣고 있음. 나머지는 9~10 정도의 대미지만 들어가고 있음.

#### 이슈 수정
- [x] 공격받을 때, 방어력이 적용되지 않는 이슈
	- `UnitEntity`의 `CurrentStats`가 있고, `Enemy`나 `Operator`의 `CurrentStats`이 따로 있음.
	- **대미지 계산 로직을 `UnitEntity`에서 가장 아래의 클래스들(`Operator, Enemy, Barricade`)로 옮김.** `UnitEntity`의 `CurrentStats`을 참조했더니 0으로 인식하는 상황이 발생하는 것 같다
	- 근데 그러면 `Health`는 왜 잘 작동하는 거지? : 사실은 `currentStats.Health`를 참조한다는 개념보다는, 초기화에서 `MaxHealth`의 값을 가져와서 사용하기 때문임
		- 이것도 프로퍼티에서 퍼블릭 세터를 놓는 대신, 체력 수정 메서드를 별도로 구현함

- [x] 디펜더 저지가 3인데 무한 저지가 가능한 현상
	- 저지 자체는 오퍼레이터에 설정된 저지수를 더 넘지 않음
	- `Enemy.CheckAndAddBlockingOperator` 메서드를 보면 `enemy`는 타일에 오퍼레이터가 있고, 자신을 저지하는 오퍼레이터가 없는 상태라면 `blockingOperator`를 설정해버림. 이 로직이 문제인 것으로 보임.
	- `blockedEnemies` 필드를 참조하는 `IReadOnlyList<Enemy>`인 `BlockedEnemies`을 `Operator`에 구현하고, `Enemy.CheckAndBlockingOperator`에서 해당 오퍼레이터의 저지 수가 현재 저지 중인 수 이상이라면 `return;`을, 아니라면 이 `Enemy`를 저지하도록 함

- [x] 버프 키니까 체력이 쫙 빠지는 버그가 생겼다. 위에서 `Health` 관련 로직 수정하는 과정에서 발생한 듯. 
	- 버프를 킬 때, 현재 체력 비율을 계산하고 최대 체력 배율 수정 -> 현재 체력값을 `최대 체력 * 현재 체력 비율`로 계산했던 게 기존 로직이었는데, 수정하는 과정에서 `현재 체력 * 현재 체력 비율`로 잘못 넣어서 그럼.

#### 공부
- [[Csharp - 퍼블릭 세터 vs 퍼블릭 메서드]]

## 250211

### 짭명방
- 이슈 몇 개 수정
	- 메인메뉴의 스테이지 버튼 서식 수정
	- 스테이지에서 자동 발동인 스킬이 사용 가능할 때 "스킬 사용 가능" 아이콘이 뜨지 않게 수정


## 250207

### 짭명방

> 예전 코드들을 보니까 스파게티의 냄새가 솔솔 난다
#### 기타 이슈 수정
- [x] 오퍼레이터 배치 방향이 이상하게 되는 이슈
	- 실제 공격 범위, 하이라이트되는 공격범위가 이상하게 표시됨
	- 한편 방향 인디케이터는 정상적으로 표시됨
	- 배치 방향과 버그 방향
		- 오른쪽 : 왼쪽
		- 왼쪽 : 왼쪽
		- 윗쪽 : 오른쪽
		- 아랫쪽 : 오른쪽
	- 회전이 2번 적용되어서 그런 것으로 보인다. `DeployableManager`의 미리보기 과정에서 오퍼레이터에 관한 값들이 적용되고 있기 때문에, `Deploy`할 때에는 그대로 적용하면 됨. 
	- 제대로 구분하기 위해 `baseOffset, rotatedOffset, CurrentAttackableGridPos`으로 구분했다.

- [x] `DeployableBox`의 초기화 이슈 : 초기 코스트가 배치 코스트보다 높은데도 `inActiveImage`가 활성화되는 현상
	- 박스의 초기화가 스테이지 시작 전에 되므로 초기 코스트를 0으로 인식하는 문제로 보인다.
	- `StageManager.OnPreparationComplete` 이벤트가 발생하는 시점에서는 스테이지의 코스트가 준비된 상태이므로, `InitializeVisuals`가 저 이벤트를 구독하면 될 듯.

- [x] StageManager에 있는 스테이지 설정 : 최초 코스트, 최대 코스트, 코스트 증가 속도 등을 `StageData`에서 가져오도록 수정
	 - 코스트 증가 속도 로직 수정 : 기존엔 `CostIncreaseInterval`이라는 이름으로, `Time.deltaTime / CostIncreaseInterval`의 값을 더하는 식으로 구현했으나, 직관적이지 않았다. 
	 - 따라서 `CostIncreaseSpeed`를 `StageData`에서 가져오도록 구현하고, 그 로직도 `Time.deltaTime`에 곱하는 방식으로 바꿨음. 변수명도 `TimeToFillCost`로 변경.

- [x] 스테이지 시작 상황 : `StageLoader`의 역할은 데이터 전달 정도로 줄이고, `StageManager`의 역할을 늘림
	- 스테이지 매니저에서 시작 로직을 다 처리하게 하기 위함인데, 현재 2가지 문제 발생 중
	- [x] 1. 로딩 화면이 없어지지 않음(메서드는 실행됨)
		- `Destroy`에서 컴포넌트를 지정했기 때문. 오브젝트를 지정해야 한다.
	- [x] 2. 맵이 로드되지 않음
		- 이미 `StageLoadingScreen` 자체에서 로딩 화면의 동작이 다 구현되어 있음. 굳이 `StageLoader`에서 별도로 구현해줄 필요는 없다. 
		- 따라서 `StageLoader`에 있는 로딩 화면 제어 로직은 다 제거해준다. `StageLoader`는 로딩 화면을 인스턴스화만 해주면 됨. 
	- [x] 3. `DeployableManager`에 `Deployable`들이 정상적으로 초기화되지 않음
		- 상태들은 다 잘 들어간다. `InitializeDeployableUI`에서 로비에서의 실행을 방지한다는 명목으로 넣은 코드 때문에 이상하게 동작한 듯. 
		- 로비에서의 실행을 방지하기 위해 넣은 코드가 아래와 같았는데, 이제 스쿼드 초기화를 `Preparation`에서 진행하게 되면서 이 메서드가 실행되지 않게 된 것으로 보인다. 
```cs
        //if (StageManager.Instance.currentState != GameState.Battle) return;
```
		- 위 코드를 빼고 `OnDestroy`에서 이 메서드의 이벤트 구독을 해제하게끔 바꿨음.
		- 좀 오래 걸렸다.
## 250206

### 짭명방

#### 밸런싱
;;; 고칠게 많아서 진도를 못 뺀다

#### 기타 이슈 수정
- [x] `SlashSkill` : 보는 방향의 시계방향 90도로 돌아가서 스킬이 나감
	- 스킬 프리팹을 보니까 스크립트 참조가 풀려 있다. 다시 설정하니 잘 작동함.

- [x] 1정예화가 됐는데도 스킬 범위의 확장이 적용되지 않음
	- 정확히는 **스테이지 씬**에서 확장된 스킬 범위가 적용되지 않음. 메인메뉴 씬에서는 잘 나타나고 있다. 
	- `Operator`에서 공격 가능한 타일은 `CurrentAttackableTiles = OwnedOperator.CurrentAttackableTiles`을 참조해야 하나, `Basedata.attackableTiles`을 참조하는 부분이 몇몇 있었다. 이 부분을 `CurrentAttackableTiles`으로 수정.

- [x] 왼쪽 패널 : 스탯 보는 부분에서 공격력 = 방어력인 현상
	- 체크리스트 
		- [x] 인스펙터 SerializeField 참조 오류 : 아님
		- [x] `ScriptableObject`에서 값 잘못 입력 : 아님
	- 발견) `OperatorGrowthSystem`에서 초기화 로직에서 `Defense`를 할당해야 하는 부분에 `AttackPower`가 할당되어 있었음
	- 이거 수정하는 과정에서, 추가로 `InventoryPanel`이나 `DetailPanel` 등에서 소수점으로 나타나는 스탯들은 모두 내림으로 보이게 구현함
		- 스탯 자체에 내림을 적용하면 레벨업 스탯이 소수점으로 올라가므로 문제가 발생할 것으로 보임. **실제 수치는 그대로 두되, 보일 때에만 내림 적용.**

- [x] 레벨업 패널 : 배경을 `CircleWithSmallAlpha` 텍스쳐를 사용했을 때, `Fill` 이미지에 쓰이는 텍스쳐도 `CircleWithSmallAlpha`를 사용했을 경우, `Fill Amount` 값이 변함에 따라 내부 원 부분이 보이는 현상이 있음
	- `CircleWithSmallAlpha`라는 텍스쳐 자체를 다시 뜯어서 내부 알파값을 제거한 텍스쳐를 만들어서 적용
	- 같은 이미지로 내부에 알파값을 준 것과 아닌 것 2개로 나눠서 구현함
	- 1개의 `Circle` 이미지만으로 스텐실 마스크 조건에 `Alpha = 0` 조건만 줘서 구현할 수 있나 테스트해봤는데, 안되는 것 같음.

- [x] `EnemyWeak`의 경우 도착 지점에 도달해도 카운트가 안 올라감
	- `Enemy`에 설정된 `destinationPosition`의 경우 `PathData`의 마지막 노드 + `Vector3.Up * 0.5f`로 설정되어 있었다. `0.5f`라는 값을 `BaseData.defaultYPosition`으로 수정
	- 추가로 `MoveAlongPath`의 목적지 도달 조건이 `nextPosition == destinationPosition`이었는데,  아래의 조건을 한 번 보면..
```cs
        // 노드 도달 확인
        if (Vector3.Distance(transform.position, nextPosition) < 0.05f)
        {
            // 목적지 도달
            if (nextPosition == destinationPosition)
            {
                ReachDestination();
            }
            // 기다려야 하는 경우
            else if (nextNode.waitTime > 0)
            {
                StartCoroutine(WaitAtNode(nextNode.waitTime));
            }
            // 노드 업데이트
            else
            {
                UpdateNextNode();
            }
        }
```
> 코드가 뭔가 부자연스러워보인다는 느낌이 있음. 상위 조건을 만족해도 정확한 위치가 아니면 `ReachDestination`이 동작하지 않음.
> - 그냥 똑같이 거리 조건을 주는 식으로 수정하기는 했는데.. `Vector3.Distance(transform.position, destinationPosition) < 0.05f` 뭔가 찝찝쓰.

---
## 250205

### 짭명방

#### 밸런싱
- 스탯 수치를 구체적으로 조정하겠음
	- 주로 몇 대 정도 쳐야 적이 죽고 아군이 죽을지 등에 대한 생각이 들어가야겠다.
- 구체적으로 어느 정도면 좋겠다.. 라는 기준이 잡히지 않아서 원본 게임의 0챕터를 보면서 맞추겠음. 
	- 적의 종류도 늘려야 할 필요가 있어 보인다.
		- 원석충 개념의 약한 몹
		- 댕댕이 개념의 빠른 몹
		- 리유니온 일반 병사 같은 일반 보병 느낌의 몹

- `EnemyWeak` 구현
	- **명일방주의 원석충$\alpha$ 의 스탯**
	- 모델 : 캡슐 대신 블렌더에서 반구를 만들어서 넣음
	- 반구로 넣으니까 `0.5` 정도로 띄운 기존의 구현이 이상해보인다. 높이를 살짝 낮춤.
	- 경로 데이터도 단순히 GridPosition 기반의 검색을 하고, Enemy에서 월드 포지션을 찾을 때에는 `Vector3.Up * 0.5f`을 더하기 때문에 저 `0.5`만 `EnemyData.DefaultYPosition`으로 바꾸면 같은 경로를 쓰더라도 `Enemy`마다 다른 높이를 갖게 경로를 짤 수 있음

- 스탯은 명방의 **5성 오퍼레이터 중**에서 구현한 직군에 맞는 스탯들로 맞춰봄
	- 스탯 차이가 나는 경우는 돌파가 됐기 때문으로, 기본적으로 0돌파를 전제함
- [[오퍼레이터들 스탯 정리]]

|          |  뱅가드 |   가드 |  디펜더 | 캐스터 | 스나이퍼 |   메딕 |
| -------- | ---: | ---: | ---: | --: | ---: | ---: |
| 기본 체력    |  727 | 1443 | 1475 | 651 |  679 |  845 |
| 기본 공격력   |  273 |  449 |  198 | 281 |  244 |  166 |
| 기본 방어력   |  139 |  116 |  245 |  45 |   61 |   62 |
| 기본 마법저항력 |    0 |    0 |    0 |  10 |    0 |    0 |
| 공격속도(초)  | 1.05 |  1.5 |  1.2 | 1.6 |    1 | 2.85 |
| 배치코스트    |   11 |   15 |   20 |  18 |   13 |   17 |
| 렙업당 체력   |  5.5 | 12.6 |  9.5 | 5.8 |    6 |  7.1 |
| 렙업당 공격력  |    2 |    4 |  0.9 | 9.3 |  2.3 |  2.2 |
| 렙업당 방어력  |  1.4 |  0.9 |  2.5 | 0.7 |  0.9 |  0.6 |
> 옵시디언의 표 기능을 되게 오랜만에 쓰는데 훨씬 안정적이다. 엔터가 아래로 안 간다는 거 빼면 굿.


#### 기타 수정
> 체크박스는 `ctrl + L`이란다. 옵시디언을 꽤 오래 썼는데 이제 알았습니다.
 
- [x] `SquadEditPanel`에서 이미 배치된 오퍼레이터 슬롯을 클릭했을 때 `OperatorInventoryPanel`로 이동하는데, 이 경우 이미 스쿼드에 들어가 있는 오퍼레이터도 나타나야 함. 

- [x] `SquadEditPanel` 초기화 및 오퍼레이터가 슬롯에 배치된 상황에서 스킬 아이콘이 제대로 초기화되지 않는 현상
	- 구체적으로, `StageSelect -> SquadEdit`으로 들어가면 흰색(디폴트) 아이콘이 표시되는데, `OperatorInventory -> SquadEdit`으로 나올 때는 표시가 잘 된다. 다시 실행 여부에 관계 없는 이슈.
	- 생각보다 원인을 찾기 힘들다. 이미지 컴포넌트의 활성화가 꺼져 있기는 함.
	- `OperatorSlot`에서 스킬 이미지 컴포넌트를 끄는 부분들을 제거했음. 현재는 모든 오퍼레이터의 스킬 아이콘이 구현된 상태이므로, 스킬 이미지 컴포넌트를 `null`로 설정할 이유가 없다.

- [x] `DeployableActionUI`에서 스킬 아이콘 이미지가 약간 붉게 나오는 현상
	- 스킬 아이콘이 있을 때는 Color를 `1, 1, 1`로 들어가게끔 수정.

- [x] `StageManager`에서, 적을 잡았을 때 목숨 카운트가 왜 올라감?? 
	- 프로퍼티 복붙할 때 발생시킬 이벤트를 수정을 안 했다. 이거 되게 오래됐을 것 같은데?

- [x] 뱅가드의 2스킬을 사용한 다음에, 뱅가드를 퇴각시킬 때 이벤트에 의해 이미 사라진 2스킬의 컨트롤러를 다시 호출하려는 문제
	- 이벤트를 구현할 때, `+=`와 `-=`을 함께 구현해야 하는 이유가 되겠다. 이벤트 구독 해제를 `Disable`에 구현하지 않아서 발생한 문제.
	- `MeteorController` 이외에도, `FieldEffectController`에도 `OnDisable`에 구독 해제를 달아놨다.

## 250204

### 짭명방
#### 맵 구현하기
- 맵 1-2, 1-3 구현, 경로 데이터 설정
	- 맵의 순서는 뒤바뀔 수도 있을 것 같음. 1-3을 너무 작게 만들었다.

- 기본적인 맵이랑 경로 설정, 경로를 따라 적이 잘 이동하는가 등등은 잘 구현되었음!
- `PathData`에서 경로를 설정할 때
	1. 프리팹에서는 경로 설정이 안됨. 맵을 하이어라키에 띄우고 설정해야 함
	2. 왜인지는 모르겠는데 `Start` 타일을 클릭하면 노드가 추가되고 편집이 멈춘다. 나머지는 괜찮음.

> 어쨌든 맵의 구성 자체도 생각보다 엄청 빨리 진행했다.

- `Spawner`의 경우 단순히 `MapEditorWindow`에서 `Start` 타일을 만들 때 함께 들어가고, `StageManager`에서 맵에 있는 모든 `EnemySpawner`를 가져와서 전체 적의 숫자를 계산하는 정도로만 구현이 되어 있다. 이 정도로만 해도 충분할 것 같음. `FindObjectsOfType`을 쓰기는 하지만.. 1회성 연산이고 맵이 크지 않아서 큰 상관이 없지 않을까?
#### 기타 오류 수정
- `MapEditorWindow`에서 `Enemy Spawner` 프리팹을 찾지 못하는 현상 수정
	- 경로 주의! 에디터 단위의 편집이므로 `SerializeField`에 오브젝트를 할당하는 방식으로 접근하지 못함

- `MapEditorWindow`에서 맵을 저장하는 방식 수정
```cs
	//PrefabUtility.SaveAsPrefabAsset(currentMap.gameObject, path); // Map을 단순히 저장만 함(프리팹과 연결 X)
	PrefabUtility.SaveAsPrefabAssetAndConnect(currentMap.gameObject, path, InteractionMode.UserAction); // 프리팹과 연결까지 수행
```
>- 기존 : `PrefabUtility.SaveAsPrefabAsset`은 오브젝트를 복사 / 붙여넣기 하는 방식으로 저장함. 즉 유니티에서 하이어라키의 오브젝트를 프로젝트에 드래그&드랍할 때 프리팹으로 연결되면서 참조되는 게 아니라, 아예 통째로 복사되고 별도의 오브젝트로 취급됨
>- **프리팹화 : `PrefabUtlility.SvaeAsPrefabAssetAndConnect`은 오브젝트를 프리팹으로 저장하고, 하이어라키에 떠 있는 오브젝트도 프리팹과 연결시킨다.** 파라미터가 1개 추가됨. API 메서드가 호출되는 모드 설정이라는 듯.

> 별 차이는 아닐 수 있는데, 이렇게 수정하지 않으면 굳이 하이어라키에 띄우고 작업하는 의미가 없어보임.

- 프리팹 수정하기
	- 수정된 프리팹을 프로젝트의 원본 프리팹에 드래그&드롭 하는 것만 알고 있었는데
	- **인스펙터 - `Overrides - Apply All`로도 가능함**
	- 이걸 몇달을 만졌는데 이제야 알았다 엌ㅋㅋ 기존 방식은 새로 프리팹을 만드는 행위였다는데?

- 유니티 상-식 보충
	- `private`으로 설정된 필드는 일반적인 저장상황에서 그 값이 보존되지 않음
	- `[SerializeField]`는 여태까지 `private` 필드를 인스펙터에 노출시키는 용도로만 써왔지만, 이외에도 **해당 필드에 할당된 값을 직렬화해서 저장**하는 역할도 함
	- 예를 들어 내 프로젝트의 `Tile.gridPosition`이 있다.

- Enemy가 여러 명일 때 이펙트가 정상적으로 나타나지 않는 현상
	- 이펙트의 오브젝트 풀링에서 각 이펙트의 태그를 만들 때, 객체마다 갖는 태그가 모두 달라야 함. 이걸 고려를 안해서 여러 객체가 같은 오브젝트 풀을 공유하는 현상이었음.
	- `GetInstanceID()`는 각 객체의 고유한 인스턴스 ID를 `int`값으로 반환함. `ToString()` 메서드로 `string`으로 바꿔서 사용.


### 구독을 해야 하나 말아야 하나
> - 딥시크가 터지면서 끌로드 구독을 끊었는데, 다시 해야 하나? 아니면 무료 분으로 충분한가 고민 중이다. 프로젝트 기능을 쓰면 질문의 수가 엄청 줄어들어서.. 5시간에 10개 조금 넘는 질문은 너무 적긴 하다.
> - 근데 이미지 작업 등을 하는 날에는 아예 안 쓰기도 한다. 이렇듯 순간적으로 많이 필요한 느낌에 가까움.
> - 그래서 1개월에 3만원이나 하는 끌로드를 써야 하나, 말아야 하나.. 아니면 딥시크가 불러올 경쟁을 믿고 가격 인하를 존버해야 하나.. 등등이 고민된다. 일단은 chatGPT랑 끌로드를 번갈아가면서 쓰면서 버티는 중. 네이버의 클로바 X도 써봤는데 조금 긴 스크립트를 넣었더니 작동하지 않더라.


## 250203

### 짭명방
- 남은 프로세스 정리
	- 맵 2개 더 추가로 구현하기
	- 도전과제 구현하기
		- 맵 1-1에서 `바리케이드만 배치해서 스테이지 클리어하기` 를 일단 생각하고 있음
		- 이걸 구현하려면 해당 타일을 밟았을 때 죽을 때까지 도트 대미지가 들어가되 공격 속도가 증가하는 타일을 만들어야 함
		- 도트 대미지가 들어가고 있음을 보여주는 알리는 아이콘?도 하나 필요함

#### 스테이지 진척도 구현
> - 여러 개의 스테이지, 스테이지 버튼 등을 구현하기에 앞서서, 스테이지의 진척도를 저장하는 규격을 하나 정의해놓고 시작하는 게 좋을 것 같다. 
> - 이를 ~~`StageProgressData`~~를 새로 만들려고 했는데 그 대신 `StageResultData`를 수정해서 사용하는 방식으로 수정하겠음

2. 이전 스테이지가 클리어됐을 때 다음 스테이지 버튼이 보여야 함
3. 사용자의 스테이지 진척도를 불러와서 체크해야 함
4. 스테이지 클리어 시 진척도를 저장해야 함

- 그리고 스테이지의 클리어 및 별 갯수는 `PlayerDataManager`에서 관리

- 결과물 1. [[스테이지 진척도 구현하기]]
- 결과물 2 - StageSelectPanel에 아래와 같은 자료를 추가
```cs
[System.Serializable]
public class StageUIData
{
    public string stageId;
    public StageButton stageButton;
    public Image lineImage; 
}
```
> - 각 스테이지의 클리어 여부를 검사, 이전 스테이지가 클리어된 경우 stageButton과 lineImage가 활성화되는 방식임
> - 만약 추가 챕터 등을 구성한다든가 하면, 이전 스테이지 자체를 넣을 수도 있을 듯. 현재는 단순히 string을 검사하는 식으로 조건을 체크하고 있다.

- 예시
- 임의로 1-1만 클리어했다고 정의해둔 상태임
![[Pasted image 20250203173957.png]]
> 여기에 아트를 추가를 한다면
- 클리어한 스테이지를 몇 성으로 클리어했나 정도만 추가하면 될 것 같음
	- 추가하고 오늘은 여기까지. 이게 생각보다 안 이쁘다.
- 저 선에 셰이더 그래프 머티리얼을 씌워서 좌->우로 가는 그라데이션을 추가하는 것도 떠오름
	- **이거 안된다.** 왜 안되는지는 모르겠음. 그라데이션 텍스쳐 + `Tiling And Offset`으로 우->좌로 그라데이션 이미지가 옮겨가도록 했는데, 그냥 검정색 이미지만 뜬다. 
	- 모든 `Source Image`에서 동일한 현상.

