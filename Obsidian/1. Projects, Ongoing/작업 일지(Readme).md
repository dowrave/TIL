# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 보고 참고할 내용
- [[오퍼레이터들 스탯 정리]]
- [[적 스탯 정리]]
- [[프로크리에이트로 텍스쳐 작업 시 유의할 점]]
## 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보고 따라했거나 직접 만듦)
	- 텍스쳐 제작 : `Procreate` / `Krita`
	- 메쉬 제작 : `Blender`

## 하고 싶은데 못한 것
- 2D로 오브젝트들 구현하기
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.
	- 현재는 정예화 패널에서만 스킬 범위를 볼 수 있음.

---
# 작업 일지

## 짭명방 

### 남은 내용

>[!todo]
>- 남은 작업
>	- 스테이지들 만지기 : 스테이지마다 특징을 부여하면 좋을 듯?
>		- `1-0`은 튜토리얼 개념
>		- `1-1`은 바리케이드를 통한 진로 변경
>		- `1-2`는 여러 종류의 적을 보여줌(방어력 강화, 저지 무시) 
>		- `1-3`은 앞에서 배운 것들 활용 + 보스 출현
>- 테스트 중 제대로 나타나지 않는 머티리얼들 수정
>- 테스트 중 발생하는 이슈/오류 수정
>- 눈에 보이는 개선이 필요한 사항 수정
>	- 소리 추가
>- 기타 이슈들 수정

### 현재 진행 중
- `Todo` 내용 작업 진행
- `ClickDetectionSystem` 정리 중
	- 유니티에서 자체적으로 동작하는 `EventSystem`과는 완전히 별도로 동작시키는 게 목표
	- `DeployableActionUI`의 버튼들이 정상적으로 동작해야 한다.  현재는 클릭해도 아무 동작이 없음.


### 간헐적인 이슈
- 이슈가 있다고 느꼈는데 다시 테스트했을 때 재현이 안된 것들을 정리함
	- 오퍼레이터 A를 배치할 때, 방향 설정 로직 중 오퍼레이터 B의 위치에서 마우스 커서를 떼면 배치되면서 해당 마우스 커서의 위치에 있는 오퍼레이터가 클릭되는 현상

---

# 251023 - 짭명방

>[!done]
>- 버그 수정
>	- 바리케이드가 풀로 돌아가지 않고 맵에 남는 현상
>	- `DeployableUnitEntity` 재배치 시 렌더링 이슈

>[!wip]
>- 버그 수정
>	- `DeployableActionUI`가 정상적으로 동작하지 않음
>		-> 클릭 감지 시스템 총체적인 개편 작업으로 이어짐

## 버그 수정

### 바리케이드 풀로 돌아가지 않는 현상
- `Barricade` 프리팹에 `Barricade``SO`가 할당되지 않아서 풀 태그가 정상적으로 등록되지 않는 현상이었다. 어제 물음표를 띄우던 거에 비해 간단하게 해결했음

### DeployableActionUI 요소가 정상적으로 동작하지 않음
- **`ActionUI`의 요소인 퇴각, 스킬 사용이 정상적으로 동작하지 않는 현상**이 있다.
- **원래 잘 되던게 안돼서 감이 안 잡힘;;**

- 기존의 `ClickDetectionSystem`에 문제가 있었다. **`EventSystem`과 충돌할 여지가 많은 스크립트였다.** 둘 중 뭐가 먼저 동작하고 나중에 동작하는지 판단하는 것도 애매함.

- 그래서 
1. `ClickDetectionSystem`에서는 `EventSystem`에서 처리하지 못하는 요소들을 처리한다.
	- `EventSystem.current.IsPointerOverGameOject()`로 `UI`요소 위에 커서가 있는지 여부를 점검할 수 있다. 이게 `true`일 때는 `ClickDetectionSystem`에서 어떠한 처리도 일어나선 안됨.
		- 메서드 이름이 애매한데, **`EventSystem`에 의해 감지되는 요소들을 잡아낸다.**
		- 추가) **`Camera`가 `Physics Raycaster` 컴포넌트를 갖고 있다면 콜라이더를 갖고 있는 모든 요소를 감지한다.**
			- 이 부분은 `ClickDetectionSystem`에서 감지하기 때문에 제거하고 진행해도 될 듯? 
2. 빈 공간을 클릭했을 때에 대해서는, 빈 패널을 `Canvas`의 최상단, 즉 가장 먼저 렌더링되어서 가장 아래로 가게끔 한 다음 해당 요소가 클릭되면 `Cancel` 동작이 일어나게 구현하면 된다.

> 기존 방식이 복잡했기 때문에 정리하는 것 자체는 찬성인데 이거는 시간이 좀 걸릴 듯? 일단 츄라이해봄.




### DeployableUnitEntity 재배치 시 렌더링 이슈
- 이거는 `Operator`에서 구현했던 `InitializeVisuals`와 동일한 방법, 즉 `Initialize`에 `InitializeVisuals`을 구현하는 방식으로 해결할 수 있을 것 같다

- 구조를 어떻게 만들지에 대한 생각을 좀 했다. 
	- 지금은 가장 부모 클래스인 `UnitEntity`에 `InitializeVisuals()`을 구현하고 자식 요소들에서 이를 사용하고 있음
	- 지금 방식을 유지할지, 아니면 가장 부모 클래스인 `UnitEntity`에 `Initialize()`으로 구현하고 자식 요소들에서 상속시킬지가 고민된다.
	- 자식 요소들에서는 `Initialize(UnitData unitData)` 같은 느낌인데, `UnitData`를 모두 별개로 쓰고 있기 때문에 `override` 개념도 아니다. 파라미터를 쓰기 때문임. 

- 원래는 **템플릿 메서드 패턴**을 쓰는게 좋은 방법이라고 하는데, 지금 같은 경우는 `SO`끼리 상속되지 않기 때문에 `Initialize()`를 굳이 구분해서 구현하고 있음
- 따라서 **기존 방식대로 `InitializeVisuals()`을 자식 클래스들의 `Initialize(Basedata baseData)`에 별개로 구현**하는 식으로 처리해두겠음

- 근데 애초에 개별 오브젝트였을텐데 왜 투명한 상태로 바리케이드가 등장했는지는 ㅁ?ㄹ

# 251022 - 짭명방

>[!done]
>- 구조 변경
>	- `DeployableInfo`를 `DeployableManager`에서 분리
>- 버그 수정
>	- `InstageInfoPanel` : 바리케이드를 클릭했는데 해당 정보로 업데이트되지 않는 현상

## 구조 변경
- `DeployableInfo`를 `DeployableManager`에서 분리해서 개별 스크립트로 만듦
```cs
    // 일반 배치 가능한 유닛일 때 할당
    public List<DeployableUnitEntity> deployedDeployables = new List<DeployableUnitEntity>();
    public DeployableUnitData? deployableUnitData;
```
> `deployedDeployables`을 `List`로 수정했다 : 바리케이드는 여러 개가 배치될 수 있기 때문임.

~~그리고 배치된 `Deployables`을 관리하는 로직도 추가해야 한다. ~~

****
## 사소한 이슈 수정

### 스테이지 `1-1` : `PreloadStageObjectPools` 이슈
![[Pasted image 20251022171238.png]]
> 항목이 있는데도 `Element`가 나타나지 않는 현상

- 해결) `MapDeployableData`에서
```cs
[SerializeField] protected GameObject deployablePrefab = default!;
[SerializeField] protected DeployableUnitData deployableData = default!; // 기본 데이터. 별도의 초기화 로직을 거치지 않게 하기 위함. 
[SerializeField] protected int maxDeployCount = 0; // 최대 배치 가능 수 
```
> 이전엔 `public`이었다가  `protected`로 바꿨는데, 이 경우 인스펙터에서 렌더링할 요소를 찾지 못한다. 그래서 앞에 `[SerializeField]`를 추가함.
> 접근자 자체는 저렇게 관리해주는 게 좋다. 외부에 의한 변경을 막아야 하기 때문.

### `DeployableUnitEntity` : 인게임 정보 인포 갱신 안됨
- 오퍼레이터 정보는 잘 나타나고 있고 바리케이드 같은 `DeployableUnitEntity`를 클릭했을 때 정보 갱신이 되지 않았음
- `InstageInfoPanel`에 로깅해보니
```cs
public void UpdateInfo(DeployableManager.DeployableInfo deployableInfo, bool IsClickDeployed)
{
	gameObject.SetActive(true); // 정보 업데이트라서 오브젝트 활성화
	currentDeployableInfo = deployableInfo;
	currentDeployableUnitState = DeployableManager.Instance!.UnitStates[currentDeployableInfo];

	if (currentDeployableUnitState == null)
	{
		Debug.LogError("현재 배치 요소의 정보가 없음");
		return;
	}

	// 박스에서 꺼내는 요소인 경우, 맵의 남은 부분을 클릭하면 현재 동작을 취소할 수 있음
	if (!IsClickDeployed)
	{
		cancelPanel.gameObject.SetActive(true);
		cancelPanel.onClick.AddListener(OnCancelPanelClicked);
	}

	// 1. 배치된 Operator
	if (currentDeployableInfo.deployedOperator != null)
	{
		Operator op = currentDeployableInfo.deployedOperator;
		op.OnDeathStarted += Hide;
		CameraManager.Instance!.AdjustForDeployableInfo(true, op);
		UpdateOperatorInfo();
	}
	// 2. 박스의 Operator
	else if (currentDeployableInfo.ownedOperator != null)
	{
		OwnedOperator op = currentDeployableInfo.ownedOperator;
		CameraManager.Instance!.AdjustForDeployableInfo(true);
		UpdateOperatorInfo();
	}
	else if (currentDeployableInfo.deployedDeployable != null)
	{
		DeployableUnitEntity deployable = deployableInfo.deployedDeployable;
		deployable.OnDeathStarted += Hide;
		CameraManager.Instance!.AdjustForDeployableInfo(true, deployable);
		UpdateDeployableInfo();
	}
	else
	{
		Debug.LogError("UpdateInfo를 클릭했으나 갱신된 정보가 없음");
	}   
}
```
> 여기서 `else`문이 출력되고 있음 
> 따라서 `deplyoableDeployable`이 제대로 할당되지 않고 있다는 거고, 이건 `DeployableManager`의 초기화에 관한 이슈로 보임

