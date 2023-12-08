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

func part1(network Network) int {
	steps := 0
	current := "AAA"
	mod := len(network.moves)
	index := 0
	for {
		if current == "ZZZ" {
			break
		}

		d := network.paths[current]
		if network.moves[index] == 'L' {
			current = d.left
		} else {
			current = d.right
		}

		index = (index + 1) % mod
		steps++
	}

	return steps
}

func (n *Network) AtGoal() bool {
	for i := range n.starts {
		if !strings.HasSuffix(n.starts[i], "Z") {
			return false
		}
	}

	return true
}

func part2(network Network) int {
	steps := 0
	mod := len(network.moves)
	index := 0
	for {
		if network.AtGoal() {
			break
		}

		for i := range network.starts {
			d := network.paths[network.starts[i]]
			if network.moves[index] == 'L' {
				network.starts[i] = d.left
			} else {
				network.starts[i] = d.right
			}

		}

		steps++
		index = (index + 1) % mod
	}

	return steps
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
	fmt.Println("Part1:", part1(network))
	fmt.Println("Part2:", part2(network))
}
