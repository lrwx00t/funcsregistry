package registry

type funcsSlice []Register

var Regs funcsSlice

type Register interface {
	Demo()
}

func StartRegisteration(r Register) {
	RegisterFunc(r)
}

// RegisterFunc can be configured to return the target func
func RegisterFunc(register Register) {
	// regs[0]()
	// regs[1]()
	Regs = append(Regs, register)
}
