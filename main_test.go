package measuregc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// helpers
func _ASSERT(b bool, s string) {
	if !b {
		panic(fmt.Sprintf("ASSERTION FAILED! (%s)", s))
	}
}

func TestPanic_(t *testing.T) {
	// test panic_ (no error)
	panic_(nil)
	fmt.Println("Tested panic_()")
	// test panic (error)
	defer func() {
		if r := recover(); r == nil {
			_ASSERT(false, "#1")
		}
	} ()
	panic_(errors.New("fakeError"))
}

func TestStartWith(t *testing.T) {
	// test StartWith() errors
	c := errorConsumer{}
	c.consume(StartWith(0.00))
	c.consume(StartWith(0.05))
	c.consume(StartWith(0.09))
	_ASSERT(c.err != nil, "#2")
	fmt.Println("Tested StartWith(float64)")
}

func TestStart(t *testing.T) {

	_Start(0.1)
	defer fmt.Println("Tested _Start(float64)")
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(float64(time.Second) * 0.3))
		bytes_, err := ioutil.ReadFile("gc.out.txt")
		_ = os.Remove("gc.out.txt")

		// test sucessfully loaded gc.out.txt
		_ASSERT(
			err == nil,
			fmt.Sprintf("%s (%v)", "#3", err),
		)

		// test 10 lines in gc.out.txt
		c := bytes.Count(bytes_, []byte("\n"))
		_ASSERT(
			c == 2,
			fmt.Sprintf("%s (%d)", "#4", c),
		)

		ch <- struct{}{}
	}()

	// wait to exit
	<-ch
}
