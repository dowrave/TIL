
### 1. 입양 시각 구하기
- 09:00부터 19:59까지, 각 시간대별로 입양이 몇 건이나 발생했는지 조회하는 SQL문을 작성해주세요. 이때 결과는 시간대 순으로 정렬해야 합니다.
```SQL
SELECT SUB.HOUR, COUNT(*) AS COUNT
  FROM (
    SELECT *, HOUR(DATETIME) AS HOUR
      FROM ANIMAL_OUTS
     WHERE HOUR(DATETIME) BETWEEN 9 AND 19
) SUB
GROUP BY HOUR
ORDER BY 1 ASC
```
>`DATETIME` 자료형에서 원하는 값(연, 월, 일, 시, 분 ,초)을 뽑아내려면
```
	-   YEAR : 연도 추출
	-   MONTH : 월 추출
	-   DAY : 일 추출 (DAYOFMONTH와 같은 함수)
	-   HOUR : 시 추출
	-   MINUTE : 분 추출
	-   SECOND : 초 추출
```


### 2. 입양 시각 구하기(2)
- 위랑 똑같은데 0시부터 23시임 : SQL에 들어있지 않은 값을 어떻게 채워넣어 줄까?
```SQL
-- 1) 0부터 23이 들어가 있는 테이블을 하나 만든다
WITH RECURSIVE HOURS AS (
    SELECT 0 AS hour
    UNION ALL 
    SELECT hour + 1
      FROM HOURS
     WHERE hour < 23
    )

-- 2)
( -- 1. 본 테이블에 없는 hour에는 count column에 0을 넣는다.
SELECT hour, (0) AS cnt 
  FROM HOURS
 WHERE hour NOT IN ( -- 요 where문이 본 테이블에 없는 시간들임
                    SELECT HOUR(DATETIME) 
                      FROM ANIMAL_OUTS
                     GROUP BY HOUR(DATETIME) 
                    )
UNION ALL -- 그것들을 UNION ALL함
-- 2. 본 테이블에 있는 hour에는 그냥 해주면 됨
SELECT HOUR(DATETIME) AS hour, COUNT(*) AS cnt
  FROM ANIMAL_OUTS
 GROUP BY HOUR(DATETIME)
)
-- 마지막에 정렬해줌
 ORDER BY hour
    
```
>- 이런 구현이 가능하겠다 : 0부터 23까지 있는 HOUR 테이블
	- 이거 어케 함? : [[임시 테이블 생성하기]]

> 방법
> 1. 0부터 23까지 돌아가는 테이블을 하나 만듦
> 2. 주어진 테이블에 없는 시간에 대해, 카운트 값으로 0을 넣는다.
> 3. 주어진 테이블에 있는 시간은 그냥 구해서 넣으면 됨
> 4. 2와 3을 `UNION ALL` 해준다. 

- 그러면 **`UNION ALL`은 row 단위로 나눠서 넣는가봉가**

- 이거는 복습 필요하겠다 : **`WITH RECURSIVE`를 만드는 방법이라든가, `UNION ALL`의 활용이라든가 모두 중요해보임**