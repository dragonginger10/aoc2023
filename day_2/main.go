package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := os.Args[1]
  m := map[string]int64{
    "red": 12,
    "green": 13,
    "blue": 14,
  }
  var ans int
  var game int

	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
    line = strings.Split(line, ":")[1]
    ok := true
    game += 1

    for _, event := range strings.Split(line, ";") {
      for _, cubes := range strings.Split(event, ",") {
        colors := strings.Split(cubes, " ")
        n := colors[1]
        color := colors[2]
        i, _ := strconv.ParseInt(n, 10, 64)

        if i > m[color] {
          ok = false
        }
      }
    }
    if ok {
      ans += game
    }
  }
  fmt.Println(ans)
}
