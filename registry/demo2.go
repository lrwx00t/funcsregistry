package registry

import (
	"fmt"
)

type Demo2 struct{}

func init() {
	var d Demo2
	fmt.Println(Register(d))
}

func (d Demo2) Demo() {
	fmt.Println("demo2 registered..", d)
}
