package main

import (
	"os/user"
	"log"
)

func GetCurrentUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.Username, err
}
