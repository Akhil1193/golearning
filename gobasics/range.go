package basics

import "fmt"

func rg() {
	name := "akhil"

	for i, v := range name {
		fmt.Println(i, v)
		fmt.Printf("index: %d, character: %c\n",i,v)
	}

}
