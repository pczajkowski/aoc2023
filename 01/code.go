package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Line struct {
	full   string
	number int
}

func readInput(file *os.File) []Line {
	scanner := bufio.NewScanner(file)
	var lines []Line
	const delta = 48

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		current := Line{full: line}
		var digit byte
		haveFirst := false
		for i := range line {
			if line[i] >= delta && line[i] <= 57 {
				digit = line[i]
				if !haveFirst {
					current.number = int(digit-delta) * 10
					haveFirst = true
				}
			}
		}

		current.number += int(digit - delta)
		lines = append(lines, current)

	}

	return lines
}

func part1(lines []Line) int {
	var sum int
	for i := range lines {
		sum += lines[i].number
	}

	return sum
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

	lines := readInput(file)
	fmt.Println("Part1:", part1(lines))
}
