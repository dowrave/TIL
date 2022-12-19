- **쿠버네티스에서 관리하는 가장 작은 배포 단위**.
- 도커와의 차이점
	- **도커는 컨테이너**를 만든다
	- k8s는 Pod을 만든다. **Pod이란 개념은 1개 이상의 컨테이너를 포함**한다.


#### Pod 빠르게 만들기
```sh
kubectl run echo --image ghcr.io/subicura/echo:v1
# pod/echo created
```
- 참고 : `k8s v1.18` 이상은 `run`이 `Pod`을 만들지만 이 이하는 `Deployment`를 만든다.

```
kubectl get pod
```
- `STATUS`는 컨테이너가 정상적으로 생성되면 Running으로 바뀌고, 오류가 있다면 에러를 표시한다.

- Pod 상태 더 상세하게 확인
```sh
kubectl desribe pod/echo
```
- `describe` 명령어는 해당 리소스의 상세한 정보를 알려준다. 
	- 주로 `Events`를 확인하게 되는데, 현재 Pod의 상태를 이벤트별로 확인할 수 있다.
```
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  108s  default-scheduler  Successfully assigned default/echo to minikube
  Normal  Pulling    106s  kubelet            Pulling image "ghcr.io/subicura/echo"
  Normal  Pulled     85s   kubelet            Successfully pulled image "ghcr.io/subicura/echo" in 20.4301379s
  Normal  Created    85s   kubelet            Created container echo
  Normal  Started    84s   kubelet            Started container echo

```

### Pod 생성 분석
- 클러스터 `minikube` 내부에 `Pod`이 있고, Pod 내부에 `컨테이너`가 있다.
![[pod-single.webp]]

> Pod의 생성 과정
>> 1. `Scheduler`는 API 서버 감시, 할당되지 않은 `Pod`을 체크함
>> 2. `Scheduler`는 할당되지 않은 `Pod`을 적절한 노드에 할당함(`minikube`는 단일노드)
>> 3. `kubelet` : 노드에 설치되어 있으며, 자신에게 할당된 `Pod`을 체크
>> 4. `kubelet`은 `Scheduler`에 의해 자신에게 할당된 `Pod`의 정보를 확인 & 컨테이너 생성
>> 5. `kubelet`은 자신에게 할당된 `Pod`의 정보를 `API 서버에 전달함`
> 이렇듯 노드가 많아지더라도 `Scheduler`만 잘 작동하면 문제가 없는 구조다.

```
kubectl delete pod/echo
```

#### YAML로 설정파일(Spec) 작성하기
- `kubectl run`은 실무에선 거의 안 쓴다.  
- 설정이 복잡하고 다양한데, 명령어로만 표현하면 금방 복잡해지고 관리하기 어렵기 때문.
- 이를 위해 `yaml`파일로 복잡한 내용을 표현 & 변경 내용을 버전으로 관리할 수 있다.
- `echo-pod.yml`
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: echo
  labels:
  app: echo
spec:
  containers:
    - name: app
      image: ghcr.io/subicura/echo:v1
```
- `run` 명령어와의 차이점은, `label`이 추가되었다는 것이다.
- `k8s`는 리소스 관리 시 `name`, `label`을 이용한다.

- **리소스 정의 4가지 필수 요소**
| 정의       | 설명          | 예시                                      |
| ---------- | ------------- | ----------------------------------------- |
| `version`  | 오브젝트 버전 | v1, app/v1, networking.k8s.io/v1 등..     |
| `kind`     | 종류          | Pod, ReplicaSet, Deployment, Service, ... |
| `metadata` | 메타데이터    | name, label, annotation(주석) 등으로 구성 |
| `spec`     | 상세명세      | 리소스 종류마다 다르다 |
>`version`
>>쿠버네티스 버전에 따라 지원하는 리소스의 버전이 다르기 때문에 실습 시 유의해야 함
>>>`Alpha(v1alpha1, ....)`
>>>`Beta(v1beta1, ...)`
>>>`Stable(v1)`

```sh
kubectl apply -f echo-pod.yml

# 생성까지 시간 좀 걸림
kubectl get pod

kubectl logs echo
kubectl logs -f echo # 로그 실시간 조회

# Pod 컨테이너 접속
kubectl exec -it echo -- sh

# 컨테이너 내부에서
ls # 파일 리스트
ps # 현재 돌아가고 있는 프로세스
exit

# Pod 제거
kubectl delete -f echo-pod.yml
```

### 컨테이너 상태 모니터링
- `컨테이너 생성`과 `서비스 준비`는 약간의 차이가 있다.
- 서버를 실행하면 짧게 수초, 길게 수분의 초기화 시간이 필요하며, **실제 접속이 가능할 때 서비스가 준비되었다**고 할 수 있다.
- 쿠버네티스는 초기화하는 동안 서비스되는 것을 막을 수 있다.

#### livenessProbe
- 컨테이너가 정상적으로 동작되는지 체크, 아니라면 컨테이너를 재시작하여 문제를 해결한다.
- 체크 방식은 여러 가지가 있으나, 여기서는 `http get` 요청 확인으로 체크함

```yml
apiVersion: v1
kind: Pod
metadata:
  name: echo-lp
  labels:
    app: echo
spec:
  containers:
    - name: app
      image: ghcr.io/subicura/echo:v1
      livenessProbe:
        httpGet:
          path: /not/exist
          port: 8080
        initialDelaySeconds: 5
        timeoutSeconds: 2 # Default 1
        periodSeconds: 5 # Defaults 10
        failureThreshold: 1 # Defaults 3
```
- 의도적으로 존재하지 않는 `path(/not/exist)`와 `port(8080)`이 입력됨
- 