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

>[!WIP]
>`Projectile`을 충돌 기반으로 변경했는데, 판정이 정상적으로 동작했는데도 `nullReferenceException`이 나타나는 현상


>[!todo]
>- 보스 이펙트 만들기
>- 기존 이펙트 수정하기

>[!todo]
>남은 이펙트 정리해보기
>1. 투사체
>	- 화살
>	- 적 평타
>	- 캐스터 평타
>	- 아틸러리 평타
>		- 폭발
>2. 스킬
>	- 범위 관련 이펙트들 : 기존에도 파티클 시스템이었을 거라 만질 게 많진 않고 바운더리에서 올라오는 이펙트들의 존재감이 너무 강하기 때문에 이것만 줄여주면 될 듯.
>	- 메테오 스킬
>	- 슬래쉬 스킬
>	- 코스트...는 냅둘까?
>3. 보스
>	- 최종적으로 보스까지 구현하면 완료인 듯?

- **직후의 계획**
1. `Projectile_Arrow` : 화살만 없애고 연기나 파티클 등 월드에 남는 부분들을 어떻게 남길 것인가?
2. `Artillery`의 평타 + 폭발 이펙트
3. `Caster`의 평타 이펙트 

### 발견한 이슈

### 예정
- `1-3` 스테이지 구현 완료
	- 보스 구현

# 250812 - 짭명방
>[!done]
>- 주로 `Projectile` 스크립트의 수정과 리팩토링을 진행했다.
>1. `Projectile`이 사라진 후에 파티클 남기기
>2. `Projectile` 구조 - 콜라이더 기반으로 변경하기
>	- 단 지금의 "마지막으로 알려진 위치 기반" 로직은 유지해야 한다. 적이 죽었을 때 이미 생성된 투사체는 적의 마지막 위치까지 날아가서 사라지도록 함.
>3. 기타 수정 사항 
>- `Operator`에서 `Projectile`의 생성 위치 수정
>- `Muzzle` 이펙트 오브젝트 풀링 
>

## Projectile 사라진 후에 파티클 남기기
1. `mainParticle.Stop(false)` 적용해봄
	- 안 됨. 
	- 1번째 파라미터는 `withChildren`인데, 파티클의 실행을 중단시킬 때 자식 파티클의 실행도 중단시키는가? 인데, `false`로 지정해도 안 되는 듯.

2. 메인 파티클이 사라지는 시점에 부모-자식 관계를 끊은 다음 메인 파티클의 실행을 중단시키고 오브젝트 풀로 돌아가기 전에 다시 부모 - 자식 관계를 설정하기

...지금 보니까 스크립트 바꿔놓고 인스펙터에서 할당 다시 안해놨다. ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ 그리고 `IEnumerator`로 메서드 작성해놓고 `StartCoroutine`으로 안 돌리고 그냥 실행시켰다 ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ

- 아무튼 위 요소들을 다 수정하고 다시 테스트해봤다. 1번과 2번이 완전히 구분되는 개념은 아니다. 자식 파티클 중에도 실행되어야 하는 파티클이 있고, 실행되지 말아야 하는 파티클이 있기 때문이다.
- `TrailRenderer`는 크게 상관 없을 것 같다.  오브젝트가 꺼지지 않는 이상 `Trail`은 남는 듯.

### 수정 완료
- 이제 `Projectile`에는 `mainParticle` 외에도 남겨야 하는 파티클을 별도로 구현한다.
```cs
[Header("VFX")]
[SerializeField] private VisualEffect? vfxGraph;
[SerializeField] private ParticleSystem? mainParticle;
[SerializeField] private List<ParticleSystem> remainingParticles; // mainParticle이 사라지더라도 표시되는 파티클
```

- 그리고 파티클이 사라질 때 이 코루틴 메서드를 실행한다.
```cs
private IEnumerator ReturnToPoolAfterSeconds(float seconds)
{

	if (vfxGraph != null)
	{
		vfxGraph.Reinit(); 
	}
	else if (mainParticle != null)
	{
		// 부모-자식 관계를 일시적으로 끊어서 부모 파티클의 실행을 멈춰도 자식 파티클은 계속 재생되게 함
		foreach (var ps in remainingParticles)
		{
			ps.transform.parent = null;
			ps.Stop(true, ParticleSystemStopBehavior.StopEmitting); // 파티클의 추가 생성을 막음
			// ps.Play() // 굳이 필요 없어서 주석 처리해봄
		}

		// 남기지 않아도 되는 파티클들 모두 제거
		mainParticle.Stop(true, ParticleSystemStopBehavior.StopEmittingAndClear);
	}

	yield return new WaitForSeconds(seconds);

	// 다시 부모 파티클에 할당
	foreach (var ps in remainingParticles)
	{
		ps.transform.parent = mainParticle.transform;
	}

	ObjectPoolManager.Instance!.ReturnToPool(poolTag, gameObject);
}
```

- 추가로 이전처럼 투사체가 바로 사라지는 게 아니기 때문에 판정에 대한 플래그도 하나 설정해준다.
```cs
private bool isReachedTarget = false;
```
> `Initialize`에서 `false`로 초기화해주고, `OnReachTarget`이 동작하면 `true`로 바뀐다. `true`로 바뀌면 `OnReachTarget`이 실행되지 않음.

