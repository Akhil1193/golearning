package basics

import "fmt"

func forLoopAsWhile() {
	//while loop

	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Print numbers from:", i)
	// 	i++
	// }

	//while loop with break
	//  i:= 1
	// for {
	// 	fmt.Println("Odd numbers:",i)
	// 	i+=2
	// 	if i == 13{
	// 		break
	// 	}
	// }

	//while loop with continue

	number := 0
	for number <= 50 {
		number += 10
		if number == 20 {
			continue
		}
		if number == 50 {
			break
		}
		fmt.Println("numbers:", number)
	}

}
