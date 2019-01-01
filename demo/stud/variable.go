package stud

import (
	"fmt"
)

/*
-------------------------- var --------------------
  1. 定义变量
  2. 初始化变量
   go 在声明变量时，自动对变量对应的内存区域进行初始化操作。
*/
func VarTest() {
	//1.1 标准格式：var 变量名 变量类型
	var aa int32 // 默认 0
	//1.2 批量格式
	var (
		a int
		b string //默认 ""
		c bool
		d []float32
		e struct {
			x int
		}
	)
	fmt.Printf("变量默认值：aa=%d,a=%d,b=%s,c=%t,d=%g,e=%c \n", aa, a, b, c, d, e)

	//2.1 标准初始化：var 变量名 类型 = 表达式
	var num int16 = 100
	//2.2 编译器推到类型
	var str = "100"
	//2.3 短变量声明并初始化
	num2 := 100 // 左值必须时没有定义过的变量(如果定义过会报错：no new variables on left side of :)
	fmt.Printf("变量初始化：var1:%d,var2:%s,var3:%d \n", num, str, num2)
	//2.4 多重赋值
	var m = 100
	var n = 200
	//	m, n = swap(m, n)
	// m,_ = swap(m,n) //如果不需要在左值中接收变量时，可以使用匿名变量
	swapP(&m, &n) //通过指针进行交换
	fmt.Printf("多重赋值：m=%d,n=%d \n", m, n)
}

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

/*
---------------- 指针 ------------------
 * 指针分类：1.类型指针 2.切片
 * 在对普通变量使用"&" 操作符取地址获得这个变量的指针后，
   可以对指针使用 "*"进行指针取值。
 * 创建指针的另一种方法：new() <new()函数可以创建一个对应类型的指针，
   创建过程会分配内存。被创建的指针指向的值为默认值>

*/

func PointTest() {
	var a int = 100
	ptr := &a
	fmt.Printf("%p %T, %d %T \n", &a, &a, *ptr, *ptr) //0xc00004c0d0 *int, 100 int

}

// 使用指针修改值
func swapP(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
