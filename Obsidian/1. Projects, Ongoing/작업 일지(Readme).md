# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 보고 참고할 내용
- [[오퍼레이터들 스탯 정리]]
- [[적 스탯 정리]]
- [[프로크리에이트로 텍스쳐 작업 시 유의할 점]]
## 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`

## 하고 싶은데 못한 것
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.
	- 현재는 정예화 패널에서만 스킬 범위를 볼 수 있음.



---
# 작업 일지

## 짭명방 예정

### 작업 중

>[!wip]
>- 머티리얼 변경하면서 제대로 나타나지 않는 이펙트들 수정 중
>- 쉴드 스킬 파티클 시스템으로 변경
>- 스테이지 밸런싱

>[!todo]
>- 보스 구현
>	- 쉴드 스킬
>		- 얘만 지금 VFX Graph로 남아 있음.
>- 남은 작업
>	- 스테이지 1-3 완성
>	- 남은 스테이지들 밸런스 수정
### 간헐적인 이슈
- 이슈가 있다고 느꼈는데 다시 테스트했을 때 재현이 안된 것들을 정리함

- 오퍼레이터 A를 배치할 때, 방향 설정 로직 중 오퍼레이터 B의 위치에서 마우스 커서를 떼면 배치되면서 해당 마우스 커서의 위치에 있는 오퍼레이터가 클릭되는 현상

# 251001 - 짭명방

>[!done]
>1. `OperatorUI` 사라지는 문제 수정 완료
>2. UI Sprite Atlas 생성 - 생각보다 엄청 빠르고 쉽게 했음
>3. 기존에 구현했던 VFX Graph나 텍스쳐 등등 정리

## OperatorUI 사라지는 문제 수정
- 다른 오퍼레이 사망 시 발생했던 이슈 계속) `InStageInfoPanel`은 유지되는데 다른 정보도 함께 사라지는 현상

```cs
    public void OnDeployableRemoved(DeployableUnitEntity deployable)
    {
        deployedItems.Remove(deployable);

        // 어제 수정한 부분
        StageUIManager.Instance.HideInfoPanelIfDisplaying(deployable);
        
        // 수정해야 하는 부분
        HideOperatorUIs();
        ResetHighlights();
        
        
    // 오퍼레이터 주위에 나타난 UI 제거
    private void HideOperatorUIs()
    {
        // null이어도 상관 없음

        if (currentActionUI != null)
        {
            Destroy(currentActionUI.gameObject);
            currentActionUI = null;
        }

        if (currentDeployingUI != null)
        {
            Destroy(currentDeployingUI.gameObject);
            currentDeployingUI = null;
        }

        currentUIState = UIState.None;
    }
```

- 이 부분은 `Instantiate`와 `Destroy`로 구현되어 있기 때문에 **오브젝트 풀링으로 바꾸는 것까지 해야 할 듯**

근데 하나씩 접근하자. 동시에 생각하면 머리가 복잡해져서 하나도 제대로 못할 때가 많다.

#### 다른 오퍼레이터가 죽으면 OperatorUI 유지
- `currentActionUI`, `currentDeployingUI`에 대해서도 어제 `InStageInfoPanel`에서 처리한 것처럼 "현재 어떤 오퍼레이터/배치 가능 요소"에 대한 정보를 띄우고 있는지에 대한 필드를 갖게 한다. `Deployable` 게터 프로퍼티를 노출시키겠다는 뜻.
- `DeployableManager`에서는 아래처럼 구현
```cs
// HideOperatorUIs 을 복붙, deployable
    private void HideOperatorUIsOnCondition(DeployableUnitEntity deployable)
    {
        if  (currentActionUI != null && currentActionUI.Deployable == deployable)
        {
            Destroy(currentActionUI.gameObject);
            currentActionUI = null;
        }

        if (currentDeployingUI != null && currentDeployingUI.Deployable == deployable)
        {
            Destroy(currentDeployingUI.gameObject);
            currentDeployingUI = null;
        }
    

        currentUIState = UIState.None;
    }
```

