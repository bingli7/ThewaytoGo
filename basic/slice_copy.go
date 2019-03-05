package main

import "fmt"

func main()  {

	var a []int
	var b []int

	a = append(a, 1,2,3)

	fmt.Printf("a: %v", a)

	// 返回值为是否成功copy
	copy(b, a)
	fmt.Printf("After copy to b: \n")
	fmt.Printf("b: %v \n", b)
	fmt.Printf("b is empty, because length of b is %d", len(b))

	var c = make([]int, 2, 5)
	// 根据c的长度，只有2个元素被copy
	copy(c, a)
	fmt.Printf("After copy to c: \n")
	fmt.Printf("c: %v \n", c)
	fmt.Printf("c is not empty, only %d elements were copied", len(c))
}