## Projectile 판정 콜라이더 기반으로 변경
- 단 기존의 `Update` 문에 있던 거리 기반 계산 로직은 남겨둔다. 투사체가 날아가는 중에 적이 사라지는 경우도 다뤄야 하기 때문이다.
```cs
private void OnTriggerEnter(Collider other)
{
	// 이미 목표에 도달했거나, 타겟이 없으면 무시
	if (isReachedTarget || target == null) return;

	// 충돌한 오브젝트가 내 타겟인지 확인
	BodyColliderController hitUnitCollider = other.GetComponent<BodyColliderController>();
	if (hitUnitCollider.ParentUnit == target)
	{
		// OnReachTarget() 대신 새로운 공용 함수 호출
		HandleHit(target.transform.position);
	}
}
```

## 기타 수정 사항
### `Projectile`의 생성 위치 수정
```cs
// Vector3 spawnPosition = transform.position + Vector3.up * 0.5f;
Vector3 spawnPosition = transform.position + transform.forward * 0.5f;
```
> 오퍼레이터가 보는 방향 쪽에서 생성. 기존엔 오퍼레이터보다 "위"에서 생성되었는데 화살의 경우 시각적으로 어색해보여서 수정함.

### Muzzle 이펙트가 생성되기만 하고 풀로 되돌아가지 않음
- 처음엔 귀찮아서 `Operator`에서 `Muzzle` 실행시킨 다음 1초 후에 풀로 되돌렸음
- 근데 `Muzzle` 개념은 `Enemy`에서도 쓸 거잖아? 그러니 `MuzzleVFXController`을 별도로 구현하는 게 낫다.
```cs
using UnityEngine;
using System.Collections;
using UnityEngine.VFX;

// 오브젝트 풀로 돌아가는 기능만 수행함
public class MuzzleVFXController : MonoBehaviour
{
    private string poolTag;

    [Header("Settings")]
    [SerializeField] private ParticleSystem ps;
    [SerializeField] private VisualEffect vfxGraph;
    [SerializeField] private float vfxLifetime = 1f;

    public void Initialize(string poolTag)
    {
        this.poolTag = poolTag;

        if (ps != null)
        {
            ps.Play(true);
        }
        else if (vfxGraph != null)
        {
            vfxGraph.Play();
        }

        StartCoroutine(WaitAndReturnToPool(vfxLifetime));
    }

    private IEnumerator WaitAndReturnToPool(float duration)
    {
        yield return new WaitForSeconds(duration);

        ObjectPoolManager.Instance.ReturnToPool(poolTag, gameObject);
    }
}
```
> 사실 오브젝트 풀에서 스폰 -> 컴포넌트 얻기 -> Initialize하기 여서 뭔가 부자연스러운 느낌은 있는 듯 한데, 대안도 없다.

### OnTriggerEnter가 null로 뜨는 이슈
```cs
    private void OnTriggerEnter(Collider other) {
    }
        // 충돌한 오브젝트가 내 타겟인지 확인
        BodyColliderController hitUnitCollider = other.GetComponent<BodyColliderController>();
```
> 위아래는 대충 생략. 저기서 `null` 예외가 발생 중.

근데 타격 판정 자체는 잘 발생했음.  **이거 마무리 못했음!!**


# 250811 - 짭명방


>[!done]
>- 이펙트
>	- `Projectile - Arrow`
>	- `Muzzle_v1`
>- `Projectile` 스크립트 수정

## Projectile - Arrow

### 김이 서리는 효과
- **어제 넣으려다가 못 넣었던 김이 서리는 효과**에 대해서 다시 구현해봤다. 
	- 레퍼런스들을 찾아봐도 너무 차가워서 김이 나는 듯한 효과는 종종 보이기도 하고, 어떻게 구현했는가도 보면서 판단해보기로 함.                                               
- 텍스쳐를 직접 그리진 않았고 `EffectTextureMaker`라는 웹 도구가 있어서 이걸 이용해서 `3x3` 텍스쳐 시트를 만들었다.
- 어제는 애니메이션으로 구현했는데 뭔가 어색해보였음. 
- 스프라이트 시트에서 랜덤한 텍스쳐 하나를 집어서 사용하는 방식이 더 있어보인다.
	- 텍스쳐 자체를 그렇게 밝게 가져가지 않기 때문에 쉐이더와 함께 사용하면 어느 정도 투명하게 나타난다. 
		- `Rotation`을 랜덤으로 주기
		- 서서히 사라지는 효과를 구현하기
		- 서서히 커지게 하기
		- 생성된 자리에서 크게 벗어나지 않기(이거는 케바케일 수 있음)
> 결과물은 아래에서 한꺼번에 정리

### 메쉬 다듬기?
- 어제는 메쉬를 썼기 때문에 `Opaque`를 적용해야 한다는 고정관념이 있었음
- 근데 꼭 `Opaque` 셰이더를 써야 할까? 단순 이펙트니까 `Transparent` 셰이더를 써도 된다. 사실은 이 쪽이 더 권장될 수도 있다. 화면에 잠깐 나타났다가 사라져야 하는 요소니까.
- 또, `Opaque` 셰이더를 쓰면 거의 필연적으로 `Lit`을 쓰게 되는데 이 경우 메쉬의 회전을 구현할 때 그림자가 지는 이슈가 있는데 이게 보기 좋은 요소는 아니다.

