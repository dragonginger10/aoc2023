package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
  var answer int
  schematic := ReadInput(os.Args[1])

  for _, l := range schematic {
    fmt.Println(l)
  }
  fmt.Println(answer)
}
