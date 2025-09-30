# 참고
- **옵시디언으로 봐야 제대로 보인다.**
	- **옵시디언으로 작성된 만큼 깃허브의 마크다운에서는 지원하지 않는 기능들이 있을 수 있다.** `[[]]`, 이미지 첨부 방식 등이 대표적.
- `[[]]` 링크는 `유니티/보관함`이나 `작업 일지/직접 작성/일지`에 대부분 있다.
---
# 작업 내용 : 짭명방

## 보고 참고할 내용
- [[오퍼레이터들 스탯 정리]]
- [[적 스탯 정리]]
- [[프로크리에이트로 텍스쳐 작업 시 유의할 점]]
## 사용 툴
- `Unity`
- UI 이미지 제작 : `Procreate`
- `VFX` 제작 사용 툴(강의 보면서 따라함)
	- `Krita`
	- `Blender`

## 하고 싶은데 못한 것
- 캐릭터 스프라이트 
	- 현재 Capsule로 Operator나 Enemy 등을 구현한 상태
	- 이걸 투명한 Quad로 바꾸고 그 위에 2D 스프라이트들을 구현하는 방식이 명방에서 쓰고 있는 방식으로 보임
	- 직접 하려면 배워서 할 수는 있겠지만 시간이 문제겠다. 
- 도전과제 구현
- 인게임 / 다른 패널에서 스킬 범위 보여주기 
	- 일관성 때문에 다른 곳에서도 구현을 해보고는 싶은데, 일단 보류.
	- 현재는 정예화 패널에서만 스킬 범위를 볼 수 있음.



---
# 작업 일지

## 짭명방 예정

### 작업 중

>[!wip]
>- 머티리얼 변경하면서 제대로 나타나지 않는 이펙트들 수정 중


>[!todo]
>- 보스 구현
>	- 쉴드 스킬
>		- 얘만 지금 VFX Graph로 남아 있음.
>- 남은 작업
>	- 스테이지 1-3 완성
>	- 남은 스테이지들 밸런스 수정
### 간헐적인 이슈
- 이슈가 있다고 느꼈는데 다시 테스트했을 때 재현이 안된 것들을 정리함

- 오퍼레이터 A를 배치할 때, 방향 설정 로직 중 오퍼레이터 B의 위치에서 마우스 커서를 떼면 배치되면서 해당 마우스 커서의 위치에 있는 오퍼레이터가 클릭되는 현상

# 250930 - 짭명방

## 이슈 수정

### 어제 적어둔 거

오퍼레이터 A를 배치할 때, **방향 설정 로직 중 오퍼레이터 B의 위치에서 마우스 커서를 떼면 배치되면서 해당 마우스 커서의 위치에 있는 오퍼레이터가 클릭되는 현상**이 있음

- 오늘 다시 테스트해봤는데 **괜찮았음** 
	- 간헐적으로 발생할 수도 있는 이슈라서 일단 기록은 해두고 다른 걸로 넘어감

### 오퍼레이터 UI 관련
- `InStageInfoUI`를 띄운 상태에서 해당 오퍼레이터가 아닌 다른 오퍼레이터가 죽었을 때 `InStageInfoUI` 클릭 상태가 초기화되는 현상
	- `DeployableManager.OnDeployableRemoved` 에 `HideDeployableInfo`가 들어가 있다. 이걸 수정해줘야 함.
```cs
// 배치된 요소가 제거되었을 때 동작함
public void OnDeployableRemoved(DeployableUnitEntity deployable)
{
	deployedItems.Remove(deployable);
	StageUIManager.Instance!.HideDeployableInfo();
	
	// ...
	
}
```

그러면 현재 제거된 `deployable`이랑 `StageUIManager`에서 현재 띄우고 있는 `deployable`이 같은지 다른지 점검하는 조건을 추가해야 함.

1. 현재 정보를 띄운 오퍼레이터가 제거되면 정보 패널이 사라져야 하고
2. 다른 정보를 띄웠을 때는 어떤 오퍼레이터가 제거되면 정보 패널이 유지되게 하기 위함

그러면 `StageUIManager`나 `InStageInfoPanel`에서 현재 띄우고 있는 오퍼레이터/배치가능요소의 정보를 갖고 있어야 함