- 그러면 어제 구현한 `UV Map`을 보면서 텍스쳐를 새로 그리는 게 나을지도?
	- 시도해봤다. 화살의 꼬리 부분만 조금 더 보이는 식으로.
	- 근데 크게 눈에 띄지 않아서 **텍스쳐를 할당하지 않는 방식으로 바꿈.** 

![[Projectile_Arrow_v1 1.gif]]
> 일단 지금까지의 구현은 이렇다.

### Trail 구현하기

#### 텍스쳐 구현

>[!note]
>- 프로크리에이트에서 `Seamless` 이미지를 만들려면
> 1. 만들려는 이미지의 가로 폭 2배의 캔버스를 준비
 >2. 그리기 가이드 - 좌우 대칭을 켜고 그리기 도우미도 켠다
 >3. 오른쪽 캔버스에 원하는 텍스쳐를 구현(왼쪽에도 좌우 대칭으로 들어감)
 >4. 전체를 가로로 4분할한다고 하면, 2, 3번째 부분만 사용할 거다.
>	- 좌우대칭이므로 3번째 부분의 오른쪽 경계와 2번째 부분의 왼쪽 경계가 자연스럽게 만남.
>
>내 경우 4번을 진행한 다음 텍스쳐를 다듬었는데, 작업은 자르기 전의 상태에서 하는 게 더 나을 듯.  잘라낸 다음에 추가로 작업하는 경우 `빛 산란` 등에서 `Seamless`가 깨질 가능성이 있어보인다. 경계 부분을 특히 조심스럽게 다뤄야 함.


...사실 Krita를 쓰는게 더 간단할지도. 

- `Trail Renderer`는 조금 더 공부해볼 필요는 있겠다. 파티클 시스템 대비 제약이 좀 있어 보인다. 예를 들면 메쉬를 넣지 못하는 걸로 보임.

![[Projectile_Arrow_v1 2.gif]]

> 일단 `Projectile_Arrow`는 이 정도까지만 작업해두겠음. 너무 시간을 많이 먹는 감이 있다.

## Projectile 스크립트 수정
- 기존에 `VFXGraph`용으로만 작성되어 있었기 떄문에 여기서 수정함
- 지금 돌아보면 예전 스크립트는 너무 복잡하게 작성된 감이 있다. 

![[Pasted image 20250811170225.png]]
ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ 문워크로 간다

```cs
lastKnownPosition - transform.position
```
> 방향을 설정하는 부분인데, **벡터의 뺄셈 `B-A`는 `A`가 `B`를 보는 방향**이다. 기존엔 반대로 설정되어 있었음. 

이펙트의 크기가 크기 때문에 줄이고, Trail도 화살의 머리부분부터 나오게 해도 이펙트가 괜찮아 보여서 그렇게 수정.

- 추가 수정 중...
1. 타겟에 도달했을 때 잠깐의 유예기간을 준다. 남는 이펙트나 Trail 등이 바로 사라지는 게 어색해보임.
2. 1.과 함께 타겟에 도달했을 때는 파티클의 생성을 막는다.
```cs
	// 새로운 파티클의 생성을 막음
	if (vfxGraph != null)
	{
		vfxGraph.Stop();
	}
	else if (ps != null)
	{
		ps.Stop();
	}

	// 공격자가 사라졌거나, 풀이 제거 예정인 경우
	if (shouldDestroy)
	{
		Destroy(gameObject, WAIT_DISAPPEAR_TIME);
	}
	else
	{
		ReturnToPoolAfterSeconds(WAIT_DISAPPEAR_TIME);
	}
```

> [!warning]
> 다른 이펙트 구현하기 전에 이 부분은 어떤 원리로 동작시킬 수 있는지 알아봐야 할 것 같다. 
> - **투사체 파티클이 사라진 다음에 다른 파티클 시스템들을 어떻게 유지시켜야 하는가? 이를 위해 파티클 시스템은 어떻게 구성해야 하는가?**

## Muzzle_v1 구현
- 이 부분은 강의에서 구현한 것들이 있으니까 메쉬랑 텍스쳐는 그대로 사용하겠음

![[Muzzle01_v1.gif]]
> 뭔가 완전 만족스러운 그림은 아닌데 그대로 사용함
> 너무 화려할 필요도 없다. 평타 개념이라서..


## Muzzle 이펙트는 Operator에 반영해야 함
- `CombatVFXController`로 `Source To Target`으로 설정하면 될 듯.
- 원거리 공격할 때 `Muzzle` 이펙트가 있는지 체크하고 있다면 실행시켜야 함. 기존엔 `Muzzle` 이펙트 실행하는 부분을 구현하지 않았기 때문에 추가.
	- `OperatorData`에 프리팹을 받는 필드 하나 추가
	- 이 구현은 오브젝트 풀과 관련된 설정들도 모두 들어가야 함.

- 중간에 태그 적용 안되는 문제 있었음
	- 태그 이름에 들어가는 값을 `ProjectileTag`을 복붙해서 써놓고 `Projectile`을 `Muzzle`이라는 문자열로 바꾸지 않아서 발생했던 문제ㅋㅋ


