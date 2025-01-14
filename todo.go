package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Src           int       `json:"src"`
	Title         string    `json:"title"`
	Completed     string    `json:"completed"`
	CreatedTime   time.Time `json:"createdTime"`
	CompletedTime time.Time `json:"completedTime"`
}

func main() {
	//* this will check wheather the toldo list file is present or not, if not then it will create it
	fileExist()

	//* now we will read data from toodList.json and add decode the json into our data structure so we can perform opns
	todoList := convertJsonToStruct()
	fmt.Println(todoList)

	var options string
	if len(os.Args) > 1 {
		options = os.Args[1]
	}

	switch options {
	case "add":
		fmt.Println(options)
	case "delete":
		fmt.Println(options)
	case "edit":
		fmt.Println(options)
	case "list":
		fmt.Println(options)
	case "com":
		fmt.Println(options)
	case "incom":
		fmt.Println(options)
	default:
		displayAllOptions()
	}
}

func fileExist() {

	filePath := "todoList.json"

	//* checking if file exist or not
	if _, err := os.Stat(filePath); err == nil {
		return
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file:", err)
		os.Exit(1)
	}

	//* if not then we will create file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating a file")
	}
	defer file.Close()

	//* now we will add basic empty array where we can append json of our todo list
	data := []byte("[]")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func convertJsonToStruct() []Todo {

	file, err := os.Open("todoList.json")
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		os.Exit(1)
	}

	var todoList []Todo
	err = json.NewDecoder(file).Decode(&todoList)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	return todoList
}

func displayAllOptions() {
	fmt.Print("Cmds for todo: \n\tadd <title_name>\n\t\tAdd a new todo and give a new title for it \n\tdel <src_no> \n\t\tDelete the todo for the given src number \n\tedit <src_no> <new_title> \n\t\tEdit the todo by giving a src number and a new title \n\tlist <src_no> \n\t\tList all the todos \n\tcom <src_no> \n\t\tToogle the todo to complete \n\tincom <src_no> \n\t\tToogle the todo to incomplete\n")
}
