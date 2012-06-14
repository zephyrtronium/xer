package xer

import (
	"math/rand"
	"testing"
)

func BenchmarkRandSrc(b *testing.B) {
	b.StopTimer()
	s := rand.NewSource(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
}

func BenchmarkXer256(b *testing.B) {
	b.StopTimer()
	s := New(0, 256)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
}

func BenchmarkXer312(b *testing.B) {
	b.StopTimer()
	s := New(0, 312) // same size as MT19937
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
}

func BenchmarkXer65536(b *testing.B) {
	b.StopTimer()
	s := New(0, 65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
}

func BenchmarkRandSrcSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.NewSource(0)
	}
}

func BenchmarkXer256Seed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(0, 256)
	}
}

func BenchmarkXer312Seed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(0, 312)
	}
}

func BenchmarkXer65536Seed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(0, 65536)
	}
}