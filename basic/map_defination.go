package main

import "fmt"

func main()  {
	// 第一种方式：
	var m1 map[int]string
	// m1元素赋值，以下方式会报错，因为没用用 make 方法来分配内存。
	// panic: assignment to entry in nil map
	// m1[1] = "i am m1"
	m1 = map[int]string{1: "I'm m1"}
	// 第二种定义方式
	m2 := make(map[int]string)
	// m2元素赋值
	m2[1] = "i am m2"
	// 第三种方式
	m3 := map[int]string{1: "aaa", 2: "bbb", 3: "ccc"}
	m3[4] = "ddd"
	// 依次打印
	fmt.Println("Printing m1:")
	for i,v := range m1 {
		fmt.Println(i, " ", v)
	}
	fmt.Println("Printing m2:")
	for i,v := range m2 {
		fmt.Println(i, " ", v)
	}
	fmt.Println("Printing m3:")
	for i,v := range m3 {
		fmt.Println(i, " ", v)
	}
}
/*
Printing m1:
1   I'm m1
Printing m2:
1   i am m2
Printing m3:
2   bbb
3   ccc
4   ddd
1   aaa
*/