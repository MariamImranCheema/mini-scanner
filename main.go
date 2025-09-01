package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to this vulnerability scanner!")

	if len(os.Args) < 2 {
		fmt.Println("Usage: run main.go <docker-image>")
		return
	}

	image := os.Args[1]
	fmt.Printf("Scanning Docker image: %s\n", image)
}
