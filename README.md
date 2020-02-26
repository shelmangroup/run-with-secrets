# Run with secrets

![](https://github.com/shelmangroup/run-with-secrets/workflows/Build/badge.svg)


Populate environment variables from [Google Secret Manager](https://cloud.google.com/secret-manager) secrets and run another program.

![Run with secrets](run.png)


## Example

```
$ run-with-secrets -s FOO=projects/myproject/secrets/foo/versions/latest -s BAR=projects/myproject/secrets/bar/versions/latest -- /usr/local/bin/awesome-app arg1 arg2 ...
```

