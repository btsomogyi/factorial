// Package factorial implements scalable efficient uint factorial function
package factorial

import (
	"math"
	"math/bits"
)

// Factorial returns m! for any uint, with overflow indicator 'ok' (true =
// correct output, false = overflow and zero return value).
func Factorial(m uint) (result uint, ok bool) {
	var floatResult float64
	floatResult = math.Gamma(float64(m + 1))
	var maxVal float64
	switch x := bits.UintSize; x {
	case 8:
		maxVal = float64(math.MaxUint8)
	case 16:
		maxVal = float64(math.MaxUint16)
	case 32:
		maxVal = float64(math.MaxUint32)
	case 64:
		maxVal = float64(math.MaxUint64)
	}
	if floatResult > maxVal {
		return 0, false
	}
	return uint(floatResult), true
}
