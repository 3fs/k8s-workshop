# [Workshop](../README.md) &raquo; Custom Helm

Helm is a tool that streamlines installing and managing Kubernetes applications.
Think of it like apt/yum/homebrew for Kubernetes.

- Helm has two parts: a client (`helm`) and a server (`tiller`)
- Tiller runs inside of your Kubernetes cluster, and manages releases (installations)
  of your charts.
- Helm runs on your laptop, CI/CD, or wherever you want it to run.
- Charts are Helm packages that contain at least two things:
  - A description of the package (`Chart.yaml`)
  - One or more templates, which contain Kubernetes manifest files
- Charts can be stored on disk, or fetched from remote chart repositories
  (same as Debian or RedHat packages)

## Prerequisites

Execute all the tasks in the preprepared container. Instructions can be found [here](../02_kubernetes/README.md#access-to-your-namespace-in-workshop-k8s-cluster)

## Tasks

1. [Create simple helm chart](./01_create_helm_chart.md)
2. [Enable Ingress resource](./02_enable_ingress.md)
3. [Add custom variables](./03_add_variables.md)
4. [Configure health probes](./04_configure_probes.md)

## Demo

Example chart is available [here](https://github.com/3fs/k8s-workshop/blob/master/04_custom_helm/README.md)

[![asciicast](https://asciinema.org/a/SY034jPcErkVmVMBeb9lXc9Xp.svg)](https://asciinema.org/a/SY034jPcErkVmVMBeb9lXc9Xp)
