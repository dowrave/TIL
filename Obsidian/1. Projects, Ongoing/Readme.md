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

>[!wip]
>- 작업 중
>	- `Localization` 지원 및 구조 변경
>- 작업 예정
>	- 스테이지 1-3 구현 끝내기
>	- 기타 이슈
>		- `Defender`의 너무 높은 방어력
>		- `BossSlashSkill`의 효과음 수정?
>		- BGM의 저음 줄이기?

>[!note]
>- 마지막 맵 테스트 전 추가로 구현하고 싶은 것들
>- `BossBGM`의 보스 스폰 후의 BGM으로 전환되더라도 크게 바뀌었는지 모르겠다. 이 부분은 수정이 필요해보임.
>- `LightningHit`의 소리가 겹칠 때 너무 커진다. 갯수는 2개로 제한했는데도 그런데, 이전처럼 관리하고 제어하는 방법밖에 없나?

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.
## 최근 작업 내용 - 짭명방
- [[짭명방_260615]]

>[!done]
>- `Skill Localization` 마무리
>	- `BattoSkill`
>	- `TacticalStrikeSkill`
>	- `HealingZoneSkill`
>- `Skill Localization` 점검 및 되지 않았던 부분 / 수정할 부분 수정
>	- `AreaHasteHealSkill` 구현 X : `HealingZoneSkill`이란 이름으로 구현
>	- 공격 속도 툴팁 수정 : 공격 쿨다운 감소율을 함께 표시하는 방식
>	- 한국어의 배치 코스트 정리 : `배치 코스트를 ~만큼 얻습니다.

## 최근 작업 내용 - 블로그
- 마지막 수정 내역 : [[블로그_260614 - 모바일 레이아웃 수정]]

>[!done]
>1. `WriteReviewModal`(새로운 리뷰 생성)과 `ReviewModal` 간의 서식 불일치 수정
>2. 모바일 레이아웃에서 텍스트박스가 차지하는 영역 설정
>3. ~~모바일 툴바와 실제 툴바 구분~~ 툴바에서 안 쓰는 기능 정리, 모바일에서도 1줄로 보이게 설정
>4. 모바일에서 키보드가 나타났을 때 텍스트 박스만 보이는지 테스트 - 굿

# 이전 일지
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크(프로젝트 자체는 Private 전환)](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
