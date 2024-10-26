package controller

import (
	"errors"
	"image"
	"image/color"
)

var (
	redBayerFilter = [][]uint32{
		{1, 0},
		{0, 0},
	}

	greenBayerFilter = [][]uint32{
		{0, 1},
		{1, 0},
	}

	blueBayerFilter = [][]uint32{
		{0, 0},
		{0, 1},
	}
)

func Mosaic(img *image.RGBA) error {
	if img == nil {
		return errors.New("image is nil")
	}

	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()

	for y := 0; y < imageHeight; y += 2 {
		redPatternRow := redBayerFilter[y%2]
		greenPatternRow := greenBayerFilter[y%2]
		bluePatternRow := blueBayerFilter[y%2]

		for x := 0; x < imageWidth; x += 2 {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img.At(x+1, y).RGBA()

			r1 = r1 * redPatternRow[0]
			g1 = g1 * greenPatternRow[0]
			b1 = b1 * bluePatternRow[0]
			r2 = r2 * redPatternRow[1]
			g2 = g2 * greenPatternRow[1]
			b2 = b2 * bluePatternRow[1]

			img.Set(x, y, color.RGBA{uint8(r1), uint8(g1), uint8(b1), 255})
			img.Set(x+1, y, color.RGBA{uint8(r2), uint8(g2), uint8(b2), 255})

		}

	}

	return nil
}
