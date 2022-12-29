## 1. Base Model Development

- 라이브러리 설치
```sh
pip install pandas scikit-learn joblib
```
- `joblib` : 파이썬에서 병렬처리를 가능하게 함

- 학습 & 평가 데이터 선정
```python
from sklearn.dataset import load_iris
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from sklearn.svm import SVC
from sklearn.metrics import accuracy_score

import joblib

X, y = load_iris(return_X_y = True,
				as_frame = True)

X_train, X_valid, y_train, y_valid = train_test_split(X, y, test_size = 0.2, random_state = 2022)

scaler = StandardScaler()

scaled_X_train = scaler.fit_transform(X_train)
scaled_X_valid = scaler.transform(X_valid)

# 서포트 벡터 머신 분류기
classifier = SVC()

classifier.fit(scaled_X_train, y_train)

train_pred = classifier.predict(scaled_X_train)
valid_pred = classifier.predict(scaled_X_valid)

# accuracy_score(y_true, y_pred)
train_acc = accuracy_score(y_train, train_pred)
valid_acc = accuracy_score(y_valid, valid_pred)

print(train_acc, valid_acc)
```
> svc가 약간 낯설긴 하나 다 아는 거니까 특별히 코멘트 없이 지나감


- **모델 저장** : `joblib` 패키지 이용
```python
import joblib

joblib.dump(scaler, 'scaler.joblib')
joblib.dump(classifier, 'classifier.joblib')
```
- 경로에 `scaler.joblib`, `classifier.joblib`이 저장됨

- **모델 불러오기** : 위에서 출력한 `acc` 값들과 동일한지 확인
```python
scaler_load = joblib.load('scaler.joblib')
classifier_load = joblib.load('classifier.joblib')
```
- 위의 훈련 코드에 적용해보면 됨 : 똑같이 분류한 데이터셋에 동일한 결과가 나오는가 보자

- 이를 코드로 옮기면
`base_train.py` : 학습 & 데이터 저장
```python
# base_train.py  
import joblib  
from sklearn.datasets import load_iris  
from sklearn.metrics import accuracy_score  
from sklearn.model_selection import train_test_split  
from sklearn.preprocessing import StandardScaler  
from sklearn.svm import SVC  
  
# 1. get data  
X, y = load_iris(return_X_y=True, as_frame=True)  
X_train, X_valid, y_train, y_valid = train_test_split(X, y, train_size=0.8, random_state=2022)  
  
# 2. model development and train  
scaler = StandardScaler()  
classifier = SVC()  
  
scaled_X_train = scaler.fit_transform(X_train)  
scaled_X_valid = scaler.transform(X_valid)  
classifier.fit(scaled_X_train, y_train)  
  
train_pred = classifier.predict(scaled_X_train)  
valid_pred = classifier.predict(scaled_X_valid)  
  
train_acc = accuracy_score(y_true=y_train, y_pred=train_pred)  
valid_acc = accuracy_score(y_true=y_valid, y_pred=valid_pred)  
  
print("Train Accuracy :", train_acc)  
print("Valid Accuracy :", valid_acc)  
  
# 3. save model  
joblib.dump(scaler, "scaler.joblib")  
joblib.dump(classifier, "classifier.joblib")
```

`base_validate_save_model.py` : 저장된 모델 검증
```python
# base_validate_save_model.py  
  
import joblib  
from sklearn.datasets import load_iris  
from sklearn.metrics import accuracy_score  
from sklearn.model_selection import train_test_split  
  
# 1. reproduce data  
X, y = load_iris(return_X_y=True, as_frame=True)  
X_train, X_valid, y_train, y_valid = train_test_split(X, y, train_size=0.8, random_state=2022)  
  
# 2. load model  
scaler_load = joblib.load("scaler.joblib")  
classifier_load = joblib.load("classifier.joblib")  
  
# 3. validate  
scaled_X_train = scaler_load.transform(X_train)  
scaled_X_valid = scaler_load.transform(X_valid)  
  
load_train_pred = classifier_load.predict(scaled_X_train)  
load_valid_pred = classifier_load.predict(scaled_X_valid)  
  
load_train_acc = accuracy_score(y_true=y_train, y_pred=load_train_pred)  
load_valid_acc = accuracy_score(y_true=y_valid, y_pred=load_valid_pred)  
  
print("Load Model Train Accuracy :", load_train_acc)  
print("Load Model Valid Accuracy :", load_valid_acc)
```

