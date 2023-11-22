package main

import (
	"regexp"
	"strings"
	"time"
)

func MatchTitle(title string, includeString string, regex string) bool {
	if regex != "" {
		match, err := regexp.MatchString(regex, title)
		if err != nil {
			Logger("Error matching regex: " + err.Error())
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

func OptionalParam(flag, value string) string {
	if value != "" {
		return " " + flag + "='" + value + "'"
	}
	return ""
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

func GetCleanDateString() string {
	return time.Now().AddDate(0, 0, 7).Format("2006-01-02")
}
