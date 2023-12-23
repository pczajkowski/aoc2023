package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	None = iota
	North
	South
	East
	West
)

type Point struct {
	y, x       int
	steps      int
	direction  int
	mustGoDown bool
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

	if p.mustGoDown && p.y+1 < height && board[p.y+1][p.x] != Rock {
		p.y++
		p.direction = None
		p.mustGoDown = false
		destinations = append(destinations, *p)
		return destinations
	}

	switch board[p.y][p.x] {
	case Up:
		p.direction = North
		p.mustGoDown = true
	case Down:
		p.direction = South
		p.mustGoDown = true
	case Left:
		p.direction = West
		p.mustGoDown = true
	case Right:
		p.direction = East
		p.mustGoDown = true
	}

	if p.direction != None {
		switch p.direction {
		case North:
			p.y--
		case South:
			p.y++
		case West:
			p.x--
		case East:
			p.x++
		}

		destinations = append(destinations, *p)
		return destinations
	}

	if p.y-1 >= 0 && board[p.y-1][p.x] != Rock {
		destinations = append(destinations, Point{y: p.y - 1, x: p.x, steps: p.steps})
	}

	if p.x+1 < width && board[p.y][p.x+1] != Rock {
		destinations = append(destinations, Point{y: p.y, x: p.x + 1, steps: p.steps})
	}

	if p.x-1 >= 0 && board[p.y][p.x-1] != Rock {
		destinations = append(destinations, Point{y: p.y, x: p.x - 1, steps: p.steps})
	}

	if p.y+1 < height && board[p.y+1][p.x] != Rock {
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
			_, ok := visited[successors[i].key()]
			if !ok {
				visited[successors[i].key()] = successors[i].steps
				frontier = append(frontier, successors[i])
			}
		}
		sort.Slice(frontier, func(i, j int) bool { return frontier[i].steps > frontier[j].steps })
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
