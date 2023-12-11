package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
	var answer int
	fmt.Println(answer)
}
