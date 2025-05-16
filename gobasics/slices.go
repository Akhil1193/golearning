package basics

import (
	"fmt"
	"slices"
)

func slice() {

	num := []int{1, 2, 3, 4, 5}
	num1 := make([]int, 5)
	num1 = []int{1, 2, 3, 4, 5}

	fmt.Println(num, num1)

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println(array, slice)

	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	fmt.Println("print slice1 values:", slice1)

	num3 := append(num, num1...)
	fmt.Println(num3)

	for i, v := range slice1 {
		fmt.Printf("for index %v Value is %v\n", i, v)
	}

	if slices.Equal(num, num1) {
		fmt.Println("num is equal to num1")
	}

}
