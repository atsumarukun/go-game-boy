package ppu

const (
	BG_WINDOW_ENABLE          uint8 = 1 << 0
	BG_TILE_MAP               uint8 = 1 << 3
	TILE_DATA_ADDRESSING_MODE uint8 = 1 << 4
	PPU_ENABLE                uint8 = 1 << 7
)

const LYC_EQ_LY uint8 = 1 << 2

type registers struct {
	lcdc uint8
	stat uint8
	scy  uint8
	scx  uint8
	ly   uint8
	lyc  uint8
	bgp  uint8
	obp0 uint8
	obp1 uint8
	wy   uint8
	wx   uint8
}
