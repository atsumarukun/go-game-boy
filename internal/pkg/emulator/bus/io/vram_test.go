package io_test

import (
	"testing"

	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus/io"
)

func TestVramRead(t *testing.T) {
	vram := io.NewVram()
	val := vram.Read(0)
	if val != 0 {
		t.Errorf("Failed to read vram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestVramWrite(t *testing.T) {
	vram := io.NewVram()
	vram.Write(0, 1)
	val := vram.Read(0)
	if val != 1 {
		t.Errorf("Failed to write vram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
