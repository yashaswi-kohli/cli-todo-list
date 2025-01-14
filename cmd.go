package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Add(file *os.File, todoList []Todo) {
	if len(os.Args) < 3 {
		log.Fatal("Title not added")
	}

	sliceTitle := os.Args[2:]
	newTitle := strings.Join(sliceTitle, " ")

	now := time.Now()
	formattedTime := now.Format("Mon 02 Jan 03:04PM")

	newTodo := Todo{
		Src:         len(todoList),
		Title:       newTitle,
		Completed:   "❌",
		CreatedTime: formattedTime,
	}
	todoList = append(todoList, newTodo)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(todoList)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
}

func Delete(file *os.File, todoList []Todo) {
	if len(os.Args) < 2 {
		log.Fatal("Id not mentioned")
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	todoList = append(todoList[:id], todoList[id+1:]...)
	for idx := range todoList {
		todoList[idx].Src = idx
	}

	err = file.Truncate(0)
	if err != nil {
		log.Fatal("Error truncating file:", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
}

func Edit(file *os.File, todoList []Todo) {
	if len(os.Args) < 4 {
		log.Fatal("New Title/Id not mentioned")
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	sliceTitle := os.Args[3:]
	newTitle := strings.Join(sliceTitle, " ")

	todoList[id].Title = newTitle

	err = file.Truncate(0)
	if err != nil {
		log.Fatal("Error truncating file:", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
}

func Toggle(file *os.File, todoList []Todo, toggle bool) {
	if len(os.Args) < 2 {
		log.Fatal("Src no. not added")
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	if toggle {
		todoList[id].Completed = "✅"

		now := time.Now()
		formattedTime := now.Format("Mon 02 Jan 03:04PM")

		todoList[id].CompletedTime = formattedTime
		log.Fatal(todoList[id].CompletedTime)

	} else {
		todoList[id].Completed = "❌"
	}

	err = file.Truncate(0)
	if err != nil {
		log.Fatal("Error truncating file:", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(todoList)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
}

func List(todoList []Todo) {
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	tbl := table.New("Src No.", "Title", "Completed", "Created At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, val := range todoList {
		if val.Completed == "❌" {
			tbl.AddRow(val.Src, val.Title, val.Completed, val.CreatedTime)
		} else {
			tbl.AddRow(val.Src, val.Title, val.Completed, val.CreatedTime, val.CompletedTime)
		}
	}

	tbl.Print()
}
