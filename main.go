package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		user, err := GetCurrentUser()
		if err != nil {
			log.Fatal(err)
		}

		hostname, err := GetHostname()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s@%s: ", user, hostname)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
		}

		if err = ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}





















