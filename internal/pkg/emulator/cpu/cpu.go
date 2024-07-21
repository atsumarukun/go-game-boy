package cpu

import (
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus"
	"github.com/atsumarukun/go-game-boy/internal/pkg/util/converter"
)

type Cpu struct {
	regs registers
	ctx  context
}

func NewCpu() *Cpu {
	return &Cpu{}
}

func (c *Cpu) Emulate(bus bus.Bus) {
	if c.ctx.isPrefixCB {
		c.cbDecode(bus)
	} else {
		c.decode(bus)
	}
}

func (c *Cpu) fetch(bus bus.Bus) {
	c.ctx.opcode = bus.Read(c.regs.readPC())
	c.ctx.isPrefixCB = false
}

func (c *Cpu) decode(bus bus.Bus) {
	switch c.ctx.opcode {
	case 0x00:
		c.nop(bus)
	case 0x01:
		c.ld16(bus, REG_BC, DST_PC_16)
	case 0x02:
		c.ld8(bus, DST_BC, REG_A)
	case 0x03:
		c.inc16(bus, REG_BC)
	case 0x04:
		c.inc8(bus, REG_B)
	case 0x05:
		c.dec8(bus, REG_B)
	case 0x06:
		c.ld8(bus, REG_B, DST_PC_8)
	case 0x08:
		c.ld16(bus, DST_DST_PC_16, REG_SP)
	case 0x0A:
		c.ld8(bus, REG_A, DST_BC)
	case 0x0B:
		c.dec16(bus, REG_BC)
	case 0x0C:
		c.inc8(bus, REG_C)
	case 0x0D:
		c.dec8(bus, REG_C)
	case 0x0E:
		c.ld8(bus, REG_C, DST_PC_8)
	case 0x11:
		c.ld16(bus, REG_DE, DST_PC_16)
	case 0x12:
		c.ld8(bus, DST_DE, REG_A)
	case 0x13:
		c.inc16(bus, REG_DE)
	case 0x14:
		c.inc8(bus, REG_D)
	case 0x15:
		c.dec8(bus, REG_D)
	case 0x16:
		c.ld8(bus, REG_D, DST_PC_8)
	case 0x18:
		c.jr(bus, N)
	case 0x1A:
		c.ld8(bus, REG_A, DST_DE)
	case 0x1B:
		c.dec16(bus, REG_DE)
	case 0x1C:
		c.inc8(bus, REG_E)
	case 0x1D:
		c.dec8(bus, REG_E)
	case 0x1E:
		c.ld8(bus, REG_E, DST_PC_8)
	case 0x20:
		c.jr(bus, NZ)
	case 0x21:
		c.ld16(bus, REG_HL, DST_PC_16)
	case 0x22:
		c.ld8(bus, DST_HL_INC, REG_A)
	case 0x23:
		c.inc16(bus, REG_HL)
	case 0x24:
		c.inc8(bus, REG_H)
	case 0x25:
		c.dec8(bus, REG_H)
	case 0x26:
		c.ld8(bus, REG_H, DST_PC_8)
	case 0x28:
		c.jr(bus, Z)
	case 0x2A:
		c.ld8(bus, REG_A, DST_HL_INC)
	case 0x2B:
		c.dec16(bus, REG_HL)
	case 0x2C:
		c.inc8(bus, REG_L)
	case 0x2D:
		c.dec8(bus, REG_L)
	case 0x2E:
		c.ld8(bus, REG_L, DST_PC_8)
	case 0x30:
		c.jr(bus, NC)
	case 0x31:
		c.ld16(bus, REG_SP, DST_PC_16)
	case 0x32:
		c.ld8(bus, DST_HL_DEC, REG_A)
	case 0x33:
		c.inc16(bus, REG_SP)
	case 0x34:
		c.inc8(bus, DST_HL)
	case 0x35:
		c.dec8(bus, DST_HL)
	case 0x36:
		c.ld8(bus, DST_HL, DST_PC_8)
	case 0x38:
		c.jr(bus, C)
	case 0x3A:
		c.ld8(bus, REG_A, DST_HL_DEC)
	case 0x3B:
		c.dec16(bus, REG_SP)
	case 0x3C:
		c.inc8(bus, REG_A)
	case 0x3D:
		c.dec8(bus, REG_A)
	case 0x3E:
		c.ld8(bus, REG_A, DST_PC_8)
	case 0x40:
		c.ld8(bus, REG_B, REG_B)
	case 0x41:
		c.ld8(bus, REG_B, REG_C)
	case 0x42:
		c.ld8(bus, REG_B, REG_D)
	case 0x43:
		c.ld8(bus, REG_B, REG_E)
	case 0x44:
		c.ld8(bus, REG_B, REG_H)
	case 0x45:
		c.ld8(bus, REG_B, REG_L)
	case 0x46:
		c.ld8(bus, REG_B, DST_HL)
	case 0x47:
		c.ld8(bus, REG_B, REG_A)
	case 0x48:
		c.ld8(bus, REG_C, REG_B)
	case 0x49:
		c.ld8(bus, REG_C, REG_C)
	case 0x4A:
		c.ld8(bus, REG_C, REG_D)
	case 0x4B:
		c.ld8(bus, REG_C, REG_E)
	case 0x4C:
		c.ld8(bus, REG_C, REG_H)
	case 0x4D:
		c.ld8(bus, REG_C, REG_L)
	case 0x4E:
		c.ld8(bus, REG_C, DST_HL)
	case 0x4F:
		c.ld8(bus, REG_C, REG_A)
	case 0x50:
		c.ld8(bus, REG_D, REG_B)
	case 0x51:
		c.ld8(bus, REG_D, REG_C)
	case 0x52:
		c.ld8(bus, REG_D, REG_D)
	case 0x53:
		c.ld8(bus, REG_D, REG_E)
	case 0x54:
		c.ld8(bus, REG_D, REG_H)
	case 0x55:
		c.ld8(bus, REG_D, REG_L)
	case 0x56:
		c.ld8(bus, REG_D, DST_HL)
	case 0x57:
		c.ld8(bus, REG_D, REG_A)
	case 0x58:
		c.ld8(bus, REG_E, REG_B)
	case 0x59:
		c.ld8(bus, REG_E, REG_C)
	case 0x5A:
		c.ld8(bus, REG_E, REG_D)
	case 0x5B:
		c.ld8(bus, REG_E, REG_E)
	case 0x5C:
		c.ld8(bus, REG_E, REG_H)
	case 0x5D:
		c.ld8(bus, REG_E, REG_L)
	case 0x5E:
		c.ld8(bus, REG_E, DST_HL)
	case 0x5F:
		c.ld8(bus, REG_E, REG_A)
	case 0x60:
		c.ld8(bus, REG_H, REG_B)
	case 0x61:
		c.ld8(bus, REG_H, REG_C)
	case 0x62:
		c.ld8(bus, REG_H, REG_D)
	case 0x63:
		c.ld8(bus, REG_H, REG_E)
	case 0x64:
		c.ld8(bus, REG_H, REG_H)
	case 0x65:
		c.ld8(bus, REG_H, REG_L)
	case 0x66:
		c.ld8(bus, REG_H, DST_HL)
	case 0x67:
		c.ld8(bus, REG_H, REG_A)
	case 0x68:
		c.ld8(bus, REG_L, REG_B)
	case 0x69:
		c.ld8(bus, REG_L, REG_C)
	case 0x6A:
		c.ld8(bus, REG_L, REG_D)
	case 0x6B:
		c.ld8(bus, REG_L, REG_E)
	case 0x6C:
		c.ld8(bus, REG_L, REG_H)
	case 0x6D:
		c.ld8(bus, REG_L, REG_L)
	case 0x6E:
		c.ld8(bus, REG_L, DST_HL)
	case 0x6F:
		c.ld8(bus, REG_L, REG_A)
	case 0x70:
		c.ld8(bus, DST_HL, REG_B)
	case 0x71:
		c.ld8(bus, DST_HL, REG_C)
	case 0x72:
		c.ld8(bus, DST_HL, REG_D)
	case 0x73:
		c.ld8(bus, DST_HL, REG_E)
	case 0x74:
		c.ld8(bus, DST_HL, REG_H)
	case 0x75:
		c.ld8(bus, DST_HL, REG_L)
	case 0x77:
		c.ld8(bus, DST_HL, REG_A)
	case 0x78:
		c.ld8(bus, REG_A, REG_B)
	case 0x79:
		c.ld8(bus, REG_A, REG_C)
	case 0x7A:
		c.ld8(bus, REG_A, REG_D)
	case 0x7B:
		c.ld8(bus, REG_A, REG_E)
	case 0x7C:
		c.ld8(bus, REG_A, REG_H)
	case 0x7D:
		c.ld8(bus, REG_A, REG_L)
	case 0x7E:
		c.ld8(bus, REG_A, DST_HL)
	case 0x7F:
		c.ld8(bus, REG_A, REG_A)
	case 0xB8:
		c.cp(bus, REG_B)
	case 0xB9:
		c.cp(bus, REG_C)
	case 0xBA:
		c.cp(bus, REG_D)
	case 0xBB:
		c.cp(bus, REG_E)
	case 0xBC:
		c.cp(bus, REG_H)
	case 0xBD:
		c.cp(bus, REG_L)
	case 0xBE:
		c.cp(bus, DST_HL)
	case 0xBF:
		c.cp(bus, REG_A)
	case 0xC1:
		c.pop(bus, REG_BC)
	case 0xC5:
		c.push(bus, REG_BC)
	case 0xC9:
		c.ret(bus)
	case 0xCB:
		c.prefix(bus)
	case 0xCD:
		c.call(bus)
	case 0xD1:
		c.pop(bus, REG_DE)
	case 0xD5:
		c.push(bus, REG_DE)
	case 0xE0:
		c.ld8(bus, DST_FF_DST_PC_8, REG_A)
	case 0xE1:
		c.pop(bus, REG_HL)
	case 0xE2:
		c.ld8(bus, DST_FF_C, REG_A)
	case 0xE5:
		c.push(bus, REG_HL)
	case 0xEA:
		c.ld8(bus, DST_DST_PC_8, REG_A)
	case 0xF0:
		c.ld8(bus, REG_A, DST_FF_DST_PC_8)
	case 0xF1:
		c.pop(bus, REG_AF)
	case 0xF2:
		c.ld8(bus, REG_A, DST_FF_C)
	case 0xF5:
		c.push(bus, REG_AF)
	case 0xFA:
		c.ld8(bus, REG_A, DST_DST_PC_8)
	case 0xFE:
		c.cp(bus, DST_PC_8)
	default:
		panic("Incorrect operation code")
	}
}

