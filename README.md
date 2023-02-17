# tour-of-senzing-go

## Differences

Visit
[DiffNow](https://www.diffnow.com/compare-urls), and compare any of these URLs:

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
