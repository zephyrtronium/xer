/ʃɛər/ - A simple and speedy pseudo-random number generator.

Xer uses the recurrence `X_n = s_i → s_{i-H mod b}`, where `s` is `X_{n-1} ⊕ … ⊕ X_{n-L}` and `H` is the Hamming weight of s. In other words, each iterate is the XOR sum of the last `L` iterates, rotated right by its popcount. The heart of this algorithm boils down to two instructions on modern processors, making it extraordinarily fast.

Example usage:

```go
import (
	"github.com/zephyrtronium/xer"
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {
	rng = xer.New(time.Now().UnixNano(), 256)
}
```

then use `rng` just as you would any other rand.Rand (or instead of the global functions in rand).

Xer comes with a hand-assembled generator on amd64 with state size 256. Benchmarks comparing xer to the source built in to math/rand show this version to be winner on my AMD FX-6100.