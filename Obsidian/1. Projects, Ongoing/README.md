# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 남은 작업

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

## 이슈? 정리
1. **`DualBlade`는 아이콘의 해상도를 256으로 올려도 살짝 깨져서 보인다.** 가우시안 블러를 미리 해놔서 그런가? 왜 그러지..
2. 현재) **이중 공격을 할 때, 적이 1타를 맞고 죽으면 2번째 공격은 나가지 않는 구조다.**
	- **적이 죽더라도 2번째 공격은 나가게 해야 하나? 아니면 지금 상태를 유지하는 게 좋나?**
	- 게임을 생각해보면 소리 같은 것도 다 2번 들리니까 그거에 맞춰서 공격을 나가도록 구현한 건지, 아니면 1대만 때리고 소리만 2대 나오는 건지 모르겠다. 현상유지해도 큰 상관은 없어보이지만 실제로 어떤 식으로 구현되는지 정리해볼 필요는 있겠다.


# 250708 - 짭명방
- `DualBlade`의 클래스 아이콘 이미지 작업.
	- 칼 2개와 명방의 마름모를 넣고 싶다. 마법 공격을 하는 2번째 직군이기도 해서.
![[DualBlade_128.png]]
> 뭔가 단순하고 마법의 맛이 없지만 그냥 이거 쓰기로 함. ChatGPT한테도 맡겨봤는데 스타일이 너무 화려해서 직접 그렸다.

> 여기서 옵시디언으로 볼 때는 괜찮은데 인게임에서 아이콘을 보면 계단현상이 좀 두드러져 보인다. `Artillery`도 비슷한 현상이 있어서, 얘네 둘은 해상도를 256으로 올려서 넣겠음.

- 스킬은 일단 나중에 구현하기로 하고 기초 테스트만 해본다
	- `ScriptableData`로 만들고, Prefab 연결하고, ResourceManager에 연결된 `OperatorIcon`에 아이콘 추가하고. `Prefab`에는 `Operator` 대신  이전에 구현했던 `DualBladeOperator` 스크립트를 연결했다. 
	- **무한 공격 이슈** : `DualBladeOperator`의 이중 공격 부분 다음에 `AttackDuration, AttackCooldown` 설정함.




# 250706 - 블로그

## 스타일 작업
- 이전에 글씨 크기를 `text-base`로 통일하는 과정에서, `.ql-editor` 클래스를 사용하는 곳들 모두에 공통적으로 적용되지 않는다는 걸 알게 되었다. 
- 글을 작성할 때랑 읽을 때의 서식을 통일해야 함 

### `format` 중에 `list`를 사용할 것인가 말 것인가? 
- 사실 크게 상관은 없는 것 같다. 왓챠에 `- 내용`처럼 서술한 다음에 복붙해도 자동으로 `li` 태그가 붙지는 않는다. 일단 내용이 그대로 들어감.
- 이게 문제가 됐던 상황은 티스토리로 옮길 때였다. 근데 지금은 크게 상관없지 않을까.
- 이거는 일단 **유지**하기로 결정. 문제가 되면 그 때 다시 보자.

### quill 사용하는 기능과 각각의 클래스들 정리
> 여기서 `quill` 클래스는 정확히는 `"quill "`으로 뒤에 공백이 하나 있다. 왜 이런지는 모르겠음.

- `my-doing`
	- 글 읽기 : `ql-snow ql-editor` (아래에 아무것도 없는 div 태그가 하나 더 있음)
		- 아래의 `div` 태그 제거해봄
	- 글 쓰기 : `quill ql-container(&ql-snow) ql-editor`

- `post`
	- 글 읽기 : `ql-snow ql-editor ql-container` 아래에 div 태그
	- 글 쓰기 : `quill ql-container(&ql-snow) ql-editor`

- `review`
	- 글 읽기 : `review-content`
	- 글 쓰기 : `quill ql-container(&ql-snow) ql-editor`

