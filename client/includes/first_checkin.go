package includes

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"io/ioutil"
	"log"
	"net/http"
)

// json struct to recive data
type CheckInResponse struct {
	AgentID string `json:"agent_id"`
}

func FirstCheckin(ip string, port string) string {
	// putting the url together lol
	url := fmt.Sprintf("https://%s:%s/first_checkin", ip, port)

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
		log.Fatalf("[-] Request failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[-] Failed to read response body: %v", err)
	}

	// reading in the json response for the client ID
	var checkInResp CheckInResponse
	err = json.Unmarshal(body, &checkInResp)
	if err != nil {
		log.Println("[-] Failed to parse JSON:", err)
		return ""
	}

	fmt.Println("[+] Agent ID:", checkInResp.AgentID)
	return checkInResp.AgentID
}
