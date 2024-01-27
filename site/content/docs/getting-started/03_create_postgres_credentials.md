---
title: "Create PostgreSQL Credentials"
weight: 3
---
</br>

In this slide we are going to provision a user and password for the database which will be stored in a Kubernetes secret, and configure our sample-pg-app to retrieve the authentication details from that secret.

## Create PostgreSQL user

We can declare a new user that will be associated to the database by creating the following `.yaml` in [apps/services/sample-pg-app/hooks](apps/services/sample-pg-app/hooks):  
> [!WARNING]  
> Again, make sure the file name you choose ends with `.yaml`, for example, `database-user.yaml`.

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

> [!NOTE]  
> You can decode the Kubernetes secret by executing

```sh
kubectl -n services get secret postgres-sample-pg-app-auth -o go-template='{{range $k,$v := .data}}{{printf "%s: " $k}}{{$v | base64decode}}{{"\n"}}{{end}}'
```  

## Assign credentials to application

Although we've created a new Postgres user and a Kubernetes secret, the health status is still degraded.  
This is because we still need to reference the secret in the sample-pg-app.  

We'll do that by adding the following `values.yaml` file to [apps/services/sample-pg-app](apps/services/sample-pg-app)

```yaml
postgres:
  auth:
    secretName: "postgres-sample-pg-app-auth"
```

> [!NOTE]  
> You can see the default Helm values of sample-pg-app in [sample-pg-app/chart/sample-pg-app/values.yaml](sample-pg-app/chart/sample-pg-app/values.yaml)

Once your changes have been synced, you should see the sample-pg-app in a healthy state!

![healthy-sample-pg-app](./images/healthy-sample-pg-app.png "healthy-sample-pg-app")

## Bonus

If you want to be extra sure the sample-pg-app is able to communicate with the database, we can set up port-forwarding to the sample-pg-app and send API requests.  

### Port Forwarding

To set up port forwarding to the sample-pg-app service, execute

```sh
kubectl -n services port-forward svc/sample-pg-app 8081:80
```

This will forward the traffic from your localhost at port 8081 to the sample-pg-app Kubernetes service at port 80.  
Now store some data in the database by sending an API request to sample-pg-app

```sh
curl --location --request POST 'localhost:8081/data' \
    --header 'Content-Type: application/json' \
    --data-raw '{"data":"k8s-operators-workshop"}'
```

> [!NOTE]  
> Remember, when sending `curl` request to `localhost:8081` the traffic will be forwarded to the sample-pg-app service

To make sure the data is actually stored in the database, retrieve the data with another API request

```sh
curl --location --request GET 'localhost:8081/data'
```

If the data was stored successfully, you should see the following response

```sh
[{"data":"k8s-operators-workshop"}]
```

---

## Learn more

You can read more about accessing Kubernetes resources using port-forwarding in the [Kubernetes documentation](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/)