---
title: "Initial Provisioning Overview"
weight: 211
---
<br>

# Overview

Once the prerequisite (see the [prerequisites section](../01_intro.md#prerequisites)) are completed, we'll need to initialize the environment.

To initialize the environment, follow the documentation under the [init-env](../init-env) folder.
In this section, we'll use _Helmfile_ to install Argo CD, this enables us to complete this workshop by following the GitOps methodology.  
Using Argo CD, we'll provision a PostgreSQL database and a PostgreSQL operator.  

## Deploy Argo CD

Open the `helmfile.yaml` file, there, make sure the `kubeContext` is configured with the Kubernetes cluster you want to work with.  
Next, install Argo-CD by executing:

```sh
helmfile apply
```

This will install the Argo-CD's Helm chart with the values under:  

```sh
./init-env/values/argo-cd/values.yaml
```

## Forward traffic to Argo CD

Now that Argo CD is installed, we'll need to forward traffic to the Argo CD service, so we could access its web UI.  
To do that, we'll use kubectl port-forward to forward traffic from localhost to the Argo CD (Kubernetes) service.  
The following command will forward traffic from _localhost_ at port 8080 to the _argocd-server_ service at port 80:

```sh
kubectl -n argocd port-forward service/argo-cd-argocd-server 8080:80 &
```

## Log in to Argo CD UI

In your browser, go to [localhost:8080](http://localhost:8080).  

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
