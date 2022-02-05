package math

import "time"

func Add(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}
