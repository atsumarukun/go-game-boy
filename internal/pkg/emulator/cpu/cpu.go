package cpu

import "go-game-boy/internal/pkg/emulator/bus"

type Cpu struct {
	regs registers
	ctx  context
}

func NewCpu() *Cpu {
	return &Cpu{}
}

func (c *Cpu) Emulate(bus *bus.Bus) {
	c.decode(bus)
}

func (c *Cpu) fetch(bus *bus.Bus) {
	c.ctx.opcode = bus.Read(c.regs.pc)
	c.regs.pc += 1
	c.ctx.isPrefixCB = false
}

func (c *Cpu) decode(bus *bus.Bus) {
	switch c.ctx.opcode {
	case 0x00:
		c.nop(bus)
	default:
		panic("Incorrect operation code")
	}
}

func (c *Cpu) read8(bus *bus.Bus, operand Operand8) *uint8 {
	switch operand {
	case REG_A:
		return &c.regs.a
	case REG_B:
		return &c.regs.b
	case REG_C:
		return &c.regs.c
	case REG_D:
		return &c.regs.d
	case REG_E:
		return &c.regs.e
	case REG_H:
		return &c.regs.h
	case REG_L:
		return &c.regs.l
	case DST_BC, DST_DE, DST_HL, DST_HL_INC, DST_HL_DEC, DST_FF_C:
		switch c.ctx.operandStep {
		case 0:
			switch operand {
			case DST_BC:
				c.ctx.temp8 = bus.Read(c.regs.readBC())
			case DST_DE:
				c.ctx.temp8 = bus.Read(c.regs.readDE())
			case DST_HL:
				c.ctx.temp8 = bus.Read(c.regs.readHL())
			case DST_HL_INC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr + 1)
				c.ctx.temp8 = bus.Read(addr)
			case DST_HL_DEC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr - 1)
				c.ctx.temp8 = bus.Read(addr)
			case DST_FF_C:
				c.ctx.temp8 = bus.Read(0xFF00 | uint16(c.regs.c))
			}
			c.ctx.operandStep = 1
		case 1:
			c.ctx.operandStep = 0
			return &c.ctx.temp8
		}
	case DST_PC_8, DST_DST_PC_8, DST_DST_FF_PC_8:
		switch c.ctx.operandStep {
		case 0:
			c.ctx.temp8 = bus.Read(c.regs.pc)
			c.regs.pc += 1
			if operand == DST_PC_8 {
				c.ctx.operandStep = 3
			} else if operand == DST_DST_FF_PC_8 {
				c.ctx.temp16 = 0xFF00 | uint16(c.ctx.temp8)
				c.ctx.operandStep = 2
			} else {
				c.ctx.operandStep = 1
			}
		case 1:
			c.ctx.temp16 = uint16(bus.Read(c.regs.pc))<<8 | uint16(c.ctx.temp8)
			c.regs.pc += 1
			c.ctx.operandStep = 2
		case 2:
			c.ctx.temp8 = bus.Read(c.ctx.temp16)
			c.ctx.operandStep = 3
		case 3:
			c.ctx.operandStep = 0
			return &c.ctx.temp8
		}
	default:
		panic("Incorrect operand")
	}
	return nil
}

func (c *Cpu) write8(bus *bus.Bus, operand Operand8, val uint8) {
	switch operand {
	case REG_A:
		c.regs.a = val
	case REG_B:
		c.regs.b = val
	case REG_C:
		c.regs.c = val
	case REG_D:
		c.regs.d = val
	case REG_E:
		c.regs.e = val
	case REG_H:
		c.regs.h = val
	case REG_L:
		c.regs.l = val
	case DST_BC, DST_DE, DST_HL, DST_HL_INC, DST_HL_DEC, DST_FF_C:
		switch c.ctx.operandStep {
		case 0:
			switch operand {
			case DST_BC:
				bus.Write(c.regs.readBC(), val)
			case DST_DE:
				bus.Write(c.regs.readDE(), val)
			case DST_HL:
				bus.Write(c.regs.readHL(), val)
			case DST_HL_INC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr + 1)
				bus.Write(addr, val)
			case DST_HL_DEC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr - 1)
				bus.Write(addr, val)
			case DST_FF_C:
				bus.Write(0xFF00|uint16(c.regs.c), val)
			}
			c.ctx.operandStep = 1
		case 1:
			c.ctx.operandStep = 0
		}
	case DST_DST_PC_8, DST_DST_FF_PC_8:
		switch c.ctx.operandStep {
		case 0:
			c.ctx.temp8 = bus.Read(c.regs.pc)
			c.regs.pc += 1
			if operand == DST_DST_FF_PC_8 {
				c.ctx.temp16 = 0xFF00 | uint16(c.ctx.temp8)
				c.ctx.operandStep = 2
			} else {
				c.ctx.operandStep = 1
			}
		case 1:
			c.ctx.temp16 = uint16(bus.Read(c.regs.pc))<<8 | uint16(c.ctx.temp8)
			c.regs.pc += 1
			c.ctx.operandStep = 2
		case 2:
			bus.Write(c.ctx.temp16, val)
			c.ctx.operandStep = 3
		case 3:
			c.ctx.operandStep = 0
		}
	default:
		panic("Incorrect operand")
	}
}

func (c *Cpu) read16(bus *bus.Bus, operand Operand16) *uint16 {
	switch operand {
	case REG_AF:
		val := c.regs.readAF()
		return &val
	case REG_BC:
		val := c.regs.readBC()
		return &val
	case REG_DE:
		val := c.regs.readDE()
		return &val
	case REG_HL:
		val := c.regs.readHL()
		return &val
	case REG_SP:
		return &c.regs.sp
	case DST_PC_16:
		switch c.ctx.operandStep {
		case 0:
			c.ctx.temp8 = bus.Read(c.regs.pc)
			c.regs.pc += 1
			c.ctx.operandStep = 1
		case 1:
			c.ctx.temp16 = uint16(bus.Read(c.regs.pc))<<8 | uint16(c.ctx.temp8)
			c.regs.pc += 1
			c.ctx.operandStep = 2
		case 2:
			c.ctx.operandStep = 0
			return &c.ctx.temp16
		}
	default:
		panic("Incorrect operand")
	}
	return nil
}

func (c *Cpu) write16(bus *bus.Bus, operand Operand16, val uint16) {
	switch operand {
	case REG_AF:
		c.regs.writeAF(val)
	case REG_BC:
		c.regs.writeBC(val)
	case REG_DE:
		c.regs.writeDE(val)
	case REG_HL:
		c.regs.writeHL(val)
	case REG_SP:
		c.regs.sp = val
	case DST_DST_PC_16:
		switch c.ctx.operandStep {
		case 0:
			c.ctx.temp8 = bus.Read(c.regs.pc)
			c.regs.pc += 1
			c.ctx.operandStep = 1
		case 1:
			c.ctx.temp16 = uint16(bus.Read(c.regs.pc))<<8 | uint16(c.ctx.temp8)
			c.regs.pc += 1
			c.ctx.operandStep = 2
		case 2:
			bus.Write(c.ctx.temp16, uint8(val))
			c.ctx.operandStep = 3
		case 3:
			bus.Write(c.ctx.temp16+1, uint8(val>>8))
			c.ctx.operandStep = 4
		case 4:
			c.ctx.operandStep = 0
		}
	default:
		panic("Incorrect operand")
	}
}

func (c *Cpu) nop(bus *bus.Bus) {
	c.fetch(bus)
}
