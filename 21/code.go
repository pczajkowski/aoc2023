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

func readInput(file *os.File) (Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var board [][]byte
	var start Point
	var index int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []byte
		for i := range line {
			if line[i] == 'S' {
				start.x = i
				start.y = index
			}

			row = append(row, line[i])
		}

		board = append(board, row)
		index++
	}

	return start, board
}

const (
	Rock = '#'
)

func (p *Point) getDestinations(board [][]byte, height int, width int, maxSteps int) []Point {
	var destinations []Point
	p.steps++
	if p.steps > maxSteps {
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

	if p.y+1 < height && board[p.y+1][p.x] != Rock {
		destinations = append(destinations, Point{y: p.y + 1, x: p.x, steps: p.steps})
	}

	return destinations
}

func (p *Point) key() string {
	return fmt.Sprintf("%d-%d", p.y, p.x)
}

func calculate(board [][]byte, start Point, maxSteps int) int {
	height := len(board)
	width := len(board[0])
	visited := make(map[string]int)
	visited[start.key()] = start.steps
	frontier := []Point{start}

	for {
		if len(frontier) == 0 {
			break
		}

		current := frontier[0]
		frontier = frontier[1:]

		successors := current.getDestinations(board, height, width, maxSteps)
		for i := range successors {
			value, ok := visited[successors[i].key()]
			if !ok || value < maxSteps && successors[i].steps != value {
				visited[successors[i].key()] = successors[i].steps
				frontier = append(frontier, successors[i])
			}
		}
	}

	count := 0
	for _, v := range visited {
		if v == maxSteps {
			count++
		}
	}

	return count
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

	start, board := readInput(file)
	fmt.Println("Part1:", calculate(board, start, 6))
}
