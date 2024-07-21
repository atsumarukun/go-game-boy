package io

import (
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/lcd"
	"github.com/atsumarukun/go-game-boy/internal/pkg/util/converter"
)

const (
	BG_WINDOW_ENABLE          uint8 = 1 << 0
	BG_TILE_MAP               uint8 = 1 << 3
	TILE_DATA_ADDRESSING_MODE uint8 = 1 << 4
	PPU_ENABLE                uint8 = 1 << 7
)

const LYC_EQ_LY uint8 = 1 << 2

type registers struct {
	lcdc uint8
	stat uint8
	scy  uint8
	scx  uint8
	ly   uint8
	lyc  uint8
	bgp  uint8
	obp0 uint8
	obp1 uint8
	wy   uint8
	wx   uint8
}

type PpuMode int

const (
	PPU_HBLANK_MODE PpuMode = iota
	PPU_VBLANK_MODE
	PPU_OAM_SCAN_MODE
	PPU_DRAWING_MODE
)

type context struct {
	mode  PpuMode
	cycle uint8
}

type Ppu interface {
	Read(uint16) uint8
	Write(uint16, uint8)
	Buffer() [lcd.WIDTH * lcd.HEIGHT]uint8
	Mode() PpuMode
	Emulate(Vram) bool
}

type ppu struct {
	regs   registers
	ctx    context
	buffer [lcd.WIDTH * lcd.HEIGHT]uint8
}

func NewPpu() Ppu {
	return &ppu{
		ctx: context{
			mode: PPU_OAM_SCAN_MODE,
		},
	}
}

func (p *ppu) Read(addr uint16) uint8 {
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

func (p *ppu) Write(addr uint16, val uint8) {
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

func (p *ppu) Buffer() [lcd.WIDTH * lcd.HEIGHT]uint8 {
	return p.buffer
}

func (p *ppu) Mode() PpuMode {
	return p.ctx.mode
}

func (p *ppu) Emulate(vram Vram) bool {
	if p.regs.lcdc&PPU_ENABLE == 0 {
		return false
	}

	p.ctx.cycle -= 1
	if p.ctx.cycle > 0 {
		return false
	}

	rendered := false
	switch p.ctx.mode {
	case PPU_HBLANK_MODE:
		p.regs.ly += 1
		if p.regs.ly < 144 {
			p.ctx.mode = PPU_OAM_SCAN_MODE
			p.ctx.cycle = 20
		} else {
			p.ctx.mode = PPU_VBLANK_MODE
			p.ctx.cycle = 114
		}
		p.checkLycEqLy()
	case PPU_VBLANK_MODE:
		p.regs.ly += 1
		if p.regs.ly > 153 {
			rendered = true
			p.regs.ly = 0
			p.ctx.mode = PPU_OAM_SCAN_MODE
			p.ctx.cycle = 20
		} else {
			p.ctx.cycle = 114
		}
		p.checkLycEqLy()
	case PPU_OAM_SCAN_MODE:
		p.ctx.mode = PPU_DRAWING_MODE
		p.ctx.cycle = 43
	case PPU_DRAWING_MODE:
		p.renderBg(vram)
		p.ctx.mode = PPU_HBLANK_MODE
		p.ctx.cycle = 51
	}
	return rendered
}

func (p *ppu) checkLycEqLy() {
	if p.regs.ly == p.regs.lyc {
		p.regs.stat |= LYC_EQ_LY
	} else {
		p.regs.stat &= ^LYC_EQ_LY
	}
}

func (p *ppu) renderBg(vram Vram) {
	if p.regs.lcdc&BG_WINDOW_ENABLE == 0 {
		return
	}

	y := p.regs.ly + p.regs.scy
	for i := 0; i < lcd.WIDTH; i++ {
		x := uint8(i) + p.regs.scx

		tileMapAddr := 0x1800 | (converter.BoolToUint[uint16](p.regs.lcdc&BG_TILE_MAP > 0) << 10)
		val := vram.Read(tileMapAddr | ((uint16(y>>3) << 5) + uint16(x>>3)))
		var index uint
		if p.regs.lcdc&TILE_DATA_ADDRESSING_MODE > 0 {
			index = uint(val)
		} else {
			index = uint(uint16(uint8(val)) + 0x100)
		}

		row := uint16((y & 7) * 2)
		col := uint16(7 - (x & 7))
		addr := uint16(index) << 4

		low := vram.Read(addr | row)
		high := vram.Read(addr | (row + 1))

		pixel := (((high >> col) & 1) << 1) | ((low >> col) & 1)
		switch (p.regs.bgp >> (pixel << 1)) & 0b11 {
		case 0b00:
			p.buffer[uint(lcd.WIDTH)*uint(p.regs.ly)+uint(i)] = 0xFF
		case 0b01:
			p.buffer[uint(lcd.WIDTH)*uint(p.regs.ly)+uint(i)] = 0xAA
		case 0b10:
			p.buffer[uint(lcd.WIDTH)*uint(p.regs.ly)+uint(i)] = 0x55
		default:
			p.buffer[uint(lcd.WIDTH)*uint(p.regs.ly)+uint(i)] = 0x00
		}
	}
}
