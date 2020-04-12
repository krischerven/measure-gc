package measuregc

import (
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

func StartMeasuringGCs() {
	go func() {
		_ = os.Remove("gc.out.txt")
		for {
			var m0 runtime.MemStats
			runtime.ReadMemStats(&m0)
			t1 := time.Now()
			runtime.GC()
			old, _ := ioutil.ReadFile("gc-out.txt")
			if old != nil {
				old = append(old, []byte("\n")...)
			}
			var m1 runtime.MemStats
			runtime.ReadMemStats(&m1)
			err := ioutil.WriteFile("gc-out.txt", append(
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
			time.Sleep(time.Second)
		}
	}()
}
