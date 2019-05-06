# [Workshop](../README.md) &raquo; [Custom helm](./README.md) &raquo; Configure health probes

In this task you will learn how to update health probes of a helm chart.

## Edit health checks

In the default helm template, health checks are already enabled. we just need to
adjust the paths used. In file templates/deployment.yaml we must edit the `path`
field:

In the `templates/deployment.yaml` file search for dictionary keys
`livenessProbe` and `readinessProbe` and update the paths to `/live` and
`/ready` respectively.

The referenced paths are preconfigured in the pre-created container and respond
to HTTP `GET` request with response `200`. This tells the k8s that the pod is
ready to start serving the requests and can be used in the Service and later on
in the Ingress.

The result should look like this:

```yaml
livenessProbe:
  httpGet:
    path: /live
    port: http
readinessProbe:
  httpGet:
    path: /ready
    port: http
```

<details>
    <summary>The whole `templates/deployment.yaml` file should look like this</summary>

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
              path: /live
              port: http
          readinessProbe:
            httpGet:
              path: /ready
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

## Demo

[![asciicast](https://asciinema.org/a/H0sgtulxJfj8K9lNPNcA1PW1J.svg)](https://asciinema.org/a/H0sgtulxJfj8K9lNPNcA1PW1J)
