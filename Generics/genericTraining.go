package main

import (
	"fmt"
	"reflect"
)

func Identity[T any](a T) T {
	return a
}

func AddGeneric[I ~int | ~float64 | ~complex128](a, b I) I {
	return a + b
}

func addFloatingString[S ~float64 | ~string] (a, b S) S{
	return a+b
}

func main() {
	// Test Identity Function
	variableUnknown := 4
	fmt.Println(Identity(reflect.TypeOf(variableUnknown)))

	// Test AddGeneric Function
	variablesUnknown := 10.5780235465413
	variablesUnknown1 := 99.9256513
	fmt.Println(AddGeneric(variablesUnknown, variablesUnknown1)) // float64

	varUnknown2 := 90
	varUnknown3 := 111
	fmt.Println(AddGeneric(varUnknown2, varUnknown3)) // int

	imaginaries1 := 15 + 5i
	imaginaries2 := 25 + 99i
	fmt.Println(AddGeneric(imaginaries1, imaginaries2)) // complex128

	float1 := 3.14
	float2 := 6.28
	str1 := "Hello"
	str2 := "World"
	fmt.Println(addFloatingString(float1, float2), addFloatingString(str1, str2)) 
}
