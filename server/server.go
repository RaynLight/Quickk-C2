package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	commands := map[string]func([]string){
		"exit":   func(args []string) { Exit_Server() },
		"quic":   func(args []string) { go StartQuic(args) },
		"agents": func(args []string) { go list_agents_cli(agentManager.ListAgents()) },
		"rm":     func(args []string) { agentManager.RemoveAgent(args[1]) },

		"help": func(args []string) { Get_Help() },
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
