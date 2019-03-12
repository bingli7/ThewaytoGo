package main

import "fmt"

func main()  {
	myCloFunc := myClosure()
	// 在多次调用中，变量 x 的值是被保留的，闭包函数保存并积累其中的变量的值；
	// 不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	fmt.Println(myCloFunc(1))		// x = 1
	fmt.Println(myCloFunc(10))		// x = 11
	fmt.Println(myCloFunc(100))		// x = 111
}

// 返回函数：func(int) int
func myClosure() func(int) int {
	var x int
	return func(a int) int {
		x = x + a
		return x
	}
}
/*
1
11
111
*/