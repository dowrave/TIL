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

- 테스트 및 수정

### 작업 중 
1. 스테이지 `1-1` 밸런싱 (이후 `1-2`, `1-3`)
2.  `OperatorDetailPanel`에서 스킬 선택이 되지 않는 현상 수정 필요

### 구현 예정

### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
	- 예를 들면 1-1에서 바리케이드만 이용해서 스테이지를 클리어하기 같은 게 있겠다 (= 오퍼레이터를 배치하지 않고 스테이지 클리어하기)


## 3월

## 250328 - 짭명방

### 스테이지 밸런싱

### 정예화 패널 구현
- 정예화 패널을 구현을,,, 했던가,,? 
- 정예화 자체는 기능하지만, 정예화 결과인 사정거리 증가 표현, 스킬 설명은 구현되어 있지 않아서 이 부분을 추가함
	- 초반엔 유동적인 레이아웃을 구성하려고 했는데, 스킬의 설명 분량이 많지는 않아서 고정된 길이로 사용함. UI가 들어가는 영역이 한정적이기도 하다.
- 이것저것 만져보는 중이라 시간이 좀 걸림.


### 이슈 수정
- [x] `OperatorInventoryView`에 스테이지 클리어 후에 다시 들어갈 때 `NullReferenceException` 오류 발생
	- 과정)
		1. `OwnedOperator`의 공격 범위가 제대로 초기화되지 않는가? (X)
		2. `attackRangeHelper`의 초기화가 제대로 되지 않는가? (O)
			- `Start`에 있었고, 스테이지 씬에 갔다가 돌아왔을 때 Start에서 디버깅을 해봣는데 아예 동작하지 않은 듯?
			- `Claude 3.7` 한테 물어보니까 - **최초에 비활성화된 요소가 활성화될 때의 동작 순서는 `Awake -> OnEnable -> Start` 순서**임. 
				- `Start`는 1번째 `Update` 호출 직전에 호출된다. 
				- `Awake`는
					1) 오브젝트가 처음에 있었고 비활성화일 때 : Awake가 호출되지 않음
					2) 인스턴스화되지만 바로 비활성화될 때 : Awake가 호출되고 비활성화됨
			- 이는 씬이 전환되었다가 다시 돌아왔을 때에도 동일하다. 씬 전환 시, 이전 씬의 모든 게임 오브젝트는 파괴되기 때문이다. 다시 돌아왔을 때에는 새로운 게임 오브젝트를 만들어서 진행하므로, 위와 상황이 달라지지 않는다.
			- `DontDestroyOnLoad` 혹은 `LoadSceneMode.Additive`로 씬을 로드한 경우는 상황이 달라짐
	- 수정 방향)
		- 위에서 말한 내용을 `Awake`, 혹은 `OnEnable`로 옮기면 된다.
		- `Awake`로 넣음.




- [x] `NotificationPanel`이 나타나야 할 상황일 때, 계속 클릭 시 패널이 계속 나타나는 현상을 수정하겠음
	- 원래는 `Dictionary`를 만들어서 중복되는 메시지가 활성화된 경우에는 나타나게 하고 그렇지 않은 경우는 없애는 방식으로 구현하려고 했지만.. 그냥 쉽게 가는 게 좋겠다. 
	- 수정) `NotificationPanel`이 나타날 때 쿨타임을 2초(총 2.6초 정도임)로 설정

```cs
    public DateTime LastNotificationTime { get; private set; }


    public void ShowNotification(string message)
    {
        // 저번 알림 패널이 2초 내에 떴다면 아무 것도 활성화하지 않음
        if (DateTime.Now - LastNotificationTime < TimeSpan.FromSeconds(2)) return;

        if (notificationPanelPrefab != null && mainCanvas != null)
        {
            GameObject notificationObj = Instantiate(notificationPanelPrefab, mainCanvas.transform);
            NotificationPanel notificationPanel = notificationObj.GetComponent<NotificationPanel>();
            notificationPanel?.Initialize(message);

            LastNotificationTime = DateTime.Now;
        }
    }
```
> 유니티에서 시간을 시간 타입으로 다룬 건 놀랍게도 처음이다.

## 250327 - 짭명방

### StageResultPanel - 좌우 스크롤 구현하기

- [x] 완료 여부

> 목표
> - 왼쪽 스크롤 끝은 1번째 아이템, 오른쪽 스크롤 끝은 마지막 아이템이 오른쪽 끝에 오는 시점에서 끝내기
> - 인스펙터 단위에서 `Scroll Rect`을 설정하는 것만으로는 안되는 것 같음. 스크롤을 어디에서 놔도 처음 위치로 돌아와버리는 현상이 있다

- ~~일단 뭘 만들기는 했는데 제대로 동작하지는 않고 있음. ItemUIElement의 클릭 동작이 되지 않는 현상도 있다.~~ 
	- 이걸 포함한 다른 문제를 먼저 해결하고 돌아왔다. 

- [[Unity - Scroll Rect]] : 유니티의 Scroll Rect을 간단하게 정리
	- 지금 발생하는 문제점은 `Elastic`이어도 맨 처음으로 스크롤이 돌아오고, `Clamped`일 때는 아예 스크롤이 움직이지 않는다. 
	- 위의 내용에 따르면, `Elastic`과 `Clamped`의 차이는 아래와 같다.
		- `Elastic` : 스크롤이 가장자리에 위치할 때 가장자리 밖으로 스크롤이 튕겨나갔다가 들어옴
		- `Clamped` : 가장자리에서 스크롤이 멈춤
	- 그러면 `Scroll Rect`가 보는 `Content`의 너비가 뷰포트보다 크지 않다?라는 게 되나?

