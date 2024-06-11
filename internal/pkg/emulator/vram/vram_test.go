package vram_test

import (
	"go-game-boy/internal/pkg/emulator/vram"
	"testing"
)

func TestRead(t *testing.T) {
	vram := vram.NewVram()
	val := vram.Read(0)
	if val != 0 {
		t.Errorf("Failed to read vram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWrite(t *testing.T) {
	vram := vram.NewVram()
	vram.Write(0, 1)
	val := vram.Read(0)
	if val != 1 {
		t.Errorf("Failed to write vram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
