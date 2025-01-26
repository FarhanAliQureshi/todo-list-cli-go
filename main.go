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
			displayTitlesOfTodoLists(&todos, true)
		} else if userSelection == 2 { // Menu: Display all tasks of a Todo List
			displayTodoListTasks(&todos, true)
		} else if userSelection == 3 { // Menu: Create a new Todo List
			createNewTodoList(&todos)
		} else if userSelection == 4 { // Menu: Manage a Todo List
			manageTodoList(&todos)
		} else if userSelection == 5 { // Menu: Delete a Todo List
			deleteTodoList(&todos)
		} else if userSelection == 6 { // Menu: Save changes to file on disk
			saveTodoListToFile(&todos)
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
	fmt.Println("2. Display all tasks of a Todo List")
	fmt.Println("3. Create a new Todo List")
	fmt.Println("4. Manage a Todo List")
	fmt.Println("5. Delete a Todo List")
	fmt.Println("6. Save changes to file on disk")
	fmt.Println("0. Exit")
	fmt.Println("")
}

func displayTitlesOfTodoLists(todos *[]todo, pause bool) {
	fmt.Printf("\nTitles of all Todo Lists:\n\n")
	for index, listTask := range *todos {
		fmt.Printf("%d. %s\n", index+1, listTask.name)
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

func displayTodoListTasks(todos *[]todo, pause bool) {
	fmt.Println("")
	fmt.Println("Select a Todo List to list all its tasks")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Main Menu")
	fmt.Println("")

	userSelection := GetUserMenuChoice(0, len(*todos))
	if userSelection != 0 {
		// Index of list starts at zero, while the user input starts with 1
		// Therefore, we must first deduct 1 from userSelection to get
		// a proper index
		selectedListIndex := userSelection - 1

		fmt.Println("")
		fmt.Printf("Selected Todo List: %s\n", (*todos)[selectedListIndex].name)
		displayTodoListTasksOfSelectedList(todos, selectedListIndex)
		if pause {
			fmt.Println("")
			PauseProgramToLetUserRead()
		}
	}
}

func createNewTodoList(todos *[]todo) {
	fmt.Println("")
	fmt.Println("Create a New Todo List")
	fmt.Println("")
	fmt.Println("(Press Enter key without writing anything to cancel)")
	newTodoListName := GetInputFromUser("Enter Name for New Todo List: ")
	if newTodoListName != "" {
		todoList := todo{}
		todoList.name = newTodoListName
		*todos = append(*todos, todoList)
	}
}

func deleteTodoList(todos *[]todo) {
	fmt.Println("")
	fmt.Println("Delete a Todo List")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Cancel")
	fmt.Println("")
	fmt.Println("Select a Todo List to delete")
	userSelection := GetUserMenuChoice(0, len(*todos))

	if userSelection > 0 {
		index := userSelection - 1
		fmt.Printf("Deleting \"%s\" ...\n", (*todos)[index].name)
		*todos = append((*todos)[:index], (*todos)[index+1:]...)
	}
}

func manageTodoList(todos *[]todo) {
	fmt.Println("")
	fmt.Println("Select a Todo List to manage")
	displayTitlesOfTodoLists(todos, false)
	fmt.Println("0. Main Menu")
	fmt.Println("")

	userSelection := GetUserMenuChoice(0, len(*todos))
	if userSelection != 0 {
		selectedListIndex := userSelection - 1
		for {
			displayManageTodoListMenu(todos, selectedListIndex)
			userSelection = GetUserMenuChoice(0, 5)
			if userSelection == 1 { // Menu: Display tasks of the todo list
				displayTodoListTasksOfSelectedList(todos, selectedListIndex)
			} else if userSelection == 2 { // Menu: Mark a task as completed
				changeStatusOfTodoListTask(todos, selectedListIndex, true)
			} else if userSelection == 3 { // Menu: Mark a task as incomplete
				changeStatusOfTodoListTask(todos, selectedListIndex, false)
			} else if userSelection == 4 { // Menu: Add a task to todo list
				addNewTodoListTask(todos, selectedListIndex)
			} else if userSelection == 5 { // Menu: Delete a task from todo list
				deleteTodoListTask(todos, selectedListIndex)
			} else if userSelection == 0 { // Menu: Back to Main Menu
				break
			}
		}
	}
}

func displayManageTodoListMenu(todos *[]todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Printf("Managing Todo List: %s\n", (*todos)[selectedListIndex].name)
	fmt.Println("")
	fmt.Println("1. Display tasks of the todo list")
	fmt.Println("2. Mark a task as completed")
	fmt.Println("3. Mark a task as incomplete")
	fmt.Println("4. Add a task to todo list")
	fmt.Println("5. Delete a task from todo list")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("")
}

func displayTodoListTasksOfSelectedList(todos *[]todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Printf("Todo List Tasks for: %s\n", (*todos)[selectedListIndex].name)
	for index, listTask := range (*todos)[selectedListIndex].tasks {
		var completed string
		if listTask.completed {
			completed = "+"
		} else {
			completed = "-"
		}
		fmt.Printf("%d. %s %s\n", index+1, completed, listTask.name)
	}
}

func changeStatusOfTodoListTask(todos *[]todo, selectedListIndex int, markCompleted bool) {
	var changeStatus string
	if markCompleted {
		changeStatus = "completed"
	} else {
		changeStatus = "incomplete"
	}
	fmt.Println("")
	fmt.Printf("Select a task to mark as %s\n", changeStatus)
	displayTodoListTasksOfSelectedList(todos, selectedListIndex)
	fmt.Println("0. Cancel")
	fmt.Println("")
	userSelection := GetUserMenuChoice(0, len((*todos)[selectedListIndex].tasks))

	if userSelection != 0 {
		selectedTask := userSelection - 1
		(*todos)[selectedListIndex].tasks[selectedTask].completed = markCompleted
	}
}

func addNewTodoListTask(todos *[]todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Println("Add a New Task")
	fmt.Println("")
	fmt.Println("(Press Enter key without writing anything to cancel)")
	newTaskName := GetInputFromUser("Enter Task Task: ")
	if newTaskName != "" {
		newTask := task{}
		newTask.name = newTaskName
		(*todos)[selectedListIndex].tasks = append((*todos)[selectedListIndex].tasks, newTask)
	}
}

func deleteTodoListTask(todos *[]todo, selectedListIndex int) {
	fmt.Println("")
	fmt.Println("Delete a Todo List's Task")
	displayTodoListTasksOfSelectedList(todos, selectedListIndex)
	fmt.Println("0. Cancel")
	fmt.Println("")
	fmt.Println("Select a Todo List's Task to delete")
	userSelection := GetUserMenuChoice(0, len((*todos)[selectedListIndex].tasks))

	if userSelection > 0 {
		index := userSelection - 1
		fmt.Printf("Deleting \"%s\" ...\n", (*todos)[selectedListIndex].tasks[index].name)
		(*todos)[selectedListIndex].tasks = append(
			(*todos)[selectedListIndex].tasks[:index],
			(*todos)[selectedListIndex].tasks[index+1:]...)
	}
}

func saveTodoListToFile(todos *[]todo) {
	var buffer, status string

	for _, todoList := range *todos {
		buffer = fmt.Sprintln(buffer + "# " + todoList.name)
		for _, listTask := range todoList.tasks {
			if listTask.completed {
				status = "+"
			} else {
				status = "-"
			}
			buffer = fmt.Sprintln(buffer + status + " " + listTask.name)
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
