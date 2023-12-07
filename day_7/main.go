package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handInfo struct {
    hand string
    handNum []int
    handValue int
    bid int
    ranking int
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

func parseHandValue(hand string) int {
	var handValue int
	counts := make(map[rune]int, 0)

	for _, card := range hand {
        counts[card]++
	}

    num := 0
    for _, count := range counts {
        if count > num {
            num = count
        }
    }

	for card, count := range counts {
        if count == num && card != 'J'{
            count += counts['J']
            num = 0
        }

		switch count {
		case 5:
			handValue = 6
			break
		case 4:
			handValue = 5
			break
		case 3:
			handValue += 3
		case 2:
			handValue++
		default:
			continue
		}
	}
	return handValue
}

func cardToInt(card rune) int {
	cardValue := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 1,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

    return cardValue[card]    
}

func handToInt(hand string) []int {
    var list []int

    for _, card := range hand {
        list = append(list, cardToInt(card))
    }

    return list
}

func main() {
	var answer int64
    var bid int64
    var handRanking []handInfo 
	listOfHands := ReadInput(os.Args[1])

	for i, line := range listOfHands {
        i++
        if line == "" {
            break
        }
		handBid := strings.Split(line, " ")
		hand := handBid[0]
		bid, _ = strconv.ParseInt(handBid[1], 10, 64)

        handRanking = append(handRanking, handInfo{
            hand: hand,
            handNum: handToInt(hand),
            handValue: parseHandValue(hand),
            bid: int(bid),
        })
	}

    sort.Slice(handRanking, func(i, j int) bool {
        if handRanking[i].handValue == handRanking[j].handValue {
            for x := 0; x < len(handRanking[i].handNum); x++ {
                if handRanking[i].handNum[x] != handRanking[j].handNum[x]{
                    return handRanking[i].handNum[x] < handRanking[j].handNum[x]
                }
            }
        }
        return handRanking[i].handValue < handRanking[j].handValue
    })


    for rank, h := range handRanking {
        rank++
        // fmt.Printf("Hand: %s\nBid: %d\nHand Value %d\nRank: %d\n", h.hand, h.bid, h.handValue, rank)
        answer += int64(rank) * int64(h.bid)
    }

        
	fmt.Println(answer)
}
