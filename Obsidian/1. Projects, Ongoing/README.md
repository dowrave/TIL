# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방
## 남은 작업 내용 
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

# 250703 - 짭명방
## 구현 예정 
- 세부 직군은 구현하지 않을 거니까 이 자체가 하나의 직군이 됨. `Artillery`와 비슷.
- 지상에서 평타 1번에 2번의 공격이 나가는 클래스 만들기. 
	- `DualBlade`가 가장 무난한가? 2번 때린다는 의미에서? 
	- 기존 게임의 소드마스터는 단어만 들었을 때는 그런 이미지는 아니니까. 

- `DualBlade` 아이콘 만들기
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