func (c *Cpu) cbDecode(bus bus.Bus) {
	switch c.ctx.opcode {
	case 0x10:
		c.rl(bus, REG_B)
	case 0x11:
		c.rl(bus, REG_C)
	case 0x12:
		c.rl(bus, REG_D)
	case 0x13:
		c.rl(bus, REG_E)
	case 0x14:
		c.rl(bus, REG_H)
	case 0x15:
		c.rl(bus, REG_L)
	case 0x16:
		c.rl(bus, DST_HL)
	case 0x17:
		c.rl(bus, REG_A)
	case 0x40:
		c.bit(bus, 0, REG_B)
	case 0x41:
		c.bit(bus, 0, REG_C)
	case 0x42:
		c.bit(bus, 0, REG_D)
	case 0x43:
		c.bit(bus, 0, REG_E)
	case 0x44:
		c.bit(bus, 0, REG_H)
	case 0x45:
		c.bit(bus, 0, REG_L)
	case 0x46:
		c.bit(bus, 0, DST_HL)
	case 0x47:
		c.bit(bus, 0, REG_A)
	case 0x48:
		c.bit(bus, 1, REG_B)
	case 0x49:
		c.bit(bus, 1, REG_C)
	case 0x4A:
		c.bit(bus, 1, REG_D)
	case 0x4B:
		c.bit(bus, 1, REG_E)
	case 0x4C:
		c.bit(bus, 1, REG_H)
	case 0x4D:
		c.bit(bus, 1, REG_L)
	case 0x4E:
		c.bit(bus, 1, DST_HL)
	case 0x4F:
		c.bit(bus, 1, REG_A)
	case 0x50:
		c.bit(bus, 2, REG_B)
	case 0x51:
		c.bit(bus, 2, REG_C)
	case 0x52:
		c.bit(bus, 2, REG_D)
	case 0x53:
		c.bit(bus, 2, REG_E)
	case 0x54:
		c.bit(bus, 2, REG_H)
	case 0x55:
		c.bit(bus, 2, REG_L)
	case 0x56:
		c.bit(bus, 2, DST_HL)
	case 0x57:
		c.bit(bus, 2, REG_A)
	case 0x58:
		c.bit(bus, 3, REG_B)
	case 0x59:
		c.bit(bus, 3, REG_C)
	case 0x5A:
		c.bit(bus, 3, REG_D)
	case 0x5B:
		c.bit(bus, 3, REG_E)
	case 0x5C:
		c.bit(bus, 3, REG_H)
	case 0x5D:
		c.bit(bus, 3, REG_L)
	case 0x5E:
		c.bit(bus, 3, DST_HL)
	case 0x5F:
		c.bit(bus, 3, REG_A)
	case 0x60:
		c.bit(bus, 4, REG_B)
	case 0x61:
		c.bit(bus, 4, REG_C)
	case 0x62:
		c.bit(bus, 4, REG_D)
	case 0x63:
		c.bit(bus, 4, REG_E)
	case 0x64:
		c.bit(bus, 4, REG_H)
	case 0x65:
		c.bit(bus, 4, REG_L)
	case 0x66:
		c.bit(bus, 4, DST_HL)
	case 0x67:
		c.bit(bus, 4, REG_A)
	case 0x68:
		c.bit(bus, 5, REG_B)
	case 0x69:
		c.bit(bus, 5, REG_C)
	case 0x6A:
		c.bit(bus, 5, REG_D)
	case 0x6B:
		c.bit(bus, 5, REG_E)
	case 0x6C:
		c.bit(bus, 5, REG_H)
	case 0x6D:
		c.bit(bus, 5, REG_L)
	case 0x6E:
		c.bit(bus, 5, DST_HL)
	case 0x6F:
		c.bit(bus, 5, REG_A)
	case 0x70:
		c.bit(bus, 6, REG_B)
	case 0x71:
		c.bit(bus, 6, REG_C)
	case 0x72:
		c.bit(bus, 6, REG_D)
	case 0x73:
		c.bit(bus, 6, REG_E)
	case 0x74:
		c.bit(bus, 6, REG_H)
	case 0x75:
		c.bit(bus, 6, REG_L)
	case 0x76:
		c.bit(bus, 6, DST_HL)
	case 0x77:
		c.bit(bus, 6, REG_A)
	case 0x78:
		c.bit(bus, 7, REG_B)
	case 0x79:
		c.bit(bus, 7, REG_C)
	case 0x7A:
		c.bit(bus, 7, REG_D)
	case 0x7B:
		c.bit(bus, 7, REG_E)
	case 0x7C:
		c.bit(bus, 7, REG_H)
	case 0x7D:
		c.bit(bus, 7, REG_L)
	case 0x7E:
		c.bit(bus, 7, DST_HL)
	case 0x7F:
		c.bit(bus, 7, REG_A)
	default:
		panic("Incorrect operation code")
	}
}

