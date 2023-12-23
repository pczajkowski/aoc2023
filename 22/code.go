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
	canBeDeleted       bool
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

func process(objects []Object) {
	sort.Slice(objects, func(i, j int) bool { return objects[i].zMin < objects[j].zMin })
	endsAt := make(map[int][]*Object)

	for i := range objects {
		stable := false
		for z := objects[i].zMin - 1; z > 0; z-- {
			below := endsAt[z]
			for j := range below {
				if intersect(objects[i], *below[j]) {
					objects[i].zRealMin = z + 1
					objects[i].zRealMax = objects[i].zRealMin + (objects[i].zMax - objects[i].zMin)
					objects[i].standsOn = append(objects[i].standsOn, below[j])
					below[j].supports = append(below[j].supports, &objects[i])
					if !stable {
						endsAt[objects[i].zRealMax] = append(endsAt[objects[i].zRealMax], &objects[i])
					}
					stable = true
				}
			}

			if stable {
				break
			}
		}

		if !stable {
			objects[i].zRealMin = 1
			objects[i].zRealMax = objects[i].zRealMin + (objects[i].zMax - objects[i].zMin)
			endsAt[objects[i].zRealMax] = append(endsAt[objects[i].zRealMax], &objects[i])
		}
	}
}

func part1(objects []Object) int {
	var result int
	process(objects)
	for i := range objects {
		if len(objects[i].supports) == 0 {
			objects[i].canBeDeleted = true
			result++
			continue
		}

		canBeDeleted := true
		for j := range objects[i].supports {
			if len(objects[i].supports[j].standsOn) < 2 {
				canBeDeleted = false
				break
			}
		}

		if canBeDeleted {
			objects[i].canBeDeleted = true
			result++
		}
	}

	return result
}

func countFallen(object *Object, fallen map[*Object]bool) int {
	var count int
	ok, _ := fallen[object]
	if !ok {
		fallen[object] = true
		count++
	}

	for i := range object.supports {
		count += countFallen(object.supports[i], fallen)
	}

	return count
}

func part2(objects []Object) int {
	var result int
	for i := range objects {
		if objects[i].canBeDeleted {
			continue
		}

		fallen := make(map[*Object]bool)
		for j := range objects[i].supports {
			result += countFallen(objects[i].supports[j], fallen)
		}
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

	objects := readInput(file)
	fmt.Println("Part1:", part1(objects))
	fmt.Println("Part2:", part2(objects))
}
