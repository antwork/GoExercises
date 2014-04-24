package main

import "fmt"

func average(s []float64) (result float64) {

	if len(s) == 0 {
		result = 0
	} else {
		result = 0.0
		for _, v := range s {
			result += v
		}
		result /= float64(len(s))
	}

	return
}

func main() {
	s := []float64{1.2, 1.4, 1.6}
	average := average(s)
	fmt.Printf("average of %f is %f\n", s, average)
}
