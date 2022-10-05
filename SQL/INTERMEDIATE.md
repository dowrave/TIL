# Intermediate

## Aggregate Function

### COUNT
1. 모든 Column 수 세기
```SQL
SELECT COUNT(*) 
  FROM (table1)
```
- `COUNT(1)`도 `COUNT(*)`과 동일하다.

2. (Null 값 제외한) 특정 Column의 수 세기 : 
```SQL
SELECT COUNT(column1)
  FROM (table1)
```

- `COUNT()`는 단순히 ROW의 수만을 셈 : 고유값을 세는 함수는 뒤에서 다룸.
- `AS`를 이용해 새로운 COLUMN으로 조회할 수도 있음.
  - 주의) 기존 column들과 **동시에 조회할 수 없음**
  - 대신 `Aggregation Function`끼리는 동시 조회 가능함

### SUM
- 어떤 **열**의 값들을 다 더함
- 참고) 행 방향으로 연산을 하고 싶다면, `BASIC`에서 다룬 단순한 연산을 이용하면 됨 ex) `colA + colB AS colC`

### MIN & MAX
- 숫자가 아닌 데이터 타입에도 적용 가능함
  - `MIN` : 가장 작은 숫자, 이른 날짜, a에 가까운 알파벳
  - `MAX`는 그 반대겠죠?

### AVG
- 몇 가지 제한 사항이 있음
1. **수치형에만 사용** 가능
2. **`null`값을 완전히 무시**함
- 따라서 사전에 결측치를 처리하긴 해야 함
```SQL
SELECT AVG(col1)
  FROM table1
 WHERE col1 IS NOT NULL -- 있든 없든 같은 결과를 낸다
```

## GROUP BY
- `Aggregation Function`은 모든 테이블에 대해 적용됨
- 근데 테이블의 일부에만 연산을 가할 상황이 있음
- 이 때 쓰는게 `GROUP BY`임
- `GROUP BY`는 데이터를 그룹으로 나누고, 서로 독립된 상황을 만듦
```SQL
SELECT year,
       month,
       COUNT(*) AS count
  FROM tutorial.aapl_historical_stock_price
 GROUP BY year, month
```
- 어떤 연도가 몇 개의 ROW를 갖는지 나눠서 보여줌
- `GROUP BY`를 쓰면 원본 COLUMN과 `AGG FUNC`을 함께 나타낼 수 있다.
- `GROUP BY` 뒤에 오는 COLUMN도 여러 개 지정할 수 있음
- 보면 알겠지만 `GROUP BY`에 쓰인 column은 `SELECT`문에도 들어감
-------------------
- 예제 : 연, 월을 오름차순으로 정렬
```SQL
SELECT year,
       month,
       SUM(volume) AS SUM_VOLUME
  FROM tutorial.aapl_historical_stock_price
 GROUP BY year, month
 ORDER BY year, month
```
--------------------
- Column의 순서를 안다면 1, 2 식으로 집어넣을 수 있으나 모든 SQL에 적용 가능한 것은 아니다.
- `GROUP BY`에 들어가는 열의 순서는 상관 없다. 그러나 `ORDER BY`에 들어가는 열의 순서는 상관 있다.
  - 위 예제는 `2000년 1월, 2000년 2월, ...` 순으로 정렬된다.
  - 그러나 `ORDER BY`가 `month, year`로 바뀐다면 `2000년 1월, 2001년 1월, 2002년 1월, ...`순으로 바뀐다.
- LIMIT과 함께 쓸 경우
  - LIMIT은 표시되는 데이터의 수에 제한을 걸기 때문에, 만약 고유값의 수가 100개가 넘는다면 정보의 일부분은 표시되지 않을 것임
  - 대충 그런 상황 조심하라는 얘기임

----------
예제2) 1년의 하루 평균 주식 가격 변화량
```SQL
SELECT year,
       AVG(close - open) AS avg_daily_change
  FROM tutorial.aapl_historical_stock_price
 GROUP BY year
 ORDER BY year
```
예제3) 매 해 매 달 주식의 최저가와 최고가
```SQL
SELECT year, 
       month,
       MAX(high) as monthly_high,
       MIN(low) as monthly_low
  FROM tutorial.aapl_historical_stock_price
 GROUP BY year, month
 ORDER BY year, month 
 ```
 ---------------

