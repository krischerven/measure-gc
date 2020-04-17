package measuregc

import (
	"fmt"
)

func _ASSERT(b bool, s string) {
	if !b {
		panic(fmt.Sprintf("ASSERTION FAILED! (%s)", s))
	}
}
