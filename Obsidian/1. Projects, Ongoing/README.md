## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 데이터가 제때 수집되고 있는지 눈팅 정도만 하면 충분할 것 같다.
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음?

## 작업 예정

### 남은 작업 내용 
- 남은 작업들
	- 스테이지 1-0 ~ 1-3 밸런싱
	- 1-3에 보스 추가
	- 보상 설정

- 보고 참조할 내용
	- [[오퍼레이터들 스탯 정리]]
	- [[적 스탯 정리]]

- 테스트 및 수정

### 작업 중 
1. 스테이지 `1-1` 밸런싱 (이후 `1-2`, `1-3`)
	- **스테이지 완성, 보상 설정, 보스 추가** 등으로 게임을 완성하는 게 젤 중요함!! 다른 건 다 부차적인 요소! 
	- ...인데 잘 실천이 안 되고 있다. 계속 고쳐야 하는 게 보임..

### 구현 예정

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다 (= 오퍼레이터를 배치하지 않고 스테이지 클리어하기)
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 드는 품을 생각하면 미뤄도 될 것 같음. 


## 4월

## 250407

### 밸런싱과 스테이지 구성
- `Operator - Medic`의 경우, 공속이 너무 느려서 코스트 값을 못하는 느낌이 든다. 
	- 공속을 2초로 수정해봄
- 전체적인 스테이지 구성
	- 시작 코스트가 너무 낮음 : `10 -> 20`으로 상향
	- 코스트 회복 속도도 1.5로 높여봄 
	- 코스트 수급을 담당하는 뱅가드가 잘리면 게임 흐름이 확 답답해진다. 뱅가드 스킬의 효율을 높이는 건 안 그래도 뱅가드 의존성이 심해서 아닌 것 같고, 기본적인 스테이지 설정 자체를 조금 더 완화할 필요가 있겠음
	- 대신 급한 불을 끄는 용도, 즉 원래 게임에서의 `스페셜리스트 - 처형자` 직군 같이 재배치 시간을 짧게 하는 구성은 뱅가드에게 준다 - 재배치 시간 30초
		- 근데 사실 테스트에서 1초로 놨어도 뱅가드가 빠져있는 시간 자체가 코스트가 느릿하게 차서 뱅가드를 재배치할 경우가 잘 없기는 했음

#### 스테이지 1-1
- 위 웨이브, 아래 웨이브로 구분이 되어 있는데 위 웨이브가 끝나고 나서 아래 웨이브의 시작 사이에 10초의 여유를 뒀음(단순한 스폰 타이밍 조정)
- 시작 코스트 높이고, 10초 여유 주니까 깰 수 있기는 함.
- `EnemyRanged` 관련 : 육성이 되지 않은 `Sniper`를 3대만에 보낼 수 있는데, 이게 맞을까?
	- `Ranged`를 `1-2` 이후로 빼는 것도 방법일 듯?
> 일단 이 부분은 보류하고 통계 매니저, 패널 관련 현상들 수정함



### 사소한 것들 생각
- 경험치 아이템을 자동으로 사용하는 로직이 현재 서비스 중인 게임들하고 비교하면 조금 다른 듯 - 보통 **가장 좋은 가치를 지닌 경험치 아이템이 먼저 소비**되는 구조를 띄고 있는데, 내 프로젝트는 **경험치 낭비를 줄이기 위한** 방식으로 짜여져 있어서 가장 낮은 가치의 경험치 아이템이 먼저 소비되는 경향이 있음
	- 다음 레벨에 도달할 정도가 아니라면, 어떤 레벨에서 사용되는 경험치 아이템들은 레벨업이 동반되지 않으므로 스탯의 상승을 가져오지 않기 때문에 아이템을 사용하더라도 실질적인 이득이 없음. 낭비라는 표현은 조금 이상하긴 한데 이미 쓴 말이니까 냅둔다.

### 이슈 수정

- [x] `SquadEditPanel` -> 스테이지 진입 시 클릭 동작으로 인해 맵 등이 정상적으로 불러와지지 않는 현상.
	- 정확한 현상은 스테이지 로드 시에, 패널이 서서히 사라지는 중에 **스테이지 진입 버튼 위에서** 클릭 동작이 이뤄질 때, 스테이지 로드 메서드가 다시 동작하는 현상임
	- 시도한 방법들
		- 스테이지 로드 중 클릭 동작 막기 
			1) UI 시스템에 대한 동작이기 떄문에 `ClickDetectionSystem`에 들어가는 것보다 먼저 동작해야 함
			- 이걸 구현하는 방법은 여러 가지가 있음
				- [x] `EventSystem.current.enabled = false`로 해서 **모든 입력**을 막는 것
				- [ ] 화면의 가장 위에 투명한 `image.raycastTarget`을 설정하는 것
				- [ ] `button.interactable = false`로 놓는 것 등등
				- [ ] 해당 버튼이 클릭됐을 때, 버튼의 리스너를 제거하는 것
					- 이 방법이 번뜩이면서 생각났는데, 스테이지 입장 버튼을 클릭한 시점에서 예상치 못한 모든 동작을 막으려면 맨 위 방법이 가장 좋아 보여서 그걸로 선택함.
	- 수정) `LoadStage`의 시작에서 모든 입력을 차단하고, `InitializeStage` 직전에 모든 입력 차단을 해제함