> 여기서 고민했던 건 `InstageInfoPanel`에서 현재 띄우고 있는 정보를 갖느냐 아니면 `StageUIManager`에서 정보를 갖느냐였는데, SRP에 의하면 **항상 자기가 처리할 정보는 자기가 갖고, 매니저는 단순한 연결고리 역할만 해주면 됨.**
> `currentDeployableInfo`라는 걸 활용하고 있기 때문에, 이에 대한 게터 프로퍼티만 패널에 만들어주고, 스테이지 매니저에서도 저 스크립트에 대한 게터 프로퍼티를 만들어주고, `DeployableManager`에서 아래처럼 연결
```cs
if (deployable == StageUIManager.Instance!.InStageInfoPanel.CurrentDeployableInfo.deployedDeployable)
{
	StageUIManager.Instance!.HideDeployableInfo();
}
```
- 코드가 지저분해졌다. 더 개선할 수 있다. `StageUIManager`에서 `InstageInfoPanel`은 어떤 정보를 갖고 있는지 관리하는 메서드를 별도로 구현하고, 여기서는 해당 메서드만 요청하면 됨.

#### 조금 더 개선
- 요런 느낌. `DeployedOperator`를 구분한 이유는 기존 구현에 `Deployable`과 `Operator`를 구분했기 때문에 이렇게 해놨음. 

- `InStageInfoPanel`
```cs
    public bool IsCurrentlyDisplaying(DeployableUnitEntity entity)
    {
        // 패널이 비활성화거나, 표시중인 정보가 없거나, 표시 중인 배치된 유닛이 없다면 false
        if (!gameObject.activeSelf ||
            CurrentDeployableInfo == null ||
            CurrentDeployableInfo.deployedDeployable == null ||
            CurrentDeployableInfo.deployedOperator == null)
        {
            return false;
        }

        if (entity is Operator op)
        {
            return CurrentDeployableInfo.deployedOperator == op;
        }

        return CurrentDeployableInfo.deployedDeployable == entity;
    }
```

- `StageUIManager`
```cs
public void HideInfoPanelIfDisplaying(DeployableUnitEntity entity)
{
	if (inStageInfoPanelScript.IsCurrentlyDisplaying(entity))
	{
		HideDeployableInfo();
	}
}
```

- `DeployableManager`
```cs
	// 현재 나타나고 있는 InstageInfoPanel의 정보가 제거된 deployable이라면 정보를 가림
	StageUIManager.Instance.HideInfoPanelIfDisplaying(deployable);
```

### InStageInfoPanel은 유지되는데 다른 정보가 사라짐
- 예를 들면 클릭 가능한 요소(스킬 버튼, 퇴각 버튼) 같은 게 사라진다. 일단 오늘은 여기까지만 작업해둠

---

# 250929 - 짭명방
- 9월이면 끝낼 수 있을 줄 알았는데 자만함과 나태함과 자신에 대한 과대평가와 이슈에 대한 과소평가 등 이것저것이 겹치고 있다. 후다닥 달리자.
- 아이패드 프로도 결국 충전 단자가 가버렸다. 수리를 맡겼는데 배터리 교체까지 포함하면 최소 30은 나올 듯.. 
	- 이라고 생각했는데 배터리 효율이 89%라고 한다. ? 거의 5년 넘게 썼는데? 18만원 선에서 수리받았다. 휴..

>[!done]
>1. `Enemy`와 `Operator`가 겹쳤을 때 `Operator` 클릭이 동작하지 않는 이슈 수정
>	- `RaycastAll` 사용 + 모든 `Operator`에 기존보다 더 큰 콜라이더 넣음
>	- 콜라이더 위치와 크기 수정 : 가능한 타일에 가깝게(z축 위아래 겹쳤을 때 위 오퍼레이터를 클릭하기 힘든 현상 때문)
>2. 기타 이슈 수정
>	- `statisticItem`의 아이콘 색상에 오퍼레이터의 색상 반영되지 않는 현상 수정
>	- 제대로 나타나지 않는 머티리얼 수정
>		- 스킬 사용 가능 아이콘
>		- `VFX_Arcane_Hit`
>		- `SlashSkill`
>		- `Projectile_Molotov`
>		- `SkillRange_Boundary`
>		- `MeteorSkill`
>			- 영역 벽 부분
>			- 메테오 파티클
>		- `StunEffect`
>		- `Projectile_Missile`
## 레이캐스트 관련
- 오퍼레이터를 클릭하려고 하는데 다른 유닛이 겹친 상황인 경우, 오퍼레이터 클릭이 동작하지 않는 현상이 있음
### 접근 시도
- 현재 코드가 이렇다. `UI` 판별 부분은 제외했음.
```cs
private void ProcessClickMapObject()
{
	// 1. 배치 중 드래깅 혹은 방향 선택 상태라면 클릭 처리 중단
	// 배치 상황은 DeployableManager에서 처리한다. 여기서는 다른 오브젝트의 클릭 동작을 막기 위해 남겨둠.
	if (DeployableManager.Instance!.IsSelectingDirection)
	{
		Debug.Log("HandleClick : 배치 중 드래깅 혹은 방향 선택 상태 - 클릭 처리 중단");
		return;
	}

	// 2. 3D 오브젝트 클릭 처리
	Ray ray = mainCamera.ScreenPointToRay(Input.mousePosition);
	if (Physics.Raycast(ray, out RaycastHit clickableHit, Mathf.Infinity, clickableLayerMask))
	{
		HandleObjectClick(clickableHit);
	}
	else
	{
		HandleEmptySpaceClick();
	}
}
```
> `Enemy`들에 대해 저 `clickableLayerMask`이 할당되어 있는 게 문제일까?

