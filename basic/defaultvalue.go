/*
各数据类型默认值：
int: 0
string: ""
bool: false
float64: 0
*/

package main

import "fmt"

func main()  {

	var a int
	var b string
	var c bool
	var d float64


	// int的默认值为：0
	fmt.Printf("int a: %d \n", a)
	// string的默认值为空字符串：""
	if b == "" {
		fmt.Printf("string b == \"\"\n")
	}
	// 布尔值bool的默认值为：false
	fmt.Println("bool c：", c)

	// 浮点型float64的默认值为：
	fmt.Println("float64 d: ", d)
}
/*
int a: 0
string b == ""
bool c： false
float64 d:  0
*/