- 성능 확인
```sh
python3 base_train.py
python3 base_validate_save_model.py
```

## 2. 모델 파이프라인
- `svc`는 사용에 앞서 `scaler`가 필요함
	- `scaler`를 사용하지 않거나 다른 `scaler`를 사용하는 경우가 발생할 수 있다

- 이런 실수를 방지할 수 있는 방법이 `파이프라인`이다.
	- 학습이 완료된 scaler와 svc를 같이 넣으면 됨
- 물론 단점이 있다.
	- 고치기 어려움
	- `scaler` 같이 여러 모델에 쓸 수 있는 걸 중복적으로 학습해야 함

`pipeline_train.py`의 일부 
```python
import joblib
from sklearn.preprocessing import StandardScaler
from sklearn.pipeline import Pipeline
from sklearn.svm import SVC
from sklearn.metrics import accuracy_score

model_pipeline = Pipeline([("scaler", StandardScaler), ("svc", SVC())])

model_pipeline.fit(X_train, y_train)

# scaler만 쓰고 싶다면 이런 식으로 쓸 수 있다
print(model_pipeline[0].transform(X_train[:1]))

# 바로 예측 & 정확도 계산
train_pred = model_pipeline.predict(X_train)  
valid_pred = model_pipeline.predict(X_valid)  
  
train_acc = accuracy_score(y_true=y_train, y_pred=train_pred)  
valid_acc = accuracy_score(y_true=y_valid, y_pred=valid_pred)  
  
print("Train Accuracy :", train_acc)  
print("Valid Accuracy :", valid_acc)

# 모델 저장
joblib.dump(model_pipeline, 'model_pipeline.joblib')
```

- 전체 `pipeline_train.py`
```python
# pipeline_train.py  
import joblib  
from sklearn.datasets import load_iris  
from sklearn.metrics import accuracy_score  
from sklearn.model_selection import train_test_split  
from sklearn.preprocessing import StandardScaler  
from sklearn.pipeline import Pipeline  
from sklearn.svm import SVC  
  
# 1. get data  
X, y = load_iris(return_X_y=True, as_frame=True)  
X_train, X_valid, y_train, y_valid = train_test_split(X, y, train_size=0.8, random_state=2022)  
  
# 2. model development and train  
model_pipeline = Pipeline([("scaler", StandardScaler()), ("svc", SVC())])  
model_pipeline.fit(X_train, y_train)  
  
train_pred = model_pipeline.predict(X_train)  
valid_pred = model_pipeline.predict(X_valid)  
  
train_acc = accuracy_score(y_true=y_train, y_pred=train_pred)  
valid_acc = accuracy_score(y_true=y_valid, y_pred=valid_pred)  
  
print("Train Accuracy :", train_acc)  
print("Valid Accuracy :", valid_acc)  
  
# 3. save model  
joblib.dump(model_pipeline, "model_pipeline.joblib")
```

- 앞에서와 마찬가지로 저장된 모델을 불러와 동일한 결과를 출력하는지 확인
```python
# pipeline_validate_save_model.py  
import joblib  
from sklearn.datasets import load_iris  
from sklearn.metrics import accuracy_score  
from sklearn.model_selection import train_test_split  
  
# 1. reproduce data  
X, y = load_iris(return_X_y=True, as_frame=True)  
X_train, X_valid, y_train, y_valid = train_test_split(X, y, train_size=0.8, random_state=2022)  
  
# 2. load model  
model_pipeline_load = joblib.load("model_pipeline.joblib")  
  
# 3. validate  
load_train_pred = model_pipeline_load.predict(X_train)  
load_valid_pred = model_pipeline_load.predict(X_valid)  
  
load_train_acc = accuracy_score(y_true=y_train, y_pred=load_train_pred)  
load_valid_acc = accuracy_score(y_true=y_valid, y_pred=load_valid_pred)  
  
print("Load Model Train Accuracy :", load_train_acc)  
print("Load Model Valid Accuracy :", load_valid_acc)
```

- 실행
```sh
python3 pipeline_train.py
python3 pipeline_validate_save_model.py
```

[여기부터](https://mlops-for-mle.github.io/tutorial/docs/model-development/load-data-from-database)