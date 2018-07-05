// Example using embedding struct and anonymous fields
package main

import "fmt"

type Point struct {X, Y int}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	// Using names
	w = Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8, // Trailing coma necessary
			},
			Radius: 5, // Trailing coma necessary
		},
		Spokes: 20, // Trailing coma necessary

	}

	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)

}