package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directions struct {
	left  string
	right string
}

type Network struct {
	moves  string
	paths  map[string]Directions
	starts []string
}

func readInput(file *os.File) Network {
	scanner := bufio.NewScanner(file)
	var network Network
	readMoves := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if !readMoves {
				readMoves = true
				continue
			}

			break
		}

		if !readMoves {
			network.moves = line
			network.paths = make(map[string]Directions)
		} else {
			fromParts := strings.Split(line, " = ")
			if len(fromParts) != 2 {
				log.Fatalf("Wrong number of fromParts: %s", line)
			}

			from := fromParts[0]
			parts := strings.Split(fromParts[1], ", ")
			if len(parts) != 2 {
				log.Fatalf("Wrong number of parts: %s", fromParts[1])
			}

			var directions Directions
			directions.left = strings.TrimLeft(parts[0], "(")
			directions.right = strings.TrimRight(parts[1], ")")
			network.paths[from] = directions

			if strings.HasSuffix(from, "A") {
				network.starts = append(network.starts, from)
			}
		}
	}

	return network
}

func atGoal(starts string, goal string) bool {
	return strings.HasSuffix(starts, goal)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	l := min(a, b)
	h := max(a, b)

	return gcd(l, h%l)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func part(network Network, starts []string, goal string) int {
	mod := len(network.moves)
	result := 1
	for i := range starts {
		steps := 0
		index := 0
		for {
			if atGoal(starts[i], goal) {
				break
			}

			turn := network.moves[index]

			d := network.paths[starts[i]]
			if turn == 'L' {
				starts[i] = d.left
			} else {
				starts[i] = d.right
			}

			steps++
			index = (index + 1) % mod
		}

		result = lcm(result, steps)
	}

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

	network := readInput(file)
	fmt.Println("Part1:", part(network, []string{"AAA"}, "ZZZ"))
	fmt.Println("Part2:", part(network, network.starts, "Z"))
}
