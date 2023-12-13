package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x int
}

type Pattern struct {
	note   [][]byte
	mirror Point
}

func readInput(file *os.File) []Pattern {
	scanner := bufio.NewScanner(file)
	var patterns []Pattern

	var current Pattern
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patterns = append(patterns, current)
			current = Pattern{}
		}

		var row []byte
		for i := range line {
			row = append(row, line[i])
		}

		current.note = append(current.note, row)
	}
	patterns = append(patterns, current)

	return patterns
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

	patterns := readInput(file)
	fmt.Println(patterns)
}
