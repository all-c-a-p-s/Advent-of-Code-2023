package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	id    string
	left  string // id of node rather than pointer
	right string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseNode(s string) node {
	parts := strings.Split(s, "=")
	id := parts[0][:3]
	children := strings.Split(parts[1][2:len(parts[1])-1], ",")
	children[1] = strings.TrimSpace(children[1])
	return node{id, children[0], children[1]}
}

func scanNodes(lines []string) (nodes []node) {
	for _, line := range lines {
		nodes = append(nodes, parseNode(line))
	}
	return nodes
}

func find(target string, nodes []node) int {
	for i, n := range nodes {
		if n.id == target {
			return i
		}
	}
	panic("Failed to find target node")
}

func move(direction byte, current node, nodes []node) int {
	switch direction {
	case 'L':
		return find(current.left, nodes)
	case 'R':
		return find(current.right, nodes)
	default:
		panic("Failed to parse direction")
	}
}

func pathFind(directions []byte, nodes []node, n node, part int) int {
	var moveCount int
	var currentNode node
	if part == 1 { // used in part 1
		for _, node := range nodes {
			if node.id == "AAA" {
				currentNode = node
			}
		}
	} else {
		currentNode = n
	}
	for i := 0; i < len(directions); i++ {
		currentNode = nodes[move(directions[i], currentNode, nodes)]
		moveCount++
		if part == 1 && currentNode.id == "ZZZ" {
			break
		} else if part == 2 && currentNode.id[2] == 'Z' {
			break
		}
		if i == len(directions)-1 {
			i = -1 // then incremented to zero by loop
		}
	}
	return moveCount
}

func part1() int {
	input, err := os.Open("input.txt")
	check(err)

	var lines []string
	var dirs string
	scanningDirections := true

	scanner := bufio.NewScanner(input)
	var start, lineNum int
	for scanner.Scan() {
		lineNum++
		lines = append(lines, scanner.Text())
		if len(scanner.Text()) == 0 {
			start = lineNum
			scanningDirections = false
		}
		if scanningDirections {
			dirs += scanner.Text()
		}
	}

	directions := []byte(strings.TrimSpace(dirs))

	nodes := scanNodes(lines[start:])
	result := pathFind(directions, nodes, node{"AAA", "AAA", "AAA"}, 1)

	return result
}

func hcf(x, y int) int {
	if y == 0 {
		return x
	}
	return hcf(y, x%y)
}

func lcm(ns []int) int {
	if len(ns) == 0 {
		panic("Slice is empty")
	}
	res := ns[0]
	for i := 1; i < len(ns); i++ {
		res = (ns[i] * res) / hcf(ns[i], res)
	}
	return res
}

func pathFindP2(directions []byte, nodes []node) int {
	var startNodes []node
	for _, node := range nodes {
		if node.id[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	var cycles []int

	for i := 0; i < len(startNodes); i++ {
		cycles = append(cycles, pathFind(directions, nodes, startNodes[i], 2))
	}

	return lcm(cycles)
	// pretty surprising that this actually works since there are multiple exits for each cycle
	// but it must be that the input is designed for this to work
}

func part2() int {
	input, err := os.Open("input.txt")
	check(err)

	var lines []string
	var dirs string
	scanningDirections := true

	scanner := bufio.NewScanner(input)
	var start, lineNum int
	for scanner.Scan() {
		lineNum++
		lines = append(lines, scanner.Text())
		if len(scanner.Text()) == 0 {
			start = lineNum
			scanningDirections = false
		}
		if scanningDirections {
			dirs += scanner.Text()
		}
	}

	directions := []byte(strings.TrimSpace(dirs))

	nodes := scanNodes(lines[start:])
	result := pathFindP2(directions, nodes)

	return result
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
