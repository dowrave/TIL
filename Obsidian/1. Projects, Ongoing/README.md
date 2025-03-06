## 참고
- **옵시디언으로 봐야 멀쩡하게 보임!!!**
- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 블로그
- 데이터가 제때 수집되고 있는지 눈팅 정도만 하면 충분할 것 같다.
- 사실 Quill을 쓰면서 되돌리기 / 붙여넣기 기능이 좀 이상하게 작동하고 있는 문제가 있긴 한데... 언제 해결할지는 모르겠음?

## 작업 예정

### 진행 중
- 스테이지 1-0 ~ 1-3 밸런싱
- 튜토리얼 구현해야 할 듯.
- 스테이지 클리어 후 경험치 / 정예화 아이템 지급 구현하기
### 부차적인 이슈
-  `FloatingText`가 나타났으면서 오퍼레이터 클릭 등으로 카메라 각도가 돌아간 상태일 때, `FloatingText`의 각도가 틀어지는 현상이 있음

### 구현 예정
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다 (= 오퍼레이터를 배치하지 않고 스테이지 클리어하기)
### 발생 중인 이슈
- (241216) 레벨업 후에 `OperatorLevelUpPanel`의 스냅핑 동작이 정상적으로 동작하지 않을 때가 있다. 
- 간헐적인 문제
	-  `Enemy` 기준, 바리케이드 파괴 로직이 동작하지 않는 현상이 있음

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 


## 3월

### 250307

#### 어제 몬한 것
- [ ] 뱅가드 스킬 - 코스트 회복 파티클 : 우측 하단에 너무 가까울 경우, 생존 시간이 남아있기 때문인지 파괴됨에도 다시 생성되는 문제

### 250306
#### 수정 사항
- [x] 스테이지 클리어 실패했는데도 `StageResultPanel`에 보상이 나타나는 문제
	- 클리어했을 때 잘 나타남
	- 클리어하고 다시 시도해서 실패했을 때 안 나타남
	- 그냥 실패했을 때 안 나타남
- [x] ItemUIElement 관련 수정 1 - 아이템의 우측에 표시되는 `ItemDetailPanel`의 오른쪽 끝이 화면에서 벗어날 경우, 화면의 왼쪽에 나타나게 하기
	- 이런 메서드를 자세하게 몰라서 AI한테 물어볼 때는, 메서드가 어떻게 동작하는지에 대해서만 알고 직접 만지는 게 훨씬 좋아보인다.
		- 예를 들면 (현재는 Gemini 2.0 Pro를 사용 중) 이거 수정할 때 캔버스의 월드 좌표, 스크린 좌표 등등의 개념을 말하면서 코드를 던져주고 이를 기반으로 수정했지만, 실제로 적용해보면 `ItemUIElement`을 기준으로 좌-우 전환을 하는 게 아니라 패널이 어디로 튀어나가 있음
		- 근데 단순히 `GetWorldCorners` 메서드를 캔버스와 패널에 적용하고, 오른쪽에 있는 `2`번이나 `3`번의 `x`좌표를 얻어서 비교하는 것만으로도 충분했음
```cs
    // 디테일 패널이 화면에서 잘리는 경우 패널을 왼쪽으로 띄우는 메서드
    private void AdjustDetailPanel()
    {
        // 월드 좌표를 쓰는 이유 : 캔버스 내의 절대적인 위치 개념임. 스크린 좌표는 해상도에 따라 값이 달라질 수 있다.

        // detailPanel의 월드 좌표의 우측 가장자리 계산
        Vector3[] detailPanelCorners = new Vector3[4];
        detailPanelRectTransform.GetWorldCorners(detailPanelCorners);
        float detailPanelRightEdgeWorldX = detailPanelCorners[2].x; // 우측 상단 코너의 x좌표

        // 캔버스의 월드 좌표에서의 우측 가장자리 계산
        Vector3[] canvasCorners = new Vector3[4];
        canvasRectTransform.GetWorldCorners(canvasCorners);
        float canvasRightEdgeWorldX = canvasCorners[2].x; // 캔버스 우측 상단의 x좌표

        // 디테일 패널의 오른쪽 변이 화면에서 벗어나는 경우, 왼쪽으로 디테일 패널을 옮김
        if (detailPanelRightEdgeWorldX > canvasRightEdgeWorldX)
        {
            Debug.Log("detailPanelRightEdge가 canvasWidth보다 크다");
            SetDetailPanelLeft();
        }
        else // 좌측으로 넘어가지 않더라도 detailPanel의 위치를 잡아준다
        {
            SetDetailPanelRight();
        }

        // 실시간으로 위치를 변화시킬 게 아니라서 이 정도로만 구현함
    }
```

