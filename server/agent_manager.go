package main

import "fmt"

// initializes the agent manager
func NewAgentManager() *AgentManager {
	return &AgentManager{
		agents: make(map[string]*Agent),
	}
}

// adds a new agent
func (m *AgentManager) AddAgent(agent *Agent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.agents[agent.ID] = agent
	fmt.Printf("[+] Agent Registered: %+v\n", agent)
}

// GetAgent retrieves an agent by ID
func (m *AgentManager) GetAgent(id string) (*Agent, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	agent, exists := m.agents[id]
	return agent, exists
}

// returns a list of all agents
func (m *AgentManager) ListAgents() []*Agent {
	m.mu.Lock()
	defer m.mu.Unlock()
	agentList := make([]*Agent, 0, len(m.agents))
	for _, agent := range m.agents {
		agentList = append(agentList, agent)
	}
	return agentList
}

// Deletes an agent
func (m *AgentManager) RemoveAgent(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.agents, id)
	fmt.Printf("[-] Agent %s removed.\n", id)
}
