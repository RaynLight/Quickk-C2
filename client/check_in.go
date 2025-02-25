package main

import (
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"io"
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

	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("[-] Request failed: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[-] Failed to read response body: %v\n", err)
		return false
	}
	defer resp.Body.Close() // Always close response body

	fmt.Printf("[+] Server Response: %s\n", string(body))

	return true
}
