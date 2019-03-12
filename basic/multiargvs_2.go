package main

import "fmt"

func main()  {
	s1 := "Li"
	s2 := "Bing"
	s3 := "Awesome"
	country1 := "China"
	// 依次传入多个string
	printName2(country1, s1, s2, s3)
}

func printName2(country string, name ...string) {
	fmt.Printf("You are from: %s \t", country)
	// 变长参数使用for range
	for _, value := range name {
		fmt.Printf("%s ", value)
	}
	fmt.Printf("\n")
}
/*
You are from: China 	Li Bing Awesome
*/