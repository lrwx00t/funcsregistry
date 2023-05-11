package registry

import (
	"fmt"
)

type Demo4 struct{}

func init() {
	StartRegisteration(Demo4{})
}

func (d Demo4) Demo() {
	fmt.Println("demo4 registered..", d)
}