>  조건문의 순서를 생각없이 반대로 넣었다가 오류가 떴다. **당연히 `null` 체크 후에 해당 객체의 필드에 접근해야 함.** 
#### 요소들 생성/파괴 수정
- **게임 내내 한꺼번에 1개씩만 띄우는 요소**이므로 오브젝트 풀링으로 구현할 것도 없고, 단순히 활성화 / 비활성화하게끔 하면 됨
- 기존에 프리팹을 할당했던 요소들을 캔버스 내에 넣고 오브젝트로 대체하고 필드로 할당함
- 해당 클래스들의 활성화 / 비활성화나 전환에 따른 상태들 관리만 잘해주면 됨. 이건 기존에 구현되어 있었을 거임.

```cs
currentActionUI.gameObject.SetActive(true);
currentActionUI.Initialize(deployable);
```
> 활성화 후 초기화가 더 권장되는 방식이라고 함
> - 유니티 생명주기라든가 UI 레이아웃 등은 활성화 시점에 맞춰서 동작하기 때문
> - 모든 게 준비된 다음 활성화하는 게 깔끔하다고 생각했는데 **활성화 -> 초기화**가 유니티에서 더 어울리는 방식이라고 함.

- `HideOperatorUIs`들도 `null`을 할당하는 부분을 제거함

- 스크립트 상에서 오류는 없는데 나타나는 위치가 이상해짐 -> 수정
	- `ActionUI`는 `WorldCanvas`인데 `Screen Canvas` 안에 집어넣어서 발생하는 문제였음
- 오퍼레이터를 클릭했는데 ActionUI가 안 나타나는 현상 수정 중
	- 수정 완료. 얘도 `Screen Canvas` 안에 집어넣어서 발생한 문제였다.

## UI 스프라이트들 통합
- **이거는 생각보다 엄청 빠르게 진행됐다. 유니티의 `Sprite Atlas` 기능 때문.**

- `Sprite Atlas`라는 게 있다고 한다. 
- 기본적으로 꺼져 있는 듯. `Project Settings - Editor - Sprite Atlas`의 `V2, Enabled`을 사용함.
- 근데 이걸 적용해봤더니 기존 이미지들의 영역에 다른 이미지가 함께 나타나는 현상이 있음. 
	- **`Tight Packing` 옵션이 이슈**였음. 이걸 켜면 정확한 영역으로 저장하고 끄면 이미지 사각형 단위로 저장하는 듯.

- `Sprite Atlas` 자체도 미리보기에서 1개만 보일 수 있는데, 우측 상단을 보면 `Page`를 전환할 수 있다. 이걸로 지정한 경로의 이미지들이 모두 들어갔는지를 볼 수 있음

- 이 `Sprite Atlas`에 포함된 `Image`는 유니티에서 원본 텍스쳐를 `Image` 컴포넌트에 할당했더라도, 유니티에서 자체적으로 `Sprite Atlas`에서 해당 이미지를 가져온다. 
	- 원본 파일의 이름이 그대로 나타나고 있어서 헷갈릴 수 있는 부분인데 유니티에서 자체적으로 처리하고 있는 상태라고 함. 
	- 만약 정말로 그런지 알고 싶다면 `Window > Analysis > Frame Debugger`로 해당 이미지를 어디서 가져오고 있는지 체크할 수 있다. 

![[Pasted image 20251001175626.png]]
`SpriteAtlas`가 표기되어 있음
### 그러면 드는 생각
- **이펙트 텍스쳐 시트도 Sprite Atlas를 이용해서 만들 수 있나?** 오늘 해본 것처럼, UI 아이콘 Atlas를 만드는 건 겁나 빠르고 편했다. 

- 제미나이에게 물어보니 **권장하지 않는다고 함**

1. **유니티의 `Sprite Atlas`는 그리드 형태가 아님.**  
	- 파티클 시스템은 그리드라고 가정하고 진행하며, `Sprite` 모드가 있긴 한데 모든 프레임을 리스트에 수동으로 복붙해야 하는 번거로움이 있다.
2. 밉맵 생성 문제
	- 각 프레임 주위에 충분한 여백이 필요하다. 밉맵이 생길 때 이웃 프레임의 색이 섞이는 현상을 방지하기 위함이다.
	- `Sprite Atlas`는 패킹 효율을 극대화하기 위해 스프라이트들을 매우 가깝게 붙인다. 


---
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
	- [[짭명방_25년 9월]]
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