
## `k3s`

#### 사전 준비
- `k3s`는 기본값으로 `containerd`를 백엔드로 이용한다. 
- 그러나 GPU를 쓰기 위해선 `docker`가 백엔드가 되어야 하므로, `--docker` 옵션을 추가한다.
```ubuntu
curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=v1.21.7+k3s1 sh -s - server --disable traefik --disable servicelb --disable local-storage --docker
```
- `k3s config` 확인
```ubuntu
sudo cat /etc/rancher/k3s/k3s.yaml
```
- 대충 apiVersion, Clusters, ... 등등 나오면 ok

#### 쿠버네티스 클러스터 셋업
```ubuntu
mkdir .kube
sudo cp /etc/rancher/k3s/k3s.yaml .kube/config
```
- 복사된 config 파일에 유저 접근 권한 부여
```ubuntu
sudo chown $USER:$USER .kube/config
```

#### 쿠버네티스 클라이언트 셋업
- 클러스터에서 설정한 `kubeconfig`를 로컬로 이동함. 
	- 로컬 경로는 `~/.kube/config`
- 처음 복사한 config파일에는 서버 ip가 `https://127.0.0.1:6443`으로 되어 있는데, 이를 클러스터의 ip에 맞게 `https://192.168.0.19:6443`으로 수정함

