# Go Functions

Provides a Go framework for executing functions with minimal tooling.


## Templates

Three configurations are provided as templates: HTTP, CloudEvent, and Blank.

### 'http' (default)

Invokes an HTTP handler function.  Suitable as a starting point for simple, stateless HTTP event handling tasks.

### 'cloudevents'

Invokes a function which handles CloudEvents requests.  The provided default configuration expects CloudEvents over the HTTP protocol serialized as JSON.

For more information see the [CloudEvents Specification](https://github.com/cloudevents/spec) and the [CloudEvents SDK Docs](https://godoc.org/github.com/cloudevents/sdk-go).  In particular, the SDK supports a wide array of available function signatures.

### 'blank'

The minimal runnable stack, suitable for use as a chron job or for experienced function developers who would prefer to start with the minimum possible preexisting tooling.

## Getting Started

Install Appsody (see the [official documentation](https://appsody.dev/docs/getting-started/installation)).

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

## Installing latest Appsody from Source

Since Appsody is written in Go, it can be built locally.  For example, to install the latest on Linux one could:
```console
$ git clone git@github.com:appsody/appsody $GOPATH/src/github.com/appsody/appsody
$ cd $GOPATH/src/github.com/appsody/appsody
$ GO111MODULE=off make build-linux
$ ln -s $PWD/build/appsody-0.0.0-linux-amd64 /usr/local/bin/appsody
```

## License

This stack is licensed under the [Apache 2.0](./image/LICENSE) license
