package io_test

import (
	"testing"

	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus/io"
)

func TestOamRead(t *testing.T) {
	oam := io.NewOam()
	val := oam.Read(0)
	if val != 0 {
		t.Errorf("Failed to read oam.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestOamWrite(t *testing.T) {
	oam := io.NewOam()
	oam.Write(0, 1)
	val := oam.Read(0)
	if val != 1 {
		t.Errorf("Failed to write oam.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
