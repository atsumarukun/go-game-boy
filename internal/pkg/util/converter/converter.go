package converter

func BoolToUint8(val bool) uint8 {
	if val {
		return 1
	}
	return 0
}