.. 이게 생각보다 복잡한 문제가 됐다. `deployedDeployable`이 1개가 아니라서 발생하고 있는 문제임. (바리케이드라는 데이터 1개로 3개의 개체까지 관리할 수 있기 때문)

- 일단 스크립트를 통째로 받아서 AI에게 넘겨주니까 큰 규모의 리팩토링을 해줬는데, **내 능지로는 이해가 안돼서 다시 어제 커밋으로 되돌린 다음 더 간단하고 이해가 가는 버전으로 수정**하고 있다. 

- `InStageInfoPanel` 부분에 관해 
	- `UpdateInfo`에 파라미터를 받아서 **배치된 요소/배치되지 않은 요소 처리 메서드를 나눴다.**
	- 그리고 각각의 메서드는 `DeployableUnitEntity`와 `Operator`를 구분해서 처리하는 방식이다.
	- 지금 보면 `StageUIManager`에서는 `Info`와 배치된 객체를 구분해서 처리했느데, `InstageInfoPanel`에서는 `UpdateInfo(DeployableInfo)`로 억지로 통합해서 처리해서 코드가 더 복잡해보이는 느낌이 있음.
	- 이런 걸 보면 **스크립트가 길어지더라도  혼동 가능성을 줄이는 게 더 바람직한 코드인 것 같음.** 지금 보면 괜히 꼬여있는 느낌이 있다.

- 일단 이렇게 처리했을 때 바리케이드를 클릭했을 때 바리케이드 UI가 나타나고 있어서 1번째 문제 해결은 성공


---
# 251021 - 짭명방

>[!done]
>- 버그 수정
>	- `Enemy` : `Operator`가 사라졌는데도 공격하는 현상
>	- `MeleeAttackEffect`와 `HitEffect`가 섞이는 현상 : 정확히는 `HitEffect`가 나타나야 할 자리에 `MeleeAttackEffect`가 섞이는 현상에 가깝다. 
## 버그 수정하기

### Enemy) Operator가 사라졌는데도 공격하는 현상
- 예전에 `static event`로 오퍼레이터가 죽었을 때 모든 적에게 알리는 코드를 구현했던 거 같은데 없다? 
#### 해결
1.  `DeployableUnitEntity`
```cs
public static event Action<DeployableUnitEntity> OnDeployableDied = delegate { };
```
> 이벤트 추가, `Despawn`에서 `Invoke`

2. `Enemy`
 - ~~`Initialize`에서 전역 이벤트 등록, `Despawn`에서 전역 이벤트 해제~~
 - `OnEnable`에서 등록, `OnDisable`에서 해제
```cs
public void HandleDeployableDied(DeployableUnitEntity disabledEntity)
{
	if (targetsInRange.Contains(disabledEntity))
	{
		targetsInRange.Remove(disabledEntity);

	}

	// 타겟이 범위에서 벗어났어도 현재 타겟으로 지정되는 상황이 있을 수 있으니 위 조건과는 별도로 구현했음
	if (CurrentTarget == disabledEntity)
	{
		CurrentTarget = null;
	}
}
```
- `CurrentTarget` 프로퍼티의 세터의 이벤트 발생 메서드는 제거했음.

> [!done]
> 1. `Operator`가 사라진 자리를 `Enemy`가 계속해서 공격하는 문제
> 2. 미리보기 상태인데도 `Enemy`가 해당 미리보기 객체를 향해 공격하는 문제

> [!tips]
> - **개별 객체에서 전역 이벤트를 등록하고 해제하는 로직은 다른 객체에게 영향을 주지 않는다.**  갑자기 궁금해져서 AI에게 물어봤음.
### HitVFX와 MeleeAttackVFX가 번갈아가며 나타나는 현상
- 어제 해결하려다가 해결 못한 문제
- 이건 답을 못 찾았기 때문에 관련 요소들 긁어다가 AI한테 던져줬다. 이 부분이 문제였음.

```cs
private IEnumerator WaitAndReturnToPool(float duration = 1f)
{
	yield return new WaitForSeconds(duration);

	if (gameObject != null)
	{
		ObjectPoolManager.Instance!.ReturnToPool(_tag, gameObject);
	}
	
	// else문은 필요 없음 - null이면 Destroy 호출 불가능
}
```
> - 여기서 `_tag` 부분이 `attackSource.HitEffectTag`으로 되어 있었다. 이걸 `_tag`로 수정해서 인풋으로 받은 태그로 풀을 돌려줬음.

> [!done]
> 1. 오브젝트 풀을 이용하지 않고 계속 인스턴스를 생성하는 문제
> 2. `hitEffect`가 나타나야 할 자리에 `MeleeAttackEffect`가 나타나는 문제

---
# 251020 - 짭명방

>[!done]
>1. 테스트하면서 나타나는 문제들 수정
>	- 재배치 로직
>		- 재배치 시 유닛 미리보기에서 방향 표시기가 나타나는 이슈
>			- **`Awake` 시점에 `StageUIManager`에 할당된 방향 표시기 프리팹을 이용해 자식 오브젝트로 생성**함
>		- 죽어서 사라진 유닛 재배치 시 머티리얼 초기화 안 된 이슈
>	- 사망 후 재배치 시 유닛 렌더링 문제

>[!wip]
>- `SO`의 오브젝트 풀을 제대로 사용하지 않는 문제 수정

## 점검과 수정
- 퇴각 후 재배치했을 때 의도한 구현들이 정상적으로 동작하는지 테스트
- 까먹을까봐 기록해 둠 : 뱅가드는 30초, 나머지는 70초

### 1. (해결)재배치 시 미리보기 상태에서 방향 표시기가 나타남
```cs
protected void CreateDirectionIndicator()
{
	// 자식 오브젝트로 들어감
	DirectionIndicator indicator = Instantiate(StageUIManager.Instance!.directionIndicator, transform).GetComponent<DirectionIndicator>();
	indicator.Initialize(this);

	// 오퍼레이터가 파괴될 때 함께 파괴되므로 전역변수로 설정하지 않아도 됨
}
```
> 방향 표시기는 `Operator.Deploy()`에서 이런 식으로 인스턴스화되는 요소임

1. 게임을 플레이하는 중에 `Instantiate`가 동작하는 게 바람직한 방식은 아니다. 
2. 프리팹에 `DirectionIndicator`를 자식 오브젝트로 할당하는 방법
	- 모든 오퍼레이터가 공통으로 사용하는 필드를 일일이 할당하는게 맞나? 이런 식으로 작업한 게 있긴 한데 바람직한 방식은 아니라는 생각이 듦. 규모가 조금만 더 커져도..

아래의 `가상의 이슈`를 거쳐서...  

정리해보면
>[!conclusion]
>1. 이미 `Operator`는 오브젝트 풀링으로 관리되고 있음
>2. 따라서 오브젝트 풀이 생성되는 시점의 `Awake`에서 `StageUIManager`에 있는 `DirectionIndicator`를 생성함

정도가 되겠다. 이 부분은 오브젝트 풀링으로 구현할 필요도 없다. 한 번 만들면 계속 `Operator`의 자식 오브젝트로 남아 있고 활성화 / 비활성화 여부만 조절해주면 되기 때문.

#### 추가 이슈
- (완료) 방향이 제대로 반영되지 않는 이슈
	- `DirectionIndicator`의 `Initialize` 시점은 배치되는 시점이어야 함. 생성되는 시점이 아님.

---
#### 가상의 이슈
>[!question]
>- 지금의 시나리오는 "모든 오퍼레이터가 동일한 방향 표시기를 사용한다"이다.
>- 만약 방향 표시기 스킨이 있다고 가정한다면 이 방향 표시기 프리팹 필드의 위치는 어디로 지정해야 할까?

- 방향 표시기 프리팹 필드의 적절한 위치
1. `전역 매니저` : 모든 오퍼레이터가 공통된 표시기를 쓰며, 커스터마이징이 없을 때
	- **디폴트 프리팹은 아예 여기에 둬도 무방!**
2. 커스텀 개념, 즉 스킨 개념이 있는 경우에는 아예 스킨만을 위한 `SO`를 별도로 구현하고, `소유한 오퍼레이터` 데이터에 어떤 스킨들을 갖고 있는지에 관한 필드를 구현하면 됨.
	- 여기서 **`SO`는 해당 오퍼레이터에 관한 모든 정보**이고, **`소유한 오퍼레이터`는 유저가 갖고 있는 정보**들이다. 성장, 스킨 보유 여부 등에 관한 정보들. 

