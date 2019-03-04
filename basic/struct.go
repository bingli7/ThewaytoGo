package main

import "fmt"

type BookInfo struct {
	title string
	author string
	price int
}

func main() {
	a := BookInfo{"Jane Eyre", "Charlotte Brontë", 33}
	b := BookInfo{title: "Three Body", author: "Liu CiXin", price: 55}
	c := BookInfo{title: "I, Robot", author: "Isaac Asimov"}

	fmt.Println("1st book, title:", a.title, ", author:", a.author, ", price:", a.price)
	fmt.Println("2nd book, title:", b.title, ", author:", b.author, ", price:", b.price)
	fmt.Println("3rd book, title:", c.title, ", author:", c.author, ", price:", c.price)
}

/*
1st book, title: Jane Eyre , author: Charlotte Brontë , price: 33
2nd book, title: Three Body , author: Liu CiXin , price: 55
3rd book, title: I, Robot , author: Isaac Asimov , price: 0
*/