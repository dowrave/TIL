- 현재 어떤 작업 중인지 기록 중
# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

---
# 작업 일지

## 짭명방 
- 지난 내역 : [짭명방 프로젝트 일지 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- [[기타 참고 사항]]

>[!note]
>
>  **"흐름"으로 설명하는 것 외에도 "원리"도 곁들이면 좋다**
> - 지금은 "어떤 이슈가 발생해서 어떻게 해결했더라"로 정리했다면
> - 추가로 **이슈가 '왜' 발생했는지**까지 정리해두면 나중에 볼 때 다시 헷갈리지 않을 수 있음
> 
> 모든 코드를 일일이 기억하는 건 불가능함! 나중에 봤을 때 '이러면 되겠다'는 감이 바로 잡히도록 메모해보자.

>[!plan]
>- (진행 중)`UnitEntity` 리팩토링 : `God Class` -> `퍼서드Facade` 패턴
>1. 스테이지 밸런싱, 버그 수정
>2. 경로 로직 관련 현재 간헐적으로 발생하는 이슈들 보이면 처리
>	- 경로가 막혔을 때 바리케이드를 제거하지 않고 그대로 지나가는 현상
>	- 경로가 바뀔 때 살짝 뒤로 빼는 현상. 
>		- 최단 경로를 찾기 때문에 뒤로 빼는 현상이 나오면 안됨.
>3. 소리 추가
>4. 프로젝트 정리글 작성

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움
> - `Enemy`가 사라질 때 풀 태그의 키가 없다는 오류 
> - `Barricade` 배치 시 가끔 배치되지 않음
> - `Operator` 체력이 다했을 때 사망처리가 되지 않았는데도 적이 지나가는 현상

## 현재 이슈
- `Enemy`의 컨트롤러 구현하기
- `PathNavigator`의 구현 관련 
	- `Enemy`나 `PathIndicator`에 구현된 필드가 너무 많아서 굳이 서브 컴포넌트로 빼둔 의미를 잘 못 느끼겠음
- `CombatController`
	- 적군 구현 필요
- `ArcaneFieldSkill` : 범위 내에 적이 없을 때 사용되면 안되게 수정 필요
	- `ArcaneFieldSkill`의 타격 이펙트에 노란색도 추가해볼까?
- `UnitEntity.ExecuteSkillSequence`을 `SkillController`에 통합하기
	- `BossSkillController`에서 사용 중. 보스가 스킬을 시전할 때 잠깐 멈추는 효과로 구현하고 있다.

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

## 260119
>[!done]
> - `BuffVFX`
> 	- 필요 없다고 생각해서 지웠는데 필요함. 복구. 
> 	- 오브젝트 풀링에서 스테이지 시작 시에 생성되도록 변경
> - `SlowBuff`
> 	- VFX 추가
> 	- 감을 다 잃었다. 이미지 작업은 자제하거나 AI님의 도움을 받자.
> 		- 제미나이는 투명한 배경 이미지를 못 만든다. grok이나 chatGPT가 더 나음.
> - 기타 이슈 / 수정사항
> 	- `AreaHasteHeal` : `Barricade`에도 힐이 들어가는 이슈
> 	- `ArcaneFieldSkill` : 둔화율 30% -> 50%로 증가

### Buff VFX 효과 복구
- 버프에 VFX를 달아뒀던 이유는 `Stun, Slow` 등의 효과를 시각화하기 위함이었음
- 이걸 모르고 주석처리해뒀다. 다행히 지우진 않았어서 원래대로 돌려놓고 주석 추가

> [!note]
> - 참고) 버프의 VFX를 가져오는 방식은 `BuffVFXManager.GetBuffVFXObject(Buff, Transform)`으로 가져옴

