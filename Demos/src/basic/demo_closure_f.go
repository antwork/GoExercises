package main

import "fmt"

func plusTwo(num int) func(a int) int {
	a := func(a int) int {
		return a + num
	}
	return a
}

func main() {
	p := plusTwo(10)
	fmt.Printf("%v\n", p(2))

	fmt.Printf("%v\n", p(3))

	fmt.Printf("%v\n", p(8))
}
