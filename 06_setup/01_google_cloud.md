# [Workshop](../README.md) &raquo;[Workshop setup](./README.md) &raquo;Set up Google Cloud

## Google cloud setup

Set up a [free trial](https://console.cloud.google.com/freetrial) of Google
Cloud. This includes filling out the billing form and creating a project.

To use the command line `gcloud` utility, Google Cloud SDK needs to be
installed. The instruction how to do that, can be found
[here](https://cloud.google.com/sdk/install).

### Authenticate to Google Cloud

Log in to Google Cloud using `gcloud`.

```bash
gcloud auth login
```

### Set the default project

Set default project with the following command (e.g. `k8s-workshop-skopje`)

```bash
gcloud config set project k8s-workshop-skopje
```

### Enable the container registry

Enable container registry by executing the following command:

```bash
gcloud services enable containerregistry.googleapis.com
```

Pushing and pulling container images is now possible. The registry address we
use is `eu.gcr.io/k8s-workshop-skopje`.

