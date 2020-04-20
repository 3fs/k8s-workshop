# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Install an Example Chart

## Task objective

In this task, we will install DokuWiki. DokuWiki is a simple to use and highly
versatile Open Source wiki software. We will use helm to deploy it and expose it
to the Internet using Ingress kubernetes object.

## Install an Example Chart

To install a chart, you can run the `helm install` command. Helm has several
ways to find and install a chart, but the easiest is to use one of the `bitnami`
charts.

In the example below, the `bitnami` helm chart repository will be added and
`bitnami/dokuwiki` chart will be installed, and the name of our new release is
`my-wiki`.

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

```bash
helm install my-wiki bitnami/dokuwiki
```

This command created the objects and send them to k8s and the output describes,
which objects were created and show some basic information about the
installation.

<details>
    <summary>Show output</summary>

```console
# helm repo add bitnami https://charts.bitnami.com/bitnami
"bitnami" has been added to your repositories

# helm install my-wiki bitnami/dokuwiki
NAME: my-wiki
LAST DEPLOYED: Tue Apr 21 08:02:44 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
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
NAME   	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART          	APP VERSION
my-wiki	default  	1       	2020-04-21 08:02:44.32098 +0200 CEST	deployed	dokuwiki-6.0.15	0.20180422.201901061035
```

## Next: [Inspect a Release](./02_inspect_upgrade.md)
