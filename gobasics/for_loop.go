package basics

import "fmt"

func forLoop() {
	//simple for loop iteration
	// for i := 1; i < 6; i++ {
	// 	fmt.Println(i)
	// }

	//for loop iteration over a slice
	// slice := []int{1,2,3,4,5,6,7}
	// for index,value:= range slice{
	// 	fmt.Printf("Index:%d, Value:%d\n", index, value)
	// }

	//break and continue
	// for i := 1; i <= 10; i = i + 2 {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	if i == 7 {
	// 		break
	// 	}
	// }

	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// for i:= range 10{
	// 	i++
	// 	fmt.Println(i)
	// }
}
