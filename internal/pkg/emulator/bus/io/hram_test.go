package io_test

import (
	"testing"

	"github.com/atsumarukun/go-game-boy/internal/pkg/emulator/bus/io"
)

func TestHramRead(t *testing.T) {
	hram := io.NewHram()
	val := hram.Read(0)
	if val != 0 {
		t.Errorf("Failed to read hram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestHramWrite(t *testing.T) {
	hram := io.NewHram()
	hram.Write(0, 1)
	val := hram.Read(0)
	if val != 1 {
		t.Errorf("Failed to write hram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
