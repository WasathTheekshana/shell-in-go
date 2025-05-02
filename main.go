package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"os/exec"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the new line character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to get the command and arguments
	args := strings.Split(input, " ")

	switch args[0] {
	case "exit":
		os.Exit(0)		

	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: Missing Arguments")
		}
		return os.Chdir(args[1])
	}

	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command
	return cmd.Run()
}
