package main

import "fmt"

func fibo(i int) (result int) {
	if i == 1 || i == 2 {
		result = 1
	} else {
		result = fibo(i-1) + fibo(i-2)
	}
	return
}

func main() {

	for i := 1; i < 50; i++ {
		result := fibo(i)
		fmt.Printf("result : %d\n", result)
	}
}
