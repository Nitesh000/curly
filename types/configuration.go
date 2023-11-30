package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func NewConfig() *RequestConfigurations {
	// NOTE: Read from the configuration file in the ~/.curly/curly.json
	var home string
	home = os.Getenv("HOME")

	// NOTE: check if the file exists or not
	if _, err := os.Stat(home + "/" + CONFIGURE_FILE_PATH + "/" + CONFIGURE_FILE_NAME); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		fmt.Println("To create this file do curly config -c")
		os.Exit(1)
	}

	file, err := os.Open(home + "/" + CONFIGURE_FILE_PATH + "/" + CONFIGURE_FILE_NAME)
	if err != nil {
		return nil
	}

	// NOTE: read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading the config file")
		os.Exit(1)
	}

	defer file.Close()

	var config RequestConfigurations
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error parsing the config file: %s", err)
		os.Exit(1)
	}

	return &RequestConfigurations{
		TotalTime:             config.TotalTime,
		NameLookupTime:        config.NameLookupTime,
		ConnectTime:           config.ConnectTime,
		AppConnectTime:        config.AppConnectTime,
		PreTransferTime:       config.PreTransferTime,
		RedirectTime:          config.RedirectTime,
		StartTransferTime:     config.StartTransferTime,
		RedirectCount:         config.RedirectCount,
		SizeUpload:            config.SizeUpload,
		SizeDownload:          config.SizeDownload,
		SpeedDownload:         config.SpeedDownload,
		SpeedUpload:           config.SpeedUpload,
		HeaderSize:            config.HeaderSize,
		RequestSize:           config.RequestSize,
		SslVerifyResult:       config.SslVerifyResult,
		FiletimeConnect:       config.FiletimeConnect,
		FiletimePretransfer:   config.FiletimePretransfer,
		FiletimeStarttransfer: config.FiletimeStarttransfer,
		FiletimeAppconnect:    config.FiletimeAppconnect,
		RemoteIp:              config.RemoteIp,
		RemotePort:            config.RemotePort,
		LocalIp:               config.LocalIp,
		LocalPort:             config.LocalPort,
		RedirectUrl:           config.RedirectUrl,
		PrimaryIp:             config.PrimaryIp,
		PrimaryPort:           config.PrimaryPort,
		CurlError:             config.CurlError,
		HttpCode:              config.HttpCode,
		EffectiveUrl:          config.EffectiveUrl,
		ContentType:           config.ContentType,
		NumConnects:           config.NumConnects,
		PrimaryIp6:            config.PrimaryIp6,
		LocalIp6:              config.LocalIp6,
	}
}

func (c *RequestConfigurations) CreateConfigureString() string {
	if c == nil {
		return ""
	}

	var configureString string

	configureString += " -w '"

	if c.TotalTime {
		configureString += "time_total: %{time_total}\n"
	}

	if c.NameLookupTime {
		configureString += "time_namelookup: %{time_namelookup}\n"
	}

	if c.ConnectTime {
		configureString += "time_connect: %{time_connect}\n"
	}

	if c.AppConnectTime {
		configureString += "time_appconnect: %{time_appconnect}\n"
	}

	if c.PreTransferTime {
		configureString += "time_pretransfer: %{time_pretransfer}\n"
	}

	if c.RedirectTime {
		configureString += "time_redirect: %{time_redirect}\n"
	}

	if c.StartTransferTime {
		configureString += "time_starttransfer: %{time_starttransfer}\n"
	}

	if c.RedirectCount {
		configureString += "redirect_count: %{redirect_count}\n"
	}

	if c.SizeUpload {
		configureString += "size_upload: %{size_upload}\n"
	}

	if c.SizeDownload {
		configureString += "size_download: %{size_download}\n"
	}

	if c.SpeedDownload {
		configureString += "speed_download: %{speed_download}\n"
	}

	if c.SpeedUpload {
		configureString += "speed_upload: %{speed_upload}\n"
	}

	if c.HeaderSize {
		configureString += "header_size: %{header_size}\n"
	}

	if c.RequestSize {
		configureString += "request_size: %{request_size}\n"
	}

	if c.SslVerifyResult {
		configureString += "ssl_verify_result: %{ssl_verify_result}\n"
	}

	if c.FiletimeConnect {
		configureString += "filetime_connect: %{filetime_connect}\n"
	}

	if c.FiletimePretransfer {
		configureString += "filetime_pretransfer: %{filetime_pretransfer}\n"
	}

	if c.FiletimeStarttransfer {
		configureString += "filetime_starttransfer: %{filetime_starttransfer}\n"
	}

	if c.FiletimeAppconnect {
		configureString += "filetime_appconnect: %{filetime_appconnect}\n"
	}

	if c.RemoteIp {
		configureString += "remote_ip: %{remote_ip}\n"
	}

	if c.RemotePort {
		configureString += "remote_port: %{remote_port}\n"
	}

	if c.LocalIp {
		configureString += "local_ip: %{local_ip}\n"
	}

	if c.LocalPort {
		configureString += "local_port: %{local_port}\n"
	}

	if c.RedirectUrl {
		configureString += "redirect_url: %{redirect_url}\n"
	}

	if c.PrimaryIp {
		configureString += "primary_ip: %{primary_ip}\n"
	}

	if c.PrimaryPort {
		configureString += "primary_port: %{primary_port}\n"
	}

	if c.CurlError {
		configureString += "curl_error: %{curl_error}\n"
	}

	if c.HttpCode {
		configureString += "http_code: %{http_code}\n"
	}

	if c.EffectiveUrl {
		configureString += "effective_url: %{effective_url}\n"
	}

	if c.ContentType {
		configureString += "content_type: %{content_type}\n"
	}

	if c.NumConnects {
		configureString += "num_connects: %{num_connects}\n"
	}

	if c.PrimaryIp6 {
		configureString += "primary_ip6: %{primary_ip6}\n"
	}

	if c.LocalIp6 {
		configureString += "local_ip6: %{local_ip6}\n"
	}

	configureString += "' -o -"

	return configureString
}
