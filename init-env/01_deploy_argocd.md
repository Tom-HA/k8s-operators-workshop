# Deploy Argo CD

In this section, we'll use _Helmfile_ to install Argo CD, this enables us to complete this workshop by following the GitOps methodology.  
Using Argo CD, we'll provision a PostgreSQL database and a PostgreSQL operator.  

## Initial provisioning

Open the `helmfile.yaml` file, there, make sure the `kubeContext` is configured with the Kubernetes cluster you want to work with.  
Next, install Argo-CD by executing:

```sh
helmfile apply
```

This will install the Argo-CD's Helm chart with the values under:  

```sh
./init-env/values/argo-cd/values.yaml
```

## Access Argo CD UI

Now that Argo CD is installed, we can access its UI.  
To do that, we'll use port forwarding to access the Kubernetes service of Argo CD's UI.  
The following command will forward traffic from _localhost_ at port 8080 to the _argocd-server_ service at port 80:

```sh
kubectl -n argocd port-forward service/argo-cd-argocd-server 8080:80 &
```


## Log in to Argo CD UI

The default admin user name is `admin`.  
The password is stored in a Kubernetes secret which is generated during the initial deployment and it's base64 encoded.  
To retrieve and decode the password, execute:

```sh
printf "%s\n" $(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" |base64 -d)
```

### :warning: Note

If you want to stop the port forwarding, execute:

```sh
kill %1
```
