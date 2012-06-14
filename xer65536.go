package xer

import "math/rand"

type xer65536 struct {
	sum   uint64
	c     int
	state [65536]uint64
}

func new65536(seed int64) rand.Source {
	x := &xer65536{}
	x.Seed(seed)
	return x
}

func (x *xer65536) Seed(seed int64) {
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
	for _, v := range x.state {
		x.sum ^= v
	}
}

func (x *xer65536) Int63() int64 {
	return int63_a65536(x)
}
