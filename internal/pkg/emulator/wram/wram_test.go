package wram_test

import (
	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/wram"
	"testing"
)

func TestRead(t *testing.T) {
	wram := wram.NewWram()
	val := wram.Read(0)
	if val != 0 {
		t.Errorf("Failed to read wram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWrite(t *testing.T) {
	wram := wram.NewWram()
	wram.Write(0, 1)
	val := wram.Read(0)
	if val != 1 {
		t.Errorf("Failed to write wram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
