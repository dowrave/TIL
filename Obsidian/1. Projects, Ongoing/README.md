# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 남은 작업
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




---
# 작업 일지


## 짭명방 예정

### 작업 중
- 보스 구현 시작

### 발견한 이슈

### 예정
- `1-3` 스테이지 구현 완료
	- 보스 구현

# 250724 ~  - 유니티 VFX 공부
- 유데미에서 예전에 VFX를 넣을 때 유튜브에서 신세를 졌던 센세의 강의를 듣기로 했다.
	- Gabriel Aguiar 센세의 강의 Visual Effects for Games in Unity - Beginner To Intermediate로 공부해보는 중.
	- [[1_유니티VFX_이펙트의 종류]]
	- [[2_유니티VFX_VFX의 기초, 원칙]]
	- [[3_유니티VFX_제작 전]]
	- [[4_유니티VFX_실습 사전 세팅]]
	- [[5_유니티VFX_스파크만들기]]
	- [[6_유니티VFX_범위효과]]
	- [[7_유니티VFX_투사체&주문]]
	- [[컬러 팔레트 참고용]]
	- 깃허브에서는 `Projects/Ongoing`으로 접근할 수 있다. 언제 `Archive`로 옮겨질지는 모르겠지만.

# 250723 - 짭명방
## 버그 수정

>[!done]
>- 장판 스킬들 범위 이펙트 이슈
>- `MeteorSkill` 1번째 파티클이 너무 빨리 떨어짐
>- 기절 VFX 중첩 문제 + Buff 관련 리팩토링
>- 공용 `Buff VFX` 들은 오브젝트 풀링으로 구현
>	- `VFX`가 달려 있는 부모 오브젝트 파괴 시에 풀로 반환되지 않고 함께 제거되는 문제 수정

### 장판 스킬들 범위 이펙트 이슈
- **풀에 등록할 때의 태그와 사용할 때의 태그를 다른 걸 사용해서 발생했던 문제. 스킬 내에서 전역 태그값을 관리하는 방식으로 변경했다.**


- 스킬이 켜졌을 때 범위 이펙트 관련 프리팹들이 동작하지 않는 것으로 보임. 
	- 스킬을 1번 켜봤을 때 이펙트의 위치 좌표를 보면 모두 `(0, 0)`인 점
	- 그렇다고 이펙트가 엉뚱한 위치에서 보이는 것도 아니었다.
- **오브젝트 풀에 넣을 때의 태그와 풀에서 만들어낼 때의 태그가 다른 듯?** 
	- 이펙트 태그로 스폰하려고 하는데 만들었을 때의 태그와 다르니까 스폰 자체가 불발되는 것으로 보인다. 

- 스킬의 오브젝트 풀 등록에 사용되는 아래의 로직이 반복되고 있다.
```cs
	MELEE_ATTACK_OVERRIDE_TAG = GetVFXPoolTag(caster, meleeAttackEffectOverride);
	if (!ObjectPoolManager.Instance.IsPoolExist(MELEE_ATTACK_OVERRIDE_TAG))
	{
		ObjectPoolManager.Instance.CreatePool(MELEE_ATTACK_OVERRIDE_TAG, meleeAttackEffectOverride);
	}
```
이거를 조금 더 축약해서, `GetVFXPoolTag -> CreatePool`로 이어지는 과정을 메서드화함. `CreatePool` 내에 조건 체크가 있기 때문에` IsPoolExist`도 사실 필요하진 않다.
```cs
/// <summary>
/// 스킬과 관련된 오브젝트 풀을 등록한다. 태그를 반환한다.
/// </summary>
protected virtual string RegisterPool(UnitEntity caster, GameObject prefab, int initialSize = 5)
{
	if (prefab == null) return null;

	string poolTag = GetVFXPoolTag(caster, prefab);
	ObjectPoolManager.Instance.CreatePool(poolTag, prefab, initialSize);
	return poolTag;
}
```

### `MeteorSkill` 1번째 파티클이 너무 빨리 떨어지는 문제
- 1번째 파티클의 시작 지점 높이만 바꿔줬다. `1f`는 생각보다 낮은 듯.

### 기절 이펙트 중첩되는 문제 + Buff 관련 리팩토링

- `Buff` 중에서 기절이 다시 들어갈 때는 이전 버프를 파괴하고 새로운 버프를 등록하는 방식으로 진행했다. 이 파괴와 생성 로직 내에 이펙트를 없애고 다시 만드는 과정도 포함되어 있음.
- 일단 기절 버프가 다시 들어갈 때 "파티클 이펙트를 중지하고 인스턴스를 파괴한다"는 **로직 자체는 디버깅 로그에서 정상적으로 출력되고 있음. 근데 파티클 이펙트가 중지되지 않음.**
1. `ParticleSystem.Clear()`, `VFXGraph.Reinit()`을 넣어봄(새로운 파티클 생성 중지) : 그래도 안 됨.

> 이거 도대체 뭐가 문제임?

- 일단 요인 하나
	- `UnitEntity.AddBuff`를 하면 그 안에 `Buff.OnApply()`가 기본적으로 함께 들어가 있는데, `OnApply`를 따로 구현해놓은 부분이 있음. 이거 제거.
	- **이거 없애니까 해결됐다.**

