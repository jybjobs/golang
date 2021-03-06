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
	log "demo/exercise/logger"
	s "demo/stud"
	"fmt" //格式化io
	"math"
	//别名
)

func main() {
	fmt.Println("hello,world!")
	fmt.Println(math.Pi)
	var a string
	var b float64
	a, b = s.Version() //也可以无需提前定义a,b 直接写成 a, b := s.Version()
	fmt.Println(a, b)
	// s.VarTest()
	// s.TypeTest()
	// s.PointTest()
	// s.ParMode()
	//s.StrTest()
	//	s.ArrayTest()
	//s.MapTest()
	//s.ListTest()
	//s.ForTest()
	//s.FuncTest()
	//s.ClosTest()
	//s.StructTest()

	//------------------- 接口测试
	f := new(s.File)                         //实例化file
	var writer s.DataWriter                  //声明接口
	writer = f                               //将接口赋值，也就是 * File 类型
	writer.WriteData("interface test info.") //write data: interface test info.
	//-------------------logger------
	l := createLogger()
	l.Log("hello")
}

//创建日志器
func createLogger() *log.Logger {
	l := log.NewLogger()
	cw := log.NewConsoleWriter()
	l.RegisterWriter(cw)
	fw := log.NewFileWriter() //创建文件写入器

	if err := fw.SetFile("info.log"); err != nil {
		fmt.Println(err)
	}
	l.RegisterWriter(fw) //注册文件写入器到日志器中
	return l
}
