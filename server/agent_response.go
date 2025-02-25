package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ResponseReply struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Output string `json:"output"`
}

func agent_response(w http.ResponseWriter, r *http.Request) {
	// Extract agent_id from URL
	agent_id := strings.TrimPrefix(r.URL.Path, "/response/")
	if len(agent_id) == 0 {
		http.Error(w, "Agent ID not provided", http.StatusBadRequest)
		return
	}

	// Get the agent from agentManager
	agent, exists := agentManager.GetAgent(agent_id)
	if !exists {
		http.Error(w, "Agent not found", http.StatusNotFound)
		return
	}
	agent.LastSeen = time.Now()

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse JSON into struct
	var response ResponseReply
	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Log the received response
	fmt.Printf("\n%s\n", response.Output)
	fmt.Printf(">> ")
}
