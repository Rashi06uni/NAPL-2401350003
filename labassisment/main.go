
package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

func AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("ID must be positive")
	}
	if item.Name == "" {
		return errors.New("item name must not be empty")
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

func RemoveItem(id int) error {
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func main() {
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

	fmt.Println("Inventory:")
	for _, item := range inventory {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	if err := UpdateItem(2, 15, 550); err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Println("\nMice updated successfully")
	}

	if err := RemoveItem(3); err != nil {
		fmt.Println("Remove error:", err)
	} else {
		fmt.Println("Keyboard removed successfully")
	}

	fmt.Println("\nFinal Inventory:")
	for _, item := range inventory {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}
}
