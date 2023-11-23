package utils

import (
	"curly/types"
	"os"
)

func CreateConfigureFile() error {
	// NOTE: create configure file in the ~/.curly/config.json
	fullPath := types.CONFIGURE_FILE_PATH + "/" + types.CONFIGURE_FILE_NAME

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

	return nil
}
