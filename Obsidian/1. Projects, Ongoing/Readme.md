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
- `PathNavigator`의 구현 관련 
	- `Enemy`나 `PathIndicator`에 구현된 필드가 너무 많아서 굳이 서브 컴포넌트로 빼둔 의미를 잘 못 느끼겠음
- `CombatController`
	- 적군 구현 필요
- `UnitEntity.ExecuteSkillSequence`을 `SkillController`에 통합하기
- 스탯 정리 : 공격 속도와 공격 쿨다운 간의 정리 필요

## 작업 내용 정리
- 여기에는 가장 최근에 작업한 것만 기록해두고, 이전의 내용은 파일로 만들어서 옮겨뒀다.
## 260115

>[!done]
>- 버그 수정
>	- `ArcaneField` 
>		- 스킬이 꺼질 때 적이 멈추는 현상
>		- 스킬 범위에 진입 / 이탈 시 적이 진행 반대 방향으로 날아가는 현상
>	- `DoubleShotSkill` 
>		- 스킬 아이콘 클릭 시 나타나는 스택 오버플로우
>	- 일부 스킬 사용 후 SP가 회복되지 않는 현상 
>		- `OpSkillController`의 종료 로직 이슈
>- 스탯 시스템 다시 정리 : `ArcaneField` 만지는 과정에서 정리했음
>	- **`Modifier`의 기준값 : 0**
>- 보류
>	- `Barricade` 배치 시 어떤 상황에서 배치되지 않는 현상
>	- `Operator`의 체력이 다했을 때 사망 처리 X인데 적이 지나가는 현상

### 버그 수정

#### <해결> `ArcaneField` :  스킬이 꺼질 때 적이 멈추는 현상
- 체크 1 : `Enemy.Update.MoveAlongPath` 동작 (갑자기 저지당한다든가 하는 이슈는 아님)

> [!think]
> - 하나씩 점검하려던 중, 번뜩인 게 하나 있다
> - `ArcaneField`는 `SlowBuff`가 포함되어 있음
> - **`SlowBuff`가 풀리는 과정에서 원상 복구가 제대로 안된 거 아닐까?** 
> - 이런 논리로 접근하면 **또다른 이슈인 범위에 진입/이탈 시에 반대로 날아가는 현상도 설명할 수 있을 것 같음**

- `SlowBuff`의 `modifier` 값이 얼마인가 봤다 : `0.3`임. (30% 감소시킨다는 의미)
- 그리고 `Buff`의 메서드들을 보면
```cs
    public void AddModifier(StatType type, float value)
    {
        float actualValue = value - 1.0f;
        
        // ...
    }

    public void RemoveModifier(StatType type, float value)
    {
        float actualValue = value - 1.0f;
        
        // ..
    }
```
> 들어온 계수값에서 1을 빼고 `modifier`에 추가하는 방식임

- `modifier`에서 값이 활용되는 방식은 **`base * (1 + modifier)`임**
- ~~그래서 모든 스탯은 1을 기준으로 생각해야 한다. 30% 증가라면 1.3, 30% 감소라면 0.7..~~

>[!question]
>- 일단 `StatModifier`를 어떻게 관리할지부터 다시 점검해보면 좋을 것 같다. `Modifier` 값은 **1을 기준으로 넣을 것인가? 아니면 0을 기준으로 할까?**

>[!answer]
>- **`0`을 기준으로 하는 게 일반적임**
>- 사용하는 쪽에서는 **저 시스템만 알고, 메서드를 호출할 때 이를 반영**하면 됨

- 예를 들어서 `SlowAmount`라는 값으로 이동 속도를 저하시키는 로직이 있다고 하자. 
- "이동 속도 저하 30%"라는 로직을 구현하겠다면, 0이 기준일 경우 `0.3`이라는 값을 넣으면 됨. 
- 이것의 실제 의미는 "원래 이동속도의 70%로 이동하게 만들겠다"이므로 
- `AddModifier()`에서 넣는 값은 `-slowAmount` 만 넣어주면 된다. 

1. **`BuffController`의 `actualValue` 부분은 전부 제거**
2. `SlowBuff`도 아래처럼 반영
```cs
	// 30%의 SlowAmount는 modifier에서 `-0.3`이 됨
	// modifier는 baseValue * (1 + modifier)이므로 0.7배가 된다
    public override void OnApply(UnitEntity owner, UnitEntity caster)
    {
        base.OnApply(owner, caster);
        owner.AddStatModifier(StatType.MovementSpeed, -slowAmount); 
    }

    public override void OnRemove()
    {
        // owner.SetMovementSpeed(originalSpeed);
        owner.RemoveStatModifier(StatType.MovementSpeed, -slowAmount);
        base.OnRemove();
    }
```

