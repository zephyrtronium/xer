// +build !amd64

package xer

func int63_a256(x *xer256) int64 {
	s := linear(x.sum)
	x.sum ^= x.state[x.c] ^ s
	x.state[x.c] = s
	x.c = (x.c + 1) % len(x.state)
	return int64(s & 0x7fffffffffffffff)
}
