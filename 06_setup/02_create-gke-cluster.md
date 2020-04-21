# [Workshop](../README.md) &raquo;[Workshop setup](./README.md) &raquo;Create GKE cluster

## Table of contents

* [Enable GKE API](#enable-gke-api)
* [Create GKE Cluster](#create-gke-cluster)
* [Official documentation](#official-documentation)
* [Set up kubectl](#set-up-kubectl)
* [Initialize Helm](#initialize-helm)
* [Nginx ingress controller](#nginx-ingress-controller)
* [Create certificate secret](#create-certificate-secret)
* [Expose kubernetes dashboard](#expose-kubernetes-dashboard)

## Enable GKE API

To be able to execute tasks below, successful `gcloud` login needs to be
performed. How to do that can be found [here](./01_google_cloud.md)

Enable Google Kubernetes Engine API:

```bash
gcloud services enable container.googleapis.com
 ```

## Create GKE Cluster

Create a Kubernetes cluster with:

* nodes with 2 vCPUs and 13 GB of memory capacity (type `n1-highmem-2`)
* location in Germany, Europe
* kubernetes dashboard addon
* cluster autoscaling enabled

```bash
gcloud container clusters create workshop-cluster \
  --machine-type=n1-highmem-2 \
  --zone europe-west3-c \
  --addons=KubernetesDashboard \
  --enable-autoscaling \
  --max-nodes=15 \
  --min-nodes=3
 ```

## Official documentation

* [GKE Quickstart
  guide](https://cloud.google.com/kubernetes-engine/docs/quickstart?authuser=1&refresh=1)

## Set up kubectl

Before we can use `kubectl` we must configure it to communicate with our GKE cluster.

```bash
gcloud container clusters get-credentials workshop-cluster
```

## Install NGINX ingress controller using helm3

Install [NGINX ingress
controller](https://github.com/nginxinc/kubernetes-ingress) in `kube-system`
namespace with `externalTrafficPolicy` set to `local`, so that ingress can filter
requests by source IP address.

```bash
helm install nginx-ingress stable/nginx-ingress \
  --namespace kube-system \
  --set controller.service.externalTrafficPolicy=Local
```

## Generate wildcard certificate for given domain

By using DNS validation execute the following command to obtain certificate.

```bash
certbot --server https://acme-v02.api.letsencrypt.org/directory \
  -d $domain --manual --preferred-challenges dns-01 \
  --cert-path . --key-path .  --fullchain-path . --chain-path . \
  --work-dir . --logs-dir . --config-dir . certonly
```

## Create certificate secret

Before any Ingress can be correctly configured to use TLS, the wildcard
certificate should be available as a secret. Apply it to `kube-system` namespace
with:

```bash
kubectl create -f k8s.3fs.si-cert.yaml -n kube-system
```

## Expose kubernetes dashboard

Apply [dashboard-ingress.yaml](./resources/dashboard-ingress.yaml) to create an ingress for the dashboard.

```bash
kubectl apply -f ./resources/dashboard-ingress.yaml
```
