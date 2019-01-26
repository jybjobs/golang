package stud

import "fmt"

/*
---------------- interface -------------------------------
  go 使用组合实现对象特性的描述。对象的内部使用结构体内嵌组合对象应该具有的特性，
  对外通过接口暴露能使用的特性。

  1. 声明接口
	 * 格式：
	   type 接口名 interface{ //接口名一般以er结尾 如：Writer/Closer
		   方法名1(参数列表1) 返回值列表1
		   方法名2(参数列表2) 返回值列表2
	   }
  2. 实现接口的条件
	 * 条件一： 接口的方法和实现接口的类型方法格式一致
	   在类型中添加与接口签名一致的方法就可以实现该方法。签名包括方法中的名称、参数列表、返回参数列表。
	 * 条件二： 接口中方法均被实现
	   当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用。
	 * 接口实现是隐式的，无需让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。
  3. 类型和接口是多对多的关系 (实例：logger)
  4. 接口的嵌套组合： 将多个接口放在一个接口内
  5. 空接口类型(interface{})
	 * 空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无需实现空接口。从实现的角度看，
	 任何之都是满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
*/

// 定义一个数据写入接口
type DataWriter interface {
	WriteData(data interface{}) error
	//CanWrite() bool //如果不实现会报错：stud.File does not implement stud.DataWriter (missing CanWrite method)
}

//定义文件结构，用于实现接口方法
type File struct {
}

// 实现 DataWriter接口的WriteData方法
func (f *File) WriteData(data interface{}) error {
	fmt.Println("write data:", data)
	return nil
}
