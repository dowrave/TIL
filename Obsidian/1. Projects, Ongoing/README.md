## 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

## 작업 내용 : 짭명방
### 남은 작업 내용 
- 1-3 밸런싱, 보스 추가
- `Artillery` 2스킬 정리
- 테스트 및 수정

- 보고 참조할 내용
	- [[오퍼레이터들 스탯 정리]]
	- [[적 스탯 정리]]
### 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`
### 하고 싶은데 못할 듯
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.

## 블로그



# 현재 작업 일지


## 250625



## 250624 - 짭명방
- 모니터 바꾸고 게임 프로젝트 다시 꺼내는 건 처음인데, 색감이 좀 많이 뿌얘졌다.. ㅋㅋ;
	- 메인 모니터에는 돈을 아끼지 말아야지,,,

- 오늘의 교훈
	- `느슨한 결합`이 보통 권장되는 구조이긴 하다.
	- 하지만 게임의 핵심 로직은 상호작용하는 두 객체의 `public` 메서드를 서로 호출하는 방식이 나중에 유지보수하기 더 편할 수도 있다. 
	- 오늘 애먹었던 저지 로직 같은 경우, "`Operator`가 `Enemy`를 저지할 때 `Enemy`는 자신을 저지하는 `Operator`를 등록한다"는 로직 자체는 이벤트 없이 둘 간의 상호작용으로 구현해도 무방함.
	- 이벤트는 일반적으로 `1:多` 구조에서 사용하면 좋다. 
### 고칠 때마다 자꾸 어딘가 에러가 나서 일단 원상복구
- 흠...

### 작업 완료

#### 실제 적용 X) 원래 커밋으로 되돌림

- `Enemy`가 죽어서 `null`인데 `SetCurrentTarget`에서 계속 참조하려고 하는 문제
```cs
CurrentTarget = enemiesInRange.OrderBy(E => E.GetRemainingPathDistance()).FirstOrDefault();
```
> 가장 간단하게 해결하는 방식은 `enemiesInRange.RemoveAll(enemy => enemy == null);`을 넣어서 항상 우선 `null`을 제거해주는 식으로 구현하면 됨

더 나은 방식으로 이벤트를 이용하는 방식이 있다.
그런데 나는 
```
1. 이미 스폰된 Enemy들에 대해 모든 배치된 오퍼레이터들에게 해당 이벤트 구독
2. Enemy가 스폰될 때마다 모든 배치된 오퍼레이터들에게 해당 이벤트 구독
3. 오퍼레이터를 배치할 때 이미 스폰된 Enemy들에 대해 해당 이벤트 구독
```
이렇게 작업해야 하는 게 아닌가 생각했는데, **클래스 단위의 이벤트를 구독하는 `static event`를 사용하는 방법이 있다.** [[Unity - 클래스 단위의 이벤트]]
- **`static event`** 로 이벤트를 생성하면, 해당 이벤트는 모든 영역에서 단 1개만 존재하는 이벤트가 된다. 이 경우 외부에서 구독할 때 인스턴스 참조가 필요하지 않고, 클래스 이름으로 직접 접근한다.

- 이거를 기존에 몰랐기 때문에 원래 코드에서 수정할 수 있는 부분도 많을 것 같음. 일단은 보이는 부분만 고친다. `디커플링 작업`임.
	- `Enemy` 사망 시 `Die`에 만든 부분들을 `Operator, Tile`에서 이벤트 구독 메서드에 구현
		- `StageManager`에서도 추가.
	- `Enemy`가 타일에 진입 / 이탈할 때에도 추가.
		- 이 이벤트를 받는 메서드에서는 다시 `Operator`에 전달하기 위한 이벤트를 발생시켜서 해당 `Enemy`가 이 `Tile`에 접근했음을 알린다.
		- `CurrentAttackableTiles`을  `HashSet`으로 바꿈 - `List`에서 쓰는 메서드와 큰 차이도 없다.
			- `HashSet`을 쓸 때, **저장은 List로 하되 런타임에서는 HashSet을 쓰는 게 표준**이다. 
	- `Operator`의 `Enemy` 저지 로직 : 오퍼레이터**가** 적을 저지하는 것이므로 `Operator`에 구현



- 이 과정에서 발생한 문제들
	 - [ ] 저지수 0인데 저지하지 않는 문제
		 - 저지가 동작했는데 적이 그냥 빠져나가는 현상이 있음.
		 - `Melee`의 경우 저지가 됐을 때에만 공격해야 하는데, 저지되지 않고 그냥 지나가는 상황에서 때리고 지나가는 현상도 있음.
		 - AI한테 다시 던져봄. 무려 3분 간의 추론 끝에 나온 결론.
			- **"느슨한 결합은 좋은 구조고 강한 결합은 나쁜 구조다"** 라는 편견은 버리고 진행하는 게 좋다. 
				- 지금의 상황에서 저지 로직은 `Operator`와 `Enemy` 간의 직접적인 상호작용이 일어나는 게 더 직관적이고 바람직한 상황임.
				- **이벤트 기반의 구조가 항상 좋은 건 아니다.** 암시적인 구조이기 때문에 수정이 필요한 상황에선 오히려 더 복잡할 수도 있다. 
			- **이벤트를 활용하는 상황은 `1:多` 상황이라고 생각하면 편할 듯.**
				- 결국 유지보수를 할 때 어떻게 작성해놔야 나중에 수정하기 편할지를 생각해보면 될 것 같다.
		- 따라서 저지 로직은 기존 구현대로 돌려놓는다. 

	- [ ] 킬 카운트 여러개씩 올라가는 문제
		- 1개가 올라갈 때도 있고, 동시에 여러 개가 올라갈 때도 있다.
		- 이건 추정인데 투사체가 체력 0 이하인 `Enemy`에 부딪힐 때 이벤트가 여러번 발생하는 거 아닐까?
		- 일단 `TakeDamage` 메서드가 `CurrentHealth <= 0`일 때는 실행되지 않게 수정함
		- 유사한 문제가 발생하고 있지는 않아서 일단 이렇게 둔다. 위의 이슈가 해결되지 않고 있음.

- **그냥 이전 커밋으로 되돌려놨음** : 저지 로직이 정상적으로 동작하지 않는 이슈가 치명적

- `Melee`라서 본인이 저지당하지 않았을 때는 때리면 안되는데 때리는 현상이 있다.
- 같은 원인으로 보이는데 저지 당할 상태가 아닌데 저지되는 현상이 있다.

## 250623 - 블로그

### 작업 중

### 작업 완료
- 구현한 기능을 직접 사용해보면서 어색하거나 추가되었으면 하는 기능을 정리해봤다.
#### 1. '보는 중' 필터링 추가
- 기존 헤더에 있는 `보는 중`이라는 범례를 버튼으로 활용, 이것도 필터링으로 사용한다.
	- `스포일러 포함`도 범례에 있는데, 이건 필터링으로 만들 필요는 없을 듯.
- 기존의 카테고리 필터와는 별개의 기능으로 구현.

