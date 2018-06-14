// This example prints and compares the SHA256 digests of "x" and "X"
package sha256

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%b %b",[]byte("x"),[]byte("X") )
}
