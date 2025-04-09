package basics

import "fmt"

type EmployeeInfo struct {
	FirstName string
	LastName  string
	Age       int
}

func namingConvention() {

	// PascalCase - structs,interfaces, enums
	// e.g: CalculateArea, UserInfo

	//snake_case - e.g: used_id, first_name

	// UPPERCASE - usecase in CONSTANTS
	const MAXTRIES = 5
	// mixedCase- e.g: javaScript, isValid etc
	var employeeId int
	fmt.Println(employeeId)
}
