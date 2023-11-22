package utils

import (
	"curly/types"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filename string) (types.RequestType, error) {
	// NOTE: open the json file
	log.Println("filename: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening the file %s", filename)
	}

	// NOTE: read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading the file %s", filename)
	}

	// NOTE: parse the json content
	var data types.RequestType
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalf("Error parsing the file %s", filename)
	}

	defer file.Close()
	command := data.GetCurlCommand()
	ch := make(chan bool)
	go types.RunCurlCommand(command, ch)

	<-ch

	return data, nil
}

func CreateFile(filename string) {
	// NOTE: create the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error while creating the file %s", err)
	}
	defer file.Close()

	// NOTE: write the content to the file
	_, err = file.WriteString(`{
  "method": "GET | POST | PUT | DELETE",
  "url": "Your url here",
  "headers": {
    "Content-Type": "application/json"
  },
    "verbose": false,
    "body": { // you can remove the "data" if you don't have a body
    "key": "value"
    }
  }`)
	if err != nil {
		log.Fatalf("Error while writing to the file %s", err)
	}
}
