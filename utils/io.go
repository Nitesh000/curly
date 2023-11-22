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
