# Custom Resources Definition

With CRDs (custom resource definitions) we can define new resources in Kubernetes.  
For example, some default resources we already have in a cluster are of kind: _pod_, _deployment_, _service_, etc.  
Using a CRD we can define a new resource of kind _database_ and its YAML specification.  
The new database resource specification can contain a key named `database`, and its value is expected the database name.  
