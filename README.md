# tour-of-senzing-go

The Senzing Go SDK supports development on any platform supporting Go.
There are caveats:

1. On a Linux system, development can be to the native Senzing SDK, a gRPC Server, or to the mock client.
1. On a non-Linux system that supports Docker, development can be to a gRPC Server or to the mock client.
1. On a non-Linux, non-Docker system, development can be to the mock client.

## Programming to an interface

The practice of
[programming to an interface](https://www.google.com/search?channel=fs&client=ubuntu&q=program+to+interfaces)
is supported by the Senzing Go SDK.
In fact, it's recommended.
Senzing has different implementations of the
[G2config](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2config),
[G2configmgr](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2configmgr),
[G2diagnostic](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2diagnostic),
[G2engine](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2engine), and
[G2product](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2product)
interfaces.

They are:

1. **Base:** [g2-sdk-go-base](https://pkg.go.dev/github.com/senzing/g2-sdk-go-base) - for
   calling Senzing Go SDK APIs natively.
1. **gRPC** [g2-sdk-go-grpc](https://pkg.go.dev/github.com/senzing/g2-sdk-go-grpc) - for
   calling Senzing SDK APIs over [gRPC](https://grpc.io/).
1. **Mock:** [g2-sdk-go-mock](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock) - a
   [test double](https://en.wikipedia.org/wiki/Test_double)
   used for unit testing calls to the Senzing Go SDK.
   It can also be used for programming against the Senzing Go SDK without
   installing the Senzing C libraries.

## Differences

To see differences in programming to the 3 implemenations, base, grpc, and mock, visit
[DiffNow](https://www.diffnow.com/compare-urls),
and compare any of these URLs:

1. **Base:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.base>
1. **gRPC:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.grpc>
1. **Mock:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.mock>

Notice that the differences among the files are mostly how the G2* objects are created.
Once the object is created, it is passed to functions that have
parameters which use the Senzing interfaces defined by the imported
`github.com/senzing/g2-sdk-go/g2api` package.
Example:

```go
import "github.com/senzing/g2-sdk-go/g2api"

func loadDatasources(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr, g2Engine g2api.G2engine) {
    :
}
```

### main.go.mock vs. main.go.base

```console
$ diff main.go.mock main.go.base
11,13c11,13
<   "github.com/senzing/g2-sdk-go-mock/g2config"
<   "github.com/senzing/g2-sdk-go-mock/g2configmgr"
<   "github.com/senzing/g2-sdk-go-mock/g2engine"
---
>   "github.com/senzing/g2-sdk-go-base/g2config"
>   "github.com/senzing/g2-sdk-go-base/g2configmgr"
>   "github.com/senzing/g2-sdk-go-base/g2engine"
140c140
<   fmt.Printf("-------------------  Completed running mock implementation --------------\n")
---
>   fmt.Printf("-------------------  Completed running base implementation --------------\n")
```

## Development

### Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

### Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=docktermj
    export GIT_REPOSITORY=tour-of-senzing-go
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

### Run Mock

1. Run Mock implementation.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make run-mock

    ```

### Run gRPC

1. To run a Senzing gRPC server, visit
   [Senzing/servegrpc](https://github.com/Senzing/servegrpc).

1. Run gRPC implementation.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make run-gRPC

    ```

### Run Base

1. Since the Senzing library is a prerequisite, it must be installed first.
    1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

1. Run Mock implementation.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make run-base
