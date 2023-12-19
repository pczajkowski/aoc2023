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

func readInput(file *os.File) map[string]Workflow {
	scanner := bufio.NewScanner(file)
	workflows := make(map[string]Workflow)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		workflow := parseWorkflow(line)
		workflows[workflow.id] = workflow
	}

	return workflows
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

	workflows := readInput(file)
	fmt.Println(workflows)
}
