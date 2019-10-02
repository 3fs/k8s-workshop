# [Workshop](../README.md) &raquo; k8s introduction

## Prerequisites

For the purpose of this workshop we have pre-created a kubernetes cluster in
Google Cloud and created multiple namespaces, named by the `$CODE`. Each
participant has access to one of the namespaces, where all the kubernetes
related tasks can be performed.

To simplify access to the cluster, we have created docker images with already
baked-in credentials to access your specific namespace. The image has
preconfigured `.kubeconfig` and all the required environmental variables set.

### Access to your namespace in workshop k8s cluster

Run the below command to start a docker container with credentials to access
your test namespace in a kubernetes cluster.

Substitute `${CODE}` with workshop code you received as part of printed
instructions.

```bash
docker run \
    -it \
    eu.gcr.io/k8s-workshop-skopje/console:${CODE}
```

If you would like to edit file locally and use them in docker container, you can
do that by downloading the workshop repository from github and mount it into
docker container. Replace `${PWD}` with path to [downloaded github
repository](../00_prerequisites/README.md#github-repository). This will mount
the repository inside container, so you will be able to edit files locally and
then apply changes using credentials inside the container.

```bash
docker run \
    -v "${PWD}":/repo \
    -it \
    eu.gcr.io/k8s-workshop-skopje/console:${CODE}
```

Users of Windows operating system may expirience issues with accessing the
repository directory. This can be solved like
[this](https://token2shell.com/howto/docker/sharing-windows-folders-with-containers/).

### Access to kubernetes dashboard

The workshop kubernetes cluster has [kubernetes
dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
enabled. To access the dashboard:

1. open [kubernetes dashboard](https://dashboard.k8s.3fs.si)
2. login using token which can be obtained from the `console` container with
   command `login-token`
3. enter the name of your namespace into the input field on the dashboard's left
   column

## Tasks

### Create a Deployment and Service objects

One of the commands for creating and updating kubernetes objects is `kubectl
apply`. Use it to create a Deployment resource named `my-deployment`, which will
run a sample container and a Service resource named `my-service`, which
abstracts access to `my-deployment` in the cluster.

```bash
kubectl apply -f /repo/02_kubernetes/files/deployment.yaml
kubectl apply -f /repo/02_kubernetes/files/service.yaml
```

Example output:

```console
# kubectl apply -f /repo/02_kubernetes/files/deployment.yaml
deployment.apps/my-deployment created
# kubectl apply -f /repo/02_kubernetes/files/service.yaml
service/my-service created
```

### Create the certificate secret

A .yaml file to create the wildcard certificate secret is located in the `/`
folder of Docker container. This object will be important in the future tasks.
Apply it using the following command:

```bash
kubectl create -f /k8s.3fs.si-cert.yaml
```

Example output:

```console
# kubectl create -f /k8s.3fs.si-cert.yaml
secret/k8s.3fs.si-certificate created
```

### Create an Ingress object

First off, edit hostnames in `/repo/02_kubernetes/files/ingress.yaml` to include
`$CODE` (e.g. `cranky-hippo.k8s.3fs.si`). The `spec` section should look similar
to:

```yaml
tls:
  - hosts:
      - cranky-hippo.k8s.3fs.si
    secretName: k8s.3fs.si-certificate
rules:
  - host: cranky-hippo.k8s.3fs.si
    http:
      paths:
        - path: "/"
          backend:
            serviceName: my-service
            servicePort: http
```

Create the Ingress object named `my-ingress` using the following command.

```bash
kubectl apply -f /repo/02_kubernetes/files/ingress.yaml
```

The deployment should now be available at the configured hostname. Example
output:

```console
# kubectl apply -f /repo/02_kubernetes/files/ingress.yaml
ingress.extensions/my-ingress created
```

### Show all resources

To list all configured resources, use `kubetl get <resource_type>`.

```bash
kubectl get all
```

Example output:

```console
# kubectl get all
NAME                                 READY   STATUS    RESTARTS   AGE
pod/my-deployment-cb847b564-w78t8    1/1     Running   0          4m

NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)     AGE
service/my-service      ClusterIP   10.39.243.107   <none>        80/TCP      2m

NAME                            DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/my-deployment   1         1         1            1           4m

NAME                                       DESIRED   CURRENT   READY   AGE
replicaset.apps/my-deployment-cb847b564    1         1         1       4m
```

### Show deployment logs

We can inspect text output made by any kubernetes resource using `kubectl logs`.
Following is an example of logs for previously deployed `my-deployment`.
Resource names can be found using `kubectl get all`.

```bash
kubectl logs deployment.apps/my-deployment
```

Example output:

```console
# kubectl logs deployment.apps/my-deployment
http: 2019/05/06 12:29:38 Server is starting on :80
http: 2019/05/06 12:32:08 GET hippo.k8s.3fs.si / 10.36.2.11:33654 Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:66.0) Gecko/20100101 Firefox/66.0
http: 2019/05/06 12:32:09 GET hippo.k8s.3fs.si /favicon.ico 10.36.2.11:33654 Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:66.0) Gecko/20100101 Firefox/66.0
```

### Delete created resources

Kubernetes objects can be removed using `kubectl delete` command.

```bash
kubectl delete ingress my-ingress
kubectl delete service my-service
kubectl delete deployment my-deployment
```

Example output:

```console
# kubectl delete ingress my-ingress
ingress.extensions "my-ingress" deleted
# kubectl delete service my-service
service "my-service" deleted
# kubectl delete deployment my-deployment
deployment.extensions "my-deployment" deleted
```

## Next: [Helm](../03_helm/README.md)
