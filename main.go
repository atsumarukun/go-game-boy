package gogameboy

import (
	"os"

	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/lcd"
)

const (
	LCD_WIDTH  = lcd.WIDTH
	LCD_HEIGHT = lcd.HEIGHT
)

func Emulate(bootromPath string, lcd lcd.Lcd) {
	bootrom, err := os.ReadFile(bootromPath)
	if err != nil {
		panic(err)
	}

	gameboy := emulator.Init(bootrom, lcd)
	gameboy.Emulate()
}
