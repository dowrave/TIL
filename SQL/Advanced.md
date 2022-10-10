# Advanced
[여기부터 시작](https://mode.com/sql-tutorial/intro-to-advanced-sql/)
## Data Types
- SQL DB는 다른 포맷의 데이터를 다른 정밀도로 보관할 수 있다. 
  - `INTEGER` 타입은 정수를 보관할 수 있음
  - `DOUBLE PRECISION (=FLOAT64)`는 물리학자가 아닌 이상 필요한 것보다 높은 정밀도로 데이터를 저장함
- MODE에서 지원하는 데이터 타입은 이러하다
- 우항이 공식 SQL 데이터 타입
1. `STRING` = `VARCHAR(1024)`
2. `DATE/TIME` = `TIMESTAMP`
    - `YYYY-MM-DD hh:mm:ss`
3. `NUMBER` = '`DOUBLE PRECISION`
4. `BOOLEAN` = `BOOLEAN`

### 쿼리로 열의 데이터타입 바꾸기
- 보통 `date`나 `number`가 `string`으로 저장되어 있을 때 변환을 이용한다.
```SQL
CAST(column_name AS integer)
column_name :: integer
```
- 위 2개 줄은 같은 역할을 수행한다. 

## DATE FORMAT
- 미국은 `MM-DD-YYYY` 포맷을 쓴다.
- 근데 나머지에서는 `YYYY-MM-DD`를 쓰죠?
- 한편, `string`으로 저장된 경우라면 순서대로 정렬할 수가 없음
- 대부분의 DB 포맷은 `YYYY-MM-DD`를 따른다.
  - 왜냐면 `DATE` 포맷이든 `STRING` 포맷이든 정렬이 잘 되기 때문임
  
- 날짜에 연산을 가하면 결과는 `interval` 데이터 타입이 된다. 기간을 나타내는 `integer`임.  

예시 : `datetime`과 `timestamp`
```SQL
SELECT companies.permalink,
       companies.founded_at_clean,
       acquisitions.acquired_at_cleaned,
       acquisitions.acquired_at_cleaned -
         companies.founded_at_clean::timestamp AS time_to_acquisition
  FROM tutorial.crunchbase_companies_clean_date companies
  JOIN tutorial.crunchbase_acquisitions_clean_date acquisitions
    ON acquisitions.company_permalink = companies.permalink
 WHERE founded_at_clean IS NOT NULL
 ```
 - `acquisitions.acquired_at_cleaned`는 `datetime`,`companies.founded_at_clean` 열은 `string`이라서 `timestamp`로 바꿔줌
   - `datetime`과 `timestamp` 간 연산이 가능하다는 것이 있고,
   - 그 결과는 `timestamp`이며 결과는 `day`로 나왔다.
  
예시 2 : `INTERVAL`
```SQL
SELECT companies.permalink,
       companies.founded_at_clean,
       companies.founded_at_clean::timestamp +
         INTERVAL '1 week' AS plus_one_week
  FROM tutorial.crunchbase_companies_clean_date companies
 WHERE founded_at_clean IS NOT NULL
 ```
- `plus_one_week` 열은 `companies_founded_at_clean`에 1주일을 더한 날짜이다.
  - 비슷하게 `x seconds`, `y month`도 사용 가능.
- `INTERVAL` 열과 `DATE`열을 더했더니 `DATE`열이 나왔다는 것도 참고.

- 그 외 `NOW()`라는 함수도 있다. 쿼리를 작동시킨 현재 시간이 들어감.
- 예제 풀이) `INTERVAL 'X years'`은 일종의 기간에 대한 변수로 볼 수 있다.
  - ex) `date2 - date1 :: timestamp <= INTERVAL 'X years'` 같은 형태로 쓸 수 있다는 것.


## Data Wrangling with SQL
- `Wrangle`이란?
  ``` 
  반자동 도구를 이용, "원시" 데이터를 보다 편리하게 사용할 수 있도록 다른 형식으로 수동을 변환하거나 매핑하는 프로세스
  ```
  - 즉, 열의 값들을 특정 방식으로 바꾸거나 여러 열을 합칠 수 있다.  
  - 사람에 의해 수집된 데이터는 실수가 많고,
  - 웹에서 수집된 데이터는 웹에 표시되는 데 최적화되지 수집하는데 최적화되지 않는다.

