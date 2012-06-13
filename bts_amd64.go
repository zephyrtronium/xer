package xer

func init() {
	if havePopcnt() {
		int63_a256 = int63_popcnt256
	} else {
		int63_a256 = int63_basic256
	}
}

var int63_a256 func(*xer256) int64

func int63_basic256(*xer256) int64
func int63_popcnt256(*xer256) int64

func havePopcnt() bool
