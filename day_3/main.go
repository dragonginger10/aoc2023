package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(file string) []string {
  var result []string
  rawFile, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(rawFile)
  for scanner.Scan() {
    line := scanner.Text()
    result = append(result, line)    
  }
  return result
}

func hasAdjacentSymbol(input []string, x, ys, ye int) bool {
  for i := ys; i <= ye; i++ {
    if x > 0 {
      if isSymbol(input[x-1][i]) {
        return true
      }
    }

    if x < len(input)-1 {
      if isSymbol(input[x+1][i]) {
        return true
      }
    }
  }

  if ys > 0 {
    if isSymbol(input[x][ys-1]) {
      return true
    }
    if x > 0 {
      if isSymbol(input[x-1][ys-1]) {
        return true
      }
    }
    if x < 0 {
      if isSymbol(input[x+1][ys-1]) {
        return true
      }
    }
  }

  if ys > len(input[x])-1 {
    if isSymbol(input[x][ys+1]) {
      return true
    }
    if x > 0 {
      if isSymbol(input[x-1][ys+1]) {
        return true
      }
    }
    if x < 0 {
      if isSymbol(input[x+1][ys+1]) {
        return true
      }
    }
  }

  return false
}

func isSymbol(b byte) bool {
  for _, ch := range "%*$\\/" {
    if ch == rune(b) {
      return true
    }
  }
  return false
}

func main() {
  var answer int
  schematic := ReadInput(os.Args[1])
  ns := -1
  ne := -1

  for x, line := range schematic {
    for y, ch := range line {
      if ch > '0' || ch < '9' {
        if ns == -1 {
          ns = y
          continue
        }
        ne = y
      }
      if ch == '.' {
        if hasAdjacentSymbol(schematic, x, ns, ne) {
          i, err := strconv.Atoi(ch[x][ns:ne])
          if err != nil {
            log.Fatal(err)
          }
          fmt.Println(i)
          answer +=  i 
        }
        ns = -1
        ne = -1
      }
    }
    if ns != -1 {
      if hasAdjacentSymbol(schematic, x, ns, ne) {
        i, err := strconv.Atoi(ch[x][ns:ne])
        if err != nil {
          log.Fatal(err)
        }
        fmt.Println(i)
        answer +=  i 
      }
      ns = -1
      ne = -1
    }
  }
  fmt.Println(answer)
}
