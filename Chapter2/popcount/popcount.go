package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte // Default initialize to 0 each position

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountEx23 returns the population count (number of set bits) of x. Using a for
func PopCountEx23(x uint64) int {
	var numBits int = 0
	for i:=uint64(0); i < 8; i++ {
		numBits += int(pc[byte(x>>(0*8))])
	}
	return numBits
}

// PopCountEx24 returns the population count (number of set bits) of x. shifting its the parameter x 64 times
func PopCountEx24(x uint64) int {
	var numBits = 0

	for i:=uint64(0); i < 64; i++ {
		numBits += int(x&1)
		x = x>>1
	}
	return numBits
}

// PopCountEx25 returns the population count (number of set bits) of x. Using the expression x&(x-1)
func PopCountEx25(x uint64) int {
	var numBits = 0

	for x !=0 {
		x = x&(x-1)
		numBits ++
	}
	return numBits
}

func main() {
	fmt.Printf("%v",PopCountEx25(18446744073709551615))
}