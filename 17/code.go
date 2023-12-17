package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const Diff = 48

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var board [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []int
		for i := range line {
			row = append(row, int(line[i]-Diff))
		}

		board = append(board, row)
	}

	return board
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

	board := readInput(file)
	fmt.Println(board)
}
