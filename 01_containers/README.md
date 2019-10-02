# [Workshop](../README.md) &raquo; container basics

This set of tasks covers the basic operations on containers with Docker. Here
you will learn:

* [How to check docker status](#how-to-check-docker-status)
* [How to run a simple container](#run-a-simple-container)
* [How to expose a container network port to the host
  OS](#expose-a-local-container-network-port)
* [How to use environment variables](#how-to-use-environment-variables)
* [How to expose local directory to a
  container](#expose-a-local-directory-to-container)
* [Various other docker commands](#various-docker-commands)

After you have installed docker, open your favourite command line terminal
(Powershell, iTerm, Terminal, ...) and execute the following command.

## How to check docker status

```bash
docker ps --all
```

> The previous command will show all the running containers on the host OS. Some
> of them might be stopped, running or dead. This is one of the basic `docker`
> commands to check the current status of containers within your operating
> system. If there are no stopped or running containers, the list will be empty.

<details>
    <summary>Show demo</summary>

[![asciicast](https://asciinema.org/a/ZiIMl4QrCuY18kBnw9M2q8r9o.svg)](https://asciinema.org/a/ZiIMl4QrCuY18kBnw9M2q8r9o)

</details>

## Run a simple container

```bash
docker run -it --name ubuntu-container ubuntu:latest
```

This will download an docker image with latest Ubuntu linux distribution, start
it and open an interactive terminal session. After that you will have an
interactive session within Ubuntu Linux operating system. When you are done with
your work, you can exit the container by executing `exit` command.

The container will go into stopped status (you can check the status of
containers with `docker ps --all` as discussed in this
[section](#how-to-check-docker-status).

By running the command below you can inspect the container and get the detailed
information about the running or stopped container.

```bash
docker inspect ubuntu-container
```

<details>
    <summary>Show demo</summary>

[![asciicast](https://asciinema.org/a/CUa8mWXTxYCu4P6W0T5zdmXjp.svg)](https://asciinema.org/a/CUa8mWXTxYCu4P6W0T5zdmXjp)

</details>

## Expose a local container network port

Container can have network services (web servers, mail servers, ...) running
inside them.

The following command will start a container, expose its TCP network port 80 to
localhost port 8080 (`-p 127.0.0.1:8080:80`), put it in the background
(`--detach`) and named it (`--name k8s-workshop-container`).

You can access the container with your browser [here](http://127.0.0.1:8080).

```bash
docker run \
    -p 127.0.0.1:8080:80 \
    --detach \
    --name k8s-workshop-hello \
    eu.gcr.io/k8s-workshop-skopje/k8s-workshop:latest
```

> The container image was prepared in advance and pushed into Google cloud
> repository.

The container is now running in the background and acting as an HTTP server,
listening on `localhost:8080`.

You can check the logs of the container by executing:

```bash
docker logs k8s-workshop-hello
```

Now the container is running in the background (check status with `docker ps
--all`).

After you are done, stop it with the following command.

```bash
docker stop k8s-workshop-hello
```

## How to use environment variables

You can pass environment variables to running container by specifying them in
command line (`--env`). In the example below, the variable `CODE` is used by the
container and displayed when you open the browser on [this
address](http://127.0.0.1/hello).

```bash
docker run \
    -p 127.0.0.1:8080:80 \
    --detach \
    --name k8s-workshop-env \
    --env CODE=k8s-workshop \
    eu.gcr.io/k8s-workshop-skopje/k8s-workshop:latest
```

After you are done, stop it with the following command.

```bash
docker stop k8s-workshop-env
```

<details>
    <summary>Show demo</summary>

[![asciicast](https://asciinema.org/a/K82r5UFM76H8fP7l2b1D2ijpP.svg)](https://asciinema.org/a/K82r5UFM76H8fP7l2b1D2ijpP)

</details>

## Expose a local directory to container

Docker provides a capability to expose (mount) a file or directory to container.
In the example below, we will mount a local directory to a specific directory in
the container. This will be used in the http server of the container.

```bash
docker run \
    -p 127.0.0.1:8080:80 \
    --detach \
    --name k8s-workshop-directory \
    --mount type=bind,source=/Users/k8s-workshop/,target=/tmp \
    eu.gcr.io/k8s-workshop-skopje/k8s-workshop:latest
```

After you are done, stop it with the following command.

```bash
docker stop k8s-workshop-directory
```

## Various docker commands

### Remove stopped containers

```bash
docker rm ubuntu-container
```

## Docker images

Each container image is stored on the disk and can be reused.

```bash
docker image ls
```

You can delete the container image by executing the following command. (e.g.
`ubuntu:latest`). You can only remove a container image, which is currently not
in use.

```bash
docker image rm ubuntu:latest
```

### Docker system cleanups

```bash
docker system prune
```

> Note: The command may also remove other container images that are not related
> to the workshop. For more details consult the
> [manual](https://docs.docker.com/engine/reference/commandline/system_prune/).

<details>
    <summary>Show demo</summary>

[![asciicast](https://asciinema.org/a/Mr98aIaF5xMEi8IYNwfktRg7c.svg)](https://asciinema.org/a/Mr98aIaF5xMEi8IYNwfktRg7c)

</details>

## Next: [Kubernetes](../02_kubernetes/README.md)