- [ ] `StageResultPanel` -> 가끔 별점 애니메이션이 동작하지 않고 이미지가 작게 유지되는 현상이 있음
	- 정확히 어떤 시점에 문제가 발생하는지 확인할 수가 없다. 클릭을 미리해서도 아니고 스크롤을 해서도 아니고 ItemUIElement의 팝업이 나타나서도 아님.

#### 통계 패널 정렬 이슈

- [x] `StageResultPanel` - `HealingDone` 탭에서 `DamageDealt`로 탭을 전환할 때, 내림차순으로 올바르게 정렬되지 않는 현상이 있음. 왜 저기서만?
	- 정렬 로직 자체가 문제인지, 이걸 보여주는 로직 자체가 문제인지 모르겠는데..
```cs
// 기존 구현 
	for (int i = 0; i < statItems.Count; i++)
	{
		var (opData, value) = sortedStats[i];
	
		// 현재 위치의 StatItem이 올바른 오퍼레이터를 표시하고 있는지 확인
		if (statItems[i].OpData != opData)
		{
			// 올바른 위치의 StatItem 찾기
			var correctItem = statItems.Find(item => item.OpData == opData);
			if (correctItem != null)
			{
				// Transform 순서 변경
				correctItem.transform.SetSiblingIndex(i);
	
				// statItems 리스트 순서도 변경
				int oldIndex = statItems.IndexOf(correctItem);
				statItems[oldIndex] = statItems[i];
				statItems[i] = correctItem;
			}
		}
	
		// 통계 표시 업데이트
		statItems[i].UpdateDisplay(currentStatType, showingPercentage);
	}
```
> - 여기서 스왑 로직 구현할 때의 이슈임 - `a <-> b` 스왑 로직을 구현할 때, 원래 `temp`라는 변수를 따로 두는 식으로 하는데 지금처럼 구현하면 중간값이 제대로 보존되지 않음

> - 따라서, 아래처럼 `temp`를 따로 둬서 중간값을 보존하면 됨
```cs
				int oldIndex = statItems.IndexOf(correctItem);
				StatisticItem temp = statItems[i];
				statItems[i] = correctItem;
				statItems[oldIndex] = temp;
```

- 더 좋은 방법을 AI가 줬는데, `Find` 메서드 때문에 `O(n^2)`이라고 한다. 
```cs

        Dictionary<string, StatisticItem> itemDict = new Dictionary<string, StatisticItem>();
        foreach (var item in statItems)
        {
            itemDict[item.OpData.entityName] = item;
        }

        List<StatisticItem> newOrder = new List<StatisticItem>();
        foreach (var (opData, _) in sortedStats)
        {
            newOrder.Add(itemDict[opData.entityName]);
        }

        for (int i=0; i<newOrder.Count; i++)
        {
            newOrder[i].transform.SetSiblingIndex(i);
            newOrder[i].UpdateDisplay(currentStatType, showingPercentage);
        }
```

> 잘 적용됨.

 - 추가 수정
 - `StatItem` 값이 0일 때에는 해당 `StatItem`을 비활성화
 - 하단 퍼센티지 바는 `fillImage`를 이용했던 것을 `slider.normalizedValue`를 이용하는 쪽으로 수정.

#### 통계 매니저 + 패널 - 받은 피해량 수정

- [x] `StageResultPanel` - **결과 패널에 나타날 피해량 값은 실제로 들어간/들어온 대미지**여야 맞는 듯. 현재는 계산 전의 값이 나타나고 있다. 
	- 즉, **방어력 / 마법저항력이 고려된 받은 피해량**을 통계에 넣는다. 쉴드 같은 개념도 그냥 저 값에서 까기 시작하는 거니까, `오퍼레이터가 받은 피해량 = 실질 피해량`으로 계산하는 게 맞다고 봄
```cs
// Operator.cs
    public override void TakeDamage(UnitEntity attacker, AttackSource attackSource, float damage)
    {
        base.TakeDamage(attacker, attackSource, damage);
        StatisticsManager.Instance!.UpdateDamageTaken(OperatorData, damage);
    }

// UnitEntity.cs
    public virtual void TakeDamage(UnitEntity attacker, AttackSource attackSource, float damage)
    {
        if (attacker is ICombatEntity iCombatEntity && CurrentHealth > 0)
        {
            // 방어 / 마법 저항력이 고려된 실제 들어오는 대미지
            float actualDamage = CalculateActualDamage(iCombatEntity.AttackType, damage);

            // 쉴드를 깎고 남은 대미지
            float remainingDamage = shieldSystem.AbsorbDamage(actualDamage);

            // 체력 계산
            CurrentHealth = Mathf.Max(0, CurrentHealth - remainingDamage);

            Debug.Log($"남은 체력 : {CurrentHealth}, 들어온 대미지 : {remainingDamage}");
            OnHealthChanged?.Invoke(CurrentHealth, MaxHealth, shieldSystem.CurrentShield);

            PlayGetHitEffect(attacker, attackSource);
        }


        if (CurrentHealth <= 0)
        {
            Die(); // 오버라이드 메서드
        }
    }

```
> - 현재 이런 구조인데, `Operator.TakeDamage`에 통계를 갱신하는 로직이 있음. 들어온 대미지를 그대로 통계에 넣는 구조
> - 저걸 `UnitEntity.cs`에 넣기 위한 방법으로, **후크 메서드**라는 게 있다. 부모 클래스에서 가상 메서드를 정의하고, 자식 클래스에서 오버라이드해서 사용하는 방식. 예전에 스킬 구조 짤 때 사용했던 방법이기도 하니 잘 기억해두자.

