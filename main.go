package main

import (
	"fmt"
)

func main() {
	tasks := Tasks{}
	tasks.add("Fix feature 1")
	tasks.add("Refactor code")
	fmt.Printf("%+v\n\n", tasks)

	tasks.delete(0)
	fmt.Printf("%+v\n\n", tasks)
}