func (c *Cpu) read8(bus bus.Bus, operand Operand8) (*uint8, bool) {
	switch operand {
	case REG_A:
		return &c.regs.a, true
	case REG_B:
		return &c.regs.b, true
	case REG_C:
		return &c.regs.c, true
	case REG_D:
		return &c.regs.d, true
	case REG_E:
		return &c.regs.e, true
	case REG_H:
		return &c.regs.h, true
	case REG_L:
		return &c.regs.l, true
	case DST_BC, DST_DE, DST_HL, DST_HL_INC, DST_HL_DEC, DST_FF_C:
		switch c.ctx.operand.step {
		case 0:
			switch operand {
			case DST_BC:
				c.ctx.operand.temp8 = bus.Read(c.regs.readBC())
			case DST_DE:
				c.ctx.operand.temp8 = bus.Read(c.regs.readDE())
			case DST_HL:
				c.ctx.operand.temp8 = bus.Read(c.regs.readHL())
			case DST_HL_INC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr + 1)
				c.ctx.operand.temp8 = bus.Read(addr)
			case DST_HL_DEC:
				addr := c.regs.readHL()
				c.regs.writeHL(addr - 1)
				c.ctx.operand.temp8 = bus.Read(addr)
			case DST_FF_C:
				c.ctx.operand.temp8 = bus.Read(0xFF00 | uint16(c.regs.c))
			}
			c.ctx.operand.step = 1
		case 1:
			c.ctx.operand.step = 0
			return &c.ctx.operand.temp8, true
		}
	case DST_PC_8, DST_DST_PC_8, DST_FF_DST_PC_8:
		switch c.ctx.operand.step {
		case 0:
			c.ctx.operand.temp8 = bus.Read(c.regs.readPC())
			if operand == DST_PC_8 {
				c.ctx.operand.step = 3
			} else if operand == DST_FF_DST_PC_8 {
				c.ctx.operand.temp16 = 0xFF00 | uint16(c.ctx.operand.temp8)
				c.ctx.operand.step = 2
			} else {
				c.ctx.operand.step = 1
			}
		case 1:
			c.ctx.operand.temp16 = uint16(bus.Read(c.regs.readPC()))<<8 | uint16(c.ctx.operand.temp8)
			c.ctx.operand.step = 2
		case 2:
			c.ctx.operand.temp8 = bus.Read(c.ctx.operand.temp16)
			c.ctx.operand.step = 3
		case 3:
			c.ctx.operand.step = 0
			return &c.ctx.operand.temp8, true
		}
	default:
		panic("Incorrect operand")
	}
	return nil, false
}

