# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방
## 남은 작업 내용 
- **지상 오퍼레이터 1개 추가**
	- 명방에서 가드 중에 2번 공격하는 직군이 있다. 그걸로 구현하고 싶음.
	- 특징) 2저지, 2번 공격함, 스킬들은 모두 공격 시에만 회복함.
	- 1스킬을 기존 강타와 동일(?)하게 구현한다면, 2스킬은 어떻게 구현할지도 고민.
		- 2번 때리는 거니까 다르게 구현해야 할지도?

- 1-3 밸런싱, 보스 추가
- 테스트 및 수정

### 보고 참고할 내용
- [[오퍼레이터들 스탯 정리]]
- [[적 스탯 정리]]
## 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`
## 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.
	- 현재는 정예화 패널에서만 스킬 범위를 볼 수 있음. 아니면 인게임에서 실행시켜도.

# 작업 내용 : 블로그



---
# 작업 일지

# 250627 - 짭명방

## 오늘 배운 점

1. `DOTween` 애니메이션을 갖는 메서드는 일종의 코루틴으로 생각해도 될 듯
	- 그래서 애니메이션이 끝나고 어떤 걸 실행시키고 싶다면 `DOTween`의 `OnComplete()` 내에 구현해야 한다. 
	- 메서드 외부에 따로 두고 해당 메서드 밑에 구현하면 애니메이션이 끝나기 전에 동작할 수 있음.

2. **불특정 다수와의 연결을 고려하면 이벤트 기반의 설계가 제일 좋긴 하다.** 무엇보다 기능을 추가하거나 수정할 때 양쪽 객체에 일일이 구현할 필요가 없다. 해당 객체의 이벤트만 구독하는 방식으로 이으면 되니까 유지보수하기 쉽고 확장성도 좋음.

3. ★★ `트리거 기준` - 어떤 오브젝트의 자식 오브젝트에 여러 콜라이더가 있다면, **부모 오브젝트의 트리거 이벤트 감지 로직은 자식 오브젝트의 트리거 이벤트 감지 로직을 모두 감지**한다.
```
Parent (Collider 1, Rigidbody)
- Child (Collider 2)
```
일 경우, **`Parent`의 `OnTriggerEnter`는 `Collider 1, Collider 2`를 모두 감지한다.** 
따라서 이런 경우에는 `Collider 1`만 감지할 자식 오브젝트를 하나 더 둬야 한다.

- 이 때 `Rigidbody`는 부모 오브젝트에 유지하는 게 좋다.
- 즉 **`Rigidbody`는 최상위 오브젝트에 유지하고 콜라이더는 용도에 따라 자식 오브젝트들로 따로 나눠서 구현해야 한다.**
## 작업 중

### 오퍼레이터 하나 추가하겠음
- 지상 3개 / 언덕 4개라서 지상에서 하나 뚫리면 그 다음은 힘들어지는 느낌이 있다.
- 근데 이건 못했다. 


## 작업 완료

### Toast 상단 패딩 설정
- 그냥 컨테이너의 `Pos Y`만 `-20`으로 잡아주면 된다. 

	
### Enemy 사망 타이밍 조절

#### 문제 상황
- 죽는 애니메이션이 끝나기 전에 킬 스코어가 올라가서, 적이 남아있는 것으로 보이는데 게임이 끝나는 현상

#### 원인
```cs
protected override void Die()
{
	StageManager.Instance!.OnEnemyDefeated(); // 사망한 적 수 +1

	// 공격 이펙트 프리팹 제거
	if (BaseData.hitEffectPrefab != null)
	{
		ObjectPoolManager.Instance!.RemovePool("Effect_" + BaseData.entityName);
	}

	// UI 제거
	if (enemyBarUI != null)
	{
		Destroy(enemyBarUI.gameObject);
	}

	base.Die();
}
```
> `base.Die()`에는 사망 애니메이션을 실행하는 메서드가 들어있다. 그 안에서 `UnitEntity`의 파괴 이벤트와 `Destroy()` 모두 동작함.
> 그런데 위의 사망한 적수 +1 로직을 `base.Die()`이후로 뺄 수 없다. 저 안에서 파괴되기 때문에.

#### 생각
- `StageManager`에서 이벤트를 구독하게 만드는 방법을 생각했음.
	- 그런데 지금의 구현에서 `Enemy`만을 구독하게 하긴 어렵고, `UnitEntity`를 통해 구독하고 이벤트 구독 메서드 내에서 `Enemy`에 대한 형 체크를 한 번 더 해야 함
	- 가능은 한데 `Enemy`에 대해서만 동작하면 될 걸 `UnitEntity` 전체를 구독시킬 필요는 없어보인다.

