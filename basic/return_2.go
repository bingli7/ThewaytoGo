// return with no value

package main

func returnTest(a, b int) {
	return
}

func main()  {
	a, b := 1, 2
	returnTest(a, b)
	return
}