package registry

type funcsSlice []Register

var Regs funcsSlice

type Register interface {
	Demo()
}

func StartRegisteration(r Register) {
	RegisterFunc(r)
}

func RegisterFunc(register Register) {
	Regs = append(Regs, register)
}