- 옵션을 보면 `Deployable, UI_Overlay, UI_WorldSpace` 이렇게만 `ClickableLayerMask`에 할당되어 있다. `Enemy`는 `Default` 레이어에 해당하므로 여기에 없는 듯.
- 클릭되어야 하는 오브젝트들은 `Deployable`에 할당되어 있음.


1. `RayCastAll`으로 수정 - 지금의 코드는 가장 위에 있는 것만 잡아서 보여주는 방식임
2. `Operator`의 콜라이더 관련
	- 기존 `Enemy`와의 충돌을 감지하기 위한 콜라이더보다 더 넓은 범위의 클릭을 감지하게 하고 싶음
	- 클릭 감지를 위한 콜라이더를 별도로 추가해야 함. 일단 모든 오퍼레이터에 `BoxCollider`만 자식 오브젝트로 추가해봄.
		- **이것만으로 잘 구현되는 것 같다.** 이 현상이 주로 오퍼레이터가 위에 있고 적이 아래에서 저지당한 상태일 때, 적의 위치를 클릭하면 아무 반응이 없는 것 때문에 신경쓰였던 부분이었다. 


### 이슈 수정하다가 추가 수정
- `BoxCollider`의 스케일을 `1, 1, 1`로 구현했는데, 이렇게 할 경우 **위 타일과 아래 타일에 오퍼레이터가 겹쳐 있으면 위 타일의 오퍼레이터를 클릭하기 매우 힘든 현상**이 있음
- 그래서 가능한 타일에 가까운 위치로 설정함 
	- 중심 위치는 `0, -0.4, 0`
	- 스케일은 `1, 0.1, 0`


## 기타 이슈 수정

### 결과 카드 아이콘 색상
- 결과 카드에서 오퍼레이터 색상 반영되는 부분이 모두 흰색으로 나타남 
	- `250926`에 바꿨던 부분 때문에 발생한 현상.

```cs
// 이전
operatorIcon.color = opData.prefab.GetComponentInChildren<Renderer>().sharedMaterial.color;

// 수정
operatorIcon.color = opData.PrimaryColor;
```

### 제대로 나타나지 않는 요소들 수정
- `SkillAvailableIcon`
	- `OperatorUI`에 머티리얼 할당이 해제되어 있었다. 텍스쳐를 지운 것도 아니고 머티리얼을 지운 것도 아닌데 왜 없어져 있었는지는 모르겠음.
- `Arcane_Hit`
	- `Trail02` 수정
	- `Trail` 셰이더 자체도 `Vertex Color` 반영되도록 수정
- `SlashSkill`
	- ...인 줄 알았는데 `Legacy`로 할당되어 있던 부분이었다. 새로 수정한 `SlashSkill`은 잘 나타남.
- `Projectile_Molotov`
	- `Trail` 너비가 너무 넓어서 줄였음
- `SkillRange_Boundary`
	- **셰이더를 삭제했었다. 휴지통을 뒤져보니 있어서 다행히 구조해옴**
	- 유니티에서 삭제했던 파일들 되돌렸을 때 원본 참조 유지하는 거 보면 신기하다. `Meta` 파일의 역할인건가?
- `MeteorSkill`
	- 스킬 시전시 한 번 들썩이는 모양 -> 수정 완료(가장 기본적인 그래디언트 텍스쳐 사용)
	- 메테오 파티클 : 계속 `Vertical Billboard`로 적용되는 현상이 있음 -> `Stretched`여야 하는데 반영이 안 된다. 왜 그런지는 휴식 겸 수리점 다시 갔다 와서 수정함
		- 반영 완료. 디폴트 속도를 `Velocity over Lifetime`에서 `-y`로 아주 작은 값을 설정해서 파티클의 방향을 설정했다. 기존엔 어떻게 했는지 기억 안 남.
- `StunEffect`
	- 파티클은 `TextureSheet`, 트레일은 `Trail02_White` 사용함
	- 이거는 가끔 머티리얼이 마젠타 색으로 나타날 때가 있었음. 그거까지 테스트해봄.
		- 이 문제는 계속 발생하진 않는 듯?
- `Projectile_Missile`
	- `Spark01_AB` : 텍스쳐 시트 기반으로 변경
	- `Smoke`도 `AnimationSheet` 기반으로 변경


