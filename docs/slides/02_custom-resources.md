# Custom Resource Definitions and Custom Resources

## Custom Resource Definitions

With CRDs (custom resource definitions) we can define new resources in Kubernetes.  
For example, some default resources we already have in a cluster are of kind: _pod_, _deployment_, _service_, etc., each with its specification.  
Using a CRD we can define a new resource of kind _database_ and its YAML specification.  
The new database resource specification can contain a key named `databaseName`, and its value is expected to be a database name.  

## Custom Resource

With CRs (custom resources) we can implement our custom resource definition.  
For example, if we defined a resource of kind _database_ using a CRD (like in the example above), we can create a YAML manifest that implements the resource specification of kind _database_.  
