package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i:=0; i<10; i++ {
		time.Sleep(time.Millisecond  * 100)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}
/*
输出没有先后顺序：
world
hello
hello
world
hello
world
world
hello
hello
world
world
hello
world
hello
hello
world
world
hello
hello
world
*/