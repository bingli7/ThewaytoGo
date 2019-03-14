package main

import "fmt"

func main() {
	// []byte类型
	mySlice := []byte{'a', 'b', 'c'}
	myString := string(mySlice)
	fmt.Printf("%v", myString)
}
/*
abc
*/