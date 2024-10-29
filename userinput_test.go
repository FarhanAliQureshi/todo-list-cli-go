package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPromptUserForMenuChoice(t *testing.T) {
	tests := []struct {
		testName, input      string
		rangeStart, rangeEnd int
		expected             int
		expectError          bool
	}{
		{"Empty input", "", 1, 5, 0, true},
		{"Lower-boundary out-of-range", "0", 1, 5, 0, true},
		{"Upper-boundary out-of-range", "6", 1, 5, 0, true},
		{"Number with negative sign", "-1", 1, 5, 0, true},
		{"Number with positive sign", "+4", 1, 5, 4, false},
		{"Two numbers with a space", "1 2", 1, 5, 1, false},
		{"Number with a comma", "1,2", 1, 5, 0, true},
		{"Number with a decimal", "1.2", 1, 5, 0, true},
		{"String input", "abcd", 1, 5, 0, true},
		{"Edge case lower-boundary", "1", 1, 5, 1, false},
		{"Edge case upper-boundary", "5", 1, 5, 5, false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s (%d, %d, %s)", test.testName,
			test.rangeStart, test.rangeEnd, test.input)

		t.Run(testName, func(t *testing.T) {
			reader := strings.NewReader(test.input + "\n")
			output, err := promptUserForMenuChoice(reader, test.rangeStart, test.rangeEnd)

			if test.expectError {
				if err == nil {
					t.Errorf("Was expecting an error but got [%d]", output)
				}
			} else if output != test.expected {
				t.Errorf("Was expecting [%d] but got [%d]", test.expected, output)
			} else if output == test.expected {
				// Test passed. Do nothing
			} else {
				t.Errorf("Something unexpected happened. Output [%d]. Error [%s]",
					output, err)
			}
		})
	}
}
