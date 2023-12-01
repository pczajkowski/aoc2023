package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Line struct {
	full   string
	number int
}

const delta = 48

func readInput(file *os.File) []Line {
	scanner := bufio.NewScanner(file)
	var lines []Line

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		current := Line{full: line}
		var digit byte
		haveFirst := false
		for i := range line {
			if line[i] >= delta && line[i] <= 57 {
				digit = line[i]
				if !haveFirst {
					current.number = int(digit-delta) * 10
					haveFirst = true
				}
			}
		}

		current.number += int(digit - delta)
		lines = append(lines, current)

	}

	return lines
}

func part1(lines []Line) int {
	var sum int
	for i := range lines {
		sum += lines[i].number
	}

	return sum
}

var digits map[string]byte

func init() {
	digits = make(map[string]byte)
	digits["one"] = '1'
	digits["two"] = '2'
	digits["three"] = '3'
	digits["four"] = '4'
	digits["five"] = '5'
	digits["six"] = '6'
	digits["seven"] = '7'
	digits["eight"] = '8'
	digits["nine"] = '9'
}

func part2(lines []Line) int {
	var sum int
	for i := range lines {
		line := lines[i].full
		var digit byte
		var number int
		index := 0
		end := len(line)
		haveFirst := false
		for {
			if index >= end {
				break
			}

			if line[index] >= delta && line[index] <= 57 {
				digit = line[index]
				if !haveFirst {
					number = int(digit-delta) * 10
					haveFirst = true
				}

				index++
				continue
			}

			found := false
			for j := 3; j < 6; j++ {
				if index+j > end {
					break
				}

				value, ok := digits[line[index:index+j]]
				if ok {
					digit = value
					if !haveFirst {
						number = int(digit-delta) * 10
						haveFirst = true
					}

					index += j
					found = true
					break
				}
			}

			if !found {
				index++
			}
		}

		number += int(digit - delta)
		sum += number
	}

	return sum
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

	lines := readInput(file)
	fmt.Println("Part1:", part1(lines))
	fmt.Println("Part2:", part2(lines))
}
