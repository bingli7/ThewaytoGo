package main

import "fmt"

func main()  {
	s := []string{"Li", "Bing"}
	country := "China"
	printName3(country, s)
	// 不能如下依次传入多个string
	// printName3(country1, s1, s2, s3)
}

func printName3(country string, name []string) {
	fmt.Printf("You are from: %s \t", country)
	// 变长参数使用for range
	for _, value := range name {
		fmt.Printf("%s ", value)
	}
	fmt.Printf("\n")
}
/*
You are from: China 	Li Bing
*/
