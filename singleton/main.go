package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	// Scanning instance lines
	fmt.Scanln()
}
