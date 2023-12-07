package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type handInfo struct {
    hand string
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

	for _, count := range counts {
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

func rankHand(h *handInfo, list []handInfo) []handInfo {
	cardValue := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
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

    for i, hand := range list {
        if hand.handValue == h.handValue {
            for i, c := range h.hand {
                valOne := cardValue[c]
                valTwo := cardValue[rune(hand.hand[i])]
                if valOne == valTwo {
                    continue
                } else if valOne > valTwo && h.ranking < hand.ranking{
                    h.ranking, hand.ranking = hand.ranking, h.ranking
                    break
                } else if valOne < valTwo && h.ranking > hand.ranking {
                    h.ranking, hand.ranking = hand.ranking, h.ranking
                    break
                }
            }
            continue
        }

        if h.handValue > hand.handValue && h.ranking < hand.ranking {
            h.ranking, hand.ranking = hand.ranking, h.ranking
        } else if h.handValue < hand.handValue && h.ranking > hand.ranking{
            h.ranking, hand.ranking = hand.ranking, h.ranking
        }

        if h.ranking == hand.ranking && h.handValue > hand.handValue{
            hand.ranking--
        }

        list[i] = hand
    }
    return list
}

func main() {
	var answer int64
    var bid int64
    var handRanking []handInfo 
	listOfHands := ReadInput(os.Args[1])

    answer += bid


	for i, line := range listOfHands {
        i++
        if line == "" {
            break
        }
		handBid := strings.Split(line, " ")
		hand := handBid[0]
        value := parseHandValue(hand)
		bid, _ = strconv.ParseInt(handBid[1], 10, 64)

        handRanking = append(handRanking, handInfo{
            hand: hand,
            handValue: value,
            bid: int(bid),
            ranking: i,
        })
	}

    copyRankings := handRanking
    for _, h := range handRanking {
        copyRankings = rankHand(&h, copyRankings)
    }
    
    fmt.Println(handRanking)
	fmt.Println(answer)
}