- 따라서 아래처럼 수정함
```cs
// Operator.cs
// 부모 클래스에서 정의된 TakeDamage에서 사용됨
protected override void OnDamageTaken(float actualDamage)
{
	StatisticsManager.Instance!.UpdateDamageTaken(OperatorData, actualDamage);
}

// UnitEntity.cs
    public virtual void TakeDamage(UnitEntity attacker, AttackSource attackSource, float damage)
    {
	    // 아래에서 써야 해서, 정의 및 할당을 밖에서 함
        float actualDamage = 0f;

        if (attacker is ICombatEntity iCombatEntity && CurrentHealth > 0)
        {
            // 방어 / 마법 저항력이 고려된 실제 들어오는 대미지
            actualDamage = CalculateActualDamage(iCombatEntity.AttackType, damage);
		
        }

        OnDamageTaken(actualDamage);
        
        // ...
	}

	// 후크 메서드 정의
    protected virtual void OnDamageTaken(float actualDamage) { }
```
## 250405 
### Barricade 퇴각 버튼이 동작하지 않는 현상
- [x] 해결

- `Operator`에서 나타날 때는 잘 작동하는데, `Barricade`에서 나타날 때는 잘 동작하지 않는다. 버튼 인식 자체가 안되는 것 같음.
- 리스너 등록까지는 잘 되는데? 버튼을 클릭하면 아무 동작이 없다는 개념에 가깝다. 
- `deployable.Retreat`이 구현되지 않았다고 보기도 어려움. `DeployableUnitEntity`의 자식으로 `Operator`나 `Barricade`가 있으니까.

- **이거 뭐가 문제지????**
	1. `Vanguard`로만 테스트해봤다. 다른 오퍼레이터는? -> 모두 잘 적용됨(퇴각 버튼)
	2. 아무리 생각해도 답이 안 나온다. **관련 요소들을 껐다 켰다 해보면서 어디서 원인이 발생하는지 찾아야 할 것 같음.** 지금 스크립트에서 **유니티에서 UI 클릭을 인식하는 동작**의 차이가 발생하는 부분이 아무리 봐도 없어..
		1) 버튼에 리스너를 등록하는 부분 주석 처리 : 이 메서드의 유무에 관계없이, 오퍼레이터에게 나타난 `DeployableActionUI`는 클릭했을 때 버튼이 클릭되면서 색상의 변화가 나타나는 반면, `Barricade`에서 나타난 `DeployableActionUI`는 아무런 변화가 없음
			- 이건 `button.interactable`의 문제도 아님) `interactable = false`이면 최초에 버튼 색깔이 더 흐릿하게 보이기 때문
		2) 스테이지 씬을 띄운 상태에서, `DeployableActionUI` 프리팹에 있는 걸 씬에 옮긴 다음 퇴각 버튼을 클릭해봄 -> **클릭 잘 됨.(위에서 말한 클릭 시의 색깔의 변화가 나타남)**
			- 즉, **`Barricade`에서만 알 수 없는 이유로 유니티에서 동작하는 버튼의 클릭 인식 동작이 작동하지 않는 현상이 발생**하고 있는 것. 
		3) `DeployableActionUI`라는 요소는 `DeployableManager`에서 프리팹으로 할당되어 있으며, 호출될 때마다 인스턴스화 되고 없어질 때마다 파괴된다. 
	3. 보통 이 정도 헤매면 뭐가 문제인지 감이라도 잡히는데 진짜 감이 하나도 안잡힌다. 이런 상황에서는 AI도 도움이 안 됨.
		1) 분기가 문제인가 싶어서 `Initialize`의 분기 부분을 모두 주석처리해도 똑같이 `Operator`에서는 동작, `DeployableUnitEntity`에서는 동작 안 함.
		2) `Operator.OnClick`과의 차이가 문제인가 하면, `DeployableUnitEntity.OnClick`으로 동작하는 게 `Barricade`이고, `Operator.OnClick`은 위의 OnClick을 상속받고 공격 범위만 하이라이트되는 게 전부임. 
		3) 심지어 `Initialize`의 모든 동작을 비활성화했을 때도 5번이랑 동작이 똑같음
		4) 심지어x2 `DeployableActionUI`의 모든 스크립트를 비활성화하고 `Initialize`만 남겼을 때에도 5번이랑 동작이 똑같음
	- **따라서 일단 `DeployableActionUI` 자체의 문제는 아님**

