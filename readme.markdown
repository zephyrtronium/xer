/ʃɛər/ - A simple and speedy pseudo-random number generator.

Xer uses the recurrence `X_n = s_i → s_{i-H mod b}`, where `s` is `X_{n-1} ⊕ … ⊕ X_{n-L}`, `H` is the Hamming weight of `s`, and `b` is bit width (always 64 in this implementation). In other words, each iterate is the XOR sum of the last `L` iterates, rotated right by its popcount. The heart of this algorithm boils down to two instructions on modern processors, making it extraordinarily fast.

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

Xer comes with hand-assembled generators on amd64 with state sizes 256 and 65536. Benchmarks comparing xer to the source built in to math/rand show xer256 to be winner on my AMD FX-6100. xer65536 was originally a joke; I thought it would thrash something fierce, but surprisingly, it was performant enough that it is staying.

No formal analysis has been done on xer's quality or periods, because I don't hold a doctorate in mathematics, but 200 000 000 random numbers generated from xer256 came through bzip2 with a compression ratio of 1.00474750625, and compression ratios between 2 000 000 xer65536 numbers through bzip2, gzip, and xz all were above 1. "Congratulations, smells like entropy." Given that it consists entirely of linear operations, however, it's safe to say that xer is not suitable for crytography.