#### 왜 이런 문제가 발생했는지 정리하고 넘어감
- **AI를 많이 쓰니까 구조 파악이 무뎌지는 느낌이 들어서 직접 정리**해본다.
- 현재의 버프 구조를 보면 아래와 같다.

1. `AddBuff` 
	- `StunBuff`가 있다면 `RemoveBuff`로 해당 버프를 제거함
	- `activeBuffs`에 해당 `buff`를 추가
	- `buff.OnApply()` 실행
		- `VFX` 실행 로직도 여기 있음
	- `OnBuffChanged?.Invoke` 실행

2. `StunBuff`는 `duration`으로 초기화한다.

3. `Buff`에는 공통적으로 `OnUpdate` 로직이 있다. `duration`이 지나면 스스로가 있는 `UnitEntity`에 `RemoveBuff`를 실행시킨다.

4. `RemoveBuff`
	- `buff.OnRemove()` 실행
		- VFX 제거 로직도 여기 있음
	- `OnBuffChanged?.Invoke` 실행

`MeteorSkillController`에서 발생했던 문제는 **1번에서 `AddBuff`를 실행한 다음 별도로 `Buff.OnApply()`를 실행했기 때문이었다.**

```cs
public override void OnApply(UnitEntity owner, UnitEntity caster)
{
	this.owner = owner;
	this.caster = caster;
	elapsedTime = 0f;

	// 적용될 때 VFX 재생
	PlayVFX();
        
	owner.AddRestriction(ActionRestriction.Stunned);
}
```
1. `AddBuff`로 `OnApply`가 먼저 실행됨 -> VFX가 실행됨
2. 1번의 상태에서 다시 `OnApply`가 실행됨 -> 시간이 초기화되고, VFX가 한 개 더 실행됨
3. 시간이 지나고 `RemoveBuff`를 실행시키면 1번에서 실행된 VFX의 참조는 실종된다. 2번에서 다시 실행시켰던 VFX만 제거되게 된다.
4. `StunBuff`가 다시 들어오더라도 1개의 이펙트가 계속 실행되는 상태에서, 추가로 다시 이펙트를 실행하게 됨. 

> `VFX`의 경우 그 자체에서 파라미터를 노출시키는 것보다는 **지금처럼 계속 실행되게 하고 스크립트로 켜고 끄는 걸 조절하는 방법이 더 좋다고 함**
---
- 여기부턴 AI와 함께 정리. 최종적으로 `Buff`에서 `Owner`가 예기치 못한 방식으로 파괴될 때까지를 반영한 `OnDestroyed` 이벤트 구독 처리까지 포함해서 구현했다.
	- 이것까지 구현한 이유는 기절 이펙트 오브젝트 풀을 10개를 만들었는데, 중간에 무슨 문제인지 1개가 사라져서 9개의 오브젝트만 보이는 현상이 있었기 때문이다. 즉, 예상치 못한 상황에서 오브젝트가 파괴되었던 것으로 보임.
```
정상 종료 (시간 초과): OnUpdate -> owner.RemoveBuff() -> OnRemove()
정상 종료 (유닛 사망): owner.OnDeathAnimationCompleted -> HandleOwnerTermination -> OnRemove()
비정상 종료 (강제 파괴, 씬 전환): owner.OnDestroyed -> HandleOwnerTermination -> OnRemove()
위의 모든 경로가 결국 OnRemove라는 단일 정리 메서드로 모입니다.
```
> 대충 이런 것들이 반영되었다~

> 이건 잡설이지만 요즘 제미나이가 사용자를 칭찬하는 방식이 부담스럽다. 옛날에 클로드가 비슷하게 부담스러웠던 기억이 있음.




### 공용 이펙트 오브젝트 풀링으로 구현
- `StunBuff` 같은 경우는 여러 상황에서 사용되므로 **스테이지가 시작되는 시점에 오브젝트 풀을 만들어놓으면 좋을 것 같음.** 
	- 보통 오퍼레이터가 사용하는 오브젝트 풀은 오퍼레이터가 배치되는 시점에 생성된다. 이것과는 차이를 두려고 함.

- 작업 방향
1. `BuffVFXDatabase` : 프리팹 정보를 딕셔너리에 저장하고, 오브젝트 풀도 생성한다.
2. `BuffVFXManager` : `ObjectPoolManager`에서 이펙트 오브젝트를 얻어 반환하는 `ReleaseBuffVFXObject`를 추가한다
3. `Buff`는 `CreateBuffVFXObject`를 호출해서 이펙트를 가져오고 버프라 끝날 때 `ReleaseBuffVFXObject`를 호출해 풀에 반환한다.

- 추가
	- 이펙트 위치 문제. 지금은 `Stun`만 쓰니까 오브젝트의 머리 위에 이펙트를 나타나게 한다. 
	- 만약 슬로우 같은 이펙트를 추가로 구현해야 한다면 오브젝트 아랫쪽에 이펙트가 나타나게 해야 함. 
	- 이를 위해 `BuffVFXDatabase.BuffVFXMapping`을 아래처럼 구현한다.
