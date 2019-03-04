package main

import "fmt"

func main()  {
	var array = []int{1, 2, 3, 4, 5}

	// 定义指针数组
	var ptr [5]*int

	for i:=0; i<5; i++ {
		ptr[i] = &array[i]
	}

	for i:=0; i<5; i++ {
		fmt.Println("address:", ptr[i], "value:", array[i])
	}
}
/*
address: 0xc000070030 value: 1
address: 0xc000070038 value: 2
address: 0xc000070040 value: 3
address: 0xc000070048 value: 4
address: 0xc000070050 value: 5
*/