## SQL String Function

### String Cleansing
#### `LEFT`
- `LEFT(대상 문자열, 출력할 마지막 인덱스(포함))`
- SQL에서 시작 인덱스는 1이다.
- 첫 인덱스 ~ 마지막 인덱스'까지'를 추출

#### `RIGHT`
- 용법은 `LEFT`와 동일.
- 가장 마지막 글자의 인덱스를 1로 본다.
- 마지막 인덱스 ~ 첫 인덱스를 추출

#### `LENGTH`
- `LENGTH(대상 문자열)` 
- 글자의 길이를 추출한다.
- SQL에서 **길이 = 마지막 인덱스**

#### `TRIM`
- `TRIM (arg1 'charcter' FROM column)`
  - `arg1`에는 어느 쪽의 문자열을 제거할까를 정한다
    - `leading` : 시작부분만
    - `trailing` : 끝부분만
    - `both` : 양쪽 모두
  - `character`에는 어떤 문자열을 제거할까를 정한다. 
    - 참고) 중복해서 들어오는 같은 문자열은 모두 제거된다. 
    - 여러 실험을 해봤는데, 정확히 어떤 기준으로 제거되는지 잘 모르겠다. 
  ```SQL
  SELECT TRIM (leading 'h' FROM 'helloh') -- elloh
  SELECT TRIM (trailing 'h' FROM 'helloh')
  -- hello
  SELECT TRIM (both 'h' FROM 'helloh')
  -- ello
  SELECT TRIM (both 'h' FROM 'hhhellohh')
  -- ello
  ```
일단 양 끝에 있는 어떤 값을 제거하기 위해 이용한다고 생각하자. 
```SQL
SELECT location,
       TRIM(both '()' FROM location)
  FROM tutorial.sf_crime_incidents_2014_01
```

#### POSITION, STRPOS
- `POSITION('character' IN column)` 
  - `character`가 처음 등장하는 인덱스 반환
- `STRPOS(column, 'character')`도 동일한 결과.
- 어떤 걸 쓰든 대문자나 소문자로 통일하고 사용하자. 영향을 받는다.

#### SUBSTR
- `SUBSTR('column', 시작 인덱스, 문자 수)`
  - 바꾸는 게 아니라 **뽑아내는 거**임
- `SUBSTR('01/31/2014', 4, 2)`의 실행 결과는 `31`이다.

-------------
예제 : `location`에서 위도, 경도 따로 뽑기
```sql
SELECT location,
      LEFT(TRIM(both '()' FROM location), STRPOS(TRIM(both '()' FROM location), ',') - 1) AS hand_lat,
      -- RIGHT(TRIM(both '()' FROM location), STRPOS(TRIM(both '()' FROM location), ',')) AS hand_lon,
      SUBSTR(location, STRPOS(location, ',') + 1, LENGTH(location) - STRPOS(location, ',') - 1) AS hand_lon,
       lat, 
       lon
  FROM tutorial.sf_crime_incidents_2014_01
```
- 뭔 함수든 `함수(COLUMN, 'CHARACTER')` 순으로 들어가는 걸 기억하자. 겁나 헷갈렸넹
- `RIGHT`가 생각보다 쓰기 까다롭다. 얘는 인덱스를 거꾸로 세는데 다른 함수들은 인덱스를 똑바로 세기 때문에 그렇다.
- 그래서 대신 `SUBSTR`을 쓰는데, 양쪽 끝의 괄호를 `TRIM`하는 게 편할 것이다.
  - 근데 SELECT에서 그런 식으로 COLUMN을 만들고 그 밑에서 인용하려고 했는데 그 방법은 먹히질 않았다. 즉 모든 COLUMN에 TRIM 문을 넣어야 했다
- 그래서 대신 `SUBSTR`문의 시작점과 끝점 인덱스를 괄호를 신경 써가면서 넣었다. 
--------------------
#### CONCAT
- 여러 열을 이어붙이는 기능.  

예시)
```SQL
SELECT incidnt_num,
       day_of_week,
       LEFT(date, 10) AS cleaned_date,
       CONCAT(day_of_week, ', ', LEFT(date, 10)) AS day_and_date
  FROM tutorial.sf_crime_incidnets_2014_01
```

