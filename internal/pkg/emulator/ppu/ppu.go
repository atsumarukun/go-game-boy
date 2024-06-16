package ppu

import (
	"go-game-boy/internal/pkg/emulator/bus"
	"go-game-boy/internal/pkg/emulator/vram"
	"go-game-boy/internal/pkg/util/converter"
)

const (
	WIDTH  = 160
	HEIGHT = 144
)

type Ppu struct {
	regs   registers
	ctx    context
	buffer [WIDTH * HEIGHT]uint8
}

func NewPpu() *Ppu {
	return &Ppu{
		ctx: context{
			mode: OAM_SCAN,
		},
	}
}

func (p *Ppu) Buffer() [WIDTH * HEIGHT]uint8 {
	return p.buffer
}

func (p *Ppu) Read(addr uint16) uint8 {
	switch addr {
	case 0xFF40:
		return p.regs.lcdc
	case 0xFF41:
		return 0x80 | p.regs.stat | uint8(p.ctx.mode)
	case 0xFF42:
		return p.regs.scy
	case 0xFF43:
		return p.regs.scx
	case 0xFF44:
		return p.regs.ly
	case 0xFF45:
		return p.regs.lyc
	case 0xFF47:
		return p.regs.bgp
	case 0xFF48:
		return p.regs.obp0
	case 0xFF49:
		return p.regs.obp1
	case 0xFF4A:
		return p.regs.wy
	case 0xFF4B:
		return p.regs.wx
	default:
		panic("Incorrect address.")
	}
}

func (p *Ppu) Write(addr uint16, val uint8) {
	switch addr {
	case 0xFF40:
		p.regs.lcdc = val
	case 0xFF41:
		p.regs.stat = (p.regs.stat & LYC_EQ_LY) | (val & 0xF8) // The lower 3 bits of F register are always 0.
	case 0xFF42:
		p.regs.scy = val
	case 0xFF43:
		p.regs.scx = val
	case 0xFF44:
		// pass
	case 0xFF45:
		p.regs.lyc = val
	case 0xFF47:
		p.regs.bgp = val
	case 0xFF48:
		p.regs.obp0 = val
	case 0xFF49:
		p.regs.obp1 = val
	case 0xFF4A:
		p.regs.wy = val
	case 0xFF4B:
		p.regs.wx = val
	default:
		panic("Incorrect address.")
	}
}

func (p *Ppu) Emulate(bus *bus.Bus) bool {
	if p.regs.lcdc&PPU_ENABLE == 0 {
		return false
	}

	p.ctx.cycle -= 1
	if p.ctx.cycle > 0 {
		return false
	}

	rendered := false
	switch p.ctx.mode {
	case HBLANK:
		p.regs.ly += 1
		if p.regs.ly < 144 {
			p.ctx.mode = OAM_SCAN
			p.ctx.cycle = 20
		} else {
			p.ctx.mode = VBLANK
			p.ctx.cycle = 114
		}
		p.checkLycEqLy()
	case VBLANK:
		p.regs.ly += 1
		if p.regs.ly > 153 {
			rendered = true
			p.regs.ly = 0
			p.ctx.mode = OAM_SCAN
			p.ctx.cycle = 20
		} else {
			p.ctx.cycle = 114
		}
		p.checkLycEqLy()
	case OAM_SCAN:
		p.ctx.mode = DRAWING
		p.ctx.cycle = 43
	case DRAWING:
		p.renderBg(bus)
		p.ctx.mode = HBLANK
		p.ctx.cycle = 51
	}
	return rendered
}

func (p *Ppu) checkLycEqLy() {
	if p.regs.ly == p.regs.lyc {
		p.regs.stat |= LYC_EQ_LY
	} else {
		p.regs.stat &= ^LYC_EQ_LY
	}
}

func (p *Ppu) renderBg(bus *bus.Bus) {
	if p.regs.lcdc&BG_WINDOW_ENABLE == 0 {
		return
	}

	y := p.regs.ly + p.regs.scy
	for i := 0; i < WIDTH; i++ {
		x := uint8(i) + p.regs.scx

		tileMapAddr := 0x1800 | (converter.BoolToUint[uint16](p.regs.lcdc&BG_TILE_MAP > 0) << 10)
		index := bus.Read(vram.VRAM_ADDR|tileMapAddr|(uint16(y>>3)<<5)|uint16(x>>3)) << 4

		row := uint16((y & 7) * 2)
		col := uint16(7 - (x & 7))
		addr := uint16(index) << 4

		low := bus.Read(vram.VRAM_ADDR | addr | row)
		high := bus.Read(vram.VRAM_ADDR | addr | (row + 1))

		pixel := (((high >> col) & 1) << 1) | ((low >> col) & 1)
		switch (p.regs.bgp >> (pixel << 1)) & 0b11 {
		case 0b00:
			p.buffer[uint(WIDTH)*uint(p.regs.ly)+uint(i)] = 0xFF
		case 0b01:
			p.buffer[uint(WIDTH)*uint(p.regs.ly)+uint(i)] = 0xAA
		case 0b10:
			p.buffer[uint(WIDTH)*uint(p.regs.ly)+uint(i)] = 0x55
		default:
			p.buffer[uint(WIDTH)*uint(p.regs.ly)+uint(i)] = 0x00
		}
	}
}
