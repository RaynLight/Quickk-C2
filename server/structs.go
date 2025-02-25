package main

import (
	"net"
	"sync"
	"time"
)

type C2Server struct {
	Address string            `json:"address"`
	Agents  map[string]*Agent `json:"agents"`
	mu      sync.Mutex        `json:"-"`
}

// Agent represents a connected C2 client
type Agent struct {
	ID       string    `json:"id"`
	Hostname string    `json:"hostname"`
	OS       string    `json:"os"`
	IP       net.IP    `json:"ip"`
	LastSeen time.Time `json:"last_seen"`
}

type AgentManager struct {
	agents map[string]*Agent
	mu     sync.Mutex
}