### 데이터 분석 언어로서의 SQL과 Python의 차이점
[원문](https://mode.com/blog/group-by-sql-python/?utm_medium=referral&utm_source=mode-site&utm_campaign=sql-tutorial)
- 여기선 결론 부분만 요약해둠
1. **복잡한 Grouping이 필요하지 않다면 SQL**을 이용, 한 언어로 계속 작업할 수 있다. 
2. 더 **복잡한 Grouping이 필요하다면 Python**으로 옮겨 작업하는 게 좋다.

## HAVING
- `Aggregate`된 Column에는 `WHERE`문이 먹히지 않는다.
- 이 때는 `HAVING`을 사용한다.
```SQL
SELECT year,
       month,
       MAX(high) AS month_high
  FROM tutorial.aapl_historical_stock_price
 GROUP BY year, month
HAVING MAX(high) > 400
 ORDER BY year, month
```
- 참고) `Subquery`에도 사용된다. 

### 쿼리문 순서
1. `SELECT`
2. `FROM`
3. `WHERE`
4. `GROUP BY`
5. `HAVING`
6. `ORDER BY`

## CASE
- SQL에서의 `if/then`문. 
- 항상 `CASE WHEN ... THEN ... (ELSE ...) END ...`식으로 사용한다.
```SQL
SELECT player_name,
       year,
       CASE WHEN year = 'SR' THEN 'yes'
            ELSE NULL END AS is_a_senior
            -- is_a_senior이라는 Column을 만든다
            -- 그 값은 year 값이 'SR'일 때 'yes',
            -- 아니라면 결측치로 남겨둔다
  FROM benn.college_football_players 
```
1. `CASE`문은 각 row에 대해 `year = 'SR'`을 체크한다
2. 조건문이 참이라면 `is_a_senior` 열에 `yes`값이 들어간다
    - 거짓이라면 null
3. 위 과정과 동시에 SQL은 `player_name`, `year` 열의 값들을 보여준다.

- 여러 조건문 넣기 
```SQL
SELECT player_name,
       weight,
       CASE WHEN weight > 250 THEN 'over 250'
            WHEN weight > 200 AND weight <= 250 THEN '201-250'
            WHEN weight > 175 AND weight <= 200 THEN '176-200'
            ELSE '175 or under' END AS weight_group
  FROM benn.college_football_players
```
- 조건의 중복을 피하게끔 코드를 짜 넣는게 좋다.
    - 즉 2번째 WHEN 문은 사실 `weight > 200`만 남겨도 동일하게 작동한다. `weight > 250`문에서 조건들이 걸러졌기 때문이다.
    - 그러나 **보는 사람이 명료하게 알 수 있게 쿼리를 짜 넣는 것이 중요**하다.
- `WHEN` 문 내부에는 `AND/OR`문을 넣을 수 있음.
---------------
- 예제 진행 중 결측치 데이터 처리
```SQL
SELECT max(height), min(height), avg(height)
  FROM benn.college_football_players 
 WHERE height != 0 
  -- 이거 있을 때 83, 57, 73.27
  -- 이거 없을 때 83, 0, 72.67
```
- `SELECT`문은 마지막에 실행되므로, `WHERE`문에서 결측치 값은 제거하고 데이터를 처리했다. 
--------------------
### CASE with AGG FUNC.
어떤 조건을 만족하는 열만을 세기
```SQL
SELECT CASE WHEN year = 'FR' THEN 'FR'
            ELSE 'Not FR' END AS year_group,
            COUNT(1) AS count
  FROM benn.college_football_players
 GROUP BY year_group
```
- `GROUP BY`문을 빼면 작동하지 않음
  - 일반 COLUMN과 `AGG FUNC`이 같이 들어가기 때문임
  - 이 때는 `SELECT`에 있는 조건문을 그대로 같이 써주면 된다.
- 위 쿼리는 `FR`과 `Not FR`에 대한 `COUNT` 결과를 같이 보여준다.
- `CASE ~ END`까지를 `GROUP BY` 문에 복사해도 똑같이 작동한다. 근데 굳이 그렇게 길게 쿼리를 작성할 필요가 없겠죠?
- `SELECT`문의 `CASE`가 맨 앞에 온다면 `GROUP BY 1`로도 가능하다.
  - 근데 명시적으로 작성해놓는 게 더 좋을 것 같다. 
