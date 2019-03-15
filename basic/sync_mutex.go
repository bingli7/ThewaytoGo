package main

import (
	"fmt"
	"sync"
)

func main()  {
	var myInfo info
	myInfo.myStr = "This is a string"
	fmt.Println(myInfo.myStr)
	// 上锁
	myInfo.myLock.Lock()
	myInfo.myStr = "This is New string"
	// 更新完成后解锁
	myInfo.myLock.Unlock()

	fmt.Println(myInfo.myStr)
}

type info struct {
	myLock sync.Mutex
	myStr string
}
/*
This is a string
This is New string
*/