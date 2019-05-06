# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Inspect a Release

Every release has its own release history. For every release with the same
name, helm creates a new release version and deploys it. If the deployment
failes, rollback to the previous version is performed.

To inspect the release history execute the following command:

```bash
helm history my-database
```

We can upgrade the release with new configuration like this. The following will
redeploy PostgreSQL with 9.6 version.

```bash
helm upgrade --set persistence.enabled=false --set image.tag=9.6 my-database stable/postgresql
```

After a successful deployment we can check history once again.

```bash
helm history my-database
```

The whole procedure can be seen here:

[![asciicast](https://asciinema.org/a/CPzz5lL76mDKJt6vcLd3hFd6I.svg)](https://asciinema.org/a/CPzz5lL76mDKJt6vcLd3hFd6I)