- 참고) SELECT에서 `cleaned_date`라는 열을 지정했을 때 아래의 `CONCAT` 문에서는 위에서 지정한 열을 이용할 수 없음.
  - 아마 원래 테이블에 있는 `column`만 이용할 수 있는 것 같다.

- 위의 CONCAT 문은 `day_of_week || ', ' || LEFT(date, 10) AS day_and_date`로도 사용할 수 있다. 

#### UPPER, LOWER
- 해당 column들의 `case`를 `upper`나 `lower`로 통일시킴  

예제 ) `Category` 열을 만드는데, 첫글자만 대문자 (`Category`열은 전체가 대문자)
```sql
SELECT CONCAT(LEFT(category, 1), lower(RIGHT(category, LENGTH(category) - 1))) AS cleaneed_category
  FROM tutorial.sf_crime_incidents_2014_01
```

#### 이외에도 이런 식의 문자를 가공하는 함수는 많은데, SQL마다 미묘한 차이가 있다. `Mode`에서는 `Postgre`를 사용하고 있음

### 문자열 -> 날짜 가공
1. `CAST`를 이용함  
`CAST(DATE_COL AS date)` = `DATE_COL :: date`
------------------------
- 갑자기 궁금 ) `date`하고 `timestamp` 차이점? [스택오버플로우](https://stackoverflow.com/questions/31761047/what-difference-between-the-date-time-datetime-and-timestamp-types)
```sql
DATE : 날짜만 필요하고 시간은 필요 없을 때
DATETIME : 날짜 + 시간 모두 필요할 때, YYYY-mm-DD HH:mm:SS 
TIMESTAMP : 날짜 + 시간 + time zone(어느 위치가 현재 어느 시간인지)을 포함함
TIME : HH:mm:SS 
```
- 단 `Mode`에서는 `datetime` 개념이 `timestamp`로 쓰이고 있음.
-------------------------------
2. `timestamp`에 `EXTRACT` 함수 적용하기
```sql
SELECT cleaned_date, -- timestamp ex) '2014-01-31'
       EXTRACT('year'   FROM cleaned_date) AS year,
       EXTRACT('month'  FROM cleaned_date) AS month,
       EXTRACT('day'    FROM cleaned_date) AS day,
       EXTRACT('hour'   FROM cleaned_date) AS hour,
       EXTRACT('minute' FROM cleaned_date) AS minute,
       EXTRACT('second' FROM cleaned_date) AS second,
       EXTRACT('decade' FROM cleaned_date) AS decade, -- decade : year // 10
       EXTRACT('dow'    FROM cleaned_date) AS day_of_week -- dow : 1 ~ 7, 1이 아마도 월요일?
  FROM tutorial.sf_crime_incidents_cleandate
```  
  
3. `timestmap`에 `DATE_TRUNC` 함수 적용하기
- `DATE_TRUNC` : 함수를 해당하는 단위까지 반올림함
```SQL
SELECT cleaned_date, -- '2014-03-31 17:25:00'
       DATE_TRUNC('year'   , cleaned_date) AS year, -- '2014-01-01 00:00:00'
       DATE_TRUNC('month'  , cleaned_date) AS month, -- '2014-03-01 00:00:00'
       DATE_TRUNC('week'   , cleaned_date) AS week, -- '2014-03-27 00:00:00'
       DATE_TRUNC('day'    , cleaned_date) AS day, -- '2014-03-31 00:00:00'
       DATE_TRUNC('hour'   , cleaned_date) AS hour, -- '2014-03-31 17:00:00'
       DATE_TRUNC('minute' , cleaned_date) AS minute,
       DATE_TRUNC('second' , cleaned_date) AS second,
       DATE_TRUNC('decade' , cleaned_date) AS decade -- '2010-01-01 00:00:00'
  FROM tutorial.sf_crime_incidents_cleandate
```
- 참고 ) DATE_TRUNC 결과에 `CAST` 함수를 적용해 `DATE`로 바꾸더라도, 시-분-초 데이터는 `00:00:00`으로 저장된다. 없어지는 게 아님
- (아마도) `PostgreSQL`기준일 듯

#### 시간 관련 자료형
```SQL
SELECT CURRENT_DATE AS date,
       CURRENT_TIME AS time,
       CURRENT_TIMESTAMP AS timestamp, -- 다른 SQL의 DATETIME 역할
       LOCALTIME AS localtime,
       LOCALTIMESTAMP AS localtimestamp, -- 다른 SQL의 TIMESTAMP 역할
       NOW() AS now
```
- `PST` : `Pacific Stantard Time`(UTC-8) 
- `now()`는 그냥 부를 수 있는 함수, 현재 시간을 `timestamp`로 불러온다.

