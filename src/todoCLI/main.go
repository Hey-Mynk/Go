package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todoList []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1. Add ToDo")
		fmt.Println("2. List ToDos")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter a ToDo: ")
			todo, _ := reader.ReadString('\n')
			todoList = append(todoList, strings.TrimSpace(todo))
		case "2":
			fmt.Println("ToDo List:")
			for i, todo := range todoList {
				fmt.Printf("%d. %s\n", i+1, todo)
			}
		case "3":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
