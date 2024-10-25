package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func format(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			return s[i:]
		}
	}
	panic("No colon found in line")
}

func calculateDistance(heldDown, totalTime int) int {
	return heldDown * (totalTime - heldDown)
}

func part1() int {
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 1

	timesStr := strings.Fields(strings.TrimSpace(format(lines[0])))
	distancesStr := strings.Fields(strings.TrimSpace(format(lines[1])))

	var times, distances []int

	for i := 0; i < len(timesStr); i++ { // works because same number of times and distances
		t, err := strconv.Atoi(timesStr[i])
		check(err)
		d, err := strconv.Atoi(distancesStr[i])
		check(err)

		times = append(times, t)
		distances = append(distances, d)
	}

	for i := 0; i < len(times); i++ {
		waysToWin := 0
		for t := 0; t < times[i]; t++ { // possible times
			distance := calculateDistance(t, times[i])
			if distance > distances[i] {
				waysToWin++
			}
		}
		total *= waysToWin
	}
	return total
}

func part2() (total int) {
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	timesStr := strings.Fields(strings.TrimSpace(format(lines[0])))
	distancesStr := strings.Fields(strings.TrimSpace(format(lines[1])))

	var rt, rd string

	for i := 0; i < len(timesStr); i++ { // works because same number of times and distances
		rt += timesStr[i]
		rd += distancesStr[i]
	}

	raceTime, err := strconv.Atoi(rt)
	check(err)
	raceDistance, err := strconv.Atoi(rd)
	check(err)

	for i := 0; i < raceTime; i++ {
		if calculateDistance(i, raceTime) > raceDistance {
			total++
		}
	}
	return total
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
