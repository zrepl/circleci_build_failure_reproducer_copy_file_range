The `copy_file_range` does not work inside Circle-CI nested Docker containers spawned using `setup_remote_docker`.

This demo contains reproducers written in Go and C.

The build jobs in .circleci/config work fine on my local machine but fail on CircleCI.

See failing builds at: https://app.circleci.com/pipelines/github/zrepl/circleci_build_failure_reproducer_copy_file_range/

## How this was discovered

When running `golang:latest` (or more precisely `golang:1.15` or newer), Go builds and/or invocations of the Linux `copy_file_range` syscall will fail like so:

https://app.circleci.com/pipelines/github/zrepl/zrepl/663/workflows/aa6a96f0-d372-4714-b5b6-cc9b8c091ec5/jobs/2872

```
go build runtime: copying /usr/local/go/src/runtime/tls_arm64.h to /tmp/go-build523120453/b008/tls_GOARCH.h: write /tmp/go-build523120453/b008/tls_GOARCH.h: copy_file_range: operation not permitted
```

The reason for these failures is that Go 1.15 starts to use copy_file_range automatically for io.Copy(*os.File, *os.File): https://github.com/golang/go/issues/36817
