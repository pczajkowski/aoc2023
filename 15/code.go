package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(data), ",")
	if len(parts) == 0 {
		log.Fatal("Bad input!")
	}

	for i := range parts {
		parts[i] = strings.TrimRight(parts[i], "\n")
	}

	return parts
}

const (
	Max = 256
)

func hash(text string) int {
	var current int
	for i := range text {
		current += int(text[i])
		current = current * 17 % Max
	}

	return current
}

func part1(steps []string) int {
	var result int
	for i := range steps {
		result += hash(steps[i])
	}

	return result
}

type Lens struct {
	label string
	power int
}

func getBoxes(steps []string) [][]Lens {
	lenses := make([][]Lens, 256)
	for i := range steps {
		if strings.Contains(steps[i], "=") {
			parts := strings.Split(steps[i], "=")
			if len(parts) != 2 {
				log.Fatalf("Problem reading step %s", steps[i])
			}

			lens := Lens{label: parts[0]}
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("Problem converting number %s: %s", parts[1], err)
			}

			lens.power = n
			fmt.Println(lens)
		} else {
			label := strings.TrimRight(steps[i], "-")
			fmt.Println(label)
		}
	}

	return lenses
}

func part2(steps []string) int {
	var result int
	getBoxes(steps)
	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	steps := readInput(os.Args[1])
	fmt.Println("Part1:", part1(steps))
	fmt.Println("Part2:", part2(steps))
}
