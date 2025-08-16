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

>[!todo]
>- 보스 이펙트 만들기
>- 기존 이펙트 수정하기

>[!todo]
>남은 이펙트 정리해보기
>1. 투사체
>	- 아틸러리 타격 이펙트 ) 폭발
>2. 스킬
>	- 범위 관련 이펙트들 : 기존에도 파티클 시스템이었을 거라 만질 게 많진 않고 바운더리에서 올라오는 이펙트들의 존재감이 너무 강하기 때문에 이것만 줄여주면 될 듯.
>	- 메테오 스킬
>	- 슬래쉬 스킬
>	- 코스트...는 냅둘까?
>3. 보스
>	- 최종적으로 보스까지 구현하면 완료인 듯?

### 발견한 이슈

### 예정
- `1-3` 스테이지 구현 완료
	- 보스 구현

# 250816 - 짭명방

> [!done]
> - 이펙트 구현
> 	- `VFX_Projectile_Missile_v1`
> - 기타 이슈 수정


## 이펙트 구현 - Projectile_Missile
- `Artillery`의 평타에 사용될 이펙트.

### Mesh
![[Pasted image 20250816135720.png]]
Gemini에게 Mark Seam을 어떻게 해야 할지를 물어봤다.

- `MarkSeam`은 절제선을 어디로 설정하는가라고 생각하면 쉽다고 함. 해당 엣지들에 가위질을 했을 때 3D 도형을 2D로 펼칠 수 있는가를 생각해보면 된다.
> 물론 말은 쉽다. ㅋㅋㅋㅋㅋ 처음엔 반구와 원기둥 사이의 경계면에만 넣었는데, 추가로 원기둥의 옆면 + 바닥면까지도 포함해야 했음.

- `Mark Seam`을 설정했다면 `Unwrap`을 할 때 `Smart UV Project` 대신 기존에 사용하던 기본적인 `UnWrap`으로 풀어내면 된다. 
> 물론 이 프로젝트에서는 저 메쉬는 단색으로 구현할 거라서 UV Map을 추출해서 사용할 일이 없겠지만, 만들어놨다면 UV Map과 함께 저장해두는 게 좋겠음.

- 원기둥 옆면에 뭔가 노랑/검정 같은 이미지 추가해서 조금 더 눈에 띄게 하고 싶어서 보는데 어디가 윗면이지? ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ
	- UV Map 기준 오른쪽이 윗면이다. ㅋㅋㅋㅋㅋㅋㅋ

![[Missile01.png]]

> 이런 이미지를 만들었음.
![[Pasted image 20250816150619.png]]
 이런 느낌으로 들어간다. 완전 처음 해본 것 치고는 괜찮은 듯.

이걸 추가하니까 머리 부분에 이펙트도 하나 추가해보고 싶다. 기존에 `ProjectileMesh`가 있었고, 이것의 Uv Map이 가운데로 모이는 모양인 걸 참고해서 가운데 부분에만 이미지가 있는 텍스쳐를 하나 쪄봄. `EffectTextureMaker`를 사용해서 만들었다.

![[Projectile_Missile_v1.gif]]

- 추가로 전체적인 파티클의 사이즈들을 줄임
- 여기서 타격 이펙트도 다시 만져볼 계획.

## 이펙트 구현 - Hit_Explosion
- 위 투사체 gif 파일에서 나오는 폭발 이펙트 대신, 새로운 폭발 이펙트를 만들어본다.
- 이런 게임에서 쓰기 좋은 이펙트는 아닌 것 같지만, 투명하지만 배경이 왜곡되는 효과를 구현해보고 싶어서 그걸 테스트해보려고 한다. 사실 그거 외에도 할 건 많지만.
	- 일단 튜토리얼 따라가보긴 했는데 셰이더에 적용이 되지 않는다? 보류.
- 일단 폭발 자체와 투명한 바닥 충격파 부분을 제외한, 현재 구현할 수 있어 보이는 것들은 구현해두겠음.

![[Hit_Explode_v1.gif]]
> 일단은 이 정도로 구현해뒀고, 여기에 추가할 사항으로는 아래 2가지가 있음.
> 1. 입체적인 폭발 이펙트
> 2. 바닥의 ShockWave는 투명하며 배경을 왜곡시키게끔



