package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const notesDir = "./notes/"

func main() {
	if err := os.MkdirAll(notesDir, os.ModePerm); err != nil {
		fmt.Println("Error creating notes directory:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Add Note")
		fmt.Println("2. List Notes")
		fmt.Println("3. View Note")
		fmt.Println("4. Update Note")
		fmt.Println("5. Delete Note")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addNote()
		case "2":
			listNotes()
		case "3":
			viewNote()
		case "4":
			updateNote()
		case "5":
			deleteNote()
		case "6":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Enter note content (end with a blank line):")
	var content string
	for {
		line, _ := reader.ReadString('\n')
		if strings.TrimSpace(line) == "" {
			break
		}
		content += line
	}

	if err := os.WriteFile(notesDir+title+".txt", []byte(content), 0644); err != nil {
		fmt.Println("Error writing note:", err)
	} else {
		fmt.Println("Note added successfully!")
	}
}

func listNotes() {
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Println("Error reading notes directory:", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("No notes found.")
		return
	}

	fmt.Println("Notes:")
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}

func viewNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title to view: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	content, err := os.ReadFile(notesDir + title + ".txt")
	if err != nil {
		fmt.Println("Error reading note:", err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(string(content))
}

func updateNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title to update: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	filepath := notesDir + title + ".txt"
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading note:", err)
		return
	}

	fmt.Println("Current content:")
	fmt.Println(string(content))

	fmt.Println("Enter new content (end with a blank line):")
	var newContent string
	for {
		line, _ := reader.ReadString('\n')
		if strings.TrimSpace(line) == "" {
			break
		}
		newContent += line
	}

	if err := os.WriteFile(filepath, []byte(newContent), 0644); err != nil {
		fmt.Println("Error updating note:", err)
	} else {
		fmt.Println("Note updated successfully!")
	}
}

func deleteNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title to delete: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	filepath := notesDir + title + ".txt"
	if err := os.Remove(filepath); err != nil {
		fmt.Println("Error deleting note:", err)
	} else {
		fmt.Println("Note deleted successfully!")
	}
}
