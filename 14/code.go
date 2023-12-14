package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var platform [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []byte
		for i := range line {
			row = append(row, line[i])
		}

		platform = append(platform, row)
	}

	return platform
}

func tiltNorth(platform [][]byte, y int, x int, height int, width int) bool {
	change := false
	for {
		prevY := y - 1
		if prevY < 0 || platform[prevY][x] == '#' || platform[prevY][x] == 'O' {
			break
		}

		platform[y][x] = '.'
		platform[prevY][x] = 'O'
		change = true
		y--
	}

	return change
}

func tiltPlatform(platform [][]byte, direction func([][]byte, int, int, int, int) bool, height int, width int) bool {
	change := false
	for y := range platform {
		for x := range platform[y] {
			if platform[y][x] == 'O' {
				if direction(platform, y, x, height, width) {
					change = true
				}
			}
		}
	}

	return change
}

func part1(platform [][]byte) int {
	height := len(platform)
	width := len(platform[0])
	tiltPlatform(platform, tiltNorth, height, width)
	var result int
	for y := range platform {
		for x := range platform[y] {
			if platform[y][x] == 'O' {
				result += height - y
			}
		}
	}

	return result
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

	platform := readInput(file)
	fmt.Println("Part1:", part1(platform))
}
