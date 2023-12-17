package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Diff = 48
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

func getNorth(board [][]int, height int, width int, lava Destination, minMoves int, maxMoves int) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == North {
		moves = lava.moves
	}

	end := lava.pos.y - maxMoves
	if end < 0 {
		end = 0
	}

	cost := lava.cost
	for y := lava.pos.y - 1; y >= end; y-- {
		cost += board[y][lava.pos.x]
		moves++
		if moves > maxMoves {
			break
		}

		if moves > minMoves {
			destinations = append(destinations, Destination{pos: Point{y: y, x: lava.pos.x}, moves: moves, cost: cost, direction: North})
		}
	}

	return destinations
}

func getEast(board [][]int, height int, width int, lava Destination, minMoves int, maxMoves int) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == East {
		moves = lava.moves
	}

	end := lava.pos.x + maxMoves
	if end >= width {
		end = width - 1
	}

	cost := lava.cost
	for x := lava.pos.x + 1; x <= end; x++ {
		cost += board[lava.pos.y][x]
		moves++
		if moves > maxMoves {
			break
		}

		if moves > minMoves {
			destinations = append(destinations, Destination{pos: Point{y: lava.pos.y, x: x}, moves: moves, cost: cost, direction: East})
		}
	}

	return destinations
}

func getSouth(board [][]int, height int, width int, lava Destination, minMoves int, maxMoves int) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == South {
		moves = lava.moves
	}

	end := lava.pos.y + maxMoves
	if end >= height {
		end = height - 1
	}

	cost := lava.cost
	for y := lava.pos.y + 1; y <= end; y++ {
		cost += board[y][lava.pos.x]
		moves++
		if moves > maxMoves {
			break
		}

		if moves > minMoves {
			destinations = append(destinations, Destination{pos: Point{y: y, x: lava.pos.x}, moves: moves, cost: cost, direction: South})
		}
	}

	return destinations
}

func getWest(board [][]int, height int, width int, lava Destination, minMoves int, maxMoves int) []Destination {
	var destinations []Destination
	moves := 0
	if lava.direction == West {
		moves = lava.moves
	}

	end := lava.pos.x - maxMoves
	if end < 0 {
		end = 0
	}

	cost := lava.cost
	for x := lava.pos.x - 1; x >= end; x-- {
		cost += board[lava.pos.y][x]
		moves++
		if moves > maxMoves {
			break
		}

		if moves > minMoves {
			destinations = append(destinations, Destination{pos: Point{y: lava.pos.y, x: x}, moves: moves, cost: cost, direction: West})
		}
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

func getDestinations(board [][]int, height int, width int, lava Destination, minMoves int, maxMoves int) []Destination {
	var destinations []Destination
	directions := getDirections(lava.direction)
	for i := range directions {
		switch directions[i] {
		case North:
			destinations = append(destinations, getNorth(board, height, width, lava, minMoves, maxMoves)...)
		case East:
			destinations = append(destinations, getEast(board, height, width, lava, minMoves, maxMoves)...)
		case South:
			destinations = append(destinations, getSouth(board, height, width, lava, minMoves, maxMoves)...)
		case West:
			destinations = append(destinations, getWest(board, height, width, lava, minMoves, maxMoves)...)
		}
	}

	return destinations
}

type Visited struct {
	pos       Point
	direction int
}

func calculate(board [][]int, minMoves int, maxMoves int) int {
	min := 1000000
	height := len(board)
	width := len(board[0])
	goal := Point{y: height - 1, x: width - 1}
	explored := make(map[Visited]int)
	lava := Destination{pos: Point{x: 0, y: 0}, moves: 0, direction: East}
	frontier := []Destination{lava}

	for {
		if len(frontier) == 0 {
			break
		}

		current := frontier[0]
		frontier = frontier[1:]

		if current.pos == goal {
			if min > current.cost {
				min = current.cost
			}
		}

		successors := getDestinations(board, height, width, current, minMoves, maxMoves)
		for i := range successors {
			v := Visited{pos: successors[i].pos, direction: successors[i].direction}
			newCost := successors[i].cost
			value, ok := explored[v]
			if !ok || value > newCost {
				explored[v] = newCost
				frontier = append(frontier, successors[i])
			}
		}
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
	fmt.Println("Part1:", calculate(board, 0, 3))
	fmt.Println("Part2:", calculate(board, 3, 10))
}
