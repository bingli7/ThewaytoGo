// 切记！ map为引用类型

package main

import "fmt"

func main()  {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	var m2 map[int]string
	// 更新m2并指向m1的值
	m2 = m1
	// 修改m1元素值
	m1[3] = "New value"
	fmt.Println(m2[3])
	// 修改m2元素值
	m2[4] = "The 4th element"
	fmt.Println("---------")
	// 循环输出m1的值，从m2定义的新值也会输出
	// 可能不会按照定义顺序输出
	for i, v := range m1 {
		fmt.Println(i, " ", v)
	}
}
/*
New value
---------
1   a
2   b
3   New value
4   The 4th element
*/