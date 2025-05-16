package basics

import "fmt"

func ifElase() {
	// age := 25
	// var input int
	// fmt.Println("Enter Age:")
	// fmt.Scan(&input)
	// if input > age {
	// 	fmt.Println("Age is greater than 25")
	// } else if input < age {
	// 	fmt.Println("age is less than 25")
	// } else {
	// 	fmt.Println("Correct age")
	// }

	//nested if

	number := 53
	if number%2 == 0 {
		if number%3 == 0 {
			fmt.Println("number is divisible by 2 and 3")
		} else {
			fmt.Println("number is divisble by 2")
		}
	} else {
		fmt.Println("number is not divisible by 2 and 3")
	}
}
