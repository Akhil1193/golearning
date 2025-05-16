package basics

import "fmt"

func panicc() {
	process(3)
	process(-1)

}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("Input must not be a negative number")
		// fmt.Println("after panic")
	}
	fmt.Println("Input is:", input)
}
