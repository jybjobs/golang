package stud

import "fmt"

/*
 -------------------------------- struct ------------------------------
 1. go 语言通过用自定义的方式形成新的类型，结构体是类型中带有成员的复合类型
 2. 字段： 结构体由一系列成员变量构成，这些成员变量也被称为 字段；
	* 字段拥有自己的类型和值
	* 字段名必须唯一
	* 字段的类型也可以是结构体，甚至可以是字段所在结构体的类型
 3. 定义结构体
    * type 类型名 struct{
		字段1 字段类型
		字段2 字段类型
	}
 4. 实例化结构体
	* 只有当结构体实例化时，才会真正的分配内存。因此必须在定义结构体并实例化后才能使用结构体的字段。
	* 实例化形式：
	  var ins T // ins 为结构体的实例，T为结构体类型
	* 创建指针类型的结构体
	  可以使用new关键字对类型进行实例化，结构体在实例化后会形成指针类型的结构体
	  ins := new(T) //ins 的类型为 *T
	* 取结构体的地址实例化
	  对结构体进行 “&”取地址操作时，是为对该类型进行一次new的实例化操作
	  ins := &T{}
 5. 初始化结构体成员变量
	* 键值对 方式初始化
	  ins := T {
		  字段1: 字段值,
		  字段2: 字段值,
	  }
	* 多个值列表初始化结构体的书写格式: 必须初始化结构体的所有字段，顺序必须一致
	  ins := T {
		  字段值,
		  字段值,
	  }
	* 初始化匿名结构体
	  ins := struct {
		  字段1 字段类型
		  字段2 字段类型
	  }{
		  初始化字段1：字段值,
		  初始化字段2: 字段值，
	  }
 6. 构造函数：
	* go的类型或结构体没有构造函数的功能。结构体的初始化过程可以使用函数封装实现。
	* 多种方式创建和初始化结构体--模拟构造函数重载
	* 带有父子关系的结构体的构造和初始化--模拟父级构造调用
 7. 方法
	* go语言中的方法(method) 是一种作用于特定类型变量的函数。这种特定类型变量叫做 接收器(Receiver)
	* 类似于其它语言中的 this 或 self
	* 每个方法只能有一个接收器（接收器类型可以为指针类型和非指针类型）
	  func (接收器变量 接收器类型) 方法名(参数列表)(返回列表){
		  // 函数体
	  }
-------------------- todo 完善 -------------------------
 8. 类型内嵌和结构体内嵌
 9.json ???
*/

type User struct {
	// Username string
	// Password string
	Username, Password string //同类型的可以写一行
}

type People struct {
	name  string
	child *People //只能包含结构的指针类型
}

func (u *User) add(username string) { // u *User 部分就是接收器
	u.Username = username
}

//8.
type Data struct {
	int
	float32
	bool
}

// test demo
func StructTest() {
	// 4 实例化
	var u1 User
	u1.Username = "xiaoming"
	u2 := new(User)
	u2.Username = "xiaoqiang"
	u3 := &User{}
	u3.Username = "xiaohua"
	fmt.Println(u1, u2, u3) //{xiaoming } &{xiaoqiang } &{xiaohua }
	// 5. 初始化
	relation := &People{
		name: "grandfather",
		child: &People{
			name: "father",
			child: &People{
				name: "me",
			},
		},
	}
	fmt.Println(relation)
	// 8.
	d := &Data{
		int:     10,
		float32: 3.14,
		bool:    false,
	}
	fmt.Println(d)
}
