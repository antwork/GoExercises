package main 

import "fmt"

func main () {

	sliceA := make([]int, 5)
	fmt.Println("SliceA :%d len:%d cap %d", sliceA, len(sliceA), cap(sliceA))

	sliceB := make([]int, 0, 5)
	fmt.Println("SliceB :%d len:%d cap %d", sliceB, len(sliceB), cap(sliceB))

	s0 := []int{0, 0}
	fmt.Println("s0: %d", s0)

	s01 := append(s0, 2)
	fmt.Println("s01 %d", s01)
	fmt.Println("s0: %d", s0)

	s02 := append(s01, 3, 5, 7)
	fmt.Println("s02 %d", s02)

	s03 := append(s02, s01...)
	fmt.Println("s03 %d", s03)

	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[2:4]
	fmt.Println("s1: %d" ,s1)

	s2 := a[1:5]
	fmt.Println("s2: %d" , s2)

	s3 := a[:]
	fmt.Println("s3: %d" , s3)

	s4 := a[:4]
	fmt.Println("s4: %d" , s4)

	s5 := s2[:]
	fmt.Println("s5: %d" , s5)
}