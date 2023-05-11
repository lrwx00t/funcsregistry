package main

import (
	"github.com/lrwx00t/funcsregistry/registry"
)

func main() {
	for _, r := range registry.Regs {
		r.Demo()
	}
}
