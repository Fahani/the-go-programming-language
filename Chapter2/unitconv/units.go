// This file defines the unit's type and the functions attached to the named types
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64
type Kilogram float64
type Pound float64

const (
	BoilingC Celsius = 100
	FreezingC Celsius = 0
	BoilingF Fahrenheit = 212
	FreezingF Fahrenheit = 32
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g ºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g ºF", f)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}