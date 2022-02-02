package math

import "math"

func Add(a int, b int) int {
	sum := 0
	for i := 0; i < int(math.Abs(float64(a))); i++ {
		if a < 0 {
			sum--
		} else {
			sum++
		}

	}
	for i := 0; i < int(math.Abs(float64(b))); i++ {
		if b < 0 {
			sum--
		} else {
			sum++
		}
	}
	return sum
}

func Minus(a int, b int) int {
	return a - b
}
