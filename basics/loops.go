package basics

import (
	"fmt"
	"time"
)

func dontTryInProdInfiniteLoop() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Please don't do this in prod")
	}
}

func forAsWhile() {
	number := 19
	fmt.Println("For as a while loop")
	for number%2 != 0 && number > 0 {
		fmt.Println(number)
		number = number - 2
	}
}

func iterateForEach() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println("\nfor each from a collection")
	for index, el := range arr {
		fmt.Println(el, index)
	}
}

func iterateNormal() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
	}
}