--> `스포일러 포함`은 버튼으로서의 기능을 넣지 않았고 `보는 중`은 버튼으로서의 기능을 넣어봤다. 우선 헷갈릴 여지가 있다. 범례를 클릭할 수 있다면 둘 모두 클릭 가능하거나 둘 모두 클릭 불가능해야 하는 게 일반적이니까.
- 그래서 범례는 범례로만 두고, 기존의 `ReviewCategoryFilter`에 "보는 중"으로 기능의 위치를 두겠음. 
- `ReviewCategoryFilter -> ReviewFilterBar.tsx`로 이름 수정, 왼쪽에 필터 버튼 추가했음.

- 백엔드에서도 기존 조회하던 기능에서 아래 기능만 추가하면 됨
```python
class ReviewViewSet(viewsets.ModelVietSet):
	
	# ...
	
	def get_queryset(self):
	
		# ...
		
		# "보는 중" 필터링 추가
		is_watching_param = self.request.query_params.get('is_watching')
		if is_watching_param == 'true':
			queryset = queryset.filter(is_watching=True)
```

#### 2. 제목 외의 검색어로 기능할 수 있는 것(별칭) 추가
- 프론트엔드에서 해당 부분을 입력받을 수 있게 하려고 함.
- `,`로 구분되는 입력을 받는 방식도 생각했는데, AI한테 물어보니까 `태그`로 구현하는 방법이 괜찮아 보임. 

백엔드 1. **`Review` 모델을 외래키로 갖는 `ReviewAlias` 모델을 추가**
```python
# 리뷰 별칭 모델
class ReviewAlias(models.Model):
    review = models.ForeignKey(
        Review,
        on_delete = models.CASCADE, # 리뷰가 삭제되면 별칭도 삭제
        related_name='aliases' # Review 모델에서 .aliases로 접근 가능
    )
    name = models.CharField(verbose_name="별칭/약어", max_length=100)

    def __str__(self):
        return f'{self.review.title}의 별칭: {self.name}'
    
    class Meta:
        verbose_name = "리뷰 별칭"
        verbose_name_plural = "리뷰 별칭들"
```
> `ForeignKey`의 `related_name`이 어떻게 동작하는지 알아보자 [[Django - 역참조]]
> - 주석에 적힌 것처럼 `Review.aliases`로 어떻게 사용할 수 있게 되는가? 가 궁금했는데 저건 추상화의 개념이고 **실제로는 `ReviewAlias` 모델에서 `Review`와 일치하는 데이터들을 찾는 구조로 돌아간다.** 즉 `Review` 테이블에 대한 쿼리가 돌아가는 게 아니라는 뜻. 그걸 `Review.Aliases`로 쉽게 이용할 수 있게 한 것 뿐이다.

- 모델을 추가/수정했으니 마이그레이션도 해준다. 
	- 내 세팅에서 `migrate`는 이미지 올릴 때 자동으로 됨
```sh
python manage.py makemigrations reviews
```

백엔드 2. **`serializers.py`**
- 생성 / 수정 시 별칭 목록을 한번에 받아서 처리하도록 `ReviewCreateUpdateSerializer`를 수정한다.
```python
aliases = serializers.ListField(
	child=serializers.CharField(max_length = 100),
	write_only=True, 
	required=False
)
```
> 이외에도 `aliases` 필드가 실제로는 `Review` 모델에 없기 때문에 이들은 `Create, Update` 메서드를 오버라이드해서 `Review`에 해당하는 필드들은 `Review` 모델에 저장하고 `ReviewAliases`에 해당하는 필드는 `ReviewAliases`에 저장하도록 구분해야 한다. 

백엔드 3. **`views.py`**
- `ReviewViewSet`에 아래 요소 추가
```python
class ReviewViewSet(viewsets.ModelViewSet):
    search_fields = [
        'title', 
        'aliases__name' # aliases로 연결된 ReviewAlias 모델의 name 필드를 검색 대상에 포함, 자동으로 JOIN과 DISTINCT를 사용해 중복이 없다.
    ] 
```

프론트 1. `reviewApi.ts`
```ts
export interface ReviewItem {
	// ...
    aliases: Alias[]; // 추가
}

export interface Alias {
    id: number;
    name: string;
}
```

프론트 2. `ReviewForm.tsx`
```tsx
const ReviewForm: React.FC<Props> = ({ initialData, onSubmit, isSubmitting }) => {
  const [aliases, setAliases] = useState<string[]>(initialData?.aliases.map(a => a.name) || []);
  const [aliasInput, setAliasInput] = useState('');
	
  // ...
  
  // 별칭 입력 후 엔터/컴마를 눌렀을 때 별칭 목록에 추가
  const handleAliasKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' || e.key ===',') {
      e.preventDefault();
      const newAlias = aliasInput.trim();
      if (newAlias && !aliases.includes(newAlias)) {
        setAliases([...aliases, newAlias]);
      } 
      setAliasInput('');
    }
  }

  // 별칭 제거
  const removeAlias = (indexToRemove: number) => {
    setAliases(aliases.filter((_, index) => index !== indexToRemove))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    // ...
  
    aliases.forEach(alias => {
      formData.append('aliases', alias);
    })
    // ...
  }
```
> 아래의 렌더링 부분에는 입력란을 만들고 엔터나 컴마를 누를 때 해당 태그가 Aliases에 들어가는 방식으로 구현되어 있음

백엔드 4. 프론트엔드에 `aliases` 배열을 보내줘야 하므로 읽기 전용인 `ReviewSerializer`에도 `aliases` 필드를 추가한다.
```python
class AliasSerializer(serializers.ModelSerializer):
    """별칭을 위한 시리얼라이저"""
    class Meta:
        model = ReviewAlias
        fields = ['id', 'name']

class ReviewSerializer(serializers.ModelSerializer):
	aliases = AliasSerializer(many=True, read_only=True)
	# ...

	class Meta:
        model = Review
        fields = [
            'category',
            'title',
            'release_year',
            'poster_image',
            'rating',
            'content',
            'is_spoiler',
            'is_watching',
            'aliases' # 추가
        ]
```

#### 3. 목록에 나타나는 개별 카드 크기 줄이기
- 아래 4번에 한꺼번에 정리.
#### 4. 한꺼번에 목록에서 불러오는 갯수 조절하기
- sm, md, lg, xl에서 각각 3, 4, 6, 8개가 나타나도록 조절. **각 카드의 크기는 `grid`의 요소들에 의해 조정된다.**
- 불러오는 갯수의 경우, **백엔드의 페이지네이션에서 저들의 최소공배수인 24개**로 설정했음.
```python
class ReviewPagination(PageNumberPagination):
    page_size = 24  # 페이지 당 리뷰 개수
    page_size_query_param = 'page_size'  
    max_page_size = 100  # 악의적인 요청에 대비한 제한 사항. 999999같은 요청을 방지함.
```


