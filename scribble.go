// Application logger for Google Go. The file logger is built to work with the
// linux logrotate process for the handling of file rotation. For log rotation
// to work a logrotation config file must be created and placed in the 
// /etc/logrotate.d directory.
//
// Example:
//    /var/log/<logfile> {
//        rotate 7
//        daily 
//        compress
//        postrotate
//            /usr/bin/killall -SIGUSR1 <application daemon defined by init.d>
//        endscript
//    }
package scribble

import (
	"runtime"
	"strconv"
	"strings"
	"time"
	"fmt"
	"os"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var (
	logfile *os.File
	loglevel int
	filename string
	write chan string
	done chan bool
	stop chan bool
)

// Builds the logfile prefix to print in front of the message. An example of the format
// is diplayed below.
//
// Example:
//    2013-04-09 05:27:31.965 [scribble_test.go:62] TRACE - Test trace message
//    2013-04-09 05:27:31.965 [scribble_test.go:63] DEBUG - Test debug message
//    2013-04-09 05:27:31.965 [scribble_test.go:64] INFO  - Test info message
//    2013-04-09 05:27:31.965 [scribble_test.go:65] WARN  - Test warn message
//    2013-04-09 05:27:31.965 [scribble_test.go:66] ERROR - Test error message
//    2013-04-09 05:27:31.965 [scribble_test.go:67] FATAL - Test fatal message
func prefix(level int) string {
	var files []string
	_, filepath, line, ok := runtime.Caller(2)
	if ok {
		files = strings.Split(filepath, "/")
	} else {
		files[0] = ""
		line = 0
	}

	t := time.Now()
	return t.Format("2006-01-02 15:04:05.999") + " [" + files[len(files)-1] + ":" + strconv.Itoa(line) + "] " + parseLevel(level) + " - "
}

// Parses the log level enum into a string value for printing to the log file.
func parseLevel(level int) string {
	switch level {
	case 0:
		return "TRACE"
	case 1:
		return "DEBUG"
	case 2:
		return "INFO "
	case 3:
		return "WARN "
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	}

	return "UNKWN"
}

// Write a trace message to the log file.
//
// Example:
//     scribble.Trace("This is test message number %d", 22)
func Trace(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= TRACE {
		write <-prefix(TRACE)+result
	}
}

// Write a debug message to the log file.
//
// Example:
//     scribble.Debug("This is test message number %d", 22)
func Debug(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= DEBUG {
		write <-prefix(DEBUG)+result
	}
}

// Write a info message to the log file.
//
// Example:
//     scribble.Info("This is test message number %d", 22)
func Info(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= INFO {
		write <-prefix(INFO)+result
	}
}

// Write a warn message to the log file.
//
// Example:
//     scribble.Warn("This is test message number %d", 22)
func Warn(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= WARN {
		write <-prefix(WARN)+result
	}
}

// Write an error message to the log file.
//
// Example:
//     scribble.Error("This is test message number %d", 22)
func Error(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= ERROR {
		write <-prefix(ERROR)+result
	}
}

// Write a fatal message to the log file.
//
// Example:
//     scribble.Fatal("This is test message number %d", 22)
func Fatal(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= FATAL {
		write <-prefix(FATAL)+result
		panic(result)
	}
}
