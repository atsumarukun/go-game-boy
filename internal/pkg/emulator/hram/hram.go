package hram

// type Hram [0x80]uint8 can be operated externally.
type Hram struct {
	ram [0x80]uint8
}

func NewHram() *Hram {
	return &Hram{
		[0x80]uint8{},
	}
}

// addr is the address of the memory map
func (h *Hram) Read(addr uint16) uint8 {
	return h.ram[addr & 0x7F]
}

// addr is the address of the memory map
func (h *Hram) Write(addr uint16, val uint8) {
	h.ram[addr & 0x7F] = val
}
