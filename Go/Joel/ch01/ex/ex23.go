package exercise

import (
	"fmt"
)

func Ex2_3() {
	fmt.Println(`
		here we do a inequation, so the idea is to know when this is true
		100n^2 < 2^n  ==>  ∈[0,∞) \ [1,15]
		it means that for n >= 15, the 100n^2 is faster than 2^n.
	`)
}