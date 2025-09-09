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
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`

## 하고 싶은데 못한 것
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.
	- 현재는 정예화 패널에서만 스킬 범위를 볼 수 있음.

---
# 작업 일지

## 짭명방 예정

### 작업 중

>[!wip]
>- 보스 스킬 이슈 수정
>	- 스킬 영역 표시, 폭발 후 도트댐 영역 표시가 정상적으로 나타나지 않고 있음

>[!todo]
>- 보스 구현
>	- 원거리 폭발 스킬
>		- 이펙트
>			- `Anticipation`, `Cast` 이펙트 구현 완료
>			- **폭발, 바닥에 틱댐 이펙트는 구현 필요**
>		- 스크립트
>			- **전체적인 흐름은 정리 완료. 세부적으로 다듬을 필요 있음**
>	- 근거리 통과 스킬
>		- 구현 필요 : **이펙트, 실제 효과**
>	- **보스 스크립트는 구현**한 듯? 이것도 테스트하면서 다듬는 단계
>	- 쉴드 스킬
>		- 유지? 
>- 남은 작업
>	- 스테이지 1-3 완성
>	- 남은 스테이지들 밸런스 수정

# 250909 - VFX 강의
- [[4. 충격 이펙트 - 플립북]]
- [[5. Aftermath 구현 - GroundCrack]]

## 오늘 배운 내용
>[!note]
>1. `Parallax Occlusion`이라는 기술이 있다. 평면적인 텍스쳐에 입체감을 부여하는 기능. 셰이더 그래프에 기본 탑재되어 있으며 위의 5번 파일에 정리해뒀다.
>2. 프로크리에이트로 텍스쳐 작업할 때 약간 "이렇게 하면 투명도가 비슷할까??" 싶은 부분이 생기는데, 눈대중으로 맞다면 크게 문제는 없을 것 같다.
>3. `MaterialMaker`에 대해 처음으로 제대로(?) 본 것 같은데 생각보다 되게 괜찮다. 어렵긴 해도. 예전 버전은 자꾸 튕겼는데 지금 새로운 버전을 받아봤는데 잘 작동함.

- 확실히 몰랐던 내용인 `Material Maker`에 대해 배우기 시작하니까 강의 진행 속도가 느려졌다. 총 44강인데 31강까지 진행했음.
# 250908 - 짭명방..이 아니라 강의 듣기
- 내용은 `유니티 VFX 공부/유데미 강의 - Gabriel Aguiar - 폭발 이펙트 강의`에 정리해두고 있다.

## 유데미에서 또 할인을 시작했다
- 지금 필요한 게 **폭발 이펙트인데 센세의 강좌 중에 폭발 이펙트에 대해 다룬 게 있다.** 이전 강좌의 심화 버전이기도 해서 이거 공부한 다음에 보스 이펙트 작업을 진행하겠음
- `MaterialMaker`에 대해선 알고 있긴 한데 자세히 알고 있다는 느낌은 아니었다. 이것도 알아두면 좋겠다 싶음.

## 오늘 배운 것들
>[!note]
>1. **파티클 시스템의 `Custom Data`, `Custom Vertex Streams` 및 셰이더 간의 설정 연동을 통해서 `Color over lifetime` 외에도 다채로운 방법으로 파티클을 서서히 사라지게 할 수 있다.**
>	- 이 때 설정은 `Custom Vertex Streams`에서 `Custom.xyzw`를 추가한 다음 `UV2`도 추가해야 함. `UV2`가 `Custom`보다 더 위에 온다.
>	- 이렇게 설정하면 셰이더 그래프의 `UV1`을 파티클 시스템의 `Custom Data`에서 입력하는 채널처럼 사용할 수 있음.
>	- 여기서 **`Custom.xyzw`는 `Custom Data` 모듈에서 설정한 값을 셰이더에 넣어주는 역할**을 한다.
>2. **`Limit Velocity over Lifetime`으로 빠르게 튀어나가는 파티클이 급격하게 멈추는 효과를 구현할 수 있다.** `Velocity over Lifetime`의 `Speed Modifier`를 `Curve`로 구현할 필요가 없음.
>3. **텍스쳐의 `MainTex` 프로퍼티에는 `Tiling And Offset`을 인스펙터 단위에서 설정할 수 있게 하는 필드가 있다.** 굳이 셰이더 그래프를 오가지 않아도 됨. 강의에선 `Seamless` 텍스쳐를 여러 번 반복시키기 위해 사용했다.
>4. 블렌더에서 `3D Cursor`의 위치를 숫자로 입력해서 지정하는 방법 :  **`Viewport`에 커서 올림 -> `N` -> `View` -> `3D Cursor`에서 설정**
>5. 블렌더에서 여러 오브젝트 합치는 방법 : `Ctrl + J`
>	- 이후에 **`Remesh`를 해주고 `Apply All Transform`도 해줘야 함!**
- 다른 것들은 대부분 배운 요소를 반복하는 개념인데, 복습이니까 좋다.
- 오늘 진행한 파트들
	- [[1. 사전 작업]]
	- [[2. 폭발과 연기 이펙트]]
	- [[3. 부가 이펙트들]]

- 총 44장 중 21장까지 진행.(물론 7장 정도는 날먹이다) **빠르면 내일, 늦어도 모레 내에는 끝낼 수 있을 것 같다.** 물론 계획대로 된다는 보장은 늘 없다.

# 250907 - 짭명방
>[!done]
>- 보스 스킬 프로토타입 완성
>	- 남은 건 이펙트 수정 정도. 이건 내일 작업해야겠다.
>- 스크립트 수정
>	- 어떤 공격으로 인한 대미지 표시를 할지에 대한 여부는 `AttackSource`에서 설정할 수 있다.
>	- 힐이 0 들어갈 때는 나타나지 않도록 했음

## 보스 스킬 프로토타입 완성
![[VFX_BossSkill_Prototype.gif]]
> 여기서 이제 **폭발 이펙트랑 폭발하고 나서의 세부 이펙트 정도만 수정**하면 될 듯
> 폭발의 경우는 대미지도 표시하겠음.

- 바닥 이펙트의 경우 `SkillRangeVFXController`로 통합했다. 모든 `SerializeField`에 할당하지 않아도 게임 자체는 돌아가게끔 구현함. 

## 스크립트 수정
- `AttackSource` 자체에 `ShowDamagePopup`이 들어가도록 수정함
	- 일단 만져놓긴 했는데 수정할 곳이 더 있을 수도 있다. 
	- 대부분의 경우 공격할 때에 `AttackSource`를 새로 설정하니까 거기에 대미지를 띄울지 말지만 표시해두었음.

여담)
- 구현해놓은 저 스킬의 대미지가 2100이 나온다. 만렙 디펜더가 겨우 살아남을 정도의 대미지인데
	1) 메딕이나 디펜더 등의 스킬에 마법 저항력을 올리는 효과를 추가한다
	2) 그냥 보스 하향하죠?
... 밸런스는 나중에 생각하겠음. 일단 금요일에 "주말에 이거는 고쳐놓자"라고 생각했던 것들을 다 고치긴 했으니까 오늘은 여기까지!


# 250906 - 짭명방
>[!done]
>- 보스 스킬 이슈 수정
>	1) 이펙트들이 한꺼번에 실행되는 현상(해결)
>- 추가) `Boss`는 스킬 시전 중, 즉 `Cast`를 시전하는 0.5초 동안은 멈춰 있음.

## 어제의 이슈 1 수정
- 어제의 이슈 1) 이펙트들이 한꺼번에 좌르륵 실행되는 현상
- 스킬 컨트롤러(사실 주로 핸들러라고 한다)에서 시간 필드를 갖고 있는데 정작 전달을 안해서 0으로 되어 있었음. 그래서 `yield return new WaitForSeconds()`이 있음에도 대기 시간이 없었음.
- 이걸 1시간 헤매서 찾았다 하 ㅋㅋㅋ


# 250905 - 짭명방
>[!done]
>- 보스 원거리 스킬 컨트롤러 큰 흐름은 완성함
>	- 세부적으로는 버그가 발생 중

>[!WIP]
>- 보스 스킬 구현
>	- 컨트롤러 설계 완성
>	- 세부 이펙트(폭발, 바닥) 이펙트 구현

## 컨트롤러의 전체 동작
```cs
private IEnumerator PlaySkillCoroutine()
{
	// 1. 영역 표시
	VisualizeSkillRange(caster);

	// 2. 해 파티클 목표 위치에 떨어지는 효과 실행
	GameObject sunParticleObj = ObjectPoolManager.Instance.SpawnFromPool(skillData.GetFallingSunVFXTag(caster), mainTarget.transform.position, Quaternion.identity);
	FallingSunVFXController sunParticleSystem = sunParticleObj.GetComponent<FallingSunVFXController>();
	if (sunParticleSystem != null)
	{
		sunParticleSystem.Initialize(fallDuration);
	}
	yield return new WaitForSeconds(fallDuration); // 낙하시간 동안 대기

	// 3. 낙하 후에 폭발 이펙트 실행, 대미지를 가함
	PlayExplosionVFX();
	ApplyInitialEffect(null); // 실제 효과

	// 4. 폭발 시 범위 타일들에 임팩트 대미지를 주고 지속 대미지를 입히는 필드가 남음
	PlayPeriodicVFX();
	StartCoroutine(PeriodicEffectCoroutine());

	yield return new WaitForSeconds(lingeringDuration); // 필드 지속시간 동안 대기

		ObjectPoolManager.Instance.ReturnToPool(skillData.GetSkillControllerTag(caster), gameObject); // 풀로 되돌림
}
```

## 스킬 범위 인식 관련
- 원래 `Enemy`에 들어가 있는 `AttackRangeCollider`를 이용하려고 했음
- 근데 공격 타입이 `Melee`면 기본적으로 비활성화하는 설정이다.
- 보스 한정으로 별도로 할당하는 `EnemyBossSkillRangeCollider`를 구현하는 쪽으로 작업
	- 원거리 스킬에 한해서만 동작함


## 아 왜 스킬 안써!!!!!!
- 로 계속 머리 싸매는 중! 


## 세부사항 구현

### 이펙트 실행 중 부모가 파괴됐을 때
```cs
ParticleSystem castVFX = castVFXObject.GetComponent<ParticleSystem>();
if (castVFX != null)
{
	castVFX.Play(true);
}