### COALESCE
- 결측치가 있는 열에 어떤 값을 추가하는 역할
- `COALESCE(column, value)`로 사용 가능, value에는 수치형, 문자형 다 가능.

## Subquery
- `Inner Query`, `Nested Query`라고도 한다.  

예시) 
```sql
SELECT sub.*
  FROM (
        SELECT *
          FROM tutorial.sf_crime_incidents_2014_01
         WHERE day_of_week = 'Friday'
       ) sub
 WHERE sub.resolution = 'NONE'
```
- 작동 순서
1. FROM 문 내부의 `Inner Query`부터 작동함 : **내부 쿼리는 별도의 쿼리로 작동함**
   - 이 내부 쿼리의 결과는 별도의 `Alias`를 지정해줘야 함(외부 쿼리에서 이용하기 위함)
2. 내부 쿼리가 작동을 마치면, **외부 쿼리는 내부 쿼리의 결과만을 이용하여 작동됨**

- 별도
- `Subquery` 실행 코드는 두 칸 띄어 들여쓰기를 해주자

- 궁금증
```sql
SELECT * 
  FROM tutorial.sf_crime_incidents_2014_01
 WHERE day_of_week = 'Friday'
   AND resolution = 'NONE'
```
- 이 쿼리랑 실행 방법은 달라도 결국 같은 거 수행하는 거 아닌가?
- 실제로 위의 내용은 `WHERE`문으로 해결할 수 있음. `SUBQUERY`를 필수로 요하지 않음.

### SUBQUERY만으로 해결할 수 있는 예제
- 여러 단계에 걸쳐 `AGGREGATE`가 필요한 경우
```SQL
SELECT LEFT(sub.date, 2) AS cleaned_month,
       sub.day_of_week,
       AVG(sub.incidents) AS average_incidents
  FROM (
        SELECT day_of_week,
               date,
               COUNT(incidnt_num) AS incidents
          FROM tutorial.sf_crime_incidents_2014_01
         GROUP BY 1,2
       ) sub
 GROUP BY 1,2
 ORDER BY 1,2
```
- 내부 쿼리 -> 외부 쿼리로 진행하는 순서가 자연스러움(내부 쿼리를 따로 실행해서)
- 원본 테이블은 `각 사건이 어떤 날짜에 발생했는가`가 저장되어 있는데, 외부 쿼리의 `AVG()`문은 `어떤 달, 어떤 요일 당 발생한 평균 사건 수`가 된다.
  - 내부 쿼리에서는 `요일 & 날짜`을 추려내 `COUNT()`를 적용했음
  - 외부 쿼리에서는 `날짜`에서 다시 `월`을 뽑아냈고, 뽑아낸 `월`을 이용해 `AVG()`를 적용했음. 
----------------
예제) 범죄 각 카테고리에 대해 월 평균 사건 수  
- 월 평균 사건 수에 필요한 정보
  1. 각 카테고리에 대해, 어떤 달에 몇 건의 사건이 발생했는지에 대한 정보가 필요함
  2. 위의 숫자 정보를 평균 내야 함.
* 서브 쿼리 없이 해결할 수 없음 : 테이블에서 뽑아낸 통계 자료를 다시 변수로 넣어야 하는 과정이 들어가는데, 단일 쿼리로는 통계 자료를 변수로 활용할 수 없기 때문임(`Group by`로 묶는 정도만 가능함)
```sql
SELECT sub.category,
       AVG(sub.counts)
  FROM (
        SELECT EXTRACT('month' FROM cleaned_date :: timestamp) as month,
          category,
          COUNT(*) AS counts
        FROM tutorial.sf_crime_incidents_cleandate 
        GROUP BY 1, 2
        ORDER BY 3 DESC
        ) sub
 GROUP BY 1
 ORDER BY 2 DESC
```
-----------------------------
1. 서브쿼리는 즉 테이블에서 **뽑아낸 통계값을 다시 변수로 활용하고 싶을 때 사용**할 수 있음(이전에 부딪힌 적 있는 문제 : s**ql에서 추려낸 정보를 다시 변수로 이용할 수 있는가?를 해결.**)

