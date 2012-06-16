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
	x.sum = doSeed(uint64(seed), x.state[:])
}

func (x *xer65536) Int63() int64 {
	return int63_a65536(x)
}
