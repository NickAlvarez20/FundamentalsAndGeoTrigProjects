package main

import "fmt"

// Define the structs for Store and Transaction
type Store struct {
	Name      string
	Inventory map[string]int // maps from item names to quantity
}

type Transaction struct {
	ItemName string
	Quantity int // quantity can be positive or negative
}

// Define function to UpdateInventory
func UpdateInventory(s *Store, t Transaction) bool {
	itemName := t.ItemName
	quantity := t.Quantity // Assignment to variables for easier to read code

	// Store currentQuantity in store and check if item exists
	currentQuantity, exists := s.Inventory[itemName]

	// If item does not exist and quantity is less than 0, return false
	if !exists && quantity < 0 {
		return false
	}

	// Calculate the updatedQuantity
	updatedQuantity := currentQuantity + quantity

	// If the updatedQuantity is negative, return false
	if updatedQuantity < 0 {
		return false
	}

	// update the store with the transaction, return true
	s.Inventory[itemName] = updatedQuantity
	return true

}

func main() {
	// Test Cases
	s := &Store{
		Name:      "Store 6",
		Inventory: map[string]int{},
	}
	t := Transaction{
		ItemName: "Item 1",
		Quantity: 5,
	}

	// Print results

	fmt.Println(UpdateInventory(s, t), s)
}
