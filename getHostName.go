package main

import (
	"strings"
	"os/exec"
)

func GetHostname() (string, error) {
	hostname, err := exec.Command("hostname").Output() 
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(hostname)), err
}
