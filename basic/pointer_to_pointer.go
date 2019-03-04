package main

import "fmt"

func main()  {

	v := 10
	//指针
	var ptr *int
	//指向指针的指针
	var pptr **int
	//指针赋值
	ptr = &v
	//指向指针的指针赋值
	pptr = &ptr

	fmt.Println("v's value:", v)
	fmt.Println("ptr:", *ptr)
	fmt.Println("pptr:", **pptr)
}

/*
v's value: 10
ptr: 10
pptr: 10
*/
