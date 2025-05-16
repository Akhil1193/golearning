package basics

import (
	"errors"
	"fmt"
)

func multipleReturnValues() {

	q, r := divide(12, 3)
	fmt.Println(q, r)

	c, err := compare(5, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(c)
	}

}

func divide(x, y int) (quotient int,reminder int) {
	quotient = x / y
	reminder = x % y
	return quotient, reminder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("both are equal")
	}
}
