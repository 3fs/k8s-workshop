# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Uninstall a Release

To uninstall a release, use the `helm delete` command:

```bash
helm delete my-database
```

This will uninstall `my-database` from Kubernetes, but you will
still be able to request information about that release:

```bash
helm status my-database
```

Because Helm tracks your releases even after you've deleted them, you
can audit a cluster's history, and even undelete a release (with `helm
rollback`).

```bash
helm history my-database
helm rollback my-database 1
```

Output of these commands can be seen here:

```console
➜  helm delete my-database
release "my-database" deleted
➜  helm status my-database
LAST DEPLOYED: Sat May  4 12:38:14 2019
NAMESPACE: default
STATUS: DELETED

NOTES:
** Please be patient while the chart is being deployed **

PostgreSQL can be accessed via port 5432 on the following DNS name from within your cluster:

    my-database-postgresql.default.svc.cluster.local - Read/Write connection
To get the password for "postgres" run:

    export POSTGRES_PASSWORD=$(kubectl get secret --namespace default my-database-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)

To connect to your database run the following command:

    kubectl run my-database-postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:9.6 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host my-database-postgresql -U postgres


To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace default svc/my-database-postgresql 5432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U postgres

➜ helm rollback my-database 1
Rollback was a success! Happy Helming!
➜ helm history my-database
REVISION      UPDATED                         STATUS          CHART                   DESCRIPTION
1               Sat May  4 12:36:58 2019        SUPERSEDED      postgresql-3.18.4       Install complete
2               Sat May  4 12:38:14 2019        SUPERSEDED      postgresql-3.18.4       Deletion complete
3               Sat May  4 12:43:15 2019        DEPLOYED        postgresql-3.18.4       Rollback to 1
```

## Next: [Custom Helm](../04_custom_helm/README.md)
