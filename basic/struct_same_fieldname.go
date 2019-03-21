package main

import "fmt"

type aStruct struct {
	a int
}

type bStruct struct {
	a int
	b int
}

type cStruct struct {
	aStruct
	bStruct
}

type dStruct struct {
	aStruct
	a int
}

func main() {
	s1 := new(cStruct)
	// 以下代码报错：同一等级有2个a
	// s.a = 1
	s1.aStruct.a = 1
	s1.bStruct.a = 2
	s1.bStruct.b = 3
	fmt.Printf("%v\n", s1)
	// 不同字段同名
	s2 := new(dStruct)
	// 为a int赋值，而不是aStruct中的a字段
	s2.a = 1
	fmt.Printf("%v\n", s2)
	// 为aStruct中的a字段赋值
	s2.aStruct.a = 2
	fmt.Printf("%v\n", s2)
}
/*
&{{1} {2 3}}
&{{0} 1}
&{{2} 1}
*/