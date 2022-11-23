package javascript

import (
	"fmt"

	"github.com/evcc-io/evcc/util"
)

type logger struct {
	log *util.Logger
}

func (l *logger) Print(v ...any) {
	fmt.Println(v...)
	l.log.TRACE.Println(v...)
	// l.log.ERROR.Println(v...)
}

func (l *logger) Error(v ...any) {
	println("++++++++")
	l.log.ERROR.Println(v...)
}
