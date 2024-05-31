package hram_test

import (
	"go-game-boy/internal/pkg/emulator/hram"
	"testing"
)

func TestRead(t *testing.T) {
	hram := hram.NewHram()
	val := hram.Read(0)
	if val != 0 {
		t.Errorf("Failed to read hram.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWrite(t *testing.T) {
	hram := hram.NewHram()
	hram.Write(0, 1)
	val := hram.Read(0)
	if val != 1 {
		t.Errorf("Failed to write hram.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
