## 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 작업 내용 : 짭명방

### 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`

### 남은 작업 내용 
- 1-3 밸런싱, 보스 추가
- `Artillery` 2스킬 정리
- 테스트 및 수정

- 보고 참조할 내용
	- [[오퍼레이터들 스탯 정리]]
	- [[적 스탯 정리]]

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.

# 현재 작업 일지


## 250606

### 작업 중

### 작업 완료

- [x] 바리케이드 아이콘 추가
![[Barricade_256.png]]
- [x] 모든 `UnitEntity` 상속받는 객체에 대해 `DeathAnimation` 추가
	- 검정색으로 변함 -> 투명해짐
```cs
// UnitEntity.cs
    protected virtual void Die()
    {
        StartCoroutine(DeathAnimation());
    }

    protected IEnumerator DeathAnimation()
    {
        Renderer renderer = GetComponentInChildren<Renderer>();

        // 동일한 머티리얼을 사용하는 모든 객체에 적용되는 걸 막고자 머티리얼 인스턴스를 만들고 진행한다.
        if (renderer != null)
        {
            Material materialInstance = new Material(renderer.material);
            renderer.material = materialInstance;

            // DOTween 사용하여 검정으로 변한 뒤 투명해지는 애니메이션 적용
            materialInstance.DOColor(Color.black, 0f);
            materialInstance.DOFade(0f, 0.5f).OnComplete(() =>
            {
                Destroy(materialInstance); // 메모리 누수 방지
                OnDestroyed?.Invoke(this);
                RemoveAllCrowdControls();
                Destroy(gameObject);
            });
        }
    }
```
- 투명해지는 효과가 구현되지 않고 있음 : 머티리얼 자체가 투명도를 지원하지 않을 가능성이 있다.
	- 머티리얼의 `Surface Type`을 `Transparent`로 바꿔봤는데, **평소에는 `Opaque`였다가 죽을 때만 `Transparent`로 바꾸는 게 나아 보인다.** 관련 코드는 Claude한테 받아서 적용.
	- **검게 변하는 것도 이상하다.** 그냥 투명도만 서서히 낮추는 식으로 하겠음.


## 250605

### 작업 중

#### 1-0 ~ 1-3 정주행하면서 테스트
- `1-3` 스테이지 몹 추가 배치
- 보스는 아직 만들지 않았음

- [ ] 바리케이드 아이콘 추가

### 작업 완료

#### 발견한 이슈들 & 수정 사항
- [x] `BaseSkill`에서 `isOnTest` 필드를 하나 추가
	- 기존 스킬 비용을 나타내는 `SPCost`의 변수 이름을 `spCost`로 변경
	- 변수 `SPCost`는 `isOnTest`에 따라 1이거나 `spCost`를 가리키는 프로퍼티로 변경함.
	- 테스트하려고 SP Cost를 바꾸는 일이 있는데, 원본을 따로 기록해두지 않아서(!) 테스트할 때마다 변동사항이 생겼기 때문에 이렇게 바꿨다. 

- [x] `SPBar`의 회복되는 UI를 보면 낮은 수치일 때는 천천히 오르다가 높은 수치일 때는 빠르게 오르는 듯한 현상이 보임.
```cs
	// spFill.rectTransform.anchorMin = new Vector2(0, 0);
	// spFill.rectTransform.anchorMax = new Vector2(valueRatio, 1);
```
> **이 부분 주석처리로 해결**
> - 저렇게 설정하면 내부 Fill 값이 1이 됐을 때도 슬라이더를 가득 채우지 못함. 딱 중간쯤 온다.

- [x] 튜토리얼 진행 중 `OperatorSlot`을 제대로 찾지 못하는 문제 수정
	- `OperatorInventoryPanel` 에서 슬롯 초기화할 때 원래 슬롯의 이름도 같이 수정했는데 그 부분이 없어졌다. 그 부분을 수정.
	- **`Slot` 자체에서 초기화 시에 이름도 설정하도록 수정 완료**

- [x] `1-1` 스테이지 바리케이드 설명 추가

- [x] 3성 클리어가 아닌데 정예화 아이템이 지급되는 현상 수정
```cs
// 아래처럼 수정 완료
if (stars < 3) return 0f;
else
{
	// 이전에 3성 클리어를 한 적이 없으면 지급
	if (resultInfo == null || resultInfo.stars < 3)
	{
		return 1f;
	}
	throw new InvalidOperationException("FirstClearPromotionItemRate의 예상치 못한 동작");
}
```

- [x] `NormalBuffVFX`의 `Lifetime`이 짧아서 버프가 켜진 중간에 스킬이 잠깐 끊기는 현상이 있음
	- **`Lifetime`을 늘려서 수정 완료**
## 250604

### 작업 중

#### 1-3 스테이지 작업
- 몹 배치
- ~~보스 작업~~