```cs
    [System.Serializable]
    public struct BuffVFXMapping
    {
        [Tooltip("버프 클래스의 정확한 이름 (네임스페이스 포함)")]
        public string exactBuffClassName;
        public GameObject vfxPrefab;
        [Tooltip("풀의 초기 생성 개수")]
        public int initialPoolSize;
        [Tooltip("타겟을 기준으로 한 VFX의 로컬 위치 오프셋")]
        public Vector3 vfxOffset; 
    }
```
> 여기서 "클래스의 정확한 이름"을 입력해야 하는 이유는
> 1. 유니티의 인스펙터에서 `System.Type`을 필드로 노출해서 드래그&드롭하는 기능을 지원하지 않음
> 2. `string`은 인스펙터에서 매우 쉽게 입력하고 수정할 수 있으며, 직렬화도 용이함. 
> 3. `DB`에서 특정 버프 클래스를 직접 참조하지 않고 문자열을 통해 런타임에 연결되기 때문에 클래스 파일의 위치 등에 의존하지 않는 느슨한 결합

...등이 있다.


#### 부모 오브젝트 파괴 시 풀로 반환되지 않고 함께 파괴되는 문제 수정
- `Owner(UnitEntity)`가 파괴되는 시점에 이펙트들을 먼저 풀로 반환(및 `owner`의 자식에서 탈출)하기 위해 `OnDeathAnimationCompleted` 이벤트를 구독시켜서 탈출함

- 여기서 알아둘 게 2가지 있는데
1. **람다 함수로 구현하기 vs 래퍼 함수로 구현하기**
- `OnDeathAnimationComplted`는 `UnitEntity`를 받아야 하는 액션이다. 반면 `RemoveVFX()`는 파라미터를 받지 않음.
- 이를 구현하는 방법은 2가지다.
1) 람다함수로 구현
```cs
owner.OnDeathAnimationCompleted += (deadOwner) => RemoveVFX();
```
2) 별도의 래퍼 함수 구현
```cs
	owner.OnDeathAnimationCompleted += RemoveVFX;
}

protected void RemoveVFX(UnitEntity owner)
{
	RemoveVFX();
}
```

그런데 여기서 `Buff`의 오브젝트 풀들은 재사용될 수 있는 요소이므로, 구독 해제가 반드시 들어가야 한다.  이 때 차이가 생기는데
1) 람다 함수로 구현
- 여기서는 별도의 함수를 하나 만들어둔 상태이기 때문에, 저 람다함수를 별도의 델리게이트 상태로 갖고 있어야 한다.
```cs
private Action<UnitEntity> onOwnerDeathAction;

// 생성 로직
// 람다를 생성하여 변수에 저장합니다.
onOwnerDeathAction = (deadOwner) => RemoveVFX();

// 저장된 변수를 사용하여 이벤트를 구독합니다.
owner.OnDeathAnimationCompleted += onOwnerDeathAction;

// 해제 로직
owner.OnDeathAnimationCompleted -= onOwnerDeathAction;
onOwnerDeathAction = null; // 참조 정리
```

2) 래퍼 함수로 구현
- 훨씬 간단하다. 이미 따로 정의를 해뒀기 때문에 그것만 그대로 이용하면 됨.
```cs
owner.OnDeathAnimationCompleted -= RemoveVFX;
```

> 참고) 이런 건 안됨
```cs
owner.OnDeathAnimationCompleted -= (deadOwner) => RemoveVFX();
```
> 왜냐하면 등록할 때의 람다 함수와 해제할 때의 람다 함수가 서로 다른 인스턴스이기 때문이다. 람다 함수라는 게 즉석에서 바로 만들어내는 방식인 걸 생각해보면 됨. 
> 똑같이 `(deadOwner) => RemoveVFX();`라는 로직을 가져도 등록할 때와 해제할 때의 2개의 함수는 다르다는 것이다.


2. **구독 해제 메서드 내에 자기 자신을 해제시키는 메서드가 들어간다면, 그 로직은 맨 앞에 놓기**
```cs
    protected virtual void RemoveVFX()
    {
        owner.OnDeathAnimationCompleted -= RemoveVFX;
```
> - 자기 자신의 재실행을 막기 위해, 구독 해제 메서드는 맨 앞에 구현한다. 



# 250722 - 짭명방

## 스킬에서 사용하는 오브젝트 풀 관련
>[!info]
> - 스킬과 버프에 **공격 이펙트가 변경되는 경우에 대비한 필드** 작성
> - 스킬의 오브젝트 풀을 생성하는 메서드는 스킬에 적어놓되 실행은 Operator가 배치되는 시점에 한다.
> - Operator가 파괴되어도 오브젝트 풀이 파괴되지는 않는다. 재배치 등의 상황이 있기 때문에 굳이 지웠다가 다시 생성할 필요는 없다.
> - 오브젝트 풀들은 스테이지가 종료되는 시점에 일괄적으로 파괴된다.
- 일단 마법 공격으로 나가는 SlashEffect는 색깔이 보라색으로 바뀌어서 나가도록 수정함(시각화)

