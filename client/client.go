package main

import (
	"fmt"
)

func main() {
	ip := "127.0.0.1"
	port := "4443"
	url := fmt.Sprintf("https://%s:%s/", ip, port)

	ID := FirstCheckin(ip, port)

	fmt.Println("My Id is", ID)

	fmt.Println("[+] Sending HTTP/3 request to:", url)

	for {
		checkIn(ip, port, ID)
		// time.Sleep(2 * time.Second)
	}
}
