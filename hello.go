// https://gobyexample.com/ is a good reference
package main

import (
	"fmt"
	"image"
	"image/draw"
)

func test() {
	var a, b int = 1, 2
	c := 3  // shorthand initialization var
	b, a = a, b
	fmt.Println(a, b, c)

	// for loops for iterating range
	for i := 1; i < 3; i++ {
	}
	for i := range 3 {
		fmt.Print(i)
	}
	fmt.Println()
	// string loop
	for i, runeVal := range "abcdef" {
		fmt.Printf("%d:%c ", i, runeVal)
	}

	// for loops are also while loops. They can take no condition (while true) or a condition.
	// they accept "continue" and "break" as well.
}

func invert(src image.Image) *image.RGBA {
	// testing function for inverting individual pixels
	bounds := src.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, src, bounds.Min, draw.Src)

	// iterate pixels (Stride = number of bytes per row)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// r,g,b,a := img.At(x, y).RGBA()
			i := (y-bounds.Min.Y)*rgba.Stride + (x-bounds.Min.X)*4
			// invert
			rgba.Pix[i+0] = 255 - rgba.Pix[i+0]
			rgba.Pix[i+1] = 255 - rgba.Pix[i+1]
			rgba.Pix[i+2] = 255 - rgba.Pix[i+2]
			// rgba.Pix[i+3] = 1  // alpha
		}
	}
	return rgba
}