1. **글 읽기에서 `ql-container`라는 게 굳이 필요할까?**
	- `Post`의 경우 오른쪽에 `HeaderLinks`로 나타나는 부분에서 `.ql-container`의 아래에 있는 헤더들을 추적했다. 그런데 그냥 `.ql-editor`로 해버리면 되니까 그렇게 변경함.
	- 대부분의 경우 `.ql-snow .ql-editor`가 공통되기 때문에 여기에만 적용해버려도 되지 않을까?
	- 또 `Post`의 글 읽기에는 `article-content`라는 게 따로 있다. 이거는 스타일에는 사용되지 않고..

- **전부 `.ql-snow .ql-editor` 으로 서식을 통일**
	- `.my-doing .ql-editor`의 경우만 헤더 크기를 줄이는 식으로 수정함
	- 이게 더 위지윅에 맞긴 하겠다.

## 리뷰 카드 중 텍스트가 있는 것과 없는 것 표시
- 리뷰 카드를 작성했지만 본문에 텍스트가 없을 수도 있다.
- 이런 경우를 구분하기 위해 **제목 끝에 아이콘 하나를 추가**하려고 함.
- 그래서 카드들에서 내용이 있는지 여부를 확인하고 글이 있다면 아이콘으로 표시해주는 게 좋을 것 같음. 위치는 제목이 끝난 바로 뒤로.
- 카드로 보내는 시리얼라이저에 아래 필드를 추가한다.
```python
class ReviewListSerializer(serializers.ModelSerializer):
    """목록 뷰(그리드)를 위한 시리얼라이저"""
    has_content = serializers.SerializerMethodField(); # 읽기 전용 필드, get_ 접두사가 붙은 메서드를 통해 동적으로 생성된다.

    class Meta:
        model = Review
        fields = [
	        # ...
	        ,
            'has_content'
            ]
        
    def get_has_content(sefl, obj: Review) -> bool:
        """
        Review 필드의 content 필드가 비어 있는지 확인한다.
        None이거나 ""이 아니라면 True
        """
        return bool(obj.content) # 프론트에서 받은 백엔드의 빈 컨텐츠는 ""으로 되어 있음
```
> - 프론트에서도 이 필드를 추가하고, `has_content` 필드를 확인해서 값이 있으면 제목 뒤에 아이콘을 표시하는 방식이다.
> - 프론트에서 비어있는 글의 백엔드의 데이터가 어떻게 오는지 확인했다. ""으로 옴.

- 제목 뒤에 표시하려고 했는데 생각보다 별로다. `별점, 스포일러, 보는 중` 이 들어가는 줄에 추가.
![[Pasted image 20250706171948.png]]


# 250704 - 짭명방
## 2번 공격하는 근접 직군 구현

- 이름은 `DualBlade`?
- 특징
	- 기본 공격은 1번 공격 시 2회 타격함
	- 모든 스킬은 공격 시 SP가 회복됨
		- 즉 공격 쿨타임이 돌 때마다 2회 공격하고, SP는 때릴 때마다 1씩 총 2가 올라감.
	- 컨셉은 살짝 바꾸겠음. 스킬을 켜지 않았을 때 물리딜을, 켜거나 사용할 때는 마법딜을 넣는 식으로.
	- 1스킬은 강타로 구현하되 마법 타입으로 공격이 바뀌어서 나감
	- 2스킬은 버프로 구현할까?

- 견본이 좀 있나 찾아봤는데 소드마스터 자체가 5성 이하로 별로 없다. 바이비크, 타찬카, 커터 정도.
	- 스탯은 바이비크의 것을 가져옴

- 스크립트 자체는 이런 구현에서 시작.
```cs
public class DualBladeOperator : Operator
{
    // 공격 사이의 간격
    private float delayBetweenAttacks = 0.15f;

    public override void Attack(UnitEntity target, float damage)
    {
        // 2회 공격 로직을 코루틴으로 구현
        StartCoroutine(DoubleAttackCoroutine(target, damage));
    }

    private IEnumerator DoubleAttackCoroutine(UnitEntity target, float damage)
    {
        bool showDamagePopup = false;
        float polishedDamage = Mathf.Floor(damage);

        base.PerformAttack(target, polishedDamage, showDamagePopup);

        yield return new WaitForSeconds(delayBetweenAttacks);

        if (target != null && target.CurrentHealth > 0)
        {
            base.PerformAttack(target, polishedDamage, showDamagePopup);
        }
    }
}
```

