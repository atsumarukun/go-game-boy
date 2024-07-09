package bootrom_test

import (
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bootrom"
	"testing"
)

func TestIsActive(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	active := bootrom.IsActive()
	if !active {
		t.Errorf("Failed to get bootrom active status.\n Expected value: %t\n Received value: %t\n", true, active)
	}
}

func TestRead(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})
	val := bootrom.Read(1)
	if val != 1 {
		t.Errorf("Failed to read bootrom.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}

func TestWrite(t *testing.T) {
	bootrom := bootrom.NewBootrom([]uint8{0, 1, 2})

	bootrom.Write(0xFF50, 0)
	active := bootrom.IsActive()
	if !active {
		t.Errorf("Failed to update active state of bootrom.\n Expected value: %t\n Received value: %t\n", true, active)
	}

	bootrom.Write(0xFF50, 1)
	active = bootrom.IsActive()
	if active {
		t.Errorf("Failed to update bootrom to inactive status\n Expected value: %t\n Received value: %t\n", false, active)
	}
}
