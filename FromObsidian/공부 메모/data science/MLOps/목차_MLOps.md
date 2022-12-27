[모두의 MLOps](https://mlops-for-all.github.io/)

- minikube 실행
```sh
sudo su # root유저로 접근(--driver 옵션 때문)

minikube start --driver=none --kubernetes-version=v1.21.7 --extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key --extra-config=apiserver.service-account-issuer=kubernetes.default.svc
```

- `/tmp/juju~~` 관련 에러 발생 시
```sh
sudo rm -rf /tmp/juju-mk*
sudo rm -rf /tmp/minikube.*
```

개념 배우기
1. [[MLOps란 무엇인가]]
2. [[MLOps의 단계]]
3. [[MLOps의 구성요소]]
4. [[왜 쿠버네티스인가]]

실습
1. [[MLOps 시스템 구축하기]]
2. [[Kubeflow 설치]]

