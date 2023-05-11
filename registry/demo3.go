package registry

import (
	"fmt"
)

type Demo3 struct{}

func init() {
	var d Demo3
	_ = Register(d)
}

func (d Demo3) Demo() {
	fmt.Println("demo3 registered..", d)
}
