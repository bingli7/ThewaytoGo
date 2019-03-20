package main

import "fmt"

type myStruct struct {
	name string
}

type myalias myStruct

func main() {
	b := myalias{"Libing"}
	// 将myalias转为myStruct
	// 当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型
	c := myStruct(b)
	fmt.Printf("%s\n", c.name)
}