package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nexInt := intSeq()

	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	// fmt.Println(intSeq()())
	// fmt.Println(intSeq()())
	// fmt.Println(intSeq()())

	newInts := intSeq()
	fmt.Println(newInts())
}
