package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func Add(file *os.File, todoList []Todo) {
	if len(os.Args) < 3 {
		log.Fatal("Title not added")
	}

	fmt.Println("hdhsfhshdfhs")

	newTitle := os.Args[2]

	newTodo := Todo{
		Src:         len(todoList),
		Title:       newTitle,
		Completed:   "❌",
		CreatedTime: time.Now(),
	}
	todoList = append(todoList, newTodo)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(todoList)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
