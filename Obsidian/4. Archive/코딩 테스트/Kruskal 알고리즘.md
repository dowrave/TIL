- **그래프 내의 모든 정점을 가장 적은 비용으로 연결**하는 데 사용한다.
	- **모든 정점은 연결되어 있고, 사이클을 형성하지 않는다.**
- `그리디 알고리즘`의 일종이다 : 오름차순으로 간선들을 정렬하고, 사이클을 형성하지 않는 선에서 순서대로 정렬된 간선을 선택한다.


```python

# edge_arr = [(w, a, b), ...]

def kruksal():

	# 1. 오름차순 정렬
	edge_arr.sort() 
	
	# 2. 간선 선택
	
	# 사이클 점검을 위해 union find를 넣음
	parent = [i for i in range(v + 1)]
	
	def find(x):
		if parent[x] != x:
			parent[x] = find(parent[x])
		return parent[x]
	
	def union(a, b):
		if a == b:
			return
	
		a = find(a)
		b = find(b)
	
		if a > b: a, b = b, a # swap
	
		parent[b] = a
	
	for w, a, b in edge_arr:
	
		# 두 노드의 루트 노드가 같다 = 연결되어 있다
		# 이들을 새로 연결한다면 사이클이 발생함
		if find(a) == find(b):
			continue
	
		union(a, b)
```
- 아이디어 : 연결된 그래프들의 노드가 모두 같은 루트 노드를 가리키게 구성하면 

`costs = [[0,1,1],[0,2,2],[1,2,5],[1,3,1],[2,3,8]]`가 주어졌다고 가정해보자

> 1. `parent` 행렬을 만든다. 각 노드가 가리키는 부모 노드로, 초기화는 모두 자기자신.
> 2. 간선을 가중치의 오름차순으로 정렬한다
> 3. 가장 가중치가 작은 간선부터, 두 노드의 부모가 같은지 체크(`find`)하고 아니라면 병합한다`union`. "**병합한다 = 두 노드가 같은 루트 노드를 가리키게 한다**"
> 4. `3`을 반복하면 사이클을 갖지 않고 모든 노드가 연결된 그래프가 완성된다. 최소 비용을 구하고 싶다면 `union`이 실행될 때마다 가중치를 더하면 됨.