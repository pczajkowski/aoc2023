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
	fmt.Println(matrix)
}
