package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(file string) []string {
	var result []string
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	result = strings.Split(string(rawFile), "\n")
	return result
}

func pushNorth(g []string) []string {
	var newGrid []string
	for r, row := range g {
		for c, ch := range row {

		}
	}
	return newGrid
}

func main() {
	var answer int
	grid := ReadInput(os.Args[1])
	fmt.Println(answer)
	for _, row := range grid {
		fmt.Println(row)
	}
}
