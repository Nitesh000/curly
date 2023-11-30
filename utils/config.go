package utils

import (
	"curly/types"
	"errors"
	"os"
)

func CreateConfigureFile() error {
	// NOTE: finding the path value of bash "HOME" variable
	home := os.Getenv("HOME")
	if home == "" {
		return errors.New("HOME environment variable is not found")
	}

	// NOTE: create configure file in the ~/.curly/config.json
	fullPath := home + "/" + types.CONFIGURE_FILE_PATH + "/" + types.CONFIGURE_FILE_NAME

	// NOTE: check if the file is already exist then ask the user to update or delete the file.
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		return errors.New("The configuration file is already exist.")
	}

	// NOTE: create a directory first if it is not exist
	if _, err := os.Stat(home + "/" + types.CONFIGURE_FILE_PATH); os.IsNotExist(err) {
		err := os.MkdirAll(home+"/"+types.CONFIGURE_FILE_PATH, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// NOTE: create the file
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// NOTE: write the content to the file
	_, err = file.WriteString(`{
    "total_time": true,
    "namelookup_time": false,
    "connect_time": true,
    "appconnect_time": false,
    "pretransfer_time": false,
    "redirect_time": false,
    "starttransfer_time": false,
    "redirect_count": false,
    "size_upload": false,
    "size_download": false,
    "speed_download": false,
    "speed_upload": false,
    "header_size": false,
    "request_size": false,
    "ssl_verify_result": false,
    "filetime_connect": false,
    "filetime_pretransfer": false,
    "filetime_starttransfer": false,
    "filetime_appconnect": false,
    "remote_ip": false,
    "remote_port": false,
    "local_ip": false,
    "local_port": false,
    "redirect_url": false,
    "primary_ip": false,
    "primary_port": false,
    "curl_error": false,
    "http_code": true,
    "effective_url": false,
    "content_type": false,
    "num_connects": false,
    "primary_ip6": false,
    "local_ip6": false
    }`)

	if err != nil {
		return errors.New("Error while writing the configuration file")
	}

	return nil
}