# 250926 - 짭명방

> [!summary]
> 1. `Add`, `AB` 머티리얼 외의 다른 머티리얼은 별도로 놓고, 텍스쳐도 `SingleTexture`로 따로 보관함
> 	- 별도의 용도로 사용되지 않는 `TextureSheet`의 텍스쳐들은 로컬에 백업해뒀으니 삭제
> 2. 각 유닛 메쉬에 사용하던 머티리얼들은 다 비슷한데 색만 다르다. 이게 드로우 콜을 유발하기 때문에 별도의 셰이더를 사용하는 `Ghost`를 제외하고 나머지 `Enemy`나 `Operator`, `Barricade`에 사용하는 머티리얼들은 모두 `BasicMesh01` 머티리얼로 통합했음. `URP/Lit`을 사용한다.
> 	- `DeployableBox`의 박스 색은 모델의 색을 가져왔기 때문에 여기에 영향이 갔다. 그것도 `DeployableData`에서 가져오도록 수정함. 
## 머티리얼 정리

### 유닛 머티리얼 정리
- 유닛에 사용하는 기본 머티리얼들은 모두 `URP/Lit`을 쓰는데, 색이 다르다는 이유로 모두 별도의 머티리얼을 사용하고 있음
- 그래서 하나의 머티리얼로 통합하고 초기화 시점에 머티리얼에 색을 할당하는 방식으로 구현하려고 함
	- `UnitEntity`에 구현하고 색을 저장하는 곳이 자식 클래스마다 다르기 떄문에(다른 `Data` 타입을 쓰기 때문에) `Start` 시점에 각각의 `Data`에서 색을 가져옴

단 `Ghost`는 투명해졌다가 나타났다가 하는 별도의 셰이더를 쓰므로 제외

- **정리 완료. 기존 머티리얼들은 삭제했음.**
- 이 과정에서 `DeployableBox`에 색깔 할당하는 로직에 영향이 갔다. 모델 머티리얼의 색을 가져와서 `DeployableBox`의 색에 할당했기 때문인데, 이걸 `deployableInfo.operatorData`에서 가져오도록 수정했다.
	- 이거는 `ScriptableObject`가 `Operator`에서 초기화되는 시점보다 빨랐던 것 때문에 그거 구현한다고 예전에 남겼던 게 있어서 이번에 쉽게 넣을 수 있었음. 
	- `Operator`에서 `OperatorData`를 할당하는 시점과는 별도로, 스테이지 단위에서 `Deployable`에 대한 `DeployableData`를 별도로 관리하는 개념으로 구현했던 게 있었다. 

# 250924 - 짭명방

## 아틀라스(텍스쳐 시트) 작업
- 텍스쳐를 정리했으니까 이제 아틀라스(텍스쳐 시트)로 만들어서 사용해본다. 
- 자동화하는 방법도 있지만 일단 배운 대로 써먹어봄. 
- 프로크리에이트로 옮겨서 정리하는데, `png` 파일은 투명한 배경이 그대로 1024 x 1024로 들어가는 게 아니라 **투명한 부분은 데이터가 없는 것처럼 취급해서 이미지의 비율이 캔버스의 비율을 유지하지 않는다.** 
	- 그렇기 때문에 **실제로 이펙트를 넣은 상황과 다르게 동작할 수 있음.** 아틀라스에 넣는 과정에서 비율이 달라졌기 때문이다. 그래서 이펙트들에 대한 전반적인 재점검이 필요함

### 참고) 아틀라스에 넣으면 안되는 텍스쳐
- `Tiling And Offset`이 적용되는 텍스쳐
	- **UV 기반으로 계산하기 때문에 한 이미지는 하나의 텍스쳐에 관한 정보만을 갖고 있어야 함**
	- 물론 커스텀 셰이더를 구현하면 특정 UV 내에서만 반복되게끔 구현하는 것도 가능함
	- 일단은 이렇게 알고 있자.

### Sprite Sheet Animation에서 주의할 점
- 1개의 텍스쳐를 골라서 띄울 때 Frame Over Time 값은 0으로 설정해야 함
	- 1로 설정하면 의도한 값(인덱스) 다음의 텍스쳐를 지정하는 느낌으로 나타나고 있음

### 기본적으로 Add랑 AB 위주로 수정
- 나머지 셰이더들을 사용한 경우는 어떻게 처리할지 생각 좀 해야 할 듯
- 보로노이로 구현한 게 있고 노이즈로 구현한 게 있고 이것저것 다양하다.

