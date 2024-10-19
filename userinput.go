package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Prompt user for input, validate user input, and return user's choice
func GetUserMenuChoice(validFrom int, validTo int) int {
	var choice int
	var err error

	for {
		choice, err = promptUserForMenuChoice(os.Stdin, validFrom, validTo)
		if err == nil {
			break
		}
	}

	return choice
}

func promptUserForMenuChoice(reader io.Reader, validFrom int, validTo int) (int, error) {
	var userInput string

	fmt.Printf("Your Choice? [%d-%d]: ", validFrom, validTo)

	// Get user input and then try to convert it into an integer.
	// Reading integer directly from user using Scanln() doesn't
	// work as expected, especially if user entered invalid input.
	fmt.Fscanln(reader, &userInput)
	choice, err := strconv.Atoi(userInput)

	// Validate user input. It should be a number between (inclusive of)
	// validFrom and validTo.
	if err != nil {
		fmt.Println("Invalid input")
		return -1, err
	} else if choice < validFrom || choice > validTo {
		fmt.Println("Invalid choice")
		return -1, errors.New("Invalid choice")
	}

	// By this point, user's choice must be between >= validFrom and <= validTo
	// Therefore, it is a valid input, and let's return it to the calling func
	return choice, nil
}
