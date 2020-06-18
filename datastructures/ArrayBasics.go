package datastructures

import (
	"fmt"
)

//executes function related to arrays and slices
func ExecuteArrayExamples() {
	// arrayBasics()
	// inferArrayLength()
	// demoValueTypeProperty()

	//slices example
	initializeSlice()
	testSlice()
}

func demoValueTypeProperty() {
	names := [3]string{"ne", "tyt", "uyuy"}
	modifyArray(names)
	fmt.Println("names = ", names)
	valueType()
}

func inferArrayLength() {
	var lengthInferredArrays = [...]int{2, 4, 5, 6, 8}
	fmt.Println(len(lengthInferredArrays), " is the length")
}

func arrayBasics() {
	var integers [2]int
	integers[0] = 19
	fmt.Println(len(integers))

	var testWithLength = [3]int{1, 2, 3}
	testWithLenArr := testWithLength
	fmt.Println(testWithLength)
	testWithLenArr[0] = 10
	fmt.Println(testWithLength)
}

/*
when you assign an array to a new variable or pass an array to a function,
the entire array is copied. So if you make any changes to this copied array,
the original array won’t be affected and will remain unchanged.
*/
func modifyArray(arr [3]string) {
	arr[0] = "changed"
}

func valueType() {
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`
	a2[1] = "German"
	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)
}
