package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Sales  int
}

func topAuthors(books []Book) map[string]int {
	result := make(map[string]int)
	for _, book := range books {
		result[book.Author] += book.Sales
	}
	topAuthors := make(map[string]int)
	for author, sales := range result {
		if sales >= 10000 {
			topAuthors[author] = sales
		}
	}
	return topAuthors
}

func main() {
	books := []Book{
		Book{
			Title:  "Book 1",
			Author: "Author 1",
			Sales:  3000,
		},
		Book{
			Title:  "Book 2",
			Author: "Author 3",
			Sales:  4000,
		},
		Book{
			Title:  "Book 3",
			Author: "Author 1",
			Sales:  2000,
		},
		Book{
			Title:  "Book 4",
			Author: "Author 3",
			Sales:  7000,
		},
	}

	fmt.Println(topAuthors(books))
}
