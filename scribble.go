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
//    2013-04-09 05:27:31.965 [TRACE] <scibble_test.go:62> - Test trace message
//    2013-04-09 05:27:31.965 [DEBUG] <scibble_test_test.go:63> - Test debug message
//    2013-04-09 05:27:31.965 [INFO ] <scibble_test.go:64> - Test info message
//    2013-04-09 05:27:31.965 [WARN ] <scibble.go:65> - Test warn message
//    2013-04-09 05:27:31.965 [ERROR] <scibble_test.go:66> - Test error message
//    2013-04-09 05:27:31.965 TRACE [scibble_test.go:62] - Test trace message
//    2013-04-09 05:27:31.965 DEBUG [scibble_test.go:63] - Test debug message
//    2013-04-09 05:27:31.965 INFO  [scibble_test.go:64] - Test info message
//    2013-04-09 05:27:31.965 WARN  [scibble_test.go:65] - Test warn message
//    2013-04-09 05:27:31.965 ERROR [scibble_test.go:66] - Test error message
//    2013-04-09 05:27:31.965 [scibble_test.go:62] <TRACE> Test trace message
//    2013-04-09 05:27:31.965 <scibble_test.go:63> [DEBUG] Test debug message
//    2013-04-09 05:27:31.965 [scibble_test.go:64] INFO - Test info message
//    2013-04-09 05:27:31.965 [scibble_test.go:65] WARN - Test warn message
//    2013-04-09 05:27:31.965 scibble_test.go:66   [ERROR] Test error message
//    2013-04-09 05:27:31.965 scibble_test.go:66   [FATAL] Test fatal message
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
	return fmt.Sprintf("%-23s %-25.25s %-8s", t.Format("2006-01-02 15:04:05.999"), "[" + files[len(files)-1] + ":" + strconv.Itoa(line) + "]", parseLevel(level))
	//return fmt.Sprintf("%-23s [%-20.20s:%4.4s] %-5s - ", t.Format("2006-01-02 15:04:05.999"), files[len(files)-1], strconv.Itoa(line), parseLevel(level))
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
		write <-BLUE+prefix(TRACE)+result+BLACK
	}
}

// Write a debug message to the log file.
//
// Example:
//     scribble.Debug("This is test message number %d", 22)
func Debug(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= DEBUG {
		write <-GREEN+prefix(DEBUG)+result+BLACK
	}
}

// Write a info message to the log file.
//
// Example:
//     scribble.Info("This is test message number %d", 22)
func Info(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= INFO {
		write <-BLACK+prefix(INFO)+result+BLACK
	}
}

// Write a warn message to the log file.
//
// Example:
//     scribble.Warn("This is test message number %d", 22)
func Warn(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= WARN {
		write <-YELLOW+prefix(WARN)+result+BLACK
	}
}

// Write an error message to the log file.
//
// Example:
//     scribble.Error("This is test message number %d", 22)
func Error(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= ERROR {
		write <-RED+prefix(ERROR)+result+BLACK
	}
}

// Write a fatal message to the log file.
//
// Example:
//     scribble.Fatal("This is test message number %d", 22)
func Fatal(msg string, data ...interface{}) {
	result := fmt.Sprintf(msg, data...)
	if loglevel <= FATAL {
		write <-PURPLE+prefix(FATAL)+result+BLACK
		panic(result)
	}
}
