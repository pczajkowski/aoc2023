package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	isExpression bool
	greater      bool
	id           string
	value        int
	wrong        *Node
	right        *Node
}

type Workflow struct {
	id         string
	expression *Node
}

func parseExpression(text string) []Node {
	var nodes []Node
	if !strings.Contains(text, ",") {
		var node Node
		var parts []string
		if strings.Contains(text, ">") {
			node.greater = true
			parts = strings.Split(text, ">")
		} else if strings.Contains(text, "<") {
			parts = strings.Split(text, "<")
		} else {
			node.id = text
			nodes = append(nodes, node)
			return nodes
		}

		node.isExpression = true
		if len(parts) != 2 {
			log.Fatalf("Can't parse expression: %s", text)
		}

		node.id = parts[0]
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Can't parse a number %s: %s", parts[1], err)
		}
		node.value = n

		nodes = append(nodes, node)
	} else {
		parts := strings.Split(text, ",")
		if len(parts) != 2 {
			log.Fatalf("Wrong number of parts from: %s", text)
		}

		for i := range parts {
			if strings.ContainsAny(parts[i], "<>") {
				nodes = append(nodes, parseExpression(parts[i])...)
			} else {
				nodes = append(nodes, Node{id: parts[i]})
			}
		}
	}

	return nodes
}

func parseWorkflow(line string) Workflow {
	var workflow Workflow
	idAndExpression := strings.Split(line, "{")
	if len(idAndExpression) != 2 {
		log.Fatalf("Can't extract id and expression from: %s", line)
	}

	workflow.id = idAndExpression[0]
	expressions := strings.Split(strings.TrimRight(idAndExpression[1], "}"), ":")
	if len(expressions) < 2 {
		log.Fatalf("Can't extract expressions from: %s", idAndExpression[1])
	}

	first := parseExpression(expressions[0])
	if len(first) != 1 {
		log.Fatalf("First expression should be single one: %s", expressions[0])
	}

	workflow.expression = &first[0]
	current := workflow.expression
	for i := 1; i < len(expressions); i++ {
		ex := parseExpression(expressions[i])
		if len(ex) != 2 {
			log.Fatalf("Need two: %s", expressions[i])
		}

		current.right = &ex[0]
		if current.right.isExpression {
			current = current.right
		}

		current.wrong = &ex[1]
		if current.wrong.isExpression {
			current = current.wrong
		}
	}

	return workflow
}

type Rating struct {
	x, m, a, s int
}

func readInput(file *os.File) (map[string]Workflow, []Rating) {
	scanner := bufio.NewScanner(file)
	workflows := make(map[string]Workflow)
	var ratings []Rating
	readingRatings := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if !readingRatings {
				readingRatings = true
				continue
			}

			break
		}

		if !readingRatings {
			workflow := parseWorkflow(line)
			workflows[workflow.id] = workflow
		} else {
			var rating Rating
			n, err := fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &rating.x, &rating.m, &rating.a, &rating.s)
			if n != 4 || err != nil {
				log.Fatalf("Bad input for rating: %s\n%s", line, err)
			}
			ratings = append(ratings, rating)
		}
	}

	return workflows, ratings
}

func (n *Node) test(value int) *Node {
	if n.greater {
		if value > n.value {
			return n.right
		}
	} else {
		if value < n.value {
			return n.right
		}
	}

	return n.wrong
}

func sortRatings(workflows map[string]Workflow, ratings []Rating) ([]Rating, []Rating) {
	var accepted []Rating
	var rejected []Rating
	for i := range ratings {
		current := workflows["in"].expression
		for {
			if current == nil {
				break
			}
			if current.isExpression {
				switch current.id {
				case "x":
					current = current.test(ratings[i].x)
				case "m":
					current = current.test(ratings[i].m)
				case "a":
					current = current.test(ratings[i].a)
				case "s":
					current = current.test(ratings[i].s)
				default:
					fmt.Println(current.id, "No match!")
				}
			} else {
				if current.id == "A" {
					accepted = append(accepted, ratings[i])
					break
				} else if current.id == "R" {
					rejected = append(rejected, ratings[i])
					break
				} else {
					current = workflows[current.id].expression
				}
			}
		}
	}

	return accepted, rejected
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

	workflows, ratings := readInput(file)
	accepted, rejected := sortRatings(workflows, ratings)
	fmt.Println(accepted, rejected)
}
