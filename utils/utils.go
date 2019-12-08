package utils

import (
	"log"
	"time"
)

// Timer for function execution
func Timer(message string) func() {
	start := time.Now()
	return func() { log.Printf("%s: %v", message, time.Since(start)) }
}
