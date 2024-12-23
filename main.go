package main

func main() {
	tasks := Tasks{}
	tasks.add("Fix feature 1")
	tasks.add("Refactor code")
	tasks.toggle(0)
	tasks.printTaskTable()
}
