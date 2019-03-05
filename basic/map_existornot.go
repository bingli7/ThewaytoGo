package main

import "fmt"

func main()  {

	myMap := make(map[int]string)

	myMap[1] = "a"
	myMap[2] = "b"
	myMap[3] = "c"

	i, j := myMap[1]
	// 返回 a true
	fmt.Printf("返回值为：")
	fmt.Println(i, j)

	i,j = myMap[4]
	if j == false {
		fmt.Printf("j = false \n")
		if i == "" {
			fmt.Printf("i is a empty string: \"\"")
		}
	}
}
/*
返回值为：a true
j = false
i is a empty string: ""
*/