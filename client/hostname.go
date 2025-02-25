package main

import "os"

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Error getting hostname"
	}
	return hostname
}
