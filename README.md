# go-shredder
## Installation

```bash
go get github.com/Marwanmhks/go-shredder

```

```go
import (
    "fmt"
    "os"
    "github.com/Marwanmhks/go-shredder"
)

func main() {

	if len(os.Args) != 2 {
        os.Exit(1)
    }
    Input := shredder.Config{Iterations: 3, Remove: true}
    Path := os.Args[1]
    err := Input.File(Path)
    if err != nil {
        fmt.Println("Error: %v", err)
    } else {
        fmt.Println("File shredded")
    }
}
```

```bash
go run yourprogram.go /path/to/your/file

```

## Run tests
Run all tests using `go test -v`.

To generate coverage report:
- Run `go test -v -coverprofile cover.out > cover.txt`
