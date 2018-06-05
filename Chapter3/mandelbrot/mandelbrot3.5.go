package main

import (
	"image/color"
	"math/cmplx"
	"fmt"
	"image/png"
	"os"
	"image"
)

func mandelbrot(z complex128) color.Color {
	const (
		iterations, contrast = 200, 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast*n - 255, contrast*n - 126,0, 255}
		}
	}
	return color.RGBA{255, 126,0,255}
}

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width*(xmax - xmin) + xmin
			z := complex(x,y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	if png.Encode(os.Stdout, img) != nil {
		fmt.Fprintln(os.Stdout, "Error encoding file")
	}
}