> 이거 일단 중지!

## Operator 리팩토링

### Skill - Operator 관계 재정립

- `BaseSkill`의 경우
```cs
	public bool autoRecover = false; // 활성화시 자동 회복, 비활성화 시 공격 시 회복
	public bool autoActivate = false; // 자동발동 여부
	public bool modifiesAttackAction = false; // 공격 액션 변경 여부
```
이런 필드들이 있다. 그러면 이 필드의 값에 따라 SP 회복 로직이 달라진다든가, 입력을 대기하거나 스킬이 자동으로 나간다거나 하는 부분은 전부 Operator 자체에서 일어나야 하는 일이다.

지금의 `SmashSkill` 같은 경우
```cs
protected override void SetDefaults()
{
	autoActivate = true;
}

// 공격에 묻어나가는 로직
public override void OnBeforeAttack(Operator op, ref float damage, ref bool showDamage)
{
	if (op.CurrentSP >= op.MaxSP) 
	{
		damage *= damageMultiplier;
		showDamage = true;
		op.CurrentSP = 0;
	}
	else
	{
		op.CurrentSP += 1;
	}
}
```
이런 식으로 스킬 자체에서 오퍼레이터의 SP를 회복하는 로직이 있는데, 이렇게 구현하지 말고 

- `Operator`에서는 **스킬 -> 오퍼레이터로 전달시켜서 회복시키는 게 아니라 오퍼레이터 자체에서 필드를 확인하고 그에 따른 동작을 수행하는 게 더 맞는 구현**이 된다.

1. `SmashSkill`의 `else` 부분을 제외한다. 회복 동작은 `BaseSkill.autoRecover = false`일 때 `Operator`에서 동작한다.

2. 공격 시 SP 회복 로직은 `Operator.PerformAttack`의 공격 후에서 구현한다.
```cs
    protected virtual void PerformAttack(UnitEntity target, float damage, bool showDamagePopup)
    {
		float spBeforeAttack = CurrentSP;
	    // 공격 전
        if (CurrentSkill != null)
        {
            CurrentSkill.OnBeforeAttack(this, ref damage, ref showDamagePopup);
        }
		
		// 실제 공격 동작
		// ...
		
		// 공격 후
        if (CurrentSkill != null)
        {
            CurrentSkill.OnAfterAttack(this);
            
            // SP 공격 시 회복 로직
            if (!CurrentSkill.autoRecover && !IsSkillOn && spBeforeAttack != MaxSP)
            {
                CurrentSP += 1; // 세터에 Clamp가 있으므로 여기서 하지 않아도 됨.
            }
        }
    }
```
> 추가로, sp가 최대일 때 나간 공격은 `OnBeforeAttack`에서 스킬이 발동되면서 sp가 0이 되므로 해당 공격에서는 SP가 회복되지 않도록 수정했다.

### 잠깐 휴식 전 정리
- 지금 신경쓰이는 거
```cs
protected void HandleSPRecovery()
{
	// ...

	if (CurrentSP != oldSP && operatorUI != null)
	{
		operatorUI.UpdateUI();
		OnSPChanged?.Invoke(CurrentSP, MaxSP);
		
		// ..
	}
}
```
> - 여기서 **OnSPChanged?.Invoke()로 operatorUi.UpdateUI까지 통합해버리는 게 좋아보인다.**
> - 그런데 `OnSpChanged.Invoke` 이벤트를 구독하는 부분은 `DeployableActionUI` 이랑 `DeployableBarUI`이다. `OperatorUI`는 `DeployableBarUI.SetSPBarColor` 를 수정하고 스킬 아이콘 활성화 여부를 결정하는데..
> - 이벤트로 다 묶어버릴 수 있을 것 같음.
> - 문제라면 `DeployableBarUI`가 따로 있다는 것인데.. 이따 생각해보자.

