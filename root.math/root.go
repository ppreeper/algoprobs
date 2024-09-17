package root

import (
	"math"
)

func Root(x, p float64) float64 {
	switch {
	case p == 2:
		return math.Sqrt(x)
	case p == 3:
		return math.Cbrt(x)
	case x == 0 || IsNaN(x) || IsInf(x, 1):
		return x
	case x < 0 && math.Mod(x, 2) == 0:
		return x
	}
	y := x / p
	for i := 0; i < 32; i++ {
		y = (x/math.Pow(y, (p-1)) + y*(p-1)) / p
	}
	return y
}
