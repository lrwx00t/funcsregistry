package main

import (
	"fmt"

	"github.com/lrwx00t/funcsregistry/registry"
)

func main() {
	fmt.Println("started ..")
	d1 := registry.Demo1{}
	d1.Demo()
	d2 := registry.Demo2{}
	d2.Demo()
}
