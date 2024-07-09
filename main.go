package gogameboy

import (
	"go-game-boy/internal/pkg/emulator"
	"go-game-boy/internal/pkg/emulator/bootrom"
	"go-game-boy/internal/pkg/emulator/bus"
	"go-game-boy/internal/pkg/emulator/cpu"
	"go-game-boy/internal/pkg/emulator/hram"
	"go-game-boy/internal/pkg/emulator/oam"
	"go-game-boy/internal/pkg/emulator/ppu"
	"go-game-boy/internal/pkg/emulator/vram"
	"go-game-boy/internal/pkg/emulator/wram"
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
