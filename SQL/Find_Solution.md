## SQL을 직접 다루면서 알게 되는 것들을 정리한 문서

## 1. 여러 행에 걸쳐 중복된 값이 나타날 때, DISTINCT에 대해
```SQL
SELECT DISTINCT user_id, device

```
- 예를 들면 어떤 유저가 어떤 제품을 주문하려 하는 과정에서 로그가 기록된다고 하자.
- **DISTINCT는 모든 COLUMN에 대해 적용된다.**
  1. `(COL1), DISTINCT (COL2)` 식으로 사용할 수 없다는 게 있고,
  2. `DISTINCT COL1, COL2, COL3, ...`인 경우, 튜플로 `COL1, COL2, COL3, ...`을 생각할 수 있을 것이다. 이 때 `DISTINCT`는 **튜플이 완전히 똑같지 않다면 제거하지 않는다.**