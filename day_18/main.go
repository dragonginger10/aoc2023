package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) []string {
	var result []string
	text, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()
	s := bufio.NewScanner(text)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}

func parseToInt(s string) int {
	number, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(number)
}

func shoelaceFormula(p [][2]int) float64 {
	var a float64
	for i := range p {
		prev := 0
		next := (i + 1) % len(p)
		if i == 0 {
			prev = len(p) - 1
		} else {
			prev = i - 1
		}
		a += float64(p[i][0] * (p[prev][1] - p[next][1]))
	}
	a = math.Abs(a) / 2
	return a
}

func picksTheorem(i, b float64) float64 {
	var a float64
	a = i - (b / 2) + 1
	return a
}

func part2(plan []string) []string {
	var newPlans []string
	directions := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}
	for _, line := range plan {
		hexCode := strings.Split(line, " ")[2]
		hexSlice := strings.Split(hexCode, "")
		dirNum := hexSlice[len(hexSlice)-2]
		dir := directions[string(dirNum)]
		dist := strings.Join(hexSlice[2:7], "")
		newLine := fmt.Sprintf("%v %s", dir, dist)
		newPlans = append(newPlans, newLine)
	}
	return newPlans
}

func main() {
	plan := ReadInput(os.Args[1])
	direction := map[string][2]int{
		"U": {-1, 0},
		"R": {0, 1},
		"D": {1, 0},
		"L": {0, -1},
	}
	points := [][2]int{{0, 0}}
	b := 0
	plan = part2(plan)
	for _, line := range plan {
		d := strings.Split(line, " ")[0]
		ch := strings.Split(line, " ")[1]
		n := parseToInt(ch)
		b += n
		dr := direction[d][0]
		dc := direction[d][1]
		r := points[len(points)-1][0]
		c := points[len(points)-1][1]
		points = append(points, [2]int{
			r + dr*n,
			c + dc*n,
		})
	}
	a := shoelaceFormula(points)
	answer := picksTheorem(a, float64(b))
	answer += float64(b)
	fmt.Println(int(answer))
}
