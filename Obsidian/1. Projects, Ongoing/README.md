# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 남은 작업
- **이펙트 공부 후 수정 중** 
	- **VFX 그래프를 썼던 것도 모두 파티클 시스템으로 변환**
- 1-3 밸런싱, 보스 추가
- 테스트 및 수정

### 보고 참고할 내용
- [[오퍼레이터들 스탯 정리]]
- [[적 스탯 정리]]
- [[프로크리에이트로 텍스쳐 작업 시 유의할 점]]
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

>[!wip]
>- 보스 구현
>	- 스킬
>		- 이펙트
>			- `Anticipation`, `Cast` 이펙트 구현 완료
>				- 참고) 스킬 시전에 소요되는 시간은 **0.5초**이며, 이 동안 보스는 움직이거나 공격하지 못한다.
>			- `Climax`, 즉 실제 폭발 효과는 구현해야 함
>		- 효과(스크립트)
>			1. 13칸에 범위 대미지
>			2. 자신을 저지한 적을 지나가기
>	- 보스 스크립트 구현(거의 다 된 듯
>- 쉴드 스킬은 냅둘지 말지 아직 모르겠음. 


>[!todo]
>- 남은 작업
>	- 스테이지 1-3 완성
>	- 남은 스테이지들 밸런스 수정

# 250902 - 짭명방
>[!Done]
>- `EnemyBoss` 스크립트 구현
>	- 대부분 세팅은 된 것 같음. 스킬 구현하면서 다듬어나가면 될 듯.

## EnemyBoss 스크립트 구현하기
- `Enemy.cs`에서 `enemyData` 기반으로 구현된 스크립트들을 모두 프로퍼티인 `BaseData` 기반으로 수정함. 
- `Skills`로 묶었던 스킬들을 `MeleeSkills`, `RangedSkills`로 구분함. 
	- 두 스킬이 돌아가는 상황이 다르기 때문에 이렇게 구현해뒀다.
		- 저지당하는 상황이라면 `MeleeSkills`가 나가야 함
		- 공격 범위 내에 적이 있고 저지당하지 않는 상황이라면 `RangedSkills`가 나가야 함


### 공용 스킬 쿨타임 적용
> - 두 스킬을 1개씩만 구현할 예정이지만, 만약 종류마다 여러 스킬이 있다고 하면 스킬의 쿨타임과 스킬 종류의 쿨타임을 이중으로 구현해서 스킬을 연속적으로 사용하지는 않게끔 구현할 수도 있겠다. 

그냥 떠오른 생각이라 적어뒀는데, 스킬이 나가는 상황을 가정해보면 **여러 개의 스킬을 한꺼번에 쏟아내는 상황을 막아야 하므로 여러 스킬들이 공유하는 쿨다운과 한 스킬이 갖는 쿨다운은 별도로 구현해야 한다.** 
- 이건 의도 차이긴 하지만 일반적으로 게임 시작하자마자 모든 스킬을 와다다다 쏟아내는 보스는 거의 본 적이 없을 거임. 

위에서는 스킬 종류별로 구현하겠다~고 했는데, 모든 스킬이 참조하는 공용 쿨다운과 개별 스킬이 사용하는 쿨다운만 이중으로 구현하면 될 것 같음. 

#### 부록
- 지금 `EnemyBoss`의 구현은
```cs
    private HashSet<EnemyBossSkill> meleeSkills;
    private Dictionary<EnemyBossSkill, float> meleeSkillCooldowns = new Dictionary<EnemyBossSkill, float>();

    private HashSet<EnemyBossSkill> rangedSkills;
    private Dictionary<EnemyBossSkill, float> rangedSkillCooldowns = new Dictionary<EnemyBossSkill, float>();
```
이런 느낌으로 들어가 있다. 만약 한 종류에 여러 가지의 스킬을 넣는다면 `HashSet`이 과연 적합할까? 예를 들면 **보스가 스킬을 사용하는 순서**를 정하고 싶을 수도 있다. 스킬 1이 나간 다음 2가 나가도록 고정하는 방식으로.

실제로 **그런 구현을 원한다면 해쉬셋보다는 리스트가 낫다**고 한다. 물론 `O(n)`이긴 한데 **`Big-O`를 따질 정도로 스킬의 개수가 많지 않음.**

그래서 얘네들은 `List`로 돌려놨다. 

#### 지식이 늘었다1
>[!note]
>1. 반복문에서 참조형 자료를 직접 수정하는 경우에는 복사본을 만들어서 작업해야 한다.
>	- 이렇게 복사해서 작업해도 주의할 점은 있음. 아래는 예시.
>		- 리스트 1에 적 1이 있다 -> 리스트 1을 복사해 리스트 2를 만든다 -> **리스트 2의 적 1의 체력을 수정한다면 리스트 1의 적 체력도 수정**된다
>		- 즉 주소를 복사하는 거지 값을 복사하는 개념이 아님. 
>		- **이 `ToList()`, `new List<T>(OriginalList>` 등은 얕은 복사임.**
>2. 읽기의 경우는 큰 상관없음. 





# 250901 - 짭명방

>[!done]
>- 보스 범위 스킬 구현
>	- `Anticipation` 이펙트(완)
>- 스킬 시전 시 보스에게 나타나는 스킬 시전 이펙트 `VFX_Cast_BossRangedSkill_v1` 구현
- 작업을 빨리 하고 싶은데 설계할 때 어떻게 해야할 지 뇌가 하얘지는 것이에요~ 이럴 때 참 답답하고 조급해진다.
- AI를 활용해서 개요를 짜고 있긴 한데 가끔 필요 없거나 기반 스크립트를 다 넘겨주지 않아서 괜히 더 복잡해지는 답변이 올 때도 있다. 
- AI를 어떻게 써야 할지는 쓸 때마다 참 고민되는 부분. 클로드 코드마냥 스크립트를 항상 박아놔야 하나. 
## 보스 범위 스킬 구현
- 스크립트(Sun Particle)
```cs
using UnityEngine;
using System.Collections;

public class BossSunSkillController : MonoBehaviour
{
    [Header("Sun Particle Settings")]
    [SerializeField] private ParticleSystem mainParticle;
    [SerializeField] private GameObject sunObject;
    [SerializeField] private float sunSpeed = 1f; // 일단 이렇게 구현. ScriptableObject를 받아서 구현해야 할 수도 있음.
    [SerializeField] private float sunDuration = 3f; // 파티클이 내려오는 시간.

    [SerializeField] private float startYPos = 3f;

    private void Start()
    {
        Initialize();
    }

    // 파티클 시스템 실행 및 효과 재생
    public void Initialize()
    {
        StopAllCoroutines();

        // 위치 초기화
        sunObject.transform.localPosition = new Vector3(0f, startYPos, 0f);

        // 파티클 시스템 재생
        mainParticle.Play(true);

        // Sun 파티클 위치 변화 코루틴 시작
        StartCoroutine(FallSunParticle());
    }

    public IEnumerator FallSunParticle()
    {
        float elapsedTime = 0f;

        while (elapsedTime <= sunDuration)
        {
            sunObject.transform.localPosition = new Vector3(
                sunObject.transform.localPosition.x,
                sunObject.transform.localPosition.y - sunSpeed * Time.deltaTime,
                sunObject.transform.localPosition.z
            );

            elapsedTime += Time.deltaTime;

            yield return null;
        }        
    }
}
```

> 이거 자체는 Boss 스킬에 포함된 것중 하나임. 본격적인 효과가 나타나기 전, 해당 위치에 대미지가 들어갈 것이라는 `Anticipation` 개념이다. 

## Boss 스크립트 설계와 리팩토링
- `Enemy`의 동작을 따르되 2개의 스킬이 추가된다.
- 근거리 스킬 : 저지당했을 때 자신을 저지한 적에게 대미지를 주고 그 다음 타일로 통과
- 원거리 스킬 : 범위 내의 오퍼레이터를 중심으로 13칸 범위에 대미지 (및 틱뎀)

- `EnemyBoss`의 거동은 `Enemy`를 상속받되 스킬에 관한 로직만 추가하면 될 것 같음. 

>[!question]
>- 기존의 `BaseSkill`은 `Operator`를 위주로 설계되었음. `Enemy`에 사용되는 스킬은 별도의 개념으로 가야할까?
>-> `BaseSkill` 위에 추상화 클래스를 하나 더 추가한다. `UnitEntity Caster`를 받게끔. 

### Update - 템플릿 메서드 패턴
- 기존에 `Enemy`에 구현한 행동 로직을 별도의 가상 메서드로 분리한다. 자식 클래스가 생겼기 때문에, 자식 클래스에서 별도의 행동이 필요하다면 거기서 오버라이드해서 사용하도록 함.
```cs
protected override void Update()
{
	if (StageManager.Instance!.currentState == GameState.Battle && // 전투 중이면서
		currentDespawnReason == DespawnReason.Null // 디스폰되고 있지 않을 때
		)
	{
		// 행동이 불가능해도 동작해야 하는 효과
		UpdateAttackDuration();
		UpdateAttackCooldown();
		base.Update(); // 버프 효과 갱신

		if (HasRestriction(ActionRestriction.CannotAction)) return;

		// 판단하고 행동하는 로직을 가상 메서드로 분리, 자식 클래스에서 별개로 구현할 수 있도록 함
		// 이를 템플릿 메서드 패턴이라고 한다.
		DecideAndPerformAction();
	}
}

// 행동 규칙. 원래는 Update에 있던 내용들이다.
protected virtual void DecideAndPerformAction()
{
	if (nextNodeIndex < pathData.nodes.Count)
	{
		if (AttackDuration > 0) return;  // 공격 모션 중

		// 공격 범위 내의 적 리스트 & 현재 공격 대상 갱신
		SetCurrentTarget();

		// 저지당함 - 근거리 공격
		if (blockingOperator != null && CurrentTarget == blockingOperator)
		{
			if (AttackCooldown <= 0)
			{
				PerformMeleeAttack(CurrentTarget!, AttackPower);
			}
		}
		else
		{
			// 바리케이트가 타겟일 경우
			if (targetBarricade != null && Vector3.Distance(transform.position, targetBarricade.transform.position) < 0.5f)
			{
				PerformMeleeAttack(targetBarricade, AttackPower);
			}

			// 타겟이 있고, 공격이 가능한 상태
			if (CanAttack())
			{
				Attack(CurrentTarget!, AttackPower);
			}

			// 이동 관련 로직.
			else if (!isWaiting)
			{
				MoveAlongPath(); // 이동
			}
		}
	}
}
```
> 이전에도 이런 방식으로 몇 번 구현한 적은 있지만, 이름은 몰랐다. 이걸 **`템플릿 메서드 패턴`이라고 함.** 

## Boss Skill 구현
### BaseSkill 리팩토링
- `BaseSkill -> OperatorSkill`로 변경
- 기반이 되는 `UnitSkill`이라는 스크립트 추가.
	- `BossSkill`은 이를 상속받아서 만듦.
- `BaseSkill`은 `OperatorSkill`으로 이름을 변경. 
	- 최상위 클래스의 `caster`는 `UnitEntity`이므로 `Operator`로 사용하고 싶다면 별도의 읽기 프로퍼티 필드만 하나 추가해주면 됨.
```cs
protected Operator Caster => caster as Operator;
```

### Boss Skill 구현
- 보스가 하나 있는 게임이라 `EnemyBoss`에 죄다 통합해버리는 방법이 더 빠르겠지만, 그래도 **설계를 해봤냐 아니냐는 차이가 있는 것 같아서 구현**함.
- 우선 보스 스킬들이 공유하는 요소를 만듦
	- 쿨타임
	- 이펙트
	- 메서드
- 스크립트들 정리하고 있다. 뭐 하고 있는지 슬슬 헷갈리므로 이건 쉴 때가 되었다는 뜻. 쉬고 마저 한다.

### CastVFX 구현
- 내 게임에는 **엔티티들의 애니메이션이 없기 때문에 스킬을 시전할 때는 스킬을 시전한다는 표시**를 하는 게 좋을 듯
- **오퍼레이터의 경우 별도로 구현하지 않았음** : 사용자의 조작 직후에 바로 스킬이 실행되기 때문에, 스킬 실행한다는 별도의 표시가 필요할 것 같지 않아서.

![[VFX_Cast_BossRangedSkill_v1.gif]]

놀랍게도 재탕할 것들이 꽤 있었다. `Cylinder`의 벽 부분이라든가, `Trail02` 같은 요소라든가. 
저 `Trail`이 각이 져서 사라지는 모습이 살짝 어색하므로 `Trail`은 파티클이 사라질 때 함께 사라지도록 수정했다. 

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