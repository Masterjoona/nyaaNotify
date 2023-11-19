package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckTitle(title string, includeString string, regex string) bool {
	if regex != "" {
		match, err := regexp.MatchString(regex, title)
		if err != nil {
			panic(err)
		}
		return match
	}
	lookForWords := strings.Split(includeString, ",")
	for _, word := range lookForWords {
		title = strings.ToLower(title)
		word = strings.ToLower(word)
		if strings.HasPrefix(word, ";") {
			if strings.Contains(title, word[1:]) {
				return false
			}
		} else {
			if !strings.Contains(title, word) {
				return false
			}
		}
	}
	return true
}

func MakeDescription(post NyaaPost, shortenerToken string, shortenerURL string) string {
	magnetOrTorrent := ""
	if shortenerToken == "" {
		magnetOrTorrent = "[Torrent](" + post.Torrent + ")"
	} else {
		shortURL, err := ShortenURL(post.Magnet, shortenerToken, shortenerURL)
		if err != nil {
			Logger("Error shortening url: " + err.Error())
			return "[Torrent](" + post.Torrent + ") | Error shortening url."
		}
		magnetOrTorrent = "[Magnet](" + shortURL + ")"
	}
	return fmt.Sprintf("Size: %s | %s | <t:%s:R>", post.Size, magnetOrTorrent, post.Date)

}

func CheckIfAlreadyPosted(post NyaaPost) bool {
	f, err := os.ReadFile(executablePath + "/" + name + "_posted.txt")
	if err != nil {
		return false
	}

	if bytes.Contains(f, []byte(post.URL)) {
		Logger("Already posted: " + post.Title)
		return true
	}
	return false
}

func StorePosted(url string) {
	f, err := os.OpenFile(executablePath+"/"+name+"_posted.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(url + "\n")
	if err != nil {
		panic(err)
	}

}

func IsOverAmount(amount string) bool {
	f, err := os.ReadFile(executablePath + "/" + name + "_posted.txt")
	if err != nil {
		return false
	}
	lines := strings.Split(string(f), "\n")
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}

	return len(lines)-1 >= amountInt
}

func Logger(message string) {
	f, err := os.OpenFile(executablePath+"/"+name+"_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func OptionalParam(flag, value string) string {
	if value != "" {
		return " " + flag + "='" + value + "'"
	}
	return ""
}

func ShouldntPost() bool {
	f, err := os.ReadFile(executablePath + "/" + name + "_posted.txt")
	if err != nil {
		return false
	}
	lines := strings.Split(string(f), "\n")
	lastLine := lines[len(lines)-2]
	lastLineDate := time.Now().Format("2006-01-02")
	return lastLine == lastLineDate
}

func CleanPosted() {
	f, err := os.OpenFile(executablePath+"/"+name+"_posted.txt", os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

var executablePath = getExecutablePath()
