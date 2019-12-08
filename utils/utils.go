package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Timer for function execution
func Timer(message string) func() {
	start := time.Now()
	return func() { log.Printf("%s: %v", message, time.Since(start)) }
}

// ScanFile - Scans lines from input txt
func ScanFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// ScanFileLinesToInt - Scans each line into list of integers
func ScanFileLinesToInt(path string, sep string) []int {
	lines := ScanFile(path)

	list := []int{}
	for _, v := range lines {
		strs := strings.Split(v, sep)
		for _, vv := range strs {
			i, err := strconv.Atoi(vv)
			if err != nil {
				log.Fatalf("Failed to convert line to integer\nline: %s\nerror: %v", vv, err)
			} else {
				list = append(list, i)
			}
		}
	}
	return list
}

// ScanFileLinesToStrings - Scans each line and splits into strings
func ScanFileLinesToStrings(path string, sep string) [][]string {
	lines := ScanFile(path)

	list := [][]string{}
	for _, v := range lines {
		list = append(list, strings.Split(v, sep))
	}
	return list
}