> 이거는 '이러이러한 상황이면 어떻게 되지?'라는 생각이 들어서 정리해둔 내용이다. 
---
#### 지식이 늘었다
>[!note]
>1. `MonoBehaviour.IsActiveAndEnabled` : 오브젝트와 스크립트 모두 활성화됐을 때만 `true`를 반환한다.
>	- 일반적으로 스크립트를 비활성화하는 경우가 잘 없어서 혼동하기 쉬운데 차이점은 명확히 알아두면 좋다.
>	- 오브젝트 활성화 : `gameObject.activeInHierarchy`
>	- 스크립트 활성화 : `script.enabled`

### 2. 사망 후 재배치 시 투명한 상태로 등장
- (수정 완료)`Initialize`에서 렌더링 초기화를 해줘야 할 듯
	- 유닛 셰이더의 `FadeAmount` 값을 수정했던 내용이므로 `Initialize`에
```cs
protected void InitializeVisuals()
{
	foreach (Renderer renderer in renderers)
	{
		renderer.GetPropertyBlock(propBlock);
		propBlock.SetFloat("_FadeAmount", 1f);
		renderer.SetPropertyBlock(propBlock);
	}
}
```
을 추가해줬다.
## SO에서 만든 풀을 제대로 사용하지 않았던 문제 수정
- 1번의 이슈 외에도 `배치 이펙트(SO에서 가져옴)`, `OperatorUI(Operator 자체에 할당)`이 로딩 시점에 인스턴스로 만들어지지 않고 인게임에서 만들어짐
- 이외에도 **`SO`의 태그를 이용해 만든 풀을 사용하지 않고 있던 요소들을 사용하도록 수정함** 


#### 배치 이펙트
- `SO`에서 가져오는 방식의 경우 예전에는 `OperatorData`를 할당하는 시점에 관한 이슈가 있었으나, 지금은 프리팹의 스크립트 자체에 `SO` 필드를 할당해놨음. 따라서 생성하는 시점을 `Awake`으로 옮길 수 있다. 
- 근데 `OperatorData.CreateObjectPools`를 구현하는 과정에서 배치 이펙트에 관한 오브젝트 풀을 만드는 로직을 이미 추가했다. 그렇기 때문에 만드는 로직은 별도로 구현할 필요가 없음

#### OperatorUI
- `DirectionIndicator`와 비슷한 방법으로 관리되면 좋을 듯 : 즉 `StageUIManager`에서 관리하고 `Operator`들이 접근하는 방식

#### Projectile
- `OperatorData`에서 만드는 그걸 안 쓰고 `Operator` 자체의 것을 사용하고 있었음

#### AttackSource
- `hitEffect` 관련, 원래는 프리팹이랑 태그를 함께 첨부했는데 **태그가 있으면 프리팹을 첨부할 필요는 없어보임**

```cs
protected virtual void PlayGetHitEffect(AttackSource attackSource)
{
	GameObject sourceHitEffectPrefab = attackSource.HitEffectPrefab;
	string sourceHitEffectTag = attackSource.HitEffectTag;

	if (sourceHitEffectPrefab == null || sourceHitEffectTag == string.Empty) return;

	string attackerName;

	if (attackSource.Attacker is Operator op)
	{
		OperatorData opData = op.OperatorData;
		attackerName = opData.entityName;
		if (sourceHitEffectPrefab == null)
		{
			sourceHitEffectPrefab = opData.hitEffectPrefab;
		}
	}
	else if (attackSource.Attacker is Enemy enemy)
	{
		EnemyData enemyData = enemy.BaseData;
		attackerName = enemyData.EntityName;
		if (sourceHitEffectPrefab == null)
		{
			sourceHitEffectPrefab = enemyData.HitEffectPrefab;
		}
	}
	else
	{
		Debug.LogError("이펙트 없음");
		return;
	}

	if (sourceHitEffectPrefab != null)
	{
		Vector3 effectPosition = transform.position;

		// 풀에서 이펙트 오브젝트 가져오기
		if (sourceHitEffectTag == string.Empty)
		{
			Debug.LogWarning("[TakeDamage] hitEffectTag 값이 null임");
		}

		GameObject? hitEffect = ObjectPoolManager.Instance!.SpawnFromPool(sourceHitEffectTag, effectPosition, Quaternion.identity);

		if (hitEffect != null)
		{
			CombatVFXController hitVFXController = hitEffect.GetComponent<CombatVFXController>();
			hitVFXController.Initialize(attackSource, this, sourceHitEffectTag);
		}
	}
}
```
> `TakeDamage`의 이펙트를 실행하는 코드인데, 상당 부분이 필요없어짐. 다 `AttackSource`의 `Tag` 하나만으로 처리가 가능해졌다.

```cs
protected virtual void PlayGetHitEffect(AttackSource attackSource)
{
	string sourceHitEffectTag = attackSource.HitEffectTag;

	if (sourceHitEffectTag != string.Empty)
	{
		Vector3 effectPosition = transform.position;
		GameObject? hitEffect = ObjectPoolManager.Instance!.SpawnFromPool(sourceHitEffectTag, effectPosition, Quaternion.identity);

		if (hitEffect != null)
		{
			CombatVFXController hitVFXController = hitEffect.GetComponent<CombatVFXController>();
			hitVFXController.Initialize(attackSource, this);
		}
	}
}
```

- 여기 부분 생각보다 고칠 게 많아서 일일이 정리하진 않겠음. 작업 내용은 크게
1. `StageUIManager`에서 공통적인 UI 요소들을 관리하고, 각 유닛에서 사용하는 거
2. `AttackSource` 관련 - 프리팹 필드 삭제
3. 이외에 유닛 자체에서 생성되는 오브젝트 풀 -> `SO` 기반으로 변경

이 되겠다. 그 중에서 특기할 문제만 기록하겠음.

#### HitVFX가 나타나야 할 자리에 MeleeAttackVFX도 번갈아가면서 나타나는 현상
- 로깅을 해봤는데 원인을 못 찾겠음

- 현상
	- 최초에 생성된 프리팹들은 정상적으로 재생됨
	- 최초에 생성된 풀 갯수만큼 재생된 후
		- 오류 1)`AttackEffect`가 새로 생성되며, 이게 `PlayMeleeAttackEffect`에 의해 나타나는 듯(오퍼레이터 위치에 휘두르는 이펙트가 나타남)
		- 오류 2)`HitVFX`가 나타나지 않고 피격 대상 위치에 `AttackEffect`가 나타남

- 근데 생성되는 프리팹 이름은 멀쩡함

...왜인지 모르겠다. 나중에 시도해봄.




# 251017 - 짭명방
>[!done]
>- 스테이지 설정 변경 : 생성 / 파괴 기반 -> 스테이지 준비 과정에서 생성한 다음 오브젝트 풀링으로 관리
>
>  여기 써놓은 건 오늘 다 구현 끝낸 듯 함
>- 앞으로는 **문제 발생 시 수정해나가는 정도로만 구현하면 충분할 듯**
>1. 스킬은 어떻게 할 건지(기존 : `UnitEntity - Caster`)
>2. 유닛들 자체도 어떻게 할 건지(태그를 어디서 갖고 있을 건지)
>3. `ObjectPoolManager`의 활성화는 스테이지 로딩 시점에 이미 되어 있어야 함
>4. `DeployableUnitEntity(Operator)`의 `Initialize` 및 `Deploy`, `Enemy`의 `Initialize` 등등의 로직들도 생성 / 파괴 기반에서 오브젝트 풀 활성화 / 비활성화로 바꿔야 함
>	- 현재 작업 중. 
>	- 생성 / 배치, 퇴각 / 사망 등이 오브젝트 풀링에 기반해서 작동되는지 오퍼레이터 / 적 별로 일일이 확인해줘야 함
>	- 251016 기준 수정 필요
>		- 퇴각해서 본 모델은 사라졌는데 자식 오브젝트들이 남아있는 현상이 있음
>5. `BossSkill`도 게임 시작 시점에 준비되어 있어야 함

## (계속) 유닛 생성 / 파괴 로직 오브젝트 풀링 기반으로 변경

### 어제의 이슈

>[!note]
>퇴각 / 사망 시 본 모델은 사라졌는데 자식 오브젝트들이 남아있는 현상

- 일단 `Retreat`에 대응할 `DieInstantly` 메서드를 `UnitEntity`에 구현해뒀다. `Die`와 비슷하게 동작하지만 애니메이션 후에 사라지는 개념이 아니라 바로 사라지는 개념.
- 이걸 연결한 다음 `Retreat`를 해봤는데 풀이 없다는 오류가 나옴
- 해결
	- `SetPoolTag()`을 구현해놓고 실행을 안 했다.
	- 원리) 오브젝트 풀이 생성되는 시점, 즉 오브젝트가 만들어지고 비활성화될 때 `Awake()`에서 먼저 동작하고 비활성화된다.
	- 실행) 따라서 `UnitEntity`에 빈 메서드 `SetUnitTag()`을 구현하고 `Awake()`에서 실행시킨 다음, 자식 클래스들에서 메서드의 구체적인 내용을 구현하면 됨.


## EnemyBoss 관련 요소들도 시작 시점에 구현
- 이 과정에서 `ScriptableObject(SO로 축약함)`에 들어간 태그 이름을 설정하는 부분을 수정했다.
	- 기존) `UnitEntity(런타임 중에만 활성화)`에 의존
	- 수정) `SO`의 엔티티 이름을 쓰거나, 스킬 이름만으로 설정
		- 스킬 이름만 쓴다면 발생할 수 있는 이슈? : 여러 오퍼레이터가 동일한 스킬(예를 들면 공격력 버프)을 사용한다면?

