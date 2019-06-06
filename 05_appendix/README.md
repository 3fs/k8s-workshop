# [Workshop](../README.md) &raquo; Appendix

## Advanced helm chart

This section contains instructions to set up a
[Wordpress](https://wordpress.org/) deployment using Helm. The intent is to
demonstrate usage of many Kubernetes features in a Helm deployment at once.

Start out by creating a custom [values.yaml](./wordpress-values.yaml) file with
custom configuration. All adjusted options are explained in that file.

Install the chart while providing custom `values.yaml` to override the default
settings and set ingress hostname.

```bash
helm install \
    --name workshop-wp \
    -f wordpress-values.yaml \
    --set "ingress.hosts[0].name=wp-${CODE}.k8s.3fs.si" \
    --set "ingress.tls[0].hosts[0]=wp-${CODE}.k8s.3fs.si" \
    --wait \
    stable/wordpress
```

Open the deployed web page at `https://wp-${CODE}.k8s.3fs.si`

Afterwards, you can delete the deployment with

```bash
helm delete --purge workshop-wp
```
