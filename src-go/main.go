package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config := loadConfig("config.json")
	port := ":8888"
	http.HandleFunc("/upload", HandleUpload(config))
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func loadConfig(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer file.Close()

	config := Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config file:", err)
	}

	return config
}

type Configuration struct {
	MediaFolder string   `json:"mediafolder"`
	Length      int      `json:"length"`
	Tokens      []string `json:"tokens"`
	Author      string   `json:"author"`
	SiteName    string   `json:"sitename"`
}
