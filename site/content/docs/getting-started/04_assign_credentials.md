---
title: "Assign credentials to application"
weight: 224
---
</br>

Although we've created a new Postgres user and a Kubernetes secret, the health status is still degraded.  
This is because we still need to reference the secret in the sample-pg-app.  

We'll do that by adding the following `values.yaml` file to the [apps/services/sample-pg-app](https://github.com/Tom-HA/k8s-operators-workshop/tree/main/apps/services/sample-pg-app) folder.

```yaml
postgres:
  auth:
    secretName: "postgres-sample-pg-app-auth"
```

{{% alert context="info" %}}
You can see the default Helm values of sample-pg-app in [sample-pg-app/chart/sample-pg-app/values.yaml](https://github.com/Tom-HA/k8s-operators-workshop/blob/main/sample-pg-app/chart/sample-pg-app/values.yaml)
{{% /alert %}}

Once your changes have been synced, you should see the sample-pg-app in a healthy state!

![healthy-sample-pg-app](./images/healthy-sample-pg-app.png "healthy-sample-pg-app")
