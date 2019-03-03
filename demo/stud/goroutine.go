package stud

import (
	"fmt"
	"time"
)

/*
--------------------- goroutine ----------------------
 轻量级线程(goroutine) 根据需要随时创建的 “线程”.

  1. 创建goroutine： 使用 go 关键字创建
    * 一个函数可以被创建多个goroutine，一个goroutine必定对应一个函数
	  格式：
		go 函数名(参数列表)  //被调用函数的返回值会被忽略
	* 使用匿名函数创建goroutine：
	  格式：
	    go func(参数列表){
			 函数体
		}(调用参数列表) // 启动goroutine时，需要向匿名函数传递的调用参数
  2. 调整并发的运行性能(GOMAXPROCS)
	* 格式： runtime.GOMAXPROCS(逻辑cpu数量) // <1 不修改任何数值



*/

func Running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

/*
------------ channel ----------------
 通道（channel) ：在多个 goroutine之间通信的管道
  * go语言倡导使用通信的方法代替共享内存
  * 特性：
	- 同时只能有一个goroutine访问通道进行发送和获取数据。
	- 先入先出
  * 创建通道：
	通道实例 := make(chan数据类型)
  * 发送数据：
	- 数据通过通道发送格式：    通道变量 <- 值
	- 发送将持续阻塞直到数据被接收
  * 接收数据
	- 通道的收发在不同的两个goroutine间进行
	- 接收将持续阻塞知道发送方发送数据
	- 一次只能接收一个数据元素
	- 写法(4种)：
	  data := <- ch  //阻塞接收数据
	  data,ok := <- ch	//非阻塞接收数据（可能造成cpu过高）
	  <- ch //忽略接收的数据
	  for data:= range ch{ //循环接收

	  }


*/

func channelTest() {
	ch1 := make(chan interface{})
	ch1 <- "hello" //字符串放入通道

}
