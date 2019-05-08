# [Workshop](../README.md) &raquo;[Workshop setup](./README.md) &raquo;Create client containers

## Workshop docker console setup

Build a container from [the Dockerfile](./resources/console/Dockerfile). It's
done automatically as part of the following workshop account creation.

## Workshop accounts creation

Prerequisites:

- logged into gcloud ([instructions](./01_google_cloud.md))
- `kubectl` and `docker` available and set up
- `awk` and `jq` must be installed
- for name generation `curl` and `go` runtime are also needed

Execute:
```bash
cd ./resources/console && ./create-users.bash 50
```
