// Using the package unitconv to convert the parameter or the keyboard input to different units
package main

import (
	"os"
	"fmt"
	"strconv"
	"./unitconv"
	"bufio"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			argToUnits(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			argToUnits(input.Text())
		}
	}
}

func argToUnits(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := unitconv.Fahrenheit(t)
	c := unitconv.Celsius(t)
	m := unitconv.Meter(t)
	ft := unitconv.Feet(t)
	kg := unitconv.Kilogram(t)
	lb := unitconv.Pound(t)

	fmt.Printf("%s = %s,\t%s = %s,\t%s = %s,\t%s = %s,\t%s = %s,\t%s = %s\t\n", f, unitconv.FToC(f), c, unitconv.CToF(c),
		m, unitconv.MToF(m), ft, unitconv.FToM(ft), kg, unitconv.KToP(kg), lb, unitconv.PToK(lb))
}