package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Dig struct {
	direction string
	length    int
	color     string
}

func readInput(file *os.File) []Dig {
	scanner := bufio.NewScanner(file)
	var plan []Dig

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var current Dig
		n, err := fmt.Sscanf(line, "%s %d %s", &current.direction, &current.length, &current.color)
		if n != 3 || err != nil {
			log.Fatalf("Wrong input: %s\n%s", line, err)
		}

		current.color = strings.Trim(current.color, "()")
		plan = append(plan, current)
	}

	return plan
}

type Point struct {
	y, x int
}

const (
	Up    = "U"
	Down  = "D"
	Left  = "L"
	Right = "R"
)

func (p *Point) getPoints(dig Dig, wasHere map[Point]bool) []Point {
	var points []Point
	for i := 0; i < dig.length; i++ {
		switch dig.direction {
		case Up:
			p.y--
		case Down:
			p.y++
		case Left:
			p.x--
		case Right:
			p.x++
		}

		if !wasHere[*p] {
			wasHere[*p] = true
			points = append(points, *p)
		}
	}

	return points
}

func plot(plan []Dig) []Point {
	var current Point
	result := []Point{current}
	wasHere := make(map[Point]bool)
	wasHere[current] = true

	for i := range plan {
		result = append(result, current.getPoints(plan[i], wasHere)...)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].y == result[j].y {
			return result[i].x < result[j].x
		}

		return result[i].y < result[j].y
	})

	return result
}

func countBetween(plot []Point) int {
	var count int
	edge := len(plot)
	prev := plot[0]
	for i := 1; i < edge; i++ {
		if prev.y == plot[i].y {
			if prev.x == plot[i].x {
				continue
			}

			c := plot[i].x - prev.x - 1
			count += c

			if c > 0 && i+1 < edge {
				prev = plot[i+1]
				continue
			}
		}
		prev = plot[i]
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

	plan := readInput(file)
	plot := plot(plan)
	fmt.Println("Part1:", countBetween(plot)+len(plot))
}
