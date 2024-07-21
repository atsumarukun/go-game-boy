package io

type Io interface {
	Read(addr uint16) uint8
	Write(addr uint16, val uint8)
}

type DefaultIo struct{}

func (d DefaultIo) Read(_ uint16) uint8 {
	return 0xFF
}

func (d DefaultIo) Write(_ uint16, _ uint8) {}
