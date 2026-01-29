- 현재 어떤 작업 중인지 기록 중
# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

---
# 작업 일지

## 블로그 고치고 싶은거

## 짭명방 
- 지난 내역 : [짭명방 프로젝트 일지 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- [[기타 참고 사항]]

>[!note]
>
>  **"흐름"으로 설명하는 것 외에도 "원리"도 곁들이면 좋다**
> - 지금은 "어떤 이슈가 발생해서 어떻게 해결했더라"로 정리했다면
> - 추가로 **이슈가 '왜' 발생했는지**까지 정리해두면 나중에 볼 때 다시 헷갈리지 않을 수 있음
> 
> 모든 코드를 일일이 기억하는 건 불가능함! 나중에 봤을 때 '이러면 되겠다'는 감이 바로 잡히도록 메모해보자.

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움
> - `Enemy`가 사라질 때 풀 태그의 키가 없다는 오류 
> - `Barricade` 배치 시 가끔 배치되지 않음
> - `Operator` 체력이 다했을 때 사망처리가 되지 않았는데도 적이 지나가는 현상

## 현재 계획 및 이슈

>[!plan]
>1. 스테이지 밸런싱, 버그 수정
>2. 소리 추가
>3. 프로젝트 정리글 작성

## 최근 작업 내용
- 블로그 : [[블로그_260127 EC2에서 Lightsail로 이사 가기]]
- 짭명방 : [[짭명방_260128]]
## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

## 260129

## 이슈 해결하기
>[!done]
>- `PathIndicator` 초기화 시 `PathData`가 정상적으로 들어가지 않는 현상(Node가 0개로 나타남)
>- 일부 `PathData`가 날아가 있음. 원상복구.
>	- `1-2 Direct`, `1-3 Direct` 데이터들이 날아가 있었다. 
>- `PathData`의 `public Nodes`을 `[SerializeField] private nodes` 및 프로퍼티로 변경, 그 과정에서 에디터의 로직도 변경
## PathIndicator 초기화 시 PathData가 정상적으로 들어가지 않음

- 한참 헤맸다. SO 정보가 갑자기 날아가 있기도 하고, 그걸 고쳤다고 잘 실행되는 것도 아니고, 디버깅도 여기저기 막 찍어봄
- `SetNewPath`랑 `HandlePathUpdated`가 있다. 리스트를 초기화하고 집어넣는 과정에서 정보가 누락되고 있는 것으로 보임. 교통정리좀 해야 할 듯.

- `HandlePathUpdated`의 내용물을 밖으로 빼뒀다. 로직이 겹쳐서 헷갈리고, 다른 곳에서 쓰이지 않기 때문임
```cs
if (newPathNodes != null && newPathNodes.Count > 0)
{
	// 현재 경로 노드 수정
	_currentPathNodes.Clear();
	_currentPathNodes.AddRange(newPathNodes);

	// 현재 경로(실제 Vector3 위치) 수정
	float floatY = _owner is Enemy enemy ? enemy.BaseData.DefaultYPosition : 0.3f; // Enemy마다 y 위치가 다름
	_currentPathPositions.Clear();
	_currentPathPositions.AddRange(_currentPathNodes.Select(node => MapManager.Instance!.ConvertToWorldPosition(node.gridPosition) + Vector3.up * floatY).ToList());

	// foreach (var pathPosition in _currentPathPositions)
	// {
	//     Logger.LogFieldStatus(pathPosition);
	// }

	// 인덱스 및 목적지 수정
	_currentPathIndex = _currentPathNodes.Count > 1 ? 1 : 0; // [테스트] 뒤로 가는 현상을 방지하기 위해 1로 놔 봄
	_currentDestination = _currentPathPositions[CurrentPathIndex];

	// HandlePathUpdated(newPathNodes, _currentPathPositions);
	return;
	// OnPathUpdated?.Invoke(_currentPathNodes, _currentPathPositions);
}
```

- 지금까지의 방식에서는 `Clear()` 및 `AddRange()`로 하나의 리스트를 계속 사용하는 방식을 취했는데, 어차피 `_currentPathNodes.Select()` 등에서 새로운 리스트를 만드는 상황

### AddRange()의 동작 원리
- `A.AddRange(B)`라고 하면, `B`의 요소들을 복사해서 하나씩 `A`에 넣는 개념임
- 주의할 점 - **얕은 복사** : `B`의 각 요소가 참조형이라면, `Bx`에서 `A`로 복사된 요소 `Ax`의 값을 수정하면 `Bx`의 값도 수정된다는 것
- 물론 이 경우는 그걸 고려하지 않고 쓸 수 있음 - `B`를 쓰지 않기 때문임

- **`new List()`와의 차이점은, 전역 변수를 일일이 새로 할당하지 않는다**는 것에 있음. 여기에 들어가는 원본 리스트는 할당될 수밖에 없음

### 수정
- `_currentPathNodes`의 경우는 그대로 사용할 수 있음. `_currentPathPositions`만 내부 요소를 하나씩 꺼내서 연산해서 넣는 방식으로 작업한다.
- 이외에도 `Claude`가 추천해준 옵션으로, 전역 리스트의 `Capacity`를 미리 할당하는 걸 권장했다. 리스트의 길이가 변하는 상황에서 `Capacity`가 2배씩 늘어나는데, 이를 사전에 할당해주면 중간중간 업데이트하는 로직을 생략할 수 있기 때문.
```cs
// _currentPathPositions 부분만

_currentPathPositions.Clear();

// Capacity 미리 확보 : 요소가 Capacity를 초과할 때 List는 배열을 2배로 재할당함
// 미리 맞춰놓으면 재할당 없이 한꺼번에 추가할 수 있음
if (_currentPathPositions.Capacity < _currentPathNodes.Count)
{
	_currentPathPositions.Capacity = _currentPathNodes.Count; 
}

// 반복문으로 요소를 _currentPathPositions으로 옮김
for (int i = 0; i < _currentPathNodes.Count; i++)
{
	Vector3 worldPos = MapManager.Instance!.ConvertToWorldPosition(_currentPathNodes[i].gridPosition);
	_currentPathPositions.Add(worldPos + Vector3.up * floatY);
}
```

### 근데 이래도 안되네?
- 이슈 내용은 똑같다. 
```cs
_currentDestination = _currentPathPositions[CurrentPathIndex];
```
이 부분에서 인덱스가 초과한다는 내용.

- 이런 의심이 들었다 : 얘가 생성되면서 데이터 다 날리는 거 아님?
- 일단 `PathData`의 `nodes`가 `public`으로 되어 있었다. 
```cs
[Header("Path Nodes")]
[SerializeField] private List<PathNode> nodes = new List<PathNode>(); 
// public List<PathNode> Nodes = new List<PathNode>();
public IReadOnlyList<PathNode> Nodes => nodes;
```

근데 이렇게 바꾸면 `PathDataEditor`에서 `public`으로 직접 조작할 수 있었던 `Nodes`을 더 이상 스크립트로 접근해서 고칠 수 없게 된다.

대충 이런 방법을 사용함.
- `SerializedProperty`로 변수를 인스펙터에서 설정하는 기능이 있음
```cs
public class PathDataEditor : Editor 
{
	private SerializedProperty nodesProperty; 

    private void OnEnable()
    {
		nodesProperty = serializedObject.FindProperty("nodes"); // PathData의 변수 이름 nodes
    }
    
    // 기타 요소 생략
    
	private void AddNode(Tile tile)
    {
        if (pathData == null) return; // Add null check

        serializedObject.Update();

        // 배열 끝에 새 요소 추가
        int newIndex = nodesProperty.arraySize;
        nodesProperty.InsertArrayElementAtIndex(newIndex);

        // 새로 추가된 요소 가져오기
        SerializedProperty newElement = nodesProperty.GetArrayElementAtIndex(newIndex);

        // PathNodes의 각 필드 설정
        newElement.FindPropertyRelative("tileName").stringValue = tile.name;
        newElement.FindPropertyRelative("gridPosition").vector2IntValue = tile.GridPosition;
        newElement.FindPropertyRelative("waitTime").floatValue = 0f;

        serializedObject.ApplyModifiedProperties();

        // serializedObject가 자동으로 처리해줌
        // Undo.RecordObject(pathData, "Add Path Node");
        
        // pathData.Nodes.Add(new PathNode
        // {
        //     tileName = tile.name,
        //     gridPosition = tile.GridPosition,
        //     waitTime = 0f
        // }
        // );

        // EditorUtility.SetDirty(pathData);
    }
    
	private void RemoveNearestNode(Vector2Int position)
    {
        InstanceValidator.ValidateInstance(pathData);

        if (pathData!.Nodes.Count == 0) return;

        // 가장 가까운 노드의 인덱스 찾기
        int nearestIndex = pathData.Nodes
            .Select((node, index) => new { Node = node, Index = index })
            .OrderBy(x => Vector2Int.Distance(x.Node.gridPosition, position))
            .First().Index;

        serializedObject.Update();

        // 해당 인덱스의 요소 제거
        nodesProperty.DeleteArrayElementAtIndex(nearestIndex);
        serializedObject.ApplyModifiedProperties();

        // Undo.RecordObject(pathData, "Remove Path Node");
        // pathData.Nodes.RemoveAt(nearestIndex);
        // EditorUtility.SetDirty(pathData);
    }
}
```

이러고 새로운 Nodes을 에디터에서 다시 클릭해서 설정했음.

#### 그리고..
- 처음에 `PathController`의 생성자에서 경로 데이터까지 받아 초기화한 다음 `Initialize()`에서 `SetNewPath()`을 하도록 설정했는데, 이러면 최초에 들어온 `pathNode`가 초기화된다. `SetNewPath`에는 항상 초기화 로직이 실행되도록 구현되어 있기 때문.

- 따라서 언제 실행되어도 상관 없는 요소들을 생성자로 뺐고
- 경로 데이터는 들어오는 시점에 초기화되도록 설정함 : 괜히 `_currentPathNodes`을 또 건드리지 않게끔.

- 수정 버전
```cs    
// 생성자에서 초기화
public PathController(MonoBehaviour owner)
{
	_owner = owner;

	// 이벤트 구독
	Barricade.OnBarricadeDeployed += OnBarricadePlaced;
	Barricade.OnBarricadeRemoved += OnBarricadeRemovedWithDelay;
}

// OnPathUpdate를 구독하는 부모 클래스의 메서드가 있기 때문에
// 생성자와 실제 실행 메서드는 별도로 구분하는 게 좋다.
public void Initialize(IReadOnlyList<PathNode> initialPathNodes)
{
	SetNewPath(initialPathNodes);

	_finalDestination = _currentPathPositions[_currentPathPositions.Count - 1]; // 최초 경로의 마지막 값

	// 최초 경로가 막혔을 때
	if (IsPathBlocked(0))
	{
		UpdatePath();
	}

	IsInitialized = true;
}
```

- **잘 된다.** 현재 `Enemy` 코드 약 400줄. `PathIndicator` 코드 약 96줄.

# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
