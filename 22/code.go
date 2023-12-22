package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Object struct {
	xmin, ymin, zmin int
	xmax, ymax, zmax int
}

func readInput(file *os.File) (map[int][]Object, map[int][]Object, []Object) {
	scanner := bufio.NewScanner(file)
	startsAt := make(map[int][]Object)
	endsAt := make(map[int][]Object)
	var objects []Object

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var object Object
		n, err := fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &object.xmin, &object.ymin, &object.zmin, &object.xmax, &object.ymax, &object.zmax)
		if n != 6 || err != nil {
			log.Fatalf("Bad input: %s\n%s", line, err)
		}

		startsAt[object.zmin] = append(startsAt[object.zmin], object)
		endsAt[object.zmax] = append(endsAt[object.zmax], object)
		objects = append(objects, object)
	}

	return startsAt, endsAt, objects
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

	startsAt, endsAt, objects := readInput(file)
	fmt.Println(startsAt, endsAt, objects)
}
