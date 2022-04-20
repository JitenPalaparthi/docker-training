# Kubernetes

### Version of Kubernetes client

```kubectl version```

### Cluster information

```kubectl cluster-info```

### Cluster complete information

```kubectl cluster-info dump```

## Config and Context

### View all configs

```kubectl config view```

```kubectl config current-context```

- To use differnet context

```kubectl config use-context demo```

- Get resources from a cluster using Kubeconfig file

```kubectl --kubeconfig $HOME/.kube/config```

## Find resources

### Get all 

```kubectl get all -A```

 - Based on a specific namespace
  
```kubectl get all -n default```

- To fetch all namespaces

```kubectl get namespaces```

- To fetch all pods 

```kubectl get pods -A```

- To create a pod 

```kubectl run nginx --image=nginx --restart=Never```

- To know more information about pod

```kubectl describe pod nginx -n test```
