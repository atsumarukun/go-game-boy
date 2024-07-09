package emulator

import (
	"go-game-boy/internal/pkg/emulator/bus"
	"go-game-boy/internal/pkg/emulator/cpu"
	"go-game-boy/internal/pkg/emulator/ppu"
	"time"
)

const (
	CPU_CLOCK_HZ  uint = 4_194_304
	M_CYCLE_CLOCK uint = 4
	M_CYCLE_NANOS uint = M_CYCLE_CLOCK * 1_000_000_000 / CPU_CLOCK_HZ
)

type Lcd interface {
	Render([ppu.WIDTH * ppu.HEIGHT]uint8)
}

type Emulator struct {
	cpu *cpu.Cpu
	ppu *ppu.Ppu
	bus *bus.Bus
	lcd Lcd
}

func NewEmulator(cpu *cpu.Cpu, ppu *ppu.Ppu, bus *bus.Bus, lcd Lcd) *Emulator {
	return &Emulator{
		cpu,
		ppu,
		bus,
		lcd,
	}
}

func (e *Emulator) Emulate() {
	start := time.Now()
	elapsed := 0

	for {
		ela := time.Since(start)
		for i := 0; i < int(ela-time.Duration(elapsed)); i++ {
			e.cpu.Emulate(e.bus)
			if e.ppu.Emulate(e.bus) {
				e.lcd.Render(e.ppu.Buffer())
			}
			elapsed += int(M_CYCLE_NANOS)
		}
	}
}
