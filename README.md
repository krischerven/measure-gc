# measure-gc
Tiny utility to measure GC latency in Go programs

## Usage
```go
package main

import (
	"github.com/krischerven/measure-gc"
)

func main() {
	measuregc.Start()
	select {}
}
```
