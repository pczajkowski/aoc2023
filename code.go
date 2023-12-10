package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func (p *Point) move(maze []string, height int, width int) bool {
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

func part1(maze []string) int {
	biggest := 0
	height := len(maze)
	width := len(maze[0])
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == 'F' {
				current := Point{y: y, x: x}
				start := Point{y: y, x: x}
				steps := 0
				for {
					if !current.move(maze, height, width) {
						break
					}
					steps++

					if current.x == start.x && current.y == start.y {
						steps /= 2
						if steps > biggest {
							biggest = steps
						}

						break
					}
				}
			}
		}
	}

	return biggest
}

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var maze []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		maze = append(maze, line)
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
