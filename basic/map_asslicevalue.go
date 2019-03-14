package main
import "fmt"

func main()  {
	// s为slice类型，slice的值类型为map[int]string
	// 初始化内存空间，需指定长度为5
	var s1 = make([]map[int]string, 5)
	m1 := map[int]string{1:"111", 2:"222"}
	s1[0] = m1
	for i,v := range s1[0] {
		fmt.Println(i, " ", v)
	}
	fmt.Printf("--------\n")
	// 初始化并将map m1加入到slice
	var s2 = []map[int]string{m1}
	for i,v := range s2[0] {
		fmt.Println(i, " ", v)
	}
}
/*
1   111
2   222
--------
1   111
2   222
*/