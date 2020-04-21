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

REVISION	UPDATED                 	STATUS  	CHART          	APP VERSION            	DESCRIPTION
1       	Tue Apr 21 08:02:44 2020	deployed	dokuwiki-6.0.15	0.20180422.201901061035	Install complete
```

We can update the release with new configuration like this. The following will
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
#     --set "ingress.enabled=true" \
#     --set "ingress.hosts[0].name=$CODE.k8s.3fs.si" \
#     --set "ingress.hosts[0].tls=true" \
#     --set "ingress.hosts[0].tlsSecret=k8s.3fs.si-certificate" \
#     my-wiki stable/dokuwiki
Release "my-wiki" has been upgraded. Happy Helming!
NAME: my-wiki
LAST DEPLOYED: Tue Apr 21 08:06:06 2020
NAMESPACE: default
STATUS: deployed
REVISION: 3
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

After a successful deployment we can check history once again.

```console
# REVISION	UPDATED                 	STATUS    	CHART          	APP VERSION            	DESCRIPTION
1       	Tue Apr 21 08:07:30 2020	superseded	dokuwiki-6.0.15	0.20180422.201901061035	Install complete
2       	Tue Apr 21 08:07:37 2020	deployed  	dokuwiki-6.0.15	0.20180422.201901061035	Upgrade complete
```

There is a way to check which user provided values were changed by executing the following command:

```bash
helm get values my-wiki
```

```console
# helm get values my-wiki
USER-SUPPLIED VALUES:
ingress:
  enabled: true
  hosts:
  - name: $code.k8s.3fs.si
    tls: true
    tlsSecret: k8s.3fs.si-certificate
```

## Next: [Delete release](./03_delete_release.md)
