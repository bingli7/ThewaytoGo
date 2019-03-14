package main

import "fmt"

func main()  {

	myMap := make(map[int]string)
	myMap[1] = "aaa"
	myMap[2] = "bbb"
	myMap[3] = "ccc"

	// 打印key
	fmt.Printf("只打印key：\n")
	// 只从range接收一个值
	for key := range myMap {
		fmt.Println(key)
	}
	// 打印key和value，注意：不一定按照顺序！
	fmt.Printf("打印key和value：\n")
	for key, value := range myMap {
		fmt.Printf("myMap[%d] = %s\n", key, value)
	}
}
/*
只打印key：
1
2
3
打印key和value：
myMap[1] = aaa
myMap[2] = bbb
myMap[3] = ccc
*/