- 휴식 끝!
###  1. Operator - OperatorUI 정리
1. `OperatorUI` 자체는 `Operator` 자체에서 생성과 파괴를 담당
2. HP 변경, SP 변경 등은 이벤트로 관리
	- 왜냐면 `1:多` 관계임. 저 사건들을 사용할 컴포넌트들을 일일이 관리하는 건 번거로움.
	- 하지만 `UI 자체는 오퍼레이터와 1:1 관계`임

- 이런 구조로 바꿨다. 기존엔 이것저것 엉켜있었음.
```
[ Operator ]  <-- (데이터 소스)
     |
     | (이벤트 발생: OnHealthChanged, OnSPChanged, OnSkillStateChanged)
     V
[ OperatorUI ] (컨트롤러: 이벤트 구독 및 로직 분배)
     |
     | (메서드 호출: UpdateHealthBar, SetSPBarColor 등)
     V
[ DeployableBarUI ] (뷰: 단순한 API 제공 및 하위 컴포넌트 관리)
     |
     +--- [ HealthBar ] (실제 뷰)
     |
     +--- [ SPBar ] (실제 뷰)
```

### 2. Skill - Operator - OperatorUI 정리
- 크게 2가지 궁금한 게 있다.
	1. **오퍼레이터의 스킬을 켜고 끄는 걸 어디서 처리해야 할까?**
		- **스킬의 시작과 끝을 처리하는 지점은 스킬 자체**다. 그래서 **어떤 시점에 어떤 동작을 해야 하는지 아는 스킬에서 오퍼레이터의 상태를 함께 관리**한다.
		- 스킬은 `Operator`의 상태를 변경해달라고 요청하면 `Operator`는 요청을 받아 상태를 변경하고 변경된 사실을 외부에 방송한다.
	2. 스킬이 켜졌을 때 SPBar의 동작은 어디에서 관리해야 할까?
		- SPBar가 어떻게 변하는가는 SPBar 자체에서 처리하면 된다. 스킬에서 일일이 관리할 필요 없다.

## 기타 버그
- [x] `Artillery` 공격이 동작했는데도 대미지가 안 들어가는 것처럼 보이는 이슈
	- `Projectile`에서 폭발하는 경우의 콜라이더 처리가 바뀌어야 한다. 원래는 `UnitEntity`를 직접 감지했으나, 이전에 `Body`를 각 객체의 자식 오브젝트로 별도로 구현하고 거기에 `BodyColliderController` 스크립트를 붙였던 적이 있다. 그 컴포넌트를 감지하도록 수정함.

- [x] `MedicOperator`의 공격이 연속적으로 쫘라락 나가는 현상
	- 공격 쿨타임 적용하는 로직을 Attack 내부로 바꾸면서 발생한 문제인 듯.
	- 오버라이드하기에는 살짝 구조가 달라서 `SetAttackDuration, SetAttackCooldown`을 똑같이 `MedicOperator`에 넣었다.

- [x] `Operator`의 저지도 이상하게 동작한다.
	- 상황) Operator가 저지 중인 적을 성공적으로 제거했을 때, 콜라이더가 겹치지 않는 상황인데 해당 적이 저지당하는 현상이 있음. 이전과 달리 근거리, 원거리를 가리지 않음.
	- `blockableEnemies`가 제대로 처리되지 않은 것으로 보인다. 즉, **콜라이더가 겹쳤을 때 저지 후보에는 들어갔는데, 콜라이더에서 이탈했는데도 저지 후보에서 제거되지 않은 것으로 보임**.
	- `Operator`의 `public override void OnBodyTriggerExit(Collider other)`가 `private`으로 돼 있긴 했었다. 체크해보고 다시 실행시켜봄.

