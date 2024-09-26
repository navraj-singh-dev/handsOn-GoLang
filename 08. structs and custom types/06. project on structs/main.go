package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"e.com/struct-project/note"
)

func main() {

	// user's inputted data from command line
	fmt.Println()
	title := getTerminalInput("Provide a title of your note:")
	content := getTerminalInput("Provide a content of your note:")

	// create a new note instance to store the input data
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// display the note on terminal
	note.Display()

	// save the note
	err = note.Save()
	if err != nil {
		fmt.Println(errors.New("error while saving the file"))
		return
	}

	fmt.Println("\nYour note as a .json file is successfully saved!")
}

// Uses bufio to get longer user input, fmt.Scan only supports one word input
func getTerminalInput(prompt string) string {
	fmt.Print(prompt, " ")

	// allow user to give longer input text in terminal using bufio
	// NewReader takes needs Std input (the terminal's command line) as argument
	reader := bufio.NewReader(os.Stdin)

	// This function needs to know when to consired the command line to be ended
	// So if user presses enter, its a line break "\n" so consider it ended at
	// at point when enter is pressed. It returns back the String|Text that was written
	// on the command line
	inputText, _ := reader.ReadString('\n')

	// The reader.ReadString includes the delimeter with the user input from cmd line
	// It must be removed
	inputText = strings.TrimSuffix(inputText, "\n")
	inputText = strings.TrimSuffix(inputText, "\r")

	// return the user's cmd line input
	return inputText
}
