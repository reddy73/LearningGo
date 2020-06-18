package datastructures

import (
	"fmt"
	"reflect"
)

func initializeSlice() {
	var mySlice = []int{1, 2, 3} // using literal
	fmt.Println(len(mySlice), " original length of slice")
	mySlice = append(mySlice, 5)
	fmt.Println(len(mySlice), " length of slice after append operation", cap(mySlice))
	fmt.Println(mySlice)

	var anotherSlice []int
	fmt.Println(copy(anotherSlice, mySlice), " bytes copied")
}

func testSlice() {
	var test = []int{1, 2, 3}
	testArr := test
	fmt.Println(test)
	testArr[0] = 10
	fmt.Println(test, reflect.TypeOf(test).Kind())

	fmt.Println("length of slice", len(test))

	fmt.Println("capacity of lengths", cap(test))

	test = append(test, 5, 6, 7)
	fmt.Println(len(test))

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i], " element from slice")
	}
}