- [x] `Operator`가 배치될 때 겹쳐진 적을 저지하지 않는 현상도 있다.
	- 현재 `BodyCollider`로 뺀 상태이고 배치될 때 이를 활성화한다고 하자. 이 때, `OnTriggerEnter`는 이미 겹쳐진 콜라이더에는 동작하지 않는다. 
	- 따라서 활성화 시점에서 겹쳐져 있으므로 동작하지 않고, 그 다음 프레임에도 `OnTriggerEnter`는 이미 겹쳐있으니까 동작하지 않는다. `OnTriggerStay`는 동작함.
	- 성능까지 고려해본다면, 활성화 시점에 겹쳐진 콜라이더를 체크해서 `OnTriggerEnter`로 넘겨주면 되지 않을까?

- `BodyColliderController`를 아래처럼 구현했다.
```cs
// 이 컨트롤러의 콜라이더 활성화 상태 결정
public void SetColliderState(bool enabled)
{
	if (bodyCollider != null)
	{
		bodyCollider.enabled = enabled;

		// 콜라이더가 켜지는 순간에는 수동 겹침 검사 실행
		if (enabled)
		{
			CheckForInitialOverlaps();
		}
	}
}

// 콜라이더가 활성화된 시점에 겹쳐져 있는 코라이더를 찾아 `OnTriggerEnter`처럼 owner에게 전달한다.
private void CheckForInitialOverlaps()
{
	if (owner == null) return;

	// 콜라이더의 타입을 확인해 Overlap 함수를 사용한다.
	if (bodyCollider is BoxCollider box)
	{
		// BoxCollider와 충돌하는 콜라이더들을 찾음
		Collider[] overlappingColliders = Physics.OverlapBox(
			transform.position + box.center,
			Vector3.Scale(box.size, transform.lossyScale) / 2, // 스케일링을 고려한 실제 크기
			transform.rotation,
			-1, // 모든 레이어
			QueryTriggerInteraction.Collide // 트리거 콜라이더와도 충돌하도록 설정
		);

		foreach (var otherCollider in overlappingColliders)
		{
			// 자기 자신은 무시
			if (otherCollider == bodyCollider) return;

			// 감지된 콜라이더를 owner에게 전달
			owner.OnBodyTriggerEnter(otherCollider);
		}
	}
}
```

- [x] 원거리 `Enemy`가 원거리 공격하지 않는 문제
	- 이건 갑자기 왜 그러는 걸까?
	- 얘도 비슷한 문제겠다. 즉, `OnTriggerEnter`에서 감지하는 게 `DeployableUnitEntity`를 직접 감지하는 게 아니라, 본체 트리거 감지 -> 그 부모에 `DeployableUnitEntity`가 있는가? 가 되는 것.
	- **오늘 발생한 대부분의 문제가 이 본체 콜라이더를 자식으로 이동시키면서 발생한 문제들**이다. 

- [x] `EnemyBarUI`의 체력 변화를 이벤트 기반 구독으로 변경
	- `Operator`에 비해 훨씬 쉽다. 체력밖에 없고 고려할 것도 많이 없음

# 250703 - 블로그

## 오늘의 배운 점
1. **모바일을 고려할 때, 화면에 띄울 요소이면서 잘리지 않기를 바란다면 `dvh`를 쓰자.**
	- `vh`는 모바일 브라우저에서 주소창이 나타나지 않았을 때를 기준으로 뷰포트의 높이를 계산하므로 의도한 것보다 화면에 작게 나타날 수 있다. 

## 리뷰 모달의 위아래가 잘려서 나타나는 현상
- 기존 모달의 높이는 `max-h-[95vh]`로 구현했음. PC에서 모바일 너비로 테스트해도 위아래에 빈 공간이 조금 있는 식으로 잘 구현됨. 하지만 실제로 모바일 환경을 보면, **모달의 위아래 빈 공간이 없고 오히려 모달의 일부 공간이 잘리는 현상이 발생했음.**

이런 현상은 모바일 웹 개발에서 발생하는 `100vh` 문제라고 한다. 

