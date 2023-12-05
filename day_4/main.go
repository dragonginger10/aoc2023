package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	prsdFile := strings.Split(string(rawFile), "\n")
	return prsdFile[:len(prsdFile)-1]
}

func cleanNumbers(numbers []string) []int {
	var cleanedNumbers []int
	for _, v := range numbers {
		s := strings.TrimSpace(v)
		i, _ := strconv.ParseInt(s, 10, 0)
		if i != 0 {
			cleanedNumbers = append(cleanedNumbers, int(i))
		}
	}
	return cleanedNumbers
}

func numsInList(num int, list []int) int {
	count := 0
	for _, i := range list {
		if i == num {
			count++
		}
	}
	return count
}

func countPoints(card []int, wins []int) int {
	points := 0
	for _, i := range wins {
		if numsInList(i, card) != 0 {
			if points == 0 {
				points += 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func copyCards(crdNum int, wins []int, card []int) []int {
	count := 0
	var copiedCards []int
	for _, i := range wins {
		if numsInList(i, card) != 0 {
			count++
			copiedCards = append(copiedCards, count+crdNum)
		}
	}
	return copiedCards
}

func main() {
	cards := ReadInput(os.Args[1])
	var totalPts int
	var copiedCards []int

	for crdNum, card := range cards {
		crdNum++ //set card to index starting at 1
		numbers := strings.Split(strings.Split(card, ":")[1], "|")
		winning := cleanNumbers(strings.Split(numbers[0], " "))
		ownedNumbers := cleanNumbers(strings.Split(numbers[1], " "))
		totalPts += countPoints(ownedNumbers, winning)

		copiedCards = append(copiedCards, copyCards(crdNum, winning, ownedNumbers)...)
		copies := numsInList(crdNum, copiedCards)
		if copies > 0 {
			for i := 0; i < copies; i++ {
				copiedCards = append(copiedCards, copyCards(crdNum, winning, ownedNumbers)...)
			}
		}
	}

	fmt.Println("pt 1:", totalPts)
	fmt.Println("pt 2:", len(copiedCards)+len(cards))
}