#### 5. 카드 서식 일치시키기
- AI한테 좀 복잡하게 설명했는데 찰떡같이 이해했고 예시도 줬다.
```
[ 현재 상태 ]
+-----------+  +-----------+  +-----------+
|   Image   |  |   Image   |  |   Image   |
+-----------+  +-----------+  +-----------+
| Title (1) |  | Title (2) |  | Title (1) |
|           |  | Line 2    |  |           |
| Meta      |  +-----------+  | Meta      |
| Date      |  | Meta      |  | Date      |
+-----------+  | Date      |  +-----------+
               +-----------+

[ 원하는 결과 ]
+-----------+  +-----------+  +-----------+
|   Image   |  |   Image   |  |   Image   |
+-----------+  +-----------+  +-----------+
| Title (1) |  | Title (2) |  | Title (1) |
| (빈 공간) |  | Line 2    |  | (빈 공간) |
| Meta      |  | Meta      |  | Meta      |
| Date      |  | Date      |  | Date      |
+-----------+  +-----------+  +-----------+
```
> 제목이 여러 줄이 오는 카드와 같은 줄에 있는 제목이 한 줄인 카드의 경우, 카드의 서식이 저런 식으로 흐트러지는 현상이 있음.

- `원하는 결과`를 위한 아이디어 : **`제목 / (메타데이터 + 날짜)`로 컨테이너를 구성하고, 두 그룹을 모두 포함하는 컨테이너에 `justify-between`을 설정**
- `Date`를 오른쪽으로 밀고 싶다면 **부모 `div`에 `flex`를 설정하고 `Date`에 해당하는 태그에 `ml-auto`를 달면 된다.**

##### 기타
- 이미 선택된 별점을 다시 클릭하면 별점을 선택하지 않은 상태로 되돌림.
- 연도도 선택 사항이므로 디폴트는 아무것도 선택되지 않은 상황으로 설정.
- `quill-markdown` 제거
	- 다른 데서 작성한 거 붙여넣기 할 때 `-`이 들어간 부분이나 `1. 2.` 처럼 작성한 부분이 계속 이상하게 들어가는 문제가 있어서 제거.
	- 해당 기능들은 Quill의 메뉴 부분에 있으니 그런 식으로 사용하면 됨.



## 250620 

### 짭명방
#### 작업 중

#### 작업 완료
##### 스냅핑이 안되는 문제 또 발생

- 이전에 알아낸 방법으로 시행착오
	1. 이전에 찾아낸 방법처럼 `Reimport All`을 눌러봤다.
		- 안된다. 비상!!!
	2. 에디터를 끄고 라이브러리 지우고 다시 실행
		- 역시 안된다. 하...
	- 또 날 잡겠다.
	- `Tile` 프리팹에 `공격 범위`를 추가했을 때 해당 부분의 레이어가 `Default`로 지정되어 있었다. `Tile`로 수정. 
		- 물론 크게 상관없을 가능성이 높다. 하이라이트되는 타일은 머티리얼을 수정하는 방식이지 레이어가 바뀌진 않을 거임
	- 에디터 다시 불러올 때 처음에 Tile과 관련된 에러들이 뜨는 현상이 있었다. 거기부터 만져봄.

- **과거 깃으로 돌아간 다음 커밋들을 비교해보고** 기능이 동작하지 않는 지점 직전의 커밋에서 수정된 내용들 하나하나 반영하면서 어디서 문제가 발생했는지 테스트해보겠음. 
	- 이전에 갔다 왔을 때 정상적으로 작동하는 브랜치가 있었음.
	- [[git - 과거로 돌아가서 새로운 브랜치로 옮기기]] 에 정리해뒀다. 기존의 `reimport all`이나 라이브러리 삭제로 되지 않아서 결국 이게 맞는 듯...

- 진행
	- `250609`와 `250606` 사이에 문제가 발생한 것으로 보인다. 6월 9일의 커밋에서는 동작하지 않고, 6월 6일의 커밋에서는 스냅핑이 동작하고 있음.
		- 날짜 적어 놓는 거 잘한 듯? 보기 편하다.
	- 그래서 6월 6일로 `checkout`한 다음 `git checkout -b main-fixed`로 시작함.
	- `cherry-pick` 자체는 불가능한 상황이다. 대신 모든 커밋은 깃허브에 오리진으로 커밋했기 때문에, **일일이 변경 사항들 반영**하는 방식으로 진행함. 
	- 6월 6일과 6월 9일 커밋 사이에 발생한 건 Operator와 Enemy의 타겟 선정 로직이었던 걸로 기억하는데.. 저 날의 작업 기록을 참고로 2단계에 걸쳐서 반영해본다. 
		1. 우선 `Operator`의 타겟 선정 로직을 수정했을 때, 스냅핑 동작은 여전히 잘 동작한다.
			- 유니티를 껐다가 켜서 잘 동작하는지까지 확인했음.
			- 아직 `main-fixed`인 상태이기 때문에, 이 상태로 커밋을 저장해봄.
		2. `Enemy`의 타겟 선정 로직 작업
			- `Enemy`의 자식 오브젝트로 `EnemyAttackRangeController`을 추가함. 여기엔 `Sphere Collider`도 같이 들어가 있어서 공격범위에 해당하는 반지름을 설정한다.
			- `Ranged`가 공격하지 않는 문제가 있었지만, 콜라이더의 `Is Trigger`가 꺼져있는 문제였음
			- 수정 직후 실행은 잘 동작함.
			- 유니티를 껐다가 켰을 때에도 잘 동작함. 
		- 250609의 수정사항을 다 반영했는데, 잘 동작한다. ???? 정작 저 커밋으로 돌아갔을 때는 동작하지 않았음.
		- 오리진에 마지막으로 커밋한 것까지 일일이 수정했고, 유니티를 껐다 켰을 때 최종 동작만 확인해봄 : 잘 됨
	- 일단 이 상태로 커밋까지 해두고, 오리진에도 이 브랜치로 덮어씌움.
	- 이랬는데 또 문제 발생하면 ㅡㅡ 하

### 블로그
- 그저께 블로그를 테스트하던 중에 흥미로운 현상이 있어서 정리한다.
- 로컬 호스트에서`isAdmin`을 켜서 글 작성을 가능하게 한 다음, Form의 양식에 맞춰 제출했을 때 정상적으로 동작했다. 로컬에선 로그인을 하지 않았는데 이게 기능한다는 게 이상했음.

- 정확한 상황은 아래와 같다.
	1. 띄운 사이트에서 로그인함
	2. 이 상태에서 로컬에서 테스트 폼을 만들고 리뷰 글을 하나 제출함
	3. 제출됨.
	4. 띄운 사이트에서 로그아웃했다면, 로컬에서 2번처럼 진행했을 때는 제출되지 않음.
#### 왜?

1. **쿠키 기반 인증 방식**
	- 실제 사이트에 로그인하면, 백엔드는 `인증 세션 쿠키`를 생성한다. 이 쿠키는 프론트엔드 도메인에 종속된다.
	- 이후 프론트엔드 도메인에서 백엔드 API로 요청을 보내면 쿠키를 함께 첨부한다. 백엔드는 이 쿠키를 보고 POST, PUT, DELETE 요청을 허용한다.
2. 로컬 환경
	- 로컬의 프론트엔드는 실제 사이트와 다른 출처`Origin`이다.
	- 프로토콜, 호스트, 포트 중 하나라도 다르면 다른 출처임.
