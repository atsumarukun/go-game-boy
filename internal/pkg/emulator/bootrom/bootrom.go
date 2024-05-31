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

func (b *Bootrom) Read(addr uint16) uint8 {
	return b.rom[addr]
}

func (b *Bootrom) IsActive() bool {
	return b.active
}

func (b *Bootrom) ToggleActive(val uint8) {
	b.active = b.active && val == 0
}