### 구현
>[!done]
>1. `Skill`에서 해당 스킬에 사용되는 이펙트 프리팹들을 필드로 정의한다.
>	- 예시
>		- 변경된 공격 이펙트`SlashEffect`
>		- 장판 이펙트 
>		- 실제 필드 효과 컨트롤러
>		- 피격 이펙트 등등
>	- 이 때 오브젝트 풀을 생성하는 메서드는 `BaseSkill`에서 정의한다. `BaseSkill`에서는 변경되는 공격 이펙트 부분만 구현하며, 나머지는 각각의 스킬에서 오버라이드해서 다시 정의한다.
>2. `Buff`에서는 스킬로부터 오버라이드된 정보를 받는다. 
>3. `Operator`가 배치되는 시점에 오브젝트 풀을 초기화한다. 이 때 현재 갖고 있는 스킬에 대한 초기화도 진행된다. 공격 이펙트가 변경되는 건 여기서 구현하는데, 이 부분만 스크립트 첨부함.

```cs
	// 기본적으로 Operator가 갖고 있는 공격 이펙트를 사용한다.
	GameObject effectPrefab = OperatorData.meleeAttackEffectPrefab;
	string effectTag = meleeAttackEffectTag;

	// [버프 이펙트 적용] 물리 공격 이펙트가 바뀌어야 한다면 바뀐 걸 적용함
	// 공격 이펙트가 바뀌는 버프는 1개만 있다고 가정한다.
	var vfxBuff = activeBuffs.FirstOrDefault(b => b.MeleeAttackEffectOverride);
	if (vfxBuff != null)
	{
		effectPrefab = vfxBuff.MeleeAttackEffectOverride;
		effectTag = vfxBuff.SourceSkill.GetVFXPoolTag(this, effectPrefab);
	}
```





# 250721 - 짭명방

## 보스 구현 아이디어만
- 기본적인 동작은 `Enemy`의 그것과 동일하다. 
- 대신 일반적인 적과 달리 2가지 스킬을 구현할 예정

- 여기선 아이디어만 쏟아내보자. 
- **모든 스킬은 자신의 쿨타임이 돌 때마다 자동으로 나간다.** 

1. 저지 당할 때 (아이디어)
	- **평타에 기절 섞여서 나감**
		- 여기서 "기절"의 효과는 오퍼레이터의 모든 행동을 막는 역할이다. 
		- 만약 저지 중인 적이 있다면 모든 저지 효과가 풀려야 함.
	- 자신을 저지하고 있는 적을 통과함
	- 단순히 대미지만 증가한 평타
2. 저지 당하지 않을 때 : **채널링 후에 13칸 범위 마법 대미지**


## 기타 밸런스 조정
- 보상 설정
	- `1-1` 클리어 시 1정예화에 필요한 재료 2개 -> 3개씩 지급


# 250719 - 짭명방

## 작업 완료
>[!done]
>`Dualblade` 스킬 아이콘 2개 만들기

- 둘 다 기존 아이콘을 기반으로 변화만 살짝 줬다.

![[Skill_MagicalSmash_128.png]]

![[Skill_MagicalAttackupWithStun_128.png]]

# 250718 - 짭명방
- 오늘 스킬 리팩토링은 끝낸다...
- 근데 수정해도 이해하기 좀 빡센 느낌은 있음.
## 작업 완료

>[!done]
> 스킬 리팩토링
> - `AmmoBasedSkill`
> - `DoubleShotSkill`
> - `AreaEffectSkill` 기반 : `Meteor, ArcaneField, AreaHasteHeal`
> - `SlashSkill`
> - `ShieldSkill`

### AmmoBasedSkill
- 탄환 기반의 스탯 강화 스킬이다.
	- 스탯 강화 버프는 이미 구현되어 있다.
	- 여기에 추가로 **"공격 횟수를 세는" 버프만 추가**하면 된다.
		- 이 버프는 자신이 동작할 때 자신과 연결된 다른 버프들을 해제시키기 위해, 동시에 실행되는 다른 버프들도 갖는다.
		- 이건 스킬에서 설정해주면 됨.
- 스킬은 이 두 개의 버프를 조합하는 방식.

