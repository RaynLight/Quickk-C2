package main

import (
	"fmt"
)

func PrintCursor() {
	if GetCurrentAgent() != "" {
		fmt.Printf(" [%s] ", GetCurrentAgent())
	} else {
		fmt.Printf(" [None] ")
	}
	fmt.Printf(">> ")
}
