package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Hand struct {
	cards string
	bid   int
}

func readInput(file *os.File) []Hand {
	scanner := bufio.NewScanner(file)
	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var hand Hand
		n, err := fmt.Sscanf(line, "%s %d", &hand.cards, &hand.bid)
		if n != 2 || err != nil {
			log.Fatalf("Failed to read hand: %s\n%s", line, err)
		}

		hands = append(hands, hand)
	}

	return hands
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

	hands := readInput(file)
	fmt.Println(hands)
}
