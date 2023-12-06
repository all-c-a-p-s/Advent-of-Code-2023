package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

type Map struct {
	source      Range
	destination Range
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func inputSeeds(start, end int, lines []string) []int { // starting and ending line numbers
	var seeds string
	for i := start; i <= end; i++ {
		seeds += lines[i]
	}

	numsStr := strings.Split(seeds[7:], " ")
	var nums []int

	for _, num := range numsStr {
		n, err := strconv.Atoi(num)
		check(err)
		nums = append(nums, n)
	}
	return nums
}

func readMap(start, end int, lines []string) []Map {
	var m []Map
	var numLists [][]int
	for i := start; i <= end; i++ {
		str := lines[i]
		ns := strings.Fields(str)
		var nums []int
		for _, n := range ns {
			num, err := strconv.Atoi(n)
			check(err)
			nums = append(nums, num)
		}
		numLists = append(numLists, nums)
	}

	for _, list := range numLists {
		d := Range{list[0], list[0] + list[2] - 1}
		s := Range{list[1], list[1] + list[2] - 1}
		newMap := Map{s, d}
		m = append(m, newMap)
	}
	return m
}

func part1() int {
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)

	var lines []string
	maps := [][]Map{}
	var firstMap int // linenum of heading of first map
	var seeds []int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		if lines[i] == "seed-to-soil map:" {
			firstMap = i
			seeds = inputSeeds(0, firstMap-2, lines)
		}
		if lines[i][len(lines[i])-4:] == "map:" {
			var end int
			for j := i + 1; i < len(lines); j++ {
				if j == len(lines) {
					end = j - 1
					break
				}
				if len(lines[j]) == 0 {
					end = j - 1
					break
				}
			}
			maps = append(maps, readMap(i+1, end, lines))
		}
	}

	var location []int

	for _, seed := range seeds {
		key := seed
		for _, m := range maps { // loop though 2d slice of maps
			for _, r := range m { // current slice
				dif := r.destination.min - r.source.min
				if key >= r.source.min && key <= r.source.max {
					key += dif
					break
				}

			}
		}
		location = append(location, key)
	}

	return slices.Min(location)
}

func findIntersection(r1, r2 Range) Range {
	return Range{max(r1.min, r2.min), min(r1.max, r2.max)}
}

func subtractRange(r1, r2 Range) []Range { // subtract r2 from r1
	if r1 == findIntersection(r1, r2) { // if they fully intersect
		return []Range{}
	}
	var subtracted []Range
	intersection := findIntersection(r1, r2)
	if intersection.min < intersection.max { // if there is intersection
		if r1.min < r2.min {
			subtracted = append(subtracted, Range{r1.min, r2.min - 1})
		}
		if r1.max > r2.max {
			subtracted = append(subtracted, Range{r2.max + 1, r1.max})
		}
	} else {
		return []Range{r1}
	}
	return subtracted
}

func part2() int {
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)

	var lines []string
	maps := [][]Map{}
	var firstMap int // linenum of heading of first map
	var seeds []int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		if lines[i] == "seed-to-soil map:" {
			firstMap = i
			seeds = inputSeeds(0, firstMap-2, lines)
		}
		if lines[i][len(lines[i])-4:] == "map:" {
			var end int
			for j := i + 1; i < len(lines); j++ {
				if j == len(lines) {
					end = j - 1
					break
				}
				if len(lines[j]) == 0 {
					end = j - 1
					break
				}
			}
			maps = append(maps, readMap(i+1, end, lines))
		}
	}

	var location []int
	var seedRanges []Range

	for i := 0; i < len(seeds)-1; i += 2 {
		r := Range{seeds[i], seeds[i] + seeds[i+1] - 1}
		seedRanges = append(seedRanges, r)
	}

	var locationRanges []Range

	for _, sr := range seedRanges {
		inputRanges := []Range{sr}
		for _, ml := range maps { // loop through 2d slice of maps
			var outputRanges []Range // all possible ranges of outputs from current layer
			unmapped := inputRanges  // values which have not yet been mapped
			for _, m := range ml {   // current map
				dif := m.destination.min - m.source.min
				for i := 0; i < len(unmapped); i++ { // go through unmapped ranges to see if any can be mapped

					intersection := findIntersection(unmapped[i], m.source)
					current := unmapped[i]
					if (intersection.min <= intersection.max) && (intersection.max > 0) { // there is intersection
						if i == len(unmapped)-1 { // last element
							s := subtractRange(unmapped[i], m.source)
							unmapped = unmapped[:len(unmapped)-1]
							i-- // decrement i so next index doesn't get messed up
							unmapped = append(unmapped, s...)
						} else {
							s := subtractRange(unmapped[i], m.source)
							unmapped = append(unmapped[:i], unmapped[i+1:]...) // remove intersection from unmapped slice
							i--                                                // as above
							unmapped = append(unmapped, s...)
						}
						if len(unmapped) == 0 {
							outputRange := Range{findIntersection(Range{}, m.source).min + dif, findIntersection(Range{}, m.source).max + dif}
							outputRanges = append(outputRanges, outputRange)
						} else {
							outputRange := Range{findIntersection(current, m.source).min + dif, findIntersection(current, m.source).max + dif}
							outputRanges = append(outputRanges, outputRange)
						}
					}
					if len(unmapped) == 0 {
						break
					}
				}
			}
			outputRanges = append(outputRanges, unmapped...) // ranges that fit in no maps

			inputRanges = outputRanges

		}
		locationRanges = append(locationRanges, inputRanges...)
	}

	for _, r := range locationRanges {
		location = append(location, r.min) // append lowest value for each range
	}

	return slices.Min(location)
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
