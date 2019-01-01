package stud

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "default", "process mode")

func ParMode() {
	flag.Parse()
	fmt.Println(*mode)
}
