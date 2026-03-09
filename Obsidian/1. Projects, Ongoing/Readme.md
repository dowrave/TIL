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

>[!note]
>
>  **"흐름"으로 설명하는 것 외에도 "원리"도 곁들이면 좋다**
> - 지금은 "어떤 이슈가 발생해서 어떻게 해결했더라"로 정리했다면
> - 추가로 **이슈가 '왜' 발생했는지**까지 정리해두면 나중에 볼 때 다시 헷갈리지 않을 수 있음
> 
> 모든 코드를 일일이 기억하는 건 불가능함! 나중에 봤을 때 '이러면 되겠다'는 감이 바로 잡히도록 메모해보자.

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움
> - `Enemy`가 사라질 때 풀 태그의 키가 없다는 오류 
> - `Barricade` 배치 시 가끔 배치되지 않음
> - `Operator` 체력이 다했을 때 사망처리가 되지 않았는데도 적이 지나가는 현상

## 현재 계획 및 이슈

>[!plan]
>1. 스테이지 밸런싱, 버그 수정
>2. 배경음, 효과음 추가 및 수정

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

### 최근 작업 내용
- 블로그 : [[블로그_260127 EC2에서 Lightsail로 이사 가기]]
- 짭명방 : [[짭명방_260305]]

>[!wip]
>1. 효과음 구현
>	- 생성 / 탐색 : `Soundly` / `ElevenLabs` / `FreeSound`
>	- 수정 : `FL Studio`
>	- 유의 : `-6&-9 dBFS`, `-20 LUFS` 등의 **수치에 대한 기준은 있지만 결국은 귀임**
>	- 현재 없는 소리
>		- `Boss` 관련 스킬, 사운드
>2. 발견한 버그
>	- 허공에 헛스윙하는 문제 : 여전히 간헐적으로 발생 중
>	- `ArcaneFieldSkill`도 간헐적으로 사운드 밀리는 게 있는 거 같음.
>3. 기타 필요성 느낀 지점

>[!note]
>- 버그 기록 및 개선할 요소
>- 유지
>	- 로비 BGM 
>		- Slap 외의 다른 베이스 테스트 & 후반부가 초딩 때의 음악 같은 인상이 있어서 손대봤으나 기존 버전이 더 낫고 긴 노트를 억지로 끊는 게 더 어색하게 들림
>		- 어쩌면 기존 BGM에 너무 익숙해진 걸지도 모르겠는데... 내 생각은 그렇다.

## 260309

>[!done]
>- 배치 미리보기 : 배치 시 배치 불가능/가능 여부에 따라 공격 범위가 흑백 / 컬러로 바뀜
>- 이슈 수정
>	- 마스터 볼륨이 -80dB임에도 제대로 적용되지 않는 현상 : `Awake` => `Start`
>	- 스킬 켜졌을 떄 SP 게이지 색깔 바뀌지 않는 현상 
>	- `AreaHasteHeal` : HP가 가득 찼을 땐 힐하지 않도록 수정

## 배치 로직 시각화
- **배치가 불가능할 땐 공격 범위를 회색으로, 배치 가능할 때 색이 들어오는 방식**으로 수정 진행
- `DeployableManager`에서 `Tile`들의 `ShowAttackRange`를 켜는 방식임
- 오해 : 이중 처리 아닌가? `DeploymentInputHandler -> Operator -> DeployableManager`라는 방식으로 하이라이트가 처리 중이었음

- 공격 범위를 처리하는 타일 자체의 로직. `isPlaceable`이 추가됨.
```cs
public void ShowAttackRange(bool isMedic, bool isPlaceable = true)
{
	Color targetColor;
	
	if (!isPlaceable)
		targetColor = disabledIndicatorColor;
	else 
		targetColor = isMedic ? medicIndicatorColor : defaultIndicatorColor;

	attackRangeIndicator.GetPropertyBlock(indicatorPropBlock);
	indicatorPropBlock.SetColor("_Color", targetColor); // URP Lit과 달리, attackRange라는 별도의 셰이더가 있고 Color라는 프로퍼티가 있음
	attackRangeIndicator.SetPropertyBlock(indicatorPropBlock);

	attackRangeIndicator.gameObject.SetActive(true);
}
```

### 테스트 및 수정
- 배치 불가능한 상태의 색을 RGB 0.3으로 잡았는데 너무 어두움 -> 0.5로 수정
- `MouseButtonDown`만 처리했는데 드래그 중인 상황도 추가로 처리
- 버튼 다운 시 방향 변경할 때만 Update되는 로직 수정
```cs
// if (placementDirection != newDirection)
// {
	placementDirection = newDirection;
	if (currentDeployableEntity is Operator op)
	{
		op.SetDirection(placementDirection);
		op.HighlightAttackRange(isPlaceable: pointerOverDistance);
	}
	deployManager.UpdatePreviewRotation(placementDirection);
// }
```
> 원래는 방향이 변경될 때만 업데이트됐는데, 이제 거리도 반영되기 때문에 변경 로직 자체가 필요없어 보임. 미리보기 동작이 끊기는 느낌도 아니라서 주석처럼 제거함.


## 이슈 수정
### 볼륨 관련
- 에디터 다시 켰을 때의 볼륨 설정이 반영되지 않는 현상

![[Pasted image 20260309152457.png]]
> - 게임을 켰을 때 값이 이렇게 잡히지만 실제로는 소리가 나오고 있음
> - 씬 변경 후 다시 돌아왔을 때는 잘 반영되고 있음

#### 과정
- `Awake` 시점에서 볼륨들이 세팅되고 있음
```cs
private void Awake()
{
	// PlayerPrefs에 저장된 값들 불러오기
	masterVolume = GetMasterVolume();
	bgmVolume = GetBGMVolume();
	sfxVolume = GetSFXVolume();

	// 값 적용하기(없으면 1)
	SetMasterVolume(masterVolume);
	SetBGMVolume(bgmVolume);
	SetSFXVolume(sfxVolume);
}
```

그런데 이 시점에서는 `AudioMixer`의 초기화가 제대로 되지 않았을 수 있음 : `Awake`이기 때문- 
- `Start`로 바꿔주니 잘 적용됨

### 스킬 켜졌을 때 SP 게이지 색깔 안 바뀌는 현상
- `OpSkillController.OnSkillStateChanged`가 있는데 사용하지 않았음
	- `bool` 필드를 받도록 변경
- 처리하는 김에 `_isSkillOn` 상태를 변화하는 로직을 전부 `SetSkillState` 메서드 내에 집어넣고 그 안에서 이벤트 발생시키는 식으로 변경

# 이전 일지
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
