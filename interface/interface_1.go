package main

import "fmt"

/*
1、方法声明的集合
2、任何类型的对象实现了在接口中声明的全部方法，则表明该类型实现了对应接口。
3、可以作为一种数据类型，实现了该接口的任何对象都可以给对应的接口类型变量赋值。
*/
type myPrint interface {
	print(int) int
}

type S struct {
	value int
}
// 实现S结构体的myPrint接口
func (s S) print(a int) int {
	fmt.Println(s.value)
	fmt.Println(a)
	return 0
}

func main()  {
	var s myPrint = S{100}

	result :=  s.print(111)

	fmt.Println("Result: ", result)
}
/*
100
111
Result:  0
*/