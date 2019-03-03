package main

import 	"fmt"

func main() {

	var a interface{}
	a = "abc"

	switch t := a.(type) {
		case string:
			fmt.Printf("string %s\n", t)
		case int:
			fmt.Printf("int %d\n", t)
		default:
			fmt.Printf("unexpected type %T", t)
	}
}

// string abc