# Helm and Helmfile

## Helm

Helm is a package manager for Kubernetes, it allows us to package multiple manifests that comprise our application.  
This package, also known as Helm chart, can be installed, upgraded and rolled back using Helm.  

## Helmfile

Helmfile is a utility that help us manage multiple Helm charts.  
For example, if we want to install three Helm charts, instead of installing each chart one by one, we can define them in a `helmfile.yaml` manifest, and using the Helmfile utility, we can install all of them with one command.

## Learn More

You can read more about Helm and Helmfile in the following links:

* [Helm documentation](https://helm.sh/docs/topics/charts/)
* [Helmfile documentation](https://helmfile.readthedocs.io/en/latest/#about)
