package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("I am a task")
}

func main() {
	// Create a new scheduler
	s := gocron.NewScheduler(time.UTC)

	// Schedule a task to run every second
	s.Every(1).Second().Do(task)

	// Schedule a task to run at
	// 10:30 every day
	s.Every(1).Day().At("10:30").Do(task)

	// Start the scheduler
	s.StartBlocking()
}
