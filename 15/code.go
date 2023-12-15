package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(data), ",")
	if len(parts) == 0 {
		log.Fatal("Bad input!")
	}

	for i := range parts {
		parts[i] = strings.TrimRight(parts[i], "\n")
	}

	return parts
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	steps := readInput(os.Args[1])
	fmt.Println(steps)
}
