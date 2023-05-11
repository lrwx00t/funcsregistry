package registry

import (
	"fmt"
)

type Demo1 struct{}

func init() {
	StartRegisteration(Demo1{})
}

func (d Demo1) Demo() {
	fmt.Println("demo1 registered..", d)
}
