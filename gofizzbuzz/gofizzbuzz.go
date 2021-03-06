package gofizzbuzz

import (
	"fmt"
	"math"
)

// GoFizzBuzz returns a string
// based on a set of simple rules
func GoFizzBuzz(i int) (w string) {
	f := float64(i)
	switch {
	case math.Mod(f, 15) == 0:
		return "fizzbuzz"
	case math.Mod(f, 5) == 0:
		return "buzz"
	case math.Mod(f, 3) == 0:
		return "fizz"
	}
	return fmt.Sprintf("%d", i)
}
