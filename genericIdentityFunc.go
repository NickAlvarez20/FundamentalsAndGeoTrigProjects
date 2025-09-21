package main

import (
	"fmt"
	"reflect"
)

// Identity returns the same value with the same type
func Identity[T any](value T) T {
	return value
}

func main() {
	word := "Hello"
	number := 123
	float := 123.123
	fmt.Println(Identity(word))
	fmt.Println("Type:", reflect.TypeOf(word))

	fmt.Println(Identity(number))
	fmt.Println("Type:", reflect.TypeOf(number))

	fmt.Println(Identity(float))
	fmt.Println("Type:", reflect.TypeOf(float))
}
