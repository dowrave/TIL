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
#### Concat부터 계속
[링크](https://mode.com/sql-tutorial/sql-string-functions-for-cleaning/)