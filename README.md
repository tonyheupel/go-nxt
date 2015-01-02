# NXT
A Go package for controlling Lego Mindstorms NXT 2.0

## Installation
Either install it from the command-line:

```shell
$ go install github.com/tonyheupel/go-nxt
```

or add as a dependency in your code files:
```go
import "github.com/tonyheupel/go-nxt"
```
```shell
$ go get
```

## Usage
The most basic usage can be done using the NXT class by connecting your NXT 2.0
brick over Bluetooth.  In the examples below, the brick is connected at
``` /dev/tty.NXT-DevB ```:


```go
/// The file bot creates a new NXT bot reference

package main

import (
	"go-nxt"
	"fmt"
)

func main() {
	n := nxt.NewNXT("foobar", "/dev/tty.NXT-DevB")

	fmt.Println(n)
}
```

Results in:

```shell
$ bot
NXT named foobar, at /dev/tty.NXT-DevB
```

