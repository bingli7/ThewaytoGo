package main

var aa = "G"

func main() {
	n1()
	m1()
	n1()
}

func n1() {
	print(aa)
}

func m1() {
	aa = "O"
	print(aa)
}
/*
GOO
*/