- [x] ItemUIElement 관련 수정 2 - `BackArea` : 기존에는 `3000, 2000`의 너비, 높이를 갖는 이미지가 패널이 나타날 때 패널 뒤를 가리는 역할을 했는데, 이게 나타나는 위치에 따라 화면을 온전히 가리지 못함
	- UI 상에서 구현하고자 한다면 Canvas의 자식 오브젝트로 관리하는 게 최선이기는 한데, 얘는 렌더링 순서를 고려해야 함
	- 위 문제를 해결하면서 `ItemUIElement`가 이제 `Canvas`를 참조하도록 변경했으니, 여기서도 `BackArea`가 나타날 때 캔버스의 너비와 높이를 모두 활용하는 방식으로 수정함
```cs
        Rect canvasRect = canvasRectTransform.rect;
        RectTransform backAreaRectTransform = detailBackArea.GetComponent<RectTransform>();
		Debug.Log($"{canvasRect.position}");
        backAreaRectTransform.position = -canvasRect.position;
        backAreaRectTransform.sizeDelta = new Vector2(canvasRect.width, canvasRect.height);
```
> 특이한 현상으로, `canvasRect.position`이 `960, 540`이 아니라 `-960, -540`으로 잡히고 있다. 왜 그런지 모르겠음. 일단 임시방편으로 저렇게 달아놓음.

- [x] ItemUIElement 관련 수정 3 - `BackArea`가 나타날 때, `ItemUIElement`의 일부 요소도 패널에 뒤덮이는 현상이 있음
	- `ItemUIElement`에 `Image`가 들어가 있는 경우, 자식 오브젝트인 `BackArea`가 더 위에 렌더링되기 때문으로 보인다. Image 부분을 분리해서 자식 오브젝트로 옮겨놓음.
	- `BackArea`의 경우는 아예 `ItemUIElement`의 바로 자식으로 오게끔 수정한다.
```
ItemUIElement
- BackArea
- ItemDetailPanel
	- 구성 요소..
```
> 기존엔 `BackArea`가 `ItemDetailPanel`의 자식에 있었는데, 이 경우 `BackArea`가 디테일 부분도 어둡게 덮는 현상이 있었음. 

- [x] Enemy 충돌 크기 조정
```
ㅁ ㅁ O
X ㅁ
```
- 지금 O에 있는 오퍼레이터가 공격범위 1칸인데도 X에 있는 적을 공격할 수 있음 : 콜라이더가 너무 크기 때문인 듯?
- 캡슐 모델의 스케일이 `0.25, 0.25, 0.25`임. 이거에 맞추면 되는데, 대신 `y`값의 경우는 타일과의 콜라이더 계산도 들어가고 있기 때문에, 콜라이더의 사이즈를 `0.25, 1, 0.25` 정도로 맞추면 될 것 같다.
- 오퍼레이터의 콜라이더도 생각을 해봐야 하겠는데, `0.8, 1, 0.8`이다. 콜라이더로 저지나 공격 로직을 바꿨던 이유가 **두 타일에 걸친 상태**를 구현하기 위함이었는데, 이렇게 해도 크게 상관은 없지 않을까 싶음.
- 테스트해보니 이 정도면 괜찮은 것 같다. 
### 250305
#### 수정 사항
- [x]  `DeployableActionUI`의 SP 상태에 따른 스킬 버튼, SP 수치 표시 등 컨트롤하기
	- 색상 관리도 `ResourceManager`로 빼뒀다. 
		- 예를 들면 SP 게이지에 사용됐던 스킬 사용 중일 때의 색과, 스킬을 사용하지 않는 상태의 색은 `SkillButton` 위에서도 똑같은 색을 사용해야 함 - 같은 상황에서의 게이지의 색을 표시하는 것이기 때문
##### 겹치는 요소 중복 클릭 문제
-  이전과 비슷한 문제 : `DeployableActionUI`의 버튼들과 `Operator`가 겹쳐졌을 때, 버튼 클릭 -> Operator 클릭 동작이 동시에 이뤄지는 문제
	- 저번에 수정한 내용으로, `DeployableActionUI`의 다이아몬드(클릭해도 동작하지 않음)가 `Operator`의 클릭을 빼앗아가는 현상이 있어서 넣은 코드가 있었는데, 얘가 문제를 일으키고 있는 것으로 보임
	- `Button`의 동작은 `EventSystem, GraphicRaycaster`로 동작되는데, 이게 `ClickDetctionSystem`에서 **레이캐스트를 쏘는 로직과는 별도로 동작함.** 따라서, 1번의 클릭만으로 2개의 병렬적인 레이캐스트가 수행되게 되는 것이다.
	- **이상한 현상이 있다.**
		- 일단 레이캐스트 자체는 `RaycastAll`을 쓰고 있기 때문에, 해당 포인터 위치에 있는 **모든** 요소를 가져오는 방식임
		1. 게임 진행 중에 스킬 버튼의 위치를 클릭하면, 레이캐스트에서 감지 못함
		2. **게임이 끝난 상황**에서 `DeployableActionUI`가 떠 있을 때 스킬 버튼의 위치를 클릭하면 **레이캐스트에서 감지**함(???)

