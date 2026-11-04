package main

import ("fmt" , "strconv")

func main() {
	// fmt.Println("Hello, I am Rashi Yadav")

	// name := "Rashi Yadav"
	// roll := 123456789

	// fmt.Printf("NAME : %s\n", name)
	// fmt.Printf("ROLL NO : %d\n", roll)

	// fmt.Println("NAME : %s\n", name)
	// fmt.Println("ROLL NO : %d\n", roll)


}








// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	// var a int = 87
// 	// b := strconv.Itoa(a)
// 	// fmt.Println(b)

// 	for i, ch := range "go!" {
// 		if i == 0 {
// 			continue
// 		}
// 		fmt.Println(i,ch)
// 	}
// }









package main

import "fmt"

// func sum(nums ...int) int {
//     total := 0

//     for _, n := range nums {
//         total += n
//     }

//     return total
// }

// func main() {
//     result := sum(1, 2, 3, 4, 5)
//     fmt.Println(result)
// }/















package main

import (
	"errors"
	"fmt"
)

// Item structure
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory slice
var inventory []Item

// Add a new item
func AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("ID must be positive")
	}
	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}
	if item.Quantity < 0 || item.Price < 0 {
		return errors.New("quantity and price cannot be negative")
	}

	for _, existing := range inventory {
		if existing.ID == item.ID {
			return errors.New("item ID already exists")
		}
	}

	inventory = append(inventory, item)
	return nil
}

// Update an existing item
func UpdateItem(id int, quantity int, price float64) error {
	if quantity < 0 || price < 0 {
		return errors.New("quantity and price cannot be negative")
	}

	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Quantity = quantity
			inventory[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

// Remove an item
func RemoveItem(id int) error {
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

// Summary statistics
type Summary struct {
	ItemCount int
	TotalValue float64
}

// Compute statistics concurrently using a goroutine and channel
func CalculateSummary(items []Item) Summary {
	result := make(chan Summary)

	go func() {
		var summary Summary

		summary.ItemCount = len(items)

		for _, item := range items {
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		result <- summary
	}()

	return <-result
}

func main() {
	// Add items
	err := AddItem(Item{1, "Laptop", 5, 50000})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(Item{2, "Mouse", 10, 500})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(Item{3, "Keyboard", 4, 1200})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Display inventory
	fmt.Println("Inventory:")
	for _, item := range inventory {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}

	// Update item
	if err := UpdateItem(2, 15, 500); err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Println("\nMouse updated successfully")
	}

	// Remove item
	if err := RemoveItem(3); err != nil {
		fmt.Println("Remove error:", err)
	} else {
		fmt.Println("Keyboard removed successfully")
	}

	// Concurrent summary
	summary := CalculateSummary(inventory)

	fmt.Println("\nInventory Summary:")
	fmt.Println("Item Count:", summary.ItemCount)
	fmt.Printf("Total Inventory Value: Rs. %.2f\n",
		summary.TotalValue)
}










Package main

import "testing"

// Test 1: Add an item
func TestAddItem(t *testing.T) {
	inventory = nil

	err := AddItem(Item{1, "Laptop", 5, 50000})
	if err != nil {
		t.Fatalf("Expected item to be added, got error: %v", err)
	}

	if len(inventory) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(inventory))
	}
}

// Test 2: Update an item
func TestUpdateItem(t *testing.T) {
	inventory = []Item{
		{1, "Laptop", 5, 50000},
	}

	err := UpdateItem(1, 10, 45000)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if inventory[0].Quantity != 10 {
		t.Errorf("Expected quantity 10, got %d",
			inventory[0].Quantity)
	}

	if inventory[0].Price != 45000 {
		t.Errorf("Expected price 45000, got %.2f",
			inventory[0].Price)
	}
}

// Test 3: Remove an item
func TestRemoveItem(t *testing.T) {
	inventory = []Item{
		{1, "Laptop", 5, 50000},
		{2, "Mouse", 10, 500},
	}

	err := RemoveItem(1)
	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	if len(inventory) != 1 {
		t.Fatalf("Expected 1 item remaining, got %d",
			len(inventory))
	}

	if inventory[0].ID != 2 {
		t.Errorf("Expected item ID 2, got %d",
			inventory[0].ID)
	}
}

// Test 4: Reject invalid input
func TestAddInvalidItem(t *testing.T) {
	inventory = nil

	err := AddItem(Item{0, "Laptop", 5, 50000})
	if err == nil {
		t.Fatal("Expected error for invalid ID")
	}

	if len(inventory) != 0 {
		t.Fatal("Invalid item should not be added")
	}
}

