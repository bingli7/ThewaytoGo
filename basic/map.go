package main

import "fmt"

func main()  {

	map1 := make(map[string]string)

	map1["a"] = "aaa"
	map1["b"] = "bbb"
	map1["c"] = "ccc"

	fmt.Printf("Printing map1's values: \n")
	for _, v := range map1 {
		fmt.Println(v)
	}
}

/*
Printing map1's values:
aaa
bbb
ccc
*/