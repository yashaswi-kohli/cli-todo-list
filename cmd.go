package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Add(file *os.File, todoList []Todo) {
	if len(os.Args) < 3 {
		log.Fatal("Title not added")
	}

	sliceTitle := os.Args[2:]
	newTitle := strings.Join(sliceTitle, " ")

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

func Delete(file *os.File, todoList []Todo) {
	if len(os.Args) < 2 {
		log.Fatal("Id not mentioned")
	}

	id, err := strconv.Atoi(os.Args[2])
	fmt.Println(id)
	if err != nil {
		log.Fatal("Invalid ID")
	}

	todoList = append(todoList[:id], todoList[id+1:]...)
	for idx := range todoList {
		todoList[idx].Src = idx
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func Edit(file *os.File, todoList []Todo) {
	if len(os.Args) < 4 {
		log.Fatal("New Title/Id not mentioned")
	}

	id, err := strconv.Atoi(os.Args[2])
	fmt.Println(id)
	if err != nil {
		log.Fatal("Invalid ID")
	}

	sliceTitle := os.Args[3:]
	newTitle := strings.Join(sliceTitle, " ")

	todoList[id].Title = newTitle

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func Toggle(file *os.File, todoList []Todo, toggle bool) {
	if len(os.Args) < 2 {
		log.Fatal("Src no. not added")
	}

	id, err := strconv.Atoi(os.Args[2])
	fmt.Println(id)
	if err != nil {
		log.Fatal("Invalid ID")
	}

	if toggle {
		todoList[id].Completed = "✅"
		todoList[id].CompletedTime = time.Now()
	} else {
		todoList[id].Completed = "❌"
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
