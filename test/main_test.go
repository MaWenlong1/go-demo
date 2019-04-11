package main

import (
	"os"
	"testing"
	"time"
)

func add(x, y int) int {
	return x + y
}
func TestAdd(t *testing.T) {
	var tests = []struct {
		x      int
		y      int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{2, 3, 5},
	}
	for _, tt := range tests {
		actual := add(tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("add(%d,%d): expect %d,actual:%d.", tt.x, tt.y, tt.expect, actual)
		}
	}
}
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}

func BenchmarkAdd(b *testing.B) {
	println("B.N =", b.N)
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}
