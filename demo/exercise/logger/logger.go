package exercise

/*
------------ Logger --------------
*/

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWriter
}

//注册日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

//数据写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

// 创建日志实例
func NewLogger() *Logger {
	return &Logger{}
}
