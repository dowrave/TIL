- Quill 자체에서 코드블럭을 제공하지만, 검은 블럭에 흰 글씨만 쓸 수 있게 해줄 뿐 코드 하이라이트 기능은 제공하지 않는다. 이를 적용한 과정.
- `highlight.js`를 적용하라고 공식 문서에도 나와 있는데, ReactQuill에 대한 건 아닌 것 같다. 관련해서 다룬 글도 있었는데, 나한테는 적용되지 않았다. 왜인지는 몰?루.

## 기본 방법

### 1. `highlight.js`를 설치함
```sh
npm install highlight.js
```

### 2. 불러오고 적용하기
```tsx
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';  // 테마는 다양하게 있으니 그 중 하나를 선택한다.

    const modules = {
            syntax: {
                highlight: text => hljs.highlightAuto(text).value
            },
            ...
    }

	return (
	<div>
		<ReactQuill
		ref = {quillRef}
		theme = 'snow'
		modules = {modules}
		formats = {formats}
		value={content}
		/>
	</div>
	)

```
> `main.tsx`에 지정하라는 말도 있고, `highlight.js`를 적용할 컴포넌트에 불러오라는 말도 있는데, 내 경우 둘 모두 지정했다. **`main.tsx`에만 지정하면 `highlight.js`를 등록하라는 오류와 함께 작동하지 않았음.**


#### 기타 이슈
- `modules`에 `syntax: true`를 넣으라는 말도 있었는데 나한텐 해당 없었음
- `ReactQuill` 컴포넌트에는 `onChange`를 지정할 수 있다. 원래는 상위 컴포넌트로부터 `setContent` 를 받은 다음 `onChange`에 적용했으나, 그렇게 구현한 경우 `syntax`와 충돌이 발생하는지 1글자라도 타이핑하는 등 위 컴포넌트가 다시 렌더링되는 상황이면 컴포넌트가 아예 사라지는 문제가 발생했음. <-- **이거 때문에 며칠 허비함**

