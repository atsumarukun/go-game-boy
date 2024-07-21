package bus

import "github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus/io"

type Bus interface {
	Read(uint16) uint8
	Write(uint16, uint8)
	Ppu() io.Ppu
	Vram() io.Vram
}

type bus struct {
	bootrom io.Bootrom
	vram    io.Vram
	wram    io.Wram
	hram    io.Hram
	oam     io.Oam
	ppu     io.Ppu
}

func NewBus(bootrom io.Bootrom, vram io.Vram, wram io.Wram, hram io.Hram, oam io.Oam, ppu io.Ppu) Bus {
	return &bus{
		bootrom,
		vram,
		wram,
		hram,
		oam,
		ppu,
	}
}

func (b *bus) find(addr uint16) io.Io {
	switch {
	case addr <= 0x00FF:
		return b.bootrom
	case 0x8000 <= addr && addr <= 0x9FFF:
		return b.vram
	case 0xC000 <= addr && addr <= 0xDFFF:
		return b.wram
	case 0xFE00 <= addr && addr <= 0xFE9F:
		return b.oam
	case 0xFF40 <= addr && addr <= 0xFF4B:
		return b.ppu
	case addr == 0xFF50:
		return b.bootrom
	case 0xFF80 <= addr && addr <= 0xFFFE:
		return b.hram
	default:
		return io.DefaultIo{}
	}
}

func (b *bus) Read(addr uint16) uint8 {
	io := b.find(addr)
	return io.Read(addr)
}

func (b *bus) Write(addr uint16, val uint8) {
	io := b.find(addr)
	io.Write(addr, val)
}

func (b *bus) Ppu() io.Ppu {
	return b.ppu
}

func (b *bus) Vram() io.Vram {
	return b.vram
}