#### 추가) 스테이지 불러올 때 Buff의 VFX들도 사전에 생성되도록 수정
```cs
// 기존 코드 : BuffVFXManager
private void Start()
{
	if (Instance == null)
	{
		Instance = this;
		if (VFXDatabase != null)
		{
			// 여기서 VFX 오브젝트 풀을 만듦	
			VFXDatabase.Initialize();
		}
	}
	else
	{
		Destroy(gameObject);
	}
}
```

이전에 오브젝트 풀링으로 관리되는 객체들은 모두 스테이지가 시작되는 시점에 생성되도록 옮겨놨는데, 버프 VFX는 해당 로직에서 빠져 있음

1. `VFXDatabase.Initlaize()` 부분이 갖고 있는 버프의 VFX들을 생성하는 과정이다. 저 부분을 `BuffVFXManager`의 별도 코드로 분리하고, `StageManager`의 오브젝트들을 생성하는 과정에 끼워넣었음.
2. (크게 중요한 부분은 아니지만) 로딩 화면에서 로딩 화면 게이지를 채우는 로직이 있고 이 로직의 기반이 오브젝트 갯수이므로 "총 생성해야 하는 오브젝트의 갯수"를 얻기 위한 코드도 별도로 구현해줬음

```cs
// StageManager.cs
private IEnumerator PreloadStageObjectPoolsCoroutine(List<SquadOperatorInfo> squadData, Action<float> onProgress)
{
	// ...
	
	float buffVFXTasks = BuffVFXManager.Instance.GetAllVFXPoolCounts();
	float totalTasks = enemyCountDict.Count + squadData.Count + (stageData?.mapDeployables?.Count ?? 0) + buffVFXTasks;
	float completedTasks = 0f;

	// ...

	// 3. Buff의 오브젝트 풀들 생성
	// 이 메서드가 Awake, OnEnable 후에 실행되므로 사용 가능함
	BuffVFXManager.Instance.CreateBuffVFXPools();
}
```

### SlowBuff VFX 구현

- 초안
- 계속 나타나야 함
- 발 밑에 나타남(버프 VFX 설정의 `offset`으로 가능)

...을 갖고 작업했으나 상상력의 부족 + 슬로우 디버프 레퍼런스가 생각보다 없어서 그냥 로고로 구현하는 걸로 진행

![[Pasted image 20260119181227.png]]


...도 생각보다 직관적이지 않다. 화살표를 겹치면 안될 것 같긴 함.

- 그것보다도 가장 단순한 느낌으로 아래 2개 화살표 텍스쳐를 하나 그리고 슬로우에 걸리면 파티클이 위에서 아래로 내려가는 형태로 만들어보겠음.

![[SlowBuff01_256.png]]
를 텍스쳐 시트에 추가

상단 아이콘도 달팽이를 추가했음. 
> AI에게 아이디어 스케치를 요청하는 방법이 꽤 괜찮아보인다. 

- 셰이더 감을 다 잃어서 생각보다 애먹었음;;
	1. AI로 투명한 배경의 달팽이 이미지를 만듦
	2. 지속시간 동안 약간의 깜빡임 구현을 위해 시간에 따른 알파값 변화를 셰이더에 넣음
	3. 근데 원본 이미지가 분명 투명했고 프로젝트 탭에서도 투명하게 나오는데 셰이더에 텍스쳐를 집어넣으면 텍스쳐에 정체불명의 흰색들이 나옴
		- 이 흰색 처리 방법 때문에 한참 헤맸다.
		- 원본 이미지에 `Color`와 `Vertex Color`을 곱한 뒤에, 알파값만 `Split`해서 따로 처리하면 됨. 
		- 이 경우, 알파값의 깜빡임을 따로 구현해놨으니까 그 값이랑 분리한 알파값을 곱해서 마스터 노드의 알파값에 전달하면 된다. 
![[Pasted image 20260119215405.png]]
이펙트 구현해서 `BuffVFXDatabase`에 추가하면 이런 느낌으로 나타난다.
이탈한 `Enemy`는 디버프가 풀려 있고, `ArcaneFieldSkill` 내부에 있는 `Enemy`는 `SlowBuff`가 들어간 모습을 시각적으로 확인할 수 있음. 