func (c *Cpu) write8(bus bus.Bus, operand Operand8, val uint8) bool {
	switch operand {
	case REG_A:
		c.regs.a = val
		return true
	case REG_B:
		c.regs.b = val
		return true
	case REG_C:
		c.regs.c = val
		return true
	case REG_D:
		c.regs.d = val
		return true
	case REG_E:
		c.regs.e = val
		return true
	case REG_H:
		c.regs.h = val
		return true
	case REG_L:
		c.regs.l = val
		return true
	case DST_BC, DST_DE, DST_HL, DST_HL_INC, DST_HL_DEC, DST_FF_C:
		switch c.ctx.operand.step {
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
			c.ctx.operand.step = 1
		case 1:
			c.ctx.operand.step = 0
			return true
		}
	case DST_DST_PC_8, DST_FF_DST_PC_8:
		switch c.ctx.operand.step {
		case 0:
			c.ctx.operand.temp8 = bus.Read(c.regs.readPC())
			if operand == DST_FF_DST_PC_8 {
				c.ctx.operand.temp16 = 0xFF00 | uint16(c.ctx.operand.temp8)
				c.ctx.operand.step = 2
			} else {
				c.ctx.operand.step = 1
			}
		case 1:
			c.ctx.operand.temp16 = uint16(bus.Read(c.regs.readPC()))<<8 | uint16(c.ctx.operand.temp8)
			c.ctx.operand.step = 2
		case 2:
			bus.Write(c.ctx.operand.temp16, val)
			c.ctx.operand.step = 3
		case 3:
			c.ctx.operand.step = 0
			return true
		}
	default:
		panic("Incorrect operand")
	}
	return false
}

