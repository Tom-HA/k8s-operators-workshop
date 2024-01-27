---
title: "Create PostgreSQL Credentials"
weight: 223
---
</br>

In this slide we are going to provision a user and password for the database which will be stored in a Kubernetes secret, and configure our sample-pg-app to retrieve the authentication details from that secret.

## Create PostgreSQL user

We can declare a new user that will be associated to the database by creating the following `.yaml` in [apps/services/sample-pg-app/hooks](https://github.com/Tom-HA/k8s-operators-workshop/tree/main/apps/services/sample-pg-app/hooks):  
{{% alert context="warning" %}}
Again, make sure the file name you choose ends with `.yaml`, for example, `database-user.yaml`.
{{% /alert %}}

```yaml
apiVersion: db.movetokube.com/v1alpha1
kind: PostgresUser
metadata:
  name: sample-pg-app-auth
  namespace: services
spec:
  role: sample-user
  database: sample-pg-app-db       # This references the Postgres CR
  secretName: postgres
  privileges: OWNER     # Can be OWNER/READ/WRITE
```

Once you've pushed your changes, you should see a new _postgresuser_ and a Kubernetes secret:

![postgres-user](./images/postgres-user.png "postgres-user")

{{% alert context="info" %}}
You can decode the Kubernetes secret by executing
{{% /alert %}}

```sh
kubectl -n services get secret postgres-sample-pg-app-auth -o go-template='{{range $k,$v := .data}}{{printf "%s: " $k}}{{$v | base64decode}}{{"\n"}}{{end}}'
```