```
보류한 이펙트 목록
- Attack에 있는 평타의 Slash 이펙트(Voronoi)
- BuffOn에 있는 Spiral 이펙트 (Scroll)
- 보스 스킬 시전 이펙트 - Trail (Scroll)
- Arrow Hit 이펙트의 Corona (Scroll)
- Hit_Explode 이펙트의 Distortion (UV Distortion)
- Hit_Fire 이펙트의 Corona (Mat: Corona01_AddScrollRadial)
- Hit_Heal 이펙트의 Spiral (Mat: Spira02_AddScroll)
- Projectile_Orb_Heal 이펙트의 Sphere (Mat : Spehre01_AddScroll)
- Explosion 스킬의 Explosion 부분 (Mat : NukeExplosion01)
- Explosion의 FallingParticle (Mat: MeteorEffect01)
- Explosion의 CrackedGround (Mat: Crack02)
	  - Parallax Occlusion Mapping 때문임
- Trail Renderer을 사용한 Trail들 : 텍스쳐 시트에 넣지 않음
- SlashThrough의 Slash 효과들
	- Slash(Outer, Inner, AB) : 커스텀 채널을 이용해 UV값에 따라 나타나는 효과
	- Aura : Dissolve 효과


기타 참고할 사항
- VFX_Cast_BossRangedSkill에 사용한 30번째 인덱스의 경우 일부러 가로를 256으로 맞추지 않았는데 빈 공간이 보이는 이슈가 있음 : 몇 픽셀인데 생각보다 눈에 띈다.
  
```

### 작업 중 텍스쳐 시트 2개로 분리
- `Animation`에 해당하는 텍스쳐들은 별도 파일로 분리함

> 일단 오늘은 여기까지
> 다른 셰이더들을 어떻게 적용할지 생각해봐야겠음


![[ForObsidian_VFXTextureSheet01 (1).png]]

![[ForObsidian_AnimationTextureSheet01 (1).png]]




# 250923 - 짭명방

## 텍스쳐, 머티리얼 정리
- 사용하지 않는 텍스쳐와 머티리얼을 정리하고 사용 중인 텍스쳐의 이름을 다듬을 예정
- 머티리얼에서 사용하지 않는 텍스쳐는 사용하지 않는 텍스쳐로 봐도 되겠지? 
- 문제라면 사용하지 않는 머티리얼의 경우인데.. `Find Reference In Projects`을 돌려봤는데 잘 동작하지 않는 듯.

- 텍스쳐 
```
baldo01-3x3
Beam01
Circle01
Cloud01
Cloud03_3x3
Corona01
Corona02
Crack01
Defender128
ExplosionDebris01
Fire01_3x3
GroundTexture01
GroundCrack02_Albedo
GroundCrack03
GroundCrack03_Normal
Gradient00
GradientDBD01
Gradient_BtmUp01
Glow01
Impact01
Impact02
Impact03
Lightning01
Lightning01-3x3
Lightning01_Seamless
medical_128
Missile01
Spark01
swords_128
SlashLeft01-2x2
SlashMesh01
Slash01_Albedo
SlashAura01-1
SlashThrough01
Smash01
Spiral01
Thunder01
Trail02
Trail03
Triangle01
Wall01
Wall03
```

여기 없는 것들을 정리함.

- 머티리얼
```
AttackIcon01_Add
Baldo01_Add
Beam01_Add
Beam01_Brighter
Beam01_AB
Crack01_Add
Crack01_AB
Crack02
Cloud01_Add
Cloud01_AB
Cloud02_Add
Circle01_Add
Circle01_AB
CrackedBall01
Corona01_Add
Corona01_AB
Corona01_AddScrollRadial
CylinderWall01_Add
CylinderWall02_Add
Debris01
Distortion01
DefenceIcon01_Add
ExplosionDebris01_AB
Fire01_Add
GroundTexture01_AB
HealIcon01_Add
Impact01_Add
Impact01_AB
Impact02_Add
Impact02_AB
Impact03_Add
Impact03_AB
Kira02_Add
Molotov01
MeteorEffect01
Missile01
NukeExplosion01
NoTexture_Add
Orb01
Spark01_Add
Spark01_AB
Spark01_AddScroll
Sphere01_AddScroll
Spiral01_AddScroll
Spiral02_AddScroll
Slash01_AddScrollRadial
Slash01_AB
Slash01_2x2_Add
Slash01_2x2_AB
Slash02_Add
Slash02_AB
SlashAura01
Smash01_Add
Smash01_AB
Snowflake01_Add
SmokeTrail01
Thunder01_Add
Thunder01_AB
Trail02_Trail
Trail03_Trail
Trail04_Trail
Trail06_Trail
Wall02_WallAnimation
Wall04_Add
```

- 새로 텍스쳐 만들면서 이미지나 머티리얼이 이중으로 관리되고 있는데 이런 부분들 수정해나가야 할 듯. 뱅가드 2스킬에 문제 발생했으니 참고.

