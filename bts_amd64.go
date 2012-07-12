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

func init() {
	if havePopcnt() {
		int63_a256 = int63_popcnt256
		int63_a65536 = int63_popcnt65536
	} else {
		int63_a256 = int63_basic256
		int63_a65536 = int63_basic65536
	}
}

var int63_a256 func(*xer256) int64
var int63_a65536 func (*xer65536) int64

func int63_basic256(*xer256) int64
func int63_popcnt256(*xer256) int64
func int63_basic65536(*xer65536) int64
func int63_popcnt65536(*xer65536) int64

func havePopcnt() bool
