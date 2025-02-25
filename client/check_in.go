package main

import (
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"net/http"
)

func checkIn(ip string, port string, agent_id string) bool {
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

	// resp, err := client.Get(url)
	client.Get(url)

	return true
}
