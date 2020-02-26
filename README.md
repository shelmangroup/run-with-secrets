# Run with secrets

[![Docker Repository on Quay](https://quay.io/repository/shelman/run-with-secrets/status "Docker Repository on Quay")](https://quay.io/repository/shelman/run-with-secrets)

Populate environment variables from [Google Secret Manager](https://cloud.google.com/secret-manager) secrets and run another program.

![Run with secrets](run.png)


## Example

```
$ run-with-secrets -s FOO=projects/myproject/secrets/foo/versions/latest -s BAR=projects/myproject/secrets/bar/versions/latest -- /usr/local/bin/awesome-app arg1 arg2 ...
```

