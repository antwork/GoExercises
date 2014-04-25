package qsort

import (
	"fmt"
)

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}

		if j >= p {
			values[p] = values[j]
			fmt.Println("i:", i, "j:", j, "p:", p, "j->p:", values)
			p = j
		}

		if values[i] < temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			fmt.Println("i:", i, "j:", j, "p:", p, "i->p:", values)
			p = i
		}
	}
	values[p] = temp
	fmt.Println("i:", i, "j:", j, "p:", p, "p->_:", values)
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
