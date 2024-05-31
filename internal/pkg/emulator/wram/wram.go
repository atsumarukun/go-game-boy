package wram

// type Wram [0x2000]uint8 can be operated externally.
type Wram struct {
	ram [0x2000]uint8
}

func NewWram() *Wram {
	return &Wram{
		[0x2000]uint8{},
	}
}

func (w *Wram) Read(addr uint16) uint8 {
	return w.ram[addr & 0x1FFF]
}

func (w *Wram) Write(addr uint16, val uint8) {
	w.ram[addr & 0x1FFF] = val
}