### 기타 이슈
- `AreaHasteHealSkill` 
	- `Barricade`에도 힐이 들어가는 이슈(이펙트가 나타남)
		- 타입 체크만 추가하면 ㅇㅋ





## 260116

>[!done]
>1. `UnitStats`으로 시작되는 `struct`들의 구조 수정
>	- 세터 프로퍼티 제거 & 생성자 추가 & 필드로 초기화하는 것까지 고려
>2. 직렬화와 인스펙터 노출 공부
>3. 스탯 정리 : 공격 속도와 쿨다운 개념 명확하게

## 스탯 정리 : 공격 속도 vs 쿨다운

### 수정 방향 
1. 캐릭터마다의 기본 공격 속도를 `BaseAttackRate`으로 생각
	- 현재의 `AttackSpeed`인데, 이를 `BaseAttackRate`으로 변경
	- 이 `BaseAttackRate`는 불변값
2. 공격 속도 스탯 개념을 추가 : `AttackSpeed`. 
	- 별도의 설정이 없으면 런타임에만 동작하며 `1`에서 시작하는 값
	- 버프로 계수가 추가되는 값
3. 최종 값은 아래의 공식을 따름 - 즉 공격 속도는 분모로 들어감
$$
최종 쿨다운 = \frac{기본쿨다운}{공격속도}
$$

- `BaseAttackRate = 2.0`이라고 하면
- 예시) 공격 속도 50% 증가 
	- 기본 공격 속도는 1이므로 50% 증가는 1.5
	- 원래 최종 쿨다운 : 2
	- 버프 최종 쿨다운 : 1.333

### 반영
1. 기존의 `AttackSpeed` => `BaseAttackCooldown`으로 변경
2. `StatController`에서 관리하는 타입은 `AttackSpeed`로 동일
	- 기본값에 `BaseAttackCooldown`을 넣음
	- 모디파이어에 `AttackSpeed`에 해당하는 값을 넣음
	- `GetStats`에서 나가는 부분만 수정해주면 됨
```cs
public float GetStat(StatType type)
{
	// 1. 오버라이드에 있는 값이라면 최우선으로 나감(덮어쓰기 값)
	// 공격속도 관련해서 생각할 게 있긴 한데 지금까진 공격속도를 덮어쓰는 로직은 없었음
	if (_overrides.TryGetValue(type, out float overrideValue))
	{
		return overrideValue;
	}

	float baseValue = _baseStats.TryGetValue(type, out float val) ? val : 0f;
	float modifierValue = _modifiers.TryGetValue(type, out float mod) ? mod : 0f;
	float calculatedValue;

	if (type == StatType.AttackSpeed)
	{
		// 공격 속도 : 기본 1
		float attackSpeed = (1 + modifierValue);

		// 최종 공격 쿨다운 : 기본 공격 쿨다운 / 공격 속도
		// 공격 속도가 빨라진다 = attackSpeed가 올라간다 = 공격 쿨다운이 줄어든다
		calculatedValue = baseValue / attackSpeed;
	}
	else
	{
		calculatedValue = baseValue * (1 + modifierValue);
	}
	
	return calculatedValue;
}
```
## 스탯 시스템 수정

>[!flow]
>- `struct`로 구현된 스탯 시스템의 각 프로퍼티의 setter 제거
>	- setter를 제거하면서 사용 불가능해진 코드가 발생 : 레벨에 따른 스탯을 갱신하는 부분이 프로퍼티의 세터로 구현되어 있었음 -> 수정 필요
>- `struct`의 생성자를 구현하는 것으로 접근
>	- `struct`는 `UnitStats` -> `DeployableUnitStats` -> `OperatorStats` 처럼 이전의 코드를 재활용하는 부분이 있음
>		- 하위 구조체에서 상위 구조체를 받아 초기화하는 방식으로 구현하는 게 가능함
>		- 하지만 이렇게만 구현하면 쓰는 입장에서, `OperatorStats`을 초기화하기 위해 `UnitStats`와 `DeployableUnitStats`을 초기화한 다음에 `OperatorStats`에 인자로 넣어야 함
>		- 그래서 **각 `struct`의 초기화 메서드는 더 상위 구조체를 받는 것 외에도, 단순히 각 필드를 받는 식의 구현도 추가**해줘야 함(편의 생성자) 

