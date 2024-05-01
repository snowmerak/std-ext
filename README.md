# std-ext

This is a collection of extensional libraries for go.

## Installation

```bash
go get github.com/snowmerak/std-ext
```

## Usage

### stream

```go
package main

import (
	"fmt"
	"github.com/snowmerak/std-ext/types/iterable/generator"
	"github.com/snowmerak/std-ext/types/iterable/stream"
)

func main() {
	g := generator.Int(0, 100, 2)
	s := stream.New[int](g)
	v := s.Filter(func(v int) bool {
		return v%3 == 0
	}).Map(func(v int) int {
		return v * 2
	}).Collect()

	fmt.Println(v)
}
```

```shell
[12 24 36 48 60 72 84 96 108 120 132 144 156 168 180 192]
```
