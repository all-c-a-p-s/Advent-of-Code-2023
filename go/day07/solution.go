package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	card int
	rank int
	hand struct {
		cards []byte
		bid   int
		score int
	}
)

func convert(char byte) int {
	m := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return m[char]
}

func convertP2(char byte) int {
	m := map[byte]int{
		'J': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	return m[char]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func exp(n int) int { // used for hex bits
	res := 1
	for i := 0; i < n; i++ {
		res *= 16
	}
	return res
}

func parseHand(s string) hand {
	f := strings.Fields(s)
	cards := []byte(f[0])
	str := f[1]
	bid, err := strconv.Atoi(str)
	check(err)
	return hand{cards, bid, 0}
}

func handRank(hand []byte) (score int) {
	distinctCards := map[byte]int{} // hashSet
	for _, c := range hand {
		if _, ok := distinctCards[c]; !ok {
			distinctCards[c] = 1
		} else {
			distinctCards[c]++
		}
	}

	l := len(distinctCards)
	switch l {
	case 1: // five of a kind only
		score += 6 * 0x100_000
	case 2: // full house or four of a kind
		quads := false
		for _, v := range distinctCards {
			if v == 4 { // four of a kind
				quads = true
			}
		}
		if quads {
			score += 5 * 0x100_000
		} else {
			score += 4 * 0x100_000
		}
	case 3: // three of a kind or two pair
		trips := false
		for _, v := range distinctCards {
			if v == 3 {
				trips = true
			}
		}
		if trips {
			score += 3 * 0x100_000
		} else {
			score += 2 * 0x100_000
		}
	case 4: // one pair
		score += 1 * 0x100_000
	default: // high card
		score += 0
	}
	return score
}

func scoreP1(hand []byte) (score int) { // scores a hand to a hex number where first hex bit is hand strength, next 5 bits are card scores
	score += handRank(hand)
	for i := 0; i < len(hand); i++ { // give values to cards in hand
		weight := exp(4 - i)
		c := convert(hand[i])
		score += weight * c
	}
	return score
}

func totalWinnings(s []string) (total int) {
	var hands []hand
	for i := 0; i < len(s); i++ {
		hands = append(hands, parseHand(s[i]))
		hands[i].score = scoreP1(hands[i].cards)
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score > hands[j].score
	})

	for i := 0; i < len(s); i++ {
		total += hands[i].bid * (len(s) - i)
	}
	return total
}

func part1() (total int) {
	input, err := os.Open("input.txt")
	check(err)

	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total = totalWinnings(lines)
	return total
}

func cards() []byte {
	return []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
}

func scoreP2(hand []byte) (score int) {
	for i := 0; i < len(hand); i++ { // give values to cards in hand before permuting
		weight := exp(4 - i)
		c := convertP2(hand[i])
		score += weight * c
	}
	var jokers []int
	handRankScore := handRank(hand) // initalise for all hands
	for i := 0; i < len(hand); i++ {
		if hand[i] == 'J' {
			jokers = append(jokers, i)
		}
	}

	if len(jokers) != 0 { // permute through joker combinations
		currentJoker := len(jokers) - 1 // index of joker currently being permuted
		for {
			handRankScore = max(handRankScore, handRank(hand))
			done := true
			for i := 0; i < len(jokers); i++ {
				if hand[jokers[i]] != 'A' {
					done = false
				}
			}
			if done {
				break
			}
			if hand[jokers[currentJoker]] == 'A' { // cannot be incremented
				if currentJoker == 0 {
					currentJoker++
				} else {
					currentJoker--
				}
			} else { // can be incremented
				hand[jokers[currentJoker]] = cards()[convertP2(hand[jokers[currentJoker]])] // next card
				for i := currentJoker + 1; i < len(jokers); i++ {
					hand[jokers[i]] = 'J'
				}

				if currentJoker != len(jokers)-1 { // not last joker
					currentJoker = len(jokers) - 1 // go to last joker
				}
			}
		}
	}
	score += handRankScore
	return score
}

func totalWinningsP2(s []string) (total int) {
	var hands []hand
	for i := 0; i < len(s); i++ {
		hands = append(hands, parseHand(s[i]))
		hands[i].score = scoreP2(hands[i].cards)
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score > hands[j].score
	})

	for i := 0; i < len(s); i++ {
		total += hands[i].bid * (len(s) - i)
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

	total = totalWinningsP2(lines)
	return total
}

func main() {
	fmt.Printf("Part 1: %d \n", part1())
	fmt.Printf("Part 2: %d \n", part2())
}
