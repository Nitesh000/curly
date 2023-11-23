package utils

import "os/exec"

func IsCurlInstalled() bool {
	_, err := exec.Command("curl", "--version").Output()
	if err != nil {
		return false
	}
	return true
}
