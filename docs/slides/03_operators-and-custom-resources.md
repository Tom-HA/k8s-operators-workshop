# Kubernetes operators and Custom Resources

## Summary

Now that we've covered the basics of Kubernetes operators, Custom Resource Definitions, and Custom Resources, let's tie it up together.  
Think of the following scenario:  
We work in microservices environment, and we need to be able to declaratively provision new databases.  
To solve it, we can create a CRD of kind _database_ with a `databaseName` key in its specification (like in the example in [custom resources slide](./02_custom-resources.md)).  
Then, we'll introduce a Kubernetes operator that will constantly watch for new custom resources of kind _database_.  
Once the operator detects a new CR of kind _database_, it will provision a new database, and it will name it according to the `databaseName` value in the CR.  

## Learn More

To go deeper into Kubernetes operators, and Custom Resources, you can read the CNCF [operator white paper](https://github.com/cncf/tag-app-delivery/blob/eece8f7307f2970f46f100f51932db106db46968/operator-wg/whitepaper/Operator-WhitePaper_v1-0.md#operator-design-pattern), and Kubernetes documentaion on [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/).
