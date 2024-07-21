package lcd

const (
	WIDTH  = 160
	HEIGHT = 144
)

type Lcd interface {
	Render([WIDTH * HEIGHT]uint8)
}
