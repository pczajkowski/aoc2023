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

func (p *Point) move(maze []string) bool {
	switch maze[p.y][p.x] {
	}
}

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var maze []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		maze = append(maze, line)
	}

	return maze
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

	maze := readInput(file)
	fmt.Println(maze)
}