- 실제로 게임에 넣어보니까 `Muzzle`에서 튀는 파티클도 좀 과한 느낌이 있어서 제거.
- `Muzzle`의 크기만 줄이고 오늘 내용은 마무리함.
- 일단 내일은 화살만 없애고 연기나 파티클 등 월드에 남는 부분들을 어떻게 남길 것인가부터 시작하면 되겠다. 그 다음은 `Artillery`의 평타 + 폭발 이펙트, `Caster`의 평타 이펙트 등을 수정하면 될 듯.

# 250810 - 짭명방

> [!doing]
> - `Projectile - Arrow` 다시 만드는 중

## Projectile - Arrow
![[Pasted image 20250810142749.png]]
- 메쉬를 새로 만들었다. 이런 느낌. 
- 가장 맨 위의 삼각뿔의 경우 삼각형들이 UV Map에 겹쳐진 모양이다. 이게 부자연스럽게 보일 수는 있겠는데... 일단 방법도 없어서 그냥 이렇게 구현해봄. 

- 메쉬 적용하고 머티리얼 보는데 좀 이상하다. 이 부분 다듬는 건 내일로 미룬다. 메쉬를 새로 만들어야 할 수도 있겠음. 아니면 단순히 UV Map의 문제거나.

- 투사체가 이동하는 과정에서 남기는 파티클도 구현한다. 기존의 `Beam` + `Snowflake`로 만듦. 
	- `Snowflake` 텍스쳐는 ChatGPT한테 따봤는데 배경이 투명한 `PNG` 파일이 아니어서 그거 참고해서 그냥 프로크리에이트에서 직접 만듦.

- `Cloud`를 이용해서 한기를 투사체의 경로에 남기는 방법도 구현해봤는데 그렇게 자연스러운 느낌은 아님. 
	- 차라리 화살촉 끝 부분에 메쉬를 하나 넣고 거기에 어떤 효과를 넣는 구현이 낫겠음

![[Projectile_Arrow_v1.gif]]

> 일단 오늘의 구현은 여기까지 한다. 지금 보이는 문제점으로는
> 1. 화살촉의 회전과 화살꼬리?의 회전이 비슷하게 돌지 않음. 화살 메쉬 자체가 정사각형이 아닌 것으로 보임.
> 2. 색상은 좀 고민되는 듯.. 또 뭐가 있었는데 까먹었다.


# 250808 - 짭명방

>[!done]
>- 이펙트 만들기
>	- 배치 이펙트
>	- `Hit_Lightning_v1` (캐스터 스킬 타격 이펙트)
>- 테스트 환경 수정

## 이펙트 제작 - 배치
- 원본 게임인 명방의 배치 이펙트를 보고 구현했다. 
	- 색깔이나 큰 `Circle`은 나름의 바리에이션.
![[Deploy_v1.gif]]

> - 게임 테스트하면서 알게 된 건데 본 게임에서 스킨이 있거나 한정 오퍼레이터는 배치 이펙트도 달랐다. 배치 이펙트가 워낙 짧은 순간에 지나가기 떄문에 몰랐던 내용인데, VFX 공부하면서 이런 몰랐던 것들이 새로 보이게 되는 게 재밌다. 

- 적용 과정) 
	- 위치 이슈 - 잘못된 거 아니었음. y 포지션이 `0.55`랑 `0.75`가 찍혀서 ?이었는데 타일 스케일까지 고려해서 들어간 값이라 오브젝트 위치만 잘 찾아가면 된다. 
	- 타일 위에 가로세로로 그려지는 파티클의 경우 오브젝트의 위치 자체가 타겟 대비 `-0.4y`이므로 `Shape`에서 추가로 만질 건 없음. 
	- 가로세로위아래 파티클들의 `lifetime`, `start speed`, `color over lifetime` 등을 추가로 만져서 조금 더 부드러운 느낌이 나게 수정했음.


## 테스트 환경 수정
- 위 작업을 진행하면서...
> 테스트 환경이랑 실제 환경이랑 다른가? `Shape - Module Gizmo` 위치 찍었을 때 이펙트가 나타나는 위치랑 기즈모가 표시되는 위치가 미묘하게 다르다. 

라는 생각이 들었다. 

이거는 예전에 각 오브젝트들 만지면서 오브젝트 하나에 컴포넌트를 몰아넣던 구조를 자식 오브젝트들에 배분하는 식의 수정을 거친 적이 있는데, 이 과정에서 VFX 테스트 환경의 오브젝트는 수정하지 않았기 때문이었다. 

- 수정 전)
	- 각 `Operator / Enemy` 오브젝트의 스케일이 `0.25, 0.25, 0.25`
- 수정 후)
	- 오브젝트의 스케일은 `1, 1, 1`
	- 대신 자식 오브젝트로 `Model`을 빼두고, 여기에서 메쉬 렌더링을 담당한다. 얘의 스케일이 `0.25, 0.25, 0.25`
	- 이펙트는 오브젝트를 기준으로 들어간다. 

여기서 `VFXTestScene`은 수정 전의 그것을 따랐고, 실제 스테이지에서는 수정 후를 따랐기 때문에 미묘한 차이가 발생했음.

**따라서** VFX 테스트 씬도 수정 후처럼 `Model`을 자식 오브젝트로 빼고 부모 오브젝트의 스케일을 `1, 1, 1`로 뒀음.