// 시전 동작 중에는 대기 - 이거 기다리는 중에 비활성화되면 그 아래는 실행되지 않음
yield return new WaitForSeconds(castTime);
ObjectPoolManager.Instance.ReturnToPool(GetCastVFXTag(caster), castVFXObject);
```
> 만약 `CastVFX`가 실행되고 도중에 `caster`가 파괴됐다면?
> - 이 코루틴의 실행 자체가 중지되기 때문에 실행 중인 `castVFX`가 풀로 돌아가지 못함

따라서 그냥 `ParticleSystem`으로 넣는 것보다는 스스로 실행 시간을 정해두고 실행이 끝나면 스스로 오브젝트 풀로 돌아가는 스크립트를 구현하는 게 좋다. 어차피 이펙트는 프리팹으로 구현하기도 하니까.

그래서 이 개념을 확장하면, **오브젝트 풀링으로 관리되는 거의 모든 객체는 이런 구조를 이용하면 좋다.** 

### Caster가 파괴되어도 스킬은 유지되게 하고 싶을 때
>[!note]
>- 참고) 함께 파괴되게 하려면 단순히 `SetParent`로 자식 오브젝트로 넣으면 가장 간단함

>[!question]
>- `caster`가 파괴되어도 스킬 효과를 유지시키고 싶다고 하자. `caster`의 대미지를 가져와야 하는데, `caster`가 파괴되었다. 어디서 가져와야 할까?

>[!answer]
>- **`Initialize`할 때 스킬 컨트롤러 자체에 `caster`의 공격력을 복사해두면 됨.**

이게 답변을 얻으면 와 되게 당연한 거네? 싶은데, 막상 궁금할 때는 이런 생각이 안 떠오른다.


# 250904 - 짭명방
>[!done]
>- 보스 스킬 구현
>- `ScriptableObject`의 정보들을 이용한 `Initialize` 관련
>	- "파라미터가 3~4개를 넘어가면 하나의 객체로 묶는 걸 고려하기 시작해야 한다" 라는 원칙이 있다고 함


## ScriptableObject의 정보를 활용한 Initialize
```cs
// ScriptableObject에서 관리하는 Skill에 관한 정보. Initialize에 개별 필드를 전달하고 있다. 
controller.Initialize(
	caster: caster,
	skillRangeGridPositions: caster.GetCurrentSkillRange(),
	fieldDuration: castTime,
	tickDamageRatio: tickDamageRatio,
	interval: damageInterval,
	castTime: castTime,
	fallDuration: fallDuration,
	lingeringDuration: lingeringDuration,
	hitEffectPrefab: caster.BaseData.HitEffectPrefab, // 임시
	hitEffectTag: $"{caster.BaseData.entityName}_{skillName}"
);
```
> 이런 식으로 구현하면 파라미터가 추가될 때마다 계속 수정해야 함. 

그래서 이 `ScriptableObject`를 직접 전달하는 방식으로 수정한다.

>[!question]
>1. 개별 필드를 `public`으로 정의, 프로퍼티를 별도로 정의할 필요 없이 직접 접근할 수 있게 한다
>2. 지금처럼 개별 필드를 `Initialize`로 전달한다
> 3. **개별 필드를 `[SerializeField] private`으로 정의하고 게터 프로퍼티를 열어놓는다**

...여태까지 해왔던 것처럼, 3번이 가장 좋은 선택이다. 외부에서 실수로라도 변경할 수 없게 해야 하기 때문임. 

## 일단 프리팹 만들고 초기화하는 것까지는 구현함
```cs
private IEnumerator PlaySkillCoroutine()
{
	// 1. 영역 표시
	VisualizeSkillRange(caster, skillRangeGridPositions);

	// 2. 해 파티클 목표 위치에 떨어지는 효과 실행
	fallingSunController.Initialize();
	yield return new WaitForSeconds(fallDuration); // 낙하시간 동안 대기

	// 3. 낙하 후에 폭발 이펙트 실행, 대미지를 가함

	ApplyExplosionDamage();



	// 4. 폭발 시 범위 타일들에 임팩트 대미지를 주고 지속 대미지를 입히는 필드가 남음

	yield return new WaitForSeconds(lingeringDuration); // 필드 지속시간 동안 대기
}
```
> 일단 전체적인 프로세스는 이런 느낌이다.

여기서 추가로 구현할 거는
1. 폭발 이펙트 구현
2. 지속 대미지 이펙트 남기기
3. 지속 대미지 효과 구현하기

... 등이 있으며 지금 고민 중인 거는
- 오브젝트 풀링을 구현하는 위치 : `CastVFX`나 `BossExplosionSkillController`, `hitVFXPrefab`, `AreaVFX` 등은 `ScriptableObject`에서 구현한다고 치는데 스킬의 움직임에 딸려 있는 해 파티클이 떨어지는 효과나 바닥에 남는 지속 대미지 이펙트 등은 어디서 구현하면 좋을지 모르겠음. 
	- 저것들을 프리팹인 `BossExplosionSkillController`에 딸려 있는 개념으로 한번에 넣어서 구현해야 하나? 아니면 아예 따로 빼서 구현해야 하나?
		- 한 번에 넣어서 구현해도 무방함 : 여러 번 사용하지만 한 번에 여러 개가 등장하는 스킬은 아니기 때문임 - 이 경우 스크립트를 어떻게 짜야 하나?
		- `AreaVFX`는 기존에 `ScriptableObject`인 `Skill`의 설계도에서 생성 로직을 넣어놨음. 지속 대미지 이펙트도 같은 로직으로 넣을 수 있음. 

일단 쉬고 와서 계속함.

### 어떤 식으로 구현해야 하는가?

1. 스킬 컨트롤러, 혹은 스킬 핸들러는 파티클 시스템이 나타나는 타이밍과 실제 스킬의 효과를 담당함.
2. 개별 이펙트 프리팹은
	1) `ScriptableObject`에서 `InitializeObjectPool`에서 생성함. 이는 `EnemyBoss`가 생성되는 타이밍에 실행된다.
	2) 스킬 컨트롤러에서는 태그만을 이용해서 원하는 위치에 해당 오브젝트를 생성시킴. 컨트롤러의 자식에 이펙트를 둘지 말지는 경우에 따라 다르다. 꼭 하나로 묶어둘 필요는 없다.

... 이런 식으로 정리하겠음

#### 추가로 ScriptableObject에서 태그를 갖게 할 건데
- 직접 필드를 갖는 것보다는, 항상 동일한 결과를 반환하는 메서드를 구현하는 구현도 가능하다고 함. 이를 `Stateless` 패턴이라고 한다.
- 
```cs
// 항상 동일한 태그를 반환하는 메서드
public string GetSkillEffectTag(UnitEntity caster) => $"{caster.name}_{skillName}_skillEffect";
public string GetHitVFXTag(UnitEntity caster) => $"{caster.name}_{skillName}_hit";
// 기타 이펙트 태그들 구현..
```
이제 스킬 컨트롤러에서는 `SkillData`만 갖고 있고, 관련 정보를 얻으려면 `SkillData`를 참조하기만 하면 됨. `SkillData`에서는 인게임 중에 변하는 정보가 없음!

개요 부분만 쫙 쓰고 일단 마무리함. 비염 기운이 또 올라왔다.
```cs
private IEnumerator PlaySkillCoroutine()
{
	// 1. 영역 표시
	VisualizeSkillRange(caster, skillRangeGridPositions);

	// 2. 해 파티클 목표 위치에 떨어지는 효과 실행
	GameObject sunParticleObj = ObjectPoolManager.Instance.SpawnFromPool(skillData.GetFallingSunVFXTag(caster), mainTarget.transform.position, Quaternion.identity);
	FallingSunVFXController sunParticleSystem = sunParticleObj.GetComponent<FallingSunVFXController>();
	if (sunParticleSystem != null)
	{
		sunParticleSystem.Initialize();
	}
	yield return new WaitForSeconds(fallDuration); // 낙하시간 동안 대기

	// 3. 낙하 후에 폭발 이펙트 실행, 대미지를 가함

	// 필요) 폭발 이펙트 실행 메서드
	ApplyExplosionDamage(); // 실제 효과

	// 4. 폭발 시 범위 타일들에 임팩트 대미지를 주고 지속 대미지를 입히는 필드가 남음
	// 필요) 타일 틱뎀 이펙트 실행 메서드
	ApplyTickDamage(); // 틱댐 실제 효과

	yield return new WaitForSeconds(lingeringDuration); // 필드 지속시간 동안 대기
}
```


# 250903 - 짭명방
>[!done]
>- EnemyBoss 스킬 구현하기
>	- 스크립트 구현 중(못 끝냄)
>- 기타 설계 관련
>	- 리팩토링 - 스킬 범위 필드 저장 위치 변경
>	- 방향에 따른 범위 계산은 정적 클래스 & 메서드로 분리

>[!note]
>오늘의 지식이 늘었다 시리즈
>1. `IEnumerable` 관련 : **`IReadOnlyCollection`으로 게터를 열어두면 됨.** `HashSet` 등으로 저장하더라도 게터에서 가져가는 타입은 저렇게 정의해두면 관련 메서드는 다 그대로 사용할 수 있다. 만약 오류가 뜬다면 **`using System.Linq`** 만 해주면 됨.
>2. 필요하지도 않을 유연성을 위해 코드를 복잡하게 만들지 말자 : `YAGNI` 원칙
>3. `SSoT단일진실공급원` 원칙 : 어떤 위치에 변수를 저장했고 **현재 논리에서 그 변수를 갖고 있더라도, 변수를 가져올 때는 저장한 위치에서 가져오는 설계가 좋다.**
>4. **메서드 오버라이드는 대체**, **메서드 오버로딩은 확장** 개념이다. 
>5. `ScriptableObject`에서 가져선 안되는 필드는 스크립트 단위의 상태에 관한 전역 변수이다. 메서드에서 갖는 지역 상태변수는 상관 없음.
>6. `ScriptableObject`에서 코루틴을 실행할 수 없지만, 코루틴을 정의하고 `MonoBehaviour`에서 실행하라고 전달하는 것 자체는 가능하다.
## 설계 관련

### 리팩토링 - 스킬 범위 관련 메서드와 필드 저장 위치

>[!question]
>- `ScriptableObject`에서 런타임에서만 동작하는 필드를 관리하는 게 좋지 않다. 라는 걸 저번에 배웠다.
>	- `ScriptableObject`는 불변의 설계도 개념이다. 런타임 중에만 동작하는 상태를 담는 요소가 아니기 때문이다.(용도에 따른 구분은 명확히 해두는 게 좋다 - `Operator`와 현재 성장 상태인 `OwnedOperator`를 구분했듯이)
>- 그러면 **어디에 스킬 범위를 저장해야 할까?** : 스킬을 사용하는 주체에 저장한다.
- `Operator`의 경우, 지금 상황에선 `Operator`말고는 없어보임. 
	- `Operator`가 배치될 때 자체에서 스킬 범위도 함께 갖도록 수정한다.

- 위치와 오퍼레이터의 회전에 따른 그리드 포지션을 계산하는 로직은 `ActiveSkill`에 있었다. 
	- 이건 `Operator`에만 사용하는 것을 가정해서 만든 로직이었음.
- 지금처럼 `Boss`에도 사용하는 상황이 생겼기 떄문에 아예 **`DirectionSystem`이라는 계산만 담당하는 `static` 클래스의 `static` 메서드로 빼둔다.** 

#### 지식이 늘었다
>[!question]
>`Operator`에서 스킬 범위를 해쉬셋으로 갖고, `Get`으로 `IReadOnlyCollection`의 형태로 가져가게끔 구현했다면, 이를 사용하는 쪽에서는 어떻게 구현해야 할까? 어차피 변환하는 개념이 된다면 해쉬셋을 사용하는 의미가 있을까?

>[!answer]
>1. **메서드의 파라미터를 `IReadOnlyCollection`으로 구현**하면 됨. 
>2. 이 떄 `Contains` 같은 메서드를 사용하지 못한다고 뜨는데, 이건 **`System.Linq`** 를 가져오면 해결된다.
>3. **변환하는 개념이 아니다. `HashSet<T>`는 `IReadOnlyCollections<T>`의 인터페이스를 구현한 형태**이기 때문임.
>4. 해쉬셋의 경우 범위 내에 적이 있는지 여부를 빠르게 판단하기 좋다. **런타임 중에 보관하는 자료구조는 `HashSet`을 사용**하는 게 좋음.


#### 캐싱
- 오퍼레이터의 위치를 중심으로 스킬 범위가 펼쳐지는 경우가 있고, 적의 위치를 중심으로 스킬 범위가 펼쳐지는 경우가 있음. **전자의 경우는 범위 계산이 1번 이뤄졌다면 그 다음은 실행될 필요가 없다.** 항상 동일한 범위로 스킬이 나가기 때문임.
- 그래서 **스킬 사용 시 스킬의 중심 위치를 기억**하게 하는 방식으로 캐싱을 하겠음
	- 중심 위치가 똑같다면 범위를 계산할 필요가 없고, 다르다면 범위를 계산하는 방식
- 중심 위치를 찾고 범위를 계산하는 메서드는 **공용으로 빼둬야 한다.** 지금 `Boss`에 구현하려는 스킬이 이런 형태이기 때문임.

#### 지식이 늘었다 2
>[!note]
> **필요하지도 않을 유연성을 위해 코드를 복잡하게 만들지 말자 : YAGNI(You Ain't Gonna Need It) 원칙**

#### 지식이 늘었다 3
- 이런 상황을 가정해보자. **`ScriptableObject`를 상속받는 스킬의 로직**이다.
```cs
caster = op;

