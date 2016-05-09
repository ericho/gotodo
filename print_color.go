package main

import (
	"fmt"
	c "github.com/TreyBastian/colourize"
)

func TodoPrint(text string) {
	// Prints the colour based on the type of the string
	if text == "GOTODO TASKS" || text == "========" {
		fmt.Println(c.Colourize(text, c.Yellow))
	} else if isDoneLine(text) {
		fmt.Println(c.Colourize(text, c.Grey, c.Underline))
	} else if isValidLine(text) {
		fmt.Println(c.Colourize(text, c.Cyan, c.Bold))
	}
}
