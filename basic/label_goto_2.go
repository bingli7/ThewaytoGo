package main

import "fmt"

func main()  {

	i := 1

	MYLABEL:
		fmt.Println(i)
		i ++
		if i <= 5 {
			goto MYLABEL
		}
}
/*
goto实现循环：如果i<=5，则退出循环：
1
2
3
4
5
*/