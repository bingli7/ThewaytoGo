package main

import "fmt"

func main()  {

	LOOP:
	for i:=1;i<=3;i++ {
		for j:=1;j<=3;j++ {
			fmt.Printf("i = %d, j = %d \n", i, j)
			if j == 2 {
				// break跳出外部循环
				break LOOP
			}
		}
	}
}
/*
i = 1, j = 1
i = 1, j = 2
*/