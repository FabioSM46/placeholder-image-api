package imagegen

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func GenerateImage(width, height int) image.Image {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a gray background
	gray := color.RGBA{128, 128, 128, 255} // Gray color
	draw.Draw(img, img.Bounds(), &image.Uniform{gray}, image.Point{}, draw.Src)

	// Create the resolution text
	resolution := fmt.Sprintf("%dx%d", width, height)

	// Draw the resolution text in the center
	addLabel(img, resolution, width, height)

	return img
}

func addLabel(img *image.RGBA, label string, width, height int) {
	// Set the font
	face := basicfont.Face7x13 

	// Calculate text dimensions
	labelWidth := font.MeasureString(face, label).Round()
	labelHeight := face.Metrics().Height.Round()

	// Calculate the position to center the text
	x := (width - labelWidth) / 2
	y := (height + labelHeight) / 2

	// Draw the text
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White), // Text color
		Face: face,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(label)
}