package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x  int
	steps int
}

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var board [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []byte
		for i := range line {
			row = append(row, line[i])
		}

		board = append(board, row)
	}

	return board
}

const (
	Up    = '^'
	Right = '>'
	Down  = 'V'
	Left  = '<'
	Rock  = '#'
)

func (p *Point) getDestinations(board [][]byte, height int, width int) []Point {
	var destinations []Point
	p.steps++
	slope := false
	switch board[p.y][p.x] {
	case Up:
		p.y--
		slope = true
	case Down:
		p.y++
		slope = true
	case Left:
		p.x--
		slope = true
	case Right:
		p.x++
		slope = true
	}

	if slope {
		destinations = append(destinations, *p)
		return destinations
	}

	if p.x-1 >= 0 && board[p.y][p.x-1] != Rock {
		destinations = append(destinations, Point{y: p.y, x: p.x - 1, steps: p.steps})
	}

	if p.x+1 < width && board[p.y][p.x+1] != Rock {
		destinations = append(destinations, Point{y: p.y, x: p.x + 1, steps: p.steps})
	}

	if p.y-1 >= 0 && board[p.y-1][p.x] != Rock {
		destinations = append(destinations, Point{y: p.y - 1, x: p.x, steps: p.steps})
	}

	if p.y+1 >= 0 && board[p.y+1][p.x] != Rock {
		destinations = append(destinations, Point{y: p.y + 1, x: p.x, steps: p.steps})
	}

	return destinations
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func calculate(board [][]byte) int {
	height := len(board)
	width := len(board[0])

	start := Point{y: 0, x: 1}
	goal := Point{y: height - 1, x: width - 2}

	visited := make(map[string]int)
	visited[start.key()] = start.steps
	frontier := []Point{start}
	var max int

	for {
		if len(frontier) == 0 {
			break
		}

		current := frontier[0]
		frontier = frontier[1:]

		if current.x == goal.x && current.y == goal.y {
			if max < current.steps {
				max = current.steps
			}

			continue
		}

		successors := current.getDestinations(board, height, width)
		for i := range successors {
			value, ok := visited[successors[i].key()]
			if !ok || visited[successors[i].key()] > value {
				visited[successors[i].key()] = successors[i].steps
				frontier = append(frontier, successors[i])
			}
		}
	}

	return max
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

	board := readInput(file)
	fmt.Println("Part1:", calculate(board))
}
