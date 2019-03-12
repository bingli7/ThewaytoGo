//使用 defer 语句实现代码追踪

package main

import "fmt"

func main() {
	funcB()
}

func funcA() {
	defer fmt.Printf("Leaving function A...\n")
	fmt.Printf("I'm in function A.\n")
}

func funcB()  {
	defer fmt.Printf("Leaving function B...\n")
	fmt.Printf("I'm in function B.\n")
	funcA()
}
/*
I'm in function B.
I'm in function A.
Leaving function A...
Leaving function B...
*/