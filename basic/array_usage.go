package main

import "fmt"

func main()  {

	var a [10]int
	var b = [5]int{1, 2, 3, 4, 5}
	c := [5]int{7, 6, 7, 3, 2}

	a[3] = 4
	fmt.Println("Printing array a:")
	for i:=0; i<10; i++ {
		fmt.Println(a[i])
	}

	fmt.Println("Printing array b:")
	for i:=0; i<5; i++ {
		fmt.Println(b[i])
	}

	fmt.Println("Printing array c:")
	for i:=0; i<5; i++ {
		fmt.Println(c[i])
	}
}
/*
Printing array a:
0
0
0
4
0
0
0
0
0
0
Printing array b:
1
2
3
4
5
Printing array c:
7
6
7
3
2
*/