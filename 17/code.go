package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	Diff     = 48
	MaxMoves = 3
)

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var board [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []int
		for i := range line {
			row = append(row, int(line[i]-Diff))
		}

		board = append(board, row)
	}

	return board
}

const (
	North = iota
	East
	South
	West
)

type Point struct {
	y, x int
}

type Destination struct {
	pos       Point
	cost      int
	moves     int
	direction int
}

func getNorth(board [][]int, height int, width int, lava Destination) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == North {
		moves = lava.moves
	}

	if moves > MaxMoves {
		return destinations
	}

	end := lava.pos.y - MaxMoves
	if end < 0 {
		end = 0
	}

	cost := lava.cost
	for y := lava.pos.y - 1; y >= end; y-- {
		cost += board[y][lava.pos.x]
		moves++
		if moves > MaxMoves {
			break
		}

		destinations = append(destinations, Destination{pos: Point{y: y, x: lava.pos.x}, moves: moves, cost: cost, direction: North})
	}

	return destinations
}

func getEast(board [][]int, height int, width int, lava Destination) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == East {
		moves = lava.moves
	}

	if moves > MaxMoves {
		return destinations
	}

	end := lava.pos.x + MaxMoves
	if end >= width {
		end = width - 1
	}

	cost := lava.cost
	for x := lava.pos.x + 1; x <= end; x++ {
		cost += board[lava.pos.y][x]
		moves++
		if moves > MaxMoves {
			break
		}

		destinations = append(destinations, Destination{pos: Point{y: lava.pos.y, x: x}, moves: moves, cost: cost, direction: East})
		if moves > MaxMoves {
			break
		}
	}

	return destinations
}

func getSouth(board [][]int, height int, width int, lava Destination) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == South {
		moves = lava.moves
	}

	if moves > MaxMoves {
		return destinations
	}

	end := lava.pos.y + MaxMoves
	if end >= height {
		end = height - 1
	}

	cost := lava.cost
	for y := lava.pos.y + 1; y <= end; y++ {
		cost += board[y][lava.pos.x]
		moves++
		if moves > MaxMoves {
			break
		}

		destinations = append(destinations, Destination{pos: Point{y: y, x: lava.pos.x}, moves: moves, cost: cost, direction: South})
	}

	return destinations
}

func getWest(board [][]int, height int, width int, lava Destination) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == West {
		moves = lava.moves
	}

	if moves > MaxMoves {
		return destinations
	}

	end := lava.pos.x - MaxMoves
	if end < 0 {
		end = 0
	}

	cost := lava.cost
	for x := lava.pos.x - 1; x >= end; x-- {
		cost += board[lava.pos.y][x]
		moves++
		if moves > MaxMoves {
			break
		}

		destinations = append(destinations, Destination{pos: Point{y: lava.pos.y, x: x}, moves: moves, cost: cost, direction: West})
	}

	return destinations
}

func getDirections(direction int) []int {
	switch direction {
	case North:
		return []int{East, North, West}
	case East:
		return []int{East, South, North}
	case South:
		return []int{East, South, West}
	case West:
		return []int{South, West, North}
	}

	return []int{}
}

func getDestinations(board [][]int, height int, width int, lava Destination) []Destination {
	var destinations []Destination
	directions := getDirections(lava.direction)
	for i := range directions {
		switch i {
		case North:
			destinations = append(destinations, getNorth(board, height, width, lava)...)
		case East:
			destinations = append(destinations, getEast(board, height, width, lava)...)
		case South:
			destinations = append(destinations, getSouth(board, height, width, lava)...)
		case West:
			destinations = append(destinations, getWest(board, height, width, lava)...)
		}
	}

	return destinations
}

func part1(board [][]int) int {
	min := 1000000
	height := len(board)
	width := len(board[0])
	goal := Point{y: height - 1, x: width - 1}
	explored := make(map[Point]int)
	lava := Destination{pos: Point{x: 0, y: 0}, moves: 0, direction: East}
	prev := lava
	frontier := []Destination{lava}

	for {
		if len(frontier) == 0 {
			break
		}

		current := frontier[0]
		frontier = frontier[1:]

		if current.direction == prev.direction && current.moves+prev.moves > MaxMoves {
			continue
		}

		prev = current
		if current.pos == goal {
			if min > current.cost {
				min = current.cost
			}
		}

		successors := getDestinations(board, height, width, current)
		fmt.Println(current, successors)
		for i := range successors {
			newCost := successors[i].cost + successors[i].moves
			value, ok := explored[successors[i].pos]
			if !ok || value > newCost {
				explored[successors[i].pos] = newCost
				frontier = append(frontier, successors[i])
			}
		}

		sort.Slice(frontier, func(i, j int) bool {
			return frontier[i].cost < frontier[j].cost
		})

	}

	return min
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
	fmt.Println("Part1:", part1(board))
}