### 100vh 문제
- `Vertical Height` 단위의 기준은 뷰포트 높이의 1%이다. 
- 그런데 **모바일 브라우저에서의 뷰포트 높이는 주소창이 축소되었을 때의 높이를 기준**으로 계산된다.
	- 모바일 브라우저는 동작에 따라 주소창이 나타날 때도 있고 숨겨질 때도 있는데, 나타날 때에 문제가 된다는 것이다. 

### 해결책
- 가장 좋은 방법은 **`dvh`라는 걸 사용**하는 것이다. `vh`를 더 깊게 들어가면
- `svh(Small vh)` : 브라우저 UI가 확장된 상태, 가장 작은 가시영역의 높이
- `lvh(Large vh)` : 브라우저 UI가 축소된 상태, 가장 큰 가시영역의 높이 
- `dvh(Dynamic vh)` : 브라우저 UI에 따라 동적으로 변하는 높이

지금 상황에서 화면에 나타났을 때 잘리지 않고 꽉 차게 보여야 하는 요소라면 `dvh`나 `svh`가 적합하다.


## 리뷰 모달이 떠 있는 상태에서 뒤로 가기나 새로고침 동작
- 현재는 리뷰 모달이 떠 있는 상태에서 뒤로 가기를 누르면 아예 현재 창이 닫힌다. 이거는 `url` 단위로 움직이기 때문으로 보임. 
- 모달이 떠 있는 상태에서 뒤로 가기 버튼이 눌리면 기존의 `url` 동작을 취소하고 현재 떠 있는 모달이 닫히는 방식으로 구현하면 될 것 같음.
- 새로고침의 경우는 어떻게 구현해야 할까?
	- 새로고침 자체가 거슬린다기보다는 **가장 윗스크롤에서 위로 화면을 더 당길 때 새로고침이 튀어나오는 게 거슬린다...** 인 것 같음.
	- 저 동작만 막으면 되지 않을까? 

### 뒤로 가기 동작 제어

**원리**
1. 모달이 열릴 때 `history.pushState()`로 브라우저의 방문 기록 상태에 가짜 상태를 하나 추가한다.
2. 사용자가 뒤로 가기 버튼을 누르면 브라우저는 가짜 상태로 돌아가려는 `popstate` 이벤트를 발생시킨다.
3. `popstate` 이벤트르 감지해서 모달을 닫는 함수를 실행한다.

위 로직을 커스텀 훅으로 만들어서 깔끔하게 적용할 수 있다.
```ts
import { useEffect } from 'react';

// 모달이 열려 있을 때 브라우저의 뒤로 가기 버튼을 누르면 모달을 닫도록 처리하는 훅
export const useModalHistory = (isOpen: boolean, onClose:() => void) => {
    useEffect(() => {
        // 모달이 열렸을 때만 로직 실행
        if (isOpen) {
            // 히스토리 스택에 상태 추가 (URL 변경 X)
            // 모달이 열렸음을 의미함
            window.history.pushState({ modalOpen: true }, '');

            // popstate 이벤트 리스너 정의
            const handlePopState = () => {
                // 뒤로 가기 시 모달 닫기 함수 호출
                onclose();
            }

            window.addEventListener('popstate', handlePopState);

            // 클린업 함수 : 컴포넌트가 언마운트되거나 isOpen = false일 때 실행
            return () => {
                window.removeEventListener('popstate', handlePopState); 

                // 히스토리 스택의 현재 상태가 여기서 추가된 모달이라면
                // 사용자가 뒤로가기 대신 다른 방법(X 버튼 등)으로 모달을 닫았다는 의미라서
                // 스택에서 해당 상태를 제거해줘야 다음 뒤로가기가 정상적으로 작동한다.
                if (window.history.state?.modalOpen) {
                    window.history.back();
                }
            }
        }
    }, [isOpen, onClose])
}
```

### 새로고침 제어

**문제 상황 분석**
1. 모달의 스크롤이 동작할 수 있는 상황이라면 `div.overflow-y-auto`에서 먼저 처리된다. 
2. **모달의 스크롤이 불가능하다면 스크롤 이벤트를 부모로 넘겨준다.**
3. 부모 요소의 스크롤이 동작한다. 예를 들면 `<body>`의 스크롤이 동작.

