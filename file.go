package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var postedFile string
var logFile string
var executablePath string

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func AlreadyPosted(post NyaaPost) bool {
	f, err := os.ReadFile(postedFile)
	if err != nil {
		return false
	}

	if bytes.Contains(f, []byte(post.URL)) {
		return true
	}
	return false
}

func IsOverAmount(amount string) bool {
	f, err := os.ReadFile(postedFile)
	if err != nil {
		return false
	}
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		Logger("Error converting amount to int: " + err.Error())
		panic(err)
	}

	return CountOccurrences(f, []byte("nyaa.si")) >= amountInt
}

func StorePosted(url string) {
	f, err := os.OpenFile(postedFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger("Error opening posted file: " + err.Error())
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(url + "\n")
	if err != nil {
		Logger("Error writing to posted file: " + err.Error())
		panic(err)
	}

}

func CheckDate() bool {
	f, err := os.ReadFile(postedFile)
	if err != nil {
		return false
	}
	if CountOccurrences(f, []byte("-")) == 0 {
		return false
	}
	lines := bytes.Split(f, []byte("\n"))
	lastLine := lines[len(lines)-1]
	if string(lastLine) == "" {
		lastLine = lines[len(lines)-2]
	}
	return string(lastLine) == GetDate()
}

func CleanPosted() {
	f, err := os.OpenFile(postedFile, os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
}

func Logger(message string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logMessage := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + message
	println(logMessage)
	_, err = f.WriteString(logMessage + "\n")
	if err != nil {
		panic(err)
	}
}
