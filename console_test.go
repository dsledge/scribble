package scribble

import (
	"testing"
	"log"
)

func TestConsole(t *testing.T) {
	log.Print("Testing console messages")
	NewConsoleLogger(TRACE)
	logMessages()
}
