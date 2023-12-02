package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cube struct {
	count int
	color string
}

type Game struct {
	id   int
	sets [][]Cube
}

func readInput(file *os.File) []Game {
	scanner := bufio.NewScanner(file)
	var games []Game

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var current Game
		n, err := fmt.Sscanf(line, "Game %d:", &current.id)
		if n != 1 || err != nil {
			log.Fatalf("Failed to read game: %s\n%s", line, err)
		}

		gameParts := strings.Split(line, ":")
		if len(gameParts) != 2 {
			log.Fatalf("Wrong input: %s", line)
		}

		sets := strings.Split(gameParts[1], ";")
		if len(sets) == 0 {
			log.Fatalf("Wrong input: %s", gameParts[1])
		}

		for i := range sets {
			var set []Cube
			cubes := strings.Split(sets[i], ",")
			if len(cubes) == 0 {
				log.Fatalf("Wrong input: %s", sets[i])
			}

			for j := range cubes {
				var cube Cube
				n, err = fmt.Sscanf(cubes[j], "%d %s", &cube.count, &cube.color)
				if n != 2 || err != nil {
					log.Fatalf("Wrong input: %s\n%s", cubes[j], err)
				}

				set = append(set, cube)
			}

			current.sets = append(current.sets, set)
		}

		games = append(games, current)
	}

	return games
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

	games := readInput(file)
	fmt.Println(games)
}
