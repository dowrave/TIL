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

# 250819 - 짭명방

>[!done]
> - 기타 이슈 수정
> 	- `UnitEntity`의 콜라이더 관련 : `0.5, 1, 0.5`로 수정
> 	- 사망 애니메이션 : 자식 메쉬 렌더러가 여러 개일 경우도 고려
> 	- `Operator` 배치 시에 공격 범위 내에 적이 있는데 공격하지 않는 현상 수정
> 	- `Caster` 레벨당 공격력 수정(9.3(??) -> 2.3)

## VFX_Projectile_Medic 구현 시작
- `Medic`도 투사체가 있었다.




## 기타 이슈 수정

### UnitEntity의 콜라이더 관련
- `Projectile`이 목표물에 명중할 때, 생각보다 좀 일찍 사라지는 느낌이 있음. `Trail`이 사라지는 지점이 목표물에 닿지 않은 지점으로 보이기 때문에 눈에 띈다.
- 특히 `Operator`에게 공격이 들어갈 때가 그렇다. 이는 `Operator`의 콜라이더 사이즈가 `0.8, 1, 0.8`이기 때문임
	- 참고로 `Enemy`는 `.2, 1, .2`를 갖는다
- 일단 `Enemy`의 콜라이더 사이즈가 대체로 캡슐의 사이즈와 비슷해 보인다. 그래서 `Enemy`한테 공격이 들어갈 때는 그렇게 어색하게 보이지 않음.
- 반면 `Operator`에게 공격이 들어갈 때는 메쉬에 닿는 것처럼 보이지 않는데 타격 판정이 발생함

- `Enemy`가 2개의 타일에 걸치는 효과도 고려해야 함.

- 일단 **지금까지 구현한 모든 프리팹에 대해서는 `0.5, 1, 0.5`에서 출발해본다.**
	- 오퍼레이터가 타일 B에 있고, 적이 타일 A에서 B로 넘어갈 때 타일과 타일의 경계에서 저지당하는 느낌이 됨
	- 여전히 캡슐에 비해 피격 범위가 넓어보이긴 할 거임. 
	- 이거 수정하다가 아래 렌더러 이슈 수정하고 왔다. 갑자기 생각나서 ㅎㅎ;
	- 일단 `0.5`로 수정했을 때 대체로 괜찮아 보인다. **오른쪽에서 공격이 들어왔을 때만 살짝 어색해보이는 수준?** 
	- 다른 방법으로 생각난 것들이 아래의 요소들임.
		- 콜라이더를 하나 더 구현해서 판정
		- 기존 구현대로 거리 기반으로 타격 여부 계산하기 -> `Update`를 계속 돌려야 하므로 성능적인 이슈가 있을 수 있음
		- **일단 대체로 괜찮은 느낌이 나오고 있기 때문에 지금 상태를 유지해본다.** 

### 사망 애니메이션 여러 메쉬 렌더러에 대해 동시 재생
- `Enemy_Tanker` 처럼 여러 개의 메쉬 렌더러를 갖는 오브젝트의 사망 시, 모든 렌더러에 대해 사망 애니메이션이 재생되게 해야 함

- 기존 코드
```cs
Renderer renderer = GetComponentInChildren<Renderer>();

// 동일한 머티리얼을 사용하는 모든 객체에 적용되는 걸 막고자 머티리얼 인스턴스를 만들고 진행한다.
if (renderer != null)
{
	Material materialInstance = new Material(renderer.material);
	renderer.material = materialInstance;

	SetMaterialToTransparent(materialInstance);

	// DOTween 사용하여 검정으로 변한 뒤 투명해지는 애니메이션 적용
	// materialInstance.DOColor(Color.black, 0f);
	materialInstance.DOFade(0f, 0.2f).OnComplete(() =>
	{
		OnDeathAnimationCompleted?.Invoke(this); // 사망할 것임을 알리는 이벤트
		Destroy(materialInstance); // 메모리 누수 방지
		Destroy(gameObject);
	});
}
```
> `GetComponentInChildren`으로 구현되므로 1개만 찾는다. 여러 개일 경우에 문제가 됨.

- 수정 방향
	- 우선, 기존의 구현은 애니메이션이 끝나면 머티리얼 인스턴스와 게임 오브젝트를 파괴시키는 방식이었다.
	- **어떻게 `OnComplete`를 여러 개에 대해 실행시킬 수 있을까?** 가 핵심이겠음.
	- **`DOTween`에서는 `Sequence()`라는 기능을 제공**한다. 함수가 호출된 시점에 **애니메이션(트윈)들을 `Join`으로 등록**해놓으면 시스템이 다음 프레임에 시퀀스를 감지해서 실행시키는 방식임.
