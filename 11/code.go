package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	y, x      int
	direction int
}

type Universe struct {
	galaxies        []Point
	occupiedRows    map[int]bool
	occupiedColumns map[int]bool
}

func readInput(file *os.File) Universe {
	scanner := bufio.NewScanner(file)
	var universe Universe
	universe.occupiedRows = make(map[int]bool)
	universe.occupiedColumns = make(map[int]bool)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if strings.Contains(line, "#") {
			universe.occupiedRows[index] = true
			for i := range line {
				if line[i] == '#' {
					universe.galaxies = append(universe.galaxies, Point{y: index, x: i})
					universe.occupiedColumns[i] = true
				}
			}
		}

		index++
	}

	return universe
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

	universe := readInput(file)
	fmt.Println(universe)
}
