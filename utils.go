package main

import (
	"regexp"
	"strings"
	"time"
)

func MatchPost(title, includeString, regex, category string, postCategories []string) bool {
	if regex != "" {
		match, err := regexp.MatchString(regex, title)
		if err != nil {
			Logger("Error matching regex: " + err.Error())
			panic(err)
		}
		return match && matchCategory(postCategories, category)
	}
	lookForWords := strings.Split(includeString, ",")
	title = strings.ToLower(title)
	for _, word := range lookForWords {
		word = strings.ToLower(word)
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			found := false
			options := strings.Split(word[1:len(word)-1], "|")
			for _, option := range options {
				if strings.Contains(title, option) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		} else if strings.HasPrefix(word, ";") {
			if strings.Contains(title, word[1:]) {
				return false
			}
		} else {
			if !strings.Contains(title, word) {
				return false
			}
		}
	}
	return true && matchCategory(postCategories, category)
}

func TestMatches(includeString, regexString, category string) {
	Logger("Testing matches...")
	var matchString string
	if matchString = includeString; matchString == "" {
		matchString = regexString
	}
	titles := strings.Split(TestMatchTitle, ";")
	for _, title := range titles {
		if MatchPost(title, matchString, "", category, []string{"", ""}) {
			Logger("Matched: " + title)
		} else {
			Logger("Didn't match: " + title)
		}
	}
}

func matchCategory(postCategories []string, category string) bool {
	if category == "" || postCategories[0] == category || postCategories[1] == category {
		return true
	}
	return false
}

func ParamCreation(flag, value string) string {
	if value != "" {
		return " -" + flag + "='" + value + "'"
	}
	return ""
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}
