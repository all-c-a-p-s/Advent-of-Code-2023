package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// I really need to get better at using recursion, it would have been ideal for this problem,
// but due to my lack of familiarity using recursion I had to use iteration instead
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func differences(s []int) (difs []int) {
	if len(s) == 1 {
		return []int{0}
	}
	for i := 0; i < len(s)-1; i++ {
		difs = append(difs, s[i+1]-s[i])
	}
	return difs
}

func history(s []int) [][]int {
	currentDifs := s
	history := [][]int{s}
	for {
		done := true
		for i := 0; i < len(currentDifs); i++ {
			if currentDifs[i] != 0 {
				done = false
			}
		}
		if done {
			break
		}
		next := differences(currentDifs)
		history = append(history, next)
		currentDifs = next
	}
	return history
}

func extrapolate(history [][]int) [][]int {
	result := history
	for i := len(history) - 1; i > 0; i-- { // loop backwards through history
		x := history[i][len(history[i])-1]
		y := history[i-1][len(history[i-1])-1]
		result[i-1] = append(result[i-1], x+y)
	}

	for i := len(history) - 1; i > 0; i-- {
		x := history[i][0]
		y := history[i-1][0]
		result[i-1] = append([]int{y - x}, result[i-1]...)
	}
	return result
}

func parseLine(line string) (sequence []int) {
	numsStr := strings.Fields(line)
	for _, str := range numsStr {
		n, err := strconv.Atoi(str)
		check(err)
		sequence = append(sequence, n)
	}
	return sequence
}

func part1() (total int) {
	input, err := os.Open("input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sequences [][]int
	for _, line := range lines {
		sequences = append(sequences, parseLine(line))
	}

	var extrapolated [][][]int
	for _, seq := range sequences {
		t := extrapolate(history(seq))
		extrapolated = append(extrapolated, t)
	}

	for _, extr := range extrapolated {
		total += extr[0][len(extr[0])-1]
	}
	return total
}

func part2() (total int) {
	input, err := os.Open("input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sequences [][]int
	for _, line := range lines {
		sequences = append(sequences, parseLine(line))
	}

	var extrapolated [][][]int
	for _, seq := range sequences {
		t := extrapolate(history(seq))
		extrapolated = append(extrapolated, t)
	}

	for _, extr := range extrapolated {
		total += extr[0][0]
	}
	return total
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