# 250922 - 짭명방

## VFX : SlashThroughSkill의 이펙트 수정
- `Cast` : 기존에는 깜빡이다가 **터지면서** 지나가는 이펙트였는데, **퍼지는 게 아니라 모이는 이펙트도 만들어봄.** 이쪽이 더 어울릴 것 같다.
	- 더 정확히는 서서히 커지다가 확 모이는 느낌의 이펙트.

![[VFX_BossSkill_SlashThrough_CastImprove.gif]]
> 좀 더 보기 나은 듯. 분명 순간이동인데 긋고 지나가는 느낌이 있으니 성공적이긴 하다.

젠존제의 미야비를 생각해보면 발도를 시작해서 지나가고 나서까지 캐릭터에 빛나는 효과(십자가)가 있다. 그것까지 추가해봄. 

기존에 `EffectTextureMaker`로 얻은 게 있는데 윤곽선이 보여서 좀 부자연스러운 면이 있다. `MaterialMaker`로 다시 만들어서 넣어 봄.
-> 아니다. `Glow04`가 있었음. 하... **텍스쳐랑 머티리얼 정리도 좀 필요해보임.**

- 원래는 머리 쯤에 띄워놓아서 넘어가고 나서도 잠깐 유지되는 방식으로 구현했는데, 뭔가 어색했다. 그래서 보스의 자식 오브젝트로 만든 다음 바닥에 이펙트가 남도록 구현함.





# 250919 - 짭명방

>[!done]
>- `Boss_SlashThroughSkill`
>	- 메인 이펙트 구현 완
>		- 솔직히 만족스럽진 않은데 더 개선할 방법도 딱히 안 떠오름 ㅠ
 ![[VFX_BossSkill_SlashThrough_MaybeFin.gif]]

## Boss_SlashThroughSkill 이펙트 구현
- 관통하고 지나간다는 점에서 젠레스에 있는 미야비 이펙트를 참고로 하긴 했는데 차이는 많이 나지만 ㅋㅋ 바닥에 그인 칼날 위로 세운 `Quad`에는 검격 텍스쳐에 `Dissolve` 셰이더를 붙였다. 
	- UV에 노이즈를 추가하는 경우 `Repeat` 텍스쳐는 V값이 높은 위치에서 다른 텍스쳐가 나타난다. 그에 대한 마스크도 함께 구현해서 들어가 있음.
	- `Color over lifetime`을 함께 넣어서 크게 티가 나지 않는 느낌이지만 커스텀 데이터  + `Alpha Threshold`를 이용한 `Dissolve` 이펙트가 들어가 있다.
	- 마스크는 원래는 `Voronoi`를 써보려고 했는데 그것보다는 `Gradient Noise`가 더 느낌이 좋아서 그걸로 구현했음. 

### 오늘의 작업 관련 후기
- 미야비 이펙트를 참고했기에, 거의 보이지 않는 바닥 텍스쳐는 사실 얼어 있는 바닥 같은 느낌으로 구현해보고 싶었음. ㅠㅠ 근데 실력이 모자라서 그런 느낌은 못 냈음
- 프로크리에이트가 업데이트되면서 브러쉬 라이브러리를 아예 통째로 새로 냈다. 그 중에 느낌이 괜찮은 텍스쳐로 대충 바닥 이미지 만들어서 넣었음. 위 gif에선 생각보다 티가 안 난다.
- 그리고 왜인지 모르겠지만 애플펜슬이 다시 충전되기 시작했다. ?? 
- 오늘은 필기가 별로 없다. 이미지 작업할 때는 말이 없어지나봄. 엌ㅋㅋ



# 250918 - 짭명방
>[!done]
> - `Boss_SlashThroughSkill`
> 	- `VFX` 구현 : `Main`(뚫고 지나가는 효과)
> 	- **추가로 구현하고 싶은 이펙트**
> 		- **이전 위치 -> 나중 위치로 순간적으로 이동하는 이펙트**
> 		- **바닥이 뭔가 허전해서 다른 이펙트도 추가하고 싶음**
> - 기타 수정 사항
> 	- 보스 통과 시 스테이지 종료 : 적이 도착지점에 들어갔을 때 깎이는 라이프 포인트 필드 구현
> 	- EnemyData 필드 접근자 수정
> 	- `AreaVFXController`에서 근처에 다른 영역이 있는데 경계선 이펙트가 나타나는 현상 수정

