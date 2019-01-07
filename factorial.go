package main

import (
	"math"
	"math/bits"
)

func factorial(m uint) (result uint, ok bool) {
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
