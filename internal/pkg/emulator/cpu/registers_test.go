package cpu_test

import (
	"go-game-boy/internal/pkg/emulator/cpu"
	"testing"
)

func TestReadAF(t *testing.T) {
	regs := cpu.Registers{}
	val := cpu.ReadAF(&regs)
	if val != 0 {
		t.Errorf("Failed to read AF register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadBC(t *testing.T) {
	regs := cpu.Registers{}
	val := cpu.ReadBC(&regs)
	if val != 0 {
		t.Errorf("Failed to read BC register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadDE(t *testing.T) {
	regs := cpu.Registers{}
	val := cpu.ReadDE(&regs)
	if val != 0 {
		t.Errorf("Failed to read DE register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestReadHL(t *testing.T) {
	regs := cpu.Registers{}
	val := cpu.ReadHL(&regs)
	if val != 0 {
		t.Errorf("Failed to read HL register.\n Expected value: %d\n Received value: %d\n", 0, val)
	}
}

func TestWriteAF(t *testing.T) {
	regs := cpu.Registers{}
	cpu.WriteAF(&regs, 0x0111) // The lower 4 bits of F register are always 0.
	val := cpu.ReadAF(&regs)
	if val != 0x0110 {
		t.Errorf("Failed to write AF register.\n Expected value: %x\n Received value: %x\n", 0x0110, val)
	}
}

func TestWriteBC(t *testing.T) {
	regs := cpu.Registers{}
	cpu.WriteBC(&regs, 0x0101)
	val := cpu.ReadBC(&regs)
	if val != 0x0101 {
		t.Errorf("Failed to write BC register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestWriteDE(t *testing.T) {
	regs := cpu.Registers{}
	cpu.WriteDE(&regs, 0x0101)
	val := cpu.ReadDE(&regs)
	if val != 0x0101 {
		t.Errorf("Failed to write DE register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestWriteHL(t *testing.T) {
	regs := cpu.Registers{}
	cpu.WriteHL(&regs, 0x0101)
	val := cpu.ReadHL(&regs)
	if val != 0x0101 {
		t.Errorf("Failed to write HL register.\n Expected value: %x\n Received value: %x\n", 0x0101, val)
	}
}

func TestGetZF(t *testing.T) {
	regs := cpu.Registers{}
	z := cpu.GetZF(&regs)
	if z {
		t.Errorf("Failed to get the zero flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, z)
	}
}

func TestGetNF(t *testing.T) {
	regs := cpu.Registers{}
	n := cpu.GetNF(&regs)
	if n {
		t.Errorf("Failed to get the subtract flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, n)
	}
}

func TestGetHF(t *testing.T) {
	regs := cpu.Registers{}
	h := cpu.GetHF(&regs)
	if h {
		t.Errorf("Failed to get the half carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, h)
	}
}

func TestGetCF(t *testing.T) {
	regs := cpu.Registers{}
	c := cpu.GetCF(&regs)
	if c {
		t.Errorf("Failed to get the carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", false, c)
	}
}

func TestSetZF(t *testing.T) {
	regs := cpu.Registers{}
	cpu.SetZF(&regs, true)
	z := cpu.GetZF(&regs)
	if !z {
		t.Errorf("Failed to set the zero flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, z)
	}
}

func TestSetNF(t *testing.T) {
	regs := cpu.Registers{}
	cpu.SetNF(&regs, true)
	n := cpu.GetNF(&regs)
	if !n {
		t.Errorf("Failed to set the subtract flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, n)
	}
}

func TestSetHF(t *testing.T) {
	regs := cpu.Registers{}
	cpu.SetHF(&regs, true)
	h := cpu.GetHF(&regs)
	if !h {
		t.Errorf("Failed to set the half carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, h)
	}
}

func TestSetCF(t *testing.T) {
	regs := cpu.Registers{}
	cpu.SetCF(&regs, true)
	c := cpu.GetCF(&regs)
	if !c {
		t.Errorf("Failed to set the carry flag in the flag register.\n Expected value: %t\n Received value: %t\n", true, c)
	}
}
