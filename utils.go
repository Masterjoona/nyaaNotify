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
	title = strings.ToLower(title)
	for _, word := range lookForWords {
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

func CreateNyaaPost(matches []string) NyaaPost {
	nyaaPost := NyaaPost{}
	nyaaPost.Category = matches[1]
	nyaaPost.CategoryImg = Url + matches[2]
	nyaaPost.URL = Url + matches[3]
	nyaaPost.Title = matches[4]
	nyaaPost.Comments = matches[5]
	nyaaPost.Torrent = Url + matches[8]
	nyaaPost.Magnet = matches[9]
	nyaaPost.Size = matches[10]
	nyaaPost.Date = matches[11]
	nyaaPost.Seed = matches[12]
	nyaaPost.Leech = matches[13]
	nyaaPost.Completed = matches[14]
	return nyaaPost
}