3. **CORS, withCredentials**
	- `axios` 설정의 `with Credentials: true` 설정은 "다른 출처로 API 요청을 보낼 때 현재 브라우저에 저장된 쿠키를 함께 첨부해줘"라는 의미다.
	- 백엔드에선 사전에 `django-cors-header` 등의 라이브러리에 `Access-Control-Allow-Credentials: true`가 설정되어 있을 것이다. 이는 "다른 출처에서 요청이 오더라도 쿠키를 포함하는 걸 허용한다"는 의미다.

- 전체적인 과정
1. 로그인 : 내 블로그 도메인에 대한 유효한 관리자 세션 쿠키가 저장됨
2. 로컬호스트에서 POST 요청을 보낸다. `withCredentials: true` 때문에 브라우저는 이 요청에 블로그 도메인의 쿠키를 첨부하려고 시도한다. 
3. 백엔드는 요청과 함께 온 쿠키를 보고 유효한 관리자라고 인식, 요청을 처리한다. **글 생성이 성공한다.**
4. 하지만 요청을 보낸 출처와 받는 곳이 `localhost`라서, 이 과정에서 브라우저의 `Same-Origin Policy`나 다른 `CORS` 관련 보안 정책 때문에 `axios`가 정상적인 응답으로 처리하지 못하고 네트워크 오류를 발생시킬 수 있다. 즉, **요청은 성공했지만 응답을 받는 과정에서 에러가 발생한다.**

위 과정은 로그아웃 시에는 브라우저의 쿠키가 없거나 만료되므로 `401, 403` 에러를 발생시켜 글 생성이 실패한다.

내가 걱정했던 부분은 **`isAdmin = true`로 외부에서 임의로 바꿔서 해당 요청을 보내는 식으로 접근했을 때 이게 보안에 이슈가 될 수 있는 부분인가?** 였다.  누군가 그런 식으로 접근해도, **해당 브라우저에는 백엔드에서 일치해야 하는 세션 쿠키를 갖고 있지 않은 상태이기 때문에 `return [permissions.IsAdminUser()]`에서 100% 차단된다.** 실제로 올린 사이트는 `useSelector`를 사용하고 있기 때문에 어드민 계정이 털리지 않는 이상에는 괜찮다고 보면 될 듯.






## 250618 - 블로그
### 1. 리뷰 제목 검색 기능
- Django에서 `Reviews/views.py/ReviewsViewSet(viewsets.ModelViewSet)`에 아래 요소 추가
```python
    # 검색 기능 추가
    """
    GET /api/reviews/items/?serach=검색어 로 요청이 들어오면 
    title 필드에 '검색어'가 포함된 리뷰만 필터링되어 반환한다.
    """
    filter_backends = [SearchFilter]
    search_Fields = ['title'] # 'title' 필드를 기준으로 검색
```

- 프론트에선 `fetchReviews`에 검색어를 의미하는 `searchTerm` 파라미터를 추가
```tsx
export const fetchReviews = async ({ pageParam = 1, categoryId, searchTerm }: {
    pageParam?: number, 
    categoryId: number | null, 
    searchTerm?: string // 이거
}): Promise<PaginatedReviewsResponse> => {
	
	// category와 검색어는 없어도 잘 동작함
    const categoryQuery = categoryId ? `&category=${categoryId}` : '';
    const searchQuery = searchTerm ? `&search=${searchTerm}` : ''; // 이거

    const { data } = await axios.get(`${backendUrl}api/reviews/items/?page=${pageParam}${categoryQuery}${searchQuery}`); 
    return data;
}
```
### 2. Navbar의 기존 게시판은 Posts로 묶기 / UI 수정
- 작업 완료.

### 3. 추가로 생각났거나 수정할 요소들
1. 작품을 다 감상하지 않았지만 글은 적어둘 수 있다. 일명 '**보는 중'인 상태.** 
	- 모델 수정 : 카테고리, 작품 이름만 필수 / 
	- `스포일러 기능` : 실질적으로 모든 리뷰에 스포일러를 기록해두는 상태이긴 한데, 추가는 해두자.
	- 그러면 실질적으로 필수 칸은 **카테고리와 작품 이름**밖에 없을 듯?

2. 제목을 활용한 검색 기능이 제대로 동작하지 않고 있음.
	- 위의 `search_Fields`는 `search_fields`여야 한다.

3. 글을 생성했을 때 '모두'에 나타나는 갯수가 업데이트되지 않음
	- `queryClient.invalidateQueries({ queryKey: ['totalReviewCount'] });`. `totalReviewCount`라는 요소를 새로 추가했으니 글을 생성 / 삭제하는 부분에 넣어준다. 수정은 크게 상관없는데 일관성을 위해 추가.

---
- 빌드 후 수정할 것들
	- `Card`의 경우 작은 화면에서 별점 / 스포일러 / 보는 중 표시와 날짜는 다른 줄에 나타나야 할 듯. 
		- 스포일러, 보는 중 아이콘은 아래처럼 `아이콘 + 텍스트`로 나타낼지 아니면 아예 외부로 빼서 설명을 적어놓든지?
	- `Modal`에 나타나는 스포일러 / 보는 중 아이콘은 `아이콘 + 텍스트`의 형태로 구성하기
	- `ReviewForm` 부분 : 이미지를 첨부하지 않아도 올릴 수 있어야 하는데 현재는 이미지가 필수인 상황
- 다 고쳤음


### 기타
- `전체(갯수)`에서 `갯수` 부분이 '현재 선택된 카테고리'의 숫자를 긁어오고 있었음. 어느 카테고리를 선택하든 항상 전체 리뷰 갯수를 긁어오도록 수정함.
```tsx
export const fetchTotalReviewCount = async (): Promise<PaginatedReviewsResponse> => {
    // count만 필요하므로 page_size=1로 설정하여 응답 데이터 크기를 최소화
    const { data } = await axios.get(`${backendUrl}api/reviews/items/?page_size=1`);
    return data;
}

// ReviewBoardPage.tsx
  // 전체 리뷰 개수를 가져오는 쿼리
  const { data: totalCountData } = useQuery({
    queryKey: ['totalReviewCount'],
    queryFn: fetchTotalReviewCount,
    staleTime: 5 * 60 * 1000, // 5분 동안은 캐시된 데이터를 사용
  })
```

- 모바일로도 서식 망가지지 않게 수정

### 작업 결과

- 전체 게시판
![[Pasted image 20250618183222.png]]

- 리뷰 작성 모달
![[Pasted image 20250618183325.png]]
![[Pasted image 20250618183356.png]]

- 리뷰 글 모달
![[Pasted image 20250618183446.png]]

- 발견한 이슈
	- `ReactQuill`에서와 달리, 포스트에서는 엔터가 한 번 더 들어가는 이슈가 있음. 해당 글을 다시 수정하기 등을 통해 `WriteReviewModal`에 들어갔을 때는 멀쩡하게 나타난다.
	- `WritePost`나 `PostDetail`을 참고해서 수정해보겠음.
	- 의외로 `prose` 태그를 사용하면서 생기는 문제였다. `prose`만 제거했음. 

