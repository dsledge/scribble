package scribble

import (
	"fmt"
)

func NewConsoleLogger(level int) {
	// Initializing the needed variables
	loglevel = level
	write = make(chan string)
	done = make(chan bool)
	stop = make(chan bool)

	go writeToConsole()
}

// Writes the data to the logfile on disk.
func writeToConsole() {
	CONSOLE:
	for {
		select {
		case msg := <-write:
			fmt.Print(msg+"\n")
		case <-stop:
			break CONSOLE
		}
	}
}