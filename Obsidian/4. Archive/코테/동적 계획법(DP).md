### DP의 목적
- 수많은 경우의 수를 따지는 과정에서 성능이 느려지는 것을 개선하고자 이용함 
- 메모리를 사용해서 중복 연산을 줄여 수행 속도를 개선한다.
	- `메모리를 사용한다` : 자료구조, 배열을 만든다
	- `중복 연산을 줄인다` : 연산 결과를 배열에 담는다

> 이름이 직관적이지 않다. 어떤 교수님은 "**기억하며 풀기**"라고 명명하셨는데 이 쪽이 훨씬 직관적이다.

### DP를 써야 하는 문제인가 파악하기
- DP는 **다양한 유형의 문제를 최적화**할 때 고려되는 알고리즘이다. 
- 기준 1. DFS나 BFS로 풀 순 있으나 **경우의 수가 너무 많은 문제**
	- DFS나 완전 탐색의 마지노선이 대충 500만 개($5*10^6$)
- 기준 2. **중복 연산이 너무 많은 문제**

## 예시 
[프로그래머스 정수삼각형 풀이 관련](https://www.youtube.com/watch?v=0bqfTzpWySY)

![[Pasted image 20230124165744.png]]
- 가장 위부터 아래로 내려가면서 최댓값을 찾는 문제임
- 아래로 내려가면서 값을 더하고, 지금까지 그 자리에 올 수 있는 연산 중 최댓값을 **저장**하는 방식임
	- 예를 들면 `(1,0)` 위치는 `0`과 `7+3` 중 `7+3`이 더 크니까 `10`을 저장하는 방식
- `(2,1)`위치를 보자
	- `7-3-1` 경로와 `7-8-1`경로가 있을 것이다
	- 처음엔 `7-3-1`의 값 `11`이 저장되나, 그 다음엔 `7-8-1` 경로의 값 `16`으로 갱신된다.
	- 이 **갱신**된다는 의미가 중요한데, 이후의 경로에서는 더 작은 값인 `7-3-1`을 더 이상 연산에 넣지 않게 된다는 것이다. 앞으로는 최댓값만을 사용하기 때문이다. 
	- DFS와 비교한다면 DFS는 이 이후로도 `7-3-1` 경로를 계속 연산에 집어넣고 값을 비교하게 될 것이다.

