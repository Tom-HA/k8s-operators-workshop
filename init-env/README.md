# Initializing the environment

Open the `helmfile.yaml` file, there, make sure the `kubeContext` is configured with the Kubernetes cluster you want to work with.  
Next, install Argo-CD by executing:
```sh
helmfile apply
```
This will install the Argo-CD's Helm chart with the values under:  
```sh
./init-env/values/argo-cd/values.yaml
```