- [[리뷰 기능 만들기]] - 함께한 제미나이가 정리해준 내용.
- 기능이 어떻게 구현되는가를 크게 공부하지 않았고, 그냥 `문제 발생 / 기능 추가 필요 -> AI에게 던져줌` 만으로 3일 컷 냈다. 
	- 블로그 만들 때 스타일이나 문제 발생 등에서 공부했던 게 도움이 됐지만.. 스타일 같은 것도 깔끔하게 잡아주는 요즘의 AI는 무섭다.
## 250617 - 블로그
- 갑자기 왼쪽 모니터가 안들어온다. 작업 중인 모니터는 8년 째 써도 멀쩡한데.. 하..
- ㅋㅋㅋㅋㅋㅋ한쪽 모니터 없이 작업하니까 매우 답답하다~

- 일단 `Not Found`부터 시작
```sh
Not Found: /api/review/categories
Not Found: /api/review/
```
```python
 re_path(r'^api/reviews/', include('reviews.urls')),
```
> 엌ㅋㅋㅋㅋ `reviews`인데 `review`로 넣었음
> 앱 이름도 `reviews`이기 때문에 `reviews`로 통합. 백엔드에서 이미지를 다시 굽는 것보다 프론트에서 처리해버리는 게 낫다.

### 모바일 그리드 작업
- 모바일 환경에선 한 줄에 2개씩 나오는데 이건 좀 많이 답답하다. 
- 테스트를 해본 결과 + 다른 사이트도 참고했을 때 3 -> 4 -> 6 정도면 적절할 것 같음.

### webp 파일 관련
- 프론트에선 그냥 잘 들어감
- 백엔드에서도 원래 `libwebp`라는 걸 별도로 설치해야 하는데, 글에 webp 파일을 첨부해도 잘 동작한 걸 봐서는 우연히 설치된 게 있어보인다. `dockerfile.aws`의 실행 환경 부분에 `libwebp-dev` 을 명시적으로 `RUN`에 넣어줌.


### 기타 작업
- 리뷰 작성하는 부분에 수정 / 삭제 기능
- 카테고리를 자동으로 등록하지 않고, 관리하는 페이지에서 별도로 등록하고 삭제하도록 함
- 리뷰 글 등록이 안되는 문제 
- 리뷰 이미지 등록이 안되는 문제 : 클라우드프론트로 연결, `settings.py`에서 s3 커스텀 도메인만 바꾸면 됨. 
### 일단 보이는 상황에서는 작업 끝난 듯?




## 250616

### 블로그 - 작품 감상문을 위한 기능 추가하기

#### 기능 개요

##### 테이블
- 필드는
	- id
	- 작품 이름
	- 작품 연도
	- 작품 이미지
	- 별점(5점 만점, 0.5점 단위)
	- 텍스트 필드
	- 글 생성일
	- 글 수정일
##### 게시판
- 프론트엔드에서는 카테고리(게임, 영상) 별로 구분해서 볼 수 있음. 기본은 "모두 보기"로 설정.
	- 영상의 경우 
		- 실사 / 애니메이션을 구분해야 할까?
		- 시리즈 / 단독 영화도 구분해야 할까?
- 기본적으로 "글 수정일"을 기준으로 내림차순, 즉 최근에 수정한 글이 가장 위로 오도록 구현
- 개별 글은 세로로 긴 이미지(포스터)를 가로 5~6개 정도로 나열함
	- 각 포스터 아래에 제목이 오고
	- 제목 아래의 왼쪽에는 평가한 별점, 오른쪽에는 최근 수정일이 나타남
- 게시판에서 `개별 글`(포스터, 제목 등이 포함된 하나의 덩어리)을 클릭하면 나타나는 글은 `모달` 형식. 
	- 즉 배경을 살짝 어둡게 처리한 상태에서 작은 창만 하나 띄우는 방식으로 구현할 예정
- 게시판은 동적으로 불러오도록 처리함
	- 스크롤이 맨 밑에 있을 때, 혹은 맨 밑 주위에 있을 때 새로운 포스터 목록을 불러오는 방식.
##### 개별 포스트(프론트)
- 제목 영역
	- 포스터 / 이미지가 왼쪽 영역을 차지
	- 오른쪽에는 아래 요소들이 표시
		- 제목
		- 별점
		- 최초 작성일자 / 최근 수정일자
- 본문 영역
	- 기존에 구현한 요소들을 이용함

#### 우려되는 포인트
- 로컬에서 테스트가 불가능한 상황이면서 새로운 앱과 모델을 추가한 상황이라면 로컬에서는 어디까지 작업해야 하는가? 
	- 일단 난 `python manage.py makemigrations`까지는 로컬에서 실행해서 마이그레이션 파일을 만들고, 이걸 실행 시에 `EC2`의 컨테이너에서 `python manage.py migrate`이 실행되는 게 맞다고 알고 있다. 기능을 다 구현해보고 이게 실제로 잘 구현되는지도 검토해봄.

#### 진행 상황
- 지금의 기능 추가는 `Google AI Studio`에 의존, **라이브러리 공부를 하지 않으면서 진행하고 있음.**
	- AI의 힘을 보고 있다.
- 백엔드에 구조를 만들어뒀다. 아직 `makemigrations`은 하지 않았다.
- 프론트에도 목업 페이지를 만들어뒀다. 글 작성 페이지도 만들어둠.

### 오늘의 활동
- 일단 일차적으로 백엔드랑 프론트엔드 구축은 다 해놨고, 이미지랑 프론트엔드 각각 빌드해서 서버에 띄워놓기는 했다.
	- AI는 강력하다. 생각보다 과거의 텍스트도 기억을 잘 하는 편인 듯.
	- 다만 너무 의존하면 실질적으로 프로그램이 어떻게 돌아가는지에 대해서는 거의 파악을 못할 것 같음. 이런 건 나중에 유지보수할 때 이슈가 될 수 있을 거라고 생각한다.

> 하지만 아직 실질적으로 등록이 되는 상태는 아니다.
> 1. `Not Found` 에러가 발생함 : 카테고리 등록의 이슈일 수도 있고 작품 등록의 이슈일 수도 있음.
> 2. `webp` 파일도 바로 등록할 수 있으면 좋을 듯? 대부분의 파일을 왓챠피디아에서 가져오는데 `다른 이름으로 저장`하면 `webp` 파일로 가져온다.
## 06.12 ~ 14 여행 다녀옴
## 250611 - 짭명방

### 작업 중

#### 코드 정리
- `ClickDetectionSystem`에서 기능 메서드별로 분리 및 메서드 이름 정리

#### 오퍼레이터 스냅핑이 갑자기 되지 않는 현상 해결하기
- [x] 어제 했던 거에서 계속 : 오퍼레이터 스냅핑이 되지 않는 현상
	- 아예 클릭 시스템부터 다시 보겠음. 
		- 클릭 시스템은 이 케이스와는 별개다. `DeployableBox`에서 드래그 중인 상태는 
	- 지금 버전에서는 방법을 못 찾을 것 같다. 그래서 되돌려본 다음 다시 하나하나 짚어나가겠음.