## 기타 수정사항
### `Enemy` 유닛 풀 
- `Enemy`의 종류와 수를 체크하고 그것에 맞게끔 정량으로 생성하면 된다. 지금은 `Enemy`마다 별도로 구현하고 있음. 
- 이런 식으로 변경하면 됨
```cs
// var uniqueEnemyPrefabs = new HashSet<GameObject>();
var enemyPrefabCounts = new Dictionary<GameObject, int>();

foreach (var spawner in currentMap.EnemySpawners)
{
	if (spawner.spawnList == null) continue;
	foreach (var spawnInfo in spawner.spawnList.spawnedEnemies)
	{
		if (spawnInfo.prefab != null)
		{
			// 이미 존재하면 갯수 추가
			if (enemyPrefabCounts.ContainsKey(spawnInfo.prefab))
			{
				enemyPrefabCounts[spawnInfo.prefab]++;
			}
			// 없다면 딕셔너리에 새로 추가
			else
			{
				enemyPrefabCounts.Add(spawnInfo.prefab, 1);
			}
		}
	}
}
```
> `ContainsKey`로 키값이 있는지 확인, 있다면 갯수만 추가하고 없다면 새로운 키-값을 만들어준다.
> 자꾸 까먹는 부분이니까 어떻게 쓰는지 체크해두자.

```cs
foreach (var entry in enemyPrefabCounts)
{
	GameObject enemyPrefab = entry.Key;
	int requiredAmount = entry.Value;
	
	// ...
}
```
> 기존 해쉬셋에서 바꾸기도 편리하다. 기존엔 `enemyPrefab`만을 사용했다면 이번엔 `Key, Value`값으로 둘 모두를 꺼내서 쓰는 방식.
> 여기서 `var`은 `System.Collections.Generic.KeyValuePair<TKey, TValue>`이라는 타입이다.

### Enemy : Operator가 없어졌는데 계속 그 자리 공격
- (미해결)원거리 적 유닛이 아군 오퍼레이터가 사라졌는데도 그 자리를 공격하는 이슈가 있음

### Enemy : Projectile 자체 생성 -> Data 기반 생성으로 변경
- 수정 완료
	-  `SO`에 의해 게임이 시작되는 시점으로 생성 시점을 옮겼는데, `Enemy` 자체에서 초기화될 때 오브젝트 풀을 생성하는 로직이 남아 있었음. 실제 사용되는 투사체도 `Enemy` 자체에서 생성한 오브젝트 풀이었다.




# 251016 - 짭명방

>[!wip]
>- 유닛 생성, 파괴 로직 오브젝트 풀링 기반으로 변경

## (계속) 유닛 생성 파괴 로직 오브젝트 풀링 기반으로 변경

### 일단
- `251014`에 작업한 내용 중 `DeployableUnitEntity` 등에서 `Data`를 가져오는 부분은 수정이 필요하다 - 얘네는 런타임 중에 할당되는 요소라서 사전 준비 작업에는 `Data`가 할당되지 않은 상태임

- 이 부분을 다 날리고
```cs
// c. 수집된 모든 배치 가능 유닛 프리팹 풀링
// foreach (var deployablePrefab in uniqueDeployablePrefabs)
// {
//     DeployableUnitEntity deployableUnit = deployablePrefab.GetComponent<DeployableUnitEntity>();
//     if (deployableUnit != null)
//     {
//         if (deployableUnit is Operator op)
//         {
//             OperatorData opData = op.OperatorData;
//             string opPoolTag = opData.GetUnitTag();
//             ObjectPoolManager.Instance.CreatePool(opPoolTag, deployablePrefab, 1);
//         }
//         else
//         {
//             DeployableUnitData deployableData = deployableUnit.DeployableUnitData;
//             string deployablePoolTag = deployableData.GetUnitTag();
//             ObjectPoolManager.Instance.CreatePool(deployablePoolTag, deployablePrefab, 1);
//         }
//     } 
// }
```

- Data에서 꺼내는 부분에서 한꺼번에 풀 생성까지 작업함
	- `Data`를 할당하는 부분의 변경
		- 기존) 인게임에서 각 객체가 초기화되는 시점에 할당되는 `Data`를 사용
		- 변경) 외부에서 `Data`를 갖는 별도의 클래스를 구성
```cs
// a. 스쿼드
foreach (var opInfo in squadData)
{
	if (opInfo.op.OperatorProgressData != null)
	{
		OperatorData opData = opInfo.op.OperatorProgressData;
		GameObject operatorPrefab = opData.prefab;

		string opPoolTag = opData.GetUnitTag();
		ObjectPoolManager.Instance.CreatePool(opPoolTag, operatorPrefab, 1);
	}
}
// b. 맵 전용
if (stageData != null && stageData.mapDeployables != null)
{
	foreach (var mapDeployableData in stageData.mapDeployables)
	{
		if (mapDeployableData != null)
		{
			DeployableUnitData deployableData = mapDeployableData.DeployableData;
			GameObject deployablePrefab = deployableData.prefab;

			string deployablePoolTag = deployableData.GetUnitTag();
			ObjectPoolManager.Instance.CreatePool(deployablePoolTag, deployablePrefab, 1);
		}
	}
}
```

- 다만 **`Enemy`의 경우**는 조금 작업이 달라지는데 `Spawner`에 할당된 독립된 유닛들을 수집해서 거기서 `Data` 정보를 빼내는 구조였다. 얘는 아예 구조를 새로 만들어줘야 할 듯?
	- 근데 **프리팹의 스크립트 필드에 이미 할당되어 있으면 저거도 쓸 수 있지 않나?**
		- 기존 프리팹의 구조를 보면 `Operator`에는 `OperatorData`가 `SerializeField`로 사전에 할당되지 않은 반면 `Enemy`에는 할당되어 있다. 
		- 일단 프리팹에 전부 자신에 해당하는 `OperatorData`나 `EnemyData`를 할당했다. 


- 일단 이 과정에서 하나 얻고 가는 지식
>[!note]
>- **`[SerializeField] protected` 필드와 이 필드에 관한 게터 프로퍼티를 구현했다면 그냥 이 2가지만 쓰면 된다. 굳이 세터 프로퍼티를 따로 구현해서 코드를 번거롭게 길게 만들 필요는 없음.**
>- 이 때 **필드 이름은 `_`을 붙이고 프로퍼티 이름은 대문자로 시작**하는 게 표준이라고 함
>	- 필드 : `_deployableData`
>	- 게터 프로퍼티 : `DeployableData => _deployableData`

하나하나 수정해나가고 있는데 뭔가 많이 꼬여 있다. 시간 좀 걸릴 듯.

### 수정된 내용들
- `Box`를 클릭했을 때, 드래그했을 때 동작들 다시 점검해봐야 할 듯

- `DeployableManager` 싹 뜯어보는 중
	1. 박스에서 다시 꺼내는 과정은 수정 완료
	2. 타일 위에 올렸을 때 배치 안되는 문제 있음 - 수정 완료
	
- 각 `UnitEntity` 상속 유닛들에 `poolTag` 설정 할당하고 `PlayDeathAnimation`에서 유닛을 `Destroy` 하는 대신 `ReturnToPool`로 수정

- 배치 중에 취소했을 때 `ReturnToPool` 구현되어 있음

- `ResetPlacement` 구현 관련, 배치 후에는 오브젝트가 풀로 반환되면 안됨
	- 그래서 이 둘을 구분해서 `ResetPlacement` 메서드를 구현할까 했는데 `currentDeployable.IsPreviewMode`로 조건을 통제하는 게 있어서 이걸로 제어함

- `Box` 색깔 반영 안되는 문제도 수정


# 251014 - 짭명방

## (계속)유닛 생성 파괴 -> 오브젝트 풀링 기반으로 변경 
- 남은 과제들
>[!note]
>1. `ObjectPoolManager`의 활성화는 스테이지 로딩 시점에 이미 되어 있어야 함
>2. `DeployableUnitEntity(Operator)`의 `Initialize` 및 `Deploy`, `Enemy`의 `Initialize` 등등의 로직들도 생성 / 파괴 기반에서 오브젝트 풀 활성화 / 비활성화로 바꿔야 함
>3. `BossSkill`은 어떻게 할 건지

### 1. ObjectPoolManager의 활성화 시점 
- `StageLoader`의 `LoadStage`를 보면
```cs
// 비동기 씬 로드
AsyncOperation asyncLoad = SceneManager.LoadSceneAsync(STAGE_SCENE);
asyncLoad.allowSceneActivation = false;

// 씬 로드 진행 - 로딩이 90% 되었을 때 다음으로 넘어간다.
while (asyncLoad.progress < 0.9f)
{
	yield return null;
}

// 최소 로딩 시간 보장
float elapsedTime = Time.time - loadStartTime;
if (elapsedTime < MIN_LOADING_TIME)
{
	yield return new WaitForSeconds(MIN_LOADING_TIME - elapsedTime);
}

asyncLoad.allowSceneActivation = true;

// 씬의 모든 활성 오브젝트의 Awake 메서드가 호출됨
// 싱글턴은 Awake에서 설정되므로 ObjectPoolManager도 이 시점에서 초기화됨
while (!asyncLoad.isDone)
{
	yield return null;
}
```

다만 더 나은 구조를 갖게 하고 싶다면, `StageManager`에서 `ObjectPoolManager` 필드를 할당하고 초기화 순서를 더 명료하게 할 수 있음. 지금은 `StageLoader`가 `StageManager`와 `ObjectPoolManager`의 초기화가 끝났다고 가정하는 상황이기 때문이다.

