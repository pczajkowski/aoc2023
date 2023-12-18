package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Dig struct {
	direction string
	length    int
	color     string
}

func readInput(file *os.File) []Dig {
	scanner := bufio.NewScanner(file)
	var plan []Dig

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var current Dig
		n, err := fmt.Sscanf(line, "%s %d %s", &current.direction, &current.length, &current.color)
		if n != 3 || err != nil {
			log.Fatalf("Wrong input: %s\n%s", line, err)
		}

		current.color = strings.Trim(current.color, "()")
		plan = append(plan, current)
	}

	return plan
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	plan := readInput(file)
	fmt.Println(plan)
}
