package basics

import "fmt"

// global variable declaration
var firstName string = "Lingareddy"

func variables(){
	var name string = "akhil"
	var age = 29

	height := 5.2
	weight := 91

	fmt.Println(name,age,height,weight,firstName)

	firstName = "reddy"
	fmt.Println(firstName)

	// datatype default values are 
	// int = 0
	// float = 0
	// string = ""
	// bool = false
	// struct, map, pointer, array, function types = nil
}