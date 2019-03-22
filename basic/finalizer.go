package main

import (
	"fmt"
	"runtime"
	"time"
)

type Test struct {
	A int
}

func test() {
	// create pointer
	t := &Test{}
	// add finalizer which just prints
	runtime.SetFinalizer(t, func(t *Test) { fmt.Println("I AM DEAD") })
}

func main() {
	test()
	// run garbage collection
	runtime.GC()
	// sleep to switch to finalizer goroutine
	time.Sleep(1 * time.Millisecond)
}
