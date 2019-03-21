// struct 匿名字段
package main

import "fmt"

type innerStuct struct {
	inner1 int
	inner2 int
}

type outerStruct struct {
	outer1 int
	int
	string
	// 除基本类型外，匿名字段还可以是struct类型
	innerStuct
}

func main() {
	s := new(outerStruct)
	s.outer1 = 1
	s.int = 2
	s.string = "3"
	// 第一种方式：
	s.innerStuct.inner1 = 4
	s.innerStuct.inner2 = 5
	// 第二种方式
	s.inner1 = 6
	s.inner2 = 7
	fmt.Printf("%v\n", s)
}
/*
&{1 2 3 {6 7}}
*/