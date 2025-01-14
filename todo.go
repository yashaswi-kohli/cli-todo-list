package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Todo struct {
	Src           int    `json:"src"`
	Title         string `json:"title"`
	Completed     string `json:"completed"`
	CreatedTime   string `json:"createdTime"`
	CompletedTime string `json:"completedTime"`
}

func main() {
	options := ""
	if len(os.Args) > 1 {
		options = os.Args[1]
	}

	var file *os.File
	createdNow := true
	filepath := "todoList.json"

	if _, err := os.Stat(filepath); err == nil {
		createdNow = false
	} else if !os.IsNotExist(err) {
		log.Fatal("Error checking file:", err)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0677)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	if createdNow {
		data := []byte("[]")
		_, err = file.Write(data)
		if err != nil {
			log.Fatal("Error writing to file:", err)
		}
	}

	file.Seek(0, 0)
	todoList := convertJsonToStruct(file)

	file.Seek(0, 0)
	switch options {
	case "add":
		Add(file, todoList)
	case "delete":
		Delete(file, todoList)
	case "edit":
		Edit(file, todoList)
	case "list":
		List(todoList)
	case "com":
		Toggle(file, todoList, true)
	case "incom":
		Toggle(file, todoList, false)
	default:
		displayAllOptions()
	}
}

func displayAllOptions() {
	fmt.Print("Cmds for todo: \n\tadd <title_name>\n\t\tAdd a new todo and give a new title for it \n\tdel <src_no> \n\t\tDelete the todo for the given src number \n\tedit <src_no> <new_title> \n\t\tEdit the todo by giving a src number and a new title \n\tlist <src_no> \n\t\tList all the todos \n\tcom <src_no> \n\t\tToogle the todo to complete \n\tincom <src_no> \n\t\tToogle the todo to incomplete\n")
}

func convertJsonToStruct(file *os.File) []Todo {

	var todoList []Todo
	err := json.NewDecoder(file).Decode(&todoList)
	if err != nil {
		log.Fatal("Error decoding JSON:", err.Error())
	}

	return todoList
}
