package cpu

type context struct {
	opcode      uint8
	isPrefixCB  bool
	operandStep uint8
	execStep    uint8
	temp8       uint8
	temp16      uint16
}
