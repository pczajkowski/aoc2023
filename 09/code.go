package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var matrix [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			log.Fatalf("Wrong input: %s", line)
		}

		var numbers []int
		for i := range parts {
			var number int
			n, err := fmt.Sscanf(parts[i], "%d", &number)
			if n != 1 || err != nil {
				log.Fatalf("Failed to read number: %s\n%s", parts[i], err)
			}

			numbers = append(numbers, number)
		}

		matrix = append(matrix, numbers)
	}

	return matrix
}

func allZeros(numbers []int) bool {
	for i := range numbers {
		if numbers[i] != 0 {
			return false
		}
	}

	return true
}

func getNext(numbers []int) int {
	var matrix [][]int
	matrix = append(matrix, numbers)
	level := 0
	for {
		if level >= len(matrix) {
			break
		}

		var newLevel []int
		for i := 1; i < len(matrix[level]); i++ {
			newLevel = append(newLevel, matrix[level][i]-matrix[level][i-1])
		}

		matrix = append(matrix, newLevel)
		if allZeros(newLevel) {
			break
		}

		level++
	}

	for i := len(matrix) - 1; i > 0; i-- {
		matrix[i-1] = append(matrix[i-1], matrix[i][len(matrix[i])-1]+matrix[i-1][len(matrix[i-1])-1])
	}

	return matrix[0][len(matrix[0])-1]
}

func part1(matrix [][]int) int {
	var result int
	for i := range matrix {
		next := getNext(matrix[i])
		matrix[i] = append(matrix[i], next)
		result += next
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
	fmt.Println("Part1:", part1(matrix))
}
