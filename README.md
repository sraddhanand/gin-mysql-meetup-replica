
## Meetup Clone using Go-Gin & MySql

#### Repo Structure 
- **gin** : Contains go-gin code
    - **api** : All controller functions exists here
    - **middleware** : JWT tokenization 
    - **model** : Schema defination & model functions
    - **routes** : API route configurations
    - **test** : Go tests 
    - **utils** : Utility functions such as create files, return responce, etc.
- **docker-compose.yaml** : Bring up application for development
- **Dockerfile** : multistage dockerfile to build image container compilied minary
- **argo.yaml** : K8S manifest to deploy the app using GitOps practices  

### Configure Dev Environment
```bash
docker-compose up -d
```

### GitOps deployment 
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/core-install.yaml
kubectl apply -n argocd -f argo.yaml
```

### Cleanup
```bash
kubectl delete -n argocd -f argocd
kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/core-install.yaml
```
