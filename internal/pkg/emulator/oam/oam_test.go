package oam_test

import (
	"go-game-boy/internal/pkg/emulator/oam"
	"testing"
)

func TestRead(t *testing.T) {
	oam := oam.NewOam()
	val := oam.Read(0)
	if val != 0 {
		t.Errorf("Failed to read oam.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWrite(t *testing.T) {
	oam := oam.NewOam()
	oam.Write(0, 1)
	val := oam.Read(0)
	if val != 1 {
		t.Errorf("Failed to write oam.\n Expected value: %d\n Received value: %d\n", 1, val)
	}
}
