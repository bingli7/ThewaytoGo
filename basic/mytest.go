package main

import "fmt"

func main() {
	var a string = "This is a string."
	var sliceA []int = a[3:10]
	fmt.Printf("%T\n", sliceA)
	for i, v := range sliceA {
		fmt.Printf("%d\t%c\n",i, v)
	}
}