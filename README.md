scribble
========

Go logger with support for linux logrotate interrupts

**Note:** A call to scribble.Fatal() will print the message to the log file and then call panic()

###Example:
```go
package main

import (
    "github.com/dsledge/scribble"
    "flag"
)

var logfile = flag.String("logfile", "console", "The default log file will log to system console")
var loglevel = flag.Int("loglevel", 2, "Sets the default log level to INFO messages and higher")

func main() {
    flag.Parse()
  
    // Configuring the log for console or file
    if *logfile == "console" {
        scribble.NewConsoleLogger(*loglevel)
    } else {
        scribble.NewFileLogger(*loglevel, *logfile)
    }
    
    scribble.Trace("Trace level log message to "%s", *logfile) 
    scribble.Debug("Debug level log message to "%s", *logfile) 
    scribble.Info("Info level log message to "%s", *logfile) 
    scribble.Warn("Warn level log message to "%s", *logfile) 
    scribble.Error("Error level log message to "%s", *logfile) 
    //scribble.Fatal("Fatal level log message to "%s", *logfile) 
}
```

###Output Format:
```bash
2014-06-18 03:53:49.648 [scribble_test.go:4] TRACE - Test trace message
2014-06-18 03:53:49.648 [scribble_test.go:5] DEBUG - Test debug message
2014-06-18 03:53:49.648 [scribble_test.go:6] INFO  - Test info message
2014-06-18 03:53:49.648 [scribble_test.go:7] WARN  - Test warn message
2014-06-18 03:53:49.648 [scribble_test.go:8] ERROR - Test error message
```
