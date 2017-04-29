package identicon

import (
	"image"
	"errors"
	"image/draw"
	"image/color"
	"image/color/palette"
)

const ColorfulTilesCount int = 7

type Colorful struct {
	width, height, tiles int
}

func NewColorful(size int) *Colorful {
	return &Colorful{size, size, ColorfulTilesCount}
}

func (b *Colorful) getTileSize() int {
	return b.width / b.tiles
}

func (b *Colorful) Build(in []byte)  (*image.RGBA, error) {
	if len(in) != 16 {
		return nil, errors.New("Input for dummy must has exactly 16 bytes")
	}

	// how big tile is
	tileSize := b.getTileSize()

	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: b.width, Y: b.height}})
	i := 0

	for x := 0; x < b.width / 2 + 1; x += tileSize {
		for y := 0; y < b.height / 2 + 1; y += tileSize {
			// choose color according to input
			c := palette.Plan9[in[i]].(color.RGBA)
			i++

			// draw 4 rectangles
			draw.Draw(im, image.Rect(x, y, x + tileSize, y + tileSize), &image.Uniform{c}, image.ZP, draw.Src)
			draw.Draw(im, image.Rect(b.width - x, y, b.width - (x + tileSize), y + tileSize), &image.Uniform{c}, image.ZP, draw.Src)
			draw.Draw(im, image.Rect(x, b.height - y, x + tileSize, b.height - (y + tileSize)), &image.Uniform{c}, image.ZP, draw.Src)
			draw.Draw(im, image.Rect(b.width - x, b.height - y, b.width - (x + tileSize), b.height - (y + tileSize)), &image.Uniform{c}, image.ZP, draw.Src)
		}
	}

	return im, nil
}
