---
title: "Kubernetes operators and Custom Resources"
weight: 103
---
<br>

## Summary

Now that we've covered the basics of Kubernetes operators, Custom Resource Definitions, and Custom Resources, let's tie it up together.  
Think of the following scenario:  
We need to dynamically provision a database when creating a new microservice.  
To solve it using the operator pattern, we can create a CRD of kind _Postgres_ with a `database` key in its specification (similar to the example in [custom resources slide](./02_custom-resources.md)),  
then, we'll introduce a Kubernetes operator that will constantly watch for new custom resources of kind _Postgres_.  
Once the operator detects a new CR of kind _Postgres_, it will provision a new database, and it will name it according to the `database` value in the CR.  

From the CNCF [operator white paper](https://github.com/cncf/tag-app-delivery/blob/eece8f7307f2970f46f100f51932db106db46968/operator-wg/whitepaper/Operator-WhitePaper_v1-0.md#operator-design-pattern):

![operator-pattern](https://github.com/cncf/tag-app-delivery/blob/eece8f7307f2970f46f100f51932db106db46968/operator-wg/whitepaper/img/02_1_operator_pattern.png?raw=true "operator-pattern")

## Learn More

To go deeper into Kubernetes operators, and Custom Resources, you can read the CNCF [operator white paper](https://github.com/cncf/tag-app-delivery/blob/eece8f7307f2970f46f100f51932db106db46968/operator-wg/whitepaper/Operator-WhitePaper_v1-0.md#operator-design-pattern) that is also referenced above, and the Kubernetes documentaion on [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/).