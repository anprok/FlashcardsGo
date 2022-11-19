package main

import "fmt"

const (
	FiveTasks  = 5
	ThreeTasks = 3
	OneTask    = 1
)

func main() {
	var activeTasks, pendingTasks int
	fmt.Scan(&activeTasks, &pendingTasks)

	totalTaskCount := activeTasks + pendingTasks

	switch {
	case totalTaskCount >= FiveTasks:
		fmt.Println("overload")
	case totalTaskCount >= ThreeTasks:
		fmt.Println("average")
	case totalTaskCount >= OneTask:
		fmt.Println("low")
	default:
		fmt.Println("idle")
	}
}
