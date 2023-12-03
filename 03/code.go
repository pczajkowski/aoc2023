package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	return lines
}

const firstDigit = 48
const lastDigit = 57

func isDigit(b byte) bool {
	return firstDigit <= b && b <= lastDigit
}

func isSymbol(b byte) bool {
	return !isDigit(b) && b != '.'
}

func symbolNear(lines []string, height int, width int, y int, start int, end int) bool {
	i := start - 1
	if i < 0 {
		i = start
	}

	if end >= width {
		end = width - 1
	}

	canGoUp := y-1 >= 0
	canGoDown := y+1 < height

	for ; i <= end; i++ {
		if isSymbol(lines[y][i]) {
			return true
		}

		if canGoUp && isSymbol(lines[y-1][i]) {
			return true
		}

		if canGoDown && isSymbol(lines[y+1][i]) {
			return true
		}
	}

	return false
}

func part1(lines []string) int {
	var result int
	height := len(lines)
	width := len(lines[0])
	edge := width - 1
	for i := range lines {
		var start, end int
		gotNumber := false
		tryRead := false
		for j := range lines[i] {
			if isDigit(lines[i][j]) {
				if !gotNumber {
					start = j
					gotNumber = true
				}

				if j == edge {
					end = j + 1
					tryRead = true
				}
			} else {
				if !gotNumber {
					continue
				}

				end = j
				gotNumber = false
				tryRead = true
			}

			if tryRead {
				if symbolNear(lines, height, width, i, start, end) {
					var d int
					n, err := fmt.Sscanf(lines[i][start:end], "%d", &d)
					if n != 1 || err != nil {
						log.Fatalf("Wrong input: %s\n%s", lines[i][start:end], err)
					}

					result += d
				}

				tryRead = false
			}

		}
	}

	return result
}

type Point struct {
	x, y int
}

func digitsFromLine(line string, start int, end int, y int) []Point {
	var points []Point
	reading := false
	for i := start; i <= end; i++ {
		if isDigit(line[i]) {
			if !reading {
				points = append(points, Point{x: i, y: y})
				reading = true
			}
		} else {
			reading = false
		}
	}

	return points
}

func getDigits(lines []string, height int, width int, y int, x int) []Point {
	var points []Point

	start := x - 1
	if start < 0 {
		start = x
	}

	end := x + 1
	if end >= width {
		end = x
	}

	if y-1 >= 0 {
		points = append(points, digitsFromLine(lines[y-1], start, end, y-1)...)
	}

	if y+1 < height {
		points = append(points, digitsFromLine(lines[y+1], start, end, y+1)...)
	}

	points = append(points, digitsFromLine(lines[y], start, end, y)...)

	return points
}

func numberFromPoint(lines []string, width int, point Point) int {
	start := point.x
	end := point.x
	var haveStart, haveEnd bool
	for {
		if haveStart && haveEnd {
			break
		}

		if !haveStart && start-1 >= 0 && isDigit(lines[point.y][start-1]) {
			start--
		} else {
			haveStart = true
		}

		if !haveEnd && end+1 < width && isDigit(lines[point.y][end+1]) {
			end++
		} else {
			haveEnd = true
		}
	}
	end++

	var d int
	n, err := fmt.Sscanf(lines[point.y][start:end], "%d", &d)
	if n != 1 || err != nil {
		log.Fatalf("Wrong input: %s\n%s", lines[point.y][start:end], err)
	}

	return d
}

func part2(lines []string) int {
	height := len(lines)
	width := len(lines[0])
	var result int
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '*' {
				digits := getDigits(lines, height, width, i, j)
				if len(digits) == 2 {
					prod := 1
					for k := range digits {
						prod *= numberFromPoint(lines, width, digits[k])
					}

					result += prod
				}
			}
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

	lines := readInput(file)
	fmt.Println("Part1:", part1(lines))
	fmt.Println("Part2:", part2(lines))
}
