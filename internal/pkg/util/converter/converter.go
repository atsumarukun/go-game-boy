package converter

func BoolToUint[T uint8 | uint16](val bool) T {
	if val {
		return 1
	}
	return 0
}