- 해결) **`Content`에 `Content Size Fitter` 컴포넌트를 추가**
	- 예전에도 종종 다뤘던 컴포넌트인데, 이번에 정확히 어떤 역할을 하는지 정리해야겠다.
	- [[Unity - Content Size Fitter]]
	- 요점은 **`Content Size Fitter`에서 원하는 방향으로의 `Preferred Size`를 설정하라는 것.** 그리고 그 기능은 **자식 오브젝트들의 변화에 따라 부모 오브젝트의 `RectTransform` 값이 변경된다**는 것이다. 
		- 스크롤 영역의 계산은 `Content의 너비/높이 - Viewport의 너비/높이`인데, 기존에는 `Content의 너비/높이`가 최초 설정 값으로 고정되기 때문에 자식 요소가 얼마나 추가되든 스크롤의 영역에는 큰 변화가 없었다. 그리고 `Content` 의 앵커 설정이 `Stretch`였기 때문에 스크롤 영역의 계산값이 `0`으로 나타났을 것이다.
		- `Content Size Fitter`를 추가하면, 앵커 설정에 관계 없이 `Content Size Fitter`에서 `RectTransform`을 동적으로 제어하게 되고, `Content의 너비` 값에 변화가 생기므로 실시간으로 스크롤 영역의 변화가 발생하게 된다.

- 추가) `ReturnToLobbyButton`이 설정된 영역보다 크게 잡혀서, 스크롤 영역을 침범하는 현상이 있음
	- 해결) **`ReturnToLobbyButton`을 `RewardItemContainer`보다 위에 놔서 더 나중에 렌더링**되게 함
		- 이전에 그렇게 구현해놓은 이유가 있었다..

### Spawner 용 ScriptableData 만들기

- [x] 완료 여부

> 불편했던 점
> - 기존에는 `Spawner` 각각을 열어서 편집해야 했음. 테스트해야 하는 상황에서 기존 스폰 정보들을 별도로 저장하는 방법이 없다.
> - 관련 정보들을 저장해야 하므로 `ScriptableData`로 스폰 정보들을 따로 저장하겠음.

- `EnemySpawnInfo`를 아래처럼 변경.
```cs
public class EnemySpawnerConfig : ScriptableObject
{
    public List<EnemySpawnData> spawnedEnemies = new List<EnemySpawnData>();
}

// 얘가 원래 EnemySpawnInfo
[System.Serializable]
public class EnemySpawnData
{
    public SpawnType spawnType;
    public float spawnTime = 0f;
    public PathData pathData = default!; // StageData에서는 경로 데이터를 직접 참조하도록 수정
    public GameObject prefab = default!; // 스폰되는 종류가 다양할 수 있기 때문에 EnemyData를 사용하지 않음

    [Tooltip("spawnType = Enemy일 때만 사용")]
    public EnemyType enemyType;
}
```

- 스테이지 `1-0`도 이거에 맞춰서 다시 작성하고 오늘 작업 마무리

### StatisticItem 색 옆에 클래스 아이콘 추가
- [x] 완료 여부

- `StatItem`은 크게 2가지 용도로 사용하고 있다.
	- `StatsPanel`에 프리팹이 아닌 고유 오브젝트로 넣음
		- 3개 이상 안 쓸 거라서 그냥 미리 패널에 할당해놓고 값이 들어오면 켜지는 방식
	- `StageResultPanel`에서 프리팹으로 인스턴스화 하는 요소

- 근데 `StatItem` 자체를 살펴보면 가장 부모 오브젝트의 크기가 `0, 0`으로 되어 있음
- 프리팹으로 이 요소를 사용하는 상황은 `StageResultPanel`이며, 셀 크기는 `Grid Layout`에 의해 결정되기 때문에 이렇게 구현했던 것으로 보인다
- 이걸 작성할 때는 오퍼레이터의 수가 8명을 초과하는 상황까지 가정해서 작성했는데, **최대 6명만 사용할 것이기 때문에** `Width, Height`를 `400, 100`으로 프리팹에 저장해놓겠음

- 구현) 
	- `ClassIconBackground`을 `OperatorIcon` 오른쪽에 하나 추가, 크기는 `45, 45`
	- `ClassIcon`을 `ClassIconBackground`의 자식으로 놓음. 중심에 놓고 크기는 `40, 40`

### 기타 이슈

- [x] 위의 과정들 수정하면서 `StageResultPanel`의 통계 부분 클릭이 동작하지 않는 현상 발견
	- 뭐가 원인인지 모르겠다.
	- 수정) 이전에 이거저거 만지는 과정에서 `OnDestroy`에 있던 로직들을 `OnDisable`에 옮겼는데, `Awake`에서 등록된 리스너들이 `StageScene`에 진입하면서 `StageResultPanel`을 비활성화하게 되고, 이 과정에서 리스너가 해제되는 문제가 발생한 듯. 
		- **해당 로직들을 다시 `OnDestroy`로 옮겼다.** 옮기지 말라는 주석도 달아놓음.


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
	- [[25년 3월]]