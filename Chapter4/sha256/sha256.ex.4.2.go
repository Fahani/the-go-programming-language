// prints the SHA256 hash of its standard input by default but accept the flag -sha to specify 384 or 512
package main

import (
	"flag"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)
var sha = flag.Int("sha", 256, "256 (default): Using 256 bits to digest the hash\n" +
	"384: Using 384 bits to digest the hash\n512: Using 512 bits to digest the hash")

func digestString(s string) {
	switch *sha {
	case 384:
		fmt.Printf("%s=%x\n", s, sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("%s=%x\n", s, sha512.Sum512([]byte(s)))
	default:
		fmt.Printf("%s=%x\n", s, sha256.Sum256([]byte(s)))
	}
}

func init(){
	flag.Parse()
}

func main(){
	for _,arg := range flag.Args() {
		digestString(arg)
	}
}