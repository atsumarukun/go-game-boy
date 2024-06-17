package cpu

type registers struct {
	a  uint8
	f  uint8
	b  uint8
	c  uint8
	d  uint8
	e  uint8
	h  uint8
	l  uint8
	sp uint16
	pc uint16
}

func (r *registers) readPC() uint16 {
	val := r.pc
	r.pc += 1
	return val
}

func (r *registers) readAF() uint16 {
	return uint16(r.a)<<8 | uint16(r.f)
}

func (r *registers) readBC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}

func (r *registers) readDE() uint16 {
	return uint16(r.d)<<8 | uint16(r.e)
}

func (r *registers) readHL() uint16 {
	return uint16(r.h)<<8 | uint16(r.l)
}

func (r *registers) writeAF(val uint16) {
	r.a = uint8(val >> 8)
	r.f = uint8(val & 0xF0) // The lower 4 bits of F register are always 0.
}

func (r *registers) writeBC(val uint16) {
	r.b = uint8(val >> 8)
	r.c = uint8(val)
}

func (r *registers) writeDE(val uint16) {
	r.d = uint8(val >> 8)
	r.e = uint8(val)
}

func (r *registers) writeHL(val uint16) {
	r.h = uint8(val >> 8)
	r.l = uint8(val)
}

func (r *registers) getZF() bool {
	return r.f&0b_1000_0000 > 0
}

func (r *registers) getNF() bool {
	return r.f&0b_0100_0000 > 0
}

func (r *registers) getHF() bool {
	return r.f&0b_0010_0000 > 0
}

func (r *registers) getCF() bool {
	return r.f&0b_0001_0000 > 0
}

func (r *registers) setZF(val bool) {
	if val {
		r.f |= 0b_1000_0000
	} else {
		r.f &= 0b_0111_1111
	}
}

func (r *registers) setNF(val bool) {
	if val {
		r.f |= 0b_0100_0000
	} else {
		r.f &= 0b_1011_1111
	}
}

func (r *registers) setHF(val bool) {
	if val {
		r.f |= 0b_0010_0000
	} else {
		r.f &= 0b_1101_1111
	}
}

func (r *registers) setCF(val bool) {
	if val {
		r.f |= 0b_0001_0000
	} else {
		r.f &= 0b_1110_1111
	}
}
