package bootrom

type Bootrom struct {
	rom []uint8
	active bool
}

func NewBootrom(rom []uint8) *Bootrom {
	return &Bootrom{
		rom,
		true,
	}
}

func (b *Bootrom) IsActive() bool {
	return b.active
}

// addr is the address of the memory map
// Address range is 0x0000~0x00FF
// Disabled by writing non-zero to 0xFF50

func (b *Bootrom) Read(addr uint16) uint8 {
	if addr == 0xFF50 {
		panic("Incorrect address.")
	}
	return b.rom[addr]
}

func (b *Bootrom) Write(addr uint16, val uint8) {
	if addr != 0xFF50 {
		panic("Incorrect address.")
	}
	b.active = b.active && val == 0
}
