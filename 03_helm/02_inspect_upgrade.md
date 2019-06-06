# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Inspect a Release

Every release has its own release history. For every release with the same
name, helm creates a new release version and deploys it. If the deployment
failes, rollback to the previous version is performed.

To inspect the release history execute the following command:

```bash
helm history my-wiki
```

We can upgrade the release with new configuration like this. The following will
redeploy DokuWiki and open it to outside world on the address `https://$CODE.k8s.3fs.si`.

```bash
➜  helm upgrade \
    --set "ingress.enabled=true" \
    --set "ingress.hosts[0].name=$CODE.k8s.3fs.si" \
    --set "ingress.hosts[0].tls=true" \
    --set "ingress.hosts[0].tlsSecret=k8s.3fs.si-certificate" \
    my-wiki stable/dokuwiki
```

After a successful deployment we can check history once again.

```bash
helm history my-wiki
REVISION    UPDATED                     STATUS        CHART             DESCRIPTION
1           Thu Jun  6 07:35:58 2019    SUPERSEDED    dokuwiki-4.3.1    Install complete
2           Thu Jun  6 07:37:22 2019    DEPLOYED      dokuwiki-4.3.1    Upgrade complete
```

## Next: [Delete release](./03_delete_release.md)
