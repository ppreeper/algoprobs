package root

import (
	"math"
	"testing"
)

var mulRootTests = []struct {
	x        float64
	p        float64
	expected float64
}{
	{5, 2, math.Sqrt(5)},
	{29, 3, math.Cbrt(29)},
	{27, 3, math.Cbrt(27)},
}

func TestRoot(t *testing.T) {
	for _, mt := range mulRootTests {
		r := Root(mt.x, mt.p)
		if r != mt.expected {
			t.Errorf("\nexpected %f got %f", mt.expected, r)
		}
	}
}

// Square Root
func BenchmarkSqrt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		math.Sqrt(29)
	}
}
func BenchmarkRoot2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Root(29, 2)
	}
}

// Cube Root
func BenchmarkCbrt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		math.Cbrt(29)
	}
}
func BenchmarkRoot3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Root(29, 3)
	}
}
