package main

import "fmt"

func bubble(datas []int) (results []int) {

	results = datas[:]
	count := len(results)

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			v := results[i]
			vTemp := results[j]
			if vTemp < v {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
	return
}

func main() {
	array := []int{5, 3, 2, 6, 7, 8, 10, 1}
	var result []int
	result = bubble(array)
	fmt.Printf("result %d", result)
}
