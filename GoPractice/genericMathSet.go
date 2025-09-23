package main

import (
	"fmt"
)

// Type definition
type Set[E comparable] struct {
	elements map[E]struct{}
}

// Constructor function
func NewSet[E comparable]() *Set[E] {
	return &Set[E]{
		elements: make(map[E]struct{}),
	}
}

// Functions:
// Add method: Takes some value and stores it in the set
func (s *Set[E]) Add(v E) {
	s.elements[v] = struct{}{}
}

// Remove method: Removes an element from the set, no effect if it doesn't exist
func (s *Set[E]) Remove(v E) {
	delete(s.elements, v)
}

// Contains method: Returns true if specified value exists in set, false otherwise
func (s *Set[E]) Contains(v E) bool {
	_, found := s.elements[v]
	return found
}

// Size method: Return the number of elements within the set

func (s *Set[E]) Size() int {
	return len(s.elements)
}

func main() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println(s.Contains(3)) // true
	fmt.Println(s.Size()) // 3
	s.Add(3)
	s.Remove(2)
	fmt.Println(s.Contains(1)) // true
	fmt.Println(s.Contains(2)) // false
	fmt.Println(s.Size())      // 2

}
