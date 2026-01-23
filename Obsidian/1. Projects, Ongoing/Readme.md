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

>[!note]
>
>  **"흐름"으로 설명하는 것 외에도 "원리"도 곁들이면 좋다**
> - 지금은 "어떤 이슈가 발생해서 어떻게 해결했더라"로 정리했다면
> - 추가로 **이슈가 '왜' 발생했는지**까지 정리해두면 나중에 볼 때 다시 헷갈리지 않을 수 있음
> 
> 모든 코드를 일일이 기억하는 건 불가능함! 나중에 봤을 때 '이러면 되겠다'는 감이 바로 잡히도록 메모해보자.

>[!plan]
>- (진행 중)`UnitEntity` 리팩토링 : `God Class` -> `퍼서드Facade` 패턴
>1. 스테이지 밸런싱, 버그 수정
>2. 경로 로직 관련 현재 간헐적으로 발생하는 이슈들 보이면 처리
>	- 경로가 막혔을 때 바리케이드를 제거하지 않고 그대로 지나가는 현상
>	- 경로가 바뀔 때 살짝 뒤로 빼는 현상. 
>		- 최단 경로를 찾기 때문에 뒤로 빼는 현상이 나오면 안됨.
>3. 소리 추가
>4. 프로젝트 정리글 작성

>[!issue]
> 간헐적인 이슈들 : 계속 발생하는 경우에는 수정하지만 아니라면 남겨둠
> - 적이 이미 사라졌는데 계속 해당 위치를 공격하는 현상
> - `ArcaneFieldSkill` : 스킬을 썼음에도 효과가 제대로 적용되지 않는 현상
> 	- 위치가 애매하게 걸쳐지는 경우가 있나? 의심은 있는데 상황을 재현하기 어려움
> - `Enemy`가 사라질 때 풀 태그의 키가 없다는 오류 
> - `Barricade` 배치 시 가끔 배치되지 않음
> - `Operator` 체력이 다했을 때 사망처리가 되지 않았는데도 적이 지나가는 현상

## 현재 이슈
- `Enemy`의 컨트롤러 구현하기
- `PathNavigator`의 구현 관련 
	- `Enemy`나 `PathIndicator`에 구현된 필드가 너무 많아서 굳이 서브 컴포넌트로 빼둔 의미를 잘 못 느끼겠음
- `ArcaneFieldSkill`의 타격 이펙트에 노란색도 추가해볼까?
- `UnitEntity.ExecuteSkillSequence`을 `SkillController`에 통합하기
	- `BossSkillController`에서 사용 중. 보스가 스킬을 시전할 때 잠깐 멈추는 효과로 구현하고 있다
- `PathIndicator`는 최초 `pathData`가 아니라 **생성된 다음 경로를 재생성하는 것으로 보임**



## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.

## 260123

>[!done]
>- `Enemy` 클래스 분리 관련 테스트
>	- 생성되었으나 이동하지 않는 현상
>	- 게임 끝났는데 `Spawn` 동작하는 현상
>	- 공격 이펙트 실행되지 않는 현상
>	- 저지되지 않는 현상
### `Enemy` 클래스 분리
- 일단 그저께까지 구현한 것부터 테스트.

- [x] `Enemy`가 스폰되지만 이동하지 않음
	- `AttackController.OnUpdate` 정리
		1) 이동 조건 체크하는 로직이 공격을 판단하는 `AttackController`에 있었음 : 분리
		2) `OnUpdate`의 `return` 조건을 상단에서 아래처럼 판단하기로 바꿨는데 정작 `OnUpdate`문 자체의 `return`이 그대로였음
			- `true` : 공격을 했을 경우
			- `false` : 공격을 하지 않았을 경우
- [x] 게임이 끝났는데 `EnemySpawner.Spawn()`이 동작
	- `Spawn()`의 동작 조건을 `GameState.Battle`로만 한정

- [x] (디)버프 이펙트 사라지는 과정에서의 오류
```
Cannot set the parent of the GameObject 'StunBuffVFX(Clone)' while activating or deactivating the parent GameObject 'Enemy_Normal_Fast(Clone)'.
```
> `OnDisable`이 실행되고 있는 도중에 자식 오브젝트의 `Hierarchy`에서의 이동이 있어서 발생하는 오류
- VFX 오브젝트를 풀로 반환하는 과정에서 발생하는 오류였다.
- 사망 과정에서 발생하는 이벤트는 `OnDeathStarted`와 `OnDisabled` 2개를 구현해뒀다. 전자는 체력이 0이 된 순간 발생하고 후자는 `OnDisable`에 직접 들어 있음
- `OnDisabled`에서 `OnDeathStarted`로 바꿔봄.

