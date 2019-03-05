package main

import "fmt"

// 接口
type Phone interface {
	call(string) int
}

type iPhone struct {
}

type Xiaomi struct {
}
//实现接口
func (i iPhone) call(s string) int {
	fmt.Printf("Calling %s from iPhone...\n", s)
	return 0
}
//实现接口
func (x Xiaomi) call(s string) int {
	fmt.Printf("Calling %s from Xiaomi...\n", s)
	return 0
}
// 函数参数为接口Phone
func callFromPhone(p Phone, s string) int {
	return p.call(s)
}

func main()  {
	// 为接口类型分别赋值为iphone和xiaomi
	var phone1 Phone = iPhone{}
	var phone2 Phone = Xiaomi{}

	callFromPhone(phone1, "libing")
	callFromPhone(phone2, "xiaoliu")
}

/*
Calling libing from iPhone...
Calling xiaoliu from Xiaomi...
*/