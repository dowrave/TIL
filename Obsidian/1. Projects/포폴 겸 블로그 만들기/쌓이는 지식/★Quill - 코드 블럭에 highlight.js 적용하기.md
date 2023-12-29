- Quill 자체에서 코드블럭을 제공하지만, 검은 블럭에 흰 글씨만 쓸 수 있게 해줄 뿐 코드 하이라이트 기능은 제공하지 않는다. 이를 적용한 과정.
- 관련해서 다룬 글도 있었는데, 나한테는 적용되지 않았다. 왜인지는 몰?루.

## 기본 방법

### 1. `highlight.js`를 설치함
```sh
npm install highlight.js
```

### 2. 불러옴
```tsx
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';  // 테마는 다양하게 있으니 그 중 하나를 선택한다.
```
> `main.tsx`에 지정하라는 말도 있고, `highlight.js`를 적용할 컴포넌트에 불러오라는 말도 있는데, 내 경우 둘 모두 지정했다. **`main.tsx`에만 지정하면 `highlight.js`를 등록하라는 오류와 함께 작동하지 않았음.**

### 3. 사용할 언어를 등록한다
```tsx
import javascript from 'highlight.js/lib/languages/javascript';
import typescript from 'highlight.js/lib/languages/typescript';
import python from 'highlight.js/lib/languages/python';

const QuillEditor: React.FC = () => {
    const languages = [
        { name: 'jsx', module: javascript },
        { name: 'tsx', module: typescript },
        { name: 'python', module: python },
    ]
    
    useEffect(() => {

        hljs.configure({
            languages: ['jsx', 'tsx', 'python', 'sql', 'csharp', 'cpp', 'html', 'css']
        })
        languages.forEach(lang => {
            hljs.registerLanguage(lang.name, lang.module);
        })
    }, [])
```
> - `hljs.configure` : 어떤 언어 모듈을 우선적으로 사용할지를 지정한다.  등록할 경우 굳이 필요하진 않은 듯?
> - **`hljs.registerLanguage` : 사용할 언어를 등록한다.**

### 4. 코드블럭에 적용한다

#### 시행착오 - highlightAuto를 적용한 케이스
```tsx
import debounce from 'lodash/debounce';

	...
	useEffect(() => {
        if (quillRef.current) {
            const editor = quillRef.current.getEditor();
            
            // 타이핑 완료 후에만 하이라이팅을 적용하는 debounce를 이용한다.
            // innerHTML은 렌더링을 다시 하게 하고, 커서 위치가 초기화되는 문제가 있음
            const debouncedHighlight = debounce(() => {
                const htmlContent = editor.root.innerHTML;
                document.querySelectorAll('pre.ql-syntax').forEach((block) => {
                    const code = block.textContent;
                    const result = hljs.highlightAuto(code)

                    block.innerHTML = `<span className="langauge-label float-right">${result.language}</span>
                    ${result.value}`;
                    
                })
                setContent(htmlContent);
            }, 1000);
            
            editor.on('text-change', debouncedHighlight);
        }
    }, [quillRef, setContent])
```
> 처음에는 코드를 자동으로 인식하면 좋겠다고 생각해서 `highlightAuto`를 이용했다.
> - 시행착오 내역
> 1. `innerHTML`을 이용하면서 렌더링을 반복하면, 커서는 텍스트 영역의 최초 위치로 계속 옮겨진다. 따라서 **커서가 고정되고 글씨가 오른쪽으로 밀리는 현상**이 발생함.
> 2. `lodash.debouce.debouce`란, 아무 입력도 없을 때에만 해당 코드가 작동하게끔 한다. 단위는 `ms`인데, 1초도 솔직히 짧다.

- 따라서 굳이 이렇게 번거롭게 넣을 바에는, 코드 블럭을 생성할 때 어떤 언어로 생성할지를 지정하게 하는 편이 더 사용자에게도 편리하고 구현하는 입장에서도 편리할 것이라 생각했다.

#### 