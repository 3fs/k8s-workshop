# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Install an Example Chart

## Task objective

In this task, we will install DokuWiki. DokuWiki is a simple to use and highly
versatile Open Source wiki software. We will use helm to deploy it and expose it
to the Internet using Ingress kubernetes object.

## Install an Example Chart

To install a chart, you can run the `helm install` command. Helm has several
ways to find and install a chart, but the easiest is to use one of the official
`stable` charts.

In the example below, the `stable/dokuwiki` chart will be released, and the
name of our new release is `my-wiki`.

```bash
helm install --name my-wiki stable/dokuwiki
```

This command created the objects and send them to k8s and the output describes,
which objects were created and show some basic information about the
installation.

<details>
    <summary>Show output</summary>

```console
# helm install --name my-wiki stable/dokuwiki

NAME:   my-wiki
LAST DEPLOYED: Thu Jun  6 07:35:58 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/PersistentVolumeClaim
NAME                       STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
my-wiki-dokuwiki-dokuwiki  Bound   pvc-f65469bf-881c-11e9-b0e6-025000000001  8Gi       RWO           hostpath      0s

==> v1/Pod(related)
NAME                               READY  STATUS             RESTARTS  AGE
my-wiki-dokuwiki-8599cdc775-kr6wh  0/1    ContainerCreating  0         0s

==> v1/Secret
NAME              TYPE    DATA  AGE
my-wiki-dokuwiki  Opaque  1     0s

==> v1/Service
NAME              TYPE          CLUSTER-IP     EXTERNAL-IP  PORT(S)                     AGE
my-wiki-dokuwiki  LoadBalancer  10.109.34.250  <pending>    80:31749/TCP,443:32051/TCP  0s

==> v1beta1/Deployment
NAME              READY  UP-TO-DATE  AVAILABLE  AGE
my-wiki-dokuwiki  0/1    1           0          0s


NOTES:

** Please be patient while the chart is being deployed **

1. Get the DokuWiki URL by running:

** Please ensure an external IP is associated to the my-wiki-dokuwiki service before proceeding **
** Watch the status using: kubectl get svc --namespace default -w my-wiki-dokuwiki **

  export SERVICE_IP=$(kubectl get svc --namespace default my-wiki-dokuwiki --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
  echo "URL: http://$SERVICE_IP/"

2. Login with the following credentials

  echo Username: user
  echo Password: $(kubectl get secret --namespace default my-wiki-dokuwiki -o jsonpath="{.data.dokuwiki-password}" | base64 --decode)

```

</details>

## Learn About Releases

You can list the deployed releases by executing the following command.

```console
# helm list
NAME           REVISION    UPDATED                     STATUS      CHART                   APP VERSION                NAMESPACE
my-wiki        1           Thu Jun  6 07:16:54 2019    DEPLOYED    dokuwiki-4.3.1          0.20180422.201901061035    cranky-hippo
```

## Next: [Inspect a Release](./02_inspect_upgrade.md)
