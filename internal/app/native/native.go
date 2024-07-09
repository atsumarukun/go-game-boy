package native

import (
	gogameboy "go-game-boy"
	"go-game-boy/internal/app/native/lcd"
	"os"
)

func Emulate() {
	if len(os.Args) != 2 {
		panic("bootrom is required")
	}

	lcd := lcd.NewLcd()
	gogameboy.Emulate(os.Args[1], lcd)
}
