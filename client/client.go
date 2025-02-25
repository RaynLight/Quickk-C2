package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	ip := "127.0.0.1"
	port := "4443"
	url := fmt.Sprintf("https://%s:%s/", ip, port)

	// Calls first checkin api to get the agent ID
	ID := FirstCheckin(ip, port)

	checkIn(ip, port, ID)
	fmt.Println("My Id is", ID)
	client := &http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				// Ignore self-signed certificate
				InsecureSkipVerify: true,
			},
		},
	}

	fmt.Println("[+] Sending HTTP/3 request to:", url)

	for {
		resp, err := client.Get(url)
		if err != nil {
			log.Fatalf("[-] Request failed: %v", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("[-] Failed to read response: %v", err)
		}
		fmt.Println("[+] Response:", string(body))

		time.Sleep(2 * time.Second)
	}
}