1. **기존 `AmmoBasedActiveSkill` 구현**
```cs
using UnityEngine;
using System.Collections;

namespace Skills.Base
{
    // 탄환 수가 제한되어 있고 지속시간이 무한인 스킬의 추상 클래스
    public abstract class AmmoBasedActiveSkill : ActiveSkill
    {
        [Header("Ammo Settings")]
        [SerializeField] protected int maxAmmo;

        [Header("Attack Modifications")]
        [SerializeField] protected float attackSpeedModifier = 1f;
        [SerializeField] protected float attackDamageModifier = 1f;

        protected int currentAmmo;

        protected float originalAttackSpeed;
        protected float originalAttackDamage;

        protected OperatorUI? operatorUI;

        public int CurrentAmmo => currentAmmo;
        public int MaxAmmo => maxAmmo;

        public override void Activate(Operator op)
        {
            caster = op;
            operatorUI = op.OperatorUI;

            if (!op.IsDeployed || !op.CanUseSkill()) return;

            PlaySkillVFX(op);
            PlayAdditionalVFX(op);

            currentAmmo = maxAmmo;

            originalAttackSpeed = op.AttackSpeed;
            originalAttackDamage = op.AttackPower;

            // 스탯 변경 적용 
            op.AttackPower *= attackDamageModifier;
            op.AttackSpeed /= attackSpeedModifier;

            // 스킬 사용 직후에는 공격 속도 / 모션 초기화
            op.SetAttackDuration(0f);
            op.SetAttackCooldown(0f);

            // SPbar UI 업데이트
            operatorUI?.SwitchSPBarToAmmoMode(maxAmmo, currentAmmo);

            op.StartCoroutine(Co_HandleSkillDuration(op)); // 코루틴은 MonoBehaviour에서만 사용 가능함
        }

        public override void OnAfterAttack(Operator op)
        {
            // 스킬이 켜졌을 때에만 동작
            if (!op.IsSkillOn) return; 

            // 공격 후 탄환 소모
            ConsumeAmmo(op);

            // 탄환이 소진되면 스킬 종료
            if (CurrentAmmo <= 0)
            {
                TerminateSkill(op);
            }
            
        }

        protected virtual void ConsumeAmmo(Operator op)
        {
            currentAmmo = Mathf.Max(0, currentAmmo - 1);

            // 탄환 소모 후 UI 업데이트
            UpdateAmmoUI(op);
        }

        protected override void OnSkillEnd(Operator op)
        {
            // 스킬 종료 시 원래 스탯으로 복원
            op.AttackPower = originalAttackDamage;
            op.AttackSpeed = originalAttackSpeed;

            base.OnSkillEnd(op);
        }

        protected virtual void UpdateAmmoUI(Operator op)
        {
            op.OperatorUI?.UpdateAmmoDisplay(currentAmmo);
        }

        public virtual void TerminateSkill(Operator op)
        {
            OnSkillEnd(op);
            operatorUI?.SwitchSPBarToNormalMode();
        }

        public override IEnumerator Co_HandleSkillDuration(Operator op)
        {
            OnSkillStart(op);
            // op.StartSkillDurationDisplay(duration?ㅉ);
            PlaySkillEffect(op);

            while (op.IsSkillOn)
            {
                // 오퍼레이터 파괴 시 동작을 멈춤
                if (op == null) yield break; 

                // 탄환이 떨어지면 스킬 종료
                if (currentAmmo <= 0)
                {
                    TerminateSkill(op); // 종료 로직은 이 안에 구현됨
                    yield break;
                }

                // 기본) 무한 지속
                yield return null;
            }
        }
    }
}
```

2. **`AttackCounterBuff` 구현** : 연결된 버프들을 관리한다. 공격 횟수가 최대 횟수에 도달하면 연결된 버프를 해제하고 UI를 원래 상태로 되돌려놓음.
```cs
using System.Collections.Generic;
using UnityEngine;

public class AttackCounterBuff : Buff
{
    private int maxAttacks;
    private int currentAttacks;

    // 이 버프가 해제될 때 함께 해제될 버프들
    private List<Buff> linkedBuffs = new List<Buff>();

    private OperatorUI operatorUI;

    public AttackCounterBuff(int maxAttacks)
    {
        this.buffName = "Attack Counter";
        this.duration = float.PositiveInfinity;
        this.maxAttacks = maxAttacks;
    }

    public void LinkBuff(Buff buffToLink)
    {
        if (buffToLink != null)
        {
            linkedBuffs.Add(buffToLink);
        }
    }

    public override void OnApply(UnitEntity owner, UnitEntity caster)
    {
        base.OnApply(owner, caster);
        currentAttacks = maxAttacks; // 최대 횟수를 정해놓고 1씩 줄이느

        // UI를 탄환 모드로 전환
        if (owner is Operator op)
        {
            operatorUI = op.OperatorUI;
            operatorUI?.SwitchSPBarToAmmoMode(maxAttacks, currentAttacks);
        }
    }

    public override void OnAfterAttack(UnitEntity owner, UnitEntity target)
    {
        base.OnAfterAttack(owner, target);

        currentAttacks = Mathf.Max(0, currentAttacks - 1);
        operatorUI?.UpdateAmmoDisplay(currentAttacks);
        if (currentAttacks <= 0)
        {
            // 연결된 모든 버프 제거
            foreach (var buff in linkedBuffs)
            {
                owner.RemoveBuff(buff);
            }

            owner.RemoveBuff(this);
        }
    }

    public override void OnRemove()
    {
        operatorUI?.SwitchSPBarToNormalMode();
        
        base.OnRemove();
    }
}
```

3. **`AmmoBasedSkill` 수정**
- 버프가 종료되는 시점에 스킬이 종료돼야 함. 
- 지속시간이 있는 스킬은 자체적으로 `duration`을 갖고 스킬 자체의 코루틴이 돌기 때문에 `OnSkillEnd`가 정상적으로 동작하지만, 지금처럼 지속시간이 무한인 스킬인 경우에는 버프가 종료되는 시점에 `OnSkillEnd`가 실행돼야 한다.
- AI한테 관련 코드를 다 던져줘도 무시해서 한 번 더 물어봄. 역시 무지성 복붙은 좋지 않다.
- 결론은 **`Buff`에서 스킬 종료 시에 이벤트를 발생시키고, 스킬에서 해당 이벤트를 구독**시키는 것.

