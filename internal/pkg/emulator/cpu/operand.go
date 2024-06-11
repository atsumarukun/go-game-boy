package cpu

type Operand8 int

const (
	REG_A Operand8 = iota
	REG_B
	REG_C
	REG_D
	REG_E
	REG_H
	REG_L
	DST_BC
	DST_DE
	DST_HL
	DST_HL_INC
	DST_HL_DEC
	DST_FF_C
	DST_PC_8
	DST_DST_PC_8
	DST_FF_DST_PC_8
)

type Operand16 int

const (
	REG_AF Operand16 = iota
	REG_BC
	REG_DE
	REG_HL
	REG_SP
	DST_PC_16
	DST_DST_PC_16
)

type Cond int

const (
	N  Cond = iota // Not conditions.
	NZ             // Not Z flag.
	Z              // Z flag.
	NC             // Not C flag.
	C              // C flag.
)
