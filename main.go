package main

import (
	"github.com/lrwx00t/funcsregistry/registry"
)

func main() {
	// registry.StartRegisteration()
	for _, r := range registry.Regs {
		r.Demo()
	}
}
