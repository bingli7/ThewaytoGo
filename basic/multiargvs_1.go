package main

import "fmt"

func main()  {
	s1 := []string{"Li", "Bing"}
	country1 := "China"
	// 传入参数为[]string：方式为：s...
	printName(country1, s1...)
}

func printName(country string, name ...string) {
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