### 작업 완료
#### 범위 이펙트 : 최초 실행 시에 나타나지 않는 문제
- 이슈를 정리해봤음. AI한테 던져주기도 했다.
```
1. 스테이지 씬에서 스킬을 최초로 실행했을 때 범위를 나타내는 이펙트가 나타나지 않음
	1-1. 이펙트가 나타나지 않은 상황에서 스킬이 종료되고, 해당 스킬을 다시 실행했을 때에도 똑같이 범위를 나타내는 이펙트가 나타나지 않음.

2. 스테이지가 종료되고 메인메뉴 씬으로 전환됐다가 다시 스테이지 씬으로 돌아갔을 때에는 범위 이펙트가 최초 실행부터 잘 나타남.

3. 유니티 엔진 상에서 아예 실행을 멈췄다가(Stop) 다시 시작(Play) 버튼을 눌렀을 때, 스테이지 씬에서 범위 이펙트는 최초 실행부터 잘 동작함.

4. 유니티 엔진을 아예 껐다가 켰을 때에는 다시 1번과 1-1번처럼 동작함.
```
- [[셰이더 컴파일, 리소스 초기화 지연 문제]]
	- **셰이더 컴파일, 리소스 초기화** 지연 이슈. 
	- 에디터에서는 셰이더가 필요할 때마다 실시간으로 컴파일되기 때문에 컴파일이 완료되지 않은 상태에서 실행되면 이펙트가 나타나지 않을 수 있음.
	- 즉 에디터 단계에서만 필요한 기능
	- **스테이지 로딩할 때 가지고 있는 이펙트나 파티클 시스템들을 모두 잠깐 실행하는 기능을 구현해봤는데 안됨**
	- 이펙트가 나타나지 않는 게 문제이지 기능은 똑같이 동작하니까 이건 해결하지 않고 보류함.
	- 아니면 스테이지 진입 전에 메인메뉴 씬에서 스킬들을 1번씩 실행시키는 것도 괜찮을 듯


#### 기타 이슈 수정
- [x] `Artillery`의 폭발 이펙트가 여러 개 발생하는 현상 수정
	- `Projectile`로 피격이 발생할 때 범위 콜라이더가 발생 -> 범위 콜라이더에 충돌한 적들에 대해 피격 이펙트가 다시 발생하기 때문
	- ~~`UnitEntity.TakeDamage`에서 `bool playGetHitEffect = true`을 추가, 특별한 경우에만 피격 이펙트를 끄도록 수정함.~~
		- 이건 아예 피격 이펙트가 재생되지 않을 거임
		- `Enemy.TakeDamage`가 `UnitEntity.TakeDamage`를 상속받지만, `base.TakeDamage()`를 이용하는 구조는 아님. 이 구조를 바꿔봐야 할 듯?
			- **`base.TakeDamage()`를 이용하도록 수정.**

- [x] `Enemy.TakeDamage` 부분에서 추가로 구현된 부분은 `Operator`가 공격했을 때 통계 패널을 업데이트하는 부분이었다. `base.TakeDamage` 내에 `virtual` 메서드를 이용하도록 넣고 `virtual` 메서드는 `Enemy`에서 구현하는 식으로 수정.
	- `OnDamageTaken(float actualDamage)`로 사용 중인 게 있었다. 이걸 **`OnDamageTaken(UnitEntity attacker, float actualDamage)`로 바꾸고** `UnitEntity`에선 비우고 **자식 클래스들인 `Operator`와 `Enemy`에서 각각 방어 통계 업데이트와 공격 통계 업데이트하는 식으로 수정함.**

- [x] `OperatorInventoryPanel`
	- **오퍼레이터를 선택해서 들어갔다면 해당 오퍼레이터의 슬롯은 현재 스쿼드에서 사용 중인 스킬로 초기화**돼야 함. 지금은 기본 지정 스킬이 나타난다.
```cs
// 자동으로 선택되는 오퍼레이터의 스킬은 현재 스쿼드에서 사용 중인 스킬로 선택됨
if (operatorToAutoSelect != null && opToShow == operatorToAutoSelect)
{
	selectedSkillIndex = GameManagement.Instance!.UserSquadManager.GetCurrentSkillIndex(operatorToAutoSelect);
}
// 나머지는 오퍼레이터의 기본 선택 스킬을 사용
else
{
	selectedSkillIndex = opToShow.DefaultSelectedSkillIndex;
}
```
> 단일 선택 상황에서만 이런 이슈가 있었고 다중 선택 상황은 스쿼드의 스킬 정보를 이미 가져오는 식으로 구현되어 있었음.

- [x] `TestManager`에서의 테스트를 위한 초기화 환경 개선
	- 모든 오퍼레이터의 정예화, 레벨을 지정해서 초기화하도록 기능 수정


## 250603

### 작업 중
### 작업 완료

