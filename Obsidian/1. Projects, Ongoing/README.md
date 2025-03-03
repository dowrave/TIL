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