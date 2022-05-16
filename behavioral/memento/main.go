package main

import "fmt"

func main() {
	history := NewHistory()

	tasks := NewTasks()

	tasks.Add("Task 1")
	history.Save(tasks.Memento())

	tasks.Add("Task 2")
	history.Save(tasks.Memento())

	fmt.Println(tasks.All())

	tasks.Restore(history.Undo())
	fmt.Println(tasks.All())

	tasks.Restore(history.Undo())
	fmt.Println(tasks.All())
}
