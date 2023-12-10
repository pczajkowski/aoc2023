package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Up = iota
	Down
	Left
	Right
)

type Point struct {
	y, x      int
	direction int
}

func (p *Point) move(maze [][]byte, height int, width int) bool {
	switch maze[p.y][p.x] {
	case '|':
		if p.direction == Up && p.y-1 >= 0 {
			p.y--
			return true
		} else if p.direction == Down && p.y+1 < height {
			p.y++
			return true
		}
	case '-':
		if p.direction == Right && p.x+1 < width {
			p.x++
			return true
		} else if p.direction == Left && p.x-1 >= 0 {
			p.x--
			return true
		}
	case 'L':
		if p.direction == Down && p.x+1 < width {
			p.x++
			p.direction = Right
			return true
		} else if p.direction == Left && p.y-1 >= 0 {
			p.y--
			p.direction = Up
			return true
		}
	case 'J':
		if p.direction == Down && p.x-1 >= 0 {
			p.x--
			p.direction = Left
			return true
		} else if p.direction == Right && p.y-1 >= 0 {
			p.y--
			p.direction = Up
			return true
		}
	case '7':
		if p.direction == Right && p.y+1 < height {
			p.y++
			p.direction = Down
			return true
		} else if p.direction == Up && p.x-1 >= 0 {
			p.x--
			p.direction = Left
			return true
		}
	case 'F':
		if p.direction == Up && p.x+1 < width {
			p.x++
			p.direction = Right
			return true
		} else if p.direction == Left && p.y+1 < height {
			p.y++
			p.direction = Down
			return true
		}

	}

	return false
}

const (
	verticalAbove   = "|7F"
	verticalBelow   = "|LJ"
	horizontalLeft  = "-LF"
	horizontalRight = "-J7"
)

func establishS(s Point, maze [][]byte, height int, width int) (byte, Point) {
	var left, right, up, down bool
	if s.x-1 >= 0 && strings.Contains(horizontalLeft, string(maze[s.y][s.x-1])) {
		left = true
	}

	if s.x+1 < width && strings.Contains(horizontalRight, string(maze[s.y][s.x+1])) {
		right = true
	}

	if s.y-1 >= 0 && strings.Contains(verticalAbove, string(maze[s.y-1][s.x])) {
		up = true
	}

	if s.y+1 < height && strings.Contains(verticalBelow, string(maze[s.y+1][s.x])) {
		down = true
	}

	var char byte
	if left {
		if right {
			char = '-'
			s.direction = Right
		} else if up {
			char = 'J'
			s.direction = Down
		} else if down {
			char = '7'
			s.direction = Up
		}
	} else if right {
		if up {
			char = 'L'
			s.direction = Down
		} else if down {
			char = 'F'
			s.direction = Up
		}
	} else if up {
		char = '|'
		s.direction = Down
	}

	return char, s
}

func part1(maze [][]byte) int {
	height := len(maze)
	width := len(maze[0])
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == 'S' {
				char, start := establishS(Point{y: y, x: x}, maze, height, width)
				maze[y][x] = char
				current := start
				steps := 0
				for {
					if !current.move(maze, height, width) {
						break
					}
					steps++

					if current.x == start.x && current.y == start.y {
						return steps / 2
					}
				}
			}
		}
	}

	return -1
}

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var maze [][]byte
	width := -1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if width < 0 {
			width = len(line)
		}

		row := make([]byte, width)
		for i := range line {
			row[i] = line[i]
		}

		maze = append(maze, row)
	}

	return maze
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

	maze := readInput(file)
	fmt.Println("Part1:", part1(maze))
}
