// +build !amd64

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
	
func int63_aany(x *xer) int64 {
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

func int63_a256(x *xer256) int64 {
	s := linear(x.sum)
	x.sum ^= x.state[x.c] ^ s
	x.state[x.c] = s
	x.c = (x.c + 1) & 255
	return int64(s & 0x7fffffffffffffff)
}

func int63_a65536(x *xer65536) int64 {
	s := linear(x.sum)
	x.sum ^= x.state[x.c] ^ s
	x.state[x.c] = s
	x.c = (x.c + 1) & 65535
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