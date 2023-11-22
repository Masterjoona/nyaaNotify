package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func CleanLogFile() {
	f, err := os.OpenFile(logFile, os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
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
	_, file, no, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		no = 0
	}
	file = filepath.Base(file)
	logMessage := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + file + ":" + fmt.Sprintf("%d", no) + " " + message
	println(logMessage)
	_, err = f.WriteString(logMessage + "\n")
	if err != nil {
		panic(err)
	}
}
