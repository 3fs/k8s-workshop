# [Workshop](../README.md) &raquo;[Workshop setup](./README.md) &raquo;Server container

## Build server container

The [Dockerfile](./resources/Dockerfile) builds http server located in `./code`.
The server response to any request will include the workshop ID.

Build the image.

```bash
docker build -t k8s-workshop -f ./resources/Dockerfile .
```

Log in to container registry.

```bash
gcloud auth configure-docker
```

Tag the image.

```bash
docker tag k8s-workshop eu.gcr.io/k8s-workshop-skopje/k8s-workshop:stable
docker tag k8s-workshop eu.gcr.io/k8s-workshop-skopje/k8s-workshop:latest
```

Push the image to registry.

```bash
docker push eu.gcr.io/k8s-workshop-skopje/k8s-workshop:stable
docker push eu.gcr.io/k8s-workshop-skopje/k8s-workshop:latest
```
