---
title: "Test Application"
weight: 225
---
</br>

Let's test that the sample-pg-app is able to communicate with the database.  
We need to establish connection with the sample-pg-app ourselves first, for that, we can set up port-forwarding to the sample-pg-app and send API requests.  

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

{{% alert context="info" %}}  
Remember, when sending `curl` request to `localhost:8081` the traffic will be forwarded to the sample-pg-app service
{{% /alert %}}

To make sure the data is actually stored in the database, retrieve the data with another API request

```sh
curl --location --request GET 'localhost:8081/data'
```

If the data was stored successfully, you should see the following response

```sh
[{"data":"k8s-operators-workshop"}]
```

{{% alert icon=":tada:" context="success" %}}

That's it! You've completed the workshop! Hope you had fun along the way :relaxed:

{{% /alert %}}

---

## Learn more

You can read more about accessing Kubernetes resources using port-forwarding in the [Kubernetes documentation](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/)