```cs
// Buff에 아래 델리게이트 추가
    public System.Action OnRemovedCallback; // 버프 종료 시에 호출되는 콜백 함수

// Skill에서 이벤트 구독
	attackCounterBuff.OnRemovedCallback += () => OnSkillEnd(op);
```
> **`OnSkillEnd(op)`가 파라미터를 받고, `OnRemovedCallBack`은 그렇지 않기 때문에 직접 받을 수 없다. 이럴 때 사용할 수 있는 게 람다식.**

- **그리고 이벤트 구독 해제 코드를 별도로 짤 필요는 없다.** 
	1. `Buff`가 해제될 때 `Operator`에게 자신을 참조에서 해제하라는 명령어를 보냄(이게 중요!)
	2. 참조에서 해제됨. 이 상태에서 이 `Buff`는 씬에 있는 어떠한 오브젝트와도 연결된 상태가 아니기 때문에 가비지 콜렉터에 의해 제거될 수 있는 요소가 된다.

- 만약 미사일과 미사일의 피격이벤트를 메서드로 구독하는 GameManager가 있다고 하면, 미사일이 사라질 때 구독 해제를 하지 않았다고 하자. 이 때 미사일 - 게임 매니저를 연결하는 요소가 여전히 남아있기 때문에 이 미사일은 가비지 콜렉터에 의해 해제되지 않는다.

- **UI에 연결하는 것도 이벤트 기반의 로직**으로 수정
	- 1. 버프가 추가/제거될 때마다 `UnitEntity`에서 이벤트를 발생시킴
	- 2. `UI`는 원하는 버프가 오퍼레이터에 추가된 걸 발견하면 `UI`의 초기화`OnApply` 및 업데이트 메서드를 등록시킴


### DoubleShotSkill
- 기존 구현에서 핵심은 `ModifiesAttackAction`이라는 게 있었고, **해당 필드가 활성화**됐을 때 공격 매커니즘이 `Attack`이 아닌 **스킬에 있는 방식을 따르는 것**이었음.
- 이걸 `Buff`로 옮기면 된다. 


### ArcaneFieldSkill
1. **`CannotAttackBuff` 만들기**
- 말 그대로 버프가 켜진 동안에 기본 공격을 하지 않음
- 단 공격 쿨타임은 계속 돌아감 - 아무것도 하지 않는데도 계속 호출될 수 있기 때문에
```cs
public class CannotAttackBuff : Buff
{
    public CannotAttackBuff(float duration)
    {
        this.buffName = "Cannot Attack";
        this.duration = duration;
    }

    public override bool ModifiesAttackAction => true;

    public override void PerformChangedAttackAction(UnitEntity owner)
    {
        // 공격 쿨타임만 돌리고 실질적으로 아무런 기능을 하지 않음을 표시함
        // Buff와 똑같기 때문에 굳이 넣어도 되지 않지만, 가독성을 위해 넣음
        // 쿨타임을 돌리는 이유는 계속된 호출을 방지하기 위함
        base.PerformChangedAttackAction(owner);
    }
}
```

2. `ArcaneFieldSkill` 수정하기
- `CannotAttackBuff`를 자신에게 걸고, `ArcaneFieldController`을 소환하는 기능
- 참고) 이 과정에서 `ActiveSkill -> AreaEffectSkill -> ...`으로 범위를 구현하는 스킬들이었던 `MeteorSkill`, `ArcaneFieldSkill`, `AreaHasteHealSkill` 등을 **단순하게 `ActiveSkill`에서 직접 상속받는 방식으로 변경한다.**
	- 지금 수정되는 방식이 Fire and Forget이므로 Skill은 필요한 객체들을 발사하기만 하고, 생명주기는 각 객체가 스스로 책임지도록 하면 된다. 지금의 `AreaEffectSkill`에서 직접 관리할 필요는 없음.
- `Controller`까지 한번에 싹 수정함. 
- 생각보다 잘 된다.

- 이 부분은 나도 작업하면서 머리가 좀 복잡해지는 내용이어서 정리해봄. 

> [!note]
> **1. 수정 전** 
> - `AreaEffectSkill`이라는 중앙 관리자가 있었다. 범위 스킬들을 묶어서 관리하기 위해 구현했던 건데, 서로 다른 방식으로 동작하는 스킬들(즉발, 지속시간 등)을 범위 스킬이라는 이유만으로 묶어서 관리하려고 했다.
> - 씬에 생성된 모든 효과 객체들을 일일이 추적해서 시전자가 죽으며 직접 청소하는 등의 책임을 졌다.
> 
>**2. 수정 후**
> - 중앙 관리자 `AreaEffectSkill`을 없앴다. 
> - `Skill` : **어떤 효과를 발동시키는 명령자.** `Buff`를 부여하고 `Controller`을 맵에 나타나게 한다.
> - `Buff` : **캐릭터에게 붙어서 특정 상태를 부여**하는 표식 / 지속 효과 등등
> - `Controller` : **씬에 생성되어서 주어진 역할을 하고 사라진다.**

