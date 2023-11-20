package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func CheckTitle(title string, includeString string, regex string) bool {
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

func MakeDescription(post NyaaPost, shortenerToken string, shortenerURL string) string {
	if shortenerToken == "" || shortenerURL == "" {
		torrent := "[Torrent](" + post.Torrent + ")"
		return fmt.Sprintf("Size: %s | %s | <t:%s:R>", post.Size, torrent, post.Date)
	}
	shortURL, err := ShortenURL(post.Magnet, shortenerToken, shortenerURL)
	if err != nil {
		Logger("Error shortening url: " + err.Error())
		return "[Torrent](" + post.Torrent + ") | Error shortening url."
	}
	magnet := "[Magnet](" + shortURL + ")"

	return fmt.Sprintf("Size: %s | %s | <t:%s:R>", post.Size, magnet, post.Date)

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

func CountOccurrences(source, target []byte) int {
	count := 0
	for i := 0; i < len(source); {
		index := bytes.Index(source[i:], target)
		if index == -1 {
			break
		}
		count++
		i += index + len(target)
	}
	return count
}