일단 그렇다는 것만 알아두고 지금 구조를 유지한다. 나중에 문제가 발생하면 그 때 수정하겠음.

### 2. 유닛 초기화 : 생성 -> 풀링 기반으로 변경
- 유저가 직접 배치하는 **배치 요소**와 **적** 2가지로 나뉜다.
- 어느 쪽이든 풀 태그를 가져와야 하는 식으로 바뀌기 때문에 풀 태그를 어떻게 가져올 건지만 설정하면 되겠다.

- 공통적으로 `UnitEntity`에 자체적으로 `PoolTag` 필드 하나 설정함
#### 배치 요소
- 기존엔 `currentDeployableInfo.Prefab`을 가져와서 인스턴스화를 했는데 여기에도 `poolTag`를 넣어주면 됨. 프리팹을 대체할 수 있을 것 같긴 한데 일단 구현하면서 보겠음

- **하단 `Box`는 이미 `SetActive`로 관리되는 요소임** - 수정할 필요 없음

작업하다가 멈췄음 - 이 날은 여기까지

# 251013 - 짭명방

>[!done]
>1. `UnitEntity` 사망 애니메이션 머티리얼 변경하는 로직 수정

>[!wip]
>1. 생성 / 파괴 -> 오브젝트 풀링 기반으로 변경 시작

## `UnitEntity` 사망 애니메이션 변경 로직 수정

- 기존 방식
	- 별도의 머티리얼 인스턴스를 만들고, 투명 렌더링 모드로 바꿔서 머티리얼들에 `DOTween`의 페이드아웃 트윈을 만들어서 시퀀스로 넣는 방식
	- 이 과정에서 머티리얼을 수정하는 메서드가 있었고, 이는 `Opaque -> Transparent`로 변경하는 로직이 들어가 있음(이후 투명화 애니메이션을 구현하기 위해)
- 수정
	1. (2번에 통합) **메쉬 머티리얼을 `Opaque -> Transparent`로 변경**
		- `Transparent`가 `Opaque`보다 더 무겁다 : `Opaque`는 `Z-Buffer`를 통해 다른 불투명 픽셀보다 뒤에 있다면 그리는 연산 자체가 생략되는 반면, `Transparent`는 한 픽셀에 여러 번의 연산이 들어가기 때문임
		- 하지만 크게 고려될 요소는 아니다. 
	2. **아예 `URP/Lit`을 사용한 셰이더를 새로 만듦**
		- 기존에 적용되던 프로퍼티들(Metalic, Smoothness 등등)의 이름은 그대로 넣고 `FadeAmount`만 새로 적용해서 `Map * Color -> Split(Alpha)`에 곱해서 알파값으로 전달하는 방식임
		- 이러면 `Ghost`, 즉 알파값이 주기적으로 커졌다가 작아졌다가 하는 셰이더와 통합할 수 있겠다는 생각도 들었는데 웬만하면 분리하는 게 좋다. **범용적으로 사용하는 셰이더와 특정 유닛만을 위한 셰이더는 따로 구현해놓는 게 나음.** 이것도 단일 책임 원칙에 해당한다. 굳이 범용적으로 사용하는 셰이더 내부에 주기적으로 불투명도를 조절하는 스위치 기능을 넣을 필요는 없다는 말이다.
		- 대신 일관성을 위해 `Ghost`에 사용되던 셰이더를 방금 만든 `Unit` 셰이더를 복붙한 다음, `Alpha`만 주기적으로 변하는 방식으로 수정했음

### 추가 발생 문제
![[Pasted image 20251013152719.png]]
이런 식으로 UI와 오브젝트가 겹칠 때, UI가 오브젝트를 뚫고 보이는 현상이 있음. 오브젝트의 알파값은 1.

#### 이유
- 위에서 말한 `Z-Buffer` 이슈로, **투명한 오브젝트는 `Z-Buffer`에 깊이 정보를 입력하지 않는다.** 만약 기록한다면 그 뒤에 그려져야 하는 다른 투명 오브젝트가 가려지는 문제가 발생하기 때문이다.
- 따라서 `Transparent`에서는 `ZWrite(=Depth Write)`가 `Off`로 꺼져 있음.
#### 해결
- **셰이더나 머티리얼의 설정에서 `Depth Write`를 보면 `Auto`로 되어 있는데 이를 `ForceEnabled`로 켜면 됨**

> 대신 만약 다른 투명 오브젝트와 겹쳐질 때 렌더링 순서에 의해 일부 오브젝트가 그려지지 않는 문제가 발생할 수 있음

## `UnitEntity`도 오브젝트 풀링 기반으로 구현

- 기존 방식
	- 생성되어야 할 때 생성되고, 파괴되어야 할 때 파괴됨

오브젝트 풀링을 사용하는 이유는 **인스턴스화 / 파괴 로직이 연산을 많이 잡아먹기 때문에 가능한 게임 중에 실행되는 걸 방지하기 위함**이었다. 그래서 이 부분도 오브젝트 풀링으로 따로 빼두겠음. 스테이지 로딩 씬에서 설정되는 게 좋겠다.

근데 그러면 **스킬 관련 오브젝트 풀도 전부 게임 시작 시점 - 즉 로딩되는 시점에 생성되어야 할 듯?** 
- 지금은 각 유닛이 등장(생성, 배치)하는 시점에 해당 요소들이 생성되는데, 이것도 저 오브젝트 풀링을 쓰는 이유랑 어긋나는 면이 있음. 얘네들도 고려해봐야 할 듯.

- 그러면 크게 2가지로 분류된다.
	1. 스테이지에서 사용되는 유닛 오브젝트 자체
	2. 유닛들이 사용하는 오브젝트들

- 기본적인 원리는 `SkillData`에서 사용하는 것과 동일하다. 
	- 단, 스킬을 사용하는 주체가 생기기 전에 오브젝트 풀이 생성되어야 하므로 참조 대상이 바뀌어야 하는데, **플레이 내내 불변의 데이터로 존재하는 `Data` 및 `Data`에서 관리되는 정보가 그 기준이 된다.**
	- 이 경우 **같은 종류의 다른 유닛은 같은 오브젝트 풀을 사용**하게 된다. 

- 일단 유닛에 딸린 보조 요소들은 거의 구현이 됐음. 
- 남은 과제는..
	1. 스킬은 어떻게 할 건지(기존 : `UnitEntity - Caster`)
	2. 유닛들 자체도 어떻게 할 건지(태그를 어디서 갖고 있을 건지)
	3. `ObjectPoolManager`의 활성화는 스테이지 로딩 시점에 이미 되어 있어야 함
	4. `DeployableUnitEntity(Operator)`의 `Initialize` 및 `Deploy`, `Enemy`의 `Initialize` 등등의 로직들도 생성 / 파괴 기반에서 오브젝트 풀 활성화 / 비활성화로 바꿔야 함
	5. `BossSkill`은 어떻게 할 건지

등등이 있겠다. 


### 1. 스킬 수정
1. `GetVFXPoolTag(UnitEntity caster, GameObject vfxPrefab)`이었는데 이걸 `OperatorData`를 받도록 함
```cs
public virtual string GetVFXPoolTag(OperatorData ownerData, GameObject vfxPrefab)
{
	if (vfxPrefab == null)
	{
		Debug.LogError("[BaseSkill.GetVFXPoolTag] vfxPrefab이 null임!!");
		return string.Empty;
	}

	return $"{ownerData.entityName}_{this.name}_{vfxPrefab.name}";
}
```
> 관련 메서드들도 수정. 어려운 이슈는 아니었는데 괜히 생각을 어렵게 한 듯..

2. 스쿼드 단위에서 스킬을 골라서 들어가는 구조이기 때문에, "선택된 스킬"에 대한 정보는 `OperatorData`에 들어가 있지 않다. 따라서 이는 `StageManager` 자체에서 선택된 스킬의 인덱스를 파악하고 넣는 방식이 됨.
```cs
// 3-b. Opereator의 종속 오브젝트
foreach (var opInfo in squadData)
{
	OperatorData opData = opInfo.op.OperatorProgressData;
	if (opData == null) continue;
	opData.CreateObjectPools();

	// 스킬 오브젝트 풀 생성
	// "선택된 스킬"이라는 정보는 여기나 스쿼드 단위에서 관리되므로 여기서 구현
	int skillIndex = opInfo.skillIndex;
	if (skillIndex >= 0 && skillIndex < 2)
	{
		// 인덱스는 0 or 1
		OperatorSkill selectedSkill = skillIndex == 0 ? opData.elite0Skill : opData.elite1Unlocks.unlockedSkill;
		if (selectedSkill != null)
		{
			selectedSkill.PreloadObjectPools(opData);
		}
	}
}
```

### 2. 유닛 태그 위치
- 유닛에서 관리되는 데이터들과 동일하게, 해당 유닛의 `Data`에서 태그 메서드를 `Stateless`로 구현해놨음.

> 인터페이스로 통합할 수도 있는데, 일단 헷갈리니까 따로 구현했음

