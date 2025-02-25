package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net"
	"net/http"
	"strings"
	"time"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func first_checkin(w http.ResponseWriter, r *http.Request) *Agent {
	ID := GenerateUUID()
	response := map[string]string{
		"agent_id": ID,
	}

	agent := &Agent{
		ID:       ID,
		Hostname: "",
		OS:       "",
		IP:       net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]),
		LastSeen: time.Now(),
	}

	agentManager.AddAgent(agent)

	json.NewEncoder(w).Encode(response)

	fmt.Printf("[+] New Agent Registered: %s, %s\n", r.RemoteAddr, agent.ID)

	return agent
}

func agent_checkin(w http.ResponseWriter, r *http.Request) {
	agent_id := r.URL.Path[len("/checkin/"):]
	if len(agent_id) == 0 {
		http.Error(w, "Agent ID not provided", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s", agent_id)
}
