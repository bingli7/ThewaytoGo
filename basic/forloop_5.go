package main

import "fmt"

func main()  {
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}
/*
Value of s:
Value of s: a
Value of s: aa
Value of s: aaa
Value of s: aaaa
*/