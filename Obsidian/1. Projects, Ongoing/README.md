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




---
# 작업 일지

## 이슈? 정리

# 짭명방 예정



# 250715 - 짭명방

## 작업 중 

>[!todo]
>- DualBlade의 2가지 스킬을 구현
>- 1번째 스킬은 똑같이 강타인데, 강타가 나갈 때는 마법 대미지가 나가는 방식
>- 2번째 스킬은 아이디어 자체를 생각해봐야 함. 
>	- 일단은 별도의 이펙트를 구현하지 않고 만들 방법을 고민 중
>		- 그래서 버프 형식이 될 것 같다. 
>		- 공격력과 공격 속도가 증가하고 스킬이 켜진 동안 공격 타입이 마법이 되는 방식이면 될 듯. 

- 1번째 스킬 구현
	- 기존의 `SmashSkill`과 동일하되, 공격 타입만 바뀌는 방식이어야 함
	- 이 과정에시 **기존의 공격 타입이 `ICombatEntity.AttackType`에 의존했는데, 대신 `AttackSource` 자체에 `AttackType`이 담기도록 변경함.**
		- 그래야 기본적으로 `Physical` 공격을 하는 캐릭터라도 어떤 효과로 인해서 `Magical` 공격을 하도록 구현할 수 있기 때문에
	- 이거 시스템을 건드린 거라서 생각보다 오래 걸릴 수도 있다.
		- 어차피 언젠가 고칠 일이긴 했기 때문에 시간을 들여서 작업 ㄱㄱ
		- 일단 시간이 없어서 여기까지 하는데 세부적으로 남은 것들 정리해봄

>[!todo]
>하다가 못 끝냄) **`Projectile`의 `AttackSource` 작업** 
>- 투사체를 생성하는 시점에 `AttackSource`를 담는 것도 생각해봤는데 그러면 판정이 발생하는 시점의 위치 정보를 못 담게 됨. 아니면 생성할 때 담되 
>	- 나중에 **대미지가 들어가는 판정 시점에 위치값만 따로 담는 것도 방법**일 듯. 
>- 나머지 더 할 일이 있는지 찾아보자. 
>- 아마 다음에 켜면 오류가 많은 상태일 거임~ 기억해 둘 건 `AttackSource`로 `attacker` 정보나 `damage` 정보들이 다 들어가 있다는 거!



## 작업 완료

>[!done]
>**1. 아이콘 계단으로 깨져서 보이는 문제**
>2. 평타가 2번 나갈 때, **1번째 공격으로 적이 죽은 상황이라면 2번째 공격의 처리**를 어떻게 해야 하는가?


### 1. 아이콘 계단으로 깨져서 보이는 문제
>- 아이콘의 해상도를 256 x 256으로 올렸는데도 계속 깨져서 보이는 문제가 있음
>- 가우시안 블러 때문일 수 있겠다.... 싶어서 1024 이미지를 수정한 다음 작업해봄
>- **가우시안 블러 이미지를 다시 선명한 이미지로 놓고, 해상도를 높여봐도 발생하는 문제 동일함.** 왜 이 아이콘에 대해서만 이런 문제가 나타나는지는 모르겠지만 크게 개선될 기미는 안 보인다. **256으로 놓고 유지하겠음**


### 2. 2연속 공격의 2번째 공격의 처리

#### 원본 게임 테스트
- 일단 지금의 구현은 1번째 공격으로 적이 죽었다면 2번째 공격은 나가지 않는 방식임.
- 원본 게임을 뜯어보면 아래와 같다. 테스트는 바이비크와 첸으로 해봤다.
1) 2개의 공격 각각에 온히트 처리를 하지 않고, **2번째 공격이 맞든 맞지 않든 SP는 1씩 참.**
	- 이건 좀 의외다. 소드마스터는 당연히 SP가 2씩 차는 줄 알았음.
2) 1번째 공격으로 적이 죽었을 경우, 2번째 공격의 피격음이 달라진다. 즉, **헛스윙에 대한 분기 자체는 있는 것으로 보임.** 

#### 구현 방향 설정
- 이건 사실 정말 어떻게 구현하느냐에 따른 차이 같긴 한데, 난 온히트로 생각하고 작업을 했으니까..
1. **온히트 효과로 구현. 적이 피격했다면 SP가 올라감**
2. 단 **2회의 공격 이펙트는 항상 나타남.** 적이 맞았는지 여부는 피격 이펙트 재생이 별도로 있으니까 따로 구현하지 않아도 괜찮을 것 같다. 
	- 이거를 하려면 공격 로직에서 적이 없어도 이펙트는 재생돼야겠다.

