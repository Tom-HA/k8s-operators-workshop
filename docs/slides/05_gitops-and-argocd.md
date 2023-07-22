# GitOps and Argo CD

## GitOps

GitOps is a deployment methodology, where we declare the desired state of our deployment with a Git repository which acts as the source of truth.  
For example, if the desired state of our deployment consists of a deployment and service manifests, once we've pushed the manifests into a Git repository, we can be sure they will be present in the Kubernetes cluster.  

## Argo CD

Argo CD is a continuous delivery tool that enables us to implement the GitOps methodology.  
It monitors resources that are defined in a Git repository, and applies them into a Kubernetes cluster.  
By continuously comparing the desired state to the current state, Argo CD can help us resolve drifts that can be caused by manual actions using visualization of the changes, and re-synchronization of the desired state.

## Learn More

You can read more about GitOps and Argo CD in [this article by Codefresh](https://codefresh.io/learn/gitops/gitops-with-kubernetes-why-its-different-and-how-to-adopt-it/).
