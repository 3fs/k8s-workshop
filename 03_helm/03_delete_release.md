# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Uninstall a Release

To uninstall a release, use the `helm delete` command:

```bash
helm delete my-wiki
```

This will uninstall `my-wiki` from Kubernetes, but you will
still be able to request information about that release:

```bash
helm status my-wiki
```

Because Helm tracks your releases even after you've deleted them, you
can audit a cluster's history, and even undelete a release (with `helm
rollback`).

```bash
helm history my-wiki
helm rollback my-wiki 1
```

Output of these commands can be seen here:

```console
# helm delete my-wiki
release "my-wiki" deleted
# helm status my-wiki
LAST DEPLOYED: Thu Jun  6 07:42:58 2019
NAMESPACE: default
STATUS: DELETED

NOTES:

** Please be patient while the chart is being deployed **

1. Get the DokuWiki URL indicated on the Ingress Rule and associate it to your cluster external IP:

   export CLUSTER_IP=$(minikube ip) # On Minikube. Use: `kubectl cluster-info` on others K8s clusters
   export HOSTNAME=$(kubectl get ingress --namespace default my-wiki-dokuwiki -o jsonpath='{.spec.rules[0].host}')
   echo "Dokuwiki URL: http://$HOSTNAME/"
   echo "$CLUSTER_IP  $HOSTNAME" | sudo tee -a /etc/hosts

2. Login with the following credentials

  echo Username: user
  echo Password: $(kubectl get secret --namespace default my-wiki-dokuwiki -o jsonpath="{.data.dokuwiki-password}" | base64 --decode)

# helm rollback my-wiki 1
Rollback was a success! Happy Helming!
# helm history my-wiki
REVISION    UPDATED                     STATUS        CHART             DESCRIPTION
1           Thu Jun  6 07:35:58 2019    SUPERSEDED    dokuwiki-4.3.1    Install complete
2           Thu Jun  6 07:37:22 2019    SUPERSEDED    dokuwiki-4.3.1    Upgrade complete
3           Thu Jun  6 07:42:58 2019    SUPERSEDED    dokuwiki-4.3.1    Deletion complete
4           Thu Jun  6 07:44:29 2019    DEPLOYED      dokuwiki-4.3.1    Rollback to 1
```

## Next: [Custom Helm](../04_custom_helm/README.md)