```cs
// 시퀀스로 만들어 여러 개의 애니메이션을 하나의 그룹으로 묶어서 관리한다
Sequence deathSequence = DOTween.Sequence();

List<Material> materialInstances = new List<Material>();

foreach (Renderer renderer in renderers)
{
	// 1. 머티리얼 인스턴스로 만들어 동일한 머티리얼을 사용하는 다른 객체에 영향을 주지 않게 한다
	Material materialInstance = new Material(renderer.material);
	renderer.material = materialInstance;
	materialInstances.Add(materialInstance);

	// 2. 투명 렌더링 모드로 전환
	SetMaterialToTransparent(materialInstance);

	// 3. 각 렌더러의 머티리얼에 대한 페이드 아웃 트윈을 생성
	Tween fadeTween = materialInstance.DOFade(0f, 0.2f);

	// 4. 시퀀스에 조인함. 
	deathSequence.Join(fadeTween);
}

deathSequence.OnComplete(() =>
{
	OnDeathAnimationCompleted?.Invoke(this);

	foreach (Material mat in materialInstances)
	{
		Destroy(mat);
	}

	Destroy(gameObject);
});
```

### Operator 배치 시에 공격 범위 내에 적이 있는데 공격하지 않는 현상
-  `Enemy`가 저지당하고 있는 상태에서 `Operator`가 배치됐을 때 저지당하고 있는 적이 공격 범위 타일 내에 있는데도 공격하지 않는 현상.
	- 정확한 원인은 몰라도 오퍼레이터가 공격 범위 내에 있는 적들을 제대로 캐치하지 못하는 현상이므로 관련 로직을 보면 될 듯.
	- 일단 `Enemy`가 밟고 있는 타일이 바뀌지 않는 상태여야 함

```cs
UpdateAttackableTiles(); // 방향에 따른 공격 범위 타일들 업데이트
RegisterTiles(); // 타일들에 이 오퍼레이터가 공격 타일로 선정했음을 알림
```
- `Deploy()`에서 실행되는 로직인데, 이 부분의 순서가 바뀌어 있었다. 이걸 바꿔보고

```cs
// 공격 범위 타일들에 이 오퍼레이터를 등록
private void RegisterTiles()
{
	foreach (Vector2Int eachPos in CurrentAttackableGridPos)
	{
		Tile? targetTile = MapManager.Instance!.GetTile(eachPos.x, eachPos.y);
		if (targetTile != null)
		{
			targetTile.RegisterOperator(this);
			
		}
	}
}
```
이 메서드에서 타일들에게 자신을 공격 범위로 삼은 오퍼레이터를 등록했는데, 여기에 추가로 타일들에서 오퍼레이터에게 자신을 밟고 있는 `Enemy`의 정보도 전달해야 할 것으로 보임.

```cs
if (targetTile != null)
{
	targetTile.RegisterOperator(this);

	// 타일 등록 시점에 그 타일에 있는 적의 정보도 Operator에게 전달함
	foreach (Enemy enemy in targetTile.EnemiesOnTile)
	{
		OnEnemyEnteredAttackRange(enemy);
	}
}
```

- 일단 이렇게 수정하니까 **비슷한 현상은 발생하지 않고 있다.**
- 추가로, **걸리적거렸던 것 중에 원거리 오퍼레이터를 배치했는데 왜 바로 공격하지 않지? 라는 게 있었다. 이것도 해결된 것으로 보임.** 





# 250818 - 짭명방

>[!done]
>1. 이펙트 구현
>	- `Hit_Explosion`
>	- `Area - Medic2ndSkill`
>		- `SkillRangeVFXController` 스크립트 수정 
>2. 이슈 수정
>	- `Distortion` 쉐이더의 이슈 
>		- 파티클 시스템에서 게임 뷰에 쓰일 때 검정색의 메쉬가 나오는 현상
>		- 해결할 필요가 없었다. 실제 게임에서는 투명한 메쉬가 잘 나옴. 몇 시간을 헤맸나..

## 이펙트 구현 - Hit_Explosion
- 남은 요소는 크게 2개였다 : 폭발 이펙트와 투명한 충격파 구현하기.

