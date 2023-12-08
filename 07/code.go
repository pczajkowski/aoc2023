package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	NumberOfCards = 5
	NumberOfRanks = 13
)

type Hand struct {
	cards string
	bid   int
	ranks [NumberOfCards]int
	score int
}

func (h *Hand) analyzeHand() {
	five := false
	four := false
	three := false
	pairs := 0
	var numberInRank [NumberOfRanks]int

	for i := range h.cards {
		numberInRank[ranks[h.cards[i]]]++
		h.ranks[i] = ranks[h.cards[i]]
	}

	for rank := 0; rank < NumberOfRanks; rank++ {
		switch numberInRank[rank] {
		case 5:
			five = true
		case 4:
			four = true
		case 3:
			three = true
		case 2:
			pairs++
		}
	}

	if five {
		h.score = FiveKind
	} else if four {
		h.score = FourKind
	} else if three && pairs == 1 {
		h.score = FullHouse
	} else if three {
		h.score = ThreeKind
	} else if pairs == 2 {
		h.score = TwoPair
	} else if pairs == 1 {
		h.score = OnePair
	} else {
		h.score = HighCard
	}
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

		hand.analyzeHand()
		hands = append(hands, hand)
	}

	return hands
}

var ranks map[byte]int

func init() {
	ranks = make(map[byte]int)
	ranks['2'] = 0
	ranks['3'] = 1
	ranks['4'] = 2
	ranks['5'] = 3
	ranks['6'] = 4
	ranks['7'] = 5
	ranks['8'] = 6
	ranks['9'] = 7
	ranks['T'] = 8
	ranks['J'] = 9
	ranks['Q'] = 10
	ranks['K'] = 11
	ranks['A'] = 12
}

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

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
