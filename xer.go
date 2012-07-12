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

// A simple and speedy pseudo-random number generator.
package xer

import "math/rand"

type xer struct {
	sum   uint64
	c     int
	state []uint64
}

// Create a xer PRNG. Using the same seed for generators of different sizes
// will produce different sequences starting possibly with the first value.
// stateSize is the number of 64-bit words of state; larger values imply
// longer periods. Recommended values are 256 or 65536 for best performance.
func New(seed int64, stateSize int) rand.Source {
	switch stateSize {
	case 256:
		return new256(seed)
	case 65536:
		return new65536(seed)
	}
	return newN(seed, stateSize)
}

func newN(seed int64, stateSize int) rand.Source {
	x := &xer{state: make([]uint64, stateSize)}
	x.Seed(seed)
	return x
}

func doSeed(s uint64, state []uint64) (sum uint64) {
	// MMIX LCG by Knuth retrieved from Wikipedia
	// I chose an LCG to seed because its state size is 1 and it has period m
	// (in this case 2⁶⁴) independent of its seed. 64-bit values are produced,
	// but only the 32 most significant bits are used. This will not run
	// through its period for state sizes less than 2⁶³ words.
	const a uint64 = 6364136223846793005
	const c uint64 = 1442695040888963407
	const fuse = 20
	for i := 0; i < fuse; i++ {
		s = a*s + c
		s = a*s + c
	}
	for i := range state {
		s = a*s + c
		state[i] ^= s >> 32
		s = a*s + c
		state[i] ^= s & 0xffffffff00000000
		sum ^= state[i]
	}
	imcheating := &xer{sum, 0, state}
	for _ = range state {
		imcheating.Int63()
	}
	return imcheating.sum
}

func (x *xer) Seed(seed int64) {
	x.sum = doSeed(uint64(seed), x.state)
}

func (x *xer) Int63() int64 {
	// X_n = s_i → s_{i-H mod b}
	// where s is X_{n-1} ⊕ … ⊕ X_{n-L} and H is the Hamming weight of s.
	// In other words, each element of the sequence is the XOR sum of the last
	// L elements, rotated cyclically right in b bits by its popcount.
	s := linear(x.sum)
	x.sum ^= x.state[x.c] ^ s
	x.state[x.c] = s
	x.c = (x.c + 1) % len(x.state)
	return int64(s & 0x7fffffffffffffff)
}

func linear(v uint64) uint64 {
	// popcount_2() from http://en.wikipedia.org/wiki/Hamming_weight
	p := v - ((v >> 1) & 0x5555555555555555)
	p = (p & 0x3333333333333333) + ((p >> 2) & 0x3333333333333333)
	p = (p + (p >> 4)) & 0x0f0f0f0f0f0f0f0f
	p += p >> 8
	p += p >> 16
	p += p >> 32
	p &= 0x7f
	// cyclic shift right in 64 bits
	return (v >> p) | (v << (64 - p))
}
