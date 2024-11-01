package main

import (
	"fmt"
)

const TODOLIST_FILENAME string = "todolist.txt"

func main() {
	var userSelection int

	// Read todo lists file and load them into memory
	todos := loadTodoListsFromFile(TODOLIST_FILENAME)

	for {
		displayMainMenu()
		userSelection = GetUserMenuChoice(0, 4)
		//fmt.Printf("User selected: %d\n", userSelection)
		if userSelection == 1 {
			displayTitlesOfTodoLists(todos)
			fmt.Println("")
		} else if userSelection == 2 {
			displayTodoListItems(todos)
		} else if userSelection == 0 {
			// Break the loop and gracefully end the program
			break
		}
	}
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
	fmt.Println("0. Exit")
	fmt.Println("")
}

func displayTitlesOfTodoLists(todos []todo) {
	fmt.Printf("\nTitles of all Todo Lists:\n\n")
	for index, item := range todos {
		fmt.Printf("%d. %s\n", index+1, item.name)
	}
}

func loadTodoListsFromFile(filename string) []todo {
	data, err := ReadDataFile(filename)
	if err != nil {
		panic(err)
	}

	todos := parseStringAndCreateTodoLists(data)
	return todos
}

func displayTodoListItems(todos []todo) {
	fmt.Println("")
	fmt.Println("Select a Todo List to list all its items")
	displayTitlesOfTodoLists(todos)
	fmt.Println("0. Main Menu")
	fmt.Println("")

	userSelection := GetUserMenuChoice(0, len(todos))
	if userSelection != 0 {
		// Index of list starts at zero, while the user input starts with 1
		// Therefore, we must first deduct 1 from userSelection to get
		// a proper index
		selectedIndex := userSelection - 1

		fmt.Println("")
		fmt.Printf("Selected Todo List: %s\n", todos[selectedIndex].name)
		for index, item := range todos[selectedIndex].items {
			var completed string
			if item.completed {
				completed = "+"
			} else {
				completed = "-"
			}
			fmt.Printf("%d. %s %s\n", index+1, completed, item.name)
		}
		fmt.Println("")
		PauseProgramToLetUserRead()
	}
}
