# Go CloudEvents Stack

THIS IS A PROOF OF CONCEPT IMPLEMENTATION

This stack provides a Go framework for executing a function that handles [CloudEvents](https://github.com/cloudevents/spec).  The function's method signature is any of those supported by the [CloudEvents Go SDK](https://godoc.org/github.com/cloudevents/sdk-go).


## Getting Started

Install Appsody (see the [official documentation](https://appsody.dev/docs/getting-started/installation)).  Note that since Appsody itself is written in Go, building from source is also fairly straightforward (see Building Appsody from Source below).

Package this Stack locally using `appsody stack package`, which will add it to the `dev.local` stack repository.

Initialize a new FaaS project
```console
mkdir greeter
cd greeter
appsody init dev.local/go-functions
```

This will create an appsody configuration file (`appsody-config.yaml`) and an example CloudEvent handling package in the local directory, along with some example tests.
```console
go test -cover ./...
ok  	function	0.015s	coverage: 100.0% of statements
```

When development is complete, the final function can be packaged for deployment as Knative serverless container using `appsody build --knative`.  It can also be deloyed directly; see the [Appsody Docs](https://appsody.dev/docs/using-appsody/building-and-deploying) for more.


## Running Locally

To run the function locally, use `appsody run`.  This will start the function listening on `localhost:8080`, and will respond to CloudEvents.
```console
$ curl -X POST -d '{"message": "OK"}' \
  -H'Content-type: application/json' \
  -H'Ce-id: 1' \
  -H'Ce-source: go-function' \
  -H'Ce-type: dev.knative.example' \
  -H'Ce-specversion: 1.0' \
  http://localhost:8080
```

## Updating the Package Name

By default, a newly initialized Go Functions project is defined as package `function`.  While updating is not required, if this function is expected to be imported as a Go module, the default module should be updated to reflect a proper Go module by either updating `go.mod` manually or recreating it using `go mod init`.  The `package` statements in the default `handler.go` and `handler_test.go` files should be updated as well.

## Building Appsody from Source

Since Appsody is written in Go, it can probably be built locally by users interested in this stack. However, it does still use `dep` rather than go modules, presenting a few additinal required steps.  For example, to build and install the latest on Linux one could:
```console
$ git clone git@github.com:appsody/appsody $GOPATH/src/github.com/appsody/appsody
$ cd $GOPATH/src/github.com/appsody/appsody
$ GO111MODULE=off make build-linux
$ ln -s $PWD/build/appsody-0.0.0-linux-amd64 /usr/local/bin/appsody
```

# Development

## Releasing Stack Updates

Version the stack in stack.yaml

Package the stack updates
```
appsody stack package --image-namespace boson --image-registry quay.io
```
Tag the new version git to match
```
git tag v[major].[minor].[patch]
```
Push the resultant images to the registry.
```
docker push quay.io/boson/go-ce-functions:latest
docker push quay.io/boson/go-ce-functions:[major]
docker push quay.io/boson/go-ce-functions:[major].[minor]
docker push quay.io/boson/go-ce-functions:[major].[minor].[patch]
```
Create a release of the given tag which includes the template archive.
```
hub release create -a ~/.appsody/stacks/dev.local/go-ce-functions.[tag].templates.default.tar.gz -m "[tag]" [tag]
```

The stack is now versioned and released, ready to be included in a released Stack index.  See `github.com/boson-project/stacks`, which is the stack index, and includes instructions on releasing.


## License

This stack is licensed under the [Apache 2.0](./image/LICENSE) license
