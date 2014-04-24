package simplemath

import (
	"math"
)

func Sqrt(x int) int {
	v := math.Sqrt(float64(x))
	return int(v)
}
