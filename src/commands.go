package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	Pri    uint
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task; Specify title")
	flag.UintVar(&cf.Pri, "pri", 1, "Set priority from 1-5; Lowest to Highest")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing task's title; Specify index and new title")
	flag.IntVar(&cf.Del, "del", -1, "Delete an existing task; Specify by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Set an existing task's to be complete/incomplete; Specify by index")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List:
		tasks.printTaskTable()
	case cf.Add != "":
		tasks.add(cf.Add, cf.Pri)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		tasks.edit(index, parts[1])
	case cf.Toggle != -1:
		tasks.toggle(cf.Toggle)
	case cf.Del != -1:
		tasks.delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
