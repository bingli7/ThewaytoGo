package main

import "fmt"

// 闭包
// FuncClosure返回函数类型func，且返回函数用到FuncClosure的形参
func FuncClosure(i int) func() int {
	return func() int{
		i++
		return i
	}
}

func main()  {
	f1 := FuncClosure(4)
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	// f2重新赋予形参
	f2 := FuncClosure(4)
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}

/*
5
6
7
5
6
7
*/