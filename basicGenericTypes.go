package main

import(
	"fmt"
)

// Create a Box[T any] type that holds one value of type T
type Box[T any] struct{
	value T
}

// Add a Value() T method to retrieve the value
func (b Box[T]) Value() T{
	return b.value
}

func NewBox[T any](value T) Box[T]{
	return Box[T]{value: value}
}

// -----------------------------------------------

// Create a Block[C any] type that holds value C
type Block[C any] struct{
	value C
}

// Add a Value() C method to retrieve the value
func (c Block[C]) Value() C{
	return c.value
}

func newBlock[C any](value C) Block[C]{
	return Block[C]{value: value}
}

func main(){
	intBox := NewBox(42)
	fmt.Println(intBox.value)

	strBox := NewBox("String")
	fmt.Println(strBox.value)

	float64Box := NewBox(42.02342354213)
	fmt.Println(float64Box.value)

	strBlock := newBlock("Generics")
	fmt.Println(strBlock.value)

	runeBlock := newBlock(rune('A'))
	fmt.Println(runeBlock)
}