func (c *Cpu) read16(bus bus.Bus, operand Operand16) (*uint16, bool) {
	switch operand {
	case REG_AF:
		val := c.regs.readAF()
		return &val, true
	case REG_BC:
		val := c.regs.readBC()
		return &val, true
	case REG_DE:
		val := c.regs.readDE()
		return &val, true
	case REG_HL:
		val := c.regs.readHL()
		return &val, true
	case REG_SP:
		return &c.regs.sp, true
	case DST_PC_16:
		switch c.ctx.operand.step {
		case 0:
			c.ctx.operand.temp8 = bus.Read(c.regs.readPC())
			c.ctx.operand.step = 1
		case 1:
			c.ctx.operand.temp16 = uint16(bus.Read(c.regs.readPC()))<<8 | uint16(c.ctx.operand.temp8)
			c.ctx.operand.step = 2
		case 2:
			c.ctx.operand.step = 0
			return &c.ctx.operand.temp16, true
		}
	default:
		panic("Incorrect operand")
	}
	return nil, false
}

func (c *Cpu) write16(bus bus.Bus, operand Operand16, val uint16) bool {
	switch operand {
	case REG_AF:
		c.regs.writeAF(val)
		return true
	case REG_BC:
		c.regs.writeBC(val)
		return true
	case REG_DE:
		c.regs.writeDE(val)
		return true
	case REG_HL:
		c.regs.writeHL(val)
		return true
	case REG_SP:
		c.regs.sp = val
		return true
	case DST_DST_PC_16:
		switch c.ctx.operand.step {
		case 0:
			c.ctx.operand.temp8 = bus.Read(c.regs.readPC())
			c.ctx.operand.step = 1
		case 1:
			c.ctx.operand.temp16 = uint16(bus.Read(c.regs.readPC()))<<8 | uint16(c.ctx.operand.temp8)
			c.ctx.operand.step = 2
		case 2:
			bus.Write(c.ctx.operand.temp16, uint8(val))
			c.ctx.operand.step = 3
		case 3:
			bus.Write(c.ctx.operand.temp16+1, uint8(val>>8))
			c.ctx.operand.step = 4
		case 4:
			c.ctx.operand.step = 0
			return true
		}
	default:
		panic("Incorrect operand")
	}
	return false
}

