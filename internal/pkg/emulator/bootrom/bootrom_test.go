package bootrom_test

import (
	"go-game-boy/internal/pkg/emulator/bootrom"
	"testing"
)

func TestRead(t *testing.T) {
	rom := []uint8{0, 1, 2}
	bootrom := bootrom.NewBootrom(rom)
	val := bootrom.Read(1)
	if val != 1 {
		t.Errorf("Failed to read bootrom.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestIsActive(t *testing.T) {
	rom := []uint8{0, 1, 2}
	bootrom := bootrom.NewBootrom(rom)
	active := bootrom.IsActive()
	if !active {
		t.Errorf("Failed to get bootrom active status.\n Expected value: %t\n Received value: %t\n", true, active)
	}
}

func TestToggleActive(t *testing.T) {
	rom := []uint8{0, 1, 2}
	bootrom := bootrom.NewBootrom(rom)

	bootrom.ToggleActive(0)
	active := bootrom.IsActive()
	if !active {
		t.Errorf("Failed to update active state of bootrom.\n Expected value: %t\n Received value: %t\n", true, active)
	}

	bootrom.ToggleActive(1)
	active = bootrom.IsActive()
	if active {
		t.Errorf("Failed to update bootrom to inactive status\n Expected value: %t\n Received value: %t\n", false, active)
	}
}