## SlashThroughSkill 메인 이펙트 구현
- 이런저런 시도를 해보려고 함. 어느 쪽이든 텍스쳐를 어떻게 구현할지는 모르겠다.
	1. TrailRenderer
		- 이걸로 구현할 경우 단순히 스킬 발동 시 오브젝트를 가져와서 활성화하고 이동한 후에는 비활성화하는 방식으로 할 수 있을 것 같음
		- 타이밍이 문제일 듯? 스크립트로 봤을 때 한 프레임만에 끝나서 보이지 않을지도 모르겠다는 생각이 듦.
	2. ParticleSystem
		- 이걸로 구현하면 역시 **방향성을 어떻게 구현하느냐**가 문제일 듯

- 일단 파티클 시스템으로 한다고 생각하고 머티리얼 메이커에서 `Slash` 텍스쳐를 제대로 하나 만들었다. 기존에는 `Beam`을 사용해서 늘이는 식으로 썼음.

![[VFX_BossSkill_SlashThrough_MidFin.gif]]
- 보스 이펙트는 바닥 영역 이펙트랑 통과하는 이펙트는 만들 거임 - 근데 일단 보류
- 저 방향 같은 경우 방향 설정 로직에 의해 오퍼레이터가 어떤 이동 노드 위에 있을 때는 저런 식으로 진행됨(일반적으로 8방향일 듯)
- 노드 위에 있지 않을 때는 단순히 현재 위치에서 다음 노드로 향하는 포지션을 반환함

- 추가) 커스텀 데이터 구현으로 바닥에 생기는 파티클은 왼쪽 -> 오른쪽으로 UV가 나타나게끔 구현
![[Pasted image 20250918224910.png]]
- 원래 `Step`으로 구현했는데 어색해보이는 게 있음(경계선이 뚜렷하기 때문에 스르륵 나타나는 듯해 보임)
- 그래서 `SmoothStep`으로 바꾸고 `Edge2`를 저렇게 지정해주고 In을 0->1로 조절해줘봤더니 괜찮아 보여서 그렇게 구현해놨다. 사실 의도적으로 속도를 느리게 하지 않은 이상 웬만해선 보이지 않는 요소이긴 함

- 근데 검은색 파티클이 안보이기 시작했다.
	- 이거 그냥 `Custom Data` 안 켜놔서 안 보이는 거였음 ㅋㅋㅋㅋㅋㅋ **`Custom Data` 모듈이 꺼져 있다면 벡터 `0, 0`이 전달된다고 함.**

## 기타 수정 사항
### 보스 통과 시 스테이지 종료되도록 수정
- `EnemyData`에 목적지에 도달했을 때 플레이어에게 주는 대미지 추가
- 대부분 1이고 보스만 3임
- `StageManager`의 관련 로직들 추가

### EnemyData의 필드 접근자들 수정
- 다 `Public`으로 되어 있던 걸 `[SerializedField] protected`으로 수정하고 프로퍼티로 접근자 생성

### 버그 
- `ArcaneFieldSkill`
	- 영역의 테두리에만 걸쳐져야 하는데 타일의 경계선에 나타나는 이펙트들이 있음
	- 수정 완료

# 250917 - 짭명방

>[!done]
>- `Boss_SlashThroughSkill`
>	- VFX 구현 : `Cast`, `Hit`
>	- 버그 수정 
>		- 지나가고 나서 평타 한 대 더 치는 문제 수정 완료
## SlashThroughSkill 이펙트 구현
- 베고 지나가는 느낌을 어떻게 구현해야 할지 아이디어가 도무지 안 잡힘
	- `Trail Renderer`을 이용해보려고 하는데 지금 애플펜슬이 고장나서 프로크리에이트를 사용할 수 없는 상태임
	- `Krita`로도 시도해봤는데 영 마음에 안 들고, AI로 이미지 뽑아봤는데도 역시 베고 지나가는 느낌까지는 잘 모르겠다.

여기서는 이렇게만 써놨는데 시간 오래 잡아먹었음 ㅠㅠ 이래서 레퍼런스랑 스케치가 중요한 건가

- 그래서 `Cast`랑 `Hit`부터 구현하겠음


- 이번엔 여러 색을 같이 써보는 식으로 시도해봤다. 녹색 - 주황색이 떠올라서 넣어봄.
### Cast 구현 & 한 대 더 때리는 이슈 수정
![[VFX_BossSkill_SlashThrough_Cast.gif]]
- 어제 **저 지나가서 한 대 때리는 이슈**가 있었다. 이 gif는 실제로는 스킬 대미지 + 평타 대미지 2번이 들어가고 있는 상황임.
- 스킬을 사용한 후에 **`Operator`는 자신이 저지하고 있는 적 리스트에서 `Enemy`를 제거시키고, `Enemy`는 자신을 저지하고 있는 `BlockingOperator`를 초기화**해주는 방식으로 평타 동작을 막았다.