1. 기존의 `SO.UnitStats` 관련 필드들의 프로퍼티의 세터들을 제거

2. 각 `struct`에 생성자 추가
```cs
[System.Serializable]
public struct UnitStats
{
    [SerializeField] private float _health;
    [SerializeField] private float _defense;
    [SerializeField] private float _magicResistance;

    public UnitStats(float health, float defense, float magicResistance)
    {
        _health = health;
        _defense = defense;
        _magicResistance = magicResistance; 
    }

    public float Health => _health;
    public float Defense => _defense; 
    public float MagicResistance => _magicResistance;
}

[System.Serializable]
public struct DeployableUnitStats
{
    [SerializeField] private UnitStats _baseStats;
    [SerializeField] private int _deploymentCost;
    [SerializeField] private float _redeployTime;

    public DeployableUnitStats(UnitStats baseStats, int deploymentCost, float redeployTime)
    {
        _baseStats = baseStats; // UnitStats 자체를 전달해주면 됨
        _deploymentCost = deploymentCost;
        _redeployTime = redeployTime;
    }

    public int DeploymentCost => _deploymentCost;
    public float RedeployTime => _redeployTime;
    public float Health => _baseStats.Health;
    public float Defense => _baseStats.Defense;
    public float MagicResistance => _baseStats.MagicResistance;
}
```

- 그런데 이렇게만 구현하면 `DeployableStats`을 초기화하기 위해 `UnitStats`의 초기화부터 해야 함. `OperatorStats`까지 있으니 **쓰는 입장에서 구조체 초기화만 3번 해야 함**

- 그래서 **`flat`한 필드들을 받는 생성자를 추가해준다. 대신 가장 하위 구조체에서 상위 구조체들을 초기화하는 로직을 넣어주면 됨.**

```cs
// 모든 유닛에 적용되는 스탯 생성자
public UnitStats(float health, float defense, float magicResistance)
{
	_health = health;
	_defense = defense;
	_magicResistance = magicResistance; 
}
// ---

// 상위 구조체를 받아 초기화
public DeployableUnitStats(UnitStats baseStats, int deploymentCost, float redeployTime)
{
	_baseStats = baseStats;
	_deploymentCost = deploymentCost;
	_redeployTime = redeployTime;
}

// 구조체 없이 필드들만 받아 초기화
public DeployableUnitStats(float health, float defense, float magicResistance, int deploymentCost, float redeployTime)
{
	_baseStats = new UnitStats(health, defense, magicResistance);
	_deploymentCost = deploymentCost;
	_redeployTime = redeployTime;
}
// ---

public OperatorStats(
	DeployableUnitStats deployableUnitStats,
	float attackPower,
	float baseAttackCooldown,
	int maxBlockableEnemies,
	float spRecoveryRate
	)
{
	_deployableUnitStats = deployableUnitStats;
	_attackPower = attackPower;
	_baseAttackCooldown = baseAttackCooldown;
	_maxBlockableEnemies = maxBlockableEnemies;
	_spRecoveryRate = spRecoveryRate;
}

public OperatorStats(
	float health,
	float defense,
	float magicResistance,
	int deploymentCost,
	float redeployTime,
	float attackPower,
	float baseAttackCooldown,
	int maxBlockableEnemies,
	float spRecoveryRate
	)
{
	_deployableUnitStats = new DeployableUnitStats(new UnitStats(health, defense, magicResistance), deploymentCost, redeployTime);
	_attackPower = attackPower;
	_baseAttackCooldown = baseAttackCooldown;
	_maxBlockableEnemies = maxBlockableEnemies;
	_spRecoveryRate = spRecoveryRate;
}
```