- 이들을 인스턴스화하는 `DeployableManager`로 간다.
```cs
    public void ShowActionUI(DeployableUnitEntity deployable)
    {
        InstanceValidator.ValidateInstance(actionUIPrefab);

        HideUIs();

        // 일관된 위치 구현하기
        Vector3 ActionUIPosition = new Vector3(deployable.transform.position.x, 1f, deployable.transform.position.z);
        currentActionUI = Instantiate(actionUIPrefab, ActionUIPosition, Quaternion.identity);
        currentActionUI.Initialize(deployable);
        currentUIState = UIState.OperatorAction;
    }
```
> 솔직히.. 별 차이가 없어... 위치 떄문에 막히는 것도 아니고..
- `Barricade`을 다른 스테이지에 넣어도 동일하게 동작함

- **발견한 듯....** **`InStageInfoPanel`의 `CancelPanel` 동작이 문제였다.**
```cs
    protected virtual void ShowActionUI()
    {
        DeployableManager.Instance!.ShowActionUI(this); // 무죄
        UIManager.Instance!.ShowDeployedInfo(this); // <-- 이 내부가 유죄
    }
```
> - `ActionUI`가 나타날 때, 해당되는 오퍼레이터의 정보를 함께 보여주는 `InStageInfoPanel`이 있다.
> - 이 패널의 구성은 `CancelPanel, GradientPanel, OperaotrInfoContent`인데, 이 중 `CancelPanel`, 즉 현재의 `ActionUI`가 나타난 상태를 취소하는 동작을 하는 패널의 영역은 (기본 투명이지만) 시각화해보면
![[Pasted image 20250405211229.png]] 
 화면 전체를 덮는다.
 이 패널을 만든 이유는 **하단의 오퍼레이터 박스를 클릭한 상태에서의 배치 동작을 취소하기 위함**이었는데 이게 현재 배치된 `Deployable`을 클릭할 때에도 `CancelPanel`이 비활성화되지 않은 게 일련의 사태의 원인으로 보임. 
 
```cs
    public void UpdateInfo(DeployableManager.DeployableInfo deployableInfo)
    {
    // ...
    
        // CancelPanel 부분만 따로 떼어놨음
        if (!currentDeployableUnitState.IsDeployed)
        {
            cancelPanel.gameObject.SetActive(true);
            cancelPanel.onClick.AddListener(OnCancelPanelClicked);
        }
        
    // ...
    }
```
> 결국 여기가 문제인건데, `UpdateInfo`를 호출하는 부분에서 "현재 배치된 요소를 클릭했는가"를 나타내는 변수를 파라미터에 하나 넣었다.

- 수정
```cs
    public void UpdateInfo(DeployableManager.DeployableInfo deployableInfo, bool IsClickDeployed)
    // ...

        // 박스에서 꺼내는 요소인 경우, 맵의 남은 부분을 클릭하면 현재 동작을 취소할 수 있음
        if (!IsClickDeployed)
        {
            cancelPanel.gameObject.SetActive(true);
            cancelPanel.onClick.AddListener(OnCancelPanelClicked);
        }
	// ...
```

그리고 이들을 사용하는 상황은 아래처럼 이용한다.
```cs
    // 배치되지 않은 유닛의 정보 보기 동작
    public void ShowUndeployedInfo(DeployableManager.DeployableInfo deployableInfo)
    {
	    // ...
        inStageInfoPanelScript.UpdateInfo(deployableInfo, false);
		// ...
    }


    // 배치된 유닛의 정보 보기 동작
    public void ShowDeployedInfo(DeployableUnitEntity deployableUnitEntity)
    {
	    // ...
        inStageInfoPanelScript.UpdateInfo(deployableUnitEntity.DeployableInfo!, true);
        // ...
    }
```

마지막으로, **4개의 상황**(오퍼레이터 박스에서 꺼내기, 배치된 오퍼레이터 클릭하기, 바리케이드 박스에서 꺼내기, 바리케이드 클릭하기)를 모두 테스트해봤고, 최초의 문제였던 바리케이드의 Retreat 동작을 포함해 **모두 원하는 목적에 따라 잘 돌아가는 걸 확인**했음. 

> 왜 그런지 문제 발견하는 데에만 4시간 정도를 썼는데 해결은 10분만에 했다. 

- [x] `StageResult - StatItem` 레이아웃
	- `StageResultPanel`의 `StatItem`의 너비, 크기는 `Grid Layout Group`에서 컨트롤이 된다. `StatItem`에서 컨트롤하지 않음.
	- 수정) `Grid Layout Group`에서 서식 조정

### 사이트 폭파 대비
- 구현했던 것들을 이미지로 캡처해서 저장해둠
	- 게시판, 자기소개, 진행했던 프로젝트, 메인 페이지 등등
## 250404 - 짭명방
### 이것저것 수정

