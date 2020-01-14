# Go CloudEvents Stack

THIS IS A PROOF OF CONCEPT IMPLEMENTATION

This stack provides a Go framework for executing a function that handles [CloudEvents Specification](https://github.com/cloudevents/spec).  The function's method signature is any of those supported by the [CloudEvents Go SDK](https://godoc.org/github.com/cloudevents/sdk-go).


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

By default, a newly initialized Go Functions project is defined as package `function`.  While updating is not required, if this function is expected to be imported or published, the default module should be updated to reflect a proper Go module by either updating `go.mod` manually or recreating it using `go mod init`.  The `package` statements in the default `handler.go` and `handler_test.go` files should be updated as well.

## Example

Initializing a new funciton instance is 


## Templates

Three configurations are provided as templates: HTTP, CloudEvent, and Blank.

### 'simple' (default)

The minimal runnable stack, suitable for use as a chron job or for experienced function developers who would prefer to start with the minimum possible preexisting tooling.


### 'http'

Invokes an HTTP handler function.  Suitable as a starting point for simple, stateless HTTP event handling tasks.

### 'cloudevents'

Invokes a function which handles CloudEvents requests.  The provided default configuration expects CloudEvents over the HTTP protocol serialized as JSON.


## Getting Started

```console
mkdir my-function
cd my-function
appsody init experimental/go cloudevents
```

This will initialize a CloudEvents listener in Go.  To run the container.

```console
appsody run
```
The code will be tested and deployed in a local container.  Changes to local source code will trigger a build and redeploy.

## Building Appsody from Source

Since Appsody is written in Go, it can probably be built locally by users interested in this stack. However, it does still use `dep` rather than go modules, presenting a few additinal required steps.  For example, to build and install the latest on Linux one could:
```console
$ git clone git@github.com:appsody/appsody $GOPATH/src/github.com/appsody/appsody
$ cd $GOPATH/src/github.com/appsody/appsody
$ GO111MODULE=off make build-linux
$ ln -s $PWD/build/appsody-0.0.0-linux-amd64 /usr/local/bin/appsody
```

## License

This stack is licensed under the [Apache 2.0](./image/LICENSE) license
