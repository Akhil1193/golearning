package basics

import "fmt"

func recoverr() {
	p()
}

func p() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic", r)
		}
	}()
	panic("Something went wrong")
	fmt.Println("After Panic")
}