func (c *Cpu) cond(cond Cond) bool {
	switch cond {
	case NZ:
		return !c.regs.getZF()
	case Z:
		return c.regs.getZF()
	case NC:
		return !c.regs.getCF()
	case C:
		return c.regs.getCF()
	default:
		panic("Incorrect conditions")
	}
}

func (c *Cpu) nop(bus bus.Bus) {
	c.fetch(bus)
}

func (c *Cpu) ld8(bus bus.Bus, dst Operand8, src Operand8) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read8(bus, src); ok {
			c.ctx.exec.temp8 = *val
			c.ctx.exec.step = 1
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write8(bus, dst, c.ctx.exec.temp8); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) ld16(bus bus.Bus, dst Operand16, src Operand16) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read16(bus, src); ok {
			c.ctx.exec.temp16 = *val
			c.ctx.exec.step = 1
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write16(bus, dst, c.ctx.exec.temp16); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) cp(bus bus.Bus, src Operand8) {
	if val, ok := c.read8(bus, src); ok {
		c.regs.setZF(c.regs.a == *val) // c.regs.a - *val == 0
		c.regs.setNF(true)
		c.regs.setHF((c.regs.a & 0xF) < (*val & 0xF))
		c.regs.setCF(c.regs.a < *val)

		c.fetch(bus)
	}
}

func (c *Cpu) inc8(bus bus.Bus, src Operand8) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read8(bus, src); ok {
			c.ctx.exec.temp8 = *val + 1
			c.regs.setZF(c.ctx.exec.temp8 == 0)
			c.regs.setNF(false)
			c.regs.setHF(*val&0xF == 0xF)

			c.ctx.exec.step = 1
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write8(bus, src, c.ctx.exec.temp8); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

// Number of cycles is the number of memory accesses + 1.
func (c *Cpu) inc16(bus bus.Bus, src Operand16) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read16(bus, src); ok {
			c.ctx.exec.temp16 = *val + 1
			c.ctx.exec.step = 1
			return
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write16(bus, src, c.ctx.exec.temp16); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) dec8(bus bus.Bus, src Operand8) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read8(bus, src); ok {
			c.ctx.exec.temp8 = *val - 1
			c.regs.setZF(c.ctx.exec.temp8 == 0)
			c.regs.setNF(true)
			c.regs.setHF(*val&0xF == 0)

			c.ctx.exec.step = 1
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write8(bus, src, c.ctx.exec.temp8); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

// Number of cycles is the number of memory accesses + 1.
func (c *Cpu) dec16(bus bus.Bus, src Operand16) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read16(bus, src); ok {
			c.ctx.exec.temp16 = *val - 1
			c.ctx.exec.step = 1
			return
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write16(bus, src, c.ctx.exec.temp16); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) rl(bus bus.Bus, src Operand8) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read8(bus, src); ok {
			c.ctx.exec.temp8 = (*val << 1) | converter.BoolToUint[uint8](c.regs.getCF())
			c.regs.setZF(c.ctx.exec.temp8 == 0)
			c.regs.setNF(false)
			c.regs.setHF(false)
			c.regs.setCF(*val >= 0x80) // The 7th bit is 1.

			c.ctx.exec.step = 1
		}
	}
	if c.ctx.exec.step == 1 {
		if ok := c.write8(bus, src, c.ctx.exec.temp8); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) bit(bus bus.Bus, bit uint8, src Operand8) {
	if val, ok := c.read8(bus, src); ok {
		c.regs.setZF(*val&(1<<bit) == 0)
		c.regs.setNF(false)
		c.regs.setHF(true)

		c.fetch(bus)
	}
}

func (c *Cpu) push(bus bus.Bus, src Operand16) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read16(bus, src); ok {
			c.ctx.exec.temp16 = *val
			c.ctx.exec.step = 1
			return
		}
	}
	if c.ctx.exec.step == 1 {
		c.regs.sp -= 1
		bus.Write(c.regs.sp, uint8(c.ctx.exec.temp16>>8))
		c.ctx.exec.temp8 = uint8(c.ctx.exec.temp16)

		c.ctx.exec.step = 2
		return
	}
	if c.ctx.exec.step == 2 {
		c.regs.sp -= 1
		bus.Write(c.regs.sp, c.ctx.exec.temp8)

		c.ctx.exec.step = 3
		return
	}
	if c.ctx.exec.step == 3 {
		c.ctx.exec.step = 0
		c.fetch(bus)
	}
}

