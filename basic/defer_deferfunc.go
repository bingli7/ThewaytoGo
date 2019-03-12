package main

import "fmt"

func main() {
	s := "I'm in mydeferFunc."
	defer mydeferFunc(s)
	fmt.Printf("Starting main function...\n")
}

func mydeferFunc(s string) {
	// func参数为string"leaving..."
	defer func(str string) {
		fmt.Printf("%s\n", str)
	}("leaving...")
	fmt.Printf("%s\n", s)
}
/*
Starting main function...
I'm in mydeferFunc.
leaving...
*/