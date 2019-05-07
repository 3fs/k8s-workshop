# [Workshop](../README.md) &raquo; [Custom helm](./README.md) &raquo; Additional variables

In this task you will learn how to add custom variables to helm chart and
reference them in resources.

## Additional variables

We can add the additional values to `values.yaml` file and reference them in the
templates.

For the workshop purposes, we will add a dictionary with the workshop `CODE`
value to reference it in the chart templates.

### Add values to `values.yaml`

Append the following block to the `values.yaml` file with the appropriate `CODE`
value:

```yaml
workshop:
  code: cranky-hippo
```

This block adds variable `workshop.code` to templating engine and can be
referenced in the template files as `{{ .Values.workshop.code }}`.

### Edit the Deployment resource

Kubernetes work with environment variables in a similar way as docker.

To expose an environment variable to a pod we must add `env` dictionary to pod
resource definition in the Deployment file. This can be achived by adding the
following after the `imagePullPolicy` key in `templates/deployment.yaml` file.

> The indentation should be alinged with `imagePullPolicy` since it is part of
> the container definition.

```yaml
env:
  - name: CODE
    value: "{{ .Values.workshop.code }}"
```

<details>
    <summary>The end result should look like this</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "workshop-chart.fullname" . }}
  labels:
    app.kubernetes.io/name: {{ include "workshop-chart.name" . }}
    helm.sh/chart: {{ include "workshop-chart.chart" . }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app.kubernetes.io/name: {{ include "workshop-chart.name" . }}
      app.kubernetes.io/instance: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{ include "workshop-chart.name" . }}
        app.kubernetes.io/instance: {{ .Release.Name }}
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          env:
            - name: CODE
              value: "{{ .Values.workshop.code }}"
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
    {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
    {{- end }}
    {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
    {{- end }}
```

</details>

After the `templates/deployment.yaml` file has been updated, we should redeploy
our chart with additional flag to recreate pods with new definitions.

```bash
helm upgrade my-project workshop-chart --recreate-pods
helm history my-project
```

You can check the result by opening your web page at your configured URL:
`https://${CODE}.k8s.3fs.si/hello`

## Demo

[![asciicast](https://asciinema.org/a/Vj4LzahOfCw8TkNLgrKmIIXFS.svg)](https://asciinema.org/a/Vj4LzahOfCw8TkNLgrKmIIXFS)

## Next: [Appendix](../05_appendix/README.md)