> 일단 지금 발생하고 있는 문제가 없는 브랜치로 돌아가는 명령어는 아래와 같다.
```sh
git checkout  c67d56cab79609f68b278e0393bba0abe573b39a

// 돌아오는 건 git checkout main(현재 브랜치 이름)
```

- 위에서 정상적으로 스냅이 동작하는 브랜치로 되돌렸다가, 다시 최근 브랜치로 돌아왔는데 **갑자기 스냅이 동작한다.** ???????????????????????????????????????
	- 이게 **운으로 해결한 거라서 비슷한 상황을 막으려면 원인이랑 해결책을 알아둘 필요는 있겠다.**
	- 한가지 재밌는 건 직전 커밋은 스냅 이슈가 발생하는 상황이라고 쳐도 그 이전 커밋으로 되돌렸을 때도 스냅이 동작하지 않았다. 스테이지 테스트를 하고 있기 때문에 배치가 안 되는 상황이 발생했다면 이를 최소한 메모라도 해뒀을 거임.

- 안되던 상황 설명 
	- 레이캐스트가 타일 레이어에 지정되어 있고, 타일의 콜라이더의 `Is Trigger`도 켜져 있고, `Queries Hit Triggers`도 켜져 있었는데도 레이캐스트가 타일을 포착하지 못하는 이슈가 있었다. 그래서 모든 오브젝트를 포착하는 레이캐스트를 켜봤는데, 아무런 오브젝트도 포착하지 못했다. 
	- 특이한 건 `Canvas`에 있는 UI 요소는 잘 포착하는 반면 맵에 있는 오브젝트는 포착하지 못했음.
	- **에디터를 껐다가 켜도 해결되지 않았다.**
	- 정상적으로 동작했을 때는 `Operator`와 `Tile`의 레이캐스트 관련 설정은 건드리지도 않았는데 갑자기 스냅이 동작하지 않는 현상이다.

- 무엇이 원인인가?
	- **이전의 작업 과정에서 라이브러리나 캐시 파일이 깨졌을 확률이 높다.**
	- `git checkout`으로 이전 브랜치에 다녀오는 과정에서 유니티가 변경된 파일들을 다시 임포트하고 라이브러리를 재구성하게 되었다.
	- 에디터를 껐다 켜서 해결되는 상황은 메모리를 초기화하면서 해결되는 것인데, 이번 경우는 그렇지 않았기 때문에 저장된 파일 자체에 이슈가 있었다는 의미다.
	- 유니티는 효율성을 위해 라이브러리 폴더의 내용이 유효하다고 생각하고 이를 신뢰하고 재사용하려고 한다.
### 작업 완료

## 250610

### 작업 중

#### 이슈
- [ ] 오퍼레이터 배치 시에 스냅핑이 되지 않는 현상

> 아니 이거 뭐임??? 갑자기 발생했는데 뭐가 원인인지도 모르겠다.
> - 깃허브에서 어제 바꾼 내용들을 봤는데 딱히 이 부분에 영향을 갈 부분을 건드리지는 않았는데? 하루 종일 봐도 모르겠음. 
> - 특이한 건 레이캐스트를 전체 레이어에 대해 진행했을 때, **타일 오브젝트가 있음에도 타일을 레이캐스트하지 못한다는 게 있음.** 
> - 그래도 뭔지 모르겠다. 어제 바꾼 부분에서 여기는 전혀 건드린 게 없는데. 내일 계속함;;

- 시도 1) 해결 X
```cs
// 해결 X
// 타일 콜라이더에 IsTrigger가 켜져 있어서 QueryTriggerInteraction.Collide 사용
if (Physics.Raycast(ray, out hit, Mathf.Infinity, layerMask, QueryTriggerInteraction.Collide))
{
	Debug.Log("타일 정보 : " + hit.collider.name);
	return hit.collider.GetComponentInParent<Tile>();
}
```


### 작업 완료



## 250609

### 작업 중

#### 스테이지 정주행하면서 밸런스 조정
### 작업 완료

#### 리팩토링 / 수정사항
- [ ] `Operator` : 공격 범위 내의 적 선정 로직
	- 기존 구현 : `EnemiesInRange` 리스트를 계속해서 비우고 공격 범위 타일 내의 적을 수집해서 채우고 목적지까지 남은 거리가 가장 짧은 적을 타겟으로 함
	- 그런데 **계속해서 비울 필요가 있을까?** 적이 어떤 타일에 들어가거나 나갈 때 그 타일이 어떤 오퍼레이터의 공격 범위 내라면 해당 적을 `EnemiesInRange`에 넣거나 빼는 식으로 구현할 수 있지 않을까?
	1. `Tile`에 자신을 공격 범위로 하는 오퍼레이터들 리스트를 저장. 적이 타일에 진입하면 타일은 오퍼레이터들에게 공격 범위 내에 적이 들어왔음을 알린다. 나갈 때도 똑같은 원리로 제거함.
	2. 오퍼레이터는 이제 `Update`에서 타일들이 갖고 있는 적들을 가져오지 않고, 타일들이 알려주는 적들을 계속 불러온다. 
	- 이렇게만 하면 **공격 범위 내에 있는 한 타일에서 나갈 때 다른 타일로 들어갈 때 일시적으로 공격 범위에 있는 적을 잃는 현상**이 있음. 
		- 지금 콜라이더의 구현을 보면 적은 두 타일에 모두 존재하는 경우가 있는데도 이런 현상이 나타난다.
		- 그래서 **타일에서 나갈 때 오퍼레이터에서는 자신의 공격 범위 내에 해당 적이 계속 있는지를 검사**하는 방식으로 이를 보완할 수 있다.
		- 그래도 계속 안돼서 헤맸음. 조건문에 대한 생각을 잘못 한 듯.

> `Operator`가 `Enemy`를 감지하는 로직을 업데이트했듯, `Enemy`에서도 `Operator`를 감지하는 로직을 업데이트하겠음.

- [x] `Enemy` : 공격 범위 내의 `Operator` 선정 로직
- 기존 구현 : `Update`마다 `Sphere Collider` 생성 -> 당연히 최적화 문제가 있음
- 과정
	- **공격 범위를 위한 `Sphere Collider`를 추가할 것**이다. `Enemy`가 스폰될 때 `Radius`값만 수정하면 됨.
	- `Enemy`에는 `Box Collider`가 있다 : 이는 `Tile`과의 충돌을 감지해서 해당 타일 위에 `Enemy`가 있음을 나타내기 위한 콜라이더임
	- `Enemy`에 추가로 `Sphere Collider`를 구현할 수 있다. 
		- 하지만 이 경우 모호함이 발생할 수 있다. `OnCollisionEnter, OnTriggerEnter` 등을 쓸 때, 어떤 콜라이더에 의해 발생했는지가 애매하다. **상대방의 콜라이더를 쓰기 때문에 조건을 붙여야 해서 코드가 지저분해진다.**
	- 그래서 이 경우는 자식 오브젝트에 공격 범위용 콜라이더를 하나 구현하는 식으로 수정한다. 이 때, **자식 오브젝트는 부모 오브젝트에게 이벤트를 전달하기 위한 별도의 스크립트를 구현해야 한다.** 