## 추가 수정
- `Projectile` 생성 위치 : `Operator`가 보는 방향 + 0.5에서 0.25로 줄임. 위 움짤을 보면 너무 앞에서 움직이는 느낌이 있다. 


# 250815 - 짭명방
>[!done]
> - 이펙트 구현
> 	- `vfx_Projectile_Molotov_v1`
> - 이슈 수정
> 	- `Enemy_Ranged`가 새롭게 배치된 `Operator`을 공격하지 않는 현상
> 	- 기타 이슈들

뭔가 적당한 타협이 됐다. 

![[Projectile_Molotov_v1.gif]]

- `Molotov`를 `Mesh`로 구현하는 게 더 나아보이긴 한다..
	- 메쉬로 구현하면 메쉬의 특정 위치에 `Fire` 파티클을 붙일 수 있기 때문인데, 지금은 일단 파티클시스템으로 메쉬 생성하고 이를 감싸는 이펙트들을 구현하는 방식이 됐음.
	- 불 위치가 조금 어색해서 메쉬 대비 +z 쪽으로 위치를 수정했다.

- `Trail` 자체는 `Gabriel Aguiar` 센세의 튜토리얼을 따라 셰이더를 만들고 텍스쳐도 이런 모양을 만들어서 넣었는데, 막상  유니티에서는 모양은 잡히지만 내용물에 알파값이 잡히지 않고 전부 다 흰색으로 인식하는 현상이 있다.  왜 그런 건지 모르겠음.. 지금 모양도 물론 만족스럽긴 하지만, 이 이슈는 한 번 생각해볼 필요는 있어 보인다.
![[ForObsidian_Trail03.png]]
> 이 부분은 궁금해서 더 자세히 테스트를 해봤는데, 아무래도 빛산란에 의한 효과인 듯?
> 애초에 **텍스쳐 자체의 알파값, 혹은 밝기가 그렇게 높으면 안되는 것 같음. 대부분 50 언더로 해야 할 듯?** 낮은 알파값들로 텍스쳐를 만들고, 빛산란으로 더 많이 겹쳐진 부분을 더 밝게 해야 할 듯.
> 강의에서도 여러 번 나왔던 내용이지만, 알파값을 그렇게 높게 설정하지 않고 작업했다. 높아봐야 30% 정도? 
> 뭐가 어디까지 맞고 아닌지 모르겠다!!!!! 알파를 겁나게 낮춰서 넣어도 뭔가뭔가 애매한 느낌이다. 이런 걸 보면 겹쳐 그리는 게 중요한 건가?

## Enemy_Ranged가 공격 몇 번 후 공격하지 않는 현상
- 딱 5번 하고 공격하지 않는다. 왜?
- 이거는 다른 현상일 수도 있다 : **`Operator`가 `Enemy`의 공격 범위 내에 배치됐을 때, `ColliderEnter`가 동작하지 않는 현상일 수 있음.**
- 후자가 맞아보임. 
- 이거 예전에 구현했던 거 같은데 아니었나보다?
	- **새롭게 배치된 유닛 시점에서만 처리했고, 기존 유닛 입장에선 처리하지 않은 듯 하다.**

- `DeployedUnitEntity`의 배치에 관한 static event를 추가하고, 이를 `EnemyAttackRangeController`에서 감지한다.
```cs
private void HandleNewlyDeployed(DeployableUnitEntity target)
{
	if (owner == null || !enabled) return;

	BodyColliderController targetColliderController = target.GetComponent<BodyColliderController>();
	if (targetColliderController != null)
	{
		Collider targetCollider = targetColliderController.BodyCollider;

		if (targetCollider != null)
		{
			// 두 콜라이더가 겹치는지 검사 수행
			bool isOverlapping = Physics.ComputePenetration(
				attackRangeCollider, transform.position, transform.rotation, // 이 콜라이더의 정보
				targetCollider, targetCollider.transform.position, targetCollider.transform.rotation, // 타겟 콜라이더의 정보
				out Vector3 direction, out float dist // 출력 변수 : 사용하지 않음
			);

			if (isOverlapping)
			{
				Debug.Log("새롭게 배치된 유닛이 Enemy의 콜라이더 내에 있어서 공격 대상으로 추가");
				owner.OnTargetEnteredRange(target);
			}
		}
	}
}
```