- 위 이슈 테스트
1. 일단 스테이지 진행 중일 때의 `DeployableActionUI`의  `Canvas, GraphicRaycaster`을 보면
	- 스테이지 진행 중
		- `Layer : UI_WorldSpace`
		- `Canvas`
			- `World Space`
			- `EventCamera : Main`
			- `Sorting Layer : Default`
			- `Order in Layer : 1`
			- `Additional Shader Channels : TexCoord1, Normal, Tangent`
			- `Vertex Color Always ...` : 체크 해제
		- `Graphic Raycaster`
			- `Ignore Reversed Graphics` : `true`
			- `Blocking Objects` : `None`
			- `Blocking Mask` :  `everything`
	- 스테이지 진행 후 비교
		- 별 차이 없음
2. 버튼 위에 다른 `Screen Overlay` 패널이 씌워진 상태에서 버튼의 위치 클릭
	- 버튼 위에 다른 패널이 없을 때 클릭 : 레이캐스트에 나타나지 않음
	- 버튼 위에 다른 패널이 있을 때 클릭 : 레이캐스트에 나타남
	- 저녁 먹고 왔다. 머리를 식히고 오니 약간 이런 현상이 의심스럽다
		- 패널이 있을 때 클릭한다 = 해당 버튼의 동작이 이뤄지지 않는다
		- 패널이 없는 상태에서 버튼 위치를 클릭했을 때, 버튼이 동작하든 하지 않든 `DeployableActionUI`가 사라진다. 
		- 그러면 **`DeployableActionUI`가 사라지기 전에 버튼의 동작이 실행되고, `DeployableActionUI`가 사라진 다음에 레이캐스트 점검 로직이 실행되는 게 아닐까?**
	- 실제로 `DeployableActionUI`의 버튼 동작과, `ClickDetectionSystem`의 동작을 로깅해보면 아래처럼 나타났다.
```cs
1. Button의 클릭 동작 수행
2. Button이 클릭됐을 때, Hide() 메서드에 의해 DeployableActionUI가 파괴됨
3. ClickDetctionSystem의 동작이 수행됨
4. 결과적으로 Button이 사라진 상태에서 ClickDetectionSystem의 동작이 이뤄지게 되고, Button의 레이캐스트가 이뤄지지 않음
```
> 위의 추측이 맞는 것 같음.
> 그렇다면 버튼에 이뤄지는 레이캐스트가 먼저 이뤄지고, 그 다음에 `ClickDetctionSystem`의 클릭 처리 동작이 작동하는 것으로 보인다. 이걸 어떻게 관리하는지가 중요하겠음.

- 이런 스크립트를 `ClickDetectionSystem`에 만들었는데, 가장 위의 로그도 나타나지 않음. 왜인지는 몰라도 제대로 동작하지 않는 듯.
```cs
    public void OnPointerClick(PointerEventData eventData)
    {
        Debug.Log("OnPointerClick 동작");
        // 클릭된 위치의 가장 위에 있는 UI 요소를 반환함
        GameObject clickedObject = EventSystem.current.currentSelectedGameObject;

        Debug.Log($"{clickedObject.name}");

        if (clickedObject != null)
        {
            Button button = clickedObject.GetComponent<Button>();
            if (button != null)
            {
                DeployableActionUI actionUi = clickedObject.GetComponentInParent<DeployableActionUI>();
                if (actionUi != null)
                {
                    uiClickHandled = true;
                    return;
                }
            }
        }
    }
```

