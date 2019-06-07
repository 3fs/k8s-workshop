FROM ubuntu:bionic as base

# Install curl.
RUN apt-get update \
  && apt-get install -y curl git bash-completion vim-tiny nano \
  && echo 'source /etc/bash_completion' >> /etc/bash.bashrc \
  && rm -rf /var/lib/apt/lists/*

# Install kubectl.
ENV K8S_RELEASE=v1.14.1 K8S_SHA512=66c3a470caecfb35ce9a995a6298e689aed5fabefbdb8aca5086adff572266ae47b997eea03ff3ce0272fdb5be8e22aced3e3ae35906b5ac90cf928d7c0c974f

RUN curl -L -o k8s.tgz https://dl.k8s.io/${K8S_RELEASE}/kubernetes-client-linux-amd64.tar.gz \
  && echo -n "${K8S_SHA512} k8s.tgz" | sha512sum -c --quiet \
  && tar xzf k8s.tgz \
  && mv kubernetes/client/bin/kubectl /usr/local/bin/kubectl \
  && rm -r kubernetes/ k8s.tgz \
  && echo 'source <(kubectl completion bash)' >> /etc/bash.bashrc

# Install helm.
ENV HELM_RELEASE=v2.13.1 HELM_SHA256=c1967c1dfcd6c921694b80ededdb9bd1beb27cb076864e58957b1568bc98925a

RUN curl -L -o helm.tgz https://storage.googleapis.com/kubernetes-helm/helm-${HELM_RELEASE}-linux-amd64.tar.gz \
  && echo -n "${HELM_SHA256} helm.tgz" | sha256sum -c --quiet \
  && tar xzf helm.tgz \
  && mv linux-amd64/helm /usr/local/bin/helm \
  && rm -r linux-amd64/ helm.tgz \
  && echo 'source <(helm completion bash)' >> /etc/bash.bashrc \
  && /usr/local/bin/helm init --client-only \
  && /usr/local/bin/helm repo update

COPY motd /etc/motd
RUN cat /etc/motd >> /etc/bash.bashrc

# Make repository available.
RUN git clone https://github.com/3fs/k8s-workshop.git /repo

FROM base as console

# Set up kubectl config.
ARG CLUSTER_ENDPOINT
ARG USER_TOKEN
ARG USER_NAMESPACE
ENV CLUSTER_ENDPOINT=${CLUSTER_ENDPOINT} USER_TOKEN=${USER_TOKEN} USER_NAMESPACE=${USER_NAMESPACE}

# Setup Tiller namespace and workshop env variable.
ENV TILLER_NAMESPACE=${USER_NAMESPACE} CODE=${USER_NAMESPACE}

COPY ca.crt /ca.crt

RUN kubectl config set-cluster default \
    --embed-certs=true \
    --server=${CLUSTER_ENDPOINT} \
    --certificate-authority=/ca.crt \
  && kubectl config set-credentials default --token=${USER_TOKEN} \
  && kubectl config set-context default-context \
    --cluster=default \
    --user=default \
    --namespace=${USER_NAMESPACE} \
  && kubectl config use-context default-context