- `Enemy` 중 `Ranged`의 공격 범위 : `3 -> 2`로 하향

## 250606

### 작업 중

### 작업 완료

- [x] 바리케이드 아이콘 추가
![[Barricade_256.png]]
- [x] 모든 `UnitEntity` 상속받는 객체에 대해 `DeathAnimation` 추가
	- 검정색으로 변함 -> 투명해짐
```cs
// UnitEntity.cs
    protected virtual void Die()
    {
        StartCoroutine(DeathAnimation());
    }

    protected IEnumerator DeathAnimation()
    {
        Renderer renderer = GetComponentInChildren<Renderer>();

        // 동일한 머티리얼을 사용하는 모든 객체에 적용되는 걸 막고자 머티리얼 인스턴스를 만들고 진행한다.
        if (renderer != null)
        {
            Material materialInstance = new Material(renderer.material);
            renderer.material = materialInstance;

            // DOTween 사용하여 검정으로 변한 뒤 투명해지는 애니메이션 적용
            materialInstance.DOColor(Color.black, 0f);
            materialInstance.DOFade(0f, 0.5f).OnComplete(() =>
            {
                Destroy(materialInstance); // 메모리 누수 방지
                OnDestroyed?.Invoke(this);
                RemoveAllCrowdControls();
                Destroy(gameObject);
            });
        }
    }
```
- 투명해지는 효과가 구현되지 않고 있음 : 머티리얼 자체가 투명도를 지원하지 않을 가능성이 있다.
	- 머티리얼의 `Surface Type`을 `Transparent`로 바꿔봤는데, **평소에는 `Opaque`였다가 죽을 때만 `Transparent`로 바꾸는 게 나아 보인다.** 관련 코드는 Claude한테 받아서 적용.
	- **검게 변하는 것도 이상하다.** 그냥 투명도만 서서히 낮추는 식으로 하겠음.


## 250605

### 작업 중

#### 1-0 ~ 1-3 정주행하면서 테스트
- `1-3` 스테이지 몹 추가 배치
- 보스는 아직 만들지 않았음

- [ ] 바리케이드 아이콘 추가

### 작업 완료

#### 발견한 이슈들 & 수정 사항
- [x] `BaseSkill`에서 `isOnTest` 필드를 하나 추가
	- 기존 스킬 비용을 나타내는 `SPCost`의 변수 이름을 `spCost`로 변경
	- 변수 `SPCost`는 `isOnTest`에 따라 1이거나 `spCost`를 가리키는 프로퍼티로 변경함.
	- 테스트하려고 SP Cost를 바꾸는 일이 있는데, 원본을 따로 기록해두지 않아서(!) 테스트할 때마다 변동사항이 생겼기 때문에 이렇게 바꿨다. 

- [x] `SPBar`의 회복되는 UI를 보면 낮은 수치일 때는 천천히 오르다가 높은 수치일 때는 빠르게 오르는 듯한 현상이 보임.
```cs
	// spFill.rectTransform.anchorMin = new Vector2(0, 0);
	// spFill.rectTransform.anchorMax = new Vector2(valueRatio, 1);
```
> **이 부분 주석처리로 해결**
> - 저렇게 설정하면 내부 Fill 값이 1이 됐을 때도 슬라이더를 가득 채우지 못함. 딱 중간쯤 온다.

- [x] 튜토리얼 진행 중 `OperatorSlot`을 제대로 찾지 못하는 문제 수정
	- `OperatorInventoryPanel` 에서 슬롯 초기화할 때 원래 슬롯의 이름도 같이 수정했는데 그 부분이 없어졌다. 그 부분을 수정.
	- **`Slot` 자체에서 초기화 시에 이름도 설정하도록 수정 완료**

- [x] `1-1` 스테이지 바리케이드 설명 추가

- [x] 3성 클리어가 아닌데 정예화 아이템이 지급되는 현상 수정
```cs
// 아래처럼 수정 완료
if (stars < 3) return 0f;
else
{
	// 이전에 3성 클리어를 한 적이 없으면 지급
	if (resultInfo == null || resultInfo.stars < 3)
	{
		return 1f;
	}
	throw new InvalidOperationException("FirstClearPromotionItemRate의 예상치 못한 동작");
}
```

- [x] `NormalBuffVFX`의 `Lifetime`이 짧아서 버프가 켜진 중간에 스킬이 잠깐 끊기는 현상이 있음
	- **`Lifetime`을 늘려서 수정 완료**
## 250604

### 작업 중

#### 1-3 스테이지 작업
- 몹 배치
- ~~보스 작업~~

### 작업 완료
#### 범위 이펙트 : 최초 실행 시에 나타나지 않는 문제
- 이슈를 정리해봤음. AI한테 던져주기도 했다.
```
1. 스테이지 씬에서 스킬을 최초로 실행했을 때 범위를 나타내는 이펙트가 나타나지 않음
	1-1. 이펙트가 나타나지 않은 상황에서 스킬이 종료되고, 해당 스킬을 다시 실행했을 때에도 똑같이 범위를 나타내는 이펙트가 나타나지 않음.

2. 스테이지가 종료되고 메인메뉴 씬으로 전환됐다가 다시 스테이지 씬으로 돌아갔을 때에는 범위 이펙트가 최초 실행부터 잘 나타남.

3. 유니티 엔진 상에서 아예 실행을 멈췄다가(Stop) 다시 시작(Play) 버튼을 눌렀을 때, 스테이지 씬에서 범위 이펙트는 최초 실행부터 잘 동작함.

4. 유니티 엔진을 아예 껐다가 켰을 때에는 다시 1번과 1-1번처럼 동작함.
```
- [[셰이더 컴파일, 리소스 초기화 지연 문제]]
	- **셰이더 컴파일, 리소스 초기화** 지연 이슈. 
	- 에디터에서는 셰이더가 필요할 때마다 실시간으로 컴파일되기 때문에 컴파일이 완료되지 않은 상태에서 실행되면 이펙트가 나타나지 않을 수 있음.
	- 즉 에디터 단계에서만 필요한 기능
	- **스테이지 로딩할 때 가지고 있는 이펙트나 파티클 시스템들을 모두 잠깐 실행하는 기능을 구현해봤는데 안됨**
	- 이펙트가 나타나지 않는 게 문제이지 기능은 똑같이 동작하니까 이건 해결하지 않고 보류함.
	- 아니면 스테이지 진입 전에 메인메뉴 씬에서 스킬들을 1번씩 실행시키는 것도 괜찮을 듯

#### 기타 이슈 수정
- [x] `Artillery`의 폭발 이펙트가 여러 개 발생하는 현상 수정
	- `Projectile`로 피격이 발생할 때 범위 콜라이더가 발생 -> 범위 콜라이더에 충돌한 적들에 대해 피격 이펙트가 다시 발생하기 때문
	- ~~`UnitEntity.TakeDamage`에서 `bool playGetHitEffect = true`을 추가, 특별한 경우에만 피격 이펙트를 끄도록 수정함.~~
		- 이건 아예 피격 이펙트가 재생되지 않을 거임
		- `Enemy.TakeDamage`가 `UnitEntity.TakeDamage`를 상속받지만, `base.TakeDamage()`를 이용하는 구조는 아님. 이 구조를 바꿔봐야 할 듯?
			- **`base.TakeDamage()`를 이용하도록 수정.**

