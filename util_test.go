package measuregc

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorConsumer(t *testing.T) {
	a, b := errors.New("a"), errors.New("b")
	c := errorConsumer{}
	c.panic()
	c.consume(a)
	assert(c.err == a, "#1")
	c.consume(b)
	assert(c.err == a, "#2")
	c.clear()
	assert(c.err == nil, "#3")
	c.consume(b)
	defer func() {
		if r := recover(); r == nil {
			assert(false, "#4")
		} else {
			fmt.Println("Tested errorConsumer{}")
		}
	}()
	c.panic()
}
