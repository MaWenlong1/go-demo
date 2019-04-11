package sprintf_test

import (
	"fmt"
	"strconv"
	"testing"
)

// go test -v -run="none" -bench="BenchmarkSprintf" -benchtime="3s" -benchmem
// go test -v -run="none" -bench=. -benchtime="3s" -benchmem
func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}
func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
