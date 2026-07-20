- 현재 어떤 작업 중인지 기록 중
# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

---
# 작업 일지
## 짭명방 
- 지난 내역 : [짭명방 프로젝트 일지 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- [[기타 참고 사항]]

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움
> - `Artillery`의 공격 효과음이 제대로 나오지 않는 현상

>[!note]
>- **마무리 작업**
>1. `Enemy`의 패턴 유한 상태 머신 기반으로 수정
>2. 최초 메인 화면 진입 전의 화면 구현
>3. 엔딩 구현

>[!note]
>- 마지막 맵 테스트 전 추가로 구현하고 싶은 것들
>- `BossBGM`의 보스 스폰 후의 BGM으로 전환되더라도 크게 바뀌었는지 모르겠다. 이 부분은 수정이 필요해보임.
>- `LightningHit`의 소리가 겹칠 때 너무 커진다. 갯수는 2개로 제한했는데도 그런데, 이전처럼 관리하고 제어하는 방법밖에 없나?

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

- **(260720) `EnemyState` 관련 앞으로 할 일**
- `StunState`, `DeathState` 추가
	- 지금까지 구현한 3개와 달리, 이들의 전환 트리거는 외부 요인이어야 함.
	- 일단 **"전역 인터럽트는 FSM 레벨에서 한 번 체크해야 한다"** (Update 이전에) 정도만 기록해둠.

## 최근 5일간 작업 내용 - 짭명방

### [[짭명방_260720]]
>[!done]
>- `EnemySkillCastState` 관련 작업 완료
>	- `CastTime`의 관리 주체, 스킬이 끊겼을 때 어떻게 처리하는가 등등 


### [[짭명방_260715]]
> [!done]
> - Enemy FSM 계속
> 	- `EnemySkillCastState` 구현 중(완료 X)
> - 이슈 수정
> 	- 길이 막힌 상태에서 `Barricade`가 도착 지점처럼 처리되는 현상
> 		- `fsm`에 구현한 `combatState` 전환 메서드를 `MoveState`로 옮김
> 		- `PathController` 리팩토링
> 	- `Barricade` 위치에 도착했음에도 파괴하지 않는 현상(간헐적 - 수정 완료)
- [[CPU 연산 - 제곱과 제곱근 비교]]


### [[짭명방_260714]]
>[!done]
>- `Combat` 상태 구현 완료
>	- FSM에서 "감지"와 "전환 판단"의 책임 분리
>	- `EnemyAttackController.OnUpdate()` 코드 리뷰
>	- 반환 타입: `bool` vs `enum`
>	- `EnemySkillController` 추가 및 `EnemyBoss` 확장
>		- `CombatState`에서 스킬과 공격을 모두 처리하도록 처리했음
>	- Combat ↔ Move 오실레이션 버그 두 가지 발견 및 수정 방향
- 아직 `Move`는 처리할 게 좀 있음

### [[짭명방_260713]]
>[!wip]
> - `Enemy`의 `FSM` 구현하기 : 학습과 구현  기초

### [[짭명방_260710]]
>[!done]
>- 유한 상태 머신 공부 및 Enemy에 적용 시작
- [[유한 상태 머신(FSM)]] 문서 분리

## 최근 작업 내용 - 블로그
- 마지막 수정 내역 : [[블로그_260622 - 이것저것]]

>[!note]
>- 메인 페이지 레이아웃, 아카이브 설정, 백엔드 최적화 등등 이것저것 처리했음.
# 이전 일지
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크(프로젝트 자체는 Private 전환)](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