```cs
// 적
foreach (var enemyPrefab in uniqueEnemyPrefabs)
{
	Enemy enemy = enemyPrefab.GetComponent<Enemy>();
	if (enemy == null || enemy.BaseData == null) continue;
	EnemyData enemyData = enemy.BaseData;

	string poolTag = enemyData.GetUnitTag();
	ObjectPoolManager.Instance.CreatePool(poolTag, enemyPrefab, 10);
}

// 배치 가능 요소들
foreach (var deployablePrefab in uniqueDeployablePrefabs)
{
	DeployableUnitEntity deployableUnit = deployablePrefab.GetComponent<DeployableUnitEntity>();
	if (deployableUnit != null)
	{
		if (deployableUnit is Operator op)
		{
			OperatorData opData = op.OperatorData;
			string opPoolTag = opData.GetUnitTag();
			ObjectPoolManager.Instance.CreatePool(opPoolTag, deployablePrefab, 1);
		}
		else
		{
			DeployableUnitData deployableData = deployableUnit.DeployableUnitData;
			string deployablePoolTag = deployableData.GetUnitTag();
			ObjectPoolManager.Instance.CreatePool(deployablePoolTag, deployablePrefab, 1);
		}
	} 
}
```

일단은 여기까지!


# 251010 - 짭명방

>[!done]
>- `UnitEntity` 피격 시 `EmissionColor`을 밝게 해서 해당 유닛이 피격되었음을 표시
>	- `Global Illumination(GI)`을 `None`으로 설정했다. `Baked`의 경우 색 변화가 인스펙터 상에는 보이는데 인게임에서는 보이지 않음.

## 피격 이펙트
- 이미 `GetHit` 이펙트가 있지만 모든 `UnitEntity`가 대미지를 받을때 잠깐 반짝이는 효과를 구현하려고 함
- 메쉬의 셰이더는 `URP Lit`을 사용하고 있기 때문에, `Emission`값만 순간적으로 높였다가 낮춰주는 방식으로 구현했다. `UnitEntity.TakeDamage`에 구현.
```cs
    protected float flashDuration = .15f;
    protected Color flashColor = new Color(.3f, .3f, .3f, 1);

    private IEnumerator PlayTakeDamageVFX()
    {
        foreach (Renderer renderer in renderers)
        {
            // 현재 렌더러의 프로퍼티 블록 상태를 가져옴 (다른 프로퍼티 유지를 위해)
            renderer.GetPropertyBlock(propBlock);
            // Emission 색상만 덮어씀
            propBlock.SetColor(EmissionColorID, flashColor);
            renderer.SetPropertyBlock(propBlock);
        }

        Debug.Log("피격당해서 색 변함");

        yield return new WaitForSeconds(1f); // 테스트용 시간

        foreach (var renderer in renderers)
        {
            // Dictionary에서 해당 렌더러의 원래 색상을 찾아옴
            Color originalColor = originalEmissionColors[renderer];
            
            renderer.GetPropertyBlock(propBlock);
            propBlock.SetColor(EmissionColorID, originalColor);
            renderer.SetPropertyBlock(propBlock);
        }

        Debug.Log("원래 색으로 돌아옴");


        _flashCoroutine = null;
    }
```
이 과정에서 **사망 애니메이션이 구현되는 방식이 별도의 머티리얼을 만들고 실행하는 방식으로 되어 있기에 이 과정도 수정이 필요**하다. 이건 나중에 하겠음. 

# 251007 - 짭명방
- 연휴라서 쉬려고 했는데 스팀 서버가 터졌다. 요즘 하는 대부분의 게임이 스팀 기반이기에 할 게 없어져서 프로젝트 진도나 조금이라도 빼야겠음.

> [!done]
> 1. 스테이지 `1-0` 다시 세팅
> 	- 스테이지별로 하나씩 목표를 정한다는 느낌으로 구현하겠음. 자세한 내용은 `WIP`에 구현.
> 2. 기타 이슈 수정
> 	- 오퍼레이터 퇴각 버튼을 클릭해도 정보가 유지되는 현상
> 		- 이 과정에서 `UnitEntity.OnDeathStarted` 이벤트 추가
> 	- 시전자 파괴 시 `SelfReturnVFXController`가 제대로 처리되지 않는 현상 수정


## 스테이지 설정
- `1-0`은 튜토리얼 개념, `1-1`은 바리케이드를 통한 진로 변경, `1-2`는 방어력이 강한 적이나 저지를 무시하고 통과하는 적 같은 개념을 보여주고 `1-3`은 앞에서 배운 것들 활용 + 보스 같이 구현해보려고 함
- 이걸 감안하면 스테이지 `1-3`은 맵 구조를 바꿔야 할 듯

### 1-0
- 기존의 5마리 적은 좀 심심한 감이 있어서 숫자를 살짝 늘림
- 원거리는 추가하면 게임이 아슬아슬해지는 느낌이 있어서 남용은 자제하는 게 좋을 듯
- 여기는 간단하게 근거리 하나, 원거리 하나 추가했다.

### 1-1
- 기존의 스테이지 `1-1`은 `1-1`의 난이도가 아니었다. 적이 우르르르 쏟아져 나옴.

## 기타 버그 수정

### 오퍼레이터 퇴각 버튼 클릭해도 정보 유지됨
- 오퍼레이터의 퇴각 버튼 클릭 시 오퍼레이터 정보가 유지되는 현상
	- `InstageInfoPanel` 관련 이슈
	- 일단 오브젝트 활성화 -> 컴포넌트 초기화 과정에서 **활성화 로직을 어디에 구현하는가?** 부터 다시 점검해봄. 기존에는 매니저에서 활성화하고 초기화는 하위 클래스에서 진행했음.
	- **자신의 활성화와 초기화는 해당 클래스 자체에서 실행**되도록 하면 좋다. 역시 캡슐화 관련 이슈임. 활성화를 별도의 클래스에 구현해놓으면, 결합도가 높아지는 이슈가 생김.
	- `InstageInfoPanel`의 요소들을 수정, 오퍼레이터의 정보를 띄우는 시점에 `Operator.OnDeathAnimationCompleted`을 구독시킴

#### 그 과정에서 발생한 이상한 버그
- 이상한 버그 발생 : 스테이지 씬에 들어가면 맵이 바로 나타나야 하는데 화면이 갑자기 흑백이 됨
	- 카메라 설정이 아예 틀어진 듯? `Main Camera`가 게임 시작 시점에 `Map`에 있는 카메라 위치 정보로 이동해야 하는데 `0, 0, 0`으로 이동한다.
	- 특이한 건 하단 `DeployableBox`를 클릭하면 원하는 방식의 구동은 된다는 거임

- **`Hide`를 구현하면서 카메라 위치 정보도 초기화하는 로직을 추가했는데, 그게 `Awake`에서 동작하는 기존 패널을 감추는 로직이랑 충돌한 것으로 보임**
```cs
// StageUIManager.cs

private void Awake()
{
	if (Instance == null)
	{
		Instance = this;
	}
	else
	{
		Destroy(gameObject);
	}

	// 비활성화 전에 참조를 넣어두기
	inStageInfoPanelScript = infoPanelObject.GetComponent<InStageInfoPanel>();

	// 최초에 꺼져야 할 패널들 비활성화
	gameOverPanelObject.SetActive(false);
	gameWinPanelObject.SetActive(false);
	inStageInfoPanelScript.Hide(); // 이 부분이 문제
	stageResultPanelObject.SetActive(false);
	confirmationReturnToLobbyPanel.gameObject.SetActive(false);

	HideItemPopup();
}


// InstageInfoPanel.cs
public void Hide(UnitEntity unit)
{
	Hide();

	if (currentDeployableInfo.deployedOperator != null)
	{
		currentDeployableInfo.deployedOperator.OnDeathAnimationCompleted -= Hide;
	}
	else if (currentDeployableInfo.deployedDeployable != null)
	{
		currentDeployableInfo.deployedOperator.OnDeathAnimationCompleted -= Hide;
	}
}

public void Hide()
{
	gameObject.SetActive(false);
	CameraManager.Instance!.AdjustForDeployableInfo(false);
}
```

- 따라서 `StageUIManager`의 저 부분은 단순히 해당 게임 오브젝트를 감추는 방식으로만 구현하면 됨.

#### 다시 원래 디버깅 진행
- 이번에는 `InstageInfoPanel`에 `Operator/Deployable`의 정보가 들어가지 않고 있음
	- 이건 쉽게 수정함 : 배치된 상황과 그렇지 않은 상황을 구분하지 않았기 때문

- 그 다음은 파괴된 시점에 UI가 정상적으로 사라지지 않는 현상임
```cs
public void Hide(UnitEntity unit)
{
	Debug.Log($"currentDeployableInfo : {currentDeployableInfo}");
	Debug.Log($"currentDeployableInfo.deployedOperator : {currentDeployableInfo.deployedOperator}");
}
```
처럼 **맨 앞에 추가해보면 메서드가 2번 호출**되는 걸 볼 수 있었다

- `UnitEntity`의 사망 이벤트 관련, 사망이 호출된 시점에 발생하는 이벤트도 하나 추가해뒀다. 즉, `UnitEntity`의 사망 이벤트가 2번 발생한다는 뜻이므로 그 부분을 봐야할 듯.

- 근데 사망은 1번 발생한다. `TakeDamage`도 현재 체력이 0 이하라면 동작하지 않는다.
- 즉 1번의 사망 이벤트에 의해 `Hide`가 2번 실행된 건데.. 정확한 원인이 어디 있는지는 모르겠음. 
- 그래서 **그냥 `currentDeployableInfo`가 `null`이 아닌 상황에서만 실행되도록 했다.** 이 오브젝트가 활성화될 때 `null`이 아니게 되고, 비활성화될 때 `null`이 되며, 그 외의 다른 경우는 없기 때문에..

### 시전자 파괴 시 SelfReturnVFXController 처리
- 원래 `UnitEntity.OnDestroyed`로 처리가 되어 있는데 이 경우 참조할 `Caster`가 없어서 반환되지 않는 듯
- 그래서 위에서 수정하면서 구현한 `UnitEntity.OnDeathStarted`로 참조할 이벤트를 바꿔봄
	- 잘 되는 듯? 근데 이거도 정확히 어떤 상황에서 발생하는지 체크하진 못했어서..



