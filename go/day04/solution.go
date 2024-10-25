package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func exp(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 2
	}
	return result
}

func coincidences(line string) int {
	var start int
	var coincidences int
	for i := 0; i < len(line); i++ {
		if line[i] == ':' {
			start = i
			break
		}
	}
	l := line[start+1:]
	nums := strings.Split(l, "|")
	winningNums := strings.Fields(nums[0])
	myNums := strings.Fields(nums[1])

	for _, n := range myNums {
		for _, r := range winningNums {
			if n == r {
				coincidences++
			}
		}
	}
	return coincidences
}

func part1() (total int) {
	input, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		c := coincidences(line)
		if c == 0 {
			total += 0
		} else {
			total += exp(c - 1)
		}
	}
	return total
}

func part2() (total int) {
	input, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	copies := make([]int, len(lines))
	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}

	for _, line := range lines {
		c := coincidences(line)
		for i := 1; i < 1+c; i++ {
			for j := 0; j < copies[0]; j++ {
				copies[i]++
			}
		}
		total += copies[0]
		copies = copies[1:]
	}

	return total
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
