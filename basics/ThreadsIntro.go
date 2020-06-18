package basics

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := n; i < 100; i = i + 2 {
		fmt.Println(n, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func StartThreads() {
	go f(1)
	go f(2)
}