## 이펙트 제작 - Lightning
- `Caster`의 2번째 스킬의 타격 효과. 공격 모션은 따로 없고 공격 판정이 나올 때마다 해당 적의 위치에서 이펙트가 재생되는 방식이다.
- 기존에 `VFX` 그래프로 구현한 게 있는데, 파티클 시스템으로 만들어 봄.
	- 이 때는 아마 `Square`을 어떻게 쪼갰었나? 그랬는데 이번엔 텍스쳐 시트를 만들어서 작업해봄.

![[ForObsidian_Thunder01.png]]

- 이런 느낌으로 레퍼런스 참고해서 만들어봤는데, 번개가 사라지는 효과에 대해서는 조금 징그러운 느낌이 들었다. 
	- **번개가 서서히 사라지는 효과를 구현할 때는 번개 줄기가 얇아지는 효과를 넣어야 더 그럴 듯해 보이는 듯.** 
	- 저런 식으로 넣었더니 **번개가 사라진다기보다는 역병에 걸린 듯한(둥근 반점이 생기면서 없어지는 듯한) 효과에 가까워 보였다.**

일단은 다시 작업하는 것보다는 저걸 갖고 그대로 진행하겠음. `Color over Lifetime`을 적용해서 중간 이후로는 알파값을 0으로 향하게 하면 뒤의 사라지는 효과는 그렇게 두드러지지는 않는다.

![[Hit_Thunder_v1.gif]]

타격 이펙트이므로 이거보다 더 화려해질 필요는 없을 것 같은데
1. 바닥에 남는 검은 이펙트인 `Debris`
2. 튀어오르는 검은 파티클

이 2개는 별도의 텍스쳐 / 메쉬를 사용해보고 싶음. 

### Crack 텍스쳐 

![[ForObsidian_Crack01.png]]

![[ForObsidian_Crack01-1.png]]
> 위 버전을 사용해봤는데 너무 이상해서 아래 버전으로 수정함. 아래가 조금 더 그럴싸해보인다.
> `Crack` 이펙트의 특징은
> 1. 중앙에서 뻗어나가는 선이 굵음, 나머지 서로를 연결하는 선들은 굉장히 얇은 편
> 2. 잔가지를 많이 그릴 필요가 없으며, **굵은 잔가지가 많을수록 균열보다는 나무에 가까워보인다.**
> **선의 굵기가 포인트**인 듯.

쉐이더로 구현하면 이펙트가 가운데서 퍼져나가는 구현도 할 수 있을 듯 하지만 원하는 타이밍을 맞출 수 있는가가 좀 애매한 부분. 가운데서 퍼져나가는 효과가 `Beam`이랑 `Particle`로 나타나고 있기 떄문에 투명했다가 나타나는 것 만으로 생각보다 자연스럽게 나타나는 구현이 가능하다.

이게 정리하면 짧게 걸린 것 같지만 거의 2시간 쓴 듯. 텍스쳐 메이커에 비슷한 느낌의 구현이 보이지 않아서 직접 그렸다.

### Particle_Debris 메쉬로 구현
- 이전에 `Slash` 스킬을 구현할 때 잔해 메쉬를 쓴 기억이 있다. 그걸 활용할 계획.
- `Slash` 스킬을 의외로 파티클 시스템으로 구현했었다? 아무튼 `Debris` 메쉬가 있기 때문에 그걸 활용함.
- `Collision`을 켜서 조금 더 실감나게끔 해봄. 

### (일단) 최종본
![[Hit_Thunder_v2.gif]]
그래도 나름 만족스러운 느낌? 여전히 바닥은 살짝 아쉽지만 얘는 여기까지만 작업해둔다.
색깔을 어떻게 설정해야 할지 감이 안 잡힌다. 



# 250807 - 짭명방
>[!done]
> - 이펙트 만들기 
> 	- `Hit_Impact_v2` - 방향성 관련 추가 구현 
> 		- `Stretched Billboard`를 써보려고 했는데 제약이 있다. 정리해뒀음.
> 	- `Attack_Slash_v1` 
> 		- 기존 `VFX Graph`로 만들었던 메쉬랑 쉐이더가 있다. 그걸 활용해보려고 함.
> - `CombatVFXController`에 `Billboard`의 회전 케이스도 추가.
> 	- 회전 관련 : `FromToRotation` vs `LookRotation`
> - 버그 수정
> 	- 텍스쳐에 그린 적 없는 선이 나타나는 현상

>[!doing]
>- 배치 이펙트 구현 중 일단 중지
## Impact_v2 구현하기
- 어제 `Stretched Billboard`에 대해 알았으니까 적용해보려고 했음
	- 텍스쳐의 방향을 우->좌로 진행하는 걸 알게 됐으니까 회전도 쉽게 적용할 수 있지 않을까? 라는 생각으로 시도해봤다.
- 하지만 **제약이 있다.** 
	- 예를 들면 크기 변경. `Pivot`의 방향으로 **날아가는 파티클**을 구현하는 용도로 만든 거라서 멈춰 있는 파티클의 크기를 변경하면 텍스쳐의 크기가 축소되긴 하는데 이동하면서 축소되는 느낌으로 나타난다.
