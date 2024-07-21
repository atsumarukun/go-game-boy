package io

type Vram interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// type vram [0x2000]uint8 can be operated externally.
type vram struct {
	ram [0x2000]uint8
}

func NewVram() Vram {
	return &vram{
		[0x2000]uint8{},
	}
}

// addr is the address of the memory map.
// Address range is 0x8000~0x9FFF.

func (v *vram) Read(addr uint16) uint8 {
	return v.ram[addr&0x1FFF]
}

func (v *vram) Write(addr uint16, val uint8) {
	v.ram[addr&0x1FFF] = val
}
