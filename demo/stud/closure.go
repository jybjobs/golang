package stud

import "fmt"

/*
-------------- closure ----------------------
闭包：引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开
了变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量
  * 函数+引用环境=闭包
*/
func ClosTest() {
	str := "hello function"

	foo := func() {
		str = "hello dude"
	}
	foo()
	fmt.Println(str)
	// 闭包实现记忆
	acc := Accumulate(1)
	fmt.Println(acc())
	fmt.Println(acc())
	fmt.Printf("%p \n", acc)

	acc2 := Accumulate(5)
	fmt.Println(acc2())
	fmt.Printf("%p \n", acc2)

}

func Accumulate(value int) func() int {
	return func() int {
		value++ //
		return value
	}
}
