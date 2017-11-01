# govknum
Turkish Tax Number Validation

## Usage

```
go get github.com/netinternet/govknum
```

## Exemple

```go
package main

import (
  "fmt"
  "github.com/netinternet/govknum"
)

func main() {
  if valid := govknum.vknum("10148343778"); valid {
    fmt.Println("Tax Number Is Valid")
  } else {
    fmt.Println("Tax Number Is Not Valid")
  }
}


```
