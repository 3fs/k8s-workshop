# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Install Helm

## Initalize Helm on k8s cluster

To initalize Tiller in a custom namespace, either `TILLER_NAMESPACE`
environnment variable or `--tiller-namespace` flag have to be set. The former is
preset in the `console` container, so the command for installing Tiller is the
following.

```bash
helm init --service-account=$CODE
```

The output should look like this:

```console
root@5796aa6a1f9a:/# helm init --service-account=$CODE
Creating /root/.helm
Creating /root/.helm/repository
Creating /root/.helm/repository/cache
Creating /root/.helm/repository/local
Creating /root/.helm/plugins
Creating /root/.helm/starters
Creating /root/.helm/cache/archive
Creating /root/.helm/repository/repositories.yaml
Adding stable repo with URL: https://kubernetes-charts.storage.googleapis.com
Adding local repo with URL: http://127.0.0.1:8879/charts
$HELM_HOME has been configured at /root/.helm.

Tiller (the Helm server-side component) has been installed into your Kubernetes Cluster.

Please note: by default, Tiller is deployed with an insecure 'allow unauthenticated users' policy.
To prevent this, run `helm init` with the --tiller-tls-verify flag.
For more information on securing your installation see: https://docs.helm.sh/using_helm/#securing-your-helm-installation
Happy Helming!
```

## Next: [Install charts](./01_install_chart.md)
