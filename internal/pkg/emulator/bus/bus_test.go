package bus_test

import (
	"go-game-boy/internal/pkg/emulator/bootrom"
	"go-game-boy/internal/pkg/emulator/bus"
	"go-game-boy/internal/pkg/emulator/hram"
	"go-game-boy/internal/pkg/emulator/wram"
	"testing"
)

func TestReadBootrom(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	val := bus.Read(1)
	if val != 1 {
		t.Errorf("Failed to read bootrom.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestWriteBootrom(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	bus.Write(0xFF50, 1)
	active := bootrom.IsActive()
	if active {
		t.Errorf("Failed to update bootrom to inactive status\n Expected value: %t\n Received value: %t\n", false, active)
	}
}

func TestReadWram(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	val := bus.Read(0xC000)
	if val != 0 {
		t.Errorf("Failed to read wram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteWram(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	bus.Write(0xC000, 1)
	val := bus.Read(0xC000)
	if val != 1 {
		t.Errorf("Failed to write wram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestReadHram(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	val := bus.Read(0xFF80)
	if val != 0 {
		t.Errorf("Failed to read hram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteHram(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	wram := wram.NewWram()
	hram := hram.NewHram()
	bus := bus.NewBus(bootrom, wram, hram)
	bus.Write(0xFF80, 1)
	val := bus.Read(0xFF80)
	if val != 1 {
		t.Errorf("Failed to write hram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
