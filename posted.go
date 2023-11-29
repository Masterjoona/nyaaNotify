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

func openAndDecodeFile() (postedJSON, *os.File, error) {
	f, err := os.Open(PostedFile)
	if err != nil {
		return postedJSON{}, nil, err
	}
	defer f.Close()

	postedJSONData := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSONData)
	if err != nil {
		return postedJSON{}, nil, err
	}

	return postedJSONData, f, nil
}

func GetField(field string) string {
	postedJSONData, _, err := openAndDecodeFile()
	if err != nil {
		return ""
	}

	switch field {
	case "LastMod":
		return postedJSONData.LastMod
	case "PostedURLs":
		return strings.Join(postedJSONData.PostedURLs, ",")
	default:
		return ""
	}
}

func SetField(field string, value string) {
	postedJSONData, f, err := openAndDecodeFile()
	if err != nil {
		Logger("Error opening/decoding file: " + err.Error())
		f, err = os.Create(PostedFile)
		if err != nil {
			Logger("Error creating file: " + err.Error())
			panic(err)
		}

		postedJSONData = postedJSON{
			PostedURLs: []string{},
			LastMod:    "",
		}
	}

	defer f.Close()

	switch field {
	case "LastMod":
		postedJSONData.LastMod = value
	case "PostedURLs":
		postedJSONData.PostedURLs = append(postedJSONData.PostedURLs, value)
	default:
		return
	}

	updatedData, err := json.MarshalIndent(postedJSONData, "", "    ")
	if err != nil {
		Logger("Error marshalling JSON: " + err.Error())
	}
	err = os.WriteFile(PostedFile, updatedData, 0644)
	if err != nil {
		Logger("Error writing file: " + err.Error())
	}
}
