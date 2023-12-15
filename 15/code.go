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

func lensFromString(text string) Lens {
	parts := strings.Split(text, "=")
	if len(parts) != 2 {
		log.Fatalf("Problem reading step %s", text)
	}

	lens := Lens{label: parts[0]}
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Problem converting number %s: %s", parts[1], err)
	}

	lens.power = n

	return lens
}

func addLens(box []Lens, lens Lens) []Lens {
	found := false
	for i := range box {
		if box[i].label == lens.label {
			box[i] = lens
			found = true
			break
		}
	}

	if !found {
		box = append(box, lens)
	}

	return box
}

func getBoxes(steps []string) [][]Lens {
	lenses := make([][]Lens, 256)
	for i := range steps {
		if strings.Contains(steps[i], "=") {
			lens := lensFromString(steps[i])
			boxIndex := hash(lens.label)
			lenses[boxIndex] = addLens(lenses[boxIndex], lens)
		} else {
			label := strings.TrimRight(steps[i], "-")
			boxIndex := hash(label)
			for i := range lenses[boxIndex] {
				if lenses[boxIndex][i].label == label {
					lenses[boxIndex] = append(lenses[boxIndex][:i], lenses[boxIndex][i+1:]...)
					break
				}
			}
		}
	}

	return lenses
}

func part2(steps []string) int {
	var result int
	lenses := getBoxes(steps)
	fmt.Println(lenses)
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
