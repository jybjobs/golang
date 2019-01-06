package stud

import (
	"fmt"
	"strings"
)

/*
-------------- 流程控制 --------------

 1. 条件判断 if
  if 条件 {

  } else if 条件{

  }else{

  }
  * 特殊写法： if 表达式之前添加一个执行语句，再根据变量值进行判断
 2. 构建循环 for
  * for 初始化语句;条件表达式;结束语句{}
  * 循环可以通过 break\goto\return\panic语句强制退出
  * 键值循环 for range：用于遍历 数组、切片、字符串、map及channel
 3. switch
  * 不仅可以基于常量进行判断，还可以基于表达式进行判断
  * Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行而是跳出整个switch,
	fallthrough强制执行后面的case代码，但不会判断下一条case的expr结果是否为true。
 4. goto 跳转到指定代码标签
 5. break 可以跳出多层循环: break lebal //表示退出到某个标签对应的块
 6. continue 继续下次循环，添加标签表示开始标签对应的循环 //仅限在for中使用

*/

/*
 if test function
*/
func ForTest() {
	//1. if
	var ten = 11
	if ten == 10 {
		fmt.Printf("%d=10 \n", ten)
	} else if ten > 10 {
		fmt.Printf("%d>10 \n", ten)
	} else {
		fmt.Printf("%d<10 \n", ten)
	}
	// 特殊写法
	if err := conn(); err != "" { // 注意这里判断string是否为空用的是 "",而不是 nil
		fmt.Println("[Error]:", err)
	}

	//2. for
	for i := 0; i < 10; i++ {
		fmt.Println("for print i=", i)
	}
	//2.1 无限循环
	var i = 10
	for {
		if i <= 0 {
			fmt.Println(i, "<=0")
			break
		}
		i--
	}
	//jjTable() //九九乘法表测试
	//channelT()
	//3. switch
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println("hello")
		//fallthrough // 加上后输出 hello  \n  world
	case "world", "China": //可以有多个值
		fmt.Println("world")
	case strings.Replace("heiio", "i", "l", -1): //表达式:将i全部替换为l
		fmt.Println("heiio replace to hello")
	default:
		fmt.Println("no hello")
	}

}

/*
 启用一个goroutine，往通道中推送数据 1 2 3 然后结束关闭；
 for range 对通道c进行遍历，直到通道关闭
*/
func channelT() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

//无实际意义
func conn() string {
	return "connect is error!"
}

// 九九乘法表
func jjTable() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
