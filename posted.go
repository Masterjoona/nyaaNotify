package main

import (
	"encoding/json"
	"os"
	"strconv"
)

type postedJSON struct {
	PostedURLs []string `json:"postedURLs"`
	CleanDate  string   `json:"cleanDate"`
}

func StoreInJSON(post NyaaPost) {
	existingData, err := os.ReadFile(postedFile)
	if err != nil {
		Logger("Error reading file: " + err.Error())
	}
	jsonData := postedJSON{}
	if err := json.Unmarshal(existingData, &jsonData); err != nil && len(existingData) > 0 {
		Logger("Error unmarshalling JSON: " + err.Error())
	}

	jsonData.PostedURLs = append(jsonData.PostedURLs, post.URL)

	updatedData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		Logger("Error marshalling JSON: " + err.Error())
	}

	err = os.WriteFile(postedFile, updatedData, 0644)
	if err != nil {
		Logger("Error writing file: " + err.Error())
	}

}

func AlreadyPosted(post NyaaPost) bool {
	f, err := os.Open(postedFile)
	if err != nil {
		return false
	}
	defer f.Close()
	postedJSON := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSON)
	if err != nil {
		return false
	}
	for _, url := range postedJSON.PostedURLs {
		if url == post.URL {
			return true
		}
	}
	return false
}

func IsAmount(amount string) bool {
	f, err := os.Open(postedFile)
	if err != nil {
		return false
	}
	defer f.Close()
	postedJSON := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSON)
	if err != nil {
		return false
	}
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		Logger("Error converting amount to int: " + err.Error())
		panic(err)
	}

	return len(postedJSON.PostedURLs) >= amountInt
}

func SetCleanDateInJSON() {
	f, err := os.Open(postedFile)
	if err != nil {
		return
	}
	defer f.Close()
	postedJSON := postedJSON{}
	err = json.NewDecoder(f).Decode(&postedJSON)
	if err != nil {
		return
	}
	postedJSON.CleanDate = GetCleanDateString()
	updatedData, err := json.MarshalIndent(postedJSON, "", "    ")
	if err != nil {
		Logger("Error marshalling JSON: " + err.Error())
	}
	err = os.WriteFile(postedFile, updatedData, 0644)
	if err != nil {
		Logger("Error writing file: " + err.Error())
	}
}

func GetCleanDate() string {
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
	return postedJSON.CleanDate
}