#### 테스트 용도) 스테이지 클리어와 보상 묶기
- [x] 완료
- 보상과 스테이지별 난이도 테스트를 위한 메서드 개선
	- 테스트용) `PlayerDataManager`에서 **특정 스테이지를 클리어한 것으로 처리했을 때, 해당 스테이지의 보상도 함께 들어오게 구현**함
	- 보상 자체는 `StageData`에 있는데 `Stars`에 따른 보상의 정도를 설정하는 로직은 `StageManager`에 있음. 이걸 `PlayerDataManager`로 옮겨야 하나?
	- **`RewardManager`를 별도로 구현함**
```cs
    public void SetAndGiveStageRewards(StageData stageData, int stars)
    {
        // 어차피 바꿀거라서 변수명은 약어 처리함
        var (fcr, bcr) = SetStageRewards(stageData, stars);

        // 리스트로 바꿔서 전달해야 직렬화 시에 더 안전하다고 함
        List<ItemWithCount> firstClearRewards = new List<ItemWithCount>(fcr);
        List<ItemWithCount> basicClearRewards = new List<ItemWithCount>(bcr);

        GameManagement.Instance!.PlayerDataManager.GrantStageRewards(firstClearRewards, basicClearRewards);
    }
```
> 다른 메서드들도 옮겨왔는데, 핵심 메서드는 이거임 
> - `PlayerDataManager`에서 기록을 `StageId`로 하고 있는데, 이걸 `StageData`로 넣게 바꿀 수 있나? 이게 되려면 최소한 `StageData`들의 리스트를 갖고 있어야 하지 않나?
> - 그래서 `StageDatabase`도 `GameManagement`의 자식으로 둔다. 
- 수정 과정에서 원래는 `Stage`마다 초기화되는 `StageManager`에서 전역적인 필드로 가졌고, 한 번 설정된 필드를 그대로 가져오기만 하면 되는 구조였다. 이게 살짝 바뀌어야 하나?
	- 패널에서 사용해야 하는 상황이 있기 때문에, `RewardManager`의 결과값 -> `StageManager`에서 `IReadOnlyList<ItemWithCount>` 전역 필드를 갖고 패널에서는 `StageManager`에 있는 값을 가져오는 방식으로 구성함

#### StageItemInfoPopup
- 스크롤 동작 안함. 텍스트 들어간 다음 레이아웃 갱신 필요.
```cs
DetailBox(Scroll Rect, viewport로 설정)
- ItemDescription(Content, content size fitter, TMP)
```
1. `scrollRect`을 스크립트에 추가, `text`가 들어간 다음 스크롤을 `normalizedVerticalPosition = 1.0f(맨 위)`으로 설정
2. `content size fitter`은 자식 오브젝트 뿐만 아니라, 이를 가진 오브젝트 자신의 컴포넌트까지도 포함해서 계산함

- `StageItemInfoPopup(이 오브젝트)`가 `Inactive`라는 에러가 뜬다.
```cs
    private void Awake()
    {
        // GetComponent 계열은 Awake에서 수행한다.
        canvas = GetComponentInParent<Canvas>();
        canvasRectTransform = canvas.GetComponent<RectTransform>();
        popupRectTransform = popupArea.GetComponent<RectTransform>();

        originalItemNamePosition = popupItemNameBackground.anchoredPosition;

        gameObject.SetActive(false);
    }

    public void Show(ItemUIElement itemUIElement)
    {
        gameObject.SetActive(true);

        this.itemUIElement = itemUIElement;
        ItemData itemData = itemUIElement.ItemData;
        popupItemNameText.text = itemData.itemName;
        popupItemDetailText.text = itemData.description;

        StartCoroutine(RebuildLayoutAndScrollToTop()); // <-- 여기서 에러 발생 중
    }
```
> ????
> - 오브젝트가 활성화되지 않았을 때 코루틴이 실행될 수 있다고 함(위에 `gameObject.SetActive`가 있음에도)
> - 코루틴으로 빼지 않고, `ActivatePopupElements`에서 팝업을 활성화하기 전에 위치를 `1.0`으로 잡아주는 식으로 수정했음
> - `StageItemInfoPopup` 자체가 좀 꼬여 있음. `BackArea`와 `PopupArea` 2개가 있는데, `PopupArea`는 내부 설정을 마치고 나타나도록 설정이 되어 있기 때문에 관련 요소들이 설정된 후에 활성화된다. 

- 마무리로, `StageItemInfoPopup` 레이아웃(패딩, 너비 등등)을 정리

### 지식이 늘었다

1. 직렬화(저장) 클래스 구성 시 유의할 점(`JsonUtility`을 사용할 때의 유의할 점)
- **`Unity` 엔진에 의한 오브젝트들`ScriptableObject, UnityEngine.Object 상속`을 직접 할당하려고 한다면, 유니티의 직렬화 시스템은 이들의 참조를 저장하지 못하므로 문제가 발생할 수 있다.** 
- 그래서 웬만하면 순수 데이터들`int, string, List 등등`을 저장하는 쪽으로 권장함. 실제로도 기존에는 `stageId`라는 `string` 값으로 스테이지 클리어 여부를 저장했다.

