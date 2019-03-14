package main

import "fmt"

func main()  {
	s := "This is a string."
	mySlice := []byte(s)
	for _, v := range mySlice {
		fmt.Printf("%c", v)
	}
}
/*
This is a string.
*/