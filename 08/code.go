package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directions struct {
	left  string
	right string
}

type Network struct {
	moves string
	paths map[string]Directions
}

func readInput(file *os.File) Network {
	scanner := bufio.NewScanner(file)
	var network Network
	readMoves := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if !readMoves {
				readMoves = true
				continue
			}

			break
		}

		if !readMoves {
			network.moves = line
			network.paths = make(map[string]Directions)
		} else {
			fromParts := strings.Split(line, " = ")
			if len(fromParts) != 2 {
				log.Fatalf("Wrong number of fromParts: %s", line)
			}

			from := fromParts[0]
			parts := strings.Split(fromParts[1], ", ")
			if len(parts) != 2 {
				log.Fatalf("Wrong number of parts: %s", fromParts[1])
			}

			var directions Directions
			directions.left = parts[0]
			directions.right = parts[1]
			network.paths[from] = directions
		}
	}

	return network
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

	network := readInput(file)
	fmt.Println(network)
}
