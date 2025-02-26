package includes

import "fmt"

// Global variable to track selected agent
var current_agent string

func Use(args []string) {
	if len(args) < 2 {
		fmt.Println("[-] Usage: use <agent_id>")
		return
	}

	agent_id := args[1]
	if _, exists := Manager.GetAgent(agent_id); exists {
		current_agent = agent_id
		fmt.Printf("[+] Now using agent %s\n", agent_id)
	} else {
		fmt.Printf("[-] Agent %s not found.\n", agent_id)
	}
}

func GetCurrentAgent() string {
	return current_agent
}
