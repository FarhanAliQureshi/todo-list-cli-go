package main

import "strings"

type task struct {
	name      string
	completed bool
}

type todo struct {
	name  string
	tasks []task
}

func parseStringAndCreateTodoLists(data string) []todo {
	todos := []todo{}
	// Before splitting the data into lines, remove CR from newline because in
	// Windows newline is CRLF, while in Linux newline is just LF
	data = strings.ReplaceAll(data, "\r", "")
	lines := strings.Split(data, "\n")
	todoList := todo{}

	for _, line := range lines {
		if len(line) == 0 {
			// Skip empty lines
			continue

		} else if len([]rune(line)) < 3 {
			// There will be a space between indicator and text (task or title)
			// Therefore, a line must have at least 3 characters to be a valid
			// entry for todo list. Ignore all lines with less than 3 characters.
			// Using rune to count for Unicode characters.
			continue

		} else if string(line[0]) == "#" {
			// Title of a todo list
			// If we were parsing a todo list then save it before starting new list
			if todoList.name != "" {
				todos = append(todos, todoList)
			}

			// Create new todo list
			todoList = todo{}
			todoList.name = string([]rune(line))[2:]

		} else if string(line[0]) == "-" {
			// List task which is not yet completed
			todoTask := task{string([]rune(line))[2:], false}
			todoList.tasks = append(todoList.tasks, todoTask)

		} else if string(line[0]) == "+" {
			// List task which is completed
			todoTask := task{string([]rune(line))[2:], true}
			todoList.tasks = append(todoList.tasks, todoTask)
		}
	}

	// Save last todo list which we were parsing (if any)
	if todoList.name != "" {
		todos = append(todos, todoList)
	}

	return todos
}
