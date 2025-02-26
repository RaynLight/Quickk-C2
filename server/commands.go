package main

import (
	"fmt"
	"os"
	"time"
)

func Exit_Server() {
	fmt.Printf("Exiting...")
	os.Exit(0)
}

func list_agents_cli(agents []*Agent) {
	if len(agents) == 0 {
		fmt.Println("[-] No agents registered.")
		return
	}

	fmt.Println("[+] Registered Agents:")
	now := time.Now()

	for _, agent := range agents {
		timeAgo := now.Sub(agent.LastSeen)
		lastSeenStr := formatTimeAgo(timeAgo)

		fmt.Printf("ID: %s | IP: %s | Last Seen: %s\n", agent.ID, agent.IP, lastSeenStr)
	}
}

func formatTimeAgo(duration time.Duration) string { // converts time duration into a human-readable format
	seconds := int(duration.Seconds())

	switch {
	case seconds < 60:
		return "<1 min ago"
	case seconds < 3600:
		return fmt.Sprintf("%d minutes ago", seconds/60)
	case seconds < 86400:
		return fmt.Sprintf("%d hours ago", seconds/3600)
	default:
		return fmt.Sprintf("%d days ago", seconds/86400)
	}
}

func HandleAddTask(command []string, agents []*Agent) {
	if len(command) < 3 {
		fmt.Println("[-] Usage: addtask <agent_id> <task>")
		return
	}

	agentID := command[1]
	task := ""
	if len(command) > 3 {
		task = fmt.Sprintf("%s", command[2:])
	} else {
		task = command[2]
	}
	addTaskToAgent(agentID, task, agents)
}

func Get_Help() { // Help command
	fmt.Println("Commands:")
	fmt.Println("\n  -- Agent Management -- ")
	fmt.Println("  agents 		- List all agents")              // working
	fmt.Println("  rm <agent_id> 	- Removes an agent by ID") // working
	fmt.Println("  use <agent_id> 	- Use an agent by ID ")
	fmt.Println("  background       - Background your agent")
	fmt.Println("  addtask <agent_id> <task> - Add a task to an agent") // working

	fmt.Println("\n  -- Agent Commands -- ")
	fmt.Println("  whoami 		- Display the user your agent us running under")
	fmt.Println("  hostname 		- Displays hostname of the system")
	fmt.Println("  ps 			- Display running processes")
	fmt.Println("  netstat 		- Display network connections")
	fmt.Println("  ls 			- List files in the current directory")
	fmt.Println("  cd <directory> 	- Change directory")
	fmt.Println("  cat <file> 		- Display the contents of a file")

	fmt.Println("\n  -- Misc -- ")
	fmt.Println("  help 			- Display this help message")
	fmt.Println("  exit 			- Exit the server")
}
