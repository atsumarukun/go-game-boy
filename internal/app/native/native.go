package native

import (
	gogameboy "github.com/atsumarukun/go-game-boy"
	"github.com/atsumarukun/go-game-boy/internal/app/native/lcd"
	"os"
)

func Emulate() {
	if len(os.Args) != 2 {
		panic("bootrom is required")
	}

	lcd := lcd.NewLcd()
	gogameboy.Emulate(os.Args[1], lcd)
}