- 그래서 좀 지저분한 해결 방식이기는 한데, `DeployableActionUI`의 버튼 클릭 동작 실행 -> `ClickDetctionSystem`의 상태 변경 -> 이를 감지해서 `HandleClick`이 동작하지 않게 하겠음
```cs
// ClickDetectionSystem.cs
    // 이미 실행된 UI가 있는 경우, 이 스크립트가 동작하지 않아도 되게 함
    public bool buttonClickedThisFrame = false;
    private bool shouldSkipHandleClick = false;

	// ...

    private void Update()
    {
        if (Input.GetMouseButtonDown(0))
        {
            HandleMouseDown();
            shouldSkipHandleClick = false; // 매 프레임 초기화
        }
        if (Input.GetMouseButtonUp(0))
        {
            // UI 클릭이 없었을 때에만 HandleClick 동작
            if (!shouldSkipHandleClick)
            {
                HandleClick();
            }

            // 다음 프레임을 위한 초기화
            buttonClickedThisFrame = false;
            shouldSkipHandleClick = false;
        }
    }

// DeployableActionUI.cs
    private void OnSkillButtonClicked()
    {
        if (deployable is Operator op)
        {
            if (op.CurrentSkill.autoActivate == false)
            {
                op.UseSkill();
                UpdateSkillButton();
            }
            ClickDetectionSystem.Instance.OnButtonClicked();
            Hide();
        }
    }

    private void OnRetreatButtonClicked()
    {
        deployable.Retreat();
        ClickDetectionSystem.Instance.OnButtonClicked();
        Hide();
    }
```
> 스킬 버튼을 클릭했음에도 카메라가 원래 위치로 돌아가지 않음 : 이건 `Hide` 메서드만 수정하면 될 듯

일단 `deployableActionUI`의 어떤 버튼을 클릭했을 때, 그 버튼 밑에 가려진 다른 요소들이 함께 클릭되는 현상은 해결했음. 다른 버튼들에도 구현...하는 게 좋겠지만 크게 필요는 없을 듯. 버튼이 사라지는 경우가 그렇게 많지 않아서..





#### 오늘 해결 안된 문제
- [ ] 스테이지 클리어 실패했는데도 `StageResultPanel`에 보상이 나타나는 문제
- [ ] 뱅가드 스킬 - 코스트 회복 파티클 : 우측 하단에 너무 가까울 경우, 생존 시간이 남아있기 때문인지 파괴됨에도 다시 생성되는 문제
- [ ] ItemUIElement 관련 수정
1. `ItemDetailArea`(설명문 부분) - 너비는 일정하게, 높이는 `Detail` 스크립트에 비례하게 수정하기
	- 아 근데 이거 `ScrollRect`라서 그냥 높이를 수정하는 방법이 훨씬 나을 듯.
2. `ItemDetailPanel`(클릭시 나타나는 아이템의 자세한 내용 패널)
	- 아이템이 나타나는 위치에 따라, 이 디테일 패널이 잘릴 우려가 있음
	- 이 패널의 가장 오른쪽이 스크린의 가장 오른쪽을 초과하면 왼쪽에 나타나도록 구현하기

> 지식이 늘었다
> - `rectTransform.position`과 `rectTransform.anchoredPosition`
> - 전자는 월드 기준, 후자는 로컬 기준임
> - 보통 로컬로 쓸 상황이 많은 듯?

### 250304

#### 자잘?한 수정 사항
- [x] `DeployableBox`에 쿨타임을 기다리는 원형 게이지 추가
![[Pasted image 20250304162840.png]]

- [ ]  `DeployableActionUI`의 SP 상태에 따른 스킬 버튼, SP 수치 표시 등 컨트롤하기
	- 하다가 마무리 못함 -> 내일 계속

#### 기타 이슈 수정
- [x] 1회 퇴각 후에 오퍼레이터를 다시 배치했을 때, `DeployableBox`가 사라지지 않는 현상
	- 오퍼레이터가 배치 가능한지 여부는 `DeployableUnitState`에서 정보를 갖고 있도록 했다. 이는 `DeployableInfo`를 키로 접근이 가능한데, `IsDeployed`로만 박스의 표시 여부를 결정하게 수정했음
	- `DeployableManager`에서 `DeployableBox`의 표시 여부`gameObject.SetActive()`를 결정시키는 대신, `DeployableBox.UpdateDisplay(DeployableUnitState unitState)`로 관련된 로직을 싹 다 옮겼다. 

### 250303

#### 밸런싱 시작
- `1-0`은 이 정도면 된 듯 함

#### 튜토리얼 구현 계획
- 1-0 진입 ~ 레벨업 및 정예화까지
	- 레벨업, 정예화까지 하면 몰려오는 상황에 대응할 수 있는 오퍼레이터를 강제로 1정예화를 만들어줄 수도 있을 듯
	- 가장 유능한 캐릭이 뱅가드 같아서 뱅가드한테 주면 좋지 않을까 싶다.
- 대충 계획
	- 코스트 회복
	- 오퍼레이터 배치법
	- 스테이지 클리어 후 아이템 사용법
