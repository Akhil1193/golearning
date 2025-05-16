package basics

import "fmt"

func deferred() {

	process(10)

}

func proces(i int) {
	defer fmt.Println("Deferred i statement:",i)
	defer fmt.Println("1st Deferred statement executed")
	defer fmt.Println("2nd Deferred statement executed")
	defer fmt.Println("3rd Deferred statement executed")
	fmt.Println("Noraml statement")
	i++
	fmt.Println("i value:", i)
}
