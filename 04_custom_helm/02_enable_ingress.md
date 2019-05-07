# [Workshop](../README.md) &raquo; [Custom helm](./README.md) &raquo; Enable ingress

Update the helm chart to expose the service using Ingress resource.

## Ingress config

Ingress part of the `values.yaml` file defines the hostnames of the exposed
Ingress service, the URI part of the service and k8s Secret resource with the
appropriate SSL certificate.

To expose the containers configured in the previous section, configure the
ingress dictionary like this:

- set `enabled` to `true` to expose the service
- set host to:  `$CODE`.`k8s.3fs.si` (e.g. `cranky-hippo.k8s.3fs.si`)
- populate the paths list with root URI - `/` so all the request to this URL
  will be served by this Ingress
- set the tls part to use the Secret with name `k8s.3fs.si-certificate` (which
  is already defined in k8s, since we have imported it in the previous tasks)
- set the hosts part in the `tls` dictionary to `$CODE.k8s.3fs.si` (e.g.
  `cranky-hippo.k8s.3fs.si`)

The end result of ingress dictionary should look similar to this:

```yaml
ingress:
  enabled: true
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: cranky-hippo.k8s.3fs.si
      paths:
        - /

  tls:
    - secretName: k8s.3fs.si-certificate
      hosts:
        - cranky-hippo.k8s.3fs.si
```

Since we have updated the chart, we need to upgrade the currently deployed
chart:

```bash
helm upgrade my-project workshop-chart
helm history my-project
```

## Demo

[![asciicast](https://asciinema.org/a/jdVnwcppufwU5hZPSwk01bH0P.svg)](https://asciinema.org/a/jdVnwcppufwU5hZPSwk01bH0P)

## Next: [Add custom variables](./03_add_variables.md)
