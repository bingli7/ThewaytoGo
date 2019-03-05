package main

import "fmt"

func main()  {

	var mySlice []int

	if mySlice == nil {
		fmt.Println("Slice is nil...")
	} else {
		fmt.Println("Slice is not nil...")
	}
}

/*
Slice is nil...
*/