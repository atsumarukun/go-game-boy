package bus_test

import (
	"go-game-boy/internal/pkg/emulator/bootrom"
	"go-game-boy/internal/pkg/emulator/bus"
	"go-game-boy/internal/pkg/emulator/hram"
	"go-game-boy/internal/pkg/emulator/oam"
	"go-game-boy/internal/pkg/emulator/vram"
	"go-game-boy/internal/pkg/emulator/wram"
	"testing"
)

func TestReadBootrom(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	bus := bus.NewBus(bootrom, nil, nil, nil, nil, nil)
	val := bus.Read(1)
	if val != 1 {
		t.Errorf("Failed to read bootrom.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestWriteBootrom(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	bus := bus.NewBus(bootrom, nil, nil, nil, nil, nil)
	bus.Write(0xFF50, 1)
	active := bootrom.IsActive()
	if active {
		t.Errorf("Failed to update bootrom to inactive status\n Expected value: %t\n Received value: %t\n", false, active)
	}
}

func TestReadVram(t *testing.T) {
	vram := vram.NewVram()
	bus := bus.NewBus(nil, vram, nil, nil, nil, nil)
	val := bus.Read(0x8000)
	if val != 0 {
		t.Errorf("Failed to read vram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteVram(t *testing.T) {
	vram := vram.NewVram()
	bus := bus.NewBus(nil, vram, nil, nil, nil, nil)
	bus.Write(0x8000, 1)
	val := bus.Read(0x8000)
	if val != 1 {
		t.Errorf("Failed to write vram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestReadWram(t *testing.T) {
	wram := wram.NewWram()
	bus := bus.NewBus(nil, nil, wram, nil, nil, nil)
	val := bus.Read(0xC000)
	if val != 0 {
		t.Errorf("Failed to read wram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteWram(t *testing.T) {
	wram := wram.NewWram()
	bus := bus.NewBus(nil, nil, wram, nil, nil, nil)
	bus.Write(0xC000, 1)
	val := bus.Read(0xC000)
	if val != 1 {
		t.Errorf("Failed to write wram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestReadHram(t *testing.T) {
	hram := hram.NewHram()
	bus := bus.NewBus(nil, nil, nil, hram, nil, nil)
	val := bus.Read(0xFF80)
	if val != 0 {
		t.Errorf("Failed to read hram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteHram(t *testing.T) {
	hram := hram.NewHram()
	bus := bus.NewBus(nil, nil, nil, hram, nil, nil)
	bus.Write(0xFF80, 1)
	val := bus.Read(0xFF80)
	if val != 1 {
		t.Errorf("Failed to write hram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestReadOam(t *testing.T) {
	oam := oam.NewOam()
	bus := bus.NewBus(nil, nil, nil, nil, oam, nil)
	val := bus.Read(0xFE00)
	if val != 0 {
		t.Errorf("Failed to read oam\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteOam(t *testing.T) {
	oam := oam.NewOam()
	bus := bus.NewBus(nil, nil, nil, nil, oam, nil)
	bus.Write(0xFE00, 1)
	val := bus.Read(0xFE00)
	if val != 1 {
		t.Errorf("Failed to write oam\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
