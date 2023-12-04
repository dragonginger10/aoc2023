package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadInput(file string) []string {
  rawFile, err := os.ReadFile(file)
  if err != nil {
    log.Fatal(err)
  }
  return strings.Split(string(rawFile), "\n")
}

func getMatrix(line []string) [][]rune {
  matrix := make([][]rune, len(line))

  for i := 0; i < len(line); i++ {
    matrix[i] = readChar(line[i])
  }

  return matrix
}

func readChar(line string) []rune {
  var converter []rune
  for _, runeValue := range line {
    converter = append(converter, runeValue)
  }
  return converter
}

func getAdjacentSymbol(row int, start_col int, end_col int, matrix [][]rune) ([]int, error){
  search_area := [][]int{{row, start_col-1}, {row, end_col}}
  for i := row-1; i < row+2; i++ {
    for j := start_col-1; j < end_col+1; j++ {
      search_area = append(search_area, []int{i, j})
    }
  }
  fmt.Println(search_area)

  for _, coordinate := range search_area {
    r := coordinate[0]
    c := coordinate[1]

    if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[r]) {
      continue
    }

    candidate := matrix[r][c]
    if ! unicode.IsDigit(candidate) && candidate != '.' {
      return []int{r, c}, nil
    }
  }
  return nil, errors.New("WTF")
}

func main() {
  var ans int
  var partNum []int
  grid := ReadInput(os.Args[1])
  matrix := getMatrix(grid)

  for _, i := range matrix {
    i = append(i, '.')
  }

  for r,row := range matrix {
    num := 0
    for c, col := range row {
      if unicode.IsDigit(col){
        num *= 10
        prsd, _ := strconv.ParseInt(string(col), 10, 0)
        num += int(prsd)
        continue
      }
      if num == 0 {
        continue
      }
      col_start := c - len(fmt.Sprintf("%d", num))
      fmt.Println(num, col_start, len(fmt.Sprintf("%d", num)))
      _, err := getAdjacentSymbol(r, col_start, c, matrix); if err == nil {
        partNum = append(partNum, num)
      }
      num = 0
    }
  }

  for _, num := range partNum {
    ans += num
  }

  fmt.Println(ans)
}
