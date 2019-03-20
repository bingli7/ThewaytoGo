package main
import "fmt"

type T struct {
	name string
	age int
}

func main() {
	// *T类型
	var t1 *T
	t1 = new(T)
	t1.name = "libing"
	t1.age = 18
	fmt.Printf("%T, %T\n", t1, *t1)
	fmt.Printf("%s, %d\n", t1.name, (*t1).age)
	// *T类型
	t2 := new(T)
	t2.name = "xiaoliu"
	t2.age = 16
	fmt.Printf("%T, %T\n", t2, *t2)
	// *t2.age会报错，需要加括号：(*t2)
	fmt.Printf("%s, %d\n", t2.name, (*t2).age)
	// T类型：类型 T 的一个实例（instance）或对象（object）。
	t3 := T{name: "vvv", age: 11}
	fmt.Printf("%T\n", t3)
	fmt.Printf("%s, %d\n", t3.name, t3.age)
}