- 구체적인 내용(제미나이 참고, 복붙은 아니고 옮겨적음)
1. `AreaEffectSkill` : 제거됨
- 대신 이것을 상속받던 스킬들 `MeteorSkill, AreaHasteHealSkill, ArcaneFieldSkill`은 `ActiveSkill`을 직접 상속받도록 변경했음
- 왜?
	- **추상화가 잘못됨** : "범위 효과"라는 기준만으로 지속형, 즉발형 스킬을 억지로 묶었다.
	- **불필요한 복잡성** : 상속이 깊어질수록 코드를 이해하고 추적하기 어려워진다.
	- **책임 과다** : 씬 객체 추적, 오브젝트 풀링, 시각화 등의 책임을 한 클래스가 모두 졌다.

2. 효과 추적 및 생명주기 관리
- `Skill`이 `Dict`와 `Event`를 통해 씬의 게임 오브젝트들을 추적하고 정리하는 로직 제거
- 어떻게 : 
	- **부모-자식 관계 활용** : 효과 객체들의 부모를 시전자로 설정해서 시전자가 파괴되면 함께 파괴되도록 수정
	- **코루틴 활용** : 각 컨트롤러는 `Update` 대신 코루틴을 활용해 자신의 생명주기를 직접 관리하도록 변경
- 왜 :
	- 안정성 : 유니티 엔진의 내장된 생명주기 관리 기능 활용 
	- 단순함 : **`Instantiate` 한 줄로 다 정리됨 - 스킬 시전 오브젝트를 부모로 둔 것** 만으로 수많은 이벤트 구독 / 해제 코드가 다 필요없어졌다.
	- 성능 : 수십 개의 객체가 `Update`를 돌리는 것보다 코루틴의 `yield return`이 성능에 더 유리함

3. 시각화 및 효과 생성 로직
- `AreaEffectSkill`에 있던 공통 메서드들이 각 개별 스킬 클래스의 `private` 헬퍼 메서드로 변경됨
- 각 스킬은 자신의 `PlaySkillEffect`에서 각각의 메서드에 정의된 함수들을 호출한다.
- 왜 :
	- 유연성 : 모든 범위 스킬이 똑같은 방식으로 시각화/효과를 생성하지 않는다. 각 스킬은 자신만의 고유한 연출과 로직을 가질 수 있는 자유가 부여됐다.
	- 응집도 향상 : 해당 스킬의 모든 로직은 해당 클래스 파일 내에 있어서 가독성과 유지보수성이 향상됐다. 



### SlashSkill
- `SlashSkill`
	- 일단 `SlashSkillController`에서 **파티클 시스템의 충돌 로직을 사용하고 있는데, 이게 성능적인 부하가 크다고 함.** 
	- 그러나 **실질적으로 일회성으로 나가는 스킬이고 판정되는 경우도 많지 않기 때문에 유지**함
```cs
using System.Collections.Generic;
using Skills.Base;
using UnityEngine;

namespace Skills.OperatorSkills
{
    [CreateAssetMenu(fileName = "New Slash Skill", menuName = "Skills/SlashSkill")]
    public class SlashSkill : ActiveSkill
    {
        [Header("Damage Settings")]
        [SerializeField] private float damageMultiplier = 2f;

        [Header("Skill Settings")]
        [SerializeField] private GameObject slashEffectPrefab = default!;
        [SerializeField] private float effectSpeed = 8f;
        [SerializeField] private float effectLifetime = 0.5f;

        protected override void SetDefaults()
        {
            duration = 0f;
        }

        protected override void PlaySkillEffect(Operator op)
        {
            if (slashEffectPrefab == null) return;

            Vector3 spawnPosition = op.transform.position + op.transform.forward * 0.5f;
            Quaternion spawnRotation = Quaternion.LookRotation(op.FacingDirection);
            GameObject effectObj = Instantiate(slashEffectPrefab, spawnPosition, spawnRotation);

            SlashSkillController controller = effectObj.GetComponent<SlashSkillController>();
            if (controller != null)
            {
                controller.Initialize(op, op.FacingDirection, effectSpeed, effectLifetime, damageMultiplier, skillRangeOffset, op.OperatorData.HitEffectPrefab, op.HitEffectTag);
            }
        }
    }
}
```
> 매우 짧아졌다. 
> 추가로 눈치 못챘던 버그가 있었는데, `Controller`에 `damageMultiplier`가 들어가야 하는데 스킬에서 대미지를 계산해서 직접 넣고 있었다. 즉 대미지 계산 배율이 수백 단위가 된 거임. 이상하게 쎄긴 하더라.



