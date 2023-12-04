package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	id      int
	winning []int
	owned   []int
	wins    int
}

func readNumbers(part string) []int {
	var numbers []int
	numberParts := strings.Split(part, " ")
	if len(numberParts) == 0 {
		log.Fatalf("Can't split numbers: %s", part)
	}

	for i := range numberParts {
		if numberParts[i] == "" {
			continue
		}

		var number int
		n, err := fmt.Sscanf(numberParts[i], "%d", &number)
		if n != 1 || err != nil {
			log.Fatalf("Can't read number: %s\n%s", numberParts[i], err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func readInput(file *os.File) []Card {
	scanner := bufio.NewScanner(file)
	var cards []Card

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var current Card
		n, err := fmt.Sscanf(line, "Card %d:", &current.id)
		if n != 1 || err != nil {
			log.Fatalf("Failed to read card id: %s\n%s", line, err)
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Can't split card: %s", line)
		}

		numberParts := strings.Split(parts[1], "|")
		if len(parts) != 2 {
			log.Fatalf("Can't split tables: %s", line)
		}

		current.winning = append(current.winning, readNumbers(numberParts[0])...)
		current.owned = append(current.owned, readNumbers(numberParts[1])...)
		cards = append(cards, current)
	}

	return cards
}

func isInArray(number int, array []int) bool {
	for i := range array {
		if array[i] == number {
			return true
		}
	}

	return false
}

func pow(x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= 2
	}

	return result
}

func part1(cards []Card) int {
	var result int
	for i := range cards {
		var count int
		for j := range cards[i].owned {
			if isInArray(cards[i].owned[j], cards[i].winning) {
				count++
			}
		}

		if count > 0 {
			result += pow(count - 1)
		}

		cards[i].wins = count
	}

	return result
}

func part2(cards []Card) int {
	var pool []int
	for i := range cards {
		if cards[i].wins > 0 {
			pool = append(pool, i)
		}
	}

	var result int
	for {
		if len(pool) == 0 {
			break
		}

		current := pool[0]
		pool = pool[:1]
		result++

		for i := 0; i < cards[current].wins; i++ {
			current++
			pool = append(pool, current)
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

	cards := readInput(file)
	fmt.Println("Part1:", part1(cards))
	fmt.Println("Part2:", part2(cards))
}
