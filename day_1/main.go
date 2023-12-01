package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    file := os.Args[1]

    content, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }

    defer content.Close()

    scanner := bufio.NewScanner(content)
    var result int64

    for scanner.Scan() {
        line := scanner.Text()
        line = replaceNumberStrings(line)
        re := regexp.MustCompile("[0-9]")
        matches := re.FindAllString(line, -1)
        number, _ := strconv.ParseInt(matches[0] + matches[len(matches)-1], 10, 64)
        result += number
    }

    fmt.Println(result)
}

func replaceNumberStrings(line string) string {
    numbers := map[string]string{
        "one": "o1e",
        "two": "t2o",
        "three": "t3e",
        "four": "f4r",
        "five": "f5e",
        "six": "s6x",
        "seven": "s7n",
        "eight": "e8t",
        "nine": "n9e",
    }

    for k, v := range numbers {
        line = strings.ReplaceAll(line, k, v)
    }

    return line
}