- 이러니까 딱 막히는 지점이 생겼음 : **공격 속도 부분**
	- 이거는 **별도로 빼둠 - 복잡할 듯**
	- 공격 속도 : 1초에 몇 번 때리냐(타수/초)
	- ..를 뒤집으면(초/타수) 1번 때리는 데 몇 초 걸리냐라는 개념이 됨
	- 그래서 **공격 속도와 쿨다운을 구분해서 생각**할 필요가 있고, 별도의 로직도 필요해보인다. 이건 일단 보류.

> 리팩토링하다가 발생한 이슈이긴 하지만
> 일일이 언제 테스트해보나 하고 막막한 순간에 번뜩하고 해결한 이슈라서 오랜만에 쾌감을 느꼈다. 

#### <해결> DoubleShotSkill 클릭 시 스택 오버플로우
- `DoubleShotBuff` 부분을 보면 아래처럼 구현됨
```cs
private IEnumerator PerformDoubleAttack(UnitEntity owner)
{
	if (owner is Operator op)
	{
		UnitEntity? target = op.CurrentTarget;
		if (target == null) yield break;

		float modifiedDamage = op.AttackPower * damageMultiplier;

		op.PerformAction(target, modifiedDamage);
		yield return new WaitForSeconds(delayBetweenShots);

		if (target != null && target.CurrentHealth >= 0)
		{
			op.PerformAction(target, modifiedDamage);
		}
	}
}
```

어제 `DualBlade`에서 고민한 문제와 정확히 동일한 문제다. `PerformAction - PerformChangedAction - PerformAction`이라는 무한히 순환하는 재귀 호출 형태임

1. `PerformAction -> PerformActualAction`으로 구현
2. 이 이슈랑은 관계 없지만 `modifiedDamage`도 아래처럼 수정 - 스탯 자체에 들어간 보정치 값에 스킬의 보정치를 곱함
```cs
float modifiedDamage = op.GetStat(StatType.AttackPower) * damageMultiplier;
```

#### <해결?> 1번 시전 후, SP가 회복되지 않는 스킬이 있음
- `OpSkillController`에서 스킬 실행이 다 종료되고 나서 실행되는 로직이 달랐다
	- 지속시간이 있으면 `CompleteActiveSkill` (내부에 `OnSkillEnd` 포함)
	- 지속시간이 없는 즉발이면 `OnSkillEnd`
- `OnSkillEnd()`을 `CompleteActiveSkill` 외부로 빼서 실행 메서드에서도 흐름을 볼 수 있게 하고, 양쪽 모두 `CompleteActiveSkill()`로 끝나게끔 수정
```cs
    private void ExecuteInstantSkill(ActiveSkill skill)
    {
        skill.OnActivated(_owner);
        skill.OnUpdate(_owner);
        skill.OnEnd(_owner);
        CompleteActiveSkill(skill);
    }
    
    private IEnumerator Co_HandleDurationSkill(ActiveSkill skill)
    {
        skill.OnActivated(_owner);

        float elapsed = 0f;
        float duration = skill.Duration; 

        while (elapsed < duration)
        {
            // 소유자 파괴 시 중단
            if (_owner == null)
            {
                CleanupSkill();
                yield break;
            }

            elapsed += Time.deltaTime;
            float progress = elapsed / duration;

            // UI용 - SP 시각적 감소
            CurrentSP = MaxSP * (1f - progress);

            // 스킬 틱 호출
            skill.OnUpdate(_owner);

            yield return null;
        }

        // 정상 종료
        skill.OnEnd(_owner);
        CompleteActiveSkill(skill);
    }    
```

- 이 문제가 발생하는 스킬에선 동일한 현상 발생 X

#### <간헐적, 보류> Barricade 배치 시 특정 상황에서 배치되지 않음
- 다시 테스트했을 때는 또 잘 된다. 보류.
#### <간헐적, 보류> Operator 체력이 없는데 안 죽고 그냥 지나감
- 얘도 잘 되는데? 흠..

#### <해결> DeployableActionUI - 스킬 버튼 관련
1. 커서 올릴 때 색깔 변하게 하기
2. 스킬 활성화 시에 주황색이 줄어들어야 함
	- **텍스쳐가 누락되면 `Image.filled` 효과가 아예 나지 않음**
	- 이전에 `square_sprite`을 지운 적이 있는데 그것 때문이었다. 로직 자체는 잘 동작하고 있음.

---
# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
