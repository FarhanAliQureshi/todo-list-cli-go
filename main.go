package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userSelection int

	displayMainMenu()
	userSelection = getUserMenuChoice(1, 5)
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

// Prompt user for input, validate user input, and return user's choice
func getUserMenuChoice(validFrom int, validTo int) int {
	var choice int
	var userInput string
	var err error

	for {
		fmt.Printf("Your Choice? [%d-%d]: ", validFrom, validTo)
		// Get user input and then try to convert it into an integer
		// Getting integer directly from user using Scanln() doesn't
		// work as expected.
		fmt.Scanln(&userInput)
		// Variable "err" should be declared before this statement. Using ":="
		// will locally declare variable "choice" for this for-loop scope. We won't
		// be able to access original variable "choice". AKA variable shadowing.
		choice, err = strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Invalid input")
		} else if choice < validFrom || choice > validTo {
			fmt.Println("Invalid choice")
		} else {
			// By this point, user's choice must be between >= validFrom and <= validTo
			// Therefore, it is a valid input, and let's return it to the calling func
			fmt.Println("Choice: ", choice)
			break
		}
	}

	return choice
}
