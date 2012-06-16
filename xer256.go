package xer

import "math/rand"

type xer256 struct {
	sum   uint64
	c     int
	state [256]uint64
}

func new256(seed int64) rand.Source {
	x := &xer256{}
	x.Seed(seed)
	return x
}

func (x *xer256) Seed(seed int64) {
	x.sum = doSeed(uint64(seed), x.state[:])
}

func (x *xer256) Int63() int64 {
	return int63_a256(x)
}
