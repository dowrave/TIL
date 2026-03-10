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
- 짭명방 : [[짭명방_260309]]

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
>- 생각해봤으나 유지하기로 결정
>	- 로비 BGM 

## 260310
>[!done]
>- 1-3 스테이지 구현
>- 버그 / 기타 수정
>	- `DoubleShotSkill` : 지속시간 VFX 나타나지 않는 현상
>	- `Operator`들 색깔 변경

## 1-3 스테이지
- 스폰 간격 늘림
	- 한꺼번에 쏟아지는 경우가 나왔을 때 그 다음의 간격을 살짝 늘려줌
	- 예를 들어 `Weak`이 3마리 나왔다(`Weak` 간의 간격은 1초)면 기존엔 3초 후에 그 다음 몹이 스폰됐는데 5초로 늘렸음

- 

## 버그 / 기타 수정
- `DoubleShotSkill` - 지속시간 VFX 나타나지 않는 현상
```cs
public override void OnActivated(Operator op)
{

	_doubleShotBuff = new DoubleShotBuff(delayBetweenShots, damageMultiplier);
	op.AddBuff(_doubleShotBuff);
	PlayDurationVFX(op, _doubleShotBuff); // 추가
}
```
### SFX 수정 : `MeleeImpact01` - 타격감이 너무 강해서 로우컷

### `Operator`들 색깔 변경

### 간헐적으로 발생하는 허공을 쏘는 이슈
- 오랜만에 다시 발생해서 점검해본다
- `Enemy` 자체는 이미 `Disabled` 된 상태, 그럼에도 빈 곳을 때리고 있는 현상임

- **일단 `Claude`에게 던져봄 : `OpAttackController`**
#### 이슈 1 : OnTargetDespawn
```cs
public override void OnTargetDespawn(UnitEntity target)
{
	if (target is Enemy enemy && enemy == CurrentTarget)
	{
		RemoveEnemyInRange(enemy); 
		CurrentTarget = null;
	}
}
```
> 여기서 `currentTarget`일 때만 현재 공격범위에서 제거하고 있다. 
> - 즉 현재 타겟이 아닌데 제거됐을 때, 공격 범위 내의 대상 리스트에는 여전히 남아 있을 가능성이 있음

- 따라서 상위 조건문에서 `and`로 묶는 게 아니라 세부 조건으로 빼줘야 함
```cs
public override void OnTargetDespawn(UnitEntity target)
{
	if (target is Enemy enemy)
	{
		RemoveEnemyInRange(enemy); 
		if (enemy == CurrentTarget)
		{
			CurrentTarget = null;
		}           
	}
}
```

#### 이슈 2 : SetCurrentTarget
```cs
if (blockedEnemies.Count > 0)
{
	for (int i = 0; i < blockedEnemies.Count; i++)
	{
		if (blockedEnemies[i]) // 이슈 1
		{
			CurrentTarget = blockedEnemies[i];
			break; 
		}
	}
	return; // 이슈 2
}
```
> - 지적 사항
> 1. `blockedEnemies[i]` 체크는 유니티 오브젝트의 bool 체크
> 	- 실제로는 **`null` 체크를 해야 하고, `activeInHierarchy` 체크도 해야 한다.** `null`의 경우 유니티와 C#의 동작이 다름. 
> 	- 오브젝트 풀링을 쓸 경우 `null`일 경우는 거의 없지만, `Destroy`가 일어나지 않는다는 보장은 아니기 때문에 조건문에서 둘 다 체크해주는 게 좋음
 > 2. `blockedEnemies`가 있으면 무조건 여기서 끝남 - 모든 `blockedEnemy`가 파괴되어 있더라도`return` 문에서 걸림(이것도 위의 유니티 - C# 동작 이슈와 연관 O)
 
 - 따라서 코드는 이렇게 수정됨
```cs
if (blockedEnemies.Count > 0)
{
	Enemy? validBlockedTarget = null;
	for (int i = 0; i < blockedEnemies.Count; i++)
	{
		// 1. 유효성 검사는 null 체크 & 활성화 모두에 대해 진행
		if (blockedEnemies[i] != null && blockedEnemies[i].gameObject.activeInHierarchy)
		{
			validBlockedTarget = blockedEnemies[i];
			break;
		}
	}

	// 2. 저지된 적이 지금 존재할 때만 타겟으로 선정하고 리턴
	// 없다면 아래 로직(공격 범위 내의 적 탐색)으로 넘어감
	if (validBlockedTarget != null)
	{
		CurrentTarget = validBlockedTarget;
		return;
	}
}
```

- 공격 범위 내의 적을 탐색하는 로직도 `activeInHierarchy` 조건을 추가해준다.
```cs
CurrentTarget = _enemiesInRange
	.Where(e => e != null && e.gameObject != null && e.gameObject.activeInHierarchy) // 조건 추가
	.OrderBy(E => E.Path.GetRemainingPathDistance()) 
	.FirstOrDefault(); 
```


#### 간헐적인 버그를 다루는 방법
>[!note]
>Claude에게 질문 : 가끔씩 발생해서 재현성이 떨어지는 현상을 어떻게 고쳐야 하는가?
>1. 로그 심기 : 버그가 발생했는지 여부를 명확히 기록하기
>2. 버그를 발생시키는 테스트 환경 만들기
>	- 이건 정확한 원인을 모르기 때문에 테스트 환경을 만든다는 것도 약간 어폐가 있다고 생각함
>3. 수정 후 검증은 부재 증명이 아니라 방어 코드로
>	- 물리적으로 버그가 발생할 수 없는 코드 구조를 만들라는 것
>4. 핵심 상태에서 Assertion 걸기

> 4번 예시
```cs
// 공격 수행 직전에 항상 타겟 유효성 검증
Debug.Assert(
	target != null && target.gameObject.activeInHierarchy,
	$"[AttackController] 유효하지 않은 타겟에 공격 시도: {target?.name}"
);
```

> 여기선 1번 작업만 수행했음 - `activeInHierarchy`가 `false`인데 메서드가 동작하는 경우에 로그를 발생하도록 처리



# 이전 일지
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