#### 스테이지 클리어 후 경험치 아이템 지급하기
- `StageData`에 `ItemData`들을 넣고, 스테이지 클리어 시 유저에게 아이템이 지급됨
- 관련 로직을 구현하고, UI도 `StageResultPanel`에 구현하기

> 여담) 귀찮아서 AI한테 코드 짜달라고 해야지 -> 이제 `Claude`를 쓰고 있지 않아서 관련 정보들을 다 찾아서 전달해줘야 함 -> 관련 정보를 찾다보니 이건 이렇게 하고 저건 저렇게 하면 되겠네?가 계획됨 -> 질문을 안 하고 직접 코드 작성
> - 초반에 유니티에 대해서 잘 모를 때는 그래도 질문을 하는 편이었는데, 최근엔 그래도 익숙해졌음을 느낀다.  
> - 문제는 UI 구현할 때마다 뇌가 리셋됨 ㅋㅋㅋㅋ 그래도 만들어 놓은 게 있어서 그걸 쓰면 되기는 함.

- 이전에 가져왔던 `UniversalBlurUI`의 경우, 캔버스의 다른 요소들을 블러처리해주지는 않는다.(기억력 이슈) 따라서 뒷배경에 클릭해서 패널이 나타난 상태를 취소시키는 버튼을 추가하고 싶다면 그냥 어둡고 알파값을 1 미만으로 놓는 게 차선책으로 보임

- UI 구현하는 게 생각보다 오래 걸렸다. **`StageResultPanel`의 나가는 버튼 기능을 별도의 내부 버튼을 구현해서 그 부분을 빼놓고, 레이아웃 정리했고, `ItemUIElement`의 로비 팝업이랑 스테이지에서 나오는 패널 구분하고, 스테이지의 패널 구현했음.**

#### 기타 이슈 수정
- [x] `EnemySpawner`에서 `EnemyPathIndicator`와 `Enemy`를 구분해서 생성하지만, 어떤 필드를 사용하느냐에 관계 없음. 즉, `EnemyPathIndicator`를 지정했지만 프리팹에는 `enemy`를 할당하는 것도 가능한 상황
	- 에디터를 고치는 것도 가능하지만, 이게 이슈가 될 수 있는 상황은 개발 상황 뿐이기 때문에, `PathIndicator`를 지정했는데 `Enemy`가 나오거나 혹은 그 역인 상황일 때 오류문구를 출력하는 정도로 충분할 듯
	- `EnemySpawner`에서 로깅 처리만 추가함

- [x] 저지한 적을 제거함 -> `Operator`의 콜라이더 내에 충돌 중인 적이 있음에도 해당 적을 저지하지 않는 문제 발생 중
	- `OnTriggerEnter`, `OnTriggerExit` 외에도, 콜라이더가 충돌 중일 때마다 호출되는 `OnTriggerStay(Collider other)`가 있다. 이를 이용함.
```cs
    private void OnTriggerEnter(Collider other)
    {
        ProcessEnemyCollision(other);
    }

    private void OnTriggerStay(Collider other)
    {
        ProcessEnemyCollision(other);
    }

    private void ProcessEnemyCollision(Collider other)
    {
        if (IsDeployed && blockedEnemies.Count < currentStats.MaxBlockableEnemies)
        { 
            Enemy collidedEnemy = other.GetComponent<Enemy>();

            if (collidedEnemy != null &&
                CanBlockEnemy(collidedEnemy.BlockCount) && // 이 오퍼레이터가 이 적을 저지할 수 있을 때 
                collidedEnemy.BlockingOperator == null) // 해당 적을 저지 중인 아군 오퍼레이터가 없을 때 
            {
                BlockEnemy(collidedEnemy); // 적을 저지
                Debug.Log($"{this.BaseData.entityName}이 {collidedEnemy}을 저지하기 시작함, 현재 저지 수 : {blockedEnemies.Count}");
                collidedEnemy.SetBlockingOperator(this);
            }
        }
    }
```
> 로직 자체는 동일하기 때문에 `ProcessEnemyCollision`이라는 별도의 메서드로 빼두고, `OnTriggerEnter`와 `OnTriggerStay`를 수정한다. 

- [x] `Sniper` 에서 맵 범위를 벗어난 곳을 공격 범위로 지정할 경우, `NullReferenceException`이 발생하며 공격을 시행하지 않는 문제
	- 그냥 `Nullable` 추가하니까 잘 됨. 엌ㅋㅋ
```cs
// ? 추가
	Tile? eachTile = MapManager.Instance.GetTile(eachPos.x, eachPos.y);
```
# 이전 일지

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