package main

import "fmt"

func main()  {

	a := []int {0, 1, 2, 3, 4}
	b := "My string"
	c := map[int]string {101: "aaa", 102: "bbb", 103: "hahaha"}

	fmt.Printf("Printing array a: \n")
	for k, v := range a{
		fmt.Printf("a[%d] = %d \n", k, v)
	}

	fmt.Printf("\nPrinting string b: \n")
	for k, v := range b {
		fmt.Printf("The %dnd character is: %c \n", k, v)
	}

	fmt.Printf("\nPrinting map c: \n")
	for k, v := range c {
		fmt.Printf("key: %d, value: %s \n", k, v)
	}
}
/*
Printing array a:
a[0] = 0
a[1] = 1
a[2] = 2
a[3] = 3
a[4] = 4

Printing string b:
The 0nd character is: M
The 1nd character is: y
The 2nd character is:
The 3nd character is: s
The 4nd character is: t
The 5nd character is: r
The 6nd character is: i
The 7nd character is: n
The 8nd character is: g

Printing map c:
key: 101, value: aaa
key: 102, value: bbb
key: 103, value: hahaha
*/