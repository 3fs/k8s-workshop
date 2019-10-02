# [Workshop](../README.md) &raquo; [Custom helm](./README.md) &raquo; create helm chart

In this task you will learn how to create a simple helm chart and update it to
use it with docker image prepared for the workshop.

## Create helm chart

`helm` provides capability of creating simple helm chart with definitions
of Deployment, Service and Ingress kubernetes resources.

To create a simple helm chart execute the following command.

```bash
helm create workshop-chart
```

The newly create directory `workshop-chart` includes:

```bash
workshop-chart
|-- Chart.yaml
|-- charts
|-- templates
|   |-- NOTES.txt
|   |-- _helpers.tpl
|   |-- deployment.yaml
|   |-- ingress.yaml
|   `-- service.yaml
`-- values.yaml
```

The structure of this directory is set up in advance and allows the developer to
create a custom chart with all the prerequisites already in place:

- `Chart.yaml` and `NOTES.txt` files are documentation files for the chart
  itself
- `charts` directory stores all the required dependencies (if there are any)
- `templates` directory includes predefined templates, which will be rendered to
  valid yaml files defining different kubernetes resources (in this case
  Deployment, Service, and Ingress)
- `values.yaml` define different variables, which will be rendered into yaml
  template files and sent to k8s

## Edit values file

Every helm chart requires `values.yaml` file, which defines the variables used
later in the definitions of k8s resources (Deployment, Service, Ingress). The
default helm chart `values.yaml` includes definitions of:

- [used images](#image-config)
- Ingress resource definitions
- resource utilization definitions (these parts will be untouched)

### Image config

The image dictionary part of `values.yaml` file defines container images, which
will be used in k8s Deployment.

To use workshop image the values needs to be updated accordingly:

1. update the repository name: `eu.gcr.io/k8s-workshop-skopje/k8s-workshop`
2. update the tag name to `latest`

The result should be similar to this:

```yaml
image:
  repository: eu.gcr.io/k8s-workshop-skopje/k8s-workshop
  tag: latest
  pullPolicy: IfNotPresent
```

After this has been configured, we can use `helm` to deploy our chart to k8s
cluster.

```bash
helm install workshop-chart --name my-project
helm list
```

## Demo

[![asciicast](https://asciinema.org/a/0WVHLySWODgjdOIuqnZFoLH57.svg)](https://asciinema.org/a/0WVHLySWODgjdOIuqnZFoLH57)

## Next: [Enable Ingress resource](./02_enable_ingress.md)
