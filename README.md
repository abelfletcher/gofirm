gofirm
===

gofirm provides a simple API for user confirmation from stdin

# Example
```
import (
    "github.com/abelfletcher/gofirm"
)

func main() {
    if gofirm.Proceed("Create directory /home/stuff") {
        // create /home/stuff
    }
}
```


# Installation
go get github.com/abelfletcher/gofirm
