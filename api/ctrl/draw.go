package ctrl

import (
	"errors"
	"github.com/fogleman/gg"
)

func DrawFromPixelArray(pixel [][]uint8, out string) error {
	originalWidth := 0
	originalHeight := 0
	ratio := 1

	if len(pixel) < 1 {
		return errors.New("pixel invalid")
	}

	originalHeight = len(pixel)
	originalWidth = len(pixel[0])

	width := originalWidth * ratio
	height := originalHeight * ratio

	dc := gg.NewContext(width, height)

	startX := (width - (originalWidth * ratio)) / 2
	startY := (height - (originalHeight * ratio)) / 2

	for y := 0; y < originalHeight; y++ {
		for x := 0; x < originalWidth; x++ {
			if pixel[y][x] == 1 {
				dc.SetRGB(0, 0, 0)
			} else {
				dc.SetRGB(1, 1, 1)
			}
			dc.DrawRectangle(float64(startX+x*ratio), float64(startY+y*ratio), float64(ratio), float64(ratio))
			dc.Fill()
		}
	}

	err := dc.SavePNG(out)
	if err != nil {
		return err
	}

	return nil
}
