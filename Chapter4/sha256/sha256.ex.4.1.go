package main

import (
	"fmt"
	"crypto/sha256"
)

// pc[i] is the population count of i.
var pc [256]byte // Default initialize to 0 each position

// Init the pc variable. The value of each position is the enable bits of the index
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Returns the number of different bits between the two sha256 arrays
func diffBitsSHA256(x *[32]byte, y *[32]byte) uint8 {
	var diffBites uint8 = 0
	for i := uint8(0); i < 32; i++ {
		diffBites += pc[x[i]^y[i]]
	}
	return diffBites
}

func main(){
	c1 := sha256.Sum256([]byte("2"))
	c2 := sha256.Sum256([]byte("0"))
	fmt.Println(diffBitsSHA256(&c1, &c2))
}