- [x] `Enemy.TakeDamage` 부분에서 추가로 구현된 부분은 `Operator`가 공격했을 때 통계 패널을 업데이트하는 부분이었다. `base.TakeDamage` 내에 `virtual` 메서드를 이용하도록 넣고 `virtual` 메서드는 `Enemy`에서 구현하는 식으로 수정.
	- `OnDamageTaken(float actualDamage)`로 사용 중인 게 있었다. 이걸 **`OnDamageTaken(UnitEntity attacker, float actualDamage)`로 바꾸고** `UnitEntity`에선 비우고 **자식 클래스들인 `Operator`와 `Enemy`에서 각각 방어 통계 업데이트와 공격 통계 업데이트하는 식으로 수정함.**

- [x] `OperatorInventoryPanel`
	- **오퍼레이터를 선택해서 들어갔다면 해당 오퍼레이터의 슬롯은 현재 스쿼드에서 사용 중인 스킬로 초기화**돼야 함. 지금은 기본 지정 스킬이 나타난다.
```cs
// 자동으로 선택되는 오퍼레이터의 스킬은 현재 스쿼드에서 사용 중인 스킬로 선택됨
if (operatorToAutoSelect != null && opToShow == operatorToAutoSelect)
{
	selectedSkillIndex = GameManagement.Instance!.UserSquadManager.GetCurrentSkillIndex(operatorToAutoSelect);
}
// 나머지는 오퍼레이터의 기본 선택 스킬을 사용
else
{
	selectedSkillIndex = opToShow.DefaultSelectedSkillIndex;
}
```
> 단일 선택 상황에서만 이런 이슈가 있었고 다중 선택 상황은 스쿼드의 스킬 정보를 이미 가져오는 식으로 구현되어 있었음.

- [x] `TestManager`에서의 테스트를 위한 초기화 환경 개선
	- 모든 오퍼레이터의 정예화, 레벨을 지정해서 초기화하도록 기능 수정


## 250603

### 작업 중
### 작업 완료

#### 기존 스테이지 리밸런싱
- 오퍼레이터가 이제 7명이기 때문에
	- [x] 한꺼번에 동시에 배치 가능한 인원을 **5명 -> 6명**으로 올림
	- [x] 보상 재설정 
		- 기존의 1, 2, 3 구성 중 어디를 바꿀까?
		- 정예화 가능 숫자 기준 `1, 3, 3`이랑 `2, 2, 3` 중에 `2, 2, 3`으로 결정
- `1-1`, `1-2` 스테이지 재설정
	- 동시에 6명을 배치할 수 있게 되면서 스테이지의 난이도를 조금 더 올릴 수 있을 것 같다.
	- ...라고는 해도 큰 수정은 안해도 될 것 같은게, 지상에 배치할 수 있는 오퍼레이터가 3기밖에 없어서 엄청 쉬워진다는 느낌은 아니다. 약간의 적만 추가했음.
- `Tanker` 방어력 550으로 상향
#### SixBulletSkill
- 스킬이 켜진 직후에 공격 속도와 모션을 0으로 만듦(초기화)
	- 즉 평타가 나간 직후에 스킬이 켜지면 바로 공격이 1번 더 나갈 수 있는 구성
#### TestManager 생성
- `PlayerDataManager`에서 테스트를 위한 초기화 부분을 `TestManager`로 별도로 떼어내서 구현함
	- AI한테 던져서 초안만 만들고 디테일은 잡아나가는 식으로 작업
	- `TestManager`는 `GameManagement`의 자식 오브젝트로 처리
#### 기타 이슈 수정
- [x] `StageSelectPanel` : `PromotionItem` 갯수가 1개로 고정되어 표기되는 문제 수정

## 250602 

### 작업 중

### 작업 완료

#### Artillery 2번째 스킬 남은 작업 처리
1. 스킬이 켜졌을 때 UI가 제대로 나타나지 않음
2. 스킬이 켜졌을 때 공격 속도 타이머가 다시 돌아가는 느낌이 있고, 공격 속도가 빨라지는지도 모르겠음.

- 우선 UI 작업부터 다시 진행함.
	- **이전에 AI가 던져준 코드는 폐기. 이해도 안 되고 제대로 동작하지도 않음.**
	- 기존의 `HealthBar`는 `Slider` 내에 들어가는 `Filled` Image만 갖고 컨트롤하는 방식이었음
	- `AmmoMode`를 켜게 되면 `Horizontal Layout Group`을 사용하고 
		- 기존의 `HealthFill`은 비활성화한다.
	- 근데 이거를 `DeployableBarUI` 단위에서 구현할지 `OperatorUI`에서 구현할지가 좀 애매한 지점이 있음. 프리팹 연결도 해제해둔 상태다.
		- `DeployableBarUI`에서 사용할 요소가 아니긴 한데, 스크립트는 `DeployableBarUI`에 구현되어 있다는 게 문제인가?

- 대충 아래처럼 작업할 예정
1. 일단 `OperatorUI`에서 구현을 할 거임
2. `DeployableBarUI`는 `OperatorUI` 외에도 자체적으로 쓰이는 상황이 있기 때문에, 이 경우는 단순하게 인터페이스처럼 동작함
3. `SPBar`를 `HealthBar`와 분리해서 스크립트를 구현함. `SPBar`에만 추가로 쓰이는 것들이 있다.
	- `HealthBar`의 기능 일부를 떼서 `SPBar`에 붙였다. 예를 들면 체력이 감소하는 효과를 나타내는 `DamageOverlay`나 `ShieldFill` 같은 부분은 `SPBar`에서는 필요하지 않으니 제외하고, `SPBar`에만 들어가는 `Ammo` 관련 `UI`들이 추가되었음.

> 발생 중인 이슈
 - [x] 스킬 켜지면 ~~SP Bar의 게이지 높이가 훨씬 커지며~~ `ammoMode`로의 전환도 제대로 안 됨
- [x] 스킬 임의로 껐을 때 기존 이미지들 제거해야 함
- [x] 탄환이 나갈 때 오른쪽에 빈 공간이 나타나는 게 아니라 계속 화면을 가득 채운 채로 현재 갯수만 감소함
	- 이거는 아예 `Child Force Expand`로 설정할 게 아니라 너비값을 설정하고 줄여나가는 식으로 해야 할 듯? 
	- 너무 어렵게 생각했다. 탄환 갯수를 계속 넣는 방식으로 업데이트하니까, **이미지를 비활성화하는 게 아니라 알파값을 0으로 만들어서 투명하게만 만들면 됨.**
	- 패딩 값은 `0.02` 정도로 잡으면 얼추 보이는 걸 알고 있다.

- `DeployableActionUI`  : 스킬 버튼에 중지 표시 구현
## 250601

### 블렌더 공부
- [[블렌더 5. UV 매핑]]


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
