// Write const declarations for KB, MB, up through YB as compactly as you can
package main

import (
	"fmt"
)

type Byte float64

const (
	_ Byte = 1 << (10 * iota)
	KB // 1024
	MB // 1048576
	GB // 1073741824
	TB // 1099511627776 (exceeds 1 << 32)
	PB // 1125899906842624
	EB // 1152921504606846976
	ZB // 1180591620717411303424 (exceeds 1 << 64)
	YB // 1208925819614629174706176
)

func main(){

	fmt.Printf("%s", YB)
}

func (b Byte) String() string {
	return fmt.Sprintf("%.0f bytes", b)
}