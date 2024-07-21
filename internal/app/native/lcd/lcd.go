package lcd

import (
	"unsafe"

	gogameboy "github.com/atsumarukun/go-game-boy"

	"github.com/veandco/go-sdl2/sdl"
)

type Lcd struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture
}

func NewLcd() *Lcd {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("GoGameBoy", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, gogameboy.LCD_WIDTH*4, gogameboy.LCD_HEIGHT*4, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_STREAMING, gogameboy.LCD_WIDTH, gogameboy.LCD_HEIGHT)
	if err != nil {
		panic(err)
	}

	return &Lcd{
		renderer: renderer,
		texture:  texture,
	}
}

func (l *Lcd) Render(buffer [gogameboy.LCD_WIDTH * gogameboy.LCD_HEIGHT]uint8) {
	var buf [gogameboy.LCD_WIDTH * gogameboy.LCD_HEIGHT * 3]byte
	for i, v := range buffer {
		for j := 0; j < 3; j++ {
			buf[i*3+j] = v
		}
	}
	l.texture.Update(nil, unsafe.Pointer(&buf), 480)
	l.renderer.Clear()
	l.renderer.Copy(l.texture, nil, nil)
	l.renderer.Present()
}
