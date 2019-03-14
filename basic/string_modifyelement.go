package main

import "fmt"

func main()  {
	myString := "This is a string."
	// 将string转换为slice
	mySlice := []byte(myString)
	// 修改slice中的某个值
	mySlice[0]= 't'
	// 将slice转换为string
	newString := string(mySlice)
	fmt.Printf("%v\n", newString)
	}
/*
this is a string.
*/