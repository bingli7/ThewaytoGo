package main

import "fmt"

var a = "你好"
var b int = 100

func main()  {
	// :=只能在函数体中出现
	c := "abcde"
	fmt.Println(a, b, c)

	//  多变量声明赋值
	d, e := 100, true
	fmt.Println(d, e)

	// _空白标识符被用于抛弃值
	_, f := d, e
	fmt.Println(f)
}

//你好 100 abcde
//100 true
//true