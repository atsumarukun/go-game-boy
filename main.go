package gogameboy

import (
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bootrom"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/cpu"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/hram"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/oam"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/ppu"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/vram"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/wram"
	"os"
)

func Emulate(bootromPath string, lcd emulator.Lcd) {
	rom, err := os.ReadFile(bootromPath)
	if err != nil {
		panic(err)
	}

	cpu := cpu.NewCpu()
	ppu := ppu.NewPpu()

	bootrom := bootrom.NewBootrom(rom)
	vram := vram.NewVram()
	wram := wram.NewWram()
	hram := hram.NewHram()
	oam := oam.NewOam()

	bus := bus.NewBus(bootrom, vram, wram, hram, oam, ppu)

	emulator := emulator.NewEmulator(cpu, ppu, bus, lcd)

	emulator.Emulate()
}
