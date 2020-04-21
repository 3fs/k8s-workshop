# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Uninstall a Release

## Uninstall the release

When a release is no longer needed, we can execute one command that will remove
all the k8s objects helm has created with the deployment.

To uninstall a release, use the `helm uninstall` command:

```bash
helm uninstall my-wiki
```

The output looks like this:

```console
# helm uninstall my-wiki
release "my-wiki" uninstalled
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

NAME: my-wiki
LAST DEPLOYED: Tue Apr 21 08:07:37 2020
NAMESPACE: default
STATUS: deployed
REVISION: 2
TEST SUITE: None
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
```

</details>

## Rollback

Because Helm tracks your previous releases you can audit a cluster's history,
and rollback to a previous release with `helm rollback` command.

```bash
helm history my-wiki
```

```console
# REVISION	UPDATED                 	STATUS    	CHART          	APP VERSION            	DESCRIPTION
1       	Tue Apr 21 08:07:30 2020	superseded	dokuwiki-6.0.15	0.20180422.201901061035	Install complete
2       	Tue Apr 21 08:07:37 2020	deployed  	dokuwiki-6.0.15	0.20180422.201901061035	Upgrade complete
```

To make a rollback of the release, execute the following command, which
will rollback the release to revision you have specified (in our case `1`).

```bash
helm rollback my-wiki 1
```

Output:

```console
# helm rollback my-wiki 1
Rollback was a success! Happy Helming!

# helm history my-wiki
REVISION	UPDATED                 	STATUS    	CHART          	APP VERSION            	DESCRIPTION
1       	Tue Apr 21 08:07:30 2020	superseded	dokuwiki-6.0.15	0.20180422.201901061035	Install complete
2       	Tue Apr 21 08:07:37 2020	superseded	dokuwiki-6.0.15	0.20180422.201901061035	Upgrade complete
3       	Tue Apr 21 08:12:01 2020	deployed  	dokuwiki-6.0.15	0.20180422.201901061035	Rollback to 1
```

## Remove a release

To remove a release, `helm uninstall` command can be used.

The following command will remove `my-wiki` release and all its traces from the cluster.

```bash
helm uninstall my-wiki
```

The output should is similar to basic delete command:

```console
# helm uninstall my-wiki
release "my-wiki" uninstalled
```

If you would like to remove all associated resources and mark the release as
deleted, but retain the release history, command line switch `--keep-history`
can be used. This gives the possibility of recovering old installation in case
of accidental removal.

```bash
helm uninstall --keep-history my-wiki
```

```console
# helm uninstall --keep-history my-wiki
REVISION	UPDATED                 	STATUS     	CHART          	APP VERSION            	DESCRIPTION
1       	Tue Apr 21 13:29:55 2020	superseded 	dokuwiki-6.0.11	0.20180422.201901061035	Install complete
2       	Tue Apr 21 13:31:12 2020	uninstalled	dokuwiki-6.0.15	0.20180422.201901061035	Uninstallation complete
```

The following command will show all the information about the previously
deployed, uninstalled (if `--keep-history` is used) and running deployments:

```bash
helm list --all
```

## Next: [Custom Helm](../04_custom_helm/README.md)
