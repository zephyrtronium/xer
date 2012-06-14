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
