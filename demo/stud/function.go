package stud

import (
	"fmt"
	"math"
	rt "runtime"
)

/*
---------------- function --------------
 1. 在括号 () 中写入 0 个或多个函数的参数（使用逗号 , 分隔），每个参数的名称后面必须紧跟着该参数的类型
 2. 函数里的代码（函数体）使用大括号 {} 括起来,左大括号 { 必须与方法的声明放在同一行
 3. 只有当某个函数需要被外部包调用的时候才使用大写字母开头，并遵循 Pascal 命名法；否则就遵循骆驼命名法，
    即第一个单词的首字母小写，其余单词的首字母大写。

 因此符合规范的函数一般写成如下的形式：
 func functionName(parameter_list) (return_value_list) {
   …
 }
 其中：
  parameter_list 的形式为 (param1 type1, param2 type2, …)
  return_value_list 的形式为 (ret1 type1, ret2 type2, …)
*/

/*
-------------- type -----------------
 1. 变量（或常量）包含数据，这些数据可以有不同的数据类型，简称类型。
 2. 使用 var 声明的变量的值会自动初始化为该类型的零值。类型定义了某个变量的值的集合与可对其进行操作的集合。
 3. 类型可以是基本类型，如：int、float、bool、string；结构化的（复合的），如：struct、array、slice、map、
 channel；只描述类型的行为的，如：interface。
	* 整型： int8,int16(类似short),int32,int64(类似long);无符号整型 uint8（byte类型）,uint16,uint32,uint64;
	* 浮点型：float32 3.4e38,float64 1.8e308
	* bool型：true，false //无法参与数值运算，无法和其他类型进行转换
	* 字符串：基于utf-8编码
	* 字符：uint8/byte 代表ascll码；rune类型 代表一个utf-8字符，实际是一个int32
	* 切片：一个拥有相同类型元素的可变长度的序列。 var name []T

 4. 结构化的类型没有真正的值，它使用 nil 作为默认值（在 Objective-C 中是 nil，在 Java 中是 null，在 C 和
  C++ 中是NULL或 0）。值得注意的是，Go 语言中不存在类型继承。
*/
func TypeTest() {
	fmt.Println(math.MaxFloat32)
	fmt.Printf("%f\n %.2f", math.Pi, math.Pi)
	//string
	var str = "Happy New Year!\n"
	//多行定义，注意为反引号不是单引号
	const str2 = `happy
	new
	year
	!`
	fmt.Println(str, str2)
	// 字符
	var c1 byte = 'a'
	var c2 rune = '你'
	fmt.Printf("byte:%d %T\n rune:%d %T\n", c1, c1, c2, c2) // %T 打印类型

	//切片
	cc := make([]int, 3) //创建一个容量为3的整型切片
	ss := "happy new year"
	fmt.Println(cc[0], ss[6:]) //0 new year

}

// Version function
func Version() (string, float64) { // 左大括号 { 必须与方法的声明放在同一行
	fmt.Println(rt.Version())
	var v = rt.Version()
	var g = 11.4
	s := float64(g) //类型转换：类型 B 的值 = 类型 B(类型 A 的值)
	return v, s
}

// Sum function
func Sum(a, b int) int { return a + b }