# 251004 - 짭명방
>[!done]
>1. 오퍼레이터 스킬 효과(VFX 포함)를 모두 오브젝트 풀링 기반으로 변경

## 오퍼레이터 스킬 오브젝트 풀링으로 변경
- `VFX`가 있는 스킬들의 기초는 `ActiveSkill`에 있다 : 기존 변수명이 `skillVFXPrefab`인데, 대부분 `duration`을 나타내는 이펙트로 쓰고 있기 때문에 `durationVFXPrefab`으로 이름 변경
- 보스 스킬에서 구현한 것과 같이
	- 이름은 `Stateless`로 저장, 상태를 저장하는 게 아니라 메서드 호출마다 똑같은 변수 이름을 반환
	- 인스턴스를 스킬에서 갖지 않고, 스킬`ScriptableObject`은 활성화만 하고 생명 주기는 활성화된 스킬 자체에서 관리되는 방식

- 기본적인 건 생각보다 간단하게 구현됐음 
	- 미리 해둔 게 있으니 보고 참고로 해서 가져오면 되고, 방법도 그렇게 어렵지 않음

- 오퍼레이터가 파괴되는 상황에서의 처리
	1. 버프 이펙트를 자식 오브젝트로 넣기
		- 지속 시간 스킬이므로 자식 오브젝트로 구현했는데, 이 경우 오브젝트가 파괴되면서 스킬 풀로 돌아가지 않는 이슈가 있음.
		- 그래서 `SelfReturnVFXController.OnDestroy()`에 `ReturnToPool()`을 넣어봤다. 이게 잘 되는지부터 확인.
			- 엔티티가 생성될 때 스킬 오브젝트 풀을 초기화하지만, `ObjectPoolManager.CreatePool`은 이미 풀이 있다면 더 이상 생성하지 않음
		- 잘 된다.
			- 이거는 기존에 적용했던 모든 `SelfReturnVFXController`에 대해서 적용된다. 
		- 근데 ㅋㅋㅋㅋㅋ `OnDestroy`는 **`Destroy`가 호출된 상황에서 그 직전에 실행되는 메서드인데 풀로 반환되더라도 반환 후에 파괴되니까 의미가 없다** : 오히려 풀이 오염될 위험이 있음. 파괴된 오브젝트를 참조하게 되니까.
	2. 종속성 없이 개별 오브젝트로 구현하기
		- `SetParent`를 빼면 되지만, 대신 `caster`가 파괴될 때 버프 이펙트도 비활성화되고 풀로 돌아가야 함
		- `UnitEntity`에는 `OnDestroyed`라는 이벤트가 있으므로 이를 구독함
```cs
    protected void ReturnToPool(UnitEntity entity)
    {
        ReturnToPool();
        caster.OnDestroyed -= ReturnToPool;
    }
```
> 정상적으로 동작하지 않는 문제 있었음. 해결
> `Skill`에는 **거의 대부분의 메서드에 시전자 파라미터를 넣어줘야 한다**고 생각하면 될 것 같다. 이 자체가 상태를 가지지 않게 해야 하니까.

- `duration` 이펙트들에 대해선 작업 완료
- 기타 이펙트들 작업 진행
	- 코스트 회복 스킬 : 기존에는 코스트 파티클의 움직임만 구현했음
		- 오브젝트 풀로 자체 반환하는 기능이 없어서 `SelfReturnVFXController`을 상속시키고 충돌하는 지점들 정리
	- 이외에는 없는 듯? `SixBulletSkill` 같은 경우도 기존 `buffOn`을 썼고, `ShieldSkill`도 `SelfReturnVFXController`을 상속시켜놓았음.

오늘은 여기까지. 생각보다 빠르게 했다. 

# 251002 - 짭명방

>[!done]
>1. `Shield VFX` 수정
>	- 기존 `VFX Graph` 기반이었던 것을 `Particle System`으로 고침
>2. 사라진 머티리얼들 복구하기
>	- `TileHighlight` : 아예 하이라이트되는 원리 자체를 바꿈 - **별도의 머티리얼을 사용하지 않고,** `Unit`들에 할당했던 방식(**동일한 머티리얼, 런타임 시점에 `MaterialPropertyBlock`을 이용해 색만 변경**)을 사용함
>	- `Shader_PathTrail` : 복구했으나 다시 제거. `Trail03`을 새로 만들고 `Intensity`를 크게 올리는 식으로 비슷하게 구현 가능
>	- `Shader_SkillAvailableIcon` : 얘는 따로 필요하므로 유지

## Shield VFX 수정
- **파티클 시스템 기반으로 수정**함
- 셰이더 그래프를 고쳐야 한다고 생각했는데, **거의 그대로 쓸 수 있었다.** 바리에이션을 구현한다면 별도의 머티리얼을 쓰기는 해야 할 건데 그럴 계획은 없으니까 괜찮을 듯

- **오브젝트 풀링 기반으로 수정해야 함**
	- 이거 `Operator` 스킬들 대부분이 `Instantiate` -> `Destroy` 방식으로 구현되어 있음 ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ 하
		- 일단 오늘 작업은 `Shield` 이펙트에 한해서만 오브젝트 풀링으로 구현함. 사실 오브젝트 풀링일 필요도 없긴 하다. 한 객체에 대해서 한 번에 1개만 나타나기 때문임.
	- `ShieldBuff`에서 이펙트 부분은 분리
		- 이 때 **쉴드가 다 까졌을 때 이펙트가 사라져야 하는 부분**은 따로 구현해야 할 것으로 보임. 이건 이펙트 단위에서 처리함.
		- `SelfReturnVFXController`을 상속받은 `ShieldVFXController`을 따로 만듦

```cs
using System.Collections;
using UnityEngine;

// 쉴드는 쉴드가 깨지는 이벤트가 발생하면 사라져야 하므로 그 부분만 추가
public class ShieldVFXController : SelfReturnVFXController
{
    private UnitEntity owner;

    public void Initialize(UnitEntity owner, float duration)
    {
        this.owner = owner;

        // 일단 쉴드 시스템이 Operator에만 구현되어 있기 때문에 이렇게 만듦
        if (owner is Operator op)
        {
            op.shieldSystem.OnShieldChanged += HandleShieldChanged;
        }

        base.Initialize(duration);
    }

    protected override void ReturnToPool()
    {
        if (owner is Operator op)
        {
            op.shieldSystem.OnShieldChanged -= HandleShieldChanged;
        }

        base.ReturnToPool();
    }

    protected void HandleShieldChanged(float currentShield, bool isShieldDepleted)
    {
        if (isShieldDepleted)
        {
            ReturnToPool();
        }
    }
}
```

> 상속받아서 실행시키는데 `ps.Play`가 정상적으로 동작하지 않는 듯(`base.Initialize` 부분)
> 런닝 뛰고 와서 마저 함
> --> 스크립트 만들어놓고 할당은 `SelfReturnVFXController`로 되어 있었음 ㅋㅋㅋ

![[VFX_Skill_Shield.gif]]

> - 스킬 만료시 쉴드 해제 & 흐릿한 버프 VFX도 해제된다.
> - 쉴드 수치를 1로 설정한 구현도 해봤고, 한 대 맞자마자 쉴드 이펙트는 사라지고 버프 이펙트는 남는 것까지 확인했음.

- 이제 **진짜로 소실된 머티리얼들이랑 스테이지 밸런싱, 소리 추가**하면 얼추 될 듯?

## 기타 오류 수정

### `TileHighlight` 소실
- 배치 가능한 타일을 초록색으로 강조해주는 효과. 파일들 정리하면서 이 머티리얼이 참고로 하는 셰이더가 없어진 것으로 보인다. 
- 근데 기존 구현이 게임 시작할 때 하이라이트용 머티리얼 리스트를 만들고, 하이라이트되어야 하는 상황에서 머티리얼 리스트를 스왑하는 구현이었다. 
- **이렇게 할 필요가 없지 않을까?** 머티리얼 하나로 구현하고 필요한 경우 색만 바꾸는 식으로 하면 충분하지 않을까?

- 유닛에서 했던 것처럼, 모든 머티리얼 정보를 별도로 구현해서 쓰는 것보다는 **머티리얼은 1개만 쓰고 런타임 시점에 `MaterialPropertyBlock`을 이용해서 색상 정보들을 할당하는 방식**이 가장 좋아보임
	- `Lit` 셰이더를 쓰므로 머티리얼의 `Emission`에 어두운 초록색을 할당하고, 하이라이트 상황에서 `Emission`을 켜는 방식만 구현하면 될 듯?
	- 이렇게 쓰는 경우에도 `materialInstance`에 접근해서 설정을 바꾸는 것보다 `MaterialPropertyBlock`을 쓰는 게 좋다고 한다. `materialInstance`의 설정을 바꾸는 순간 자신만의 고유한 머티리얼 인스턴스를 갖게 되며, 드로우 콜의 수가 증가하게 됨

- `Tile` 수정하는 과정에서 난리가 났음
1. **`Emission`에 `Color`을 반영해도 `Emission Color`에는 반영됐는데 실제 화면에는 나타나지 않는 현상**
	- `VFXTestScene`이나 프리팹 단위에서는 `Green`값을 `0.01 ~ 0.03`만 봐도 눈에 확 띄었는데, 인게임에서는 티가 안 난다. 
	- `Tile` 스크립트에 하이라이트 색상을 내장해뒀음. 최종값은 `0f, 0.25f, 0f`
		- 인스펙터에선 쨍한 연두색이 나오는데 인게임에선 괜찮게 나온다. 
