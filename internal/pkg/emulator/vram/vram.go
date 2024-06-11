package vram

// type Vram [0x2000]uint8 can be operated externally.
type Vram struct {
	ram [0x2000]uint8
}

func NewVram() *Vram {
	return &Vram{
		[0x2000]uint8{},
	}
}

// addr is the address of the memory map.
// Address range is 0x8000~0x9FFF.

func (v *Vram) Read(addr uint16) uint8 {
	return v.ram[addr&0x1FFF]
}

func (v *Vram) Write(addr uint16, val uint8) {
	v.ram[addr&0x1FFF] = val
}
