package main

import "fmt"

func main()  {
	a := 10

	// 指针*
	var ptr *int

	// 变量存储地址&
	ptr = &a

	fmt.Println(ptr)
}

//0xc00004a080