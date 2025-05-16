package basics

import "fmt"

func switchCase() {

	// fruit := "orange"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an Apple.")
	// case "banana":
	// 	fmt.Println("It's a Banana.")
	// default:
	// 	fmt.Println("Unknown Fruit!")
	// }

	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a Weekday.")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// num := 29

	// switch {
	// case num == 10:
	// 	fmt.Println("Number is 10.")
	// case num >= 10 && num < 30:
	// 	fmt.Println("Number is in between 10 and 30.")
	// default:
	// 	fmt.Println("No number found!")
	// }

	//To check two conditions
	// num1 := 2
	// switch {
	// case num1 > 1:
	// 	fmt.Println("Number is 1")
	// 	fallthrough
	// case num1 <= 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Number not found")
	// }

	checkType(32)
	checkType(6.5)
	checkType("akhil")
	checkType(true)
}

func checkType(x interface{}){
	switch x. (type){
	case  int:
		fmt.Println("X is of type int")
	case float64:
		fmt.Println("x is of type float")
	case string:
		fmt.Println("x is of type string")
	default:
		fmt.Println("unknown type")
	}
}
