package io

type Oam interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// type Oam [0xA0]uint8 can be operated externally.
type oam struct {
	data [0xA0]uint8
}

func NewOam() Oam {
	return &oam{
		[0xA0]uint8{},
	}
}

// addr is the address of the memory map.
// Address range is 0xFE00~0xFE9F.

func (o *oam) Read(addr uint16) uint8 {
	return o.data[addr&0xFF]
}

func (o *oam) Write(addr uint16, val uint8) {
	o.data[addr&0xFF] = val
}
