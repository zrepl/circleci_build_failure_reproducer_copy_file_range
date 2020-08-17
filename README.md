When running `golang:latest` (or more precisely `golang:1.15` or newer), Go builds and/or invocations of the Linux `copy_file_range` syscall will fail like so:

https://app.circleci.com/pipelines/github/zrepl/zrepl/663/workflows/aa6a96f0-d372-4714-b5b6-cc9b8c091ec5/jobs/2872
```
go build runtime: copying /usr/local/go/src/runtime/tls_arm64.h to /tmp/go-build523120453/b008/tls_GOARCH.h: write /tmp/go-build523120453/b008/tls_GOARCH.h: copy_file_range: operation not permitted
```

The reason for these failures is that Go 1.15 starts to use copy_file_range automatically for io.Copy(*os.File, *os.File): https://github.com/golang/go/issues/36817

---

I put together a reproducer repository with CircleCI enabled: https://github.com/zrepl/circleci_build_failure_reproducer_copy_file_range

The command in .circleci/config.yml works fine on my local machine but fails in CircleCI.

Failing build in that repository is: https://app.circleci.com/pipelines/github/zrepl/circleci_build_failure_reproducer_copy_file_range/1/workflows/e37d0fa4-0afe-4a23-8d15-662776520c87/jobs/1
