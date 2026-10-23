
package main

import "testing"


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