```SQL
SELECT COUNT(1) AS fr_count
  FROM benn.college_football_players
 WHERE year = 'FR'
```
- 이런 식으로 `FR`과 `Not FR`을 각각 돌려 count 값을 얻을 수 있지만, 한 번에 1개의 값 밖에 얻을 수 없다. 

CASE 문 결과를 맨 앞에 넣기
```sql
SELECT CASE WHEN year = 'FR' THEN 'FR'
            WHEN year = 'SO' THEN 'SO'
            WHEN year = 'JR' THEN 'JR'
            WHEN year = 'SR' THEN 'SR'
            ELSE 'No Year Data' END AS year_group,
            *
  FROM benn.college_football_players
```
- *을 빼고 AGG FUNC을 넣고, GROUP BY를 넣을 수 있을 것이다.  

AGG 함수 내에 CASE 넣기(`피벗 테이블`)
```SQL
SELECT CASE WHEN year = 'FR' THEN 'FR'
            WHEN year = 'SO' THEN 'SO'
            WHEN year = 'JR' THEN 'JR'
            WHEN year = 'SR' THEN 'SR'
            ELSE 'No Year Data' END AS year_group,
            COUNT(1) AS count
  FROM benn.college_football_players
 GROUP BY 1
 ```
 - 결과 `year_group` 열에 `FR, SO, JR, SR, No Year Data` 값이 있을 거고 오른쪽의 `count` 열에 각 항목의 갯수가 있을 것이다.

```SQL
SELECT COUNT(CASE WHEN year = 'FR' THEN 1 ELSE NULL END) AS fr_count,
       COUNT(CASE WHEN year = 'SO' THEN 1 ELSE NULL END) AS so_count,
       COUNT(CASE WHEN year = 'JR' THEN 1 ELSE NULL END) AS jr_count,
       COUNT(CASE WHEN year = 'SR' THEN 1 ELSE NULL END) AS sr_count
  FROM benn.college_football_players
```
- **AGG 함수 내에 CASE문을 넣으면**, 4개의 열(`fr_count, so_count, jr_count, sr_count`)이 생기고 1개의 행에 각 count값이 들어가게 된다. 

----------------------
예제1 : 각 주(state)에 대해 전체 선수의 수와 연차 별 선수의 수를 나타내라.
(정렬은 전체 선수 수가 많은 순서대로 하라)
```SQL
SELECT state,
       COUNT(1) as total_count,
       COUNT(CASE WHEN year = 'FR' THEN 1 END) AS fr_count,
       COUNT(CASE WHEN year = 'SO' THEN 1 END) AS so_count,
       COUNT(CASE WHEN year = 'JR' THEN 1 END) AS jr_count,
       COUNT(CASE WHEN year = 'SR' THEN 1 END) AS sr_count
  FROM benn.college_football_players
 GROUP BY state
 ORDER BY total_count DESC 
```
예제2 : 이름이 A부터 M까지 시작하는 학교의 선수 수와, N부터 Z까지로 시작하는 학교의 선수 수를 나타내라
```SQL
SELECT CASE WHEN school_name < 'N' THEN 'A-M'
            WHEN school_name >= 'N' THEN 'N-Z'
            END AS name_group,
            count(1) AS number_of_players
  FROM benn.college_football_players
 GROUP BY name_group
```
- 문자 검색 조건이 깔끔하니 참고해두자.
----------------------------
## DISTINCT
- SELECT문에서만 쓸 수 있고, 각 COLUMN의 고유값을 찾는다.
- 당연히 여러 COLUMN에 대해 쓸 수 있음.
- **새로운 데이터셋을 탐색할 때 유용**하다 : 자주 쓰임

Agg Func과의 조합
```sql
SELECT COUNT(DISTINCT month) AS unique_months
  FROM tutorial.aapl_historical_stock_price
```
- 고유값 갯수 셀 때 `COUNT(DISTINCT ...)`를 이용한다.
- 참고 ) AGGREGATION에서 DISTINCT를 쓰는 건 쿼리 속도를 다소(`a bit`) 저하시킬 수 있다. 
---------------------------
## JOIN부터 ㄱㄱ