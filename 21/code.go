package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x  int
	steps int
}

func readInput(file *os.File) (Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var board [][]byte
	var start Point
	var index int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []byte
		for i := range line {
			if line[i] == 'S' {
				start.x = i
				start.y = index
			}

			row = append(row, line[i])
		}

		board = append(board, row)
		index++
	}

	return start, board
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

	start, board := readInput(file)
	fmt.Println(start, board)
}
