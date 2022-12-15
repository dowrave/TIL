- 대표적인 오픈소스 설치 & 서로 연동하여 사용하는 부분만을 주로 다룰 것
- 2022년 기준 아직 대표적인 표준 MLOps는 없음! 적절한 툴은 상황에 맞춰 취사선택할 것.

### 1. 쿠버네티스 실습 환경 구축
- 쿠버네티스 환경 구축 도구에는
	- 프로덕션 레벨 : `kubeadm`
	- 쉬운 버전 : `kubespray`, `kops`
	- 학습 목적 : `k3s`, `minikube`, `microk8s`, `kind`

- 여기서는 `kubeadm`, `k3s`, `minikube` 3가지 도구를 비교해볼 것
- 모든 기능을 사용하고 노드 구성도 하고 싶다면 `kubeadm`을 권장함

### 2. 사전에 설치할 것들
1. 포트포워팅을 위해 **클러스터**에 설치할 것들
```ubuntu
sudo apt-get update
sudo apt-get install -y socat
```
- `socat`은 포트포워딩을 위한 패키지

2. 도커 설치
1) 도커에 필요한 APT 패키지 설치
```ubuntu
sudo apt-get update && sudo apt-get install -y ca-certificates curl gnupg lsb-release
```
2) 도커의 공식 GPG key 추가
```ubuntu
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```
3) apt 패키지 매니저로 도커 설치 시 stable Repository에서 받도록 설정
```ubuntu
echo \ "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \ $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
4) 현재 설치할 수 있는 도커 버전 확인
```ubuntu
sudo apt-get update && apt-cache madison docker-ce
```
- `5:20.10.11~3-0~ubuntu-focal` 버전 있는지 확인
 ```ubuntu
 apt-cache madison docker-ce | grep 5:20.10.11~3-0~ubuntu-focal

// 정상 추가시 아래처럼 출력
docker-ce | 5:20.10.11~3-0~ubuntu-focal | https://download.docker.com/linux/ubuntu focal/stable amd64 Packages
```

5) 도커 설치
```ubuntu
sudo apt-get install -y containerd.io docker-ce=5:20.10.11~3-0~ubuntu-focal docker-ce-cli=5:20.10.11~3-0~ubuntu-focal
```

6) 도커 설치 확인
```ubuntu
sudo docker run hello-world
```
- Hello from Docker! 뜨면 ㅇㅋ

7) Sudo 없이 docker 커맨드 사용하도록 권한 추가


```ubuntu
sudo groupadd docker 
sudo usermod -aG docker $USER 
newgrp docker
```

8) (7) 확인하기
```ubuntu
docker run hello-world
```

9) `kubelet`의 정상 작동을 위해, 클러스터 노드에서 `swap`이라고 불리는 가상 메모리를 꺼둔다. 
```ubnutu
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab sudo swapoff -a
```
- 클러스터 - 클라이언트가 같은 데스크톱일 때 swap을 끄면 속도 저하가 있을 수 있음

## 3. Kubectl 설치
- 클라이언트 노드에 설치해야 함

1) `kubectl v1.21.7` 다운
```ubuntu
curl -LO https://dl.k8s.io/release/v1.21.7/bin/linux/amd64/kubectl
```

2) 파일 권한 & 위치 변경
```ubuntu

sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

```

3) 설치 확인
```ubuntu
kubectl version --client

// 정상 출력 메시지
Client Version: version.Info{Major:"1", Minor:"21", GitVersion:"v1.21.7", GitCommit:"1f86634ff08f37e54e8bfcd86bc90b61c98f84d4", GitTreeState:"clean", BuildDate:"2021-11-17T14:41:19Z", GoVersion:"go1.16.10", Compiler:"gc", Platform:"linux/amd64"}
```