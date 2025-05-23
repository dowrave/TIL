#Unity 

1. 크기
- 실제 사용할 크기의 2배수로 제작
- 항상 짝수 크기로 제작
- 2의 제곱수 크기 권장(16, 32, 64, 128, ...)

2. 선명도
- `픽셀 퍼펙트`를 위해, 온전한 픽셀 값 사용(반픽셀 위치 피하기)
- 아이콘의 주요 요소는 가능한 `픽셀 그리드`에 맞추기
- 작은 크기에서도 식별화 가능한, 단순화된 디자인 사용

3. 여백
- 아이콘 주변에 1~2픽셀의 여백 확보
- 심볼의 크기는 전체 캔버스의 약 8-90% 정도를 차지하게끔 설계
- 동일 세트의 아이콘은 일관된 규칙 적용

4. 색상
- 단색 아이콘은 순수 흰색`#ffffff, (255, 255, 255)` 사용
- 그라데이션이나 부드러운 효과 피하기
- 안티앨리어싱 최소화

5. 파일 형식
- PNG-24 포맷 사용
- 투명도가 필요하다면 알파채널 포함
- 불필요한 메타데잍 제거

6. 유니티 임포트 설정
- `Texture Type` : `Sprite(2D and UI)`
- `Sprite Mode` : `Single`
- `Pixels Per Unit` : 프로젝트 규칙에 맞게끔 일관되게 설정
- `Filter Mode` : `Point`
- `Compression` : `None` (또는 UI용도에 맞는 압축)

7. 조직화
- 같은 카테고리의 아이콘은 동일한 크기로 제작
- 아이콘 세트별로 일관된 디자인 언어 사용
- 명확한 네이밍 규칙 적용

--- 
## 추가 궁금증과 답변

1. 여러 사이즈에 사용될 것을 대비하면 최초에 아이콘을 만들 때 어떤 사이즈로 만드는 게 좋은가?
	- **(512, 512), (1024, 1024)를 권장** 
		- 충분히 크므로 축소해도 디테일 보존 가능
		- 2의 제곱수라 최적화에 유리
		- 고해상도 디스플레이 대응 가능

2. 픽셀 퍼펙트, 픽셀 그리드란?
- `픽셀 퍼펙트` : 각 픽셀이 정확히 한 픽셀로 표시되며 블러나 앨리어싱 없이 선명하게 표현됨
```
좋은 예: ■■□□  (픽셀 퍼펙트)
나쁜 예: ■▨▨□  (블러된 상태)
```
- `픽셀 그리드` : 이미지를 구성하는 정사각형 픽셀들의 눈금망
```
□□□□
□■■□  <- 픽셀 그리드에 맞춘 아이콘
□■■□
□□□□
```

3. `Max Size` 설정 관련
	- 원본 사이즈를 지정하면 된다. 512면 512, 64면 64.