위의 과정을 구현하면서 이런 생각이 들었다 : `Operator` 입장에서 `enemy`가 생겼을 때도 비슷한 현상이 일어나지 않을까?

- 근데 `BodyColliderController`라는 모든 엔티티에게 있는 요소는 자신이 활성화된 시점에 자신과 겹친 콜라이더를 체크하는 로직을 넣어놨다. 
- `Enemy`의 경우 이 로직에 의해 자신이 생성된 시점에 자신과 겹친 타일과의 충돌 트리거가 발동하고, 해당 타일을 공격범위로 하는 오퍼레이터에 의해 공격 대상으로 지정될 수 있다. 이것과는 다른 로직임.
- 결론은 기억력 이슈지만, 충분히 헷갈릴 수 있는 지점이었다고 생각함. 해결 완료.


## 기타 이슈
1. `Molotov` 투사체의 판정이 발생하지 않음
	- `Rigid Body`, `Sphere Collider`을 넣지 않아서 발생한 문제

2. `Projectile_Arrow`의 `Trail`이 1번 생성된 이후 사라지는 현상
	- **`Trail Renderer`의 `AutoDestruct`가 켜져 있으면 파괴**되기 때문이다. **오브젝트 풀을 이용할 거라면 저 기능은 사용하지 않는 게 좋겠다.**
3. 저지 로직도 이상하게 동작하는 듯?
	- 보통 한 군데서 버그가 생기면 다른 곳에서 이어질 가능성이 있다. 2번부터 체크해봄.



# 250814 - 짭명방

>[!done]
> - 이펙트 구현
> `vfx_hit_fire_v1`

>[!WIP]
>- `Projectile_Enemy` 
>	- 메쉬는 만들었다. 천을 애써서 넣었지만 인게임에선 잘 안 보이는 게 함정.
>	- 막힌 지점은 투사체의 불 이펙트. 텍스쳐 시트의 애니메이션 기능을 이용하자니 뭔가 어색하고 불 파티클도 뭔가 애매한 느낌이고. 좀 고민되는 지점.


## 시작 전
- `Enemy` 투사체, `Artillery` 투사체 및 폭발 이펙트를 구현하게 될 듯
- 전부 화염 관련 이펙트들로 구현할 듯. 
	- `Enemy`의 경우 `Orb`의 팔레트 스왑만을 써도 되겠지만 명방 초반의 리유니온 원거리 적의 화염병 같은 게 있었다. 그런 걸 구현해보고 싶음.

## VFX_Projectile_Molotov

### 화염병 메쉬

![[Pasted image 20250814140444.png]]

유리병은 쉽게 만들 수 있는데, 천을 어떻게 얹느냐가 궁금했던 부분. 검색해보니 생각보다 쉬웠다. 물론 말이 생각보다 쉬웠다는 거고 처음이라서 1시간 넘게 걸림.

1. `Plane`을 만듦. 물리엔진과 애니메이션을 이용해 구부릴 것이므로 `Subdivide`를 많이 넣어준다.
2. 화염병 안에 `Plane`이 들어가게끔 배치시킨다. `Plane`의 가장 아랫 부분을 고정한다.
	- 고정 방법 
		- 가장 아래 버텍스들을 잡고 `Data - Vertex Groups`에 새로운 그룹 `Pin`을 추가한 후 `Assign`
3. `Physics - Cloth` 을 켠다. 2번에서 설정한 `Vertex Group`은 `Shape`의 `Pin Group`으로 넣는다. 
4. 이전에 만들었던 유리병의 `Collider`를 켠다. 
5. `Object Mode`로 놓고 애니메이션을 재생시킨다. 원하는 장면이 나온 시점에서 프레임을 멈춘다.
6. 이전에 적용했던 `Modifier` 들을 적용시킨다. 

