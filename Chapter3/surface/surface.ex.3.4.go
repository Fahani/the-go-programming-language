// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"math"
	"fmt"
	"net/http"
	"log"
)

const (
	width, height = 600, 320 // canvas size in pixels
	cells = 100 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale = width/2/xyrange // pixels per x or y unit
	zscale = height * 0.4 // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at the corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)
	// If the z is inf then return error
	if math.IsInf(z, 0) {
		return 0, 0, z, true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, false
}

func color(height float64) int {

	if height > 0.9 {
		return 0xff0000
	} else if height > 0.07 {
		return 0xff1919
	} else if height > 0.06 {
		return 0xff3232
	} else if height > 0.05 {
		return 0xff4c4c
	} else if height > 0.04 {
		return 0x6666ff
	} else if height > 0.03 {
		return 0x4c4cff
	} else if height > 0.02 {
		return 0x3232ff
	} else if height > 0.01 {
		return 0x1919ff
	} else {
		return 0x0000ff
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, err1 := corner(i+1, j)
			bx, by, z2, err2 := corner(i, j)
			cx, cy, z3, err3 := corner(i, j+1)
			dx, dy, z4, err4 := corner(i+1, j+1)

			heightAverage := z1 + z2 + z3 + z4
			color := color(heightAverage)

			// If corner returns error, then don't print the polygon
			if err1 || err2 || err3 || err4 {
				continue
			}
			fmt.Fprintf(w,"<polygon style='fill:#%06x' points='%g,%g %g,%g %g,%g %g,%g' />\n",color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w,"</svg>")
}

func main(){
	http.HandleFunc("/", handler) // Each request calls the handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}