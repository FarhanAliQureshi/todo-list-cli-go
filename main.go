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
		userSelection = GetUserMenuChoice(0, 6)
		if userSelection == 1 { // Menu: List titles of all Todo Lists
			displayTitlesOfTodoLists(todos, true)
		} else if userSelection == 2 { // Menu: Display all items of a Todo List
			displayTodoListItems(todos, true)
		} else if userSelection == 3 { // Menu: Create a new Todo List
			todos = createNewTodoList(todos)
		} else if userSelection == 4 { // Menu: Manage a Todo List
			todos = manageTodoList(todos)
		} else if userSelection == 5 { // Menu: Delete a Todo List
			todos = deleteTodoList(todos)
		} else if userSelection == 6 { // Menu: Save changes to file on disk
			saveTodoListToFile(todos)
		} else if userSelection == 0 { // Menu: Exit
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
	fmt.Println("3. Create a new Todo List")
	fmt.Println("4. Manage a Todo List")
	fmt.Println("5. Delete a Todo List")
	fmt.Println("6. Save changes to file on disk")
	fmt.Println("0. Exit")
	fmt.Println("")
}

func displayTitlesOfTodoLists(todos []todo, pause bool) {
	fmt.Printf("\nTitles of all Todo Lists:\n\n")
	for index, item := range todos {
		fmt.Printf("%d. %s\n", index+1, item.name)
	}
	if pause {
		fmt.Println("")
		PauseProgramToLetUserRead()
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

func displayTodoListItems(todos []todo, pause bool) {
	fmt.Println("")
	fmt.Println("Select a Todo List to list all its items")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Main Menu")
	fmt.Println("")

	userSelection := GetUserMenuChoice(0, len(todos))
	if userSelection != 0 {
		// Index of list starts at zero, while the user input starts with 1
		// Therefore, we must first deduct 1 from userSelection to get
		// a proper index
		selectedListIndex := userSelection - 1

		fmt.Println("")
		fmt.Printf("Selected Todo List: %s\n", todos[selectedListIndex].name)
		displayTodoListItemsOfSelectedList(todos, selectedListIndex)
		if pause {
			fmt.Println("")
			PauseProgramToLetUserRead()
		}
	}
}

func createNewTodoList(todos []todo) []todo {
	fmt.Println("")
	fmt.Println("Create a New Todo List")
	fmt.Println("")
	fmt.Println("(Press Enter key without writing anything to cancel)")
	newTodoListName := GetInputFromUser("Enter Name for New Todo List: ")
	if newTodoListName != "" {
		todoList := todo{}
		todoList.name = newTodoListName
		todos = append(todos, todoList)
	}

	return todos
}

func deleteTodoList(todos []todo) []todo {
	fmt.Println("")
	fmt.Println("Delete a Todo List")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Cancel")
	fmt.Println("")
	fmt.Println("Select a Todo List to delete")
	userSelection := GetUserMenuChoice(0, len(todos))

	if userSelection > 0 {
		index := userSelection - 1
		fmt.Printf("Deleting \"%s\" ...\n", todos[index].name)
		return append(todos[:index], todos[index+1:]...)
	}

	return todos
}

func manageTodoList(todos []todo) []todo {
	fmt.Println("")
	fmt.Println("Select a Todo List to manage")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Main Menu")
	fmt.Println("")

	userSelection := GetUserMenuChoice(0, len(todos))
	if userSelection != 0 {
		selectedListIndex := userSelection - 1
		for {
			displayManageTodoListMenu(todos, selectedListIndex)
			userSelection = GetUserMenuChoice(0, 5)
			if userSelection == 1 { // Menu: Display items of the todo list
				displayTodoListItemsOfSelectedList(todos, selectedListIndex)
			} else if userSelection == 2 { // Menu: Mark an item as completed
				todos = changeStatusOfTodoListItem(todos, selectedListIndex, true)
			} else if userSelection == 3 { // Menu: Mark an item as incomplete
				todos = changeStatusOfTodoListItem(todos, selectedListIndex, false)
			} else if userSelection == 4 { // Menu: Add an item to todo list
				todos = addNewTodoListItem(todos, selectedListIndex)
			} else if userSelection == 5 { // Menu: Delete an item from todo list
				todos = deleteTodoListItem(todos, selectedListIndex)
			} else if userSelection == 0 { // Menu: Back to Main Menu
				break
			}
		}
	}

	return todos
}

func displayManageTodoListMenu(todos []todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Printf("Managing Todo List: %s\n", todos[selectedListIndex].name)
	fmt.Println("")
	fmt.Println("1. Display items of the todo list")
	fmt.Println("2. Mark an item as completed")
	fmt.Println("3. Mark an item as incomplete")
	fmt.Println("4. Add an item to todo list")
	fmt.Println("5. Delete an item from todo list")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("")
}

func displayTodoListItemsOfSelectedList(todos []todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Printf("Todo List Items for: %s\n", todos[selectedListIndex].name)
	for index, item := range todos[selectedListIndex].items {
		var completed string
		if item.completed {
			completed = "+"
		} else {
			completed = "-"
		}
		fmt.Printf("%d. %s %s\n", index+1, completed, item.name)
	}
}

func changeStatusOfTodoListItem(todos []todo, selectedListIndex int, markCompleted bool) []todo {
	var changeStatus string
	if markCompleted {
		changeStatus = "completed"
	} else {
		changeStatus = "incomplete"
	}
	fmt.Println("")
	fmt.Printf("Select an item to mark as %s\n", changeStatus)
	displayTodoListItemsOfSelectedList(todos, selectedListIndex)
	fmt.Println("0. Cancel")
	fmt.Println("")
	userSelection := GetUserMenuChoice(0, len(todos[selectedListIndex].items))

	if userSelection != 0 {
		selectedItem := userSelection - 1
		todos[selectedListIndex].items[selectedItem].completed = markCompleted
	}
	return todos
}

func addNewTodoListItem(todos []todo, selectedListIndex int) []todo {
	fmt.Println("")
	fmt.Println("Add a New Item")
	fmt.Println("")
	fmt.Println("(Press Enter key without writing anything to cancel)")
	newItemName := GetInputFromUser("Enter Item Name: ")
	if newItemName != "" {
		newItem := item{}
		newItem.name = newItemName
		todos[selectedListIndex].items = append(todos[selectedListIndex].items, newItem)
	}

	return todos
}

func deleteTodoListItem(todos []todo, selectedListIndex int) []todo {
	fmt.Println("")
	fmt.Println("Delete a Todo List's Item")
	displayTodoListItemsOfSelectedList(todos, selectedListIndex)
	fmt.Println("0. Cancel")
	fmt.Println("")
	fmt.Println("Select a Todo List's Item to delete")
	userSelection := GetUserMenuChoice(0, len(todos[selectedListIndex].items))

	if userSelection > 0 {
		index := userSelection - 1
		fmt.Printf("Deleting \"%s\" ...\n", todos[selectedListIndex].items[index].name)
		todos[selectedListIndex].items = append(todos[selectedListIndex].items[:index], todos[selectedListIndex].items[index+1:]...)
	}

	return todos
}

func saveTodoListToFile(todos []todo) {
	var buffer, status string

	for _, todoList := range todos {
		buffer = fmt.Sprintln(buffer + "# " + todoList.name)
		for _, item := range todoList.items {
			if item.completed {
				status = "+"
			} else {
				status = "-"
			}
			buffer = fmt.Sprintln(buffer + status + " " + item.name)
		}
		// Add an empty line after each Todo List
		buffer = fmt.Sprintln(buffer)
	}

	err := WriteDataFile(TODOLIST_FILENAME, buffer)
	if err != nil {
		fmt.Println("Error writing to file: " + TODOLIST_FILENAME)
	} else {
		fmt.Println("Saved Todo Lists to file: " + TODOLIST_FILENAME)
	}
}
