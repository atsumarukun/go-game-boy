package emulator

import (
	"time"

	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus/io"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/cpu"
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/lcd"
)

const (
	CPU_CLOCK_HZ  uint = 4_194_304
	M_CYCLE_CLOCK uint = 4
	M_CYCLE_NANOS uint = M_CYCLE_CLOCK * 1_000_000_000 / CPU_CLOCK_HZ
)

type Emulator interface {
	Emulate()
}

type emulator struct {
	cpu *cpu.Cpu
	bus bus.Bus
	lcd lcd.Lcd
}

func Init(bootromBody []byte, lcd lcd.Lcd) Emulator {
	cpu := cpu.NewCpu()
	ppu := io.NewPpu()

	bootrom := io.NewBootrom(bootromBody)
	vram := io.NewVram()
	wram := io.NewWram()
	hram := io.NewHram()
	oam := io.NewOam()

	bus := bus.NewBus(bootrom, vram, wram, hram, oam, ppu)

	return &emulator{
		cpu,
		bus,
		lcd,
	}
}

func (e *emulator) Emulate() {
	start := time.Now()
	elapsed := 0

	for {
		ela := time.Since(start)
		for i := 0; i < int(ela-time.Duration(elapsed)); i++ {
			e.cpu.Emulate(e.bus)
			if e.bus.Ppu().Emulate(e.bus.Vram()) {
				e.lcd.Render(e.bus.Ppu().Buffer())
			}
			elapsed += int(M_CYCLE_NANOS)
		}
	}
}
