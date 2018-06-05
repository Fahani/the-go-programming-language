package main

import (
	"image/color"
	"math/cmplx"
	"fmt"
	"image/png"
	"image"
	"net/http"
	"log"
	"strconv"
)

func mandelbrot(z complex128) color.Color {
	const (
		iterations, contrast = 200, 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request){
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

	var width, height = 1024, 1024


	if r.URL.Query().Get("width") != ""{
		w, err := strconv.Atoi(r.URL.Query().Get("width"))
		if err == nil && w > 0 {
			width = int(w)
		}
	}
	if r.URL.Query().Get("height") != ""{
		h, err := strconv.Atoi(r.URL.Query().Get("height"))
		if err == nil && h > 0 {
			height = int(h)
		}
	}

	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / float64(width)*(xmax - xmin) + xmin
			z := complex(x,y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	if png.Encode(w, img) != nil {
		fmt.Fprintln(w, "Error encoding file")
	}
}

func main(){
	http.HandleFunc("/", mandelbrotHandler) // Each request calls the handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}