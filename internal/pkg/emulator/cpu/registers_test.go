package cpu_test

import (
	"go-game-boy/internal/pkg/emulator/cpu"
	"testing"
)

func TestReadAF(t *testing.T) {
	registers := cpu.Registers{}
	val := cpu.ReadAF(&registers)
	if val != 0 {
		t.Errorf("Failed to read AF register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadBC(t *testing.T) {
	registers := cpu.Registers{}
	val := cpu.ReadBC(&registers)
	if val != 0 {
		t.Errorf("Failed to read BC register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadDE(t *testing.T) {
	registers := cpu.Registers{}
	val := cpu.ReadDE(&registers)
	if val != 0 {
		t.Errorf("Failed to read DE register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadHL(t *testing.T) {
	registers := cpu.Registers{}
	val := cpu.ReadHL(&registers)
	if val != 0 {
		t.Errorf("Failed to read HL register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteAF(t *testing.T) {
	registers := cpu.Registers{}
	cpu.WriteAF(&registers, 0x0111) // The lower 4 bits of F register are always 0.
	val := cpu.ReadAF(&registers)
	if val != 0x0110 {
		t.Errorf("Failed to write AF register.\n Expected value: %x\n Received value: %x\n", 0x0110, val)
	}
}

func TestWriteBC(t *testing.T) {
	registers := cpu.Registers{}
	cpu.WriteBC(&registers, 0x0101)
	val := cpu.ReadBC(&registers)
	if val != 0x0101 {
		t.Errorf("Failed to write BC register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestWriteDE(t *testing.T) {
	registers := cpu.Registers{}
	cpu.WriteDE(&registers, 0x0101)
	val := cpu.ReadDE(&registers)
	if val != 0x0101 {
		t.Errorf("Failed to write DE register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestWriteHL(t *testing.T) {
	registers := cpu.Registers{}
	cpu.WriteHL(&registers, 0x0101)
	val := cpu.ReadHL(&registers)
	if val != 0x0101 {
		t.Errorf("Failed to write HL register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestGetZF(t *testing.T) {
	registers := cpu.Registers{}
	z := cpu.GetZF(&registers)
	if z {
		t.Errorf("Failed to get the zero flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, z)
	}
}

func TestGetNF(t *testing.T) {
	registers := cpu.Registers{}
	n := cpu.GetNF(&registers)
	if n {
		t.Errorf("Failed to get the subtract flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, n)
	}
}

func TestGetHF(t *testing.T) {
	registers := cpu.Registers{}
	h := cpu.GetHF(&registers)
	if h {
		t.Errorf("Failed to get the half carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, h)
	}
}

func TestGetCF(t *testing.T) {
	registers := cpu.Registers{}
	c := cpu.GetCF(&registers)
	if c {
		t.Errorf("Failed to get the carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, c)
	}
}

func TestSetZF(t *testing.T) {
	registers := cpu.Registers{}
	cpu.SetZF(&registers, true)
	z := cpu.GetZF(&registers)
	if !z {
		t.Errorf("Failed to set the zero flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, z)
	}
}

func TestSetNF(t *testing.T) {
	registers := cpu.Registers{}
	cpu.SetNF(&registers, true)
	n := cpu.GetNF(&registers)
	if !n {
		t.Errorf("Failed to set the subtract flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, n)
	}
}

func TestSetHF(t *testing.T) {
	registers := cpu.Registers{}
	cpu.SetHF(&registers, true)
	h := cpu.GetHF(&registers)
	if !h {
		t.Errorf("Failed to set the half carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, h)
	}
}

func TestSetCF(t *testing.T) {
	registers := cpu.Registers{}
	cpu.SetCF(&registers, true)
	c := cpu.GetCF(&registers)
	if !c {
		t.Errorf("Failed to set the carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, c)
	}
}
