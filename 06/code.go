package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const Time = "Time:      "
const Distance = "Distance:  "

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var matrix [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if strings.HasPrefix(line, Time) {
			line = strings.TrimLeft(line, Time)
		} else {
			line = strings.TrimLeft(line, Distance)
		}

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			log.Fatalf("Wrong input: %s", line)
		}

		var numbers []int
		for i := range parts {
			if parts[i] == "" {
				continue
			}

			var number int
			n, err := fmt.Sscanf(parts[i], "%d", &number)
			if n != 1 || err != nil {
				log.Fatalf("Failed to read number: %s\n%s", parts[i], err)
			}

			numbers = append(numbers, number)
		}

		matrix = append(matrix, numbers)
		if len(matrix) == 2 {
			break
		}
	}

	return matrix
}

func part1(matrix [][]int) int {
	result := 1

	for i := range matrix[0] {
		min := matrix[1][i] / matrix[0][i]
		for {
			if min*(matrix[0][i]-min) > matrix[1][i] {
				break
			}

			min++
		}

		result *= matrix[0][i] - min - min + 1
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

	matrix := readInput(file)
	fmt.Println(part1(matrix))
}
