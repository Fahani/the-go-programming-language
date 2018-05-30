package popcount_test

import (
	"testing"
	"theGoProgrammingLanguage/Chapter2/popcount"
)

func BenchmarkPopCount(b *testing.B){
	for i := 0; i < b.N; i++ {
		popcount.PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountEx23(b *testing.B){
	for i := 0; i < b.N; i++ {
		popcount.PopCountEx23(18446744073709551615)
	}
}
