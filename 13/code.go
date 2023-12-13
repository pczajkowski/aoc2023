package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pattern struct {
	note       [][]byte
	vertical   int
	horizontal int
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
			continue
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

func checkVertical(index int, note [][]byte, width int, canMissOne bool) bool {
	left := 0
	prev := index - 1
	right := index + prev
	if right >= width {
		right = width - 1
	}

	for {
		if prev < left || index > right {
			break
		}

		for i := range note {
			if note[i][prev] != note[i][index] {
				if canMissOne {
					canMissOne = false
					continue
				}

				return false
			}
		}

		prev--
		index++
	}

	return true
}

func checkHorizontal(index int, note [][]byte, height int, width int, canMissOne bool) bool {
	up := 0
	prev := index - 1

	down := index + prev
	if down >= height {
		down = height - 1
	}

	for {
		if prev < up || index > down {
			break
		}

		for i := range note[index] {
			if note[index][i] != note[prev][i] {
				if canMissOne {
					canMissOne = false
					continue
				}

				return false
			}
		}

		prev--
		index++
	}

	return true
}

func parts(patterns []Pattern, canMissOne bool) int {
	var result int
	for i := range patterns {
		width := len(patterns[i].note[0])
		height := len(patterns[i].note)

		vertical := 0
		for index := 1; index < width; index++ {
			if index == patterns[i].vertical {
				continue
			}

			if checkVertical(index, patterns[i].note, width, canMissOne) {
				vertical = index
				patterns[i].vertical = index
				break
			}
		}

		horizontal := 0
		for index := 1; index < height; index++ {
			if index == patterns[i].horizontal {
				continue
			}

			if checkHorizontal(index, patterns[i].note, height, width, canMissOne) {
				horizontal = index
				patterns[i].horizontal = index
				break
			}
		}

		if horizontal > 0 {
			result += horizontal * 100
		} else if vertical > 0 {
			result += vertical
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

	patterns := readInput(file)
	fmt.Println("Part1:", parts(patterns, false))
	fmt.Println("Part2:", parts(patterns, true))
}
