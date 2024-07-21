package io

type Wram interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// type Wram [0x2000]uint8 can be operated externally.
type wram struct {
	ram [0x2000]uint8
}

func NewWram() Wram {
	return &wram{
		[0x2000]uint8{},
	}
}

// addr is the address of the memory map.
// Address range is 0xC000~0xDFFF.

func (w *wram) Read(addr uint16) uint8 {
	return w.ram[addr&0x1FFF]
}

func (w *wram) Write(addr uint16, val uint8) {
	w.ram[addr&0x1FFF] = val
}
