package registry

import (
	"fmt"
)

type Demo2 struct{}

func init() {
	StartRegisteration(Demo2{})
}

func (d Demo2) Demo() {
	fmt.Println("demo2 registered..", d)
}
