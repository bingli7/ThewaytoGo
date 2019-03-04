package main

import "fmt"

type Books struct {
	title string
	author string
	price int
}

func main()  {

	var book1 Books
	var book2 Books

	book1.title = "二手时间"
	book1.author = "S.A.阿列克谢耶维奇"
	book1.price = 67

	book2.title = "Brief History of Humankind"
	book2.author = "Yuval Noah Harari"
	book2.price = 43

	book_pointer_1 := &book1
/*	book_pointer_2 := &book2*/

	fmt.Println("1st book:")

	// 结构体指针，与结构体变量一样访问成员
	fmt.Println("Title: ", book_pointer_1.title)
	fmt.Println("Title:", book1.title)
}
/*
1st book:
Title:  二手时间
Title: 二手时间
*/