2. **공격 범위 미리보기도 `MaterialPropertyBlock`으로 바꾸는 과정에서 색상이 다르게 나옴**
	- 메딕이면 파란색, 공격 가능 오퍼레이터는 주황색이 나와야 함
	- 일단 타일에 할당된 머티리얼이 셰이더 자체에 붙어있는 머티리얼이라서 셰이더를 통해 머티리얼을 새로 만들고 모든 타일에 다시 할당해봄 -> 안됨
	- AI가 던져준 코드를 그냥 아니까 바로 쓰면 되겠다고 해서 받아먹었는데, `AttackRange`는 별도의 셰이더를 썼다. 여기에 쓰는 색깔 프로퍼티는 이름이 `Color`이므로 `_BaseColor`가 아니라 `_Color`를 바꿔야 함.
```cs
public void ShowAttackRange(bool isMedic)
{
	Color targetColor = isMedic ? medicIndicatorColor : defaultIndicatorColor;

	attackRangeIndicator.GetPropertyBlock(indicatorPropBlock);
	indicatorPropBlock.SetColor("_Color", targetColor); // URP Lit과 달리, attackRange라는 별도의 셰이더가 있고 Color라는 프로퍼티가 있음
	attackRangeIndicator.SetPropertyBlock(indicatorPropBlock);

	attackRangeIndicator.gameObject.SetActive(true);
}
```

> 참고) Tile의 `ExecuteAlways`를 꺼뒀음. 헷갈린다.

### 다른 소실된 머티리얼들도 복구
- 휴지통에 버렸다가 '이거 여기에 쓰였네?'로 구출된 셰이더들.
	- `Shader_PathTrail`
	- `Sahder_SkillAvailableIcon`




# 251001 - 짭명방

>[!done]
>1. `OperatorUI` 사라지는 문제 수정 완료
>2. UI Sprite Atlas 생성 - 생각보다 엄청 빠르고 쉽게 했음
>3. 기존에 구현했던 VFX Graph나 텍스쳐 등등 정리

## OperatorUI 사라지는 문제 수정
- 다른 오퍼레이 사망 시 발생했던 이슈 계속) `InStageInfoPanel`은 유지되는데 다른 정보도 함께 사라지는 현상

```cs
    public void OnDeployableRemoved(DeployableUnitEntity deployable)
    {
        deployedItems.Remove(deployable);

        // 어제 수정한 부분
        StageUIManager.Instance.HideInfoPanelIfDisplaying(deployable);
        
        // 수정해야 하는 부분
        HideOperatorUIs();
        ResetHighlights();
        
        
    // 오퍼레이터 주위에 나타난 UI 제거
    private void HideOperatorUIs()
    {
        // null이어도 상관 없음

        if (currentActionUI != null)
        {
            Destroy(currentActionUI.gameObject);
            currentActionUI = null;
        }

        if (currentDeployingUI != null)
        {
            Destroy(currentDeployingUI.gameObject);
            currentDeployingUI = null;
        }

        currentUIState = UIState.None;
    }
```

- 이 부분은 `Instantiate`와 `Destroy`로 구현되어 있기 때문에 **오브젝트 풀링으로 바꾸는 것까지 해야 할 듯**

근데 하나씩 접근하자. 동시에 생각하면 머리가 복잡해져서 하나도 제대로 못할 때가 많다.

#### 다른 오퍼레이터가 죽으면 OperatorUI 유지
- `currentActionUI`, `currentDeployingUI`에 대해서도 어제 `InStageInfoPanel`에서 처리한 것처럼 "현재 어떤 오퍼레이터/배치 가능 요소"에 대한 정보를 띄우고 있는지에 대한 필드를 갖게 한다. `Deployable` 게터 프로퍼티를 노출시키겠다는 뜻.
- `DeployableManager`에서는 아래처럼 구현
```cs
// HideOperatorUIs 을 복붙, deployable
    private void HideOperatorUIsOnCondition(DeployableUnitEntity deployable)
    {
        if  (currentActionUI != null && currentActionUI.Deployable == deployable)
        {
            Destroy(currentActionUI.gameObject);
            currentActionUI = null;
        }

        if (currentDeployingUI != null && currentDeployingUI.Deployable == deployable)
        {
            Destroy(currentDeployingUI.gameObject);
            currentDeployingUI = null;
        }
    

        currentUIState = UIState.None;
    }
```

>  조건문의 순서를 생각없이 반대로 넣었다가 오류가 떴다. **당연히 `null` 체크 후에 해당 객체의 필드에 접근해야 함.** 
#### 요소들 생성/파괴 수정
- **게임 내내 한꺼번에 1개씩만 띄우는 요소**이므로 오브젝트 풀링으로 구현할 것도 없고, 단순히 활성화 / 비활성화하게끔 하면 됨
- 기존에 프리팹을 할당했던 요소들을 캔버스 내에 넣고 오브젝트로 대체하고 필드로 할당함
- 해당 클래스들의 활성화 / 비활성화나 전환에 따른 상태들 관리만 잘해주면 됨. 이건 기존에 구현되어 있었을 거임.

```cs
currentActionUI.gameObject.SetActive(true);
currentActionUI.Initialize(deployable);
```
> 활성화 후 초기화가 더 권장되는 방식이라고 함
> - 유니티 생명주기라든가 UI 레이아웃 등은 활성화 시점에 맞춰서 동작하기 때문
> - 모든 게 준비된 다음 활성화하는 게 깔끔하다고 생각했는데 **활성화 -> 초기화**가 유니티에서 더 어울리는 방식이라고 함.

- `HideOperatorUIs`들도 `null`을 할당하는 부분을 제거함

- 스크립트 상에서 오류는 없는데 나타나는 위치가 이상해짐 -> 수정
	- `ActionUI`는 `WorldCanvas`인데 `Screen Canvas` 안에 집어넣어서 발생하는 문제였음
- 오퍼레이터를 클릭했는데 ActionUI가 안 나타나는 현상 수정 중
	- 수정 완료. 얘도 `Screen Canvas` 안에 집어넣어서 발생한 문제였다.

## UI 스프라이트들 통합
- **이거는 생각보다 엄청 빠르게 진행됐다. 유니티의 `Sprite Atlas` 기능 때문.**

- `Sprite Atlas`라는 게 있다고 한다. 
- 기본적으로 꺼져 있는 듯. `Project Settings - Editor - Sprite Atlas`의 `V2, Enabled`을 사용함.
- 근데 이걸 적용해봤더니 기존 이미지들의 영역에 다른 이미지가 함께 나타나는 현상이 있음. 
	- **`Tight Packing` 옵션이 이슈**였음. 이걸 켜면 정확한 영역으로 저장하고 끄면 이미지 사각형 단위로 저장하는 듯.

- `Sprite Atlas` 자체도 미리보기에서 1개만 보일 수 있는데, 우측 상단을 보면 `Page`를 전환할 수 있다. 이걸로 지정한 경로의 이미지들이 모두 들어갔는지를 볼 수 있음

- 이 `Sprite Atlas`에 포함된 `Image`는 유니티에서 원본 텍스쳐를 `Image` 컴포넌트에 할당했더라도, 유니티에서 자체적으로 `Sprite Atlas`에서 해당 이미지를 가져온다. 
	- 원본 파일의 이름이 그대로 나타나고 있어서 헷갈릴 수 있는 부분인데 유니티에서 자체적으로 처리하고 있는 상태라고 함. 
	- 만약 정말로 그런지 알고 싶다면 `Window > Analysis > Frame Debugger`로 해당 이미지를 어디서 가져오고 있는지 체크할 수 있다. 

![[Pasted image 20251001175626.png]]
`SpriteAtlas`가 표기되어 있음
### 그러면 드는 생각
- **이펙트 텍스쳐 시트도 Sprite Atlas를 이용해서 만들 수 있나?** 오늘 해본 것처럼, UI 아이콘 Atlas를 만드는 건 겁나 빠르고 편했다. 

- 제미나이에게 물어보니 **권장하지 않는다고 함**

1. **유니티의 `Sprite Atlas`는 그리드 형태가 아님.**  
	- 파티클 시스템은 그리드라고 가정하고 진행하며, `Sprite` 모드가 있긴 한데 모든 프레임을 리스트에 수동으로 복붙해야 하는 번거로움이 있다.
2. 밉맵 생성 문제
	- 각 프레임 주위에 충분한 여백이 필요하다. 밉맵이 생길 때 이웃 프레임의 색이 섞이는 현상을 방지하기 위함이다.
	- `Sprite Atlas`는 패킹 효율을 극대화하기 위해 스프라이트들을 매우 가깝게 붙인다. 


---
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
	- [[1. Projects, Ongoing/유니티 - 작은 명방 구현하기/작업 일지/직접 작성/25년 5월|25년 5월]]
	- [[짭명방_25년 6월]]
	- [[짭명방_25년 7월]]
	- [[짭명방_25년 8월]]
	- [[짭명방_25년 9월]]
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
	- [[1. Projects, Ongoing/블로그 만들기/월별 작업 기록/25년 4월|25년 4월]]
	- [[1. Projects, Ongoing/블로그 만들기/월별 작업 기록/25년 5월|25년 5월]]
	- [[블로그_25년 6월|블로그_25년 6월]]
	- [[블로그_25년 7월]]