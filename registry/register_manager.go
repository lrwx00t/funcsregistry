package registry

type funcsSlice []Register

var regs funcsSlice

type Register interface {
	Demo()
}

func RegisterFunc(register Register) {
	// regs[0]()
	// regs[1]()
	regs = append(regs, register)
}