#### 실제 구현
- 이걸 어디서 구현해야 할까?
	 - 처음엔 `Operator.cs`에서 구현해야 한다고 생각했다. 이중공격을 하는 상황 외에도 헛스윙이 발생할 수 있는 상황이 내가 예상치 못한 곳에서 있지 않을까? 라는 생각으로.
	 - 그런데 **기존 코드가 헛스윙으로 인한 문제가 발생한 적이 없었고, 이중 공격을 하는 상황에서는 헛스윙이 발생할 걸 예측할 수 있는 상황**이므로 거기다 우선적으로 구현해놓는 게 맞다고 한다. 만약 나중에 비슷한 문제가 발생한다면 근본적인 메서드를 고쳐야 할 수는 있겠지만.

- 따라서 `DualBladeOperator`에서 2번째 공격이 빗나갔을 때의 처리를 별도로 정의해둔다.

- 대신 기존의 `Operator.PlayMeleeEffect`는 수정이 필요하다. `UnitEntity` 자체를 받는 걸로 설정이 되어 좌표를 수정하는 방식으로 수정함. 여기에 엮여있는 `CombatVFXController`까지 수정한다.




# 250714 - 블로그
- `ReviewCard`에서 별점 + 텍스트, 회차 + 텍스트를 묶게 되면서, 별점/회차의 유무에 따라 해당 줄의 높이가 달라지는 현상이 있다. 텍스트로 묶인 요소가 있을 때 줄이 차지하는 높이가 더 큼
	- 높이값을 일괄적으로 부여하는 방식으로 해결하나?
	- 그게 맞다. `h-`값을 부여하는 방식으로 해결함. `h-6` 정도를 줬다. 

- 그리드 로딩이 2번씩 되는 현상
	- 한번에 24개의 데이터를 불러와야 하는데 48개씩 불러오는 현상이 있다.
	- 수정 완료. 조회 명령을 수행한 다음 딜레이를 줬다. 
		- 이 딜레이를 줄 때 크게 2개를 고려했다.
			1. 연속해서 2번 불러오면 안 됨
			2. 휠을 클릭해서 스크롤바를 계속 아래로 당겼을 때 막히지 않아야 함

2번 상황의 경우, **딜레이가 너무 짧으면 딜레이가 돌기 전에 스크롤바가 최하단에 닿아서 불러오기 기능이 정상적으로 동작하지 않는다.**  테스트해서 구한 시간은 **150ms** 정도.
> 물론 굳이 이런 걸 신경쓰면서 구현할 필요가 없을 수도 있지만, 일일이 스크롤바를 아래로 당기는 방식보다 이 방식을 이용할 수도 있을 것 같아서 고려해봤음.

# 250712 - 블로그

## 리뷰 게시판
- 별점 기준 수정
	- "아쉽지만 재밌다"의 기준을 2.0으로 내림. 
	- 1회차 만점 3.5
	- 그리고 다회차 감상한 작품은 4.0, 4.5로 이분화
	- **보다가 끊은 작품은 어디에 위치해야 할까**를 모르겠다. 
		- **다 본 게 아니니까 별점이 없다---가 바람직할듯?**

> [!note]
> 1. **다회차 표시 기능을 추가**했음.
> 	- **카드에서는 2회차 이상**일 때 나타남(그리드가 너무 번잡해지는 문제가 있음)
> 	- **모달에서는 1회차 이상**일 때 나타남
> 2. 이미 값이 채워진 자리에 빈 값을 보낼 때 변화가 반영되지 않는 현상 수정 
> 	- 이건 그냥 값이 있을 때만 폼에 데이터를 추가해서 그렇다. 
> 	- 수정 모드일 때에 값이 없다면 `''`을 보내도록 해서 수정함.




## 기타 이슈 수정
- 글을 **수정**할 때, 기존엔 글이 있었고 수정 후에는 글이 없는 상태였다면 수정된 글에는 아무런 내용이 없어야 하는데, 기존 글이 남아 있는 현상이 있음. 


# 250710 - 블로그

## 오늘의 교훈
- 기능 구현은 간단해보여도 스타일이라든가 의외의 문제라든가 등등이 발생할 수 있으니 생각보다 시간을 많이 잡아먹는 일이다. 
- 더 수정할 일이 없기를 바랍니다,,,, 이거 1시간이면 끝낼 줄 알았음... 곧 있으면 자정임,,,

## 스포일러인 경우 블러 표시 & 스크롤 전파 디버깅
- 스포일러가 활성화된 경우, 최초에는 본문을 보이지 않게 구현하겠음.
	- 모달을 펼쳤을 때 블러 처리하는 외부 태그 + 클릭 시 내용이 보이게 하도록 구성할 계획.
