# Custom Resource Definitions and Custom Resources

## Custom Resource Definitions

With CRDs (custom resource definitions) we can define new resources in Kubernetes.  
For example, some default resources we already have in a cluster are of kind: _pod_, _deployment_, _service_, etc., each with its specification.  
Using a CRD we can define a new resource of kind _Postgres_ and its YAML specification.  
The new database resource specification can contain a key named `database`, and its value is expected to be a database name.  
Here is an example from the [postgres-operator project](https://github.com/movetokube/postgres-operator/blob/master/charts/ext-postgres-operator/crds/db.movetokube.com_postgres_crd.yaml):  

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: postgres.db.movetokube.com
spec:
  group: db.movetokube.com
  names:
    kind: Postgres      # <------

...
        spec:
        description: PostgresSpec defines the desired state of Postgres
        properties:
            database:     # <------
            type: string
            dropOnDelete:
            type: boolean
```

## Custom Resource

With CRs (custom resources) we can implement our custom resource definition.  
For example, if we defined a resource of kind _Postgres_ using a CRD (like in the example above), we can create a YAML manifest that implements the resource specification of kind _Postgres_.  

Here is another example from the [postgres-operator project](https://github.com/movetokube/postgres-operator#postgres):

```yaml
apiVersion: db.movetokube.com/v1alpha1
kind: Postgres      # <------
metadata:
  name: my-db
  namespace: app
  annotations:
    postgres.db.movetokube.com/instance: POSTGRES_INSTANCE
spec:
  database: test-db     # <------
  dropOnDelete: false 
...
```
