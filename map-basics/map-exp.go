package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"

	fmt.Println(m[0])
	fmt.Println(len(m))
	m1 := make(map[string]int)

	fmt.Println(m1)

	fmt.Println(m[0])
	modifyMap(m)
	fmt.Println(m[0])
}

func modifyMap(m map[int]string) {
	m[0] = "Modified Zero"
}
