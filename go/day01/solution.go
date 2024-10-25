package main 

import (
  "fmt"
  "os"
  "bufio"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}
func part1() (total int) {
  var numbers []int
  digits := map[byte]int{
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
  }

  input, err := os.Open("input.txt")
  check(err)
  scanner := bufio.NewScanner(input)

  for scanner.Scan() {
    line := scanner.Text()
    var number string
    for i := 0;i < len(line);i++{
      if _, ok := digits[line[i]]; ok {
        number += string(line[i])
      }
    }
    n := digits[number[0]] * 10 + digits[number[len(number)-1]]
    numbers = append(numbers, n)
     
  }
  for _, r := range numbers {
    total += r
  }
  return total
}

func part2() (total int) {

  digits := map[string]int {
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,

    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
  }

  input, err := os.Open("input.txt")
  check(err)

  var calibrations []int

  scanner := bufio.NewScanner(input)

  for scanner.Scan() {
    line := scanner.Text()
    min := len(line)
    max := -1
    var first int
    var last int
    for key := range digits {
      for i := 0;i <= len(line) - len(key);i++{
        slice := line[i:i+len(key)]
        if a, ok := digits[slice]; ok && i < min && slice == key {
          first = a
          min = i
        }
      }

      for i := len(line) - len(key);i >= 0;i--{
        slice := line[i:i+len(key)]
        if b, ok := digits[slice]; ok && i > max && slice == key {
          last = b
          max = i
        }
      }
    }


    calibration := first * 10 + last
    calibrations = append(calibrations, calibration)
  }

  for _, r := range calibrations {
    total += r
  }
  return total
}

func main(){
  fmt.Printf("Part 1: %d \n", part1())
  fmt.Printf("Part 2: %d \n", part2())
}
