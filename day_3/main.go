package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file := os.Args[1]

	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
