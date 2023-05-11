package registry

import (
	"fmt"
)

type Demo1 struct{}

func init() {
	var d Demo1
	fmt.Println(Register(d))
}

func (d Demo1) Demo() {
	fmt.Println("demo1 registered..", d)
}
