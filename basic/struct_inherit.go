package main

import "fmt"

type A struct {
	a1, a2 int
}

type B struct {
	A
	b1, b2 int
}

func main() {
	// 第一种方式：分别为每个字段赋值
	s1 := new(B)
	s1.a1 = 1
	s1.a2 = 2
	s1.b1 = 3
	s1.b2 = 4
	fmt.Printf("%v\n", s1)
	// 第二种方式
	s2 := B{A{11,22}, 33,44}
	fmt.Printf("%v\n", s2)
}
/*
&{{1 2} 3 4}
{{11 22} 33 44}
*/