# [Workshop](../README.md) &raquo; Helm basics

This guide covers how you can quickly get started using Helm.

## Prerequisites

The following prerequisites are required for a successful and properly secured
use of Helm.

1. A Kubernetes cluster
2. Preconfigured container used for task in previous task: [Instructions](../02_kubernetes/README.md#access-to-your-namespace-in-workshop-k8s-cluster)
3. [Installed and configured Helm and Tiller, the cluster-side service](./00_install_helm.md).
4. Installed helm client (already installed in the container)

## Introduction

In this tasks we will cover installing PostgreSQL database server into k8s
cluster from official chart available
[here](https://github.com/helm/charts/tree/master/stable/postgresql).

The chart provides multiple configuration options. For the purpose of the
following tasks, we will use change few of them.

## Tasks

- [Install charts](./01_install_chart.md)
- [Inspect upgrade](./02_inspect_upgrade.md)
- [Delete release](./03_delete_release.md)

## Demo of all tasks

[![asciicast](https://asciinema.org/a/wBacx1Pd7gEgfqvA2BUjMEG67.svg)](https://asciinema.org/a/wBacx1Pd7gEgfqvA2BUjMEG67)

## Next: [Custom Helm](../04_custom_helm/README.md)
