package stud

import (
	"fmt"
)

type InvokerS struct {
}

// 实现Invoker call
func (s *InvokerS) Call(p interface{}) {
	fmt.Println("from InvokerS", p)
}
