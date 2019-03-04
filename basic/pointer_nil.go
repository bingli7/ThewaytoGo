package main

import "fmt"

func main()  {
	var ptr *int
	// 未赋值
	fmt.Println("Before:")
	if ptr == nil {
		fmt.Println("ptr is nil")
	} else {
		fmt.Println("ptr is NOT nil")
		fmt.Println("ptr =", *ptr)
	}
	i := 10
	ptr = &i
	// 赋值后
	fmt.Println("After:")
	if ptr == nil {
		fmt.Println("ptr is nil")
	} else {
		fmt.Println("ptr is NOT nil")
		fmt.Println("ptr =", *ptr)
	}
}
/*
Before:
ptr is nil
After:
ptr is NOT nil
ptr = 10
*/