# Yammer 데이터 분석 
[데이터리안 무료 강의](https://www.inflearn.com/course/%EB%8D%B0%EC%9D%B4%ED%84%B0-%EB%B6%84%EC%84%9D-sql-%EC%8B%A4%EC%A0%84/unit/72141?tab=curriculum)
- 이 강의를 왜 듣는가?
  - 데이터를 딱 받았을 때 어떻게 접근해야 할 지가 좀 막막한 점이 있었다.
  - SQL의 숙련도 이슈도 물론 있다.
  - 일단은 정확도에 집중한 다음 속도에 집중해야 하는데, 자꾸 2가지를 한꺼번에 잡으려고 하니까 산만해지는 점이 있다.
  - 따라서, 전문가는 어떻게 접근했는지 알아보고자 듣게 되었다.


## 여는 강의

- 따릉이 등의 공공데이터 분석도 의미가 있는데, 회사의 데이터를 다룰 기회가 많이 없음
- `Yammer` : 페북과 유사한 사내 메신저
- 현업에서 푸는 문제와 유사하고, 지금 풀고 있는 문제를 얘기할 수 있는 기회라 생각하여 강의를 기획함

### 케이스 3가지
1. 활성 유저의 감소 원인
2. 검색 기능 분석
3. A/B 테스트 결과 검증 

### 프로젝트 5단계
1. 준비 : 사전 자료 스스로 학습
    - 사전 자료를 스스로 학습하는 게 전체 학습량의 90% 이상임
    - 나오는 전문 용어를 직접 찾아보는 것도 필요
    - 베스트는 이 사전 자료만으로 공부를 마쳤다!라고 판단이 드는 거고
    - 강의 내용은 이해되지 않은 내용을 확인하거나, 강사가 제안하는 새로운 관점을 보는 정도로 파악하면 좋다.
2. 문제상황 : 프로젝트의 문제상황 파악
3. 분석 : 분석 해설
4. 주요 SQL 해설 : 분석에 사용한 주요 SQL 해설
5. 프로젝트 요약

## 1. 활성 유저의 감소 원인
- 시작하기 전에, 회사에 대해 알아두자.
### 1. 문제 파악
- `User Engagement Dashboard`
  - `DAU : Daily Active User / Weekly AU / Monthly AU` 
  - 특정 기간 동안 활성 유저를 파악함  

(문서 중) `Yammer`의 `Engagement` 정의 : `Server Call`을 발생시키는 모든 행위, `events` 테이블에 `Engagement`로 정의되어 있다.
- 해당 지표가 감소하고 있으니, 어떤 일이 있었는지 밝히자. (해결책도 같이 제시하면 좋을 것이다.)
- 실제로도 **매일 확인하는 지표(매출 등)는 사장이나 매니저가 계속 보는 요소**이다. 즉 많이 발생하는 케이스임.  

(문서 중) 데이터에 접근하기 전, **가능한 원인들을 쭉 적어보고 테스트할 순서를 만들어라.** 
- 여러 원인이 있을 수 있는데, 어떤 순서로 테스트할 지 생각해보고, 그 이유도 생각하라.
- 차트가 보여주는 것과 보여주지 않는 것도 생각하라.  
- **생각한 모든 것을 테스트해볼 순 없음.** 중요하다고 생각하는 것과 그 순서 및 이유를 생각해보자.

(강의 내용) `Engagement`의 정의는 회사마다 다를 수 있다. 로그인을 했을 때, 어떤 기능을 썼을 때, 결제를 했을 때 등 `Active`에 대한 정의가 다르다는 것.

### 2. 테이블 분석
(강의 내용) 4개의 테이블이 주어지는데, 3개의 테이블만 사용할 것.  
(강의 내용) 이 예제에는 테이블의 이름 및 각 Column에 대한 **설명**이 있지만, 이것도 사람이 하는 일이기 때문에 **여러분이 다니는 회사에는 없을 수도 있다**.   

- 테이블 설명은 내가 해본 거니까 생략함  


(문서 내용) 차트의 `View Query Results`를 누르면 쿼리 결과 테이블이 나오고 이는 복사해서 이용할 수 있음  

(문서 내용)
1. 처음에 세운 가설이 추가적인 질문을 하게 만드는가?
2. 만약 그렇다면 그게 무엇이고 어떻게 테스트할지?
3. 데이터만 갖고 대답할 수 없는 가설이라면, 어떻게 가설을 테스트할 수 있는가?(회사에 정말 다닌다고 생각하면 어떤 것들을 할 수 있을까?)
4. 어떤게 `Engagement Dip`에 가장 큰 영향을 줬는가?
5. 회사는 어떻게 대응해야 할까?  
(강의 내용) 이 질문들에 어떻게 대답할지 고민해보면 좋을 것 같다. 문서의 `Answers`가 마냥 정답은 아니다.

### 3. 쿼리하는 법
- 일단 문서에는 차트와 해당 차트를 만든 쿼리가 같이 있다. 쿼리 쓰는 법은 아니까 패스
- 직접 이런저런 쿼리를 작성하고 데이터를 뽑아보자.

### 4. 강사의 요약 내용
[분석 내용](https://docs.google.com/spreadsheets/d/1JccUfsUFzX_Mc80LqDRUpFEzqGtVVzhhW7tu8pk1RD0/edit#gid=231328655)
- 스프레드시트 깔끔하다..!
- 일단 어떤 부분을 분석했는가는 문서 내용과 동일하기 때문에 별도로 필기하진 않겠음
- 근데 쿼리 내용이 다르다는 걸 본 기억이 있어서 그건 검토해봄  
- 더 나아간 내용
  1. 8월 4일 이후의 WAU 감소가 정말 문제인가? 매 년, 매 분기 반복되는 패턴일 수 있다.
  2. 데이터를 분해하자
    - 전체를 부분으로 나누면 전체를 볼 때는 알 수 없었던 변동 원인을 관찰할 수 있음
    - 분석이 잘 안되면 `feature`를 몇 개 추가해서 데이터를 더 깊게 살펴볼 수 있다.
    - `WAU`, `유저 코호트`, `이메일 인게이지먼트` 등을 더 깊게 파고들 수 있다.
      - 고객 유입 채널(광고), 고객 정보(성, 연령, 사용기기, 거주 지역 등), 국가, 코호트 범위 등

### 5. 쿼리 내용
(강의 내용)  
- 어떻게 원하는 데이터 추출을 했는가를 추적
- 강사님이 보고 치는 게 아니라 직접 어떤 걸 조회할지 생각하면서 친 내용임
1. WAU 추적
```sql
SELECT DATE_TRUNC('week', occurred_at) AS week,
       COUNT(DISTINCT user_id) as weekly_active_user
  FROM tutorial.yammer_events e 
 WHERE occurred_at BETWEEN '2014-04-28 00:00:00' AND '2014-08-25 23:59:59'-- time_window
   AND event_type = 'engagement'
 GROUP BY week
 ORDER BY week
```
- `Mode`는 `PostgreSQL`을 씀 : `MySQL`과는 일부 용법이 다르다
  - 예를 들면 `DATE_TRUNC()`은 `MySQL`에는 없나봄?
- 문서에 있는 `WAU` 차트의 경우, `time window`를 주지 않았는데, 상용 테이블에서 이러면 큰일난다 : 데이터 조회 비용이 어마어마해짐
  - 따라서 **조회할 기간을 정해놓는 건 가장 먼저 할 일**이다.
- `GROUP BY`, `ORDER BY`에 쓰이는 `COLUMN`은 명시를 하는 것이 좋다 : 복잡한 쿼리일수록 어떤 기준으로 묶고 나열했는지 직관적으로 파악할 수 있도록!

2. 날짜와 가입자 수(ACTIVE, PENDING 포함)
```sql
SELECT DATE_TRUNC('day', created_at) AS signup_date
     , COUNT(user_id) AS signup_users
     , COUNT(CASE WHEN activated_at IS NOT NULL THEN user_id ELSE NULL END) as activated_users-- 활성화 유저 수
     
  FROM tutorial.yammer_users
 WHERE created_at BETWEEN '2014-06-01 00:00:00' AND '2014-08-31 23:59:59'
 GROUP BY signup_date
```
- 쿼리를 치며 유념할 것은, `스케일`에 관한 것이다.
  - 예를 들어 수십 명 단위를 생각하고 쿼리를 짰는데 결과가 수백 명이 나왔다면, 데이터가 잘못됐을 확률보다 쿼리가 잘못됐을 확률이 더 높다.
- **결과물을 예상하면서 쿼리를 짜야 한다.**  

3. 이메일 관련 쿼리 2가지 (ANSWER의 4번 하단, 5번 하단)
1) 단순 이메일 종류에 따른 숫자 카운트
```SQL
SELECT DATE_TRUNC('week', occurred_at) AS WEEK
     , action
     , COUNT(DISTINCT user_id) AS cnt_user
  FROM tutorial.yammer_emails em
 WHERE occurred_at BETWEEN '2014-04-28 00:00:00' AND '2014-08-25 23:59:59'
 GROUP BY week, action
```
- 문서의 내용은 쿼리 결과 표를 `날짜 - 종류1 카운트 - 종류2 카운트 - ...` 식으로 보고 싶어서 저렇게 짠 거임(피벗팅)
  - 피벗팅이 더 직관적이긴 하다. 대신 코드가 조금 더 복잡해지는 점은 있다.
- 근데 코드를 위처럼 짜고 그 결과를 복사해서 스프레드시트에 넣은 뒤 피벗테이블 기능을 이용할 수 있음!
  - 피벗테이블 기능이 이해가 안간다면 꼭 SQL로 할 필요는 없다!
- 실제로 위처럼 쿼리를 짠 다음 시각화 기능을 이용해도 비슷하게 나옴


2) 이메일이 보내진 뒤 5분 이내에 메일을 클릭했는가 아닌가
```sql
SELECT DATE_TRUNC('week', em1.occurred_at) as week
    , COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em1.user_id ELSE NULL END) as sent_digest_emails
    , COUNT(CASE WHEN em2.action = 'email_open' THEN em2.user_id ELSE NULL END) as sent_digest_opened
    , COUNT(CASE WHEN em3.action = 'email_clickthrough' THEN em3.user_id ELSE NULL END) as sent_digest_clickthrough
    -- 강의의 쿼리는 아래와 같다 : 아래 3개 주석도 위와 똑같이 작동함
    -- , COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em1.user_id ELSE NULL END) as lecture_digest_emails
    -- , COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em2.user_id ELSE NULL END) as lecture_digest_opened
    -- , COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em3.user_id ELSE NULL END) as lecture_digest_clickthrough

    -- 퍼센트로 표시 : float로 데이터를 꼭 바꿔줘야 한다
    , COUNT(CASE WHEN em2.action = 'email_open' THEN em2.user_id ELSE NULL END) / COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em1.user_id ELSE NULL END) :: float || '%' as open_pct
    , COUNT(CASE WHEN em3.action = 'email_clickthrough' THEN em3.user_id ELSE NULL END) / COUNT(CASE WHEN em1.action = 'sent_weekly_digest' THEN em1.user_id ELSE NULL END) :: float || '%'as ct_pct
FROM tutorial.yammer_emails em1
    LEFT JOIN tutorial.yammer_emails em2
      ON em1.user_id = em2.user_id 
     AND em2.occurred_at BETWEEN em1.occurred_at AND em1.occurred_at + INTERVAL '5 min'
     AND em2.action = 'email_open'
    LEFT JOIN tutorial.yammer_emails em3
      ON em1.user_id = em3.user_id 
     AND em3.occurred_at BETWEEN em2.occurred_at AND em2.occurred_at + INTERVAL '5 min'
     AND em3.action = 'email_clickthrough'
WHERE em1.occurred_at BETWEEN DATE('2014-07-01') AND DATE('2014-08-31')
  AND em1.action = 'sent_weekly_digest'
GROUP BY week

```
- 문서 내용은 엄청 복잡함(서브쿼리에 서브쿼리에 서브쿼리)
- 근데 `SELF JOIN`과 `ON` 조건을 이용하면 꽤 많이 추릴 수 있음
- 특히 `LEFT JOIN`과 `Condition`을 어떻게 활용하는지를 집중적으로 보자 : 원본의 매우 복잡한 쿼리를 단순화하는 좋은 쿼리라고 생각함
- `SELECT`문의 위, 아래 두 방법 모두 잘 작동함(왜인지 모를 이슈로 아래 쿼리가 작동 안했음)
- 강의에서 구글 스프레드시트를 활용할 것을 권장하고 있음
  - SQL로 피벗테이블을 하려면 해당되는 테이블 하나하나를 CASE WHEN 문에 집어넣어줘야 함
  - 스프레드시트로 피벗테이블을 하겠다면 SQL로 `GROUP BY`까지만 하고 그 결과를 옮기면 더 쉽고 빠르게 결과를 볼 수 있음.
- 집계 값들을 연산하고자 할 때(특히 백분율) -> `float`로 값을 **꼭 !** 변환해야 한다. 변환 안하면 그냥 0으로 나옴.

