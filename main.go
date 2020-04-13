package measuregc

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

var (
	printer = message.NewPrinter(language.English)
)

func panic_(err error) {
	if err != nil {
		panic(err)
	}
}

// measure GC latency every interval seconds
func StartWith(interval float64) error {
	if interval >= 0.1 {
		_Start(interval)
		return nil
	} else {
		return errors.New("Error: measuregc.Start() interval cannot be < 0.1 second.")
	}
}

// measure GC latency every second
func Start() {
	_Start(1)
}

// internal, unexported function
func _Start(interval float64) {
	go func() {
		_ = os.Remove("gc.out.txt")
		for {
			var m0 runtime.MemStats
			runtime.ReadMemStats(&m0)
			t1 := time.Now()
			runtime.GC()
			old, _ := ioutil.ReadFile("gc.out.txt")
			if old != nil {
				old = append(old, []byte("\n")...)
			}
			var m1 runtime.MemStats
			runtime.ReadMemStats(&m1)
			err := ioutil.WriteFile("gc.out.txt", append(
				old,
				[]byte(
					printer.Sprintf(
						"GC: %.2f ms | Heap⁰: %d | Heap¹: %d",
						float64(time.Now().Sub(t1).Microseconds())/1000,
						m0.HeapAlloc,
						m1.HeapAlloc,
					),
				)...,
			), 0644)
			panic_(err)
			time.Sleep(time.Duration(float64(time.Second) * interval))
		}
	}()
}
