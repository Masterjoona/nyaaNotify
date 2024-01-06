package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func MatchPost(title, includeString, regex string, categories []string, attributes [][]string) bool {
	if regex != "" {
		match, err := regexp.MatchString(regex, title)
		if err != nil {
			Logger("Error matching regex: " + err.Error())
			panic(err)
		}
		return match && matchCategory(categories) && matchAttributes(attributes)
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
	return true && matchCategory(categories) && matchAttributes(attributes)
}

func TestMatches(includeString, regexString, category string) {
	var matchString string
	if matchString = includeString; matchString == "" {
		matchString = regexString
	}
	titles := strings.Split(TestMatchTitle, ";")
	for _, title := range titles {
		if MatchPost(title, matchString, "", []string{category, "", ""}, [][]string{{"", ""}, {"", ""}}) { // Kinda cursed
			println("Matched: " + title)
		} else {
			println("Didn't match: " + title)
		}
	}
}

func matchCategory(categories []string) bool {
	desiredCategory := categories[0]
	if desiredCategory == "" || desiredCategory == categories[1] || desiredCategory == categories[2] {
		return true
	}
	return false
}

func matchAttribute(configValue, postValue string) bool {
	postValue = strconv.FormatBool(postValue == "Yes")
	if configValue == "" || configValue == postValue {
		return true
	}
	return false
}

func matchAttributes(attributes [][]string) bool {
	return matchAttribute(attributes[0][0], attributes[1][0]) && matchAttribute(attributes[0][1], attributes[1][1])
}

// cursed

func ParamCreation(flag, value string) string {
	if value != "" {
		return " -" + flag + "='" + value + "'"
	}
	return ""
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}
