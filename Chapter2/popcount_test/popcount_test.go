package popcount_test

import (
	"testing"
	"theGoProgrammingLanguage/Chapter2/popcount"
)

func BenchmarkPopCount(b *testing.B){
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(b.N))
	}
}