- 따라서 **`Mesh`를 쓰지 않으면서 방향성이 있는 파티클을 구현하려면 `Billboard + 코드`로 제어하는 게 지금 상황에서는 가장 좋은 방법**인 듯. `Start Rotation` 값을 만지는 방식으로.

- `vfx_Hit_Impact_v2`를 만들었다. 일단은 이 정도에서 만족.

- 게임 테스트) 내부 컴포넌트인 `Particles`과 달리 방향이 돌아가지 않는 문제가 있어서 그거 만지는 중 -> 아래 스크립트 구현으로 완료.

### 결과물
![[Hit_Smash_v2.gif]]

## Attack_Slash_v1 구현하기
- 메쉬는 기존에 만들어놨던 게 있다. UV Map만 어떻게 생겼는지 확인했음.
- 방향 설정에 관한 이슈가 있기는 한데, 기존 쉐이더도 보면 극좌표를 이용해서 정중앙 부분을 밝게, 바깥 부분을 어둡게 구현한 게 있음.
- 이를 참고해서 쉐이더에 극좌표 관련 기능을 추가해서 중앙 부분을 밝게 하고 바깥 부분을 어둡게 만드는 구현만 더하면 되지 않을까 싶다. `AddScroll`에 노이즈 관련 기능은 이미 붙어 있기 때문에 이를 기반으로 쉐이더 그래프를 하나 더 만들었다. 이를 `AddScrollRadial`이라고 표기.

- 이전 예제에서는 텍스쳐 없이 보로노이 노이즈 + 극좌표만을 이용했는데, 이번에 텍스쳐는 새로 그렸다. 

![[ForObsidian_SlashMesh01.png]]

왜 이렇게 만들었냐면 이전에 만들어놓은 메쉬의 UV Map이 아래처럼 생겼기 때문이다.
![[Pasted image 20250807174832.png]]
![[Pasted image 20250807174801.png]]
그래서 나머지 절반 부분의 텍스쳐는 있을 필요가 없고, 왼쪽 절반만 그리고 회전만 맞추면 될 것 같았음.

![[Attack_Slash_v1.gif]]
> `gif` 파일이라 느려 보이는데 실제로는 0.1초 정도만 나타나는 이펙트이므로 매우 빠르게 지나가며, 보로노이 노이즈가 적용된 상황이라 특정한 모양의 도형이 빙글빙글 도는 느낌만 나는 것도 아니다.

생각보다 더 만족스럽게 나왔다. 도형이 도는 문제나 이펙트가 너무 일정해보이는 문제는 `size`, `lifetime`, 텍스쳐의 노이즈 설정 등을 조절해서 해결했음.

휘두르는 방향의 바리에이션은 `3D Start Rotation`의 `Y`값을 바꿔서 설정할 수 있었다. `0 ~ 180`으로 설정.

여기서 문제는, 저 이펙트에 `Alpha Blended`가 적용된 메쉬를 하나 더 띄우려고 할 때인데, `Sub Emitter`에 `Rotation` 설정이 있긴 하지만 그건 이럴 때 쓰는 게 아니다. 

이것도 스크립트로 제어가 가능하다고 함. 근데 너무 품이 많이 드는 것 같아서 일단은 저대로 구현하겠음. 

- 원래는 파티클 튀는 것도 구현했는데, `Hit` 이펙트에 파티클이 이미 튀기 때문에 여기서는 저것과 알파 블렌디드만 적용하고 마무리.
- Ally, Enemy, Magical(DualBlade 스킬)에 대해 이펙트 구현 및 적용 테스트 완료. 


## CombatVFXController - Billboard의 회전도 반영되도록 수정
```cs
// 방향 계산 로직은 이전에 올렸으니 생략
// Billboard의 회전이 필요한 파티클 시스템은 인스펙터에서 등록할 수 있도록 했음.

if (direction != Vector3.zero)
{
	// 파티클 시스템 오브젝트의 회전
	// Quaternion objectRotation = Quaternion.FromToRotation(baseDirection, direction);
	Quaternion objectRotation = Quaternion.LookRotation(direction);
	transform.rotation = objectRotation;

	// 빌보드 파티클의 회전
	// 오브젝트의 Y축 회전값을 라디안으로 변환, startRotationZ값을 업데이트한다. 
	float billboardRotationInRadians = objectRotation.eulerAngles.y * Mathf.Deg2Rad;

	// 캐싱된 모든 모듈의 startRotation에 반영
	for (int i = 0; i < billboardParticles.Count; i++)
	{
		if (billboardParticles[i] != null)
		{
			var ps = billboardParticles[i];
			var mainModule = ps.main;
			mainModule.startRotationZ = new ParticleSystem.MinMaxCurve(billboardRotationInRadians);
		}
	}
}
```
> `FromToRotation`과 `LookRotation`의 차이 : [[유니티 - LookRotation vs FromToRotation]]
- 요약) 대부분의 경우 `LookRotation(forward, upwards)`을 쓴다고 생각하면 된다. 로컬 Z축을 `forward` 방향으로 향하게 하면서 `upward`을 가능한 `0, 1, 0`을 보게 하려는 동작이다.

## 기타 버그(?) 수정

### 그린 적 없는 선이 나타남
![[Pasted image 20250807132920.png]]
- 여기서 왼쪽에 나타난 선. 저 부분은 그린 적 없다.
- 해결 방법 : **텍스쳐의 `Wrap Mode - Repeat`가 켜져 있다면 `Clamp`로 바꾸기.** 
# 250806 - 짭명방

