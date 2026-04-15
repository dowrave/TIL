- 현재 어떤 작업 중인지 기록 중
# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.

---
# 작업 일지

## 블로그 고치고 싶은거
## 짭명방 
- 지난 내역 : [짭명방 프로젝트 일지 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- [[기타 참고 사항]]

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

### 최근 작업 내용
- 블로그 : [[블로그_260127 EC2에서 Lightsail로 이사 가기]]
- 짭명방 : [[짭명방_260415]]
	- BossBGM 2가지 버전으로 만든 후 보스 스폰 시의 BGM 전환 로직 구현 중

>[!note]
>1. 효과음 구현
>	- 생성 / 탐색 : `Soundly` / `FreeSound`
>	- 수정 : `FL Studio`
>	- 유의 : `-6&-9 dBFS`, `-20 LUFS` 등의 **수치에 대한 기준은 있지만 결국은 귀임**
>	- 현재 없는 소리
>2. 발견한 버그
>	- 허공에 헛스윙하는 문제 : 여전히 간헐적으로 발생 중
>	- `ArcaneFieldSkill`도 간헐적으로 사운드 밀리는 게 있는 거 같음.
>3. 기타 필요성 느낀 지점

>[!wip]
>- 작업 중
>	- `BossStage` BGM
>		- 보스가 나타났을 때 / 나타나지 않았을 때의 음악 전환
>- 작업 예정
>	- `Defender`의 방어력이 너무 높은 듯?
>	- SFX
>		- 스테이지 클리어 시점의 사운드
>		- 스테이지 클리어 후 통계 패널 사운드(BGM?)
>		- `EnemyBoss`
>			- `SlashThroughSkill`
>				- 긋는 소리 : `Slash01/02`와 다른 걸 추가해야 할 듯?




# 이전 일지
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크(Private 전환)](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
