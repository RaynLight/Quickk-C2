package main

import (
	"Deanscup/server/includes"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	commands := map[string]func([]string){
		"exit":   func(args []string) { includes.Exit_Server() },
		"quic":   func(args []string) { go includes.StartQuic(args) },
		"agents": func(args []string) { go includes.List_agents_cli(includes.Manager.ListAgents()) },
		"rm":     func(args []string) { includes.Manager.RemoveAgent(args[1]) },

		"help":       func(args []string) { includes.Get_Help() },
		"addtask":    func(args []string) { includes.HandleAddTask(args, (*includes.Manager).ListAgents()) },
		"use":        func(args []string) { includes.Use(args) },
		"background": func(args []string) { includes.Background() },
	}

	for {
		includes.PrintCursor()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		words := strings.Fields(input)

		// checking if user actually inputted something
		if len(words) == 0 {
			continue
		}

		if includes.GetCurrentAgent() != "" && words[0] != "use" && words[0] != "exit" && words[0] != "background" {
			task := strings.Join(words, " ")
			includes.HandleAddTask([]string{"addtask", includes.GetCurrentAgent(), task}, (*includes.Manager).ListAgents())
			continue
		}
		if cmd, exists := commands[strings.ToLower(words[0])]; exists {
			cmd(words)
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
