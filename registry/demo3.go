package registry

import (
	"fmt"
)

type Demo3 struct{}

func init() {
	StartRegisteration(Demo3{})
}

func (d Demo3) Demo() {
	fmt.Println("demo3 registered..", d)
}
