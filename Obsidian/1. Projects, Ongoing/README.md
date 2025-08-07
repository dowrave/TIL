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
>[!todo]
>- 보스 이펙트 만들기
>- 기존 이펙트 수정하기
### 발견한 이슈

### 예정
- `1-3` 스테이지 구현 완료
	- 보스 구현

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