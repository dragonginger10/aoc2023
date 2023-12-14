package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func (c *coordinates) isEmpty() bool {
	return c != &coordinates{0, 0}
}

func runeInString(r rune, s string) bool {
	for _, b := range s {
		if b == r {
			return true
		}
	}
	return false
}

func coordInList(c coordinates, list []coordinates) bool {
	for _, l := range list {
		if l == c {
			return true
		}
	}
	return false
}

func getStartingType(g []string, s coordinates) (rune, error) {
	var connected string
	directions := map[string]coordinates{
		"N|7F": coordinates{-1, 0},
		"E-7J": coordinates{0, 1},
		"S|JL": coordinates{1, 0},
		"W-LF": coordinates{0, -1},
	}
	for k, v := range directions {
		pipe := g[s.x+v.x][s.y+v.y]

		if runeInString(rune(pipe), k) {
			connected += string(k[0])
		}
	}

	switch connected {
	case "NE", "EN":
		return 'L', nil
	case "NW", "WN":
		return 'J', nil
	case "NS", "SN":
		return '|', nil
	case "SW", "WS":
		return '7', nil
	case "SE", "ES":
		return 'F', nil
	case "WE", "EW":
		return '-', nil
	}
	return 'x', errors.New("How did we get here")
}

func getStartPoint(g []string) coordinates {
	var startCrd coordinates
	for r, row := range g {
		for c, ch := range row {
			if ch == 'S' {
				startCrd = coordinates{x: c, y: r}
				break
			}
		}
		if startCrd.isEmpty() {
			continue
		}
		break
	}
	return startCrd
}

func getSteps(r rune) []coordinates {
	result := make([]coordinates, 0)
	directions := map[string]coordinates{
		"|LJ": coordinates{-1, 0},
		"-LF": coordinates{0, 1},
		"|7F": coordinates{1, 0},
		"-J7": coordinates{0, -1},
	}

	for k, v := range directions {
		if runeInString(r, k) {
			result = append(result, v)
		}
	}

	return result
}

func nextStep(g []string, loc coordinates, passed []coordinates) coordinates {
	var c coordinates
	steps := getSteps(rune(g[loc.x][loc.y]))
	for _, s := range steps {
		c = coordinates{
			s.x + loc.x,
			s.y + loc.y,
		}
		if coordInList(c, passed) {
			continue
		} else {
			break
		}
	}
	return c
}

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
	answer := 1
	var passed []coordinates
	grid := ReadInput(os.Args[1])
	startCrd := getStartPoint(grid)
	passed = append(passed, startCrd)
	startingCh, err := getStartingType(grid, startCrd)
	if err != nil {
		log.Fatal(err)
	}

	possibleSteps := getSteps(startingCh)
	stepA := coordinates{
		possibleSteps[0].x + startCrd.x,
		possibleSteps[0].y + startCrd.y,
	}
	stepB := coordinates{
		possibleSteps[1].x + startCrd.x,
		possibleSteps[1].y + startCrd.y,
	}
	passed = append(passed, stepA, stepB)

	for stepA != stepB {
		answer++
		log.Println(stepA, stepB)
		stepA = nextStep(grid, stepA, passed)
		stepB = nextStep(grid, stepB, passed)
		passed = append(passed, stepB, stepA)
	}

	fmt.Println(startingCh, string(startingCh))
	fmt.Println(startCrd)
	fmt.Println(answer)
}
