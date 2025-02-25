package main

import (
	"Deanscup/server"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	commands := map[string]func([]string){
		"exit": func(args []string) { Server.Exit_Server() },
		"quic": func(args []string) { go Server.StartQuic(args) },
	}

	for {
		fmt.Printf(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		words := strings.Fields(input)

		// checking if user actually inputted something
		if len(words) == 0 {
			continue
		}

		if cmd, exists := commands[strings.ToLower(words[0])]; exists {
			cmd(words)
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