> 일단 같은 내용의 오류가 발생하지 않아서 이렇게 유지해둠

- [x] Enemy 공격 이펙트 실행 안됨
####  저지가 안됨
- 어떻게 수정하느냐에 따라서 이슈가 왔다갔다 했는데, "저지당했을 때 적이 멈추지 않았다"로 요약할 수 있을 듯

1. `EnemyAttackController`의 저지당했을 때의 `OnUpdate` 로직 수정
```cs
// 저지당함 - 근거리 공격
if (_blockingOperator != null)
{
	if (CanAttack())
	{
		PerformMeleeAttack(CurrentTarget!, _owner.AttackPower);
	}

	// 저지 당한 상태라면 공격 여부에 관계 없이 추가 동작을 막기 위해 true 반환
	return true; 
}
else 
{
	// ...
}
```
> - 기존엔 `return true`문이 `PerformMeleeAttack` 내부에 있었는데 "저지 당한 상태에서는 이동하지 않는다"여야 하므로 항상 true를 반환하도록 수정함

2. `UnitOverlapSolver` 관련
	- 저지가 동작할 때 **`Enemy`가 겹치는 현상을 방지하기 위해 구현**했던 요소
	- 그런데 부자연스러운 경우가 좀 많다. 
		- 이미 저지당한 적 위치에 새로운 적이 들어갈 때 그대로 밀고 들어감
		- 저지당했을 때에만 동작하면서 기존 적을 밀리지 않게 구현했을 경우, 콜라이더의 동작 시점이 가장자리에 위치했을 때이기 때문에 `저지함 <-> 하지 않음`을 오가면서 버벅이는 현상
	- 이 부분은 아래처럼 작업해봄
		1) 저지당하는 순간 Enemy의 위치가 살짝 바뀌도록 동작
		2) `UnitOverlapSolver`의 기능은 살려놓되 오퍼레이터 배치 시에 밀려나는 기능은 살려놓음, `Enemy` 간의 충돌 기능은 꺼둠

> 일단 이 정도로 구현

```cs
// OpBlockController.TryBlockNextEnemy()

// 이 적을 저지할 수 있는지 확인
if (CanBlockEnemy(candidateEnemy))
{
	BlockEnemy(candidateEnemy);
	candidateEnemy.UpdateBlockingOperator(_owner);

	// 저지 중인 적이 2명 이상일 때, 2번째 적부터는 약간의 사이드 이동(완전히 겹치는 현상 방지)
	if (blockedEnemies.Count >= 2)
	{
		candidateEnemy.SmoothAvoidance(_owner);
	}
}
```

```cs
// Enemy 컨테이너에 추가
public void SmoothAvoidance(Operator op)
{
	// 저지 시에만 동작
	if (op != null)
	{
		// 겹쳤을 때를 고려한 위치 이동
		if (_adjustmentCoroutine != null)
		{
			StopCoroutine(_adjustmentCoroutine);
		}

		_adjustmentCoroutine = StartCoroutine(SmoothAvoidanceCoroutine(op.transform.position));
	}
}

// Enemy 끼리 겹쳤을 때 살짝의 이동
protected IEnumerator SmoothAvoidanceCoroutine(Vector3 targetPos)
{
	float duration = 0.1f;
	float elapsed = 0f;   

	Vector3 startPos = transform.position;
	
	// 방향, 수직 벡터
	Vector3 direction = startPos - targetPos;
	direction.y = 0f;

	// 수직 벡터 : (x, z)에 수직인 벡터는 (-z, x) 또는 (z, -x)
	Vector3 perpendicular = new Vector3(-direction.z, 0, direction.x).normalized;

	// 최종 목적지
	float randomSide = UnityEngine.Random.Range(0, 2) == 0 ? 1f : -1f;
	float randomDistance = UnityEngine.Random.Range(0.03f, 0.05f);
	Vector3 targetOffset = perpendicular * randomSide * randomDistance;
	Vector3 finalDestination = startPos + targetOffset;

	// 시간 동안 부드럽게 이동
	while (elapsed < duration)
	{
		elapsed += Time.deltaTime;
		float t = elapsed / duration;

		transform.position = Vector3.Lerp(startPos, finalDestination, t);
		yield return null;
	}

	transform.position = finalDestination;
	_adjustmentCoroutine = null;
}
```



---
# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
