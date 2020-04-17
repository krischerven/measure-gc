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
	_ASSERT(c.err == a, "#1")
	c.consume(b)
	_ASSERT(c.err == a, "#2")
	c.clear()
	_ASSERT(c.err == nil, "#3")
	c.consume(b)
	defer func() {
		if r := recover(); r == nil {
			_ASSERT(false, "#4")
		} else {
			fmt.Println("Tested errorConsumer{}")
		}
	}()
	c.panic()
}