#### 최종?
1. `UnitEntity.PlayDeathAnimation`에서 파괴 이벤트 발생
```cs
materialInstance.DOFade(0f, 0.3f).OnComplete(() =>
{
	OnDeathAnimationCompleted?.Invoke(this); // 사망할 것임을 알리는 이벤트
	// OnDestroyed?.Invoke(this); // 위의 이벤트로 통합
	Destroy(materialInstance); // 메모리 누수 방지
	RemoveAllCrowdControls();
	Destroy(gameObject);
});
```
> `OnDeathAnimationCompleted`는 크게 2가지 갈래로 나간다.
> 1. 기존 `OnDestroyed`에서 작동되던 것들 : 타일에서의 적 제거 로직, 오퍼레이터의 공격 범위 내의 적 / 현재 공격 대상 설정 해제 로직, `Projectile`의 공격자 / 타겟 제거 로직 등등
> 2. **이 이벤트를 구독하는 `Enemy`의 메서드가 있고, 이 메서드에서 다시 이벤트가 발생한다.** 이거 외에도 마음대로 구현하면 됨.

> 추가로, `Destroy(gameObject)`가 애니메이션의 동작 끝에 구현되어 있기 때문에 **파괴 이벤트를 발생시키겠다면 이 앞에서 발생시켜야 한다.** `DOTween`의 애니메이션 메서드는 코루틴처럼 생각하면 편하다. 즉, 코루틴 메서드 아래에 다른 메서드가 온다면 비동기 작업으로 인해 코루틴이 완료되기 전에 다른 메서드가 실행될 수 있다.


2. `Enemy`에서 해당 이벤트를 구독하는 메서드에서 다시 `Enemy`가 사라지는 이벤트를 발생시킨다.
```cs
public enum DespawnReason
{
    Null, // 디폴트
    Defeated, // 처치됨
    ReachedGoal // 목적지 도달
}

protected override void Awake() 
{
	OnDeathAnimationCompleted += HandleDeathAnimationCompleted;
}

// PlayDeathAnimation이 끝나고 호출되는 이벤트에 의해 실행되는 메서드
protected void HandleDeathAnimationCompleted(UnitEntity unitEntity)
{
	// 자기 자신의 이벤트인지 확인
	if (unitEntity == this)
	{
		OnEnemyDespawned?.Invoke(this, currentDespawnReason);
	}

	// 다른 로직도 추가 가능
}

// 

protected void OnDestroy()
{
	// OnEnemyDestroyed?.Invoke(this);
	RemoveObjectPool();

	OnDeathAnimationCompleted -= HandleDeathAnimationCompleted;
}
```
> 이렇게 구현하면 `UnitEntity`에서 1개의 이벤트를 발생시켜서 `Enemy`에 대한 여러 개의 이벤트를 발생시킬 수 있음. 지금은 1개지만.

- `Enemy`의 `Die` 메서드를 구분한다 : 체력이 다해서 죽은 경우와 목적지에 도달한 경우를 구분하며, 어느 쪽이든 사라지는 메서드로 이어진다. 대신 현재 상태만 바꾸는 방식.
```cs
    // 사라지는 로직 관리
    private void Despawn()
    {
        // 공격 이펙트 프리팹 제거
        if (BaseData.hitEffectPrefab != null)
        {
            ObjectPoolManager.Instance!.RemovePool("Effect_" + BaseData.entityName);
        }

        // UI 제거
        if (enemyBarUI != null)
        {
            Destroy(enemyBarUI.gameObject);
        }

        PlayDeathAnimation(); // 내부 이벤트 발생으로 인해 HandleDeathAnimationCompleted도 실행됨.
    }

    protected override void Die()
    {
        // 사망 이벤트 처리
        currentDespawnReason = DespawnReason.Defeated;
        Despawn();
    }

    private void ReachDestination()
    {
        currentDespawnReason = DespawnReason.ReachedGoal;
        Despawn();
    }
```
> 이 구조에서 `HandleDeathAnimationCompleted`이 실행되는 시점은 `UnitEntity`의 `OnComplete` 이후이다.

3. `StageManager`에서는 `OnEnemyDespawned`를 구독, 어떻게 디스폰되었는지에 따라 다르게 설정한다.
```cs
    private void Awake()
    {
		// ...
		
        Enemy.OnEnemyDespawned += HandleEnemyDespawned;
    }

    private void HandleEnemyDespawned(Enemy enemy, DespawnReason reason)
    {
        switch (reason)
        {
            case DespawnReason.Defeated:
                OnEnemyDefeated(enemy);
                break;
            case DespawnReason.ReachedGoal:
                OnEnemyReachDestination(enemy);
                break;
            default:
                Debug.LogError("처리되면 안되는 듯?");
                break;
        }
    }


	private void OnDestroy()
	{
		stageLoadingScreen!.OnHideComplete -= StartStageCoroutine;
		Enemy.OnEnemyDespawned -= HandleEnemyDespawned;
	}
```

