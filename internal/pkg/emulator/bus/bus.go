package bus

import (
	"go-game-boy/internal/pkg/emulator/bootrom"
	"go-game-boy/internal/pkg/emulator/hram"
	"go-game-boy/internal/pkg/emulator/wram"
)

type io interface {
	Read(addr uint16) uint8
	Write(addr uint16, val uint8)
}

type Bus struct {
	bootrom *bootrom.Bootrom
	wram *wram.Wram
	hram *hram.Hram
}

func NewBus(bootrom *bootrom.Bootrom, wram *wram.Wram, hram *hram.Hram) *Bus {
	return &Bus{
		bootrom,
		wram,
		hram,
	}
}

func (b *Bus) find(addr uint16) io {
	switch {
	case addr <= 0x00FF:
		return b.bootrom
	case 0xC000 <= addr && addr <= 0xDFFF:
		return b.wram
	case addr == 0xFF50:
		return b.bootrom
	case 0xFF80 <= addr && addr <= 0xFFFE:
		return b.hram
	default:
		panic("Incorrect address.")
	}
}

func (b *Bus) Read(addr uint16) uint8 {
	io := b.find(addr)
	return io.Read(addr)
}

func (b *Bus) Write(addr uint16, val uint8) {
	io := b.find(addr)
	io.Write(addr, val)
}