// 중심 위치
Vector2Int centerPos = GetCenterGridPos(caster);

// 스킬 범위 계산
if (centerPos == op.LastSkillCenter)
{
	HashSet<Vector2Int> skillRange = PositionCalculationSystem.CalculateRange(skillRangeOffset, centerPos, op.FacingDirection);
	op.SetCurrentSkillRange(skillRange);
}

// 중간에 계산된 skillRange를 바로 사용
VisualizeSkillRange(op, skillRange);
```
> `skillRange`에 관한 이슈를 다룰 건데, 여기선 `Scope` 때문에 어차피 쓰지 못하겠지만 저 조건문이 없더라도 발생할 수 있는 이슈라서 짚고 넘어간다.

>[!question]
>- 여기서 계산된 `skillRange`를 바로 사용할 수 있는데, 이 값을 바로 사용하는 게 맞을까? 아니면 `Operator`에 저장했으니 거기서 일일이 가져오는 게 맞을까?

>[!answer]
>- **SSoT(단일 진실 공급원)** 이라는 핵심 원칙이 있다.
>- `Operator`는 **상태의 소유자**다.
>	- 자신의 현재 스킬 범위가 무엇인가?를 대답할 수 있는 건 `Operator`만이 할 수 있어야 한다.
>- `ActiveSkill`은 **행위의 제공자**다.
>	- 어떻게 범위를 계산하고, 상태를 변경할 것인지에 대한 설계도를 제공한다.
>	- 설계도 개념이므로 자신이 상태를 갖고 있어선 안된다. 만약 **어떤 상태에 대한 정보가 필요하다면, 그 정보를 갖고 있는 객체에게서만 가져와야 한다.**

이런 원칙에 따라, 

```cs
// 중간에 계산된 skillRange를 바로 사용
VisualizeSkillRange(op, skillRange);
```

이 부분은

```cs
// 중간에 계산된 skillRange를 바로 사용
VisualizeSkillRange(op, op.GetCurrentSkillRange());
```

이렇게 사용하는 게 좋다는 의미다. 유지보수를 위해 이런 습관을 들이는 게 좋다. 

단, **여러 번 사용할 경우 일일이 가져오라는 의미는 아니다.** **하나의 메서드나 논리적인 작업 단위 내에서는 상태를 처음에 한 번 가져와서 지역 변수로 저장한 다음 계속 사용하면 된다.** 안 그러면 코드가 늘어지니까. 

#### 범위 관련해서 UnitSkill 구조 리팩토링 해보려다가 포기
- 뭐는 넣고 뭐는 빼고 이런 걸 생각하는 게 너무 복잡함;; 너무 삼천포로 빠지는 것 같아서 다시 원래 하던 거 구현한다.

#### UnitSkill 관련 - 설계도니까 Caster라는 스킬 사용 주체 필드도 여기서 갖고 있으면 안됨!!
- 제곧내. 관련 부분들은 전부 수정함. 생각보다 그걸 사용하는 부분이 많지는 않았다.


- 다시 스킬 구현으로 돌아감. 

## 보스 스킬 구현
- `ArcaneFieldSkill`의 설계를 참고하면 도움이 될 듯? 
- 이펙트랑 효과가 순차적으로 정해진 타이밍에 나와야 하는 것만 기억해두자. 

- `ArcaneFieldController`는 `FieldEffectController`의 상속이다. 그래서 얘를 상속받게 함.
- 대신 기존 구현이 되게 이상하게 되어 있었어서 이를 다듬으면서 수정해나간다. 
	- AI에게 툭 던져놓고 이해 못한 채로 코드를 작성하니까 고칠 때 잘 모르겠는 부분이 많아서, 가급적 직접 구현하는 방식으로 접근하고 있음. 
	- 이상하게 된 구현들이 있다. 일단 구현되게 하고 다듬으려면 시간이 꽤 걸릴 듯.
		- 오브젝트 풀링 코드를 작성해놓고 오브젝트를 파괴한다든가..
		- 필드 다 정의해놓고 초기화할 때 필드를 할당하지 않았다든가..

### 지식이 늘었다 4
>[!note]
>- 메서드 오버라이드 : 부모 클래스의 메서드를 `대체`하는 개념
>- 메서드 오버로딩 : 부모 클래스의 메서드를 `확장`하는 개념

- 예시
- `Parent`
```cs
public virtual void Initialize(param1 a, param2 b) 
{
	// 할당
}
```

- `Child`
```cs
public virtual void Initialize(param1 a, param2 b, param3 c) 
{
	base.Initialize(a, b);
	// 할당
}
```
> 자식 클래스의 `Initialize`는 부모 클래스의 메서드를 이용하되 파라미터가 추가됐다. 이렇게 활용하는 걸 `메서드 오버로딩`이라고 함

### ScriptableObject(설계도)에서 코루틴을 실행시키고 싶다면
>[!note]
>1. `MonoBehaviour`을 상속받는 요소에 코루틴을 실행시키는 메서드를 추가함.
>2. 설계도에서는 코루틴을 정의하고, 해당 요소에 코루틴을 전달만 하는 방식으로 구현

- 1번
```cs
public void ExecuteSkillSequence(IEnumerator skillCoroutine)
{
	StartCoroutine(skillCoroutine);
}
```

- 2번
```cs
public override void Activate(EnemyBoss caster)
{
	UnitEntity target = caster.CurrentTarget; // 수정 필요할 수 있음

	if (target == null) return;

	IEnumerator sequence = ActivateSequence(caster, target);
	caster.ExecuteSkillSequence(sequence);
}
```

> 이전까지는 `ScriptableObject`에서는 코루틴을 실행시킬 수 없다! 정도로만 알았다. 그게 맞지만, 코루틴을 정의할 수 없다는 건 아니다. **실행할 코루틴을 정의하는 것 자체는 가능하고, 그걸 실행할 `MonoBehaviour` 객체만 넘겨주면 됨.**

### 지식이 늘었다 5
>[!note]
>**"`ScriptableObject`에서 상태 필드를 가지면 안된다"의 정확한 의미는 스크립트 자체의 전역적인 상태 필드를 갖지 말라는 의미**이다. 메서드 등에서 일시적으로 동작하는 지역변수는 상관 없음.




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