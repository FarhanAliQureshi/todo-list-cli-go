package main

import (
	"fmt"
)

func main() {
	var userSelection int

	displayMainMenu()
	userSelection = GetUserMenuChoice(1, 5)
	fmt.Printf("User selected: %d\n", userSelection)
}

// Displays main menu. Input of choice selection should be handled separately.
func displayMainMenu() {
	fmt.Println("Todo List CLI in Golang")
	fmt.Println("")
	fmt.Println("1. List titles of all Todo Lists")
	fmt.Println("2. Display all items of a Todo List")
	fmt.Println("2. Create a new Todo List ")
	fmt.Println("3. Add items to a Todo List")
	fmt.Println("4. Delete a Todo List")
	fmt.Println("5. Exit")
	fmt.Println("")
}
