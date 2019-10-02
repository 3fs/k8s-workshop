#!/usr/bin/env bash

set -eu

function create_account() {
    local name cluster_endpoint secret_name user_token
    name=$1
    cluster_endpoint=$2

    # Ensure namespace and service account are present.
    if ! kubectl get namespace "$name" > /dev/null 2>&1; then
        kubectl create namespace "$name"
        kubectl create serviceaccount "$name" --namespace "$name"

        # Bind "edit" role to account in its own namespace.
        kubectl create rolebinding "${name}-role-binding" --clusterrole=edit --serviceaccount="${name}:${name}" --namespace="$name"
    fi


    # Fetch cert.
    secret_name=$(kubectl get serviceaccount "$name" --namespace "$name" -o json | jq -r .secrets[].name)
    kubectl get secret "$secret_name" --namespace "$name" -o json | jq -r '.data["ca.crt"] | @base64d' > ca.crt

    # Fetch user token.
    user_token=$(kubectl get secret "$secret_name" --namespace "$name" -o json | jq -r '.data["token"] | @base64d')


    # Build the container.
    docker build \
           -t "eu.gcr.io/k8s-workshop-skopje/console:${name}" \
           --build-arg CLUSTER_ENDPOINT="${cluster_endpoint}" \
           --build-arg USER_NAMESPACE="${name}" \
           --build-arg USER_TOKEN="${user_token}" \
           .

    # Push the container to registry.
    docker push "eu.gcr.io/k8s-workshop-skopje/console:${name}"

    # Clean up.
    rm ca.crt
}

function get_cluster_endpoint() {
    local context cluster_name
    context=$(kubectl config current-context)
    cluster_name=$(kubectl config get-contexts "$context" | awk 'END {print $3}')
    kubectl config view -o jsonpath="{.clusters[?(@.name == \"${cluster_name}\")].cluster.server}"
}

# Create service account names.
function create_names() {
    echo "it's a secret"
}

[ ! -f names.txt ] && create_names "$1"

CLUSTER_ENDPOINT="$(get_cluster_endpoint)"
while read -r name; do
    create_account "$name" "$CLUSTER_ENDPOINT"
done < names.txt
