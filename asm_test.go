// +build amd64

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