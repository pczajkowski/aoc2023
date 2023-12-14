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

func tiltNorth(platform [][]byte, y int, x int, height int, width int) {
	for {
		prevY := y - 1
		if prevY < 0 || platform[prevY][x] == '#' || platform[prevY][x] == 'O' {
			break
		}

		platform[y][x] = '.'
		platform[prevY][x] = 'O'
		y--
	}
}

func tiltSouth(platform [][]byte, y int, x int, height int, width int) {
	for {
		nextY := y + 1
		if nextY >= height || platform[nextY][x] == '#' || platform[nextY][x] == 'O' {
			break
		}

		platform[y][x] = '.'
		platform[nextY][x] = 'O'
		y++
	}
}

func tiltEast(platform [][]byte, y int, x int, height int, width int) {
	for {
		nextX := x + 1
		if nextX >= width || platform[y][nextX] == '#' || platform[y][nextX] == 'O' {
			break
		}

		platform[y][x] = '.'
		platform[y][nextX] = 'O'
		x++
	}
}

func tiltWest(platform [][]byte, y int, x int, height int, width int) {
	for {
		nextX := x - 1
		if nextX < 0 || platform[y][nextX] == '#' || platform[y][nextX] == 'O' {
			break
		}

		platform[y][x] = '.'
		platform[y][nextX] = 'O'
		x--
	}
}

func tiltPlatform(platform [][]byte, direction func([][]byte, int, int, int, int), height int, width int) {
	for y := range platform {
		for x := range platform[y] {
			if platform[y][x] == 'O' {
				direction(platform, y, x, height, width)
			}
		}
	}
}

func calculate(platform [][]byte, height int) int {
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

func part1(platform [][]byte) int {
	height := len(platform)
	width := len(platform[0])
	tiltPlatform(platform, tiltNorth, height, width)

	return calculate(platform, height)
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
