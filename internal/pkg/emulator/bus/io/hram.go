package io

type Hram interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// type Hram [0x80]uint8 can be operated externally.
type hram struct {
	ram [0x80]uint8
}

func NewHram() Hram {
	return &hram{
		[0x80]uint8{},
	}
}

// addr is the address of the memory map.
// Address range is 0xFF80~0xFFFE.

func (h *hram) Read(addr uint16) uint8 {
	return h.ram[addr&0x7F]
}

func (h *hram) Write(addr uint16, val uint8) {
	h.ram[addr&0x7F] = val
}
