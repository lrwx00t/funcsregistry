package registry

type funcsSlice []Register

var Regs funcsSlice

type Register interface {
	Demo()
}

func StartRegisteration() {
	d1 := Demo1{}
	d2 := Demo2{}
	d3 := Demo3{}
	RegisterFunc(d1)
	RegisterFunc(d2)
	RegisterFunc(d3)
}

// RegisterFunc can be configured to return the target func
func RegisterFunc(register Register) {
	// regs[0]()
	// regs[1]()
	Regs = append(Regs, register)
}
