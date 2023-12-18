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

func (p *Point) getPoints(dig Dig) []Point {
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

		points = append(points, *p)
	}

	return points
}

func plot(plan []Dig) []Point {
	var current Point
	result := []Point{current}

	for i := range plan {
		result = append(result, current.getPoints(plan[i])...)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].y == result[j].y {
			return result[i].x < result[j].x
		}

		return result[i].y < result[j].y
	})

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

	plan := readInput(file)
	fmt.Println(plot(plan))
}
