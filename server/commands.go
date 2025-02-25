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

// formatTimeAgo converts time duration into a human-readable format
func formatTimeAgo(duration time.Duration) string {
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

func Get_Help() {
	fmt.Println("Commands:")
	fmt.Println("\n  -- Agent Management -- ")
	fmt.Println("  agents 		- List all agents")
	fmt.Println("  rm <agent_id> 	- Removes an agent by ID")
	fmt.Println("  use <agent_id> 	- Use an agent by ID ")

	fmt.Println("\n  -- Agent Commands -- ")
	fmt.Println("  whoami 			- Display the user your agent us running under")
	fmt.Println("  hostname 		- Displays hostname of the system")
	fmt.Println("  ps 				- Display running processes")
	fmt.Println("  netstat 			- Display network connections")
	fmt.Println("  ls 				- List files in the current directory")
	fmt.Println("  cd <directory> 	- Change directory")
	fmt.Println("  cat <file> 		- Display the contents of a file")

	fmt.Println("\n  -- Misc -- ")
	fmt.Println("  help 			- Display this help message")
	fmt.Println("  exit 			- Exit the server")
}
