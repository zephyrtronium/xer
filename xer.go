package xer

import "math/rand"

type xer struct {
	sum   uint64
	c     int
	state []uint64
}

func New(seed int64, stateSize int) rand.Source {
	switch stateSize {
	case 256:
		return new256(seed)
	}
	return newN(seed, stateSize)
}

func newN(seed int64, stateSize int) rand.Source {
	x := &xer{state: make([]uint64, stateSize)}
	x.Seed(seed)
	return x
}

func (x *xer) Seed(seed int64) {
	// MMIX LCG by Knuth retrieved from Wikipedia
	// I chose an LCG to seed because its state size is 1 and it has period m
	// (in this case 2⁶⁴) independent of its seed. 64-bit values are produced,
	// but only the 32 most significant bits are used. This will not run
	// through its period for state sizes less than 2⁶³ words.
	const a uint64 = 6364136223846793005
	const c uint64 = 1442695040888963407
	s := uint64(seed)
	for _ = range x.state {
		s = a*s + c
		s = a*s + c
	}
	for i := range x.state {
		s = a*s + c
		x.state[i] ^= s >> 32
		s = a*s + c
		x.state[i] ^= s & 0xffffffff00000000
	}
	// Calculate the sum once. Each new sum can be calculated with two XORs
	// per value after this.
	for _, v := range x.state {
		x.sum ^= v
	}
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
