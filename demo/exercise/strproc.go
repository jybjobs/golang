package exercise

/*
 字符串处理类: 操作和数据分离
*/

//传入切片和处理链
func StringProccess(list []string, chain []func(string) string) {
	for index, str := range list { //遍历切片
		result := str
		for _, proc := range chain { //遍历处理链
			result = proc(result)
		}
		list[index] = result //结果放回切片
	}
}