### 조건문을 사용한 서브쿼리
```sql
SELECT *
  FROM tutorial.sf_crime_incidents_2014_01
 WHERE Date = (SELECT MIN(date)
                 FROM tutorial.sf_crime_incidents_2014_01
              )
--  WHERE Date = MIN(date) -- 사용 불가
```
- 서브쿼리문으로 만족되는 값이 1개이기 때문에 `=`을 쓸 수 있다.
- 서브쿼리문이 여러 개라면, `IN`을 써야 함(`WHERE`의 원래 용법과 동일함)
- `WHERE`문의 서브쿼리는 테이블로 취급하지 않기 때문에, `ALIAS`를 사용하지 않는다.(`값`으로 취급됨)

### 서브쿼리 JOIN하기
```SQL
SELECT *
  FROM tutorial.sf_crime_incidents_2014_01 incidents
  JOIN ( SELECT date
           FROM tutorial.sf_crime_incidents_2014_01
          ORDER BY date
          LIMIT 5
       ) sub
    ON incidents.date = sub.date
```
- `JOIN`문의 서브쿼리는 테이블이니 `ALIAS`를 붙여준다.

예제 ) 가장 적은 사고가 일어난 3개의 카테고리에 대해 모든 정보를 보여주셈  
```SQL
SELECT main.*,
       sub.incidents
  FROM tutorial.sf_crime_incidents_2014_01 main
  JOIN (
        SELECT category,
             count(*) as incidents
          FROM tutorial.sf_crime_incidents_2014_01
         GROUP BY 1
         ORDER BY 2 
         LIMIT 3 
        ) sub
  ON main.category = sub.category
```
- 서브쿼리로 통계치 값이 별도의 테이블로 빠졌기 때문에, 이를 main 테이블의 정보와 같이 표시해도 잘 출력됨 
- 메인 테이블에서 각 레포트에 해당하는 정보가 출력되고, 서브 테이블에서 어떤 카테고리가 몇 개 나왔는지 나타난다.

#### 쿼리 최적화
- 2개의 예시를 보고 비교
```sql

-- 실행 시간 : 19s
SELECT COALESCE(acquisitions.acquired_month, investments.funded_month) AS month,
      COUNT(DISTINCT acquisitions.company_permalink) AS companies_acquired,
      COUNT(DISTINCT investments.company_permalink) AS investments
  FROM tutorial.crunchbase_acquisitions acquisitions
  FULL JOIN tutorial.crunchbase_investments investments
    ON acquisitions.acquired_month = investments.funded_month
GROUP BY 1
```
```sql
-- 실행 시간 : 803ms
SELECT COALESCE(acquisitions.month, investments.month) AS month,
      acquisitions.companies_acquired,
      investments.companies_rec_investment
  FROM (
        SELECT acquired_month AS month,
              COUNT(DISTINCT company_permalink) AS companies_acquired
          FROM tutorial.crunchbase_acquisitions
        GROUP BY 1
      ) acquisitions

  FULL JOIN (
        SELECT funded_month AS month,
              COUNT(DISTINCT company_permalink) AS companies_rec_investment
          FROM tutorial.crunchbase_investments
        GROUP BY 1
      )investments

    ON acquisitions.month = investments.month
ORDER BY 1 DESC
```
- 전자는 `FULL JOIN`한 다음 `AGG` 함수를 적용한 것이고, 후자는 각 테이블에 `AGG`함수를 적용한 다음에 `FULL JOIN`한 것이다.
1. `COUNT(DISTINCT ~~)` : 혹시 모르게 있을 중복된 데이터 값들에 대해, 하나로 처리해줌
2. `FULL JOIN`할 때 `COUNT()`와 `COUNT(DISTINCT )`의 차이점이 있다 : 그냥 `COUNT()`가 훨씬 빠르다(`6s` vs `19s`)
  - 이거는 뒤에서 다시 다룰 예정
3. 아무튼 외적 연산이 들어가기 떄문에, **테이블을 합치기 전에 조건을 추려놓는 것이 좋다**는 의미 같다.

#### Subquery with UNION
- 보통 테이블을 합쳐놓고 작업하니까, `UNION`문은 `FROM`문에 서브쿼리로 집어넣을 수 있다. 

## Windows Function
[링크](https://mode.com/sql-tutorial/sql-window-functions/)