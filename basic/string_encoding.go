package main

import "fmt"

func main()  {

	var s string = "Hello 李兵"
	// len = 12，每个汉字占3个字节
	fmt.Printf("The length of s is: %d \n", len(s))

	for index, rune := range s {

		fmt.Printf("index: %d, value: %c, 十六进制：%X\n", index, rune, rune)
	}
}
/*
The length of s is: 12
index: 0, value: H, 十六进制：48
index: 1, value: e, 十六进制：65
index: 2, value: l, 十六进制：6C
index: 3, value: l, 十六进制：6C
index: 4, value: o, 十六进制：6F
index: 5, value:  , 十六进制：20
index: 6, value: 李, 十六进制：674E
index: 9, value: 兵, 十六进制：5175
*/