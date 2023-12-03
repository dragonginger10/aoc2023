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
  var ans2 int64

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
    bag := make(map[string]int64)

    for _, event := range strings.Split(line, ";") {
      for _, cubes := range strings.Split(event, ",") {
        colors := strings.Split(cubes, " ")
        n := colors[1]
        color := colors[2]
        i, _ := strconv.ParseInt(n, 10, 64)
        // fewest cubes in bag to get this game
        j, prs := bag[color]
        if prs {
          if j < i {
            bag[color] = i
          }
        } else {
          bag[color] = i
        }

        if i > m[color] {
          ok = false
        }
      }
    }
    if ok {
      ans += game
    }
    power := bag["blue"] * bag["red"] * bag["green"]
    ans2 += power
  }
  fmt.Println("sum of games playable:",ans)
  fmt.Println("Sum of powers of bags:",ans2)
}
