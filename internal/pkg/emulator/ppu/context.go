package ppu

type Mode int

const (
	HBLANK Mode = iota
	VBLANK
	OAM_SCAN
	DRAWING
)

type context struct {
	mode Mode
	cycle uint8
}
