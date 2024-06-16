package cpu

type internalContext struct {
	step   uint8
	temp8  uint8
	temp16 uint16
}

type context struct {
	opcode     uint8
	isPrefixCB bool
	operand    internalContext
	exec       internalContext
}
