package basics

import "fmt"

func array() {
	// names := [5]int{5, 6, 7, 8, 9}
	// fmt.Println(names)
	// names[2] = 10
	// fmt.Println(names)

	// fruits := [4]string{"apple", "banana", "mango", "pineapple"}
	// fmt.Println(fruits)

	// originalArray := [4]int{1, 2, 3, 4}
	// copiedArray := originalArray
	// copiedArray[0] = 100
	// fmt.Println(originalArray)
	// fmt.Println(copiedArray)

	//underscore is balnk identifier
	// for _, v := range copiedArray {
	// 	fmt.Printf("Value:%d\n", v)
	// }

	//compare arrays
	// array1 := [5]int{1, 2, 3, 4, 5}
	// array2 := [5]int{3, 4, 5, 6, 7}
	// fmt.Println(array1 == array2)

	//nested arrays
	// matrix := [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(matrix)

	number:= [3]int{1,2,3}
	var number1  *[3]int
	number1 = &number
	number1[1]=5
	fmt.Println(number,*number1)

}
