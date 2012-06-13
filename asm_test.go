// +build amd64

package xer

import "testing"

func TestMyASMActuallyWorks(t *testing.T) {
	xn := newN(0, 256)
	x256 := new256(0)
	t.Log("have popcnt:", havePopcnt())
	for i := 0; i < 1<<16; i++ {
		if a, b := x256.Int63(), xn.Int63(); a != b {
			t.Fatalf("mismatched generators (%d iters): got %d, expected %d\n", i, a, b)
		}
	}
	t.Log("forcing no popcnt")
	temp := int63_a256
	int63_a256 = int63_basic256
	for i := 0; i < 1<<16; i++ {
		if a, b := x256.Int63(), xn.Int63(); a != b {
			t.Fatalf("mismatched generators (%d iters): got %d, expected %d\n", i, a, b)
		}
	}
	int63_a256 = temp
}

func BenchmarkXer256ForceNoPopcnt(b *testing.B) {
	b.StopTimer()
	if !havePopcnt() {
		b.Log("don't have popcnt anyway; this is a throwaway benchmark")
	}
	s := new256(0)
	temp := int63_a256
	int63_a256 = int63_basic256
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
	b.StopTimer()
	int63_a256 = temp
}