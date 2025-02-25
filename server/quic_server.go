package Server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/quic-go/quic-go/http3"
)

func Exit_Server() {
	fmt.Printf("Exiting...")
	os.Exit(0)
}

func StartQuic(args []string) {

	port := 4443
	if len(args) > 1 {
		if p, err := strconv.Atoi(args[1]); err == nil {
			port = p
		} else {
			fmt.Println("Invalid port number, using default 4443.")
		}
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.RemoteAddr)
		fmt.Printf("[+] Connections from: %s\n", r.RemoteAddr)
	})

	fmt.Printf("[+] Starting QUIC server on port %d\n", port)

	err := http3.ListenAndServeQUIC(fmt.Sprintf(":%d", port), "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}
}
