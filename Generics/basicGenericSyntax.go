package main

import (
	"fmt"
)

// Swap func takes two arguments of the same type T
// return them in reverse order
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	int1, int2 := Swap(1, 2)
	fmt.Println(int1, int2)

	str1, str2 := Swap("hello", "generics")
	fmt.Println(str1, str2)
}
