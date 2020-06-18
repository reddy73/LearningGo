package basics

import "fmt"

func ScopeBasics() {
	scopeBasics()
}

func ControlStructuresExamples(s string) {
	switch s {
	case "tryIf":
		tryIf()
	case "tryIfM":
		tryIfWithMultipleTemporaryAssignments()
	case "trySwitch":
		trySwitch()
	case "tryFallthrough":
		trySwitchFallThrough()
	default:
		fmt.Println("Do Nothing")
	}
}

func RunLoops(s string) {
	switch s {
	case "for":
		iterateNormal()
	case "while":
		forAsWhile()
	case "for-each":
		iterateForEach()
	case "infinite":
		dontTryInProdInfiniteLoop()
	default:
		fmt.Println("Not executing any loops")
	}
}
