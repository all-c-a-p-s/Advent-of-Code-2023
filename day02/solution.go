package main 

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func part1() (sum int) {
  input, err := os.Open("input.txt")
  check(err)

  scanner := bufio.NewScanner(input) 
  for scanner.Scan() {
    var blue int 
    var red int
    var green int

    var index int
    var gamesStr string
    
    line := scanner.Text()
    for i := 0; i < len(line);i++{
      if line[i] == ':' {        
        gamesStr = line[i+1:]
        indexStr := line[5:i] 
        i, err := strconv.Atoi(indexStr)
        check(err)
        index = i
        break
      }
      
    }

    games := strings.Split(gamesStr, ";")
    for _, game := range games {
      draws := strings.Split(game, ", ")
      for _, draw := range draws {
        fields := strings.Fields(draw)

        n, err := strconv.Atoi(fields[0])
        check(err)
      
        switch fields[1] {
        case "blue":
          blue = max(blue, n)
        case "red":
          red = max(red, n)
        case "green": 
          green = max(green, n)
        default:
          panic("Error parsing colour")
       } 
      }
    } 
    if red <= 12 && green <= 13 && blue <= 14{
      sum += index
    }

  }
  return sum
}

func part2() (sum int) {
  input, err := os.Open("input.txt")
  check(err)

  scanner := bufio.NewScanner(input) 
  for scanner.Scan() {
    var blue int 
    var red int
    var green int 
    
    var gamesStr string
    
    line := scanner.Text()
    for i := 0; i < len(line);i++{
      if line[i] == ':' {        
        gamesStr = line[i+1:] 
        break
      }
      
    }

    games := strings.Split(gamesStr, ";")
    for _, game := range games {
      draws := strings.Split(game, ", ")
      for _, draw := range draws {
        fields := strings.Fields(draw)

        n, err := strconv.Atoi(fields[0])
        check(err)
      
        switch fields[1] {
        case "blue":
          blue = max(blue, n)
        case "red":
          red = max(red, n)
        case "green": 
          green = max(green, n)
        default:
          panic("Error parsing colour")
       } 
      }
    } 
    power := red * blue * green
    sum += power

  }
  return sum
}

func main() {
  fmt.Printf("Part 1: %d \n", part1())
  fmt.Printf("Part 2: %d \", part2())
}
