package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
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
	// fileExist()

	var options string
	if len(os.Args) > 1 {
		options = os.Args[1]
	}

	//* Here we are creating table
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	tbl := table.New("Src No.", "Title", "Completed", "Created At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	switch options {
	case "add":
		fmt.Println(options)
	case "delete":
		fmt.Println(options)
	case "edit":
		fmt.Println(options)
	case "list":
		listAll(tbl)
	case "com":
		fmt.Println(options)
	case "incom":
		fmt.Println(options)
	default:
		displayAllOptions()
	}
}

func displayAllOptions() {
	fmt.Print("Cmds for todo: \n\tadd <title_name>\n\t\tAdd a new todo and give a new title for it \n\tdel <src_no> \n\t\tDelete the todo for the given src number \n\tedit <src_no> <new_title> \n\t\tEdit the todo by giving a src number and a new title \n\tlist <src_no> \n\t\tList all the todos \n\tcom <src_no> \n\t\tToogle the todo to complete \n\tincom <src_no> \n\t\tToogle the todo to incomplete\n")
}
