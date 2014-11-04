package scribble

import (
	"fmt"
)

const (
	BLACK  	= "\033[0;30m"
	RED    	= "\033[0;31m"
	GREEN  	= "\033[0;32m"
	YELLOW 	= "\033[0;33m"
	BLUE	= "\033[0;34m"
	PURPLE 	= "\033[0;35m"
	CYAN 	= "\033[0;36m"
	RESET  	= "\033[0;0m"
)

func NewConsoleLogger(level int) {
	// Initializing the needed variables
	loglevel = level
	write = make(chan string)
	done = make(chan bool)
	stop = make(chan bool)
	color = true

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
