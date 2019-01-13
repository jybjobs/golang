package exercise

/* 函数类型实现接口 */

//调用器接口
type Invoker interface {
	Call(interface{}) //interface{} 这种类型变量表示任意类型的值
}
