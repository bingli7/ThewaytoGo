package main

import "fmt"

func main() {

	l := [10]int{1, 2, 3, 4, 5}

	for x,y := range l {
		fmt.Printf("l[%d] = %d\n", x, y)
	}
}

/*
l[0] = 1
l[1] = 2
l[2] = 3
l[3] = 4
l[4] = 5
l[5] = 0
l[6] = 0
l[7] = 0
l[8] = 0
l[9] = 0
*/