### 투명한 충격파
- `Heat Distortion`이라고 부를 수도 있는 요소인데, 그저께 찾아봤는데 답을 쉽게 얻지 못했다. 
- **센세께서 올려두신 영상이 있었다.** [# Unity VFX Graph - Heat Distortion Effect Tutorial](https://www.youtube.com/watch?v=CXCyVDEplyM)
	- 이 중에서는 쉐이더에 관한 내용만 가져가면 된다. URP와 HDRP의 내용이 다른 부분이 꽤 있기 때문에 URP에 있는 부분만 가져가면 됨.

- 프로젝트의 `Default Render Pipeline`에 해당하는 파일에서, **`Rendering` 탭에 있는 `Depth Texture`, `Opaque Texture`를 켜야 한다.**
	- 이게 꺼져 있으면 셰이더 그래프에서 `Scene Color`를 적용해도 아무 변화가 없음. 반대로 켜져 있으면 해당 메쉬의 뒷쪽이 그대로 투명하게 비친다. 알파값과는 무관함.

- **구현도 간단하다!** [[투명한 왜곡 효과]]에 정리.
![[Pasted image 20250818155210.png]]
- `Scene Color`를 이용하면 배경의 색을 그대로 보여준다. 여기에 변화된 `Screen Position`을 `Scene Color`의 UV 인풋에 보내면 해당 메쉬에는 왜곡된 배경이 나타난다.

#### 이 과정에서 발생한 이슈
- 씬 뷰에서는 투명하게 잘 나타나는데, 게임 뷰에서는 검정색으로 이펙트가 나타남
	- 어떻게 해결해야 하나 한참 찾아 헤맸는데, **실제 게임 & 게임 뷰에서는 아무 이상 없어서 해결하지 않아도 되는 이슈**였음. 몇 시간을 썼는데,,,
### 메쉬
- 블렌더에서 `Torus`를 만들었다. 버텍스 숫자를 좀 줄인 상태로.


- `ExplosionDebris`라는 폭발 흔적 텍스쳐를 하나 만들었고, `Crack01`도 텍스쳐 자체의 밝기? 불투명도를 크게 줄인 버전으로 수정했다.

![[Hit_Explode_v1 1.gif]]
연기가 너무 많나 싶기는 한데..? 일단 인게임 테스트.
여기서 검게 보이는 부분은 인게임에서는 투명해짐. 
- 연기랑 파티클이 너무 높게 올라가는 것처럼 보여서 추가로 수정

![[Hit_Explode_v1_Ingame.gif]]
이런 느낌으로 마무리해본다. gif 파일로 보니까 생각보다 더 괜찮은 것 같음.

## 이펙트 구현 - Area_Medic2ndSkill
- 이건 기존에 파티클 시스템으로 구현되어 있었기 때문에 이번엔 수정만 한다.
- 아마 `Area`마다 다르겠지만 `Medic2ndSkill`의 경우 파티클 1개를 띄워놓고 실행시키는 방식이었는데, 굳이 이렇게 구현할 필요는 없어보인다. `Quad`를 변에 세우고 셰이더만 적용하면 될 거임.

### SkillRangeVFXController 수정
- 그래서 기존에 이러한 타일 기반 범위 이펙트에 구현했던 `SkillRangeVFXController`도 수정한다. 
- 기존 : 파티클 시스템 기반
	- 오브젝트 기반으로 구현하고, 방향에 따라 이펙트를 실행하는 스크립트에서 해당 오브젝트에 파티클 시스템이 있는지를 검사해서 오브젝트를 활성화시키는 방식으로 수정함.
```cs
private void SetUpVisuals(Vector2Int position, HashSet<Vector2Int> effectRange)
{
	// 유효하지 않은 위치는 아무것도 표시하지 않음
	if (MapManager.Instance.CurrentMap == null || !MapManager.Instance.CurrentMap.IsTileAt(position.x, position.y))
	{
		return;
	}

	floorImage.gameObject.SetActive(false);

	// 방향에 따른 타일 검사로 이펙트 실행 여부 결정
	foreach (var direction in directions)
	{
		Vector2Int neighborPos = position + direction;
		bool showEffect = !effectRange.Contains(neighborPos) || !MapManager.Instance.CurrentMap.IsTileAt(neighborPos.x, neighborPos.y);

		var (effectObject, boundary) = directionEffects[direction];

		effectObject.SetActive(showEffect);
		boundary.gameObject.SetActive(showEffect);

		// 파티클 시스템으로 구현된 경우 파티클 시스템을 실행시킴
		ParticleSystem directionParticleSystem = effectObject.GetComponent<ParticleSystem>();
		if (directionParticleSystem != null)
		{
			PrewarmTrailAndPlayVFX(directionParticleSystem); // effect.Play() 포함
		}
	}

	// 언덕 타일 위치 보정
	Tile? currentTile = MapManager.Instance.GetTile(position.x, position.y);
	if (currentTile != null && currentTile.data.terrain == TileData.TerrainType.Hill)
	{
		transform.position += Vector3.up * 0.2f;
	}
}
```

![[Area_Heal_v1.gif]]
> 언덕 타일에서 살짝 띄워져보이는 게 어색해보이긴 한다. 수정해야 할까는 고민 대상인 듯.

일단 오늘은 여기까지!!


## 기타 수정
- `Projectile`의 콜라이더 크기 `0.1`로 통일
	- 기존) `0.25` : 오브젝트에 부딪히기 전에 사라지는 것처럼 보여서.




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