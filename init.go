package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	storage "url_shortner/storage"
)

func Initialize() {
	InitilizeEnv()
	storage.InitializeStorage()
}

func InitilizeEnv() {
	file, err := os.Open("go_config.json")
	if err != nil {
		panic("Failed to open go_config.json")
	}
	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Failed to read go_config.json")
	}
	configData := make(map[string]string)
	err = json.Unmarshal(dataBytes, &configData)
	if err != nil {
		panic("Failed to serialize json")
	}
	for key, val := range configData {
		os.Setenv(key, val)
	}
}
