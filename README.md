# GolangServer
A web server in go. The webserver handles requequests asynchronously to backends, but the frontend is synchronous

# Testing

This instructions assume that $GOROOT, $GOPATH and $GOBIN are correctly set.

Build and install the models and test executables:

```shell
cd src/asynctask/bin
go build
go install test.go
go install testworker.go
```

Execute the test workder

```shell
testworkder 9090
```

Execute the test dispatcher

```shell
test 500
```