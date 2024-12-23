package main

func main() {
	tasks := Tasks{}
	storage := NewStorage[Tasks]("tasks.json")
	storage.Load(&tasks)
	tasks.printTaskTable()
	storage.Save(tasks)
}
