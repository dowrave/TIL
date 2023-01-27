1. DP에 들어가는 값의 정의 : `DP[i]`는 i번째 인덱스에서 **끝나는** 최장 증가 부분 수열의 길이를 의미함
	- i번째 인덱스 값을 포함한다는 의미
2. `DP[i]`값은 1부터 시작함

- 따라서 dp의 각 인덱스`i`에 대해서 (반복문 1)
- 처음부터 `i-1`번째 인덱스까지 반복문을 돌린다(반복문 2)
	- 반복문 2의 구성 : `dp[i]`와 이전 인덱스의 값`dp[j]`를 비교했을 때, `dp[i]`가 `dp[j]`보다 크다면
		- `dp[i]`는 `dp[j] + 1`과 `dp[i]` 값 중 큰 값을 선택한다

```python
lst = [] # 배열
dp = [1] * n
for i in range(n):
	for j in range(i):
		if lst[i] > lst[j]:
			dp[i] = max(dp[i], dp[j] + 1)
```

- 위 방법은 $O(N^2)$ 이기 때문에, 이분 탐색을 이용하는 방법이 있다.

## 이분 탐색 이용

- 
```python
def bi_search(lst, target):
	"""
	lst는 dp array
	target은 원래 리스트의 값 
	"""
	left = 0
	right = len(lst) - 1
	
    while left < right:
        mid = (left + right) // 2
        if dp[mid] < target:
            left = mid + 1
        else:
            right = mid
    return right

dp = []
dp.append(lst[0]) # dp에 lst[0]을 넣는다
dp_idx = 0 # dp에 값을 넣는 기준 인덱스 : 마지막 값

for i in range(1, n):
    if dp[dp_idx] < lst[i]:
        dp.append(lst[i])
        dp_idx += 1
    else:
        new_idx = bi_search(dp, lst[i]) # lst[i]가 들어갈 곳을 찾는다
        dp[new_idx] = lst[i]

# dp의 길이를 물어보니까
print(dp_idx + 1)

```

![[LIS-1 1.jpg]]