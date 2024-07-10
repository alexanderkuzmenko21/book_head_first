package main

import (
	"fmt"
	"reflect"
)

func functionA(a int, b int) {
	fmt.Println(a + b)
}

func functionB(a int, b int) {
	fmt.Println(a * b)
	fmt.Println(reflect.TypeOf(a))
}

func functionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Println(a)
	}
	// fmt.Println()
}

func main() {
	functionA(2, 3)
	functionB(2, 3)
	functionC(true)
	functionD("$", 4)
	functionA(5, 6)
	functionB(5, 6)
	functionC(false)
	functionD("ha", 3)
}
