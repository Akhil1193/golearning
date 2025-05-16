package basics

import (
	"fmt"
	"os"
)

func exxit() {

	defer fmt.Println("deferred statement")

	fmt.Println("Statement before exit")
	//status code 1
	os.Exit(1)

	fmt.Println("Statement after exit")
}
