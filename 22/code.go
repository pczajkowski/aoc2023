package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Object struct {
	xMin, yMin, zMin   int
	xMax, yMax, zMax   int
	zRealMin, zRealMax int
	supports           []*Object
	standsOn           []*Object
}

func readInput(file *os.File) []Object {
	scanner := bufio.NewScanner(file)
	var objects []Object

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var object Object
		n, err := fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &object.xMin, &object.yMin, &object.zMin, &object.xMax, &object.yMax, &object.zMax)
		if n != 6 || err != nil {
			log.Fatalf("Bad input: %s\n%s", line, err)
		}

		objects = append(objects, object)
	}

	return objects
}

func intersect(a, b Object) bool {
	return a.xMin <= b.xMax && a.xMax >= b.xMin && a.yMin <= b.yMax && a.yMax >= b.yMin
}

func part1(objects []Object) int {
	var result int
	sort.Slice(objects, func(i, j int) bool { return objects[i].zMin < objects[j].zMin })
	fmt.Println(objects)
	fmt.Println(objects[1], objects[2], intersect(objects[1], objects[2]))

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

	objects := readInput(file)
	fmt.Println("Part1:", part1(objects))
}
