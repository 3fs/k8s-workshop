# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Inspect a Release

Every release has its own release history. For every release with the same
name, helm creates a new release version and deploys it. If the deployment
failes, rollback to the previous version is performed.

To inspect the release history execute the following command:

```bash
helm history my-wiki
```

```console
# helm history my-wiki
REVISION    UPDATED                     STATUS      CHART             DESCRIPTION
1           Thu Jun  6 07:35:58 2019    DEPLOYED    dokuwiki-4.3.1    Install complete
```

We can upgrade the release with new configuration like this. The following will
redeploy DokuWiki and open it to outside world on the address `https://$CODE.k8s.3fs.si`.

```bash
helm upgrade \
    --set "ingress.enabled=true" \
    --set "ingress.hosts[0].name=$CODE.k8s.3fs.si" \
    --set "ingress.hosts[0].tls=true" \
    --set "ingress.hosts[0].tlsSecret=k8s.3fs.si-certificate" \
    my-wiki stable/dokuwiki
```

<details>
    <summary>Show output</summary>

```console
# helm upgrade \
#    --set "ingress.enabled=true" \
#    --set "ingress.hosts[0].name=$CODE.k8s.3fs.si" \
#    --set "ingress.hosts[0].tls=true" \
#    --set "ingress.hosts[0].tlsSecret=k8s.3fs.si-certificate" \
#    my-wiki stable/dokuwiki

Release "my-wiki" has been upgraded.
LAST DEPLOYED: Mon Jun 10 08:53:18 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Pod(related)
NAME                            READY  STATUS   RESTARTS  AGE
my-wiki-dokuwiki-57c79d9-4zx6b  1/1    Running  0         3m22s

==> v1/Secret
NAME              TYPE    DATA  AGE
my-wiki-dokuwiki  Opaque  1     3m22s

==> v1/Service
NAME              TYPE          CLUSTER-IP    EXTERNAL-IP  PORT(S)                     AGE
my-wiki-dokuwiki  LoadBalancer  10.106.72.55  localhost    80:32526/TCP,443:31330/TCP  3m22s

==> v1beta1/Deployment
NAME              READY  UP-TO-DATE  AVAILABLE  AGE
my-wiki-dokuwiki  1/1    1           1          3m22s


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

After a successful deployment we can check history once again.

```console
# helm history my-wiki
REVISION    UPDATED                     STATUS        CHART             DESCRIPTION
1           Thu Jun  6 07:35:58 2019    SUPERSEDED    dokuwiki-4.3.1    Install complete
2           Thu Jun  6 07:37:22 2019    DEPLOYED      dokuwiki-4.3.1    Upgrade complete
```

## Next: [Delete release](./03_delete_release.md)
