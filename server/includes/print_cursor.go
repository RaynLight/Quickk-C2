package includes

import (
	"fmt"
)

func PrintCursor() {
	agent := GetCurrentAgent()
	if agent != "" {
		if len(agent) > 8 {
			agent = agent[:8]
		}
		fmt.Printf(" [%s] ", agent)
	} else {
		fmt.Printf(" [None] ")
	}
	fmt.Printf(">> ")
}
