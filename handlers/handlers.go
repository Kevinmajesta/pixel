package handlers

import (
	"bytes"
	"image"
	"image/png"

	"github.com/Kevinmajesta/pixel/utils"
	"github.com/gofiber/fiber/v2"
)

// Struktur data yang diterima dari frontend
type PixelData struct {
	Pixels     []string `json:"pixels"`
	GridWidth  int      `json:"gridWidth"`
	GridHeight int      `json:"gridHeight"`
	PixelSize  int      `json:"pixelSize"`
}

func SavePixelArt(c *fiber.Ctx) error {
	var data PixelData
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Buat gambar dengan ukuran yang benar
	img := image.NewRGBA(image.Rect(0, 0, data.GridWidth*data.PixelSize, data.GridHeight*data.PixelSize))

	// Loop untuk mengisi warna pada tiap piksel
	for i, hexColor := range data.Pixels {
		if hexColor == "" {
			continue // Jika ada warna kosong, skip
		}

		// Perhitungan posisi x dan y
		x := (i % data.GridWidth) * data.PixelSize
		y := (i / data.GridWidth) * data.PixelSize
		parsedColor := utils.ParseHexColor(hexColor)

		// Mengisi kotak piksel dengan ukuran yang sesuai
		for dx := 0; dx < data.PixelSize; dx++ {
			for dy := 0; dy < data.PixelSize; dy++ {
				img.Set(x+dx, y+dy, parsedColor)
			}
		}
	}

	// Simpan gambar ke buffer agar bisa dikirim sebagai response
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	// Kirim file PNG sebagai response dengan header download
	c.Set("Content-Type", "image/png")
	c.Set("Content-Disposition", "attachment; filename=pixel_art.png")
	return c.Send(buf.Bytes())
}
