package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type postedJSON struct {
	PostedURLs []string `json:"postedURLs"`
	LastMod    string   `json:"lastMod"`
	CleanDate  string   `json:"cleanDate"`
}

func AlreadyPosted(checkURL string) bool {
	postedURLsString := GetField("PostedURLs")
	if postedURLsString == "" {
		return false
	}
	postedURLs := strings.Split(postedURLsString, ",")
	for _, url := range postedURLs {
		if url == checkURL {
			return true
		}
	}
	return false
}

func IsAmount(amount string) bool {
	postedURLsString := GetField("PostedURLs")
	if postedURLsString == "" {
		return false
	}
	postedURLs := strings.Split(postedURLsString, ",")
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		Logger("Error converting amount to int: " + err.Error())
		panic(err)
	}

	return len(postedURLs) >= amountInt
}

func GetField(field string) string {
	f, err := os.Open(postedFile)
	if err != nil {
		return ""
	}
	defer f.Close()

	postedJSON := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSON)
	if err != nil {
		return ""
	}

	switch field {
	case "CleanDate":
		return postedJSON.CleanDate
	case "LastMod":
		return postedJSON.LastMod
	case "PostedURLs":
		return strings.Join(postedJSON.PostedURLs, ",")
	default:
		return ""
	}
}

func SetField(field string, value string) {
	f, err := os.Open(postedFile)
	if err != nil {
		Logger("Error opening file: " + err.Error())
		f, err = os.Create(postedFile)
		if err != nil {
			Logger("Error creating file: " + err.Error())
			panic(err)
		}
	}
	defer f.Close()

	postedJSONUpdate := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSONUpdate)
	if err != nil {
		postedJSONUpdate = postedJSON{
			PostedURLs: []string{},
			LastMod:    "",
			CleanDate:  "",
		}
	}
	switch field {
	case "CleanDate":
		postedJSONUpdate.CleanDate = value
	case "LastMod":
		postedJSONUpdate.LastMod = value
	case "PostedURLs":
		postedJSONUpdate.PostedURLs = append(postedJSONUpdate.PostedURLs, value)
	default:
		return
	}

	updatedData, err := json.MarshalIndent(postedJSONUpdate, "", "    ")
	if err != nil {
		Logger("Error marshalling JSON: " + err.Error())
	}
	err = os.WriteFile(postedFile, updatedData, 0644)
	if err != nil {
		Logger("Error writing file: " + err.Error())
	}
}
