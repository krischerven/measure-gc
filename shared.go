package measuregc

import (
	"fmt"
)

func assert(b bool, s string) {
	if !b {
		panic(fmt.Sprintf("ASSERTION FAILED! (%s)", s))
	}
}