생성자에 필드를 일일이 나열하는 게 번거롭긴 한데 그걸 빼면 **생각보다 간단하게 구현할 수 있다.** 괜찮은 방법 같음.

>[!question]
>- 지금처럼 `Nested`한 구조를 유지해야 하는가? 아니면 각 `Stats`을 개별 필드로 받아서 곧바로 쓸 수 있는 `Flat`한 구조를 받아줘야 하는가?

>[!answer] 
> - `Flat` 구조는 초기화가 쉽고, 직접 접근이 미세하게 빠르지만 구조적인 의미를 잃고 코드의 중복도 발생한다
> - `Nested` 구조는 계층 구조가 명확하고 재사용성이 좋은 반면, 중첩 생성자 및 약간의 메모리 오버헤드가 생김(사소함)
> - Claude는 `Nested + 편의 생성자`를 쓸 걸 추천

### private [SerializeField] struct에 관해
- 구조가 헷갈려서 정리함. `struct` 내부에서도 필드의 접근자를 정의할 수 있어서.

```cs
// SO의 설정
[SerializeField] protected OperatorStats stats;
    
// OperatorStats    
[System.Serializable]
public struct OperatorStats
{
    [SerializeField] private DeployableUnitStats _deployableUnitStats;
    [SerializeField] private float _attackPower;
    [SerializeField] private float _baseAttackCooldown; // 기본 공격 쿨다운
    [SerializeField] private int _maxBlockableEnemies;
    [SerializeField] private float _spRecoveryRate;
    
    // ...
}
```

1. `직렬화Serialize` : 메모리의 데이터를 파일에 저장 가능한 형태로 변환하고, 나중에 다시 불러올 수 있게 만드는 과정
	- 저장하고 불러올 수 있게 만드는 과정
2. **`struct`는 유니티에서 기본적으로 직렬화되지 않지만, 직렬화 자체는 가능하다.**
	- 이 부분은 구분을 잘 해야 한다 : "직렬화 가능 여부"와 "유니티의 기본 설정"
	- 유니티에서 아무 설정 없이 직렬화할 수 있는 타입들이 있음(예시는 안 듦)
	- `struct`, `class`는 직렬화할 수 있지만, **기본 설정은 "직렬화하지 않는다"** 임
		- 안전성, 성능, 호환성 등의 이슈 때문. 
		- 직렬화하고 싶다면 따로 명시하라는 의미.
	- `Dictionary`나 `Delegate`는 **직렬화 자체가 불가능.**
	- 주의) 여기서 얘기하는 직렬화와 "`struct`가 값 타입이다"는 완전히 별개의 얘기임
		- 즉 "값 타입이므로 직렬화되지 않는다, 참조 타입이니까 직렬화된다" 라는 말이 아니라는 뜻. 
3. **인스펙터에 노출되는 필드** : 아래 요소들을 다 갖춤
	- 직렬화 가능한 타입의 필드이면서
	- `public`이거나 `[SerializeField]` 가 붙은 `protected/private`이며
	- `const, static, 프로퍼티`가 아니며
	- `[HideInInspector], [NonSerialized]`가 없어야 한다.
4. **`[System.Serializable]`은 해당 타입을 직렬화 가능하게 할 뿐**이다. 내부 필드들은 원래의 규칙을 따른다. 
	- `[System.Serializable]`은 타입 자체를 "직렬화 가능하다"라고 등록하는 것임
	- 만약 `[System.Serializable]`이 없는 타입의 필드가 `[SerializeField]`이 붙으면, 직렬화되지 않고 인스펙터에도 나타나지 않는다.

---
# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