- 근데 ㅋㅋㅋㅋ 이렇게 구현하고 나니까 적 체력이 0이 된 시점에도 바로 킬 카운트가 올라가지 않는 점이 거슬린다. 
- 여기에 시간을 좀 오래 쓰긴 했지만(!) `OnDeathAnimationCompleted`을 받아서 다시 이벤트를 발생시키는 대신, `Despawn`에서 `PlayDeathAnimation`을 실행하기 전에 먼저 시작함
- 대신 이전에 생각했던 **`적이 사라지기 전에 게임이 끝나는 문제`는 단순히 `StageManager`에서 관리하는 게임 종료 시점이나 딜레이를 조금 더 주는 방식으로 해결하겠음.**
```cs
// 기존엔 yield return null이었음(한 프레임)
    private IEnumerator GameWinAfterDelay()
    {
        yield return new WaitForSeconds(0.5f); // 적이 사라지는 시간이 0.3초니까 그거보다 조금 더 길게
        GameWin();
    }

    private IEnumerator GameOverAfterDelay()
    {
        yield return new WaitForSeconds(0.5f); // 적이 사라지는 시간이 0.3초니까 그거보다 조금 더 길게
        GameOver();
    }
```

- 아무튼 구조에 대해 더 생각할 수 있었으니 좋았쓰!... 인가? **지금의 이벤트 기반 구조가 더 좋아보이긴 한다.**

- 추가 발생 문제) 적이 도착지점에 도착했을 때 1개 이상 카운트됨
	-> `currentDespawnReason`이 `0`번(`Null`)이 아니라면 `Enemy.Update`를 막음.
```cs
    protected void Update()
    {
        if (StageManager.Instance!.currentState == GameState.Battle && // 전투 중
            currentDespawnReason == DespawnReason.Null // 디스폰되고 있지 않을 때
            )
```


---
### 공격 범위 밖의 적을 공격하는 현상 수정 - 콜라이더 이슈
- 테스트는 `Caster`로 한다.

#### 원인 찾기

- **모든 상황이 발생할 때도 있고 아닐 때도 있어서 정확히 추적하기는 어려운 상황이다.**

1. ~~최초 위치에서의 공격 범위와 재배치한 위치에서의 공격 범위가 다른데 최초 공격 범위에 있는 적을 공격하는 현상~~ **공격범위 밖의 적을 타격하는 현상**
	- 공격 범위 타일 선정 로직 자체에는 문제 없어보임. 재배치했을 때 목록이 깔끔하게 초기화된다.
	- 타일에서도 자신을 공격범위로 설정하고 해제하는 오퍼레이터를 추적하는 로직이 정상적으로 작동하는 것으로 보임.
	- 추측) `Enemy`의 공격 범위 콜라이더인 `AttackRangeCollider`에 얘가 반응 하는 거 아닐까? **근접 `Enemy`에는 반응이 없는데 원거리 `Enemy`에는 반응한다.** 
		- 실제로 그런 것 같은게, `AttackRangeCollider`에 충돌하기만 하면 근접 오퍼레이터든 원거리 오퍼레이터든 관계 없이 다 해당 `Enemy`를 타겟으로 선정하고 있음.
- 추적해보면 공격 범위 밖의 대상을 공격 대상으로 선정하는 로직은 아래의 과정을 거친다.
```
[OPERATOR] Caster(Clone)가 Enemy_Easy_Ranged(Clone)를 공격 가능 리스트(enemiesInRange)에 추가함
UnityEngine.Debug:Log (object)
Operator:OnEnemyEnteredAttackRange (Enemy) (at Assets/Scripts/Entities/Operators/Operator.cs:986)
Tile:EnemyEntered (Enemy) (at Assets/Scripts/Map/Tile.cs:270)
Enemy:OnTriggerEnter (UnityEngine.Collider) (at Assets/Scripts/Entities/Enemy/Enemy.cs:875)
```
> 근데 해당 타일은 오퍼레이터의 공격범위 밖임. ???

- 구체적으로 이해가 가지 않아서 AI한테 `Operator, Enemy, Tile`을 던져줬고 답을 받았다. 
- 요점은 **`OnTriggerEnter`는 모든 자식 오브젝트의 콜라이더에서 발생하는 이벤트를 모두 수신한다는 것.** 즉, `Enemy` 본체에 있는 `BoxCollider` 외에도 공격범위 콜라이더에 있는 `SphereCollider`까지 `Enemy.OnTriggerEnter`에서 처리한다는 것이다.
- 따라서 **`원거리 Enemy`가 밟고 있는 타일로 인식되는 지점은 본체에 있는 `BoxCollider` 외에도 공격 범위를 반지름으로 갖는 구에 충돌하는 모든 타일이다.**
- 이걸 해결하는 가장 심플한 방법 : **2개 이상의 콜라이더를 한 오브젝트 아래에 둔다면 각 콜라이더는 자식 오브젝트로 두라는 것**이다. 

