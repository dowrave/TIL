- 스타일을 만지다 보니 이런 걸 발견했다.
```css
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
```
- 이전에도 이런 걸 보고 그냥 지나친 적이 있는데, 어떤 의미인지 궁금해서 chatGPT한테 물어봄

---
CSS에서 `font-family` 속성을 사용하여 글꼴을 지정할 때, 여러 글꼴 이름을 나열하는 것은 브라우저가 해당 글꼴들을 순서대로 찾아 사용하도록 지시합니다. 이렇게 글꼴을 나열하는 이유는 다음과 같습니다:

1. **폴백 메커니즘(Fallback Mechanism)**: 첫 번째 글꼴이 사용자의 시스템에 설치되어 있지 않은 경우, 브라우저는 리스트에서 다음 글꼴을 찾아보고, 그것도 없으면 다음 글꼴을 찾는 식으로 계속 진행합니다. 이러한 방식을 통해, 어떤 사용자의 시스템에서든 최적의 글꼴을 표시할 수 있습니다.
    
2. **글꼴 이름에 따옴표 사용**: CSS에서 글꼴 이름에 공백이 포함되어 있거나, 여러 단어로 구성된 경우 따옴표(`""`)로 묶어야 합니다. 예를 들어, `"Courier New"`와 같이 공백이 포함된 글꼴 이름은 따옴표로 묶어줍니다. 반면, `Monaco`나 `Consolas`와 같이 한 단어로 된 글꼴 이름에는 따옴표가 필요하지 않습니다.
    
3. **특정 글꼴에서 최적의 표현을 위해**: 때때로 특정 글꼴은 특정 문자나 스타일을 더 잘 표현할 수 있습니다. 디자이너나 개발자는 이러한 글꼴을 우선적으로 사용하도록 지정하면서, 해당 글꼴이 없는 환경을 고려해 폴백 옵션을 제공합니다.
    
4. **범용성과 호환성**: 모든 사용자의 시스템에는 동일한 글꼴이 설치되어 있지 않기 때문에, 여러 글꼴을 나열함으로써 다양한 운영체제와 환경에서도 비슷한 사용자 경험을 제공할 수 있습니다.
    

예시에서 나열된 글꼴들은 주로 모노스페이스(monospace) 글꼴로, 각 문자가 동일한 너비를 차지하는 것이 특징입니다. 이러한 글꼴은 코드 편집기나 개발 환경에서 일반적으로 사용됩니다.