다만 `UV Map`이 문제이긴 한데, 지금처럼 디테일하게 들어가지 않고 모형 정도만 필요한 경우에는 크게 상관이 없다. 
더 정교한 작업이 필요하다면 이전에 블렌더 강의에서 봤던 `Mark Seam`을 이용해서 특정 부분들을 자르는 방식을 이용하지만(실사 소주병을 따라 만든다든가..) 지금 같이 메쉬만 필요하고 텍스쳐는 단색으로 처리하려고 하는 경우에는 `Smart UV Project`가 유용할 수 있음.

![[Pasted image 20250814142220.png]]

### 불 효과
- 이게 생각보다 어렵네... 자연스러운 느낌이 안 나서 애 먹는 중.
- 병 자체의 화염 효과는 살짝 미루고 일단 다른 것들부터 구현함
- **으아아 너무 어렵다 일단 보류**

## vfx_hit_fire_v1
- 불 이펙트가 타격했을 때의 효과를 구현해본다. `hit_ice`를 응용해서 만들어봄.
![[Hit_Fire_v1.gif]]
구름 텍스쳐는 아주 요긴하게 써먹고 있다.

> 추가 수정) 
> 1) `Smoke` 파티클들의 경우 y포지션은 `-0.25` 정도로 잡는 게 적당하다. `-0.3` 이하로 내려가면 피격당한 객체의 높이가 `0.3`인 애가 있어서 타일이랑 겹쳐보이는 문제가 발생함
> - 사실 이거는 월드의 오브젝트와 이펙트가 겹쳤을 때 **이펙트가 먼저 보이면 되긴 한다.** `Opaque`랑 `Transparent` 간의 순서로 이미 적용돼야 하는 걸로 알고 있는데 안 되서 모르겠음.
> 2) `Start Delay`를 넣고 싶다면 조심스럽게 접근해야 한다. 적은 움직이고 있기 때문에 이미 지나갔는데 연기가 바닥에서 솟아오르는 건 어색해보일 수 있음. 그래서 일단 `Delay`만 `0.3` 정도로 줄여놨음.

만족스럽진 않은데 일단 여기까지. 



# 250813 - 짭명방

>[!done]
>- 이펙트 구현
>	- `Arrow_Hit`
>	- `Caster` 평타 : `vfx_Projectile_Orb_v1` 작업
>- 이슈 / 수정 사항
>	- `Projectile`의 `Collider` 관련 이슈 수정
>	- `Projectile`에 `needToRotate` 필드 추가 : 회전이 필요한 파티클만 체크

## 이펙트 - VFX_Hit_Ice
- `Sniper`의 평타가 들어갈 때마다 실행될 이펙트.
![[Hit_Icy_v1.gif]]
- 텍스쳐 
	- `EffectTextureMaker`의 `Cross`
	- 가운데가 비어 있는 연기 효과는 `Corona`
	- 연기는 기존에 `Arrow`를 구현하는 과정에서 썼던 걸 다시 썼음.

일단은 이런 느낌 정도로 마무리.
인게임 테스트도 꽤 자연스럽다. 

>[!note]
>- `Stretched Billboard` 설정의 파티클이 날아가다가 궤적이 짧아지는 효과 같은 걸 구현하는 방법
>	- `Velocity over Lifetime`의 `Speed Modifier`을 `Curve`로 구현

## 이펙트 - VFX_Projectile_Orb
- `Sphere` 메쉬를 쓰지 않고 구현할 수 있는가?를 생각해봤는데 솔직히 구현하는 게 훨씬 나을 듯
- 블렌더에서 만드는 건 쉽다. `UV Sphere`를 만들면 `UV Map`도 깔끔하게 뽑히는 구가 나온다. 
	- 대신 이 점은 유의해야 함 : 프로젝트에 사용하는 구


