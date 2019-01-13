package stud

import (
	t "demo/exercise"
	"errors"
	"fmt"
	rt "runtime"
	"strings"
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
// Version function
func Version() (string, float64) { // 左大括号 { 必须与方法的声明放在同一行
	fmt.Println(rt.Version())
	var v = rt.Version()
	var g = 11.4
	s := float64(g) //类型转换：类型 B 的值 = 类型 B(类型 A 的值)
	return v, s
}

/*
--------------- function --------------------
 1. 声明函数（普通函数需要先声明后调用）：
    func 函数名(参数列表) (返回参数列表) {
		函数体
	}
	* 函数名：字母/数字/下划线组成，且开头不能为数字
	* 返回参数列表：可以是变量名+type，也可以是返回值类型表
 2. 调用函数
	返回值变量列表 = 函数名(参数列表)	//多个返回值逗号隔开
	* 传入和返回参数在调用和返回时都使用值传递，这里需要注意的是指针、
	切片和map等引用型对象指向的内容在参数中不会发生复制，而是将指针进
	行复制，类似创建一次引用。
	* 在go语言中，函数也是一种类型
 3. 匿名函数：没有函数名的函数
    func (参数列表)(返回参数列表){
		函数体
	}
	* 在需要使用函数时，再定义函数，匿名函数没有函数名只有函数体，函数可以
	被作为一种类型赋值给函数类型的变量
	* 被用于实现回调函数和闭包等。
	* 匿名函数可以在声明后调用
	* 匿名函数用途广泛，其本身是一种值，可以方便地保存在各个容器中实现回调函数和操作封装
	* 函数类型实现接口：把函数作为接口来调用

 4.可变参数
    func 函数名(固定参数列表,v ... T) (返回参数列表){
		函数体
	}
	* fmt 包
 5. 延迟执行语句 defer
	* 先被defer的语句最后被执行，最后defer的语句，最先被执行
	* defer 语句正好是在函数退出时执行的语句，所有使用defer能非常方便地处理资源释放问题
 6. 处理运行时发生的错误
	* 一个可能造成错误的函数，需要返回值中返回一个错误接口(error).如果调用是成功的，
	错误接口将返回nil，否则返回错误。
	* 在函数调用后需要检查错误，如果发生错误，进行必要的错误处理
 7. 宕机 panic //类似 throw
	* 宕机有时是一种合理的止损方法
 8. 宕机恢复 recover //类似 try/catch
*/
func FuncTest() {
	//2.1 函数调用
	_, hour, minute := t.ResolveTime(3600)
	fmt.Println(hour, minute)
	// 2.2
	var f func() (int, int)
	f = str
	a, b := f()
	fmt.Println(a, b)
	// 字符串链式处理
	list := []string{
		"go scanner",
		"go parser",
		"go compilter",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace, //移除空格
		strings.ToUpper,
	}
	t.StringProccess(list, chain)

	for _, s := range list {
		/* 输出：
		SCANNER
		PARSER
		COMPILTER
		PRINTER
		FORMATER
		*/
		fmt.Println(s)
	}
	//3.1 匿名函数
	func(data int) {
		fmt.Println("匿名函数输出：", data)
	}(100)
	// 3.2 匿名函数用作回调
	visit([]int{1, 2, 3, 4}, func(a int) {
		fmt.Println(a)
	})
	//3.3 把函数作为接口调用
	var invoker t.Invoker //声明接口变量
	is := new(InvokerS)
	invoker = is
	invoker.Call("hello")
	//3.4 函数体实现接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from FuncCaller", v)
	})
	invoker.Call("hello")
	// defer
	/*
			defer begin
		    defer end
		    第二个 defer
		    第一个 defer
	*/
	fmt.Println("defer begin")
	defer fmt.Println("第一个 defer") //将defer放入延迟调用栈
	defer fmt.Println("第二个 defer")
	fmt.Println("defer end")
	//error 处理
	err := errors.New("this is an error.")
	if err != nil {
		fmt.Println(err)
	}
	//fanic
	defer func() { //defer 会在宕机之前执行，可以通过延迟执行来处理异常，恢复宕机
		e := recover()           //获取panic传递的上下文
		fmt.Println("Error:", e) //Error: this error need exit.
	}()
	panic("this error need exit.")
	//t.ProTest()
}

//1.1 Sum function
func Sum(a, b int) int { return a + b }

//1.2带变量名的返回值
func str() (a, b int) {
	a = 1
	b = 2
	return
}

//移除前缀
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// 3.2 遍历切片元素，通过给定函数进行访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
