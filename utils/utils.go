package utils

import (
	"fmt"
	"image/color"
)

// Konversi warna HEX menjadi format RGBA
func ParseHexColor(hex string) color.RGBA { // GANTI HURUF AWAL MENJADI BESAR
	var r, g, b uint8
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return color.RGBA{r, g, b, 255}
}
