package main

import "fmt"

func compareNum(a, b int) (int, int) {
	if a >= b {
		return b, a
	} else {
		return a, b
	}
}

func main() {
	resultA, resultB := compareNum(2, 7)
	fmt.Printf("result: %d %d\n", resultA, resultB)

	resultC, resultD := compareNum(7, 2)
	fmt.Printf("result: %d %d\n", resultC, resultD)
}
