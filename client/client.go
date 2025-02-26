package main

import (
	"Deanscup/client/includes"
	"fmt"
	"time"
)

func main() {
	ip := "127.0.0.1"
	port := "4443"
	url := fmt.Sprintf("https://%s:%s/", ip, port)

	ID := includes.FirstCheckin(ip, port)

	fmt.Println("My Id is", ID)

	fmt.Println("[+] Sending HTTP/3 request to:", url)

	for {
		includes.CheckIn(ip, port, ID)

		time.Sleep(250 * time.Millisecond)
	}

}