>[!note]
>1. 클릭 가능한 별도의 패널을 따로 만듦
>2. 본문은 스포일러 상태에 따라 블러가 활성화되거나 비활성화됨. 

추가) 지금 방식은 글의 길이에 영향을 받는다.   
예를 들어서 한 줄짜리 스포일러 글이라면 스포일러 경고 영역의 내용이 잘리고, 너무 길다면 스포일러 경고 문구는 스크롤을 내려야 볼 수 있다.  

---

따라서 스포일러 경고 영역 자체는 고정된 높이로 구현하고, 이를 클릭하면 본문이 들어가는 방식으로 구현해보겠음.

>[!note]
>제미나이에게 물어봤을 때 바로 나온 대답은
>1. 실제 컨텐츠 영역의 높이를 측정한다
>2. 경고가 활성화될 때는 해당 높이를 가진다
>3. 본문을 직접 보여줄 때는 측정한 높이를 사용한다 
>   
>이었다. 근데 이렇게 구현할 거면 그냥 **경고문을 보여줘서 해당 높이를 가진다** -> **본문을 보여줄 때 컨텐츠 영역에 실제 내용을 넣는다**가 훨씬 간단하게 구현하는 방법 같음. 

뭔가 만족스럽지 않다. 또 생각이 많아진다. 방법이 크게 2가지가 떠오른다.
>[!note]
>1. 컨텐츠 영역 자체에만 스크롤을 구현하는 것. 
>2. (현재)모달 전체에 스크롤을 유지하되 위처럼 상태에 맞춰 영역을 바꾸는 것.

1번이 나아보이기는 하는데, 본문 이외의 영역을 제목 부분이 항상 차지한다는 게 걸리는 느낌이다.
그런데 1번처럼 구현하면 **푸터 영역을 제외**할 수 있다. 지금의 구현 자체가 푸터에 이중구현이 되어 있기는 함. 
- 수정/삭제 버튼은 나만 볼 수 있게끔 윗쪽에 나타나 있고, 
- 닫기 버튼도 이미 우측 상단에 있으니까. 

그럼 짧은 글에 대해서는 어떻게 대처할 것이냐? css에서 최소 높이를 설정해주면 되지 않을까?

그런데 예전에 생각했던 것 중에 그게 있었다. 본문이 끝나는 지점이 화면의 아랫쪽에 있으면 뭔가 이상한 느낌이었음. 

>[!warning]
> 현재 문제 상황 : 모달이 활성화된 상태에서 모달 외부에 마우스 커서가 있을 때
> 1. 모달 자체가 스크롤이 가능한 상황 - 외부에 커서를 놓으면 스크롤 가능
> 2. 모달 자체가 스크롤이 불가능한 상황(내용이 적은 이유로) - 외부에 커서를 놓으면 스크롤 불가능
> 3. 모달 내부에 커서가 있는 상태에서는 외부로 이벤트 버블링이 발생하는 건 아님
> **여기서 문제가 되는 상황은 1번.** 

ReviewModal의 스크립트를 보면
```tsx
  usePreventScroll(!!reviewId);
  useModalHistory(!!reviewId, isEditing ? handleAttemptClose : onClose);
```
이 되어 있다. useEffect 문으로 감싼 요소는 아님.

그러면 `usePreventScroll`(body의 스크롤을 막음)이 동작한 다음에 review의 content 들이 들어오면서 content가 있는 영역에 스크롤이 필요해지면 다시 렌더링되면서 `usePreventScroll`의 효과가 풀리는 게 아닐까... 라는 가설을 세워본다.
즉, **내용이 들어온 다음에 usePreventScroll이 동작하면 되지 않을까?** 라는 생각으로 실험해봄.

>[!done]
> 오래 걸렸다. 온갖 쇼를 다 해봤는데, `review-outer-modal`에서 상위 요소로 이벤트가 전파되는 상황이 있으면 막고, `review-inner-modal`에서도 자신부터 자신의 상위요소들을 차례대로 올라가면서 스크롤이 가능한 요소인지 체크, 불가능한 경우에는 스크롤을 다 막는 방식으로 구현했다.
> - 내가 감을 잡고 짚어나가야 했다. AI님께서 계속 안되는 코드를 던져줘서..
> - 그건 그렇고 제미나이 프로가 갑자기 왜 이렇게 아부를 떠는지 모르겠다. 부담스럽다.