func (c *Cpu) pop(bus bus.Bus, dst Operand16) {
	if c.ctx.exec.step == 0 {
		c.ctx.exec.temp8 = bus.Read(c.regs.sp)
		c.regs.sp += 1

		c.ctx.exec.step = 1
		return
	}
	if c.ctx.exec.step == 1 {
		c.ctx.exec.temp16 = uint16(bus.Read(c.regs.sp))<<8 | uint16(c.ctx.exec.temp8)
		c.regs.sp += 1

		c.ctx.exec.step = 2
		return
	}
	if c.ctx.exec.step == 2 {
		if ok := c.write16(bus, dst, c.ctx.exec.temp16); ok {
			c.ctx.exec.step = 0
			c.fetch(bus)
		}
	}
}

func (c *Cpu) jr(bus bus.Bus, cond Cond) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read8(bus, DST_PC_8); ok {
			if cond != N && !c.cond(cond) {
				c.fetch(bus)
				return
			}
			c.regs.pc += uint16(int8(*val)) // Add signed integers.

			c.ctx.exec.step = 1
			return
		}
	}
	if c.ctx.exec.step == 1 {
		c.ctx.exec.step = 0
		c.fetch(bus)
	}
}

// Number of cycles is the number of memory accesses + 1.
func (c *Cpu) call(bus bus.Bus) {
	if c.ctx.exec.step == 0 {
		if val, ok := c.read16(bus, DST_PC_16); ok {
			c.ctx.exec.temp16 = *val
			c.ctx.exec.step = 1
			return
		}
	}
	if c.ctx.exec.step == 1 {
		c.regs.sp -= 1
		bus.Write(c.regs.sp, uint8(c.regs.pc>>8))
		c.ctx.exec.temp8 = uint8(c.regs.pc)

		c.ctx.exec.step = 2
		return
	}
	if c.ctx.exec.step == 2 {
		c.regs.sp -= 1
		bus.Write(c.regs.sp, c.ctx.exec.temp8)

		c.ctx.exec.step = 3
		return
	}
	if c.ctx.exec.step == 3 {
		c.regs.pc = c.ctx.exec.temp16
		c.ctx.exec.step = 0
		c.fetch(bus)
	}
}

// Number of cycles is the number of memory accesses + 1.
func (c *Cpu) ret(bus bus.Bus) {
	if c.ctx.exec.step == 0 {
		c.ctx.exec.temp8 = bus.Read(c.regs.sp)
		c.regs.sp += 1

		c.ctx.exec.step = 1
		return
	}
	if c.ctx.exec.step == 1 {
		c.ctx.exec.temp16 = uint16(bus.Read(c.regs.sp))<<8 | uint16(c.ctx.exec.temp8)
		c.regs.sp += 1

		c.ctx.exec.step = 2
		return
	}
	if c.ctx.exec.step == 2 {
		c.regs.pc = c.ctx.exec.temp16
		c.ctx.exec.step = 3
		return
	}
	if c.ctx.exec.step == 3 {
		c.ctx.exec.step = 0
		c.fetch(bus)
	}
}

func (c *Cpu) prefix(bus bus.Bus) {
	val, ok := c.read8(bus, DST_PC_8)
	if !ok {
		return
	}

	c.ctx.opcode = *val
	c.ctx.isPrefixCB = true
	c.cbDecode(bus)
}
