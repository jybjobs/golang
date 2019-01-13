package stud

//函数定义为类型
type FuncCaller func(interface{})

//实现invoker 的call
func (f FuncCaller) Call(p interface{}) {
	f(p) //调用f()
}
