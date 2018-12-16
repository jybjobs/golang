package main // 指明包名,小写

/*
------------- import ------------------
// 1. 不是 . 或 / 开头，如 "fmt" 会在全局文件进行查找
// 2. 可写为:
     import "fmt"
     import "runtime"
// 3. 可见性：
    当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，
    如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端
    程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果
    以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的
   （像面向对象语言中的 private ）。
*/
import (
	"fmt" //格式化io
	"math"
	rt "runtime" //别名
)

func main() {
	fmt.Println("hello,world!")
	fmt.Println(math.Pi)
	var a string
	var b float64
	a, b = Version()
	fmt.Println(a, b)
}

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
 4. 结构化的类型没有真正的值，它使用 nil 作为默认值（在 Objective-C 中是 nil，在 Java 中是 null，在 C 和
  C++ 中是NULL或 0）。值得注意的是，Go 语言中不存在类型继承。
*/

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
