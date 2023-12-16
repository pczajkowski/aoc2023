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

type Beam struct {
	y, x      int
	direction int
}

func (b *Beam) canContinue(board [][]byte, height int, width int) []Beam {
	var beams []Beam
	if b.x < 0 || b.x >= width || b.y < 0 || b.y >= height {
		return beams
	}

	switch board[b.y][b.x] {
	case Horizontal:
		if b.direction != East && b.direction != West {
			b.direction = East
			beams = append(beams, *b)
			b.direction = West
			beams = append(beams, *b)

			return beams
		}
	case Vertical:
		if b.direction != South && b.direction != North {
			b.direction = South
			beams = append(beams, *b)
			b.direction = North
			beams = append(beams, *b)

			return beams
		}
	case Slash:
		switch b.direction {
		case North:
			b.direction = East
		case South:
			b.direction = West
		case East:
			b.direction = North
		case West:
			b.direction = South
		}
	case Backslash:
		switch b.direction {
		case North:
			b.direction = West
		case South:
			b.direction = East
		case East:
			b.direction = South
		case West:
			b.direction = North
		}
	}

	return append(beams, *b)
}

func (b *Beam) move(board [][]byte, height int, width int) []Beam {
	switch b.direction {
	case North:
		b.y--
	case South:
		b.y++
	case East:
		b.x++
	case West:
		b.x--
	}

	return b.canContinue(board, height, width)
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

func part1(board [][]byte) int {
	height := len(board)
	width := len(board[0])
	var result int
	trackBoard := emptyBoard(height, width)
	beams := []Beam{Beam{y: 0, x: 0, direction: East}}

	for {
		change := false
		var newBeams []Beam
		for i := range beams {
			if trackBoard[beams[i].y][beams[i].x] != Mark {
				trackBoard[beams[i].y][beams[i].x] = Mark
				change = true
			}

			newBeams = append(newBeams, beams[i].move(board, height, width)...)
		}

		if !change {
			break
		}

		beams = newBeams
	}

	fmt.Println(count(trackBoard), trackBoard)
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

	board := readInput(file)
	fmt.Println(part1(board))
}