```tsx
import { useLayoutEffect } from 'react';

// 모달 배경의 클래스 이름
const OUTER_MODAL_CLASS = 'review-outer-modal';
const INNER_MODAL_CLASS = 'review-inner-modal';

export const usePreventScroll = (isLocked: boolean) => {
  useLayoutEffect(() => {
    if (!isLocked) {
      return;
    }
    // --- 1. 기존의 스타일 제어 (시각적 처리 및 기본 방어) ---
    const originalOverflow = document.body.style.overflow;
    const originalOverscrollBehavior = document.body.style.overscrollBehaviorY;

    document.body.style.overflow = 'hidden';
    document.body.style.overscrollBehaviorY = 'contain'; // 스크롤 체이닝 방지

    // --- 2. 문제가 되는 상황을 위한 정밀 타격 (이벤트 제어) ---
    const handleWheel = (e: WheelEvent) => {
      const target = e.target as HTMLElement;

      const innerModal = target.closest(`.${INNER_MODAL_CLASS}`);

      // 1. 이벤트가 모달 내부에서 발생하지 않은 경우
      if (!innerModal) {
        // 경로에 outer-modal이 있다면 배경 스크롤을 시도하는 것이므로 차단함.
        if (target.closest(`.${OUTER_MODAL_CLASS}`)) {
            e.preventDefault();
        }
        return;
      }

      // 2. 이벤트가 모달 내부에서 발생한 경우
      let current: HTMLElement | null = target;
      let isScrollLegitimate = false;
      
      // 2. target부터 innerModal까지 올라가면서 스크롤 가능한 부모가 있는지 확인
      while (current && current !== innerModal.parentElement) {

        const style = window.getComputedStyle(current);

        const isOverflowing = current.scrollHeight > current.clientHeight;
        const isScrollAllowed = style.overflowY === 'auto' || style.overflowY === 'scroll';

        // 요소의 영역이 보이는 영역보다 크다 + 스크롤을 허용한다
        if (isOverflowing && isScrollAllowed) {
            isScrollLegitimate = true;
            break; // 합법적인 스크롤이므로 검사 중단
        }

        if (current === innerModal) break; // innerModal이 종점 

        current = current.parentElement;
      }

      // 합법적인 스크롤 경로를 찾지 못했다면 외부로 이벤트가 나가려는 시도이므로 차단
      if (!isScrollLegitimate) {
        e.preventDefault();
      }
    };

    // passive: false 옵션으로 preventDefault가 동작할 수 있음을 브라우저에 알림
    window.addEventListener('wheel', handleWheel, { passive: false });

    // --- 3. 클린업 함수: 모든 것을 원래대로 복원 ---
    return () => {
      document.body.style.overflow = originalOverflow;
      document.body.style.overscrollBehaviorY = originalOverscrollBehavior;

      window.removeEventListener('wheel', handleWheel);
    };
  }, [isLocked]);
};
```

## 스크롤바로 인한 서식 변경 방지하기
```tsx
<div className={`flex-grow relative min-h-[150px]
${review.is_spoiler && !isSpoilerRevealed ? 'overflow-y-hidden' : 'overflow-y-auto'}
`}>
```
의 상태 변화로 인해서 오른쪽에 스크롤바가 생기는 경우, 글의 서식에 살짝 변화가 발생한다. 이 변화를 방지하고자 함.

### 해결법1) scrollbar-gutter
- 브라우저에게 "이 영역에 스크롤바가 생길 수 있으니 미리 공간을 확보해라"라고 알려준다. 실제 스크롤바가 나타나도 콘텐츠 너비는 변하지 않는다.

### 해결법2) 스크롤바 스타일링
- 크롬에서 모바일 서식의 스크롤바가 마음에 든다고 던져줬더니 커스텀 스크롤바 CSS를 던져줬다.
- 근데 제대로 적용되지 않는 느낌이다. 파이어폭스용으로 던진 서식만 적용되는데 왜지?
- 아무튼 좀 지친 상태라 너무 깊게 들어가지는 않음

>[!done]
>- 결과적으로 1, 2번을 함께 적용함
>- 추가로 발생한 문제로, `scrollbar-gutter`를 추가하면서 스포일러 경고 영역도 오른쪽에 스크롤바를 위한 영역을 남겨두는 현상이 발생
>	- 살짝 야매를 부렸다. 스포일러 영역에 `-inset-x-3` 정도를 부여해서 스크롤바에 해당되는 영역까지 덮게 만들었다. 이렇게 구현할 경우 x축으로 스크롤바가 생기기 때문에 이것에 영향을 받는 요소들에 `overflow-x-hidden`으로 가로 스크롤을 막았다.


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