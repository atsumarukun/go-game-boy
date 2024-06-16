package bus

type Io interface {
	Read(addr uint16) uint8
	Write(addr uint16, val uint8)
}

type defaultIo struct{}

func (d defaultIo) Read(_ uint16) uint8 {
	return 0xFF
}

func (d defaultIo) Write(_ uint16, _ uint8) {}

type Bus struct {
	bootrom Io
	vram    Io
	wram    Io
	hram    Io
	oam     Io
	ppu     Io
}

func NewBus(bootrom Io, vram Io, wram Io, hram Io, oam Io, ppu Io) *Bus {
	return &Bus{
		bootrom,
		vram,
		wram,
		hram,
		oam,
		ppu,
	}
}

func (b *Bus) find(addr uint16) Io {
	switch {
	case addr <= 0x00FF:
		return b.bootrom
	case 0x8000 <= addr && addr <= 0x9FFF:
		return b.vram
	case 0xC000 <= addr && addr <= 0xDFFF:
		return b.wram
	case 0xFE00 <= addr && addr <= 0xFE9F:
		return b.oam
	case 0xFF40 <= addr && addr <= 0xFF4B:
		return b.ppu
	case addr == 0xFF50:
		return b.bootrom
	case 0xFF80 <= addr && addr <= 0xFFFE:
		return b.hram
	default:
		return defaultIo{}
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
