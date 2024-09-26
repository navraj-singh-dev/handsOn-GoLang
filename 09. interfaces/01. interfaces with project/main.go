package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"e.com/interface-project/note"
)

// ----- MAIN -----
func main() {
	title := getTerminalInput("\nProvide a title of your note:")
	content := getTerminalInput("Provide a content of your note:")

	// create instance of Note
	noteInstance, err := note.New(title, content)
	if err != nil {
		fmt.Println("\n", err)
		return
	}

	// display then save using interface functions
	err = note.DisplayAndSave(noteInstance) // note struct satisfies the interface OutputAndSave
	if err != nil {
		fmt.Println("\n", err)
		return
	}
	fmt.Println("-- Note Displayed And Saved Successfully --")

	// create instance of Todo
	text := getTerminalInput("\nProvide a small single line text for your todo:")
	todoInstance, err := note.NewTodo(text)
	if err != nil {
		fmt.Println("\n", err)
		return
	}

	// display then save using interface funtions
	err = note.DisplayAndSave(todoInstance) // todo struct also stisfies the interface OutputAndSave
	if err != nil {
		fmt.Println("\n", err)
		return
	}
	fmt.Println("-- Todo Displayed And Saved Successfully --")
}

// ----- FUNCTIONS -----
func getTerminalInput(prompt string) string {
	fmt.Print(prompt, " ")
	reader := bufio.NewReader(os.Stdin)
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSuffix(inputText, "\n")
	inputText = strings.TrimSuffix(inputText, "\r")

	return inputText
}