이런 원리로 모달의 스크롤이 가장 위에 있는 상태에서 더 위로 스크롤하려고 할 때, 모바일 브라우저의 경우 `<body>`의 최상단에서 위로 스크롤하는 제스쳐는 `당겨서 새로고침(Pull-To-Refresh)` 기능과 연결되어 있고, 그래서 새로고침하려고 하는 것이다.

이런 식으로 스크롤 동작이 불가능할 때 상위 요소로 스크롤 이벤트가 전달되는 것을 **`스크롤 체이닝`** 이라고 한다.

**해결 원리** : `overscroll-behavior-y: contain`
- `overscroll-behavior` : 스크롤 경계에 도달했을 때 브라우저의 동작을 정의하는 css 속성
- `contain` : 스크롤 경계에 닿으면, 스크롤 이벤트를 부모 요소로 전파하지 말고 여기서 끝내라는 지시.

위 동작을 모달이 열렸을 때 `<body>` 태그에 추가하고, 닫힐 때 `<body>` 태그에서 제거한다.
이 로직도 커스텀 훅으로 구현할 수 있다.
```ts
import { useEffect } from 'react';

// 특정 조건에서 body의 스크롤을 방지하는 훅.
export const usePreventScroll = (isLocked: boolean) => {
    useEffect(() => {
        if (isLocked) {
            const originalStyle = window.getComputedStyle(document.body).overflow;
            const originalOverscrollBehavior = window.getComputedStyle(document.body).overscrollBehaviorY;

            // 스크롤 방지 스타일 적용
            document.body.style.overflow = 'hidden';
            document.body.style.overscrollBehaviorY = 'contain';

            // 클린업 함수 : 원래 함수로 복원
            return () => {
                document.body.style.overflow = originalStyle;
                document.body.style.overscrollBehaviorY = originalOverscrollBehavior;
            } 
        }
    }, [isLocked])
}
```

### 적용
- 모달 컴포넌트에 2개의 훅을 적용하면 된다. 페이지 단위로 적용하지 않아도 됨.
```ts
import { useModalHistory } from '../../hooks/useModalHistory'; // 뒤로가기 동작 관련
import { usePreventScroll } from '../../hooks/usePreventScroll'; // 배경 스크롤 방지

const ReviewModal: React.FC<Props> = ({ reviewId, onClose, mockData }) => {
	// ...
	
	usePreventScroll(!!reviewId); // 모달이 열린 동안(reviewId가 있을 때) 배경 스크롤 / 새로고침 방지
	useModalHistory(!!reviewId, onClose); // 뒤로가기 버튼에 모달 닫는 기능 추가
	
	//...
  }
```
> `!!`는 어떤 값을 `boolean`으로 나타내기 위한 이중 변환이다. `falsy`한 값을 `false`로, `truthy`한 값을 `true`로 쓰기 위한 관용적인 표현`idiom`이라고 함.


## 추가 수정 사항

### 1. 모달을 켜고 다른 탭에 다녀올 때 현재 모달이 꺼지는 문제
- 위에서 추가한 `useModalHistory`를 비활성화했을 때는 이 문제가 발생하지 않으므로 새로 추가한 기능이었던 이게 원인이 맞다. 
- 일부 브라우저(크롬 포함)는 다른 탭으로 갔다가 다시 돌아올 때, 내부적으로 히스토리 상태를 확인하는 과정에서 `popstate` 이벤트를 발생시키는 경우가 있다. 
- 따라서 `popstate` 이벤트가 발생했을 때 무조건 모달을 닫는 게 아니라, **이 이벤트가 실제로 사용자의 뒤로 가기 이벤트였는지를 추적**하는 게 더 좋다.