> [!done]
> - `CombatVFXController` 수정 
> 	- 이펙트의 방향 설정 가능
> - 기존 이펙트 수정 : 파티클 날아가는 속도, `Simluation Speed` 통일 등등
> - `Slash` 대신 사용할 `Smash` 이펙트 구현
> - 버그 수정 : `Enemy_Tanker`에 `Attack_Smash` 이펙트를 할당했는데 이전에 구현했던 `Slash`가 나가는 현상

> [!review]
> 하 오늘 이펙트 하나밖에 못했네;;

## 이펙트 작업 

### `Smash` 이펙트
![[ForObsidian_Smash01 (3).png]]
> 이런 텍스쳐를 만들었다. 베는 컨셉의 캐릭터가 아니라면 `Slash` 대신 사용할 거임

- 부모 오브젝트에 `Rotation`을 줘도 `Billboard` 설정이라서 텍스쳐가 회전하지는 않는다.
- 일단 `Start Rotation` 값을 설정해서 텍스쳐 자체를 회전시키는 건 가능한데, 게임과 연동해야 함.

- 중요) **텍스쳐와 이펙트의 방향**
>[!done]
>- 방향성이 있는 텍스쳐 + Billboard를 쓸 때는 텍스쳐의 머리 부분이 이미지의 왼쪽을 향하도록 그려야 한다. (완전히 새로운 파티클 시스템을 만들어서 테스트도 해 봄)
1. `Billboard`로 텍스쳐를 띄울 경우 항상 카메라를 보며, 회전은 적용되지 않음
	- 파티클 시스템 내의 `Rotation` 값으로만 설정할 수 있음
2. **`Stretched Billboard`**를 쓸 경우, 방향에 따라 파티클의 회전이 달라짐
	1) 파티클 시스템에서 설정하는 `Rotation` 관련 값들은 적용되지 않음
	2) **텍스쳐의 방향 : -X 방향, 즉 왼쪽을 보도록 그려야 파티클의 이동 방향과 텍스쳐의 머리 방향이 일치하는 걸 확인할 수 있었다.**
		- 이 부분에서 혼동이 많이 왔는데, ChatGPT도 Gemini도 모두 `+y` 방향으로 그리랬다가, `+x` 방향으로 그리랬다가... 테스트해보니까 아래처럼 해야 그나마 맞음.

![[Pasted image 20250806163228.png]]

- [[유니티 파티클 시스템 - 방향성이 있는 텍스쳐]]에도 정리.

#### 실제 적용 시 발생하는 문제
- 오퍼레이터 / 적의 Rotation 값이 적용되지 않은 듯? 이펙트가 계속 +z 방향으로 나간다.
- 이게 어떤 상황이냐면
```cs
case VFXRotationType.sourceToTarget:
	Vector3 hitDirection = (transform.position - attackSource.Position).normalized;
	if (hitDirection != Vector3.zero)
	{
		transform.rotation = Quaternion.LookRotation(hitDirection);
	}
	break;
```
> `transform.position = attackSource.Position` 이라서 방향 계산이 안 된다. 

- 아래 `CombatVFXController` 수정을 다시 수정했음.

- 이펙트의 `Z`값 설정을 만져서 적당히 튀어나가면서 `Hit_Impact`와 겹쳐지게끔 구현함. 



## 프로젝트 적용 시 문제 상황들
- `Hit` 이펙트를 구현하면서 파티클이 튀는 방향을 설정하는 구현 
	- 특정 이펙트는 방향이 적용되어야 하고 다른 이펙트는 기본 설정을 따르게 해야 함

### `CombatVFXController` 수정
- 타격, 피격 이펙트 등에 들어가는 VFX 실행 스크립트. 방향 설정과 재생을 담당한다.

```cs
public enum VFXRotationType
{
    None,
    targetToSource, // 타겟을 향함
    sourceToTarget, // 공격이 온 방향을 바라봄
}
```
을 인스펙터에서 설정할 수 있게 함. 이 값에 따라

```cs
private void PlayPS()
{
	if (ps != null)
	{
		Vector3 baseDirection = Vector3.forward; // 모든 이펙트는 +Z축으로 진행된다고 가정함
		switch (rotationType)
		{
			// 옵션 1) 피격자 -> 공격자 방향의 이펙트 진행
			case VFXRotationType.targetToSource:

				Vector3 directionToSource = (attackSource.Position - targetPosition).normalized;
				if (directionToSource != Vector3.zero)
				{
					Quaternion rotation = Quaternion.FromToRotation(baseDirection, directionToSource);
					transform.rotation = rotation;
				}                
				break;

			// 옵션 2) 공격자 -> 피격자 방향의 이펙트 진행
			case VFXRotationType.sourceToTarget:
				Vector3 directionToTarget = (targetPosition - attackSource.Position).normalized;
				if (directionToTarget != Vector3.zero)
				{
					Quaternion rotation = Quaternion.FromToRotation(baseDirection, directionToTarget);
					transform.rotation = rotation;
				}
				break;

			// 옵션 3) 별도 설정 필요 없음
			case VFXRotationType.None:
				break;
		}

		ps.Play(true); // true 시 모든 자식 이펙트까지 한꺼번에 재생함
	}
}
```
이러한 구현을 따르도록 한다. 

