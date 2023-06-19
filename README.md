# k8s-operators-workshop
A Kubernetes operators workshop

## pre-requisites

Before you start the workshop, you'll need to have a Kubernetes cluster available -   
You can run a local cluster using [K3s with K3d](https://k3d.io/v5.4.6/#installation), [Kubernetes with Docker desktop](https://docs.docker.com/desktop/kubernetes/), [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) or any other method you prefer.

Make sure you have the following installed:

* [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
* [helm](https://helm.sh/docs/intro/install/)
* [helmfile](https://github.com/helmfile/helmfile#installation)
* [helm-diff plugin](https://github.com/databus23/helm-diff)

## initializing the environemnt

Go to the `init-env` directory and open the `helmfile.yaml` file, there, make sure the `kubeContext` is configured with the correct cluster.  
Next, install Argo-CD by executing:
```sh
helmfile apply
```
This will install the Argo-CD Helm chart with the values under `./init-env/values/argo-cd/values.yaml`