```tsx
useEffect(() => {
	// 모달이 열렸을 때만 로직 실행
	if (isOpen) {

		// historyState가 없으면 추가한다. 불필요한 pushState 호출을 막는다.
		if (!isModalStatePushed.current) {
			window.history.pushState({ modalOpen: true }, '');
			isModalStatePushed.current = true;
		}

		// popstate 이벤트 핸들러 정의. 사용자가 뒤로 가기를 했다는 의미.
		const handlePopState = (event: PopStateEvent) => {
			// popstate가 발생했으나 현재 state가 이 모달 state라면
			// 즉 실제 뒤로 가기가 아닌 탭 전환 등의 이벤트라면 무시한다.
			if (event.state?.modalOpen) {
				return;
			}

			// 진짜 뒤로가기가 실행된 상황
			onClose();
		}
		
		window.addEventListener('popstate', handlePopState);

		return () => {
			// 클린업 함수 : 컴포넌트가 언마운트되거나 isOpen = false일 때 실행
			window.removeEventListener('popstate', handlePopState);
		}
	}
	// -- 모달이 닫혔을 때의 로직 --
	else {

		if (isModalStatePushed.current) { 
			isModalStatePushed.current = false; // 상태를 먼저 변경해 무한 루프를 방지한다.
			window.history.back(); // popstate 이벤트 발생
		}
	}
}, [isOpen, onClose])
```

> 원리 이해하기
- 모달을 열면 useModalHistory 훅이 실행된다. `window.history.pushState({ modalOpen: true }, '')`
```
[맨 위]  -> { state: { modalOpen: true }, URL: /reviews }  <-- 현재 위치 (모달 열림)
[아래]   -> { state: null, URL: /reviews }                <-- 모달 열기 전 상태
[그 아래] -> { state: null, URL: / }                       <-- 이전 페이지
```

**시나리오 A : 사용자의 뒤로 가기 버튼 입력**
1. 히스토리 스택의 현재 위치를 한 단계로 내린다. `[아래]` 부분이 됨.
2. 히스토리 변경을 알리기 위해 `popstate`를 발생시킨다.
3. `handlePopState(event)`가 동작한다. `event`는 새롭게 이동한 위치의 상태가 담겨있다.
4. `event.state`는 따라서 `null`이 된다. 

**시나리오 B : 사용자가 다른 탭으로 갔다가 돌아옴**
1. 일부 브라우저는 탭이 다시 활성화될 때, 현재 히스토리 상태를 확인하고 일관성을 유지하기 위해 `popstate` 이벤트를 발생시킬 수 있다.
2. 이 때, 히스토리 스택은 여전히 `[맨 위]`다. 히스토리 스택 이동이 없었으니까.
3. 따라서 `event.state`는 `{ modalOpen: true }`가 된다.

### 2. 글 작성 시 번호나 리스트 (1. 2. 3. 이나 -)를 쓰면 자동으로 서식이 지정되는 현상
- `HTML`과 부드럽게 연결되지 않으므로 이런 것들은 그냥 문자열들로 구현하려고 함
- `QuillEditor`에서 `formats`에 있는 `list, bullet` 을 제거했다. 전자는 순서 리스트, 후자는 순서 없는 리스트.

### 3. 헤더 스타일이 ReviewModal에 적용됐으면 좋겠음. 오른쪽 링크는 필요 없고.
- `ReviewModal`의 본문 내용을 감싸는 `div.review-content` 클래스를 하나 만들어줬다. 
- 스타일은 대부분 `ql-snow ql-editor h1~h3`에 해당하는 것을 가져와 수정했음.

### 4. WriteReviewModal의 경우 배경 클릭 시
- 배경을 클릭해도 모달이 꺼지지 않도록 구현한다. 왜냐면 작업 중인데 오조작 한 방에 작업 중이던 내용이 날아갈 수 있기 때문임.




## 실패) 모바일 테스트 구현
- 핸드폰을 usb에 연결해서 로컬 호스트에 접속 가능하게 하면 모바일 환경에서도 직접 `npm run dev` 환경으로 테스트할 수 있을 거라는 생각이 들어서 시도해봤다.
- 연결 자체는 가능했는데, 중간에 무한 로딩이 걸려버려서 중지.



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