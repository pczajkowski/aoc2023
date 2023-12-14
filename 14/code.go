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

func tiltNorth(platform [][]byte) [][]byte {
	newPlatform := [][]byte{platform[0]}
	height := len(platform)
	for i := 1; i < height; i++ {
		newPlatform = append(newPlatform, platform[i])
		for x := range newPlatform[i] {
			if newPlatform[i][x] == 'O' {
				y := i - 1
				for {
					if y < 0 || newPlatform[y][x] == '#' || newPlatform[y][x] == 'O' {
						break
					}

					newPlatform[y+1][x] = '.'
					newPlatform[y][x] = 'O'
					y--
				}
			}
		}
	}

	return newPlatform
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
	fmt.Println(tiltNorth(platform))
}
