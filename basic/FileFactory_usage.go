package main

import (
	. "ThewaytoGo/basic/filefactory"
	"fmt"
)

func main() {
	myFile := NewFile(10, "my.txt")
	fmt.Printf("%v\n", myFile)
}
/*
&{10 my.txt}
*/