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
		userSelection = GetUserMenuChoice(1, 5)
		//fmt.Printf("User selected: %d\n", userSelection)
		if userSelection == 1 {
			displayTitlesOfTodoLists(todos)
		} else if userSelection == 5 {
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
	fmt.Println("5. Exit")
	fmt.Println("")
}

func displayTitlesOfTodoLists(todos []todo) {
	fmt.Printf("\nTitles of all Todo Lists\n")
	for index, item := range todos {
		fmt.Printf("%d. %s\n", index+1, item.name)
	}
	fmt.Println("")
}

func loadTodoListsFromFile(filename string) []todo {
	data, err := ReadDataFile(filename)
	if err != nil {
		panic(err)
	}

	todos := parseStringAndCreateTodoLists(data)
	return todos
}
