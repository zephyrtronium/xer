// Copyright 2012 Branden J Brown
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
//    1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would be
//    appreciated but is not required.
//
//    2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
//    3. This notice may not be removed or altered from any source
//    distribution.

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