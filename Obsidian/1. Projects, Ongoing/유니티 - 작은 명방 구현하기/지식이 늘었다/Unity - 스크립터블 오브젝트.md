
> - 간단 요약 
> 1. `ScriptableObject`가 `health`라는 필드를 직접 갖는다
> -> 객체마다 이 값을 이용해 초기화해야 한다. 
> 2. `ScriptableObject`가 `struct`를 갖고 `struct`가 `health`라는 필드를 갖는다
> -> `ScriptableObject.struct`를 초기화해주는 것만으로 해당 객체만이 갖는 `health` 가 생긴다



- `Unity`에서 제공하는 데이터 컨테이너 클래스.


- 유니티에서 사용되는 데이터 저장 및 관리를 위한 클래스이다.
- 프로젝트 내에서 `에셋`으로 저장되며, `씬`에 독립적이다.

- 장점
	- 메모리 효율성 : 여러 인스턴스에서 동일한 데이터를 공유할 수 있어 메모리 사용을 줄일 수 있다
	- 데이터 일관성
	- 편집 용이성 : 유니티 에디터에서 직접 데이터를 편집할 수있다
	- 모듈화

- 생성 방법
	- `C#` 스크립트를 만든 뒤 `ScriptableObject`를 상속받는다.
	- `[CreateAssetMenu] Attribute`를 사용해 Unity 메뉴에서 생성할 수 있게 한다.

## 실사용
- 런타임 중 변하는 값에는 직접 사용하면 안된다. 초기화에는 이를 이용해 필드를 정의하되, 런타임 중에 변하는 값은 정의한 필드를 사용하는 방식을 취해야 한다.

- 예를 들면 `data.stats.health` 라는 값이 있다고 하자. 
	- 게임 중이 아닐 때, 이는 `최대 체력`을 나타내는 스탯이 된다.
	- 게임 중일 때, 저 값을 사용해 `현재 체력`과 `최대 체력`을 초기화할 수 있다. 그러나 **게임 중에는 `현재 체력`, `최대 체력` 모두 변할 수 있는 값이 될 수 있다.**
	- 따라서 실제 게임 중에 변하는 값을 사용한다면, `currentHealth`와 `maxHealth`라는 별도의 필드를 사용하고, 이 둘의 초기화에만 `data.stats.health`를 사용하되 게임 중에는 `currentHealth, maxHealth`만을 추적하라는 의미다.

### 그런데 구조체를 갖는다면?
- `struct` : `값`이다. 참조 변수가 아니다.
- 따라서 인스턴스에 구조체 필드를 초기화하는 것만으로, 해당 인스턴스만이 갖는 값들이 발생한다. 
- 단 구조체가 참조형 필드를 갖는다면, 이건 위의 방식과 동일하게 객체에서 별도의 필드로 구현하는 게 좋다. 예상치 못한 오류가 발생할 수 있다고 함.

- 위의 예시로, `stats` 값은 `struct`였고 `health`는 `int`였음.
	- 그렇다면 `data.stats`을 초기화하는 것만으로 `currentHealth`는 `data.stats.health`와 동일한 의미가 된다는 뜻. `maxHealth`는 별도의 필드로 정의해야 하지만.