2. `Content Size Fitter` - 자식 오브젝트 외에도, 자기 자신 오브젝트의 컴포넌트들에 의한 레이아웃 변화까지 다 반영해서 너비/높이 값이 설정됨.

## 250403 - 짭명방

### 이것저것 수정
- [x] `OperatorInventoryPanel, OperatorDetailPanel` : 숫자 폰트 출력이 일관적이지 않음
	- `LiberationSans SDF`가 기본 설정으로 되어있었는데, 일부 숫자는 해당 폰트의 숫자가 출력되지 않음. 프로젝트 설정이랑 충돌하는 듯 하다.
	- 수정) 해당 부분들을 프로젝트에서 사용하는 기본 폰트인 `SUITE - Regular SDF`으로 변경
		- 이전에 버전 업데이트를 하면서 `SUITE - Regular SDF Test`를 만들어서 썼던 적이 있는 걸로 기억하는데, `Test`가 아니라 그냥 버전으로 사용하는 게 맞는 듯.

- `OperatorInventoryPanel`
	- [x] 패널 진입 시 `ClearSideView` 한 번 실행 추가
		- `ClearSideView` 내부에 인디케이터 보이지 않게 하기 추가


- [x] `StageResultPanel`
	- `StageId` 아래에 왜 스테이지 설명 부분이 들어가 있지?? `StageName`으로 수정

- `OperatorDetailPanel`
	- [x] 0정예화에서 2번째 스킬 이미지가 `nullImage`가 나타나지 않는 현상 수정 
		- `Awake`에서 `noSkillIcon` 변수에 스프라이트가 할당되고(프리팹에 기본으로 설정되어 있음)
		- `OnEnable`에서 `ResetSkillIcon`이 실행되고 `UnlockedSkill`이 1개니까 `noSkillIcon`이 나타나야 하는데 아무 스프라이트도 할당되지 않은 상태가 됨.
		- 오늘 수정하는 사항들 중에서 가장 이해가 안되고 있음
![[Pasted image 20250403191431.png]]

- 갑자기 된다? 왜 될까?
![[Pasted image 20250403201350.png]]


- `OperatorPromotionPanel`
	- [x] 아이템을 갖고 있지 않은 경우, `ItemUIElement`을 살짝 흐릿하게 처리하는 건 어떤가?
		- 흐릿한 처리보다는, 아이템 갯수가 모자라다면 단순히 해당 `ItemUIElement`의 중앙에 `ImageNotEnough`을 띄우는방식으로 구현함
		- `ShowNotEnough` 이미지가 뜨지 않는 현상이 있었는데, `OnEnable`에 해당 이미지를 `SetActive(false)`로 했기 때문이다. 따라서 `OperatorPromotionPanel`에서 `초기화 -> 활성화`가 아닌, `활성화 -> 초기화`로 메서드 실행 순서를 바꿈.

![[Pasted image 20250403185751.png]]

- [x] `OperatorLevelUpPanel`
	- 50레벨에 도달했을 때 현재 경험치는 0으로 고정해야 함, 지금은 남는 값이 있음
	- 사실 패널의 문제라기 보다는 시스템 부분을 만져야 함
	- `ExpCalculationSystem`의 `CalculateOptimalItemUsage`의 남은 경험치 계산 부분
```cs
        // 현재 정예화의 최대 레벨에 도달했다면 현재 경험치는 0으로 설정
        if (finalLevel == OperatorGrowthSystem.GetMaxLevel(phase))
        {
            optimalPlan.remainingExp = 0;
        }
        else
        {
            optimalPlan.remainingExp = remainingExp;
        }
```

> Before
![[Pasted image 20250403190438.png]]

> After
![[Pasted image 20250403190540.png]]

#### 오퍼레이터 디테일 패널에서 인벤토리 패널로 빠져나올 때, 디테일 패널의 오퍼레이터 클릭 상태 유지하기
- [x] 완료

