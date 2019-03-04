package main

import "fmt"

// 本示例定义一个结构体类型和该类型的一个方法

const PI  = 3.14

/* 定义结构体 */
type Circle struct {
	radius float64
}

// 定义Circle类型的一个方法
func (c Circle) getArea() float64 {
	return PI * c.radius * c.radius
}

func main()  {
	var c Circle
	c.radius = 10

	area := c.getArea()
	fmt.Println("area = ", area)
}