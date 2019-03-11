package main

import "fmt"

func main()  {

	myValue := 5
	// 初始化赋值 a := 10
	if a := 10; a < myValue {
		fmt.Printf("Smaller~\n")
	}  else {
		fmt.Printf("Bigger~\n")
	}
	// 以下语句会报错：undefined: a
	// fmt.Println(a)

	b := 3
	// 重新赋值 b = 2
	if b = 2; b > myValue {
		fmt.Printf("Bigger~\n")
	} else {
		fmt.Printf("Smaller~\n")
	}
	// 若b在if语句前已定义，if语句中的赋值作用范围只在if语句内部
	fmt.Println(b)
}

/*
Bigger~
Smaller~
2
*/