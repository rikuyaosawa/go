package main

import (
	"encoding/json"
	"fmt"
	"os"

	"practice/note/Note"
)

const noteFile = "note.json"

func main() {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	fmt.Println("Debug:", title, content)
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	newNote.DisplayNote()

	b, err := json.Marshal(newNote)
	if err != nil {
		fmt.Print(err)
		return
	}

	writeToJson(b)

	var text string
	json.Unmarshal([]byte(noteFile), &text)
	fmt.Println(text)
}

func getUserInput(promptText string) string {
	var userInput string
	fmt.Print(promptText)
	fmt.Scanln(&userInput)
	return userInput
}

func writeToJson(note []byte) error {
	return os.WriteFile(noteFile, note, 0o644)
}
