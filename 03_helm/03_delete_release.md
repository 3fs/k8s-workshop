# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Uninstall a Release

## Uninstall the release

When a release is no longer needed, we can execute one command that will remove
all the k8s objects helm has created with the deployment.

To uninstall a release, use the `helm delete` command:

```bash
helm delete my-wiki
```

The output looks like this:

```console
# helm delete my-wiki
release "my-wiki" deleted
```

This will uninstall `my-wiki` from Kubernetes, but you will
still be able to request information about that release.

```bash
helm history my-wiki
```

```console
# helm history my-wiki
REVISION  UPDATED                   STATUS      CHART           DESCRIPTION
1         Mon Jun 10 08:49:55 2019  SUPERSEDED  dokuwiki-4.3.1  Install complete
2         Mon Jun 10 08:53:18 2019  DELETED     dokuwiki-4.3.1  Deletion complete
```

Additionally, status of the deployment can be checked:

```bash
helm status my-wiki
```

<details>
    <summary>The output of this command should be like this. Note the `STATUS: DELETED`.</summary>

```console
# helm status my-wiki
LAST DEPLOYED: Thu Jun  6 09:25:24 2019
NAMESPACE: default
STATUS: DELETED

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

## Rollback

Because Helm tracks your releases even after you've deleted them, you
can audit a cluster's history, and even undelete a release with `helm
rollback` command.

```bash
helm history my-wiki
```

```console
# helm history my-wiki
REVISION  UPDATED                   STATUS      CHART           DESCRIPTION
1         Mon Jun 10 08:49:55 2019  SUPERSEDED  dokuwiki-4.3.1  Install complete
2         Mon Jun 10 08:53:18 2019  DELETED     dokuwiki-4.3.1  Deletion complete
```

To make a undelete/rollback of the release, execute the following command, which
will rollback the release to revision you have specified (in our case `2`).

```bash
helm rollback my-wiki 2
```

Output:

```console
# helm rollback my-wiki 2
Rollback was a success! Happy Helming!

# helm history my-wiki
REVISION    UPDATED                     STATUS        CHART             DESCRIPTION
1           Thu Jun  6 07:35:58 2019    SUPERSEDED    dokuwiki-4.3.1    Install complete
3           Thu Jun  6 07:42:58 2019    SUPERSEDED    dokuwiki-4.3.1    Deletion complete
4           Thu Jun  6 07:44:29 2019    DEPLOYED      dokuwiki-4.3.1    Rollback to 2
```

## Purge a release

To totally remove helm deployed release from the cluster `--purge` command line option should be used with `delete`.

The following command will remove `my-wiki` release and all its traces from the cluster.

```bash
helm delete --purge my-wiki
```

The output should is similar to basic delete command:

```console
# helm delete --purge my-wiki
release "my-wiki" deleted
```

## Next: [Custom Helm](../04_custom_helm/README.md)
