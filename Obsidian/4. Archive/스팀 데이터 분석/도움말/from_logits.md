- 모델 출력이 활성화함수로 `softmax`나 `sigmoid`가 적용되지 않은 상황이라면, 출력은 `로짓Logit` 형태이다. 이는 순수한 모델의 선형 출력 값이다.

- 모델의 출력이 로짓 형태라면, `from_logits = True`로 변환해서 모델 출력을 확률로 변환할 수 있다 : 일부 손실 함수(크로스 엔트로피 등)에서 인풋으로 로짓을 기대한다.