# tour-of-senzing-go

The Senzing Go SDK supports development on any platform supporting Go.
Now there are caveats:

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

1. [g2-sdk-go-base](https://pkg.go.dev/github.com/senzing/g2-sdk-go-base) - for
   calling Senzing Go SDK APIs natively.
1. [g2-sdk-go-mock](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock) - a
   [test double](https://en.wikipedia.org/wiki/Test_double)
   used for unit testing calls to the Senzing Go SDK.
   It can also be used for programming against the Senzing Go SDK without
   installing the Senzing C libraries.
1. [g2-sdk-go-grpc](https://pkg.go.dev/github.com/senzing/g2-sdk-go-grpc) - for
   calling Senzing SDK APIs over [gRPC](https://grpc.io/).

## Differences

To see differences in the 3 implemenations, mock, grpc, and base, visit
[DiffNow](https://www.diffnow.com/compare-urls),
and compare any of these URLs:

1. **Mock:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.mock>
1. **gRPC:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.grpc>
1. **Base:** <https://raw.githubusercontent.com/docktermj/tour-of-senzing-go/main/main.go.base>

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
