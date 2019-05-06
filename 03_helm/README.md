# [Workshop](../README.md) &raquo; Helm basics

This guide covers how you can quickly get started using Helm.

## Prerequisites

The following prerequisites are required for a successful and properly secured
use of Helm.

1. A Kubernetes cluster
2. [Installed and configured Helm and Tiller, the cluster-side service](./00_install_helm.md).
3. Installed helm client

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