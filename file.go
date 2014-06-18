package scribble

import (
	"os/signal"
	"syscall"
	"log"
	"os"
)

// This function creates a new logger and sets the log level required. If a log
// file already exist then it is closed and a new file is opened.
func NewFileLogger(level int, filename string) {
	// Checking if the logfile is already open, if so it will be closed.
	if logfile != nil {
		done <-true
		closeLog()
	}

	// Initializing the needed variables
	loglevel = level
	write = make(chan string)
	done = make(chan bool)
	stop = make(chan bool)

	// Setting up the signal handler for log rotate
	sigRecv := make(chan os.Signal)
	signal.Notify(sigRecv, syscall.SIGUSR1)
	go sigListener(sigRecv)

	// Opening the logfile for writing
	openLog(filename)
}

// Opens a new log file for writing
func openLog(fn string) {
	filename = fn
	var err error
	logfile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
        if err != nil {
            log.Fatal("ERROR: ", err)
        }
	go writeToFile(logfile)
}

// Closes an open log file. It sends a true status to the stop channel to stop
// the file writer goroutine from writing to the log file.
func closeLog() {
	stop <-true
	logfile.Close()
}

// Handler for the linux SIGUSR1 signal to facilitate logrotation.
func sigListener(sig chan os.Signal) {
	SIGNAL:
	for {
		select {
		case <-sig:
			closeLog()
			openLog(filename)
		case <-done:
			break SIGNAL
		}
	}
}

// Writes the data to the logfile on disk.
func writeToFile(file *os.File) {
	WRITE:
	for {
		select {
		case msg := <-write:
			_, err := file.WriteString(msg+"\n")
			if err != nil {
				log.Print("Error: ", err)
			}
		case <-stop:
			break WRITE
		}
	}
}
