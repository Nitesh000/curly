package types

import (
	"fmt"
	"log"
	"os/exec"
)

func (t *RequestType) GetCurlCommand() string {
	// var command string

	command := "curl"

	command += " -X " + t.Method

	command += " " + t.Url

	for key, value := range t.Headers {
		command += " -H '" + key + ": " + value + "'"
	}

	if t.Data != nil {
		if *t.Data != "" {
			command += " -d '" + *t.Data + "'"
		}
	}

	if t.Verbose {
		command += " -v"
	}

	return command
}

func (t *RequestType) GetCurlCommandWithAuth(username string, password string) string {
	command := t.GetCurlCommand()

	command += " -u " + username + ":" + password

	return command
}

func (t *RequestType) GetCurlCommandWithAuthAndCookie(username string, password string, cookie string) string {
	command := t.GetCurlCommandWithAuth(username, password)

	command += " -b " + cookie

	return command
}

func RunCurlCommand(command string, ch chan bool) {
	log.Println("command is running...")
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Printf("ERROR::%+s", err)
	}
	fmt.Println(string(out))
	defer close(ch)
	ch <- true
}
