package logger

import (
	"log"
)

func init() {
	// Adds timestamps + file:line to log output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Info(format string, v ...any) {
	log.Printf("[INFO] "+format, v...)
}

func Error(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}
