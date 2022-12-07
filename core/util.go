package core

func RGBtoGray(r, g, b uint32) uint8 {
	return uint8((r + g + b) / 3)
}