- 지금까지의 구현에서 `UnitEntity` 자체에 본체 충돌 콜라이더를 뒀고, `Enemy`의 경우에만 자식 콜라이더를 하나 둔 상태인데.. **아예 `UnitEntity`부터 고치는 게 나중의 확장성을 고려할 때도 더 맞는 설계가 될 것 같다.** 확장을 얼마나 할지는 모르겠지만.

#### 수정
- **모든 `UnitEntity`의 자식 오브젝트로 `Body`를 두고, 거기서 콜라이더를 관리한다.**
	- `Model`에 구현하는 방법도 있지만 `Body`라는 개별 오브젝트를 관리하는 방식이 더 표준적인 패턴이라고 함
	- **`Rigidbody`는 부모 오브젝트에 유지한다.**
- 자식 오브젝트에는 `BodyColliderController.cs`라는 스크립트와 부모 오브젝트에 있던 `BoxCollider`를 기존 사양에 맞게 붙여둔다. 이 트리거의 `OnTriggerEnter`, `OnTriggerExit`이 동작하면 `UnitEntity`의 `OnBodyTriggerEnter, OnBodyTriggerExit`이 동작하도록 한다.
- 기존 부모 오브젝트들에 있던 `OnTriggerEnter, OnTriggerExit`을 이름만 `OnBodyTriggerEnter, OnBodyTriggerExit`으로 바꾼다. 
- 이 과정에서 컴포넌트 지우고 옮기고 하는 건 수작업이다. 어쩔 수 없음.

#### 결과
- **타일 밖에 있는 원거리 적을 때리는 현상은 사라졌다.**
- 대신 저지가 안되네? -> 
	1. `Enemy` 대신 `BodyColliderController`를 감지시키고, 그걸 갖고 있는 부모 오브젝트가 `Enemy`인지 확인하도록 로직 수정.
	2. 콜라이더가 본체랑 떨어져서 나타나는 이슈가 있었다. 정확히는 프리팹 자체의 위치 설정 문제였고, 수정 완료.
- 이번엔 `DeployableUnitEntity` 클릭이 안되네?
	- `ClickDetectionSystem.HandleObjectClick`을 보면
```cs
    private void HandleObjectClick(RaycastHit hit)
    {
        // 레이캐스트를 맞은 콜라이더의 부모 오브젝트에서 DeployableUnitEntity를 찾도록 수정
        // DeployableUnitEntity? clickable = hit.collider.GetComponent<DeployableUnitEntity>();
        DeployableUnitEntity? clickable = hit.collider.GetComponentInParent<DeployableUnitEntity>();
```
> 원래는 콜라이더의 위치에 `DeployableUnitEntity`가 있었으나 지금은 그것보다 자식 오브젝트가 레이캐스트를 맞는다. 따라서 부모 오브젝트에서 `DeployableUnitEntity`를 찾아야 함.

#### 이외의 복기
```cs
    private bool HandleUIClick(List<RaycastResult> uiResults)
    {
        foreach (var result in uiResults)
        {
			// ... 생략
            // 2. OperatorUI 관련 요소 클릭 처리 - Deployable.OnClick이 동작하도록 수정
            
            DeployableUnitEntity? associatedDeployable = GetAssociatedDeployableUnitEntity(result.gameObject);
            if (associatedDeployable != null )
            {
                associatedDeployable.OnClick();
                return true;
            }
        }

        return false;
    }
```
> - 오랜만에 코드를 보니까 `HandleUIClick`에 왜 `DeployableUnitEntity.OnClick`을 구현했는가? 가 헷갈린다.
> - 왜냐면 `DeployableUnitEntity`가 갖는 `OperatorUI`가 있음 : `HealthBar`라든가, 스킬 활성화 가능 표시 버튼이라든가. 게임 상의 저 요소를 그대로 클릭하면 `DeployableUnitEntity`를 클릭한 것처럼 동작하는 게 아니라 아무 반응도 없음. 왜냐하면 유니티의 이벤트 시스템이 그 UI를 클릭한 것으로 처리하기 때문이다. 
> - 따라서 여기서는 그러한 **UI를 클릭하더라도 해당 `DeployableUnitEntity`를 클릭한 것과 같은 반응을 시키겠다~ 라는 의미로 넣은 코드**다. 없으면 오히려 불편하게 될 거임.



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
