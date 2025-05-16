package basics

import "fmt"

func variadicFunctions() {
	// ... ellipsis
	s1, t1 := add(1, 5, 6, 7, 8)
	fmt.Println(s1, t1)

	numbers := []int{1, 2, 3, 4, 5, 6, 6}
	s2, t2 := add(2, numbers...)
	fmt.Println(s2, t2)
}

func add(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