> 근데 지나가고 나서 뒤에서 한 대 치는 저런 느낌도 나쁘진 않은 듯
> 물론 바로 뒤에 오퍼레이터가 배치되어 있다면 또 다른 얘기지만..

### Hit도 구현
![[VFX_BossSkill_SlashThrough_CastAndHit.gif]]
- `0.2`초는 너무 빨리 지나가는 감이 있어서 시간을 좀 키워봤는데 때리는 느낌이 덜 난다. 일단 이렇게 놓고 아쉽다는 생각이 들면 더 수정하겠음
- 전조 이펙트가 있으니까 뭔가 큰 이펙트가 나와야 할 것 같은데 소소한 이펙트가 나오니까 좀 그런 느낌은 있음. 근데 이건 메인 이펙트를 구현한 다음에 봐야겠다.

메인 이펙트가 어떤 방식으로 동작해야 하는지 등은 생각해봐야 함. 내일 진행함.

## 잡설
- 요즘 아이패드랑 애플펜슬이 오락가락했다. 아이패드는 충전이 안 되는 이슈(근데 어떤 단자에 한해서는 충전이 되고 있어서 버티는 중), 애플펜슬은 스스로 연결이 끊기는 이슈.
- 근데 **애플펜슬이 아예 가버린 것 같다.** 부착해도 Apple Pencil 이 연결되는 인터페이스가 나타나긴 하는데 연결이 되면 시작되는 충전이 아예 진행되지 않음. 패드를 껐다 켜도 마찬가지고 미니에 연결해도 비슷해서 패드 이슈가 아니라 펜슬 이슈인 듯.
- 요즘 정품은 거의 20만원인데, 그림을 그릴 용도라서 대안도 없다. 
- 다행히 여기저기 뒤지다가 g마켓에서 직구라면서 거의 반값에 파는 걸 발견했다. 이게 말이 되나 싶으면서도 리뷰가 많고 걱정했는데 괜찮다는 글이 많아서 구매.
	- 어떻게 이 가격이 가능한가..는 궁금하다. 대량으로 사뒀다가 안 팔려서 추석 특가로 내놓은 걸까?
- 이제 아이패드가 죽으면 어떡하나.. 싶은 걱정이 생김.

이전 글들은 [[짭명방_25년 9월]]로 옮김


# 이전 일지

- 깃허브 링크는 향후 프로젝트 폴더 링크 이동에 따라 손상될 수 있음
- 이 경우 대부분 `4.Archive` 폴더에서 볼 수 있다.
## 짭명방
- [짭명방 프로젝트 일지 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%EC%9C%A0%EB%8B%88%ED%8B%B0%20-%20%EC%9E%91%EC%9D%80%20%EB%AA%85%EB%B0%A9%20%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0/%EC%9E%91%EC%97%85%20%EC%9D%BC%EC%A7%80/%EC%A7%81%EC%A0%91%20%EC%9E%91%EC%84%B1)
- 옵시디언 링크
	- [[24년 7월]]
	- [[24년 8월]]
	- [[24년 9월]]
	- [[24년 10월]]
	- [[24년 11월]]
	- [[24년 12월]]
	- [[25년 1월]]
	- [[25년 2월]]
	- [[25년 3월]]
	- [[1. Projects, Ongoing/유니티 - 작은 명방 구현하기/작업 일지/직접 작성/25년 4월|25년 4월]]
	- [[1. Projects, Ongoing/유니티 - 작은 명방 구현하기/작업 일지/직접 작성/25년 5월|25년 5월]]
	- [[짭명방_25년 6월]]
	- [[짭명방_25년 7월]]
	- [[짭명방_25년 8월]]
	- [[짭명방_25년 9월]]
## 블로그
- [React + Django 프로젝트 일지 월별 작업 기록 깃허브 링크](https://github.com/dowrave/TIL/tree/main/Obsidian/1.%20Projects%2C%20Ongoing/%ED%8F%AC%ED%8F%B4%20%EA%B2%B8%20%EB%B8%94%EB%A1%9C%EA%B7%B8%20%EB%A7%8C%EB%93%A4%EA%B8%B0/%EC%9B%94%EB%B3%84%20%EC%9E%91%EC%97%85%20%EA%B8%B0%EB%A1%9D)
- 옵시디언 링크
	- [[23년 12월]]
	- [[24년 1월]]
	- [[24년 2월]]
	- [[24년 3월]]
	- [[24년 4월]]
	- [[24년 5월]]
	- [[1. Projects, Ongoing/블로그 만들기/월별 작업 기록/25년 4월|25년 4월]]
	- [[1. Projects, Ongoing/블로그 만들기/월별 작업 기록/25년 5월|25년 5월]]
	- [[블로그_25년 6월|블로그_25년 6월]]
	- [[블로그_25년 7월]]