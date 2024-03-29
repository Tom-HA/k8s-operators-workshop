---
title: "Kubernetes Operators"
weight: 101
---
<br>

Operators are applications that enable us to extend the Kubernetes API with its domain-specific knowledge.  
For example, an operator that manages the creation and deletion of databases, by monitoring the desired state in custom resources (within a Kubernetes cluster), and matching it to the actual state.

Here is another description of operators from the [Kubernetes documentation](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/):  
"_Operators are software extensions to Kubernetes that make use of custom resources to manage applications and their components. Operators follow Kubernetes principles, notably the control loop._"

![k8s operator](./images/k8s-workshop-operator.png "k8s operator")