> 일단 대부분의 이펙트는 부모 오브젝트의 회전을 따르기 때문에 크게 상관 없지만, 타격 이펙트의 파티클 튀는 방향을 설정하기 위해서 집어넣은 옵션이다. 대부분 `None`으로 해도 무방함.

## 버그 수정
- `Enemy_Tanker`에 `Smash` 이펙트를 할당했는데 `Slash`가 나가는 현상이 있음
	- `Enemy_Tanker` 프리팹에 할당되는 데이터가 `Enemy_Default`였다.  엌ㅋㅋㅋ 
	- 모델링은 프리팹에 있고 눈에 보이지 않는 데이터만 다른 곳에서 가져오니까 헷갈릴 만 했다.
	- `Enemy_Tanker`로 수정하고 다시 테스트.
	- 수정 완료.

# 250805 - 짭명방

## 이펙트 작업

>[!done]
> - 기본 버프 지속 이펙트(공격, 방어, 힐)
> - 이펙트 적용 위치 스크립트 수정
> - 타격 이펙트 오브젝트 풀 태그 꼬이는 현상 수정
> - 타격 이펙트 : 힐 구현

### 구현된 이펙트들
- `Hit_Heal`
![[Hit_Heal.gif]]


- `Buff_Attack`
![[Buff_Attack.gif]]

- `Buff_Defense`
![[Buff_Defense.gif]]

- `Buff_Heal`
![[Buff_Heal.gif]]
### Rate over time으로 적은 수의 파티클을 다룰 때

[[파티클 시스템 - Rate over Time과 적은 수의 파티클]]에 아래 내용 정리.

>[!done]
>1. **`Burst`는 0초에 파티클을 바로 생기게 할 수 있다.** 
>2. **`Rate over Time`은 1보다 작은 소수값으로도 설정할 수 있다.** 
>	- 1개의 파티클을 유지하고자 한다면 `Start Lifetime * Rate over Time = 1`을 유지하도록 값을 설정하면 됨
>3. **`Rate over Time`으로 적은 수의 파티클을** 다룰 경우, 0초에 파티클이 나타나지 않을 수 있다.
>	- **`Prewarm` 옵션으로 해결**한다. (`Looping`이 켜져 있어야 함)
>		- 단 이 경우 파티클 시스템을 멈추고 `Playback Time`을 다시 테스트할 때 안 보이는 현상이 있다.
>	- `Burst`로 구현해도 가능은 한데 파티클이 겹치는 문제가 있음. 이건 왜인지는 모르겠다.

- 공통 버프 작업을 다시 진행 중인데, `Beam` 이미지가 커졌다가 작아졌다가 하는 식으로 바닥 파티클을 구현하는 중.

![[Pasted image 20250805134300.png]]

여기서 `Size over Lifetime`, `Rate over Time`, `Lifetime` 관련 이슈가 있었다.
- `Rate over Time` : 1초당 생성하는 파티클의 갯수
- `Size over Lifetime` : `Lifetime` 동안 파티클의 크기 변화
- `Lifetime` : 개별 파티클의 지속 시간

`Rate over Time`이 **1초당**에 해당하는 개념이므로, 만약 `Lifetime`이 1초보다 길다면 바닥의 파티클이 겹치는 문제가 발생했다. 그렇다고 `Lifetime`을 줄이자니 `Size over Lifetime`의 변화 속도가 너무 빨라지는 문제도 있다.

제미나이님께 물어보니, **`Rate over Time`은 소수로도 설정할 수 있다.** 예를 들어 `Start Lifetime = 2s`, `Rate over Time = 0.5`로 설정한다면 파티클은 2초에 한 번 생긴다. 

이렇게 설정하고 테스트해보면 파티클 시스템을 재생해도 0 ~ 2초까지는 이펙트가 보이지 않는다.

이 문제는 `Prewarm` 옵션을 켜는 것으로 해결할 수 있다.
- 단순히 `Start Delay`가 있을 때 이펙트가 자연스럽게 실행되도록 할 때만 쓰는 기능인 줄 알았는데, 이럴 때도 쓸 수 있다.

여기서 파티클 시스템의 `Playback Time`을 다시 조작했을 때, **0 ~ 2초 사이의 파티클이 사라지는 현상**이 있다. 이건 **`Playback Time`을 조작할 때 `Prewarm`에 의한 계산은 고려하지 않기 때문이다.**

### 이펙트 풀 공유 문제 수정
- 태그가 계속 꼬여 있는 것 같다.
- 이게 정확히는 이펙트가 잘 재생되다가 갑자기 어느 순간부터 다른 이펙트가 재생된다는 느낌임
- 어제 수정했던 내용인데, 저장 내용이 반영되지 않았던 것으로 보임.
	- 어제 작업 중에 컴퓨터를 절전으로 만들어뒀다가 나중에 다시 켜려니까 그대로 맛이 갔던 현상이 있었다. 
	- 요점은 `UnitEntity`가 갖는 `hitEffectTag`와 `attackSource`가 갖는 `hitEffectTag`을 구분해서 넣어줘야 한다는 거였음. 또 발생하면 다시 언급함.


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