- `OperatorDetailPanel`에 들어갔다가 나올 때, 현재 선택된 오퍼레이터를 유지하는 게 좋을까? 
	- 즉, 인벤토리 패널 -> 디테일 패널 -> 인벤토리 패널로 돌아올 때 현재 커서를 유지할까 여부를 결정하는 것임
	- 스쿼드 편집 상황에서의 인벤토리 패널 -> 디테일 패널로 들어가는 이유에 대해 생각해보면 되지 않을까?
		- 오퍼레이터의 성장이 주 목적이 될 것
		- 그러면 성장을 시키고 다시 인벤토리 패널로 나왔을 때, **커서가 초기화되는 것보다는 유지되는 게 더 자연스러운 동작**이지 않을까?  성장시킨 오퍼레이터를 스쿼드에 즉시 편성할 수 있게 하는 거니까
		- 반대로 유지하지 않을 이유는 딱히 생각나지는 않는 듯. 스쿼드 편집 상황이 아니라면, "오퍼레이터를 선택하는 걸 유지하는 상황" 자체가 필요 없다. 오퍼레이터 슬롯을 클릭하면 바로 디테일 패널로 넘어가기 때문에.
	- 따라서 선택된 오퍼레이터를 유지하는 것으로 결정. 
	- 그러면 이 로직을 어디서 구현하는 게 좋을까?
		- `MainMenuManager` - 가장 먼저 생각나는 후보. 메인메뉴 씬에서 유지해야 하는 정보이기 때문에.
		- `UserSquadManager` - 얘는 스쿼드 관련이고..
		- `PlayerDataManager` - 얘도 저장된 정보를 가져오는 개념에 가까움.
	- `MainMenuManager`로 결정.
		- `CurrentEditingOperator`이라는 프로퍼티를 만들고 세터 메서드도 만듦
		- `OperatorInventoryPanel`에서 디테일 버튼을 눌렀을 때 `CurrentEditingOperator`를 저장
		- `OperatorInventoryPanel.PopulateOperators` : 다시 `OperatorInventoryPanel`로 나왔을 때 `CurrentEditingOperator`가 있는지를 검사해서 있다면 리스트의 맨 앞으로 보냄
			- 리스트의 맨 앞으로 보낼 때 `CurrentEditingOperator`는 `null`이 된다.
			- 기존의 `existingOperator`, 즉 `squadEditPanel`에서 들어왔을 때 가장 앞으로 오는 오퍼레이터 슬롯은 `else if`문으로 설정해서 2순위로 그렇게 동작하게끔 변경함.

	- [x] 이슈) `editingOperator`가 2번 나옴
		- 수정) `PopulateOperators`에서 조회된 오퍼레이터 중, 리스트에 있는 기존 오퍼레이터를 삭제하고 다시 맨 앞에 넣음 

	- [x] 이슈2) `existingOperator`로 들어간 다음, `existingOperator`와 다른 `editingOperator`를 편집하고 나왔을 때 `existingOperator`가 노출되지 않는 현상
		- `PopulateOperator`에서 `if - else if`문으로 구현했기 때문으로 보임
		- 수정) 순서 : **`existingOperator`가 있다면 우선 0번에 넣고, 그 다음에 `editingOperator`가 있으면 얘를 0번에 넣음**
			-> `existingOperator`는 자연스럽게 1번으로 밀려나는 구성
## 250402 - 짭명방

### 뭔가 허전한 요소들 채우는 중 
- `StageResultPanel` - `StatsContainer`에 제목 텍스트 추가
- `StageSelectPanel` - 스테이지 클릭 시 나타나는 우측 패널에 `StageName`도 추가
- `SkillData`로 생성한 스킬들의 `Description` 정리 
- `OperatorDetailPanel`
	- 스킬 이름 표시 : 스킬 아이콘 아래에
	- 스킬 아이콘 클릭 동작 변경
		- 현재 동작) 스킬 아이콘 클릭 시 - 즉시 해당 스킬을 디폴트 스킬로 지정
		- 하지만 스킬 아이콘을 클릭하는 이유 중에는 단순히 그 스킬에 대한 설명을 보고 싶기 때문일 수도 있다
		- 따라서 원본 게임처럼 **스킬의 설명 아래에 `디폴트로 지정`하는 별도의 버튼을 하나 추가**함
	- 원본 반영하기 : 스킬 아이콘 우측 하단에 지속시간 항목 추가
		- 소모 SP는 표시하지 않음
		- 이거를 구현할거면 아예 `SkillIcon`이라는 항목을 따로 만드는 것도 방법이겠는데?
	- `SkillIconBox.cs`을 따로 만들었음. 
		- `SkillSelectionIndicator`는 아예 패널의 요소로 빼고, 어떤 스킬을 선택하면 해당 `SkillIconBox`의 1번째 자식으로 배치시킴(지속 시간 박스와 겹치면 보이지 않게 렌더링)
		- 이를 사용하는 코드랑 레이아웃도 일부 수정. 
		- 원본 게임은 지속 시간 아이콘 옆에 숫자만 넣었는데, 나는 `s`까지 넣어서 뭘 나타내는지 조금 더 뚜렷하게 해보겠음
		- 스킬 아이콘이 들어가는 요소들을 다 이걸로 대체해야 하나? `OperatorInventoryPanel`까지는 스킬 지속시간이 보이는 게 좋아보이니 수정하고 마무리. 
		- `OperatorSlot`이나 인게임 스킬 아이콘에서도 넣어야 할까? 굳이 그럴 필요는 없을 듯?

- `OperatorDetailPanel`에 넣음
![[Pasted image 20250402220917.png]]

- `OperatorInventoryPanel`에 반영
![[Pasted image 20250403002824.png]]

### 지식이 늘었다
- `Content Size Fitter` 등에 의한 너비/높이 변화 즉시 반영하기
```cs
// ContentSizeFitter에 너비 변화 반영
LayoutRebuilder.ForceRebuildLayoutImmediate(durationBox.rectTransform);
```
> 관련된 항목을 설정 `ex) 텍스트 할당` 하고, 레이아웃을 즉시 업데이트를 한 번 해준다. 