- 메쉬에 사용할 쉐이더 구현
- 어떻게 구, 오브 효과를 위한 쉐이더를 만들 수 있을까?
- 이건 [선생님의 강의](https://www.artstation.com/artwork/NyKnwJ)가 있으니 이걸 따라가본다. 
	- 강의에서는 텍스쳐를 별도로 `MaterialMaker`를 이용해서 만들었는데, `Voronoi` 의 경우 유니티 쉐이더 그래프에서도 지원하는 기능이다. `Tiling And Offset`의 `Tiling`을 줄이면 늘어진 보로노이 효과를 넣을 수 있음.

- 일단 `Orb` 기초를 놓고, `Arcane` 버전을 따로 구현했다. 색놀이 수준이지만.

![[Caster_Attack_v1.gif]]

> UI에 이상한 밑줄 그이는 거는 GIF 파일에서만 보이는 현상이다. 인게임에선 안 나타남.

1. 저기서 `Hit_Lightning` 부분은 `Hit_Arcane`을 따로 구현할지 아니면 그냥 쓸지 고민은 되는 듯?


## 이펙트 - VFX_Hit_Arcane
- 임팩트 프레임같이 검은색이 가운데에 오고 흰색이 밖에 오는 구성으로 겹칠 수 없나? `Alpha Blended`로는 안되는 것 같다. 그냥 `Opaque`로 쳐야 하나?

- 일단은 그냥 진행.
- 위에서 생각한 거 대신 `Particle System` 자체에 있는 `Trail`을 이용해서 모양을 만드는 방식으로 만들었다. 뭔가 이거저거 공은 들어갔고 분명 새로운 걸 시도했는데, 했던 거 재탕하는 느낌이라 석연치 않은 느낌도 조금 있다. 그래도 이렇게 조금씩 늘어가는 거 아닐까?

![[Hit_Arcane_v1.gif]]
> 가운데 모양은 파티클 시스템의 `Trail`과 `Velocity over Lifetime`의 `Orbital`을 이용했다. `World`좌표,  `Orbital y`로 구현함. 파티클 자체는 투명하며, `Trail` 자체가 색을 갖는 식이다.
> 노란색을 넣은 근거는 `Orb`의 흰색을 만드는 요소가 보라색 + 노란색이기 때문이다. 거의 보이지 않지만, 그래서 궤적 파티클도 노란색으로 남겨뒀음.

일단은 이 정도로 만족하고 다음으로 넘어가면 될 듯. **오늘은 여기까지!**



## 이슈 / 수정사항

### 어제의 이슈

> [!note]
> - `Projectile`을 `Collider` 기반으로 변경했는데, `nullReferenceException`이 뜨는 현상

```cs
	// 충돌한 오브젝트가 내 타겟인지 확인
	BodyColliderController hitUnitCollider = other.GetComponent<BodyColliderController>();
	
>   if (hitUnitCollider != null && // 다른 콜라이더는 거름
		hitUnitCollider.ParentUnit == target)
	{
		// OnReachTarget() 대신 새로운 공용 함수 호출
		HandleHit(target.transform.position);
	}
```
> `hitUnitCollider`의 `null`체크 부분이 들어가지 않아서 발생했던 이슈였다.

### 의문점
- 근데 **오류 로그에서 가리키는 지점은 두 줄 위인 hitUnitCollider를 정의하는 부분**이었는데 왜 저걸로 해결되는지는 모르겠음. 아래의 `hitUnitCollider == target`에서 오류가 발생한다고 적어놔야 맞지 않나?

- AI에게 물어보니 크게 2가지로 정리되는데 
	1. **버그를 더 쉽게 찾아주기 위해 `hitUnitCollider`를 선언한 부분을 찍어준다**는 얘기.
		- 만약 오류 발생 지점을 찍는다면 `hitUnitCollider`가 `null`인지 `target`이 `null`인지 이중으로 찾아봐야 한다는 얘기.
	2. **컴파일러에서는 `hitUnitCollider` 선언 -> 사용하는 부분을 하나의 흐름으로 묶어서 처리한다**는 것. 그래서 흐름의 시작점인 선언 부분을 오류 로그로 찍어준다는 것이다.

뭔가 시원하게 알았다는 느낌은 아니긴 하다. 그냥 이렇게 찍어두고 넘어감. 

## 이슈 발견
- `Hit_Lightning` - 피격체의 `Rotation`에 의해 번개가 떨어지는 지점이 변하는 현상이 있음.
	- 근데 인게임에서는 원하는 위치에 잘 나타나서 굳이 수정 안 해도 될 듯?

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