### ShieldSkill
- `ShieldBuff`
```cs
using UnityEngine;
using UnityEngine.VFX;

public class ShieldBuff : Buff
{
    private float shieldAmount;
    private GameObject shieldEffectPrefab;
    private GameObject currentShieldEffect;

    public ShieldBuff(float amount, float duration, GameObject vfxPrefab)
    {
        this.shieldAmount = amount;
        this.duration = duration;
        this.shieldEffectPrefab = vfxPrefab;
        this.buffName = "Shield";
    }

    public override void OnApply(UnitEntity owner, UnitEntity caster)
    {
        base.OnApply(owner, caster);
        if (owner is Operator op)
        {
            op.shieldSystem.OnShieldChanged += HandleShieldChanged;
            op.ActivateShield(shieldAmount);
            PlayShieldVFX(op);
        }
    }

    public override void OnRemove()
    {
        if (owner is Operator op)
        {
            op.DeactivateShield();
            op.shieldSystem.OnShieldChanged -= HandleShieldChanged;
            RemoveShieldVFX();
        }
        base.OnRemove();
    }

    private void PlayShieldVFX(Operator op)
    {
        if (shieldEffectPrefab != null)
        {
            currentShieldEffect = GameObject.Instantiate(shieldEffectPrefab, op.transform.position, Quaternion.identity, op.transform);
            currentShieldEffect.GetComponent<VisualEffect>().Play();
        }
    }

    private void RemoveShieldVFX()
    {
        if (currentShieldEffect != null)
        {
            currentShieldEffect.GetComponent<VisualEffect>()?.Stop();
            GameObject.Destroy(currentShieldEffect, 1f);
        }
    }

    private void HandleShieldChanged(float currentShield, bool isShieldDepleted)
    {
        if (isShieldDepleted)
        {
            owner.RemoveBuff(this);
        }
    }
}
```

- `ShieldSkill`
```cs
using UnityEngine;
using Skills.Base;
using UnityEngine.VFX;

namespace Skills.OperatorSkills
{
    [CreateAssetMenu(fileName = "New Shield Skill", menuName = "Skills/Shield Skill")]
    public class ShieldSkill : ActiveSkill
    {
        [Header("Shield Settings")]
        [SerializeField] private float shieldAmount = 500f;

        [Header("Stat Boost Settings")]
        [SerializeField] private StatModifierSkill.StatModifiers statMods;

        [Header("Shield Visual Effects")]
        [SerializeField] private GameObject shieldEffectPrefab = default!;

        private ShieldBuff? _shieldBuff;
        private StatModificationBuff _statBuff;

        protected override void PlaySkillEffect(Operator op)
        {
            _statBuff = new StatModificationBuff(this.duration, statMods);
            op.AddBuff(_statBuff);

            _shieldBuff = new ShieldBuff(shieldAmount, this.duration, shieldEffectPrefab);
            op.AddBuff(_shieldBuff);
        }

        protected override void OnSkillEnd(Operator op)
        {
            // 지속시간이 다 되어서 끝나는 경우 - 적용된 버프들을 제거함
            if (_statBuff != null) op.RemoveBuff(_statBuff);
            if (_shieldBuff != null) op.RemoveBuff(_shieldBuff);
            _statBuff = null;
            _shieldBuff = null;

            base.OnSkillEnd(op);
        }

    }
}
```
> 이렇게 구현하는 경우 `ShieldBuff`가 깨지더라도 `StatBuff`는 지속시간 동안 유지된다. 
> 만약 `ShieldBuff`가 깨졌을 때 `StatBuff`도 깨지게 하고 싶다면? 아래처럼 구현하면 된다. 여기선 사용하지 않음.
```cs
        protected override void PlaySkillEffect(Operator op)
        {
            _statBuff = new StatModificationBuff(this.duration, statMods);
            op.AddBuff(_statBuff);

            _shieldBuff = new ShieldBuff(shieldAmount, this.duration, shieldEffectPrefab);
            op.AddBuff(_shieldBuff);

            // _shieldBuff.LinkBuff(_statBuff);
            // _shieldBuff.OnRemovedCallback += () => OnSkillEnd(op);
        }
```

- 버프가 제거될 때 아래의 메서드가 호출되고
```cs
    public virtual void OnRemove()
    {
        // 연결된 버프들이 있다면 우선 제거함
        foreach (var buff in linkedBuffs.ToList())
        {
            owner.RemoveBuff(buff);
        }

        // 스킬의 후처리 콜백 호출
        OnRemovedCallback?.Invoke();
        RemoveVFX();
    }
```

- 이를 호출하는 `UnitEntity.RemoveBuff`는
```cs
    public void RemoveBuff(Buff buff)
    {
        if (activeBuffs.Contains(buff))
        {
            buff.OnRemove(); // 만약 연결된 다른 버프들이 있다면 여기서 먼저 제거됨
            if (activeBuffs.Remove(buff))
            {
                OnBuffChanged?.Invoke(buff, false);
            }
        }
    }
```

의 구조를 가진다. 그래서 위처럼 `Shield -> StatModifier`로 연결된 상태이고, 지속시간이 만료되지 않았는데 Shield가 파괴될 상황이라면

1. `ShieldBuff`에 연결된 `StatModiferBuff`부터 제외됨
2. 마지막으로 `ShieldBuff`가 파괴됨
3. `ShieldBuff`에 달아뒀던 `OnSkillEnd(op)`에 의해 스킬 종료 메서드가 실행됨

의 구조를 갖는다.

- 쉴드랑 코스트 회복 잘 되는지 보고 마무리하겠음.
- 스택 오버플로우 이슈가 있었다. 해결함. 근데 `Defender`가 이상하게 단단한 느낌도 있음.



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