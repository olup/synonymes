package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	handler "github.com/olup/api/handler"
)

func main() {
	dictionary := handler.Open()
	log.Println("Loaded Dictionary")

	os.MkdirAll("./build", os.ModePerm)

	for word, entry := range dictionary {
		jsonData, _ := json.MarshalIndent(entry, "", " ")
		ioutil.WriteFile("./build/"+word, jsonData, 0644)
	}

	log.Println("File generated")

}
