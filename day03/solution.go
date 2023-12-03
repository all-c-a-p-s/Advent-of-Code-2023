package main 

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

type charType int



const (
  digit charType = iota
  period
  symbol
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func getType(char byte) charType {
  if char >= 48 && char <= 57 {
    return digit
  } else if char == 46 {
    return period
  }
  return symbol
}

func adjacentSymbol(lines []string, vertical, horizontal int) bool {
  if getType(lines[vertical][horizontal]) == symbol {
    return true
  }
  return false
}

func part1() (total int){
  input, err := os.Open("input.txt")
  check(err)

  scanner := bufio.NewScanner(input) 
  var lines []string
  
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  for lineNum, line := range lines {
    var numStr string //current number 
 
    for i := 0;i < len(line);i++{ 
      if getType(line[i]) == digit { //digit
        numStr += string(line[i])
      }
      if getType(line[i]) != digit || i == len(line) - 1 { //end of number or end of line reached 
        if len(numStr) != 0 { //there was a number and not just symbols/periods before current index
          numStart := i - (len(numStr) - 1) - 1
          numEnd := i - 1 
          var top, left, right, bottom bool
          if lineNum == 0 {
            top = true
          } else if lineNum == len(lines) -1 {
            bottom = true
          }
          if numStart == 0 {
            left = true
          } 
          if numEnd == len(lines[lineNum]) -1 {
            right = true
          }

          var ok bool 

          for j := numStart; j <= numEnd;j++{
            if !top && !left { //north west
              ok = adjacentSymbol(lines, lineNum - 1, j-1) 
              if ok {
                break
              }
            }
            if !top {//north
              ok = adjacentSymbol(lines, lineNum - 1, j) 
              if ok {
                break
              } 
            }

            if !top && !right { //north east
              ok = adjacentSymbol(lines, lineNum - 1, j+1) 
              if ok {
                break
              } 
            }

            if !right {//east
              ok = adjacentSymbol(lines, lineNum, j+1) 
              if ok {
                break
              }
 
            }

            if !right && !bottom { //south east
              ok = adjacentSymbol(lines, lineNum + 1, j+1) 
              if ok {
                break
              }

            }

            if !bottom { //south
              ok = adjacentSymbol(lines, lineNum + 1, j) 
              if ok {
                break
              }

            }     

            if !bottom && !left { //south west
              ok = adjacentSymbol(lines, lineNum + 1, j-1) 
              if ok { 
                break
              }
            }

            if !left { //west
              ok = adjacentSymbol(lines, lineNum, j-1) 
              if ok {
                break
              }

            }
          }

          n, err := strconv.Atoi(numStr)
          check(err)
  
          if ok {
            total += n
          }
        }
        numStr = ""
      }

    }
  }
  return total
}

func readNum(line string, index int) int {
  var numStr string
  var numStart int = index
  for i := index-1;i >= 0;i--{
    if getType(line[i]) == digit {
      numStart=i
    } else {
      break
    }
  }
  for i := numStart;i < len(line);i++{
    if getType(line[i]) == digit {
      numStr += string(line[i])
    } else {
      break
    }
  } 
  n, err := strconv.Atoi(numStr)
  check(err)
  return n
}

func part2() (total int){
  input, err := os.Open("input.txt")
  check(err)

  scanner := bufio.NewScanner(input) 
  var lines []string
  
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  for lineNum, line := range lines {
 
    for i := 0;i < len(line);i++{ 
      if line[i] == '*' {
        var adjacentNums map[int]bool = map[int]bool {}
        var top, left, right, bottom bool
        if lineNum == 0 {
          top = true
        } else if lineNum == len(lines) -1 {
          bottom = true
        }
        if i == 0 {
          left = true
        } 
        if i == len(lines[lineNum]) - 1 {
          right = true
        }

        if !top && !left {
          if getType(lines[lineNum - 1][i-1]) == digit { 
            adjacentNums[readNum(lines[lineNum-1], i-1)] = true
          }
        }
        if !top {
          if getType(lines[lineNum - 1][i]) == digit {
            adjacentNums[readNum(lines[lineNum-1], i)] = true
          }
        }
        if !top && !right {
          if getType(lines[lineNum - 1][i+1]) == digit {
            adjacentNums[readNum(lines[lineNum-1], i+1)] = true
          }
        }
        if !right {
          if getType(lines[lineNum][i+1]) == digit {
            adjacentNums[readNum(lines[lineNum], i+1)] = true
          }
        }
        if !right && !bottom {
          if getType(lines[lineNum + 1][i+1]) == digit {
            adjacentNums[readNum(lines[lineNum+1], i+1)] = true
          }
        }
        if !bottom {
          if getType(lines[lineNum + 1][i]) == digit {
            adjacentNums[readNum(lines[lineNum+1], i)] = true
          }
        }
        if !bottom && !left {
          if getType(lines[lineNum + 1][i-1]) == digit {
            adjacentNums[readNum(lines[lineNum+1], i-1)] = true
          }
        }
        if !left {
          if getType(lines[lineNum][i-1]) == digit {
            adjacentNums[readNum(lines[lineNum], i-1)] = true
          }
        } 
        if len(adjacentNums) == 2{
          n := 1
          for r := range adjacentNums{
            n *= r
          } 
          total += n
        }
      } 
    }
  }
  return total
}

func main () {
  fmt.Printf("Part 1: %d \n", part1())
  fmt.Printf("Part 2: %d \n", part2())
}
