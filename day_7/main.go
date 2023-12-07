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
  defer rawFile.Close()

  scanner := bufio.NewScanner(rawFile)
  for scanner.Scan() {
    readLine := scanner.Text()
    result = append(result, readLine)
  }
  return result
}

func main() {
    var answer int

    fmt.Println(answer)
}