#### 기존 스테이지 리밸런싱
- 오퍼레이터가 이제 7명이기 때문에
	- [x] 한꺼번에 동시에 배치 가능한 인원을 **5명 -> 6명**으로 올림
	- [x] 보상 재설정 
		- 기존의 1, 2, 3 구성 중 어디를 바꿀까?
		- 정예화 가능 숫자 기준 `1, 3, 3`이랑 `2, 2, 3` 중에 `2, 2, 3`으로 결정
- `1-1`, `1-2` 스테이지 재설정
	- 동시에 6명을 배치할 수 있게 되면서 스테이지의 난이도를 조금 더 올릴 수 있을 것 같다.
	- ...라고는 해도 큰 수정은 안해도 될 것 같은게, 지상에 배치할 수 있는 오퍼레이터가 3기밖에 없어서 엄청 쉬워진다는 느낌은 아니다. 약간의 적만 추가했음.
- `Tanker` 방어력 550으로 상향
#### SixBulletSkill
- 스킬이 켜진 직후에 공격 속도와 모션을 0으로 만듦(초기화)
	- 즉 평타가 나간 직후에 스킬이 켜지면 바로 공격이 1번 더 나갈 수 있는 구성
#### TestManager 생성
- `PlayerDataManager`에서 테스트를 위한 초기화 부분을 `TestManager`로 별도로 떼어내서 구현함
	- AI한테 던져서 초안만 만들고 디테일은 잡아나가는 식으로 작업
	- `TestManager`는 `GameManagement`의 자식 오브젝트로 처리
#### 기타 이슈 수정
- [x] `StageSelectPanel` : `PromotionItem` 갯수가 1개로 고정되어 표기되는 문제 수정

## 250602 

### 작업 중

### 작업 완료

#### Artillery 2번째 스킬 남은 작업 처리
1. 스킬이 켜졌을 때 UI가 제대로 나타나지 않음
2. 스킬이 켜졌을 때 공격 속도 타이머가 다시 돌아가는 느낌이 있고, 공격 속도가 빨라지는지도 모르겠음.

- 우선 UI 작업부터 다시 진행함.
	- **이전에 AI가 던져준 코드는 폐기. 이해도 안 되고 제대로 동작하지도 않음.**
	- 기존의 `HealthBar`는 `Slider` 내에 들어가는 `Filled` Image만 갖고 컨트롤하는 방식이었음
	- `AmmoMode`를 켜게 되면 `Horizontal Layout Group`을 사용하고 
		- 기존의 `HealthFill`은 비활성화한다.
	- 근데 이거를 `DeployableBarUI` 단위에서 구현할지 `OperatorUI`에서 구현할지가 좀 애매한 지점이 있음. 프리팹 연결도 해제해둔 상태다.
		- `DeployableBarUI`에서 사용할 요소가 아니긴 한데, 스크립트는 `DeployableBarUI`에 구현되어 있다는 게 문제인가?

- 대충 아래처럼 작업할 예정
1. 일단 `OperatorUI`에서 구현을 할 거임
2. `DeployableBarUI`는 `OperatorUI` 외에도 자체적으로 쓰이는 상황이 있기 때문에, 이 경우는 단순하게 인터페이스처럼 동작함
3. `SPBar`를 `HealthBar`와 분리해서 스크립트를 구현함. `SPBar`에만 추가로 쓰이는 것들이 있다.
	- `HealthBar`의 기능 일부를 떼서 `SPBar`에 붙였다. 예를 들면 체력이 감소하는 효과를 나타내는 `DamageOverlay`나 `ShieldFill` 같은 부분은 `SPBar`에서는 필요하지 않으니 제외하고, `SPBar`에만 들어가는 `Ammo` 관련 `UI`들이 추가되었음.

> 발생 중인 이슈
 - [x] 스킬 켜지면 ~~SP Bar의 게이지 높이가 훨씬 커지며~~ `ammoMode`로의 전환도 제대로 안 됨
- [x] 스킬 임의로 껐을 때 기존 이미지들 제거해야 함
- [x] 탄환이 나갈 때 오른쪽에 빈 공간이 나타나는 게 아니라 계속 화면을 가득 채운 채로 현재 갯수만 감소함
	- 이거는 아예 `Child Force Expand`로 설정할 게 아니라 너비값을 설정하고 줄여나가는 식으로 해야 할 듯? 
	- 너무 어렵게 생각했다. 탄환 갯수를 계속 넣는 방식으로 업데이트하니까, **이미지를 비활성화하는 게 아니라 알파값을 0으로 만들어서 투명하게만 만들면 됨.**
	- 패딩 값은 `0.02` 정도로 잡으면 얼추 보이는 걸 알고 있다.

- `DeployableActionUI`  : 스킬 버튼에 중지 표시 구현
## 250601

### 블렌더 공부
- [[블렌더 5. UV 매핑]]


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

## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
	- [[1. Projects, Ongoing/포폴 겸 블로그 만들기/월별 작업 기록/25년 4월|25년 4월]]
	- [[1. Projects, Ongoing/포폴 겸 블로그 만들기/월별 작업 기록/25년 5월|25년 5월]]
