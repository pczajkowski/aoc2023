package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	Horizontal = '-'
	Vertical   = '|'
	Slash      = '/'
	Backslash  = '\\'
	Empty      = '.'
	Mark       = '#'
)

const (
	North = iota
	South
	East
	West
)

type Point struct {
	y, x      int
	direction int
}

type Beam struct {
	pos     Point
	wasHere map[Point]bool
}

func (b *Beam) directions(board [][]byte, height int, width int, pastBeams map[Point]bool) []Beam {
	switch board[b.pos.y][b.pos.x] {
	case Horizontal:
		if b.pos.direction != East && b.pos.direction != West {
			var beams []Beam
			b.pos.direction = East
			if !pastBeams[b.pos] {
				pastBeams[b.pos] = true
				beams = append(beams, *b)
			}
			b.pos.direction = West
			if !pastBeams[b.pos] {
				pastBeams[b.pos] = true
				beams = append(beams, *b)
			}

			return beams
		}
	case Vertical:
		if b.pos.direction != South && b.pos.direction != North {
			var beams []Beam
			b.pos.direction = South
			if !pastBeams[b.pos] {
				pastBeams[b.pos] = true
				beams = append(beams, *b)
			}
			b.pos.direction = North
			if !pastBeams[b.pos] {
				pastBeams[b.pos] = true
				beams = append(beams, *b)
			}

			return beams
		}
	case Slash:
		switch b.pos.direction {
		case North:
			b.pos.direction = East
		case South:
			b.pos.direction = West
		case East:
			b.pos.direction = North
		case West:
			b.pos.direction = South
		}
	case Backslash:
		switch b.pos.direction {
		case North:
			b.pos.direction = West
		case South:
			b.pos.direction = East
		case East:
			b.pos.direction = South
		case West:
			b.pos.direction = North
		}
	}

	return []Beam{*b}
}

func (b *Beam) move(board [][]byte, height int, width int, pastBeams map[Point]bool) []Beam {
	b.wasHere[b.pos] = true

	var beams []Beam
	directions := b.directions(board, height, width, pastBeams)
	for i := range directions {
		switch directions[i].pos.direction {
		case North:
			directions[i].pos.y--
		case South:
			directions[i].pos.y++
		case East:
			directions[i].pos.x++
		case West:
			directions[i].pos.x--
		}

		if directions[i].wasHere[directions[i].pos] || directions[i].pos.x < 0 || directions[i].pos.x >= width || directions[i].pos.y < 0 || directions[i].pos.y >= height {
			continue
		}

		beams = append(beams, directions[i])
	}

	return beams
}

func emptyBoard(height int, width int) [][]byte {
	var board [][]byte
	for i := 0; i < height; i++ {
		board = append(board, make([]byte, width))
	}

	return board
}

func count(board [][]byte) int {
	var result int
	for y := range board {
		for x := range board[y] {
			if board[y][x] == Mark {
				result++
			}
		}
	}

	return result
}

func processBeams(board [][]byte, height int, width int, beam Beam) int {
	trackBoard := emptyBoard(height, width)
	pastBeams := make(map[Point]bool)
	beams := []Beam{beam}

	for {
		if len(beams) == 0 {
			break
		}

		var newBeams []Beam
		for i := range beams {
			pastBeams[beams[i].pos] = true
			if trackBoard[beams[i].pos.y][beams[i].pos.x] != Mark {
				trackBoard[beams[i].pos.y][beams[i].pos.x] = Mark
			}

			newBeams = append(newBeams, beams[i].move(board, height, width, pastBeams)...)
		}

		beams = newBeams
	}

	return count(trackBoard)
}

func part1(board [][]byte) int {
	height := len(board)
	width := len(board[0])

	return processBeams(board, height, width, Beam{pos: Point{y: 0, x: 0, direction: East}, wasHere: make(map[Point]bool)})
}

func part2(board [][]byte) int {
	height := len(board)
	bottomEdge := height - 1
	width := len(board[0])
	rightEdge := width - 1
	max := 0

	var beams []Beam
	for x := 0; x < width; x++ {
		beams = append(beams, Beam{pos: Point{y: 0, x: x, direction: South}, wasHere: make(map[Point]bool)})
		if x == 0 {
			beams = append(beams, Beam{pos: Point{y: 0, x: x, direction: East}, wasHere: make(map[Point]bool)})
		}

		if x == rightEdge {
			beams = append(beams, Beam{pos: Point{y: 0, x: x, direction: West}, wasHere: make(map[Point]bool)})
		}
	}

	for x := 0; x < width; x++ {
		beams = append(beams, Beam{pos: Point{y: bottomEdge, x: x, direction: North}, wasHere: make(map[Point]bool)})
		if x == 0 {
			beams = append(beams, Beam{pos: Point{y: bottomEdge, x: x, direction: East}, wasHere: make(map[Point]bool)})
		}

		if x == rightEdge {
			beams = append(beams, Beam{pos: Point{y: bottomEdge, x: x, direction: West}, wasHere: make(map[Point]bool)})
		}
	}

	for y := 0; y < height; y++ {
		beams = append(beams, Beam{pos: Point{y: y, x: 0, direction: East}, wasHere: make(map[Point]bool)})
		if y == 0 {
			beams = append(beams, Beam{pos: Point{y: y, x: 0, direction: South}, wasHere: make(map[Point]bool)})
		}

		if y == bottomEdge {
			beams = append(beams, Beam{pos: Point{y: y, x: 0, direction: North}, wasHere: make(map[Point]bool)})
		}

	}

	for y := 0; y < height; y++ {
		beams = append(beams, Beam{pos: Point{y: y, x: rightEdge, direction: West}, wasHere: make(map[Point]bool)})
		if y == 0 {
			beams = append(beams, Beam{pos: Point{y: y, x: rightEdge, direction: South}, wasHere: make(map[Point]bool)})
		}

		if y == bottomEdge {
			beams = append(beams, Beam{pos: Point{y: y, x: rightEdge, direction: North}, wasHere: make(map[Point]bool)})
		}

	}

	for i := range beams {
		energy := processBeams(board, height, width, beams[i])
		if energy > max {
			max = energy
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
	fmt.Println("Part1:", part1(board))
	fmt.Println("Part2:", part2(board))
}