- **`OperatorInventoryPanel`을 활성화한 채로 테스트에 들어갈 때, `UIHelper`보다 `Awake` 메서드가 먼저 동작**하는 이슈
	- 일단 `UIHelper`를 이용하는 메서드와 `UIHelper`의 싱글턴 인스턴스를 설정하는 메서드 모두 `Awake`에 있음
	- **2개의 활성화된 오브젝트의 `Awake`의 순서는 명확히 보장되지 않음**
		1. `Project Settings - Script Execution Order`에서 설정한 순서
		2. 하이어라키에서의 위치 : 상단의 오브젝트가 먼저 실행되는 "경향"이 있음
		3. 오브젝트 생성 시점 : 먼저 생성된 오브젝트가 먼저 실행되는 "경향"이 있음
	- 비활성화된 경우는 호출되지 않다가, 활성화될 때 1회만 호출됨. 재활성화 시에는 호출되지 않음.
	- 일단 이 문제는 해당 메서드를 `OnEnable`으로 빼는 것으로 바꿨음. 패널이 켜질 때마다 실행되긴 해서 좋은 건 아니겠지만..



## 250401 - 짭명방

### OperatorPromotionPanel 
- 좌측 하단에 사용 아이템 표시
	- 처음에는 아이템이 들어가는 부분에 대한 이미지를 검정 -> 투명 그라데이션 으로 넣어보려고 했는데, 그냥 아무 것도 없는 곳에 아이템이 나타나게 하는 것도 괜찮은 듯?

- 추가로 `OperatorData`에 `Promotion`에 필요한 아이템을 추가함
	- 실제로 정예화에 `PromotionItem 1개` 외의 아이템을 추가할 생각은 없어서 이렇게 작업할 필요는 없다. `OperatorPromotionPanel`의 좌측 하단을 고정시키면 되니까
	- 그래도 개발을 이것만 하고 말 게 아니니까 실제로 데이터를 가져와서 나타내는 것까지는 작업을 해야 공부가 될 것 같다. 
```cs
// OperatorData
    [Header("Promotion Required Items")]
    public List<PromotionItems> itemDatas = default!; 

    [System.Serializable]
    public class PromotionItems
    {
        public ItemData itemData;
        public int count;
    }
```

- 이렇게 수정한 부분을 바탕으로 아래의 과정들이 진행됨
1. `OperatorPromotionPanel`의 좌측 하단에 `ItemUIElement`을 초기화하는 것
	- 처음에는 인스턴스화 -> 파괴를 반복했는데, 그냥 미리 `ItemUIElement`들을 넣어놓고 필요한 만큼만 활성화를 하는 방식으로 수정.
2. 정예화가 진행되는 로직도 `PromotionItems`에 있는 아이템들을 사용하게 만들기

- 코파일럿) 이 때 아이템을 갖고 있는지를 점검하는 메서드는 `PlayerDataManager`에 구현하는 게 더 맞다고 한다. 정예화 시에 갖고 있는 아이템을 검사하는 거니까 `OperatorGrowthSystem`이 맞을 것 같았는데, 아이템을 정리하는 로직 자체가 `PlayerDataManager`에 있기 때문에 거기서 진행하는 게 맞다고 함. 
	- 클로드한테도 물어보니까 비슷한 답변이 돌아왔음. 재활용한 코드가 갯수는 반영되지 않았었는데, 이것도 고쳐줬다. 무서운 시대.

### 지식이 늘었다
1. `IEnumerable.ToDictionary`라는 메서드가 있다. 컬렉션을 딕셔너리로 바꿀 때 유용할 것 같음.
```cs
var itemsToUse = op.OperatorProgressData.promotionItems.ToDictionary(
	promotionItem => promotionItem.itemData.itemName,
	promotionItem => promotionItem.count);
```
> `promotionItems`가 더 여러 개의 필드를 갖고 있더라도, 그 중에서 2개의 필드만 뽑아서 키, 밸류 쌍으로 묶을 수 있다고 한다. 

2. `ItemUIElement`들을 생성 -> 파괴에서 활성화 -> 비활성화로 바꿨다. `ItemUIElement`들을 리스트로 관리하고 관련 오브젝트들을 인스펙터에서 집어넣는 방식으로 구현하고 있다.
- 이 때, **인스펙터에서 리스트의 길이를 5로 하고 일부 값이 비어있다면 리스트의 길이는 여전히 5이다.** 즉,  내부 원소가 null이더라도 리스트는 null인 원소를 갖고 있는 상태가 된다. 이거는 값이 들어갔다가 나간 상태인 `Missing`도 마찬가지.
> 활성화 - 비활성화로 오브젝트들을 다룰 때, 개발 과정에서 일부 오브젝트를 만들거나 없애는 경우에 주의해야 할 듯. 오늘 같은 경우도 5개를 쓰려다가 3개로 수정했는데 리스트는 (지금 보면 당연히) 길이 5로 유지됨.

# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `Archive` 폴더로 들어갔을 듯.
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

## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
