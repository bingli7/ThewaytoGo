package main

var aaa string

func main() {
	aaa = "G"
	print(aaa)
	f1()
}

func f1() {
	aaa := "O"
	print(aaa)
	f2()
}

func f2() {
	print(aaa)
}

/*
GOG
*/