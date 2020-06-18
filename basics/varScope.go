package basics

import (
	"fmt"
)

var testNo = 25 // Global scope
func scopeBasics() {
	fmt.Println(testNo)
	if i, j := 9, 10; j > i {
		var test = 50 //block scope
		fmt.Println("Test number is: ", test, ".")
	}
	// fmt.Println(test) // test is not accessible, thros error
	testFunctionScope()
}

func testFunctionScope() {
	fmt.Println(testNo) // prints 25
}
