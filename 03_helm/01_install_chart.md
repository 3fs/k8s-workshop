# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Install an Example Chart

## Install an Example Chart

To install a chart, you can run the `helm install` command. Helm has several
ways to find and install a chart, but the easiest is to use one of the official
`stable` charts.

In the example below, the `stable/postgresql` chart will be released, and the
name of our new release is `my-database`.

```bash
helm repo update
helm install --set persistence.enabled=false --name my-database stable/postgresql
```

The newly deployed release has `persistence.enabled` set to false to skip the
creation of Google cloud disk, where PostgreSQL would tipically store database
data. The rest of the supported values can be seen on
[github](https://github.com/helm/charts/tree/master/stable/postgresql).

The output should look like this:

```console
➜  helm repo update
Hang tight while we grab the latest from your chart repositories...
...Skip local chart repository
...Successfully got an update from the "jetstack" chart repository
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈ Happy Helming!⎈
➜  helm install --set persistence.enabled=false --name my-database stable/postgresql

NAME:   my-database
LAST DEPLOYED: Sat May  4 12:20:48 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Pod(related)
NAME                      READY  STATUS             RESTARTS  AGE
my-database-postgresql-0  0/1    ContainerCreating  0         0s

==> v1/Secret
NAME                    TYPE    DATA  AGE
my-database-postgresql  Opaque  1     0s

==> v1/Service
NAME                             TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)   AGE
my-database-postgresql           ClusterIP  10.107.229.85  <none>       5432/TCP  0s
my-database-postgresql-headless  ClusterIP  None           <none>       5432/TCP  0s

==> v1beta2/StatefulSet
NAME                    READY  AGE
my-database-postgresql  0/1    0s


NOTES:
** Please be patient while the chart is being deployed **

PostgreSQL can be accessed via port 5432 on the following DNS name from within your cluster:

    my-database-postgresql.default.svc.cluster.local - Read/Write connection
To get the password for "postgres" run:

    export POSTGRES_PASSWORD=$(kubectl get secret --namespace default my-database-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)

To connect to your database run the following command:

    kubectl run my-database-postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:10.7.0 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host my-database-postgresql -U postgres



To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace default svc/my-database-postgresql 5432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U postgres
```

</details>

## Learn About Releases

You can list the deployed releases by executing the following command.

```console
➜  helm list
NAME        REVISION    UPDATED                     STATUS      CHART               APP VERSION NAMESPACE
my-database 1           Sat May  4 12:20:48 2019    DEPLOYED    postgresql-3.18.4   10.7.0      default
```
