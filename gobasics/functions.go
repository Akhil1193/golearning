package basics

import "fmt"

func functions() {

	addition := addNumbers(5, 6)
	fmt.Println(addition)

	 greet := func(){
		fmt.Println("Hello anynonums functions!")
	}

	greet()

	add := applyFunction(5, 3, addNumbers)
	fmt.Println(add)

}

func addNumbers(x, y int) int {
	return x + y
}

func applyFunction(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}
