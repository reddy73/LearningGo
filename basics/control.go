package basics

import (
	"fmt"
)

func controlBasics() {

	fmt.Println("\n\nFrom Conditions.go")
	tryIf()

	tryIfWithMultipleTemporaryAssignments()

	trySwitch()

	trySwitchFallThrough()

	fmt.Println("Try main again now")
}

func trySwitchFallThrough() {
	switch "test" {
	case "test":
		fmt.Println("It matched") //prints this
		fallthrough
	case "one":
		fmt.Println("fall through at one") //prints this too
		fallthrough
	case "two":
		fmt.Println("fall through at two") //prints this too
	case "three":
		fmt.Println("no fall through at three") //deosnt print
	default:
		fmt.Println("It's default") //doesn't print
	}
}

func trySwitch() {
	i := 10
	j := 11
	switch j {
	case i + 1:
		fmt.Println("It matched with 11") //prints this
	case i:
		fmt.Println("It matched with 10")
	case i - 1:
		fmt.Println("It matched with 9")
	default:
		fmt.Println("Defaulting")
	}
}

func tryIf() {
	if 4%2 == 0 {
		fmt.Println("Yes its divisible") //prints this statement
	}
}

func tryIfWithMultipleTemporaryAssignments() {
	if i, j := 10, 19; i < j {
		fmt.Println("It works too")
	}
	//fmt.Print(i, j) //i & j are not accessible from here
}
