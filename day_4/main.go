package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(rawFile), "\n")
}
func main() {
	cards := ReadInput(os.Args[1])

	fmt.Println(cards)
}
