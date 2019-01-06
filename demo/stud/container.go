package stud

import (
	"container/list"
	"fmt"
	"sync"
)

/*
 ------------------------- container -----------------
 容器：存储和组织数据的方式
  1. 数组： 一段固定长度的连续内存区域(大小不可变)
	 * 声明  var name [num]T  // T可以是任意基本类型
	 * 遍历 range
  2. 切片（slice）： 动态分配大小的连续空间
	 * golang 中切片的内部结构包括地址、大小和容量。
	 * 切片一般用于快速地操作一块数据集合
	 * slice [开始位置:结束位置] //slice 表示目标切片对象
	 * 声明 var name []T
	 * make() 构造切片 make([]T,size,cap) //size 分配大小，cap 预分配的元素数量(提前分配空间)
	 * 使用append()函数为切片添加元素,扩容
	 * 复制切片元素到另一个切片 copy(destSlice,srcSlice[]T) int
	 * 切片中删除元素: go 中没有提供，需要根据特性删除

*/
func ArrayTest() {
	fmt.Println("container list demo.")
	//1.array
	var a [4]string
	a[0] = "hello"
	a[1] = ","
	a[2] = "world"
	a[3] = "!"
	fmt.Printf("a type is %T or %T\n", &a, a) //a type is *[4]string or [4]string
	//var b=[4]string{"hello",",","world","!"}
	var b = [...]string{"hello", ",", "world", "!"} //...编译器自动计算数组长度
	fmt.Println(a, b)
	for k, v := range a {
		fmt.Println(k, v)
	}
	//2. slice
	// 2.1 从数组生成
	aa := a[2:3]                       //从数组生成切片，最后一个元素 slice[len(slice)-1]
	ab := a[2:]                        //中间到结尾
	ac := a[:3]                        //开头到中间
	ad := a[:]                         // 原切片
	ae := a[0:0]                       //重置切片
	fmt.Printf("aa type is %T \n", aa) //aa type is []string
	fmt.Println(aa, ab, ac, ad, ae)    //[world] [world !] [hello , world] [hello , world !] []
	// 2.2 声明
	var strList []string
	var numList = []int{}

	if numList == nil {
		fmt.Println("numList=", numList)
	} else if strList == nil {
		fmt.Println("strList", strList)
	}
	// 2.3 make
	aList := make([]int, 4)
	bList := make([]int, 4, 10)
	fmt.Println(aList, bList) //[0 0 0 0] [0 0 0 0]
	fmt.Println("len:", len(aList), len(bList),
		"cap:", cap(aList), cap(bList)) //len: 4 4 cap: 4 10
	// 2.4 append
	/*
		len: 5 cap: 8 pointer: 0xc00007a080
		len: 6 cap: 8 pointer: 0xc00007a080
		len: 7 cap: 8 pointer: 0xc00007a080
		len: 8 cap: 8 pointer: 0xc00007a080
		len: 9 cap: 16 pointer: 0xc00008c000
		len: 10 cap: 16 pointer: 0xc00008c000
	*/
	for i := 0; i < 6; i++ {
		aList = append(aList, i)
		fmt.Printf("len: %d cap: %d pointer: %p \n", len(aList), cap(aList), aList)
	}
	aList = append(aList, bList...) //"..." 表示将一个slice加到另一个后
	fmt.Println(aList)              //[0 0 0 0 0 1 2 3 4 5 0 0 0 0]
	// 2.5 remove
	aList = rmItem(aList, 5)
	fmt.Println(aList) //[0 0 0 0 0 2 3 4 5 0 0 0 0]
}

func rmItem(srcList []int, index int) (list []int) {
	li := append(srcList[:index], srcList[index+1:]...)
	return li
}

/*
3. 映射(map)：建立事物关联的容器，使用hash实现
	 * 声明 map[keyType]valueType
	 * 遍历 range
	 * 删除 delete(map,key)
	 * 线程安全的map  sync.Map：Store() Load() Delte() Range()
*/
func MapTest() {
	a := make(map[string]string)
	a["name"] = "xiaoming"
	fmt.Println(a["name"])
	m := map[string]string{
		"name":  "xiaoming",
		"age":   "10",
		"email": "xm.qq.com", //！ 最后也要加上逗号
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "age") //删除
	for _, v := range m {
		fmt.Println(v)
	}
	for k := range m { //无需将值改为匿名变量形式
		fmt.Println(k)
	}
	//sync.Map
	var s sync.Map
	s.Store("name", "xiaoming") //保存
	s.Store("age", "18")
	s.Store("email", "xiaoming.qq.com")
	fmt.Println(s.Load("email"))          //查询  xiaoming.qq.com true
	s.Delete("age")                       //删除
	s.Range(func(k, v interface{}) bool { //遍历
		fmt.Println("range:", k, v)
		return true
	})

}

/*
   4. 列表(list) ： 可以快速增删的非连续空间的容器
	* 列表是一种非连续存储的容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，如：单链表、双链表
	* 列表没有元素类型的限制
	* 初始化： New 和 声明，即： 变量名 :=list.New(), var 变量名 list.List
*/
func ListTest() {
	li := list.New()
	//var li2 list.List
	li.PushBack("hello")
	element := li.PushFront("world") //将world 插入到列表前部，并将这个元素的内部结构保存到element
	li.InsertAfter(",", element)
	//li.Remove(element)
	//遍历
	for i := li.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
