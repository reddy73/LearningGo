package basics

import (
	"fmt"
	"runtime"
)

//NoOfCores prints available cores
func NoOfCores() {
	fmt.Println(runtime.NumCPU())
}
