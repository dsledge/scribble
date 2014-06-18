package scribble

import (
	"testing"
	"time"
	"log"
)

func TestFormatMessage(t *testing.T) {
	log.Print("Testing formatted messages with paramters")
	NewFileLogger(DEBUG, "format.log")
	Debug("formatting the number %d", 23)
	Debug("formatting the float %f", 23.3)
	Debug("formatting the string %s", "I printed a string")
	Debug("formatting a boolean %t", true)
	Debug("formatting multiple values %s %t", "|Test string|", true)
}

func TestTrace(t *testing.T) {
	log.Print("Testing trace message level")
	NewFileLogger(TRACE, "trace.log")
	logMessages()
}

func TestDebug(t *testing.T) {
	log.Print("Testing debug message level")
	NewFileLogger(DEBUG, "debug.log")
	logMessages()
}

func TestInfo(t *testing.T) {
	log.Print("Testing info message level")
	NewFileLogger(INFO, "info.log")
	logMessages()
}

func TestWarn(t *testing.T) {
	log.Print("Testing warn message level")
	NewFileLogger(WARN, "warn.log")
	logMessages()
}

func TestError(t *testing.T) {
	log.Print("Testing error message level")
	NewFileLogger(ERROR, "error.log")
	logMessages()
}

func TestFatal(t *testing.T) {
	log.Print("Testing fatal message level")
	NewFileLogger(FATAL, "fatal.log")
	logMessages()
}

func TestLogging(t *testing.T) {
	log.Print("Testing continuous concurrent logging for 30 seconds")
	done := make(chan bool)
	NewFileLogger(DEBUG, "test.log")

	go func() {
		for {
			time.Sleep(30e9)
			done <-true
		}
	}();

	TEST:
	for {
		select {
		case <-time.After(0.00001e9):
			go logMessages()
